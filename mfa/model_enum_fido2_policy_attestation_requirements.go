/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// EnumFIDO2PolicyAttestationRequirements Set to `DIRECT` if you want to require attestation. Set to `NONE` if you don't want to require attestation. If you set `attestationRequirements` to `NONE`, you should also set `mdsAuthenticatorsRequirements.option` to `NONE`.
type EnumFIDO2PolicyAttestationRequirements string

// List of EnumFIDO2PolicyAttestationRequirements
const (
	ENUMFIDO2POLICYATTESTATIONREQUIREMENTS_DIRECT EnumFIDO2PolicyAttestationRequirements = "DIRECT"
	ENUMFIDO2POLICYATTESTATIONREQUIREMENTS_NONE   EnumFIDO2PolicyAttestationRequirements = "NONE"
)

// All allowed values of EnumFIDO2PolicyAttestationRequirements enum
var AllowedEnumFIDO2PolicyAttestationRequirementsEnumValues = []EnumFIDO2PolicyAttestationRequirements{
	"DIRECT",
	"NONE",
}

func (v *EnumFIDO2PolicyAttestationRequirements) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFIDO2PolicyAttestationRequirements(value)
	for _, existing := range AllowedEnumFIDO2PolicyAttestationRequirementsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFIDO2PolicyAttestationRequirements", value)
}

// NewEnumFIDO2PolicyAttestationRequirementsFromValue returns a pointer to a valid EnumFIDO2PolicyAttestationRequirements
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFIDO2PolicyAttestationRequirementsFromValue(v string) (*EnumFIDO2PolicyAttestationRequirements, error) {
	ev := EnumFIDO2PolicyAttestationRequirements(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFIDO2PolicyAttestationRequirements: valid values are %v", v, AllowedEnumFIDO2PolicyAttestationRequirementsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFIDO2PolicyAttestationRequirements) IsValid() bool {
	for _, existing := range AllowedEnumFIDO2PolicyAttestationRequirementsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFIDO2PolicyAttestationRequirements value
func (v EnumFIDO2PolicyAttestationRequirements) Ptr() *EnumFIDO2PolicyAttestationRequirements {
	return &v
}

type NullableEnumFIDO2PolicyAttestationRequirements struct {
	value *EnumFIDO2PolicyAttestationRequirements
	isSet bool
}

func (v NullableEnumFIDO2PolicyAttestationRequirements) Get() *EnumFIDO2PolicyAttestationRequirements {
	return v.value
}

func (v *NullableEnumFIDO2PolicyAttestationRequirements) Set(val *EnumFIDO2PolicyAttestationRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFIDO2PolicyAttestationRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFIDO2PolicyAttestationRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFIDO2PolicyAttestationRequirements(val *EnumFIDO2PolicyAttestationRequirements) *NullableEnumFIDO2PolicyAttestationRequirements {
	return &NullableEnumFIDO2PolicyAttestationRequirements{value: val, isSet: true}
}

func (v NullableEnumFIDO2PolicyAttestationRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFIDO2PolicyAttestationRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
