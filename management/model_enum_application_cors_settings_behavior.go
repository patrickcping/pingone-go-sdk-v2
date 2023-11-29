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

// EnumApplicationCorsSettingsBehavior The behavior of CORS for the application. `ALLOWS_NO_ORIGINS` rejects all CORS requests.  `ALLOW_SPECIFIC_ORIGINS` rejects all CORS requests except those listed in `corsSettings.origins`.
type EnumApplicationCorsSettingsBehavior string

// List of EnumApplicationCorsSettingsBehavior
const (
	ENUMAPPLICATIONCORSSETTINGSBEHAVIOR_NO_ORIGINS EnumApplicationCorsSettingsBehavior = "ALLOW_NO_ORIGINS"
	ENUMAPPLICATIONCORSSETTINGSBEHAVIOR_SPECIFIC_ORIGINS EnumApplicationCorsSettingsBehavior = "ALLOW_SPECIFIC_ORIGINS"
)

// All allowed values of EnumApplicationCorsSettingsBehavior enum
var AllowedEnumApplicationCorsSettingsBehaviorEnumValues = []EnumApplicationCorsSettingsBehavior{
	"ALLOW_NO_ORIGINS",
	"ALLOW_SPECIFIC_ORIGINS",
}

func (v *EnumApplicationCorsSettingsBehavior) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationCorsSettingsBehavior(value)
	for _, existing := range AllowedEnumApplicationCorsSettingsBehaviorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumApplicationCorsSettingsBehavior(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumApplicationCorsSettingsBehaviorFromValue returns a pointer to a valid EnumApplicationCorsSettingsBehavior
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationCorsSettingsBehaviorFromValue(v string) (*EnumApplicationCorsSettingsBehavior, error) {
	ev := EnumApplicationCorsSettingsBehavior(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationCorsSettingsBehavior: valid values are %v", v, AllowedEnumApplicationCorsSettingsBehaviorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationCorsSettingsBehavior) IsValid() bool {
	for _, existing := range AllowedEnumApplicationCorsSettingsBehaviorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationCorsSettingsBehavior value
func (v EnumApplicationCorsSettingsBehavior) Ptr() *EnumApplicationCorsSettingsBehavior {
	return &v
}

type NullableEnumApplicationCorsSettingsBehavior struct {
	value *EnumApplicationCorsSettingsBehavior
	isSet bool
}

func (v NullableEnumApplicationCorsSettingsBehavior) Get() *EnumApplicationCorsSettingsBehavior {
	return v.value
}

func (v *NullableEnumApplicationCorsSettingsBehavior) Set(val *EnumApplicationCorsSettingsBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationCorsSettingsBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationCorsSettingsBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationCorsSettingsBehavior(val *EnumApplicationCorsSettingsBehavior) *NullableEnumApplicationCorsSettingsBehavior {
	return &NullableEnumApplicationCorsSettingsBehavior{value: val, isSet: true}
}

func (v NullableEnumApplicationCorsSettingsBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationCorsSettingsBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

