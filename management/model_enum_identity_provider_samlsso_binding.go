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

// EnumIdentityProviderSAMLSSOBinding A string that specifies the binding for the authentication request. Options are HTTP_POST and HTTP_REDIRECT.
type EnumIdentityProviderSAMLSSOBinding string

// List of EnumIdentityProviderSAMLSSOBinding
const (
	ENUMIDENTITYPROVIDERSAMLSSOBINDING_POST EnumIdentityProviderSAMLSSOBinding = "HTTP_POST"
	ENUMIDENTITYPROVIDERSAMLSSOBINDING_REDIRECT EnumIdentityProviderSAMLSSOBinding = "HTTP_REDIRECT"
)

// All allowed values of EnumIdentityProviderSAMLSSOBinding enum
var AllowedEnumIdentityProviderSAMLSSOBindingEnumValues = []EnumIdentityProviderSAMLSSOBinding{
	"HTTP_POST",
	"HTTP_REDIRECT",
}

func (v *EnumIdentityProviderSAMLSSOBinding) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIdentityProviderSAMLSSOBinding(value)
	for _, existing := range AllowedEnumIdentityProviderSAMLSSOBindingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumIdentityProviderSAMLSSOBinding", value)
}

// NewEnumIdentityProviderSAMLSSOBindingFromValue returns a pointer to a valid EnumIdentityProviderSAMLSSOBinding
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIdentityProviderSAMLSSOBindingFromValue(v string) (*EnumIdentityProviderSAMLSSOBinding, error) {
	ev := EnumIdentityProviderSAMLSSOBinding(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIdentityProviderSAMLSSOBinding: valid values are %v", v, AllowedEnumIdentityProviderSAMLSSOBindingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIdentityProviderSAMLSSOBinding) IsValid() bool {
	for _, existing := range AllowedEnumIdentityProviderSAMLSSOBindingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIdentityProviderSAMLSSOBinding value
func (v EnumIdentityProviderSAMLSSOBinding) Ptr() *EnumIdentityProviderSAMLSSOBinding {
	return &v
}

type NullableEnumIdentityProviderSAMLSSOBinding struct {
	value *EnumIdentityProviderSAMLSSOBinding
	isSet bool
}

func (v NullableEnumIdentityProviderSAMLSSOBinding) Get() *EnumIdentityProviderSAMLSSOBinding {
	return v.value
}

func (v *NullableEnumIdentityProviderSAMLSSOBinding) Set(val *EnumIdentityProviderSAMLSSOBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIdentityProviderSAMLSSOBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIdentityProviderSAMLSSOBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIdentityProviderSAMLSSOBinding(val *EnumIdentityProviderSAMLSSOBinding) *NullableEnumIdentityProviderSAMLSSOBinding {
	return &NullableEnumIdentityProviderSAMLSSOBinding{value: val, isSet: true}
}

func (v NullableEnumIdentityProviderSAMLSSOBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIdentityProviderSAMLSSOBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

