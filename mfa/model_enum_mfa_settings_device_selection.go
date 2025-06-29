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

// EnumMFASettingsDeviceSelection A string that defines the device selection method. Options are DEFAULT_TO_FIRST (this is the default setting) and PROMPT_TO_SELECT.
type EnumMFASettingsDeviceSelection string

// List of EnumMFASettingsDeviceSelection
const (
	ENUMMFASETTINGSDEVICESELECTION_DEFAULT_TO_FIRST EnumMFASettingsDeviceSelection = "DEFAULT_TO_FIRST"
	ENUMMFASETTINGSDEVICESELECTION_PROMPT_TO_SELECT EnumMFASettingsDeviceSelection = "PROMPT_TO_SELECT"
)

// All allowed values of EnumMFASettingsDeviceSelection enum
var AllowedEnumMFASettingsDeviceSelectionEnumValues = []EnumMFASettingsDeviceSelection{
	"DEFAULT_TO_FIRST",
	"PROMPT_TO_SELECT",
}

func (v *EnumMFASettingsDeviceSelection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumMFASettingsDeviceSelection(value)
	for _, existing := range AllowedEnumMFASettingsDeviceSelectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumMFASettingsDeviceSelection(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumMFASettingsDeviceSelectionFromValue returns a pointer to a valid EnumMFASettingsDeviceSelection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumMFASettingsDeviceSelectionFromValue(v string) (*EnumMFASettingsDeviceSelection, error) {
	ev := EnumMFASettingsDeviceSelection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumMFASettingsDeviceSelection: valid values are %v", v, AllowedEnumMFASettingsDeviceSelectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumMFASettingsDeviceSelection) IsValid() bool {
	for _, existing := range AllowedEnumMFASettingsDeviceSelectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumMFASettingsDeviceSelection value
func (v EnumMFASettingsDeviceSelection) Ptr() *EnumMFASettingsDeviceSelection {
	return &v
}

type NullableEnumMFASettingsDeviceSelection struct {
	value *EnumMFASettingsDeviceSelection
	isSet bool
}

func (v NullableEnumMFASettingsDeviceSelection) Get() *EnumMFASettingsDeviceSelection {
	return v.value
}

func (v *NullableEnumMFASettingsDeviceSelection) Set(val *EnumMFASettingsDeviceSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumMFASettingsDeviceSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumMFASettingsDeviceSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumMFASettingsDeviceSelection(val *EnumMFASettingsDeviceSelection) *NullableEnumMFASettingsDeviceSelection {
	return &NullableEnumMFASettingsDeviceSelection{value: val, isSet: true}
}

func (v NullableEnumMFASettingsDeviceSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumMFASettingsDeviceSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
