/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumApplicationOIDCResponseType the model 'EnumApplicationOIDCResponseType'
type EnumApplicationOIDCResponseType string

// List of EnumApplicationOIDCResponseType
const (
	TOKEN EnumApplicationOIDCResponseType = "TOKEN"
	ID_TOKEN EnumApplicationOIDCResponseType = "ID_TOKEN"
	CODE EnumApplicationOIDCResponseType = "CODE"
)

// All allowed values of EnumApplicationOIDCResponseType enum
var AllowedEnumApplicationOIDCResponseTypeEnumValues = []EnumApplicationOIDCResponseType{
	"TOKEN",
	"ID_TOKEN",
	"CODE",
}

func (v *EnumApplicationOIDCResponseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationOIDCResponseType(value)
	for _, existing := range AllowedEnumApplicationOIDCResponseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumApplicationOIDCResponseType", value)
}

// NewEnumApplicationOIDCResponseTypeFromValue returns a pointer to a valid EnumApplicationOIDCResponseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationOIDCResponseTypeFromValue(v string) (*EnumApplicationOIDCResponseType, error) {
	ev := EnumApplicationOIDCResponseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationOIDCResponseType: valid values are %v", v, AllowedEnumApplicationOIDCResponseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationOIDCResponseType) IsValid() bool {
	for _, existing := range AllowedEnumApplicationOIDCResponseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationOIDCResponseType value
func (v EnumApplicationOIDCResponseType) Ptr() *EnumApplicationOIDCResponseType {
	return &v
}

type NullableEnumApplicationOIDCResponseType struct {
	value *EnumApplicationOIDCResponseType
	isSet bool
}

func (v NullableEnumApplicationOIDCResponseType) Get() *EnumApplicationOIDCResponseType {
	return v.value
}

func (v *NullableEnumApplicationOIDCResponseType) Set(val *EnumApplicationOIDCResponseType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationOIDCResponseType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationOIDCResponseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationOIDCResponseType(val *EnumApplicationOIDCResponseType) *NullableEnumApplicationOIDCResponseType {
	return &NullableEnumApplicationOIDCResponseType{value: val, isSet: true}
}

func (v NullableEnumApplicationOIDCResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationOIDCResponseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

