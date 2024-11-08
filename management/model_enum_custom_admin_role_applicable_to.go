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

// EnumCustomAdminRoleApplicableTo The scope types to which the role can be applied. Options are ORGANIZATION, ENVIRONMENT, and POPULATION.
type EnumCustomAdminRoleApplicableTo string

// List of EnumCustomAdminRoleApplicableTo
const (
	ENUMCUSTOMADMINROLEAPPLICABLETO_ORGANIZATION EnumCustomAdminRoleApplicableTo = "ORGANIZATION"
	ENUMCUSTOMADMINROLEAPPLICABLETO_ENVIRONMENT EnumCustomAdminRoleApplicableTo = "ENVIRONMENT"
	ENUMCUSTOMADMINROLEAPPLICABLETO_POPULATION EnumCustomAdminRoleApplicableTo = "POPULATION"
)

// All allowed values of EnumCustomAdminRoleApplicableTo enum
var AllowedEnumCustomAdminRoleApplicableToEnumValues = []EnumCustomAdminRoleApplicableTo{
	"ORGANIZATION",
	"ENVIRONMENT",
	"POPULATION",
}

func (v *EnumCustomAdminRoleApplicableTo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCustomAdminRoleApplicableTo(value)
	for _, existing := range AllowedEnumCustomAdminRoleApplicableToEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumCustomAdminRoleApplicableTo(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumCustomAdminRoleApplicableToFromValue returns a pointer to a valid EnumCustomAdminRoleApplicableTo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCustomAdminRoleApplicableToFromValue(v string) (*EnumCustomAdminRoleApplicableTo, error) {
	ev := EnumCustomAdminRoleApplicableTo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCustomAdminRoleApplicableTo: valid values are %v", v, AllowedEnumCustomAdminRoleApplicableToEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCustomAdminRoleApplicableTo) IsValid() bool {
	for _, existing := range AllowedEnumCustomAdminRoleApplicableToEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCustomAdminRoleApplicableTo value
func (v EnumCustomAdminRoleApplicableTo) Ptr() *EnumCustomAdminRoleApplicableTo {
	return &v
}

type NullableEnumCustomAdminRoleApplicableTo struct {
	value *EnumCustomAdminRoleApplicableTo
	isSet bool
}

func (v NullableEnumCustomAdminRoleApplicableTo) Get() *EnumCustomAdminRoleApplicableTo {
	return v.value
}

func (v *NullableEnumCustomAdminRoleApplicableTo) Set(val *EnumCustomAdminRoleApplicableTo) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCustomAdminRoleApplicableTo) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCustomAdminRoleApplicableTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCustomAdminRoleApplicableTo(val *EnumCustomAdminRoleApplicableTo) *NullableEnumCustomAdminRoleApplicableTo {
	return &NullableEnumCustomAdminRoleApplicableTo{value: val, isSet: true}
}

func (v NullableEnumCustomAdminRoleApplicableTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCustomAdminRoleApplicableTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
