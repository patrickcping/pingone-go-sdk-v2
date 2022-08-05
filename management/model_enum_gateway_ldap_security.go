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

// EnumGatewayLDAPSecurity (Optional) A string that specifies the connection security type. Options are None, TLS, and StartTLS. The default value is None.
type EnumGatewayLDAPSecurity string

// List of EnumGatewayLDAPSecurity
const (
	ENUMGATEWAYLDAPSECURITY_NONE EnumGatewayLDAPSecurity = "None"
	ENUMGATEWAYLDAPSECURITY_TLS EnumGatewayLDAPSecurity = "TLS"
	ENUMGATEWAYLDAPSECURITY_START_TLS EnumGatewayLDAPSecurity = "StartTLS"
)

// All allowed values of EnumGatewayLDAPSecurity enum
var AllowedEnumGatewayLDAPSecurityEnumValues = []EnumGatewayLDAPSecurity{
	"None",
	"TLS",
	"StartTLS",
}

func (v *EnumGatewayLDAPSecurity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumGatewayLDAPSecurity(value)
	for _, existing := range AllowedEnumGatewayLDAPSecurityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumGatewayLDAPSecurity", value)
}

// NewEnumGatewayLDAPSecurityFromValue returns a pointer to a valid EnumGatewayLDAPSecurity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumGatewayLDAPSecurityFromValue(v string) (*EnumGatewayLDAPSecurity, error) {
	ev := EnumGatewayLDAPSecurity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumGatewayLDAPSecurity: valid values are %v", v, AllowedEnumGatewayLDAPSecurityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumGatewayLDAPSecurity) IsValid() bool {
	for _, existing := range AllowedEnumGatewayLDAPSecurityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumGatewayLDAPSecurity value
func (v EnumGatewayLDAPSecurity) Ptr() *EnumGatewayLDAPSecurity {
	return &v
}

type NullableEnumGatewayLDAPSecurity struct {
	value *EnumGatewayLDAPSecurity
	isSet bool
}

func (v NullableEnumGatewayLDAPSecurity) Get() *EnumGatewayLDAPSecurity {
	return v.value
}

func (v *NullableEnumGatewayLDAPSecurity) Set(val *EnumGatewayLDAPSecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumGatewayLDAPSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumGatewayLDAPSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumGatewayLDAPSecurity(val *EnumGatewayLDAPSecurity) *NullableEnumGatewayLDAPSecurity {
	return &NullableEnumGatewayLDAPSecurity{value: val, isSet: true}
}

func (v NullableEnumGatewayLDAPSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumGatewayLDAPSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
