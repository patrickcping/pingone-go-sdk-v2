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

// EnumAttributeMappingType A string that specifies the mapping type of the attribute. Options are CORE, SCOPE, and CUSTOM. The CORE and SCOPE mapping types are for reserved attributes managed by the API and cannot be removed. Attribute values for these mapping types can be updated. The CUSTOM mapping type is for user-defined attributes. Attributes of this type can be updated and deleted.
type EnumAttributeMappingType string

// List of EnumAttributeMappingType
const (
	ENUMATTRIBUTEMAPPINGTYPE_CORE EnumAttributeMappingType = "CORE"
	ENUMATTRIBUTEMAPPINGTYPE_SCOPE EnumAttributeMappingType = "SCOPE"
	ENUMATTRIBUTEMAPPINGTYPE_CUSTOM EnumAttributeMappingType = "CUSTOM"
)

// All allowed values of EnumAttributeMappingType enum
var AllowedEnumAttributeMappingTypeEnumValues = []EnumAttributeMappingType{
	"CORE",
	"SCOPE",
	"CUSTOM",
}

func (v *EnumAttributeMappingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAttributeMappingType(value)
	for _, existing := range AllowedEnumAttributeMappingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumAttributeMappingType", value)
}

// NewEnumAttributeMappingTypeFromValue returns a pointer to a valid EnumAttributeMappingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAttributeMappingTypeFromValue(v string) (*EnumAttributeMappingType, error) {
	ev := EnumAttributeMappingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAttributeMappingType: valid values are %v", v, AllowedEnumAttributeMappingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAttributeMappingType) IsValid() bool {
	for _, existing := range AllowedEnumAttributeMappingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAttributeMappingType value
func (v EnumAttributeMappingType) Ptr() *EnumAttributeMappingType {
	return &v
}

type NullableEnumAttributeMappingType struct {
	value *EnumAttributeMappingType
	isSet bool
}

func (v NullableEnumAttributeMappingType) Get() *EnumAttributeMappingType {
	return v.value
}

func (v *NullableEnumAttributeMappingType) Set(val *EnumAttributeMappingType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAttributeMappingType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAttributeMappingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAttributeMappingType(val *EnumAttributeMappingType) *NullableEnumAttributeMappingType {
	return &NullableEnumAttributeMappingType{value: val, isSet: true}
}

func (v NullableEnumAttributeMappingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAttributeMappingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

