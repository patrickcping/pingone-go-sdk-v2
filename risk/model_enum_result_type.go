/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumResultType A string that specifies the risk evaluation result type. Options are `VALUE`.
type EnumResultType string

// List of EnumResultType
const (
	ENUMRESULTTYPE_VALUE EnumResultType = "VALUE"
)

// All allowed values of EnumResultType enum
var AllowedEnumResultTypeEnumValues = []EnumResultType{
	"VALUE",
}

func (v *EnumResultType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumResultType(value)
	for _, existing := range AllowedEnumResultTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumResultType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumResultTypeFromValue returns a pointer to a valid EnumResultType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumResultTypeFromValue(v string) (*EnumResultType, error) {
	ev := EnumResultType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumResultType: valid values are %v", v, AllowedEnumResultTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumResultType) IsValid() bool {
	for _, existing := range AllowedEnumResultTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumResultType value
func (v EnumResultType) Ptr() *EnumResultType {
	return &v
}

type NullableEnumResultType struct {
	value *EnumResultType
	isSet bool
}

func (v NullableEnumResultType) Get() *EnumResultType {
	return v.value
}

func (v *NullableEnumResultType) Set(val *EnumResultType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumResultType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumResultType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumResultType(val *EnumResultType) *NullableEnumResultType {
	return &NullableEnumResultType{value: val, isSet: true}
}

func (v NullableEnumResultType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumResultType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

