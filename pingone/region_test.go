package pingone

import (
	"reflect"
	"testing"
)

func TestFindRegionByName(t *testing.T) {

	codeTests := map[string]RegionMapping{
		"Europe": {
			Region:    "Europe",
			URLSuffix: "eu",
			APICode:   "EU",
		},
		"NorthAmerica": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
		"Canada": {
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   "CA",
		},
		"AsiaPacific": {
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   "AP",
		},
		"DOES_NOT_EXIST": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
		"": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
	}

	for regionCode, regionMap := range codeTests {

		v := FindRegionByName(regionCode)

		if v != regionMap {
			t.Fatalf("TestFindRegionByName resulted in %v, expected %v", v, regionMap)
		}

	}
}

func TestFindRegionByAPICode(t *testing.T) {

	codeTests := map[string]RegionMapping{
		"EU": {
			Region:    "Europe",
			URLSuffix: "eu",
			APICode:   "EU",
		},
		"NA": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
		"CA": {
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   "CA",
		},
		"AP": {
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   "AP",
		},
		"DOES_NOT_EXIST": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
		"": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   "NA",
			Default:   true,
		},
	}

	for regionCode, regionMap := range codeTests {

		v := FindRegionByAPICode(regionCode)

		if v != regionMap {
			t.Fatalf("TestFindRegionByAPICode resulted in %v, expected %v", v, regionMap)
		}

	}
}

func TestAvailableRegionsList(t *testing.T) {

	expectedList := []string{"AsiaPacific", "Canada", "Europe", "NorthAmerica"}

	v := AvailableRegionsList()
	if !reflect.DeepEqual(v, expectedList) {
		t.Fatalf("AvailableRegionsList resulted in %v, expected %v", v, expectedList)
	}

}
