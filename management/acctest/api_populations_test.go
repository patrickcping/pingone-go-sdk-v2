package management_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"slices"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/acctest"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

func TestAccPopulationsAPI_PagedResults(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	client, err := acctest.TestClient(ctx)
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	environment := management.NewEnvironment(
		*management.NewEnvironmentLicense(os.Getenv("PINGONE_LICENSE_ID")),
		"go-sdk-test-populations-api-paged_results",
		management.EnvironmentRegion{
			EnumRegionCode: &client.Region.APICode,
		},
		management.ENUMENVIRONMENTTYPE_SANDBOX,
	)

	environmentResponse, _, err := client.ManagementAPIClient.EnvironmentsApi.CreateEnvironmentActiveLicense(ctx).Environment(*environment).Execute()
	if err != nil {
		t.Fatalf("Error creating environment: %v", err)
	}

	createdPopulationIDs := make([]string, 0)

	for i := 1; i <= 100; i++ {

		populationName := fmt.Sprintf("Test Population %d", i)
		population := management.NewPopulation(populationName)

		populationResponse, _, err := client.ManagementAPIClient.PopulationsApi.CreatePopulation(ctx, environmentResponse.GetId()).Population(*population).Execute()
		if err != nil {
			t.Logf("Error creating population %s: %v", populationName, err)
			t.Fail()
		}

		createdPopulationIDs = append(createdPopulationIDs, populationResponse.GetId())
	}

	slices.Sort(createdPopulationIDs)

	// Run the test
	readPopulationIDs := make([]string, 0)
	pagedIterator := client.ManagementAPIClient.PopulationsApi.ReadAllPopulations(ctx, environmentResponse.GetId()).Limit(20).Execute()

	for populationsPageCursor, err := range pagedIterator {
		log.Printf("Main loop")
		if err != nil {
			t.Logf("Error reading populations: %v", err)
			t.Fail()
			break
		}

		log.Printf("Nil checks")
		if populationsPageCursor.EntityArray == nil {
			t.Logf("EntityArray not returned")
			t.Fail()
			break
		}

		if populationsPageCursor.HTTPResponse == nil {
			t.Logf("HTTPResponse not returned")
			t.Fail()
			break
		}

		log.Printf("Add populations")
		for _, population := range populationsPageCursor.EntityArray.Embedded.GetPopulations() {
			log.Printf(" - Adding population: %s", population.GetName())
			readPopulationIDs = append(readPopulationIDs, population.GetId())
		}
	}

	slices.Sort(readPopulationIDs)

	if !slices.Equal(createdPopulationIDs, readPopulationIDs) {
		t.Logf("Expected Count %d, Actual Count %d", len(createdPopulationIDs), len(readPopulationIDs))
		t.Logf("Expected: %v", createdPopulationIDs)
		t.Logf("Actual: %v", readPopulationIDs)
		t.Fail()
	}

	// Cleanup
	for _, populationID := range createdPopulationIDs {
		_, err := client.ManagementAPIClient.PopulationsApi.DeletePopulation(ctx, environmentResponse.GetId(), populationID).Execute()
		if err != nil {
			t.Logf("Error deleting population %s: %v", populationID, err)
		}
	}

	_, err = client.ManagementAPIClient.EnvironmentsApi.DeleteEnvironment(ctx, environmentResponse.GetId()).Execute()
	if err != nil {
		t.Logf("Error deleting environment: %v", err)
	}
}
