/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumEnabledStatus A string that specifies whether device integrity detection takes place on mobile devices, for the application's enrollment and authentication events ENABLED, DISABLED
type EnumEnabledStatus string

// List of EnumEnabledStatus
const (
	ENABLED EnumEnabledStatus = "ENABLED"
	DISABLED EnumEnabledStatus = "DISABLED"
)

// All allowed values of EnumEnabledStatus enum
var AllowedEnumEnabledStatusEnumValues = []EnumEnabledStatus{
	"ENABLED",
	"DISABLED",
}

func (v *EnumEnabledStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumEnabledStatus(value)
	for _, existing := range AllowedEnumEnabledStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumEnabledStatus", value)
}

// NewEnumEnabledStatusFromValue returns a pointer to a valid EnumEnabledStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumEnabledStatusFromValue(v string) (*EnumEnabledStatus, error) {
	ev := EnumEnabledStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumEnabledStatus: valid values are %v", v, AllowedEnumEnabledStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumEnabledStatus) IsValid() bool {
	for _, existing := range AllowedEnumEnabledStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumEnabledStatus value
func (v EnumEnabledStatus) Ptr() *EnumEnabledStatus {
	return &v
}

type NullableEnumEnabledStatus struct {
	value *EnumEnabledStatus
	isSet bool
}

func (v NullableEnumEnabledStatus) Get() *EnumEnabledStatus {
	return v.value
}

func (v *NullableEnumEnabledStatus) Set(val *EnumEnabledStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumEnabledStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumEnabledStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumEnabledStatus(val *EnumEnabledStatus) *NullableEnumEnabledStatus {
	return &NullableEnumEnabledStatus{value: val, isSet: true}
}

func (v NullableEnumEnabledStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumEnabledStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

