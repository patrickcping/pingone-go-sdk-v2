/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumOrganizationType A string that specifies the organization type. If the organization has any paid licenses, the type property value is set to `PAID`. Otherwise, the property value is set to `TRIAL`.
type EnumOrganizationType string

// List of EnumOrganizationType
const (
	ENUMORGANIZATIONTYPE_PAID EnumOrganizationType = "PAID"
	ENUMORGANIZATIONTYPE_TRIAL EnumOrganizationType = "TRIAL"
)

// All allowed values of EnumOrganizationType enum
var AllowedEnumOrganizationTypeEnumValues = []EnumOrganizationType{
	"PAID",
	"TRIAL",
}

func (v *EnumOrganizationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumOrganizationType(value)
	for _, existing := range AllowedEnumOrganizationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumOrganizationType", value)
}

// NewEnumOrganizationTypeFromValue returns a pointer to a valid EnumOrganizationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumOrganizationTypeFromValue(v string) (*EnumOrganizationType, error) {
	ev := EnumOrganizationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumOrganizationType: valid values are %v", v, AllowedEnumOrganizationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumOrganizationType) IsValid() bool {
	for _, existing := range AllowedEnumOrganizationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumOrganizationType value
func (v EnumOrganizationType) Ptr() *EnumOrganizationType {
	return &v
}

type NullableEnumOrganizationType struct {
	value *EnumOrganizationType
	isSet bool
}

func (v NullableEnumOrganizationType) Get() *EnumOrganizationType {
	return v.value
}

func (v *NullableEnumOrganizationType) Set(val *EnumOrganizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumOrganizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumOrganizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumOrganizationType(val *EnumOrganizationType) *NullableEnumOrganizationType {
	return &NullableEnumOrganizationType{value: val, isSet: true}
}

func (v NullableEnumOrganizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumOrganizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

