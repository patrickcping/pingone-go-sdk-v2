package model

import (
	"reflect"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

func TestFindRegionByName(t *testing.T) {

	codeTests := map[string]RegionMapping{
		"Europe": {
			Region:    "Europe",
			URLSuffix: "eu",
			APICode:   management.ENUMREGIONCODE_EU,
		},
		"NorthAmerica": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
			Default:   true,
		},
		"Canada": {
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   management.ENUMREGIONCODE_CA,
		},
		"AsiaPacific": {
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   management.ENUMREGIONCODE_AP,
		},
		"DOES_NOT_EXIST": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
			Default:   true,
		},
		"": {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
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

	codeTests := map[management.EnumRegionCode]RegionMapping{
		management.ENUMREGIONCODE_EU: {
			Region:    "Europe",
			URLSuffix: "eu",
			APICode:   management.ENUMREGIONCODE_EU,
		},
		management.ENUMREGIONCODE_NA: {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
			Default:   true,
		},
		management.ENUMREGIONCODE_CA: {
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   management.ENUMREGIONCODE_CA,
		},
		management.ENUMREGIONCODE_AP: {
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   management.ENUMREGIONCODE_AP,
		},
		management.ENUMREGIONCODE_AU: {
			Region:    "Australia-AsiaPacific",
			URLSuffix: "com.au",
			APICode:   management.ENUMREGIONCODE_AU,
		},
		management.EnumRegionCode("DOES_NOT_EXIST"): {
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
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

	expectedList := []string{"AsiaPacific", "Australia-AsiaPacific", "Canada", "Europe", "NorthAmerica"}

	v := RegionsAvailableList()
	if !reflect.DeepEqual(v, expectedList) {
		t.Fatalf("AvailableRegionsList resulted in %v, expected %v", v, expectedList)
	}

}
