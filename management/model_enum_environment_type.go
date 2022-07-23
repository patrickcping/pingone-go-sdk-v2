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

// EnumEnvironmentType A string that specifies the type of environment to use.
type EnumEnvironmentType string

// List of EnumEnvironmentType
const (
	ENUMENVIRONMENTTYPE_PRODUCTION EnumEnvironmentType = "PRODUCTION"
	ENUMENVIRONMENTTYPE_SANDBOX EnumEnvironmentType = "SANDBOX"
)

// All allowed values of EnumEnvironmentType enum
var AllowedEnumEnvironmentTypeEnumValues = []EnumEnvironmentType{
	"PRODUCTION",
	"SANDBOX",
}

func (v *EnumEnvironmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumEnvironmentType(value)
	for _, existing := range AllowedEnumEnvironmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumEnvironmentType", value)
}

// NewEnumEnvironmentTypeFromValue returns a pointer to a valid EnumEnvironmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumEnvironmentTypeFromValue(v string) (*EnumEnvironmentType, error) {
	ev := EnumEnvironmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumEnvironmentType: valid values are %v", v, AllowedEnumEnvironmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumEnvironmentType) IsValid() bool {
	for _, existing := range AllowedEnumEnvironmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumEnvironmentType value
func (v EnumEnvironmentType) Ptr() *EnumEnvironmentType {
	return &v
}

type NullableEnumEnvironmentType struct {
	value *EnumEnvironmentType
	isSet bool
}

func (v NullableEnumEnvironmentType) Get() *EnumEnvironmentType {
	return v.value
}

func (v *NullableEnumEnvironmentType) Set(val *EnumEnvironmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumEnvironmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumEnvironmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumEnvironmentType(val *EnumEnvironmentType) *NullableEnumEnvironmentType {
	return &NullableEnumEnvironmentType{value: val, isSet: true}
}

func (v NullableEnumEnvironmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumEnvironmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

