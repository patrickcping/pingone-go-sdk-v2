/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// EnumApplicationResourceParentType The application resource's parent type. Options are PING_ONE_RESOURCE.
type EnumApplicationResourceParentType string

// List of EnumApplicationResourceParentType
const (
	ENUMAPPLICATIONRESOURCEPARENTTYPE_PING_ONE_RESOURCE EnumApplicationResourceParentType = "PING_ONE_RESOURCE"
)

// All allowed values of EnumApplicationResourceParentType enum
var AllowedEnumApplicationResourceParentTypeEnumValues = []EnumApplicationResourceParentType{
	"PING_ONE_RESOURCE",
}

func (v *EnumApplicationResourceParentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationResourceParentType(value)
	for _, existing := range AllowedEnumApplicationResourceParentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumApplicationResourceParentType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumApplicationResourceParentTypeFromValue returns a pointer to a valid EnumApplicationResourceParentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationResourceParentTypeFromValue(v string) (*EnumApplicationResourceParentType, error) {
	ev := EnumApplicationResourceParentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationResourceParentType: valid values are %v", v, AllowedEnumApplicationResourceParentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationResourceParentType) IsValid() bool {
	for _, existing := range AllowedEnumApplicationResourceParentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationResourceParentType value
func (v EnumApplicationResourceParentType) Ptr() *EnumApplicationResourceParentType {
	return &v
}

type NullableEnumApplicationResourceParentType struct {
	value *EnumApplicationResourceParentType
	isSet bool
}

func (v NullableEnumApplicationResourceParentType) Get() *EnumApplicationResourceParentType {
	return v.value
}

func (v *NullableEnumApplicationResourceParentType) Set(val *EnumApplicationResourceParentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationResourceParentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationResourceParentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationResourceParentType(val *EnumApplicationResourceParentType) *NullableEnumApplicationResourceParentType {
	return &NullableEnumApplicationResourceParentType{value: val, isSet: true}
}

func (v NullableEnumApplicationResourceParentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationResourceParentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
