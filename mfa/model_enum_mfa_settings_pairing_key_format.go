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

// EnumMFASettingsPairingKeyFormat String that controls the type of pairing key issued. The valid values are NUMERIC (12-digit key) and ALPHANUMERIC (16-character alphanumeric key).
type EnumMFASettingsPairingKeyFormat string

// List of EnumMFASettingsPairingKeyFormat
const (
	ENUMMFASETTINGSPAIRINGKEYFORMAT_NUMERIC EnumMFASettingsPairingKeyFormat = "NUMERIC"
	ENUMMFASETTINGSPAIRINGKEYFORMAT_ALPHANUMERIC EnumMFASettingsPairingKeyFormat = "ALPHANUMERIC"
)

// All allowed values of EnumMFASettingsPairingKeyFormat enum
var AllowedEnumMFASettingsPairingKeyFormatEnumValues = []EnumMFASettingsPairingKeyFormat{
	"NUMERIC",
	"ALPHANUMERIC",
}

func (v *EnumMFASettingsPairingKeyFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumMFASettingsPairingKeyFormat(value)
	for _, existing := range AllowedEnumMFASettingsPairingKeyFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumMFASettingsPairingKeyFormat(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumMFASettingsPairingKeyFormatFromValue returns a pointer to a valid EnumMFASettingsPairingKeyFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumMFASettingsPairingKeyFormatFromValue(v string) (*EnumMFASettingsPairingKeyFormat, error) {
	ev := EnumMFASettingsPairingKeyFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumMFASettingsPairingKeyFormat: valid values are %v", v, AllowedEnumMFASettingsPairingKeyFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumMFASettingsPairingKeyFormat) IsValid() bool {
	for _, existing := range AllowedEnumMFASettingsPairingKeyFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumMFASettingsPairingKeyFormat value
func (v EnumMFASettingsPairingKeyFormat) Ptr() *EnumMFASettingsPairingKeyFormat {
	return &v
}

type NullableEnumMFASettingsPairingKeyFormat struct {
	value *EnumMFASettingsPairingKeyFormat
	isSet bool
}

func (v NullableEnumMFASettingsPairingKeyFormat) Get() *EnumMFASettingsPairingKeyFormat {
	return v.value
}

func (v *NullableEnumMFASettingsPairingKeyFormat) Set(val *EnumMFASettingsPairingKeyFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumMFASettingsPairingKeyFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumMFASettingsPairingKeyFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumMFASettingsPairingKeyFormat(val *EnumMFASettingsPairingKeyFormat) *NullableEnumMFASettingsPairingKeyFormat {
	return &NullableEnumMFASettingsPairingKeyFormat{value: val, isSet: true}
}

func (v NullableEnumMFASettingsPairingKeyFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumMFASettingsPairingKeyFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

