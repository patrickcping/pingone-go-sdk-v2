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

// EnumIdentityProviderPKCEMethod The method for PKCE. Options are `NONE` or `S256`. The default is `NONE`. This value auto-populates from a discovery endpoint if the OpenID Provider includes `S256` in its `code_challenge_methods_supported` claim. The plain method is not currently supported.
type EnumIdentityProviderPKCEMethod string

// List of EnumIdentityProviderPKCEMethod
const (
	ENUMIDENTITYPROVIDERPKCEMETHOD_S256 EnumIdentityProviderPKCEMethod = "S256"
	ENUMIDENTITYPROVIDERPKCEMETHOD_NONE EnumIdentityProviderPKCEMethod = "NONE"
)

// All allowed values of EnumIdentityProviderPKCEMethod enum
var AllowedEnumIdentityProviderPKCEMethodEnumValues = []EnumIdentityProviderPKCEMethod{
	"S256",
	"NONE",
}

func (v *EnumIdentityProviderPKCEMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIdentityProviderPKCEMethod(value)
	for _, existing := range AllowedEnumIdentityProviderPKCEMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumIdentityProviderPKCEMethod(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumIdentityProviderPKCEMethodFromValue returns a pointer to a valid EnumIdentityProviderPKCEMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIdentityProviderPKCEMethodFromValue(v string) (*EnumIdentityProviderPKCEMethod, error) {
	ev := EnumIdentityProviderPKCEMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIdentityProviderPKCEMethod: valid values are %v", v, AllowedEnumIdentityProviderPKCEMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIdentityProviderPKCEMethod) IsValid() bool {
	for _, existing := range AllowedEnumIdentityProviderPKCEMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIdentityProviderPKCEMethod value
func (v EnumIdentityProviderPKCEMethod) Ptr() *EnumIdentityProviderPKCEMethod {
	return &v
}

type NullableEnumIdentityProviderPKCEMethod struct {
	value *EnumIdentityProviderPKCEMethod
	isSet bool
}

func (v NullableEnumIdentityProviderPKCEMethod) Get() *EnumIdentityProviderPKCEMethod {
	return v.value
}

func (v *NullableEnumIdentityProviderPKCEMethod) Set(val *EnumIdentityProviderPKCEMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIdentityProviderPKCEMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIdentityProviderPKCEMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIdentityProviderPKCEMethod(val *EnumIdentityProviderPKCEMethod) *NullableEnumIdentityProviderPKCEMethod {
	return &NullableEnumIdentityProviderPKCEMethod{value: val, isSet: true}
}

func (v NullableEnumIdentityProviderPKCEMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIdentityProviderPKCEMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
