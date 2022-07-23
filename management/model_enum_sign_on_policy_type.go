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

// EnumSignOnPolicyType A string that specifies the type of action. Options are `LOGIN`, `MULTI_FACTOR_AUTHENTICATION`, `IDENTIFIER_FIRST`, `IDENTITY_PROVIDER` `AGREEMENT` and `PROGRESSIVE_PROFILING`.
type EnumSignOnPolicyType string

// List of EnumSignOnPolicyType
const (
	ENUMSIGNONPOLICYTYPE_LOGIN EnumSignOnPolicyType = "LOGIN"
	ENUMSIGNONPOLICYTYPE_MULTI_FACTOR_AUTHENTICATION EnumSignOnPolicyType = "MULTI_FACTOR_AUTHENTICATION"
	ENUMSIGNONPOLICYTYPE_IDENTIFIER_FIRST EnumSignOnPolicyType = "IDENTIFIER_FIRST"
	ENUMSIGNONPOLICYTYPE_IDENTITY_PROVIDER EnumSignOnPolicyType = "IDENTITY_PROVIDER"
	ENUMSIGNONPOLICYTYPE_PROGRESSIVE_PROFILING EnumSignOnPolicyType = "PROGRESSIVE_PROFILING"
	ENUMSIGNONPOLICYTYPE_AGREEMENT EnumSignOnPolicyType = "AGREEMENT"
)

// All allowed values of EnumSignOnPolicyType enum
var AllowedEnumSignOnPolicyTypeEnumValues = []EnumSignOnPolicyType{
	"LOGIN",
	"MULTI_FACTOR_AUTHENTICATION",
	"IDENTIFIER_FIRST",
	"IDENTITY_PROVIDER",
	"PROGRESSIVE_PROFILING",
	"AGREEMENT",
}

func (v *EnumSignOnPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSignOnPolicyType(value)
	for _, existing := range AllowedEnumSignOnPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumSignOnPolicyType", value)
}

// NewEnumSignOnPolicyTypeFromValue returns a pointer to a valid EnumSignOnPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSignOnPolicyTypeFromValue(v string) (*EnumSignOnPolicyType, error) {
	ev := EnumSignOnPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSignOnPolicyType: valid values are %v", v, AllowedEnumSignOnPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSignOnPolicyType) IsValid() bool {
	for _, existing := range AllowedEnumSignOnPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSignOnPolicyType value
func (v EnumSignOnPolicyType) Ptr() *EnumSignOnPolicyType {
	return &v
}

type NullableEnumSignOnPolicyType struct {
	value *EnumSignOnPolicyType
	isSet bool
}

func (v NullableEnumSignOnPolicyType) Get() *EnumSignOnPolicyType {
	return v.value
}

func (v *NullableEnumSignOnPolicyType) Set(val *EnumSignOnPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSignOnPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSignOnPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSignOnPolicyType(val *EnumSignOnPolicyType) *NullableEnumSignOnPolicyType {
	return &NullableEnumSignOnPolicyType{value: val, isSet: true}
}

func (v NullableEnumSignOnPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSignOnPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

