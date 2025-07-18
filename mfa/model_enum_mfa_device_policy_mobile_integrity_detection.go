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

// EnumMFADevicePolicyMobileIntegrityDetection Controls how authentication or registration attempts should proceed if a device integrity check does not receive a response. Set the value to `permissive` if you want to allow the process to continue. Set the value to `restrictive` if you want to block the user in such situations.
type EnumMFADevicePolicyMobileIntegrityDetection string

// List of EnumMFADevicePolicyMobileIntegrityDetection
const (
	ENUMMFADEVICEPOLICYMOBILEINTEGRITYDETECTION_PERMISSIVE  EnumMFADevicePolicyMobileIntegrityDetection = "permissive"
	ENUMMFADEVICEPOLICYMOBILEINTEGRITYDETECTION_RESTRICTIVE EnumMFADevicePolicyMobileIntegrityDetection = "restrictive"
)

// All allowed values of EnumMFADevicePolicyMobileIntegrityDetection enum
var AllowedEnumMFADevicePolicyMobileIntegrityDetectionEnumValues = []EnumMFADevicePolicyMobileIntegrityDetection{
	"permissive",
	"restrictive",
}

func (v *EnumMFADevicePolicyMobileIntegrityDetection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumMFADevicePolicyMobileIntegrityDetection(value)
	for _, existing := range AllowedEnumMFADevicePolicyMobileIntegrityDetectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumMFADevicePolicyMobileIntegrityDetection(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumMFADevicePolicyMobileIntegrityDetectionFromValue returns a pointer to a valid EnumMFADevicePolicyMobileIntegrityDetection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumMFADevicePolicyMobileIntegrityDetectionFromValue(v string) (*EnumMFADevicePolicyMobileIntegrityDetection, error) {
	ev := EnumMFADevicePolicyMobileIntegrityDetection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumMFADevicePolicyMobileIntegrityDetection: valid values are %v", v, AllowedEnumMFADevicePolicyMobileIntegrityDetectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumMFADevicePolicyMobileIntegrityDetection) IsValid() bool {
	for _, existing := range AllowedEnumMFADevicePolicyMobileIntegrityDetectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumMFADevicePolicyMobileIntegrityDetection value
func (v EnumMFADevicePolicyMobileIntegrityDetection) Ptr() *EnumMFADevicePolicyMobileIntegrityDetection {
	return &v
}

type NullableEnumMFADevicePolicyMobileIntegrityDetection struct {
	value *EnumMFADevicePolicyMobileIntegrityDetection
	isSet bool
}

func (v NullableEnumMFADevicePolicyMobileIntegrityDetection) Get() *EnumMFADevicePolicyMobileIntegrityDetection {
	return v.value
}

func (v *NullableEnumMFADevicePolicyMobileIntegrityDetection) Set(val *EnumMFADevicePolicyMobileIntegrityDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumMFADevicePolicyMobileIntegrityDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumMFADevicePolicyMobileIntegrityDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumMFADevicePolicyMobileIntegrityDetection(val *EnumMFADevicePolicyMobileIntegrityDetection) *NullableEnumMFADevicePolicyMobileIntegrityDetection {
	return &NullableEnumMFADevicePolicyMobileIntegrityDetection{value: val, isSet: true}
}

func (v NullableEnumMFADevicePolicyMobileIntegrityDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumMFADevicePolicyMobileIntegrityDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
