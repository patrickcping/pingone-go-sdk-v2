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

// EnumSignOnPolicyNoDeviceMode A string that specifies the device mode for the MFA flow. Options are `BYPASS` to allow MFA without a specified device, or `BLOCK` to block the MFA flow if no device is specified. To use this configuration option, the authorize request must include a signed `login_hint_token` property. For more information, see Authorize (Browserless and MFA Only Flows)
type EnumSignOnPolicyNoDeviceMode string

// List of EnumSignOnPolicyNoDeviceMode
const (
	ENUMSIGNONPOLICYNODEVICEMODE_BYPASS EnumSignOnPolicyNoDeviceMode = "BYPASS"
	ENUMSIGNONPOLICYNODEVICEMODE_BLOCK EnumSignOnPolicyNoDeviceMode = "BLOCK"
)

// All allowed values of EnumSignOnPolicyNoDeviceMode enum
var AllowedEnumSignOnPolicyNoDeviceModeEnumValues = []EnumSignOnPolicyNoDeviceMode{
	"BYPASS",
	"BLOCK",
}

func (v *EnumSignOnPolicyNoDeviceMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSignOnPolicyNoDeviceMode(value)
	for _, existing := range AllowedEnumSignOnPolicyNoDeviceModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumSignOnPolicyNoDeviceMode(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumSignOnPolicyNoDeviceModeFromValue returns a pointer to a valid EnumSignOnPolicyNoDeviceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSignOnPolicyNoDeviceModeFromValue(v string) (*EnumSignOnPolicyNoDeviceMode, error) {
	ev := EnumSignOnPolicyNoDeviceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSignOnPolicyNoDeviceMode: valid values are %v", v, AllowedEnumSignOnPolicyNoDeviceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSignOnPolicyNoDeviceMode) IsValid() bool {
	for _, existing := range AllowedEnumSignOnPolicyNoDeviceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSignOnPolicyNoDeviceMode value
func (v EnumSignOnPolicyNoDeviceMode) Ptr() *EnumSignOnPolicyNoDeviceMode {
	return &v
}

type NullableEnumSignOnPolicyNoDeviceMode struct {
	value *EnumSignOnPolicyNoDeviceMode
	isSet bool
}

func (v NullableEnumSignOnPolicyNoDeviceMode) Get() *EnumSignOnPolicyNoDeviceMode {
	return v.value
}

func (v *NullableEnumSignOnPolicyNoDeviceMode) Set(val *EnumSignOnPolicyNoDeviceMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSignOnPolicyNoDeviceMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSignOnPolicyNoDeviceMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSignOnPolicyNoDeviceMode(val *EnumSignOnPolicyNoDeviceMode) *NullableEnumSignOnPolicyNoDeviceMode {
	return &NullableEnumSignOnPolicyNoDeviceMode{value: val, isSet: true}
}

func (v NullableEnumSignOnPolicyNoDeviceMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSignOnPolicyNoDeviceMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

