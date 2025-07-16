package model

import (
	"log"
	"slices"
	"sort"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

type RegionMapping struct {
	// Deprecated: Use APICode instead
	Region    string
	URLSuffix string
	APICode   management.EnumRegionCode
	Default   bool
}

var regionMappingList []RegionMapping

func init() {
	regionMappingList = []RegionMapping{
		{
			Region:    "Europe",
			URLSuffix: "eu",
			APICode:   management.ENUMREGIONCODE_EU,
		},
		{
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.ENUMREGIONCODE_NA,
			Default:   true,
		},
		{
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   management.ENUMREGIONCODE_AP,
		},
		{
			Region:    "Australia-AsiaPacific",
			URLSuffix: "com.au",
			APICode:   management.ENUMREGIONCODE_AU,
		},
		{
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   management.ENUMREGIONCODE_CA,
		},
		{
			Region:    "Singapore",
			URLSuffix: "sg",
			APICode:   management.ENUMREGIONCODE_SG,
		},
	}
}

// Deprecated: Use FindRegionByAPICode instead
func FindRegionByName(region string) RegionMapping {

	idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Region == region })

	if idx < 0 {

		idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Default })
		log.Printf("[INFO] Region code %s not found, defaulting to %s", region, regionMappingList[idx].Region)

		return regionMappingList[idx]
	}

	return regionMappingList[idx]

}

func FindRegionByAPICode(apiCode management.EnumRegionCode) RegionMapping {

	idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.APICode == apiCode })

	if idx < 0 {

		idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Default })
		log.Printf("[INFO] API code %s not found, defaulting to %s", apiCode, regionMappingList[idx].APICode)

		return regionMappingList[idx]
	}

	return regionMappingList[idx]

}

func RegionsAvailableList() []string {

	regionList := make([]string, 0)

	for _, region := range regionMappingList {

		regionList = append(regionList, region.Region)

	}

	sort.Strings(regionList)

	return regionList
}
