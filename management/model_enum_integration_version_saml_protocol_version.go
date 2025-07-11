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

// EnumIntegrationVersionSAMLProtocolVersion The SAML protocol version supported 2.0, 1.1, or 1.0.
type EnumIntegrationVersionSAMLProtocolVersion string

// List of EnumIntegrationVersionSAMLProtocolVersion
const (
	ENUMINTEGRATIONVERSIONSAMLPROTOCOLVERSION__2_0 EnumIntegrationVersionSAMLProtocolVersion = "2.0"
	ENUMINTEGRATIONVERSIONSAMLPROTOCOLVERSION__1_1 EnumIntegrationVersionSAMLProtocolVersion = "1.1"
	ENUMINTEGRATIONVERSIONSAMLPROTOCOLVERSION__1_0 EnumIntegrationVersionSAMLProtocolVersion = "1.0"
)

// All allowed values of EnumIntegrationVersionSAMLProtocolVersion enum
var AllowedEnumIntegrationVersionSAMLProtocolVersionEnumValues = []EnumIntegrationVersionSAMLProtocolVersion{
	"2.0",
	"1.1",
	"1.0",
}

func (v *EnumIntegrationVersionSAMLProtocolVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumIntegrationVersionSAMLProtocolVersion(value)
	for _, existing := range AllowedEnumIntegrationVersionSAMLProtocolVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumIntegrationVersionSAMLProtocolVersion(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumIntegrationVersionSAMLProtocolVersionFromValue returns a pointer to a valid EnumIntegrationVersionSAMLProtocolVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumIntegrationVersionSAMLProtocolVersionFromValue(v string) (*EnumIntegrationVersionSAMLProtocolVersion, error) {
	ev := EnumIntegrationVersionSAMLProtocolVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumIntegrationVersionSAMLProtocolVersion: valid values are %v", v, AllowedEnumIntegrationVersionSAMLProtocolVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumIntegrationVersionSAMLProtocolVersion) IsValid() bool {
	for _, existing := range AllowedEnumIntegrationVersionSAMLProtocolVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumIntegrationVersionSAMLProtocolVersion value
func (v EnumIntegrationVersionSAMLProtocolVersion) Ptr() *EnumIntegrationVersionSAMLProtocolVersion {
	return &v
}

type NullableEnumIntegrationVersionSAMLProtocolVersion struct {
	value *EnumIntegrationVersionSAMLProtocolVersion
	isSet bool
}

func (v NullableEnumIntegrationVersionSAMLProtocolVersion) Get() *EnumIntegrationVersionSAMLProtocolVersion {
	return v.value
}

func (v *NullableEnumIntegrationVersionSAMLProtocolVersion) Set(val *EnumIntegrationVersionSAMLProtocolVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumIntegrationVersionSAMLProtocolVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumIntegrationVersionSAMLProtocolVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumIntegrationVersionSAMLProtocolVersion(val *EnumIntegrationVersionSAMLProtocolVersion) *NullableEnumIntegrationVersionSAMLProtocolVersion {
	return &NullableEnumIntegrationVersionSAMLProtocolVersion{value: val, isSet: true}
}

func (v NullableEnumIntegrationVersionSAMLProtocolVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumIntegrationVersionSAMLProtocolVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
