/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumIntegrationPingProductName The Ping product associated with the integration.
type EnumIntegrationPingProductName string

// List of EnumIntegrationPingProductName
const (
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGID EnumIntegrationPingProductName = "PINGID"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGONE_ENTERPRISE EnumIntegrationPingProductName = "PINGONE_ENTERPRISE"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGONE EnumIntegrationPingProductName = "PINGONE"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGACCESS EnumIntegrationPingProductName = "PINGACCESS"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGFEDERATE EnumIntegrationPingProductName = "PINGFEDERATE"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGDIRECTORY EnumIntegrationPingProductName = "PINGDIRECTORY"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGDATAGOVERNANCE EnumIntegrationPingProductName = "PINGDATAGOVERNANCE"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGINTELLIGENCE_FOR_APIS EnumIntegrationPingProductName = "PINGINTELLIGENCE_FOR_APIS"
	ENUMINTEGRATIONPINGPRODUCTNAME_PINGONE_ADVANCED_SERVICES EnumIntegrationPingProductName = "PINGONE_ADVANCED_SERVICES"
)

// All allowed values of EnumIntegrationPingProductName enum
var AllowedEnumIntegrationPingProductNameEnumValues = []EnumIntegrationPingProductName{
	"PINGID",
	"PINGONE_ENTERPRISE",
	"PINGONE",
	"PINGACCESS",
	"PINGFEDERATE",
	"PINGDIRECTORY",
	"PINGDATAGOVERNANCE",
	"PINGINTELLIGENCE_FOR_APIS",
	"PINGONE_ADVANCED_SERVICES",
}

func (v *EnumIntegrationPingProductName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIntegrationPingProductName(value)
	for _, existing := range AllowedEnumIntegrationPingProductNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumIntegrationPingProductName(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumIntegrationPingProductNameFromValue returns a pointer to a valid EnumIntegrationPingProductName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIntegrationPingProductNameFromValue(v string) (*EnumIntegrationPingProductName, error) {
	ev := EnumIntegrationPingProductName(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIntegrationPingProductName: valid values are %v", v, AllowedEnumIntegrationPingProductNameEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIntegrationPingProductName) IsValid() bool {
	for _, existing := range AllowedEnumIntegrationPingProductNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIntegrationPingProductName value
func (v EnumIntegrationPingProductName) Ptr() *EnumIntegrationPingProductName {
	return &v
}

type NullableEnumIntegrationPingProductName struct {
	value *EnumIntegrationPingProductName
	isSet bool
}

func (v NullableEnumIntegrationPingProductName) Get() *EnumIntegrationPingProductName {
	return v.value
}

func (v *NullableEnumIntegrationPingProductName) Set(val *EnumIntegrationPingProductName) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIntegrationPingProductName) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIntegrationPingProductName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIntegrationPingProductName(val *EnumIntegrationPingProductName) *NullableEnumIntegrationPingProductName {
	return &NullableEnumIntegrationPingProductName{value: val, isSet: true}
}

func (v NullableEnumIntegrationPingProductName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIntegrationPingProductName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
