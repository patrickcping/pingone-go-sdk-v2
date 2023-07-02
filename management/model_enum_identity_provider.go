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

// EnumIdentityProvider A string that identifies the type of identity provider used to authenticate the user. The default value of `PING_ONE` is set when a value for identityProvider.id is not provided. The `PING_ONE` value is the default for all pre-existing users. There is currently no search filter support for this attribute.
type EnumIdentityProvider string

// List of EnumIdentityProvider
const (
	ENUMIDENTITYPROVIDER_FACEBOOK EnumIdentityProvider = "FACEBOOK"
	ENUMIDENTITYPROVIDER_GOOGLE EnumIdentityProvider = "GOOGLE"
	ENUMIDENTITYPROVIDER_LINKEDIN EnumIdentityProvider = "LINKEDIN"
	ENUMIDENTITYPROVIDER_APPLE EnumIdentityProvider = "APPLE"
	ENUMIDENTITYPROVIDER_TWITTER EnumIdentityProvider = "TWITTER"
	ENUMIDENTITYPROVIDER_AMAZON EnumIdentityProvider = "AMAZON"
	ENUMIDENTITYPROVIDER_YAHOO EnumIdentityProvider = "YAHOO"
	ENUMIDENTITYPROVIDER_MICROSOFT EnumIdentityProvider = "MICROSOFT"
	ENUMIDENTITYPROVIDER_PAYPAL EnumIdentityProvider = "PAYPAL"
	ENUMIDENTITYPROVIDER_GITHUB EnumIdentityProvider = "GITHUB"
	ENUMIDENTITYPROVIDER_OPENID_CONNECT EnumIdentityProvider = "OPENID_CONNECT"
	ENUMIDENTITYPROVIDER_SAML EnumIdentityProvider = "SAML"
	ENUMIDENTITYPROVIDER_PING_ONE EnumIdentityProvider = "PING_ONE"
)

// All allowed values of EnumIdentityProvider enum
var AllowedEnumIdentityProviderEnumValues = []EnumIdentityProvider{
	"FACEBOOK",
	"GOOGLE",
	"LINKEDIN",
	"APPLE",
	"TWITTER",
	"AMAZON",
	"YAHOO",
	"MICROSOFT",
	"PAYPAL",
	"GITHUB",
	"OPENID_CONNECT",
	"SAML",
	"PING_ONE",
}

func (v *EnumIdentityProvider) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIdentityProvider(value)
	for _, existing := range AllowedEnumIdentityProviderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumIdentityProvider(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumIdentityProviderFromValue returns a pointer to a valid EnumIdentityProvider
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIdentityProviderFromValue(v string) (*EnumIdentityProvider, error) {
	ev := EnumIdentityProvider(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIdentityProvider: valid values are %v", v, AllowedEnumIdentityProviderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIdentityProvider) IsValid() bool {
	for _, existing := range AllowedEnumIdentityProviderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIdentityProvider value
func (v EnumIdentityProvider) Ptr() *EnumIdentityProvider {
	return &v
}

type NullableEnumIdentityProvider struct {
	value *EnumIdentityProvider
	isSet bool
}

func (v NullableEnumIdentityProvider) Get() *EnumIdentityProvider {
	return v.value
}

func (v *NullableEnumIdentityProvider) Set(val *EnumIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIdentityProvider(val *EnumIdentityProvider) *NullableEnumIdentityProvider {
	return &NullableEnumIdentityProvider{value: val, isSet: true}
}

func (v NullableEnumIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

