package model

import (
	"reflect"
	"testing"

	"github.com/patrickcping/pingone-go-sdk-v2/management"
)

func TestFindProductByName_Success(t *testing.T) {

	codeTests := map[string]ProductMapping{
		"SSO": {
			APICode:     management.ENUMPRODUCTTYPE_ONE_BASE,
			ProductCode: "SSO",
			Default:     true,
			Selectable:  true,
		},
		"Provisioning": {
			APICode:     management.ENUMPRODUCTTYPE_ONE_PROVISIONING,
			ProductCode: "Provisioning",
		},
		"DaVinci": {
			APICode:     management.ENUMPRODUCTTYPE_ONE_DAVINCI,
			ProductCode: "DaVinci",
			Selectable:  true,
		},
	}

	for productCode, productMap := range codeTests {

		v, err := FindProductByName(productCode)

		if err != nil {
			t.Fatalf("TestTestFindProductByName resulted in %v", err)
		}

		if v != productMap {
			t.Fatalf("TestTestFindProductByName resulted in %v, expected %v", v, productMap)
		}

	}
}

func TestFindProductByName_Error(t *testing.T) {

	codeTests := []string{
		"DOESNOTEXIST", "",
	}

	for _, productCode := range codeTests {

		v, err := FindProductByName(productCode)

		if err == nil {
			t.Fatalf("TestTestFindProductByName resulted in %v, expected error", v)
		}

	}
}

func TestFindProductByAPICode_Success(t *testing.T) {

	codeTests := map[management.EnumProductType]ProductMapping{
		management.ENUMPRODUCTTYPE_ONE_BASE: {
			APICode:     management.ENUMPRODUCTTYPE_ONE_BASE,
			ProductCode: "SSO",
			Default:     true,
			Selectable:  true,
		},
		management.ENUMPRODUCTTYPE_ONE_PROVISIONING: {
			APICode:     management.ENUMPRODUCTTYPE_ONE_PROVISIONING,
			ProductCode: "Provisioning",
		},
		management.ENUMPRODUCTTYPE_ONE_DAVINCI: {
			APICode:     management.ENUMPRODUCTTYPE_ONE_DAVINCI,
			ProductCode: "DaVinci",
			Selectable:  true,
		},
	}

	for productCode, productMap := range codeTests {

		v, err := FindProductByAPICode(productCode)

		if err != nil {
			t.Fatalf("TestFindProductByAPICode resulted in %v", err)
		}

		if v != productMap {
			t.Fatalf("TestFindProductByAPICode resulted in %v, expected %v", v, productMap)
		}

	}
}

func TestProductsSelectableList(t *testing.T) {

	expectedList := []string{"APIIntelligence", "Authorize", "Credentials", "DaVinci", "MFA", "PingAccess", "PingAuthorize", "PingCentral", "PingDirectory", "PingFederate", "PingID", "PingID-v2", "Risk", "SSO", "Verify"}

	v := ProductsSelectableList()
	if !reflect.DeepEqual(v, expectedList) {
		t.Fatalf("ProductsSelectableList resulted in %v, expected %v", v, expectedList)
	}

}
