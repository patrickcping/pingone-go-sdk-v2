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

// EnumGatewayVendor A string that specifies the LDAP vendor. Options are `PingDirectory`, `Microsoft Active Directory`, `Oracle Directory Server Enterprise Edition`, `Oracle Unified Directory`, `CA Directory`, `OpenDJ Directory Server`, `IBM (Tivoli) Security Directory Server`, and `LDAPv3-compliant Directory Server`.
type EnumGatewayVendor string

// List of EnumGatewayVendor
const (
	ENUMGATEWAYVENDOR_PING_DIRECTORY                             EnumGatewayVendor = "PingDirectory"
	ENUMGATEWAYVENDOR_MICROSOFT_ACTIVE_DIRECTORY                 EnumGatewayVendor = "Microsoft Active Directory"
	ENUMGATEWAYVENDOR_ORACLE_DIRECTORY_SERVER_ENTERPRISE_EDITION EnumGatewayVendor = "Oracle Directory Server Enterprise Edition"
	ENUMGATEWAYVENDOR_ORACLE_UNIFIED_DIRECTORY                   EnumGatewayVendor = "Oracle Unified Directory"
	ENUMGATEWAYVENDOR_CA_DIRECTORY                               EnumGatewayVendor = "CA Directory"
	ENUMGATEWAYVENDOR_OPEN_DJ_DIRECTORY_SERVER                   EnumGatewayVendor = "OpenDJ Directory Server"
	ENUMGATEWAYVENDOR_IBM__TIVOLI_SECURITY_DIRECTORY_SERVER      EnumGatewayVendor = "IBM (Tivoli) Security Directory Server"
	ENUMGATEWAYVENDOR_LDAPV3_COMPLIANT_DIRECTORY_SERVER          EnumGatewayVendor = "LDAPv3-compliant Directory Server"
)

// All allowed values of EnumGatewayVendor enum
var AllowedEnumGatewayVendorEnumValues = []EnumGatewayVendor{
	"PingDirectory",
	"Microsoft Active Directory",
	"Oracle Directory Server Enterprise Edition",
	"Oracle Unified Directory",
	"CA Directory",
	"OpenDJ Directory Server",
	"IBM (Tivoli) Security Directory Server",
	"LDAPv3-compliant Directory Server",
}

func (v *EnumGatewayVendor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumGatewayVendor(value)
	for _, existing := range AllowedEnumGatewayVendorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumGatewayVendor(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumGatewayVendorFromValue returns a pointer to a valid EnumGatewayVendor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumGatewayVendorFromValue(v string) (*EnumGatewayVendor, error) {
	ev := EnumGatewayVendor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumGatewayVendor: valid values are %v", v, AllowedEnumGatewayVendorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumGatewayVendor) IsValid() bool {
	for _, existing := range AllowedEnumGatewayVendorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumGatewayVendor value
func (v EnumGatewayVendor) Ptr() *EnumGatewayVendor {
	return &v
}

type NullableEnumGatewayVendor struct {
	value *EnumGatewayVendor
	isSet bool
}

func (v NullableEnumGatewayVendor) Get() *EnumGatewayVendor {
	return v.value
}

func (v *NullableEnumGatewayVendor) Set(val *EnumGatewayVendor) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumGatewayVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumGatewayVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumGatewayVendor(val *EnumGatewayVendor) *NullableEnumGatewayVendor {
	return &NullableEnumGatewayVendor{value: val, isSet: true}
}

func (v NullableEnumGatewayVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumGatewayVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
