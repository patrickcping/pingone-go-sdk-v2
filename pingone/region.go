package pingone

import (
	"log"
	"sort"

	"golang.org/x/exp/slices"
)

type RegionMapping struct {
	Region       string
	RegionSuffix string
	Default      bool
}

var regionMappingList []RegionMapping

func init() {
	regionMappingList = []RegionMapping{
		{
			Region:       "EU",
			RegionSuffix: "eu",
		},
		{
			Region:       "US",
			RegionSuffix: "com",
			Default:      true,
		},
		{
			Region:       "ASIA",
			RegionSuffix: "asia",
		},
		{
			Region:       "CA",
			RegionSuffix: "ca",
		},
	}
}

func regionToRegionSuffix(region string) string {

	idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Region == region })

	if idx < 0 {

		idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Default })
		log.Printf("[INFO] Region code %s not found, defaulting to %s", region, regionMappingList[idx].Region)

		return regionMappingList[idx].RegionSuffix
	}

	return regionMappingList[idx].RegionSuffix

}

func AvailableRegionsList() []string {

	regionList := make([]string, 0)

	for _, region := range regionMappingList {

		regionList = append(regionList, region.Region)

	}

	sort.Strings(regionList)

	return regionList
}
