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

// EnumUserRequestInclude the model 'EnumUserRequestInclude'
type EnumUserRequestInclude string

// List of EnumUserRequestInclude
const (
	ENUMUSERREQUESTINCLUDE_MEMBER_OF_GROUP_IDS EnumUserRequestInclude = "memberOfGroupIDs"
	ENUMUSERREQUESTINCLUDE_MEMBER_OF_GROUP_NAMES EnumUserRequestInclude = "memberOfGroupNames"
)

// All allowed values of EnumUserRequestInclude enum
var AllowedEnumUserRequestIncludeEnumValues = []EnumUserRequestInclude{
	"memberOfGroupIDs",
	"memberOfGroupNames",
}

func (v *EnumUserRequestInclude) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserRequestInclude(value)
	for _, existing := range AllowedEnumUserRequestIncludeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserRequestInclude(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserRequestIncludeFromValue returns a pointer to a valid EnumUserRequestInclude
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserRequestIncludeFromValue(v string) (*EnumUserRequestInclude, error) {
	ev := EnumUserRequestInclude(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserRequestInclude: valid values are %v", v, AllowedEnumUserRequestIncludeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserRequestInclude) IsValid() bool {
	for _, existing := range AllowedEnumUserRequestIncludeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserRequestInclude value
func (v EnumUserRequestInclude) Ptr() *EnumUserRequestInclude {
	return &v
}

type NullableEnumUserRequestInclude struct {
	value *EnumUserRequestInclude
	isSet bool
}

func (v NullableEnumUserRequestInclude) Get() *EnumUserRequestInclude {
	return v.value
}

func (v *NullableEnumUserRequestInclude) Set(val *EnumUserRequestInclude) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserRequestInclude) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserRequestInclude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserRequestInclude(val *EnumUserRequestInclude) *NullableEnumUserRequestInclude {
	return &NullableEnumUserRequestInclude{value: val, isSet: true}
}

func (v NullableEnumUserRequestInclude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserRequestInclude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

