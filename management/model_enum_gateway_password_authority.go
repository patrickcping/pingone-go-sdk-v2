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

// EnumGatewayPasswordAuthority This can be either `PING_ONE` or `LDAP`. If set to `PING_ONE`, PingOne authenticates with the external directory initially, then PingOne authenticates all subsequent sign-ons.
type EnumGatewayPasswordAuthority string

// List of EnumGatewayPasswordAuthority
const (
	ENUMGATEWAYPASSWORDAUTHORITY_PING_ONE EnumGatewayPasswordAuthority = "PING_ONE"
	ENUMGATEWAYPASSWORDAUTHORITY_LDAP     EnumGatewayPasswordAuthority = "LDAP"
)

// All allowed values of EnumGatewayPasswordAuthority enum
var AllowedEnumGatewayPasswordAuthorityEnumValues = []EnumGatewayPasswordAuthority{
	"PING_ONE",
	"LDAP",
}

func (v *EnumGatewayPasswordAuthority) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumGatewayPasswordAuthority(value)
	for _, existing := range AllowedEnumGatewayPasswordAuthorityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumGatewayPasswordAuthority(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumGatewayPasswordAuthorityFromValue returns a pointer to a valid EnumGatewayPasswordAuthority
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumGatewayPasswordAuthorityFromValue(v string) (*EnumGatewayPasswordAuthority, error) {
	ev := EnumGatewayPasswordAuthority(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumGatewayPasswordAuthority: valid values are %v", v, AllowedEnumGatewayPasswordAuthorityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumGatewayPasswordAuthority) IsValid() bool {
	for _, existing := range AllowedEnumGatewayPasswordAuthorityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumGatewayPasswordAuthority value
func (v EnumGatewayPasswordAuthority) Ptr() *EnumGatewayPasswordAuthority {
	return &v
}

type NullableEnumGatewayPasswordAuthority struct {
	value *EnumGatewayPasswordAuthority
	isSet bool
}

func (v NullableEnumGatewayPasswordAuthority) Get() *EnumGatewayPasswordAuthority {
	return v.value
}

func (v *NullableEnumGatewayPasswordAuthority) Set(val *EnumGatewayPasswordAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumGatewayPasswordAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumGatewayPasswordAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumGatewayPasswordAuthority(val *EnumGatewayPasswordAuthority) *NullableEnumGatewayPasswordAuthority {
	return &NullableEnumGatewayPasswordAuthority{value: val, isSet: true}
}

func (v NullableEnumGatewayPasswordAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumGatewayPasswordAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
