package pingone

import (
	"log"
	"sort"

	"golang.org/x/exp/slices"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

type RegionMapping struct {
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
			APICode:   management.EnumRegionCode("EU"),
		},
		{
			Region:    "NorthAmerica",
			URLSuffix: "com",
			APICode:   management.EnumRegionCode("NA"),
			Default:   true,
		},
		{
			Region:    "AsiaPacific",
			URLSuffix: "asia",
			APICode:   management.EnumRegionCode("AP"),
		},
		{
			Region:    "Canada",
			URLSuffix: "ca",
			APICode:   management.EnumRegionCode("CA"),
		},
	}
}

func FindRegionByName(region string) RegionMapping {

	idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Region == region })

	if idx < 0 {

		idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Default })
		log.Printf("[INFO] Region code %s not found, defaulting to %s", region, regionMappingList[idx].Region)

		return regionMappingList[idx]
	}

	return regionMappingList[idx]

}

func FindRegionByAPICode(apiCode string) RegionMapping {

	idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.APICode == management.EnumRegionCode(apiCode) })

	if idx < 0 {

		idx := slices.IndexFunc(regionMappingList, func(c RegionMapping) bool { return c.Default })
		log.Printf("[INFO] API code %s not found, defaulting to %s", apiCode, regionMappingList[idx].APICode)

		return regionMappingList[idx]
	}

	return regionMappingList[idx]

}

func AvailableRegionsList() []string {

	regionList := make([]string, 0)

	for _, region := range regionMappingList {

		regionList = append(regionList, region.Region)

	}

	sort.Strings(regionList)

	return regionList
}
