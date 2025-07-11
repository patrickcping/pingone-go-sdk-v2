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

// EnumIntegrationVersionType The type of integration for this version. Currently, the only valid values are PRODUCT_INTEGRATION_KIT and SAML.
type EnumIntegrationVersionType string

// List of EnumIntegrationVersionType
const (
	ENUMINTEGRATIONVERSIONTYPE_PRODUCT_INTEGRATION_KIT EnumIntegrationVersionType = "PRODUCT_INTEGRATION_KIT"
	ENUMINTEGRATIONVERSIONTYPE_SAML                    EnumIntegrationVersionType = "SAML"
)

// All allowed values of EnumIntegrationVersionType enum
var AllowedEnumIntegrationVersionTypeEnumValues = []EnumIntegrationVersionType{
	"PRODUCT_INTEGRATION_KIT",
	"SAML",
}

func (v *EnumIntegrationVersionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIntegrationVersionType(value)
	for _, existing := range AllowedEnumIntegrationVersionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumIntegrationVersionType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumIntegrationVersionTypeFromValue returns a pointer to a valid EnumIntegrationVersionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIntegrationVersionTypeFromValue(v string) (*EnumIntegrationVersionType, error) {
	ev := EnumIntegrationVersionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIntegrationVersionType: valid values are %v", v, AllowedEnumIntegrationVersionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIntegrationVersionType) IsValid() bool {
	for _, existing := range AllowedEnumIntegrationVersionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIntegrationVersionType value
func (v EnumIntegrationVersionType) Ptr() *EnumIntegrationVersionType {
	return &v
}

type NullableEnumIntegrationVersionType struct {
	value *EnumIntegrationVersionType
	isSet bool
}

func (v NullableEnumIntegrationVersionType) Get() *EnumIntegrationVersionType {
	return v.value
}

func (v *NullableEnumIntegrationVersionType) Set(val *EnumIntegrationVersionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIntegrationVersionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIntegrationVersionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIntegrationVersionType(val *EnumIntegrationVersionType) *NullableEnumIntegrationVersionType {
	return &NullableEnumIntegrationVersionType{value: val, isSet: true}
}

func (v NullableEnumIntegrationVersionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIntegrationVersionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
