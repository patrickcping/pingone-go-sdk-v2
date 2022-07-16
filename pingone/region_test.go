package pingone

import (
	"reflect"
	"testing"
)

func TestRegionToRegionSuffix(t *testing.T) {

	codeTests := map[string]string{
		"EU":             "eu",
		"US":             "com",
		"CA":             "ca",
		"ASIA":           "asia",
		"DOES_NOT_EXIST": "com",
		"":               "com",
	}

	for regionCode, suffixCode := range codeTests {

		v := regionToRegionSuffix(regionCode)
		if v == "" {
			t.Fatalf("regionToRegionSuffix returned with blank value but expected %s", suffixCode)
		} else {

			if v != suffixCode {
				t.Fatalf("regionToRegionSuffix resulted in %v, expected %s", v, suffixCode)
			}
		}
	}
}

func TestAvailableRegionsList(t *testing.T) {

	expectedList := []string{"ASIA", "CA", "EU", "US"}

	v := AvailableRegionsList()
	if !reflect.DeepEqual(v, expectedList) {
		t.Fatalf("AvailableRegionsList resulted in %v, expected %v", v, expectedList)
	}

}
