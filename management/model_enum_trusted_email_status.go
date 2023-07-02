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

// EnumTrustedEmailStatus A string that specifies the status of the trusted email address.
type EnumTrustedEmailStatus string

// List of EnumTrustedEmailStatus
const (
	ENUMTRUSTEDEMAILSTATUS_ACTIVE EnumTrustedEmailStatus = "ACTIVE"
	ENUMTRUSTEDEMAILSTATUS_VERIFICATION_REQUIRED EnumTrustedEmailStatus = "VERIFICATION_REQUIRED"
)

// All allowed values of EnumTrustedEmailStatus enum
var AllowedEnumTrustedEmailStatusEnumValues = []EnumTrustedEmailStatus{
	"ACTIVE",
	"VERIFICATION_REQUIRED",
}

func (v *EnumTrustedEmailStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumTrustedEmailStatus(value)
	for _, existing := range AllowedEnumTrustedEmailStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumTrustedEmailStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumTrustedEmailStatusFromValue returns a pointer to a valid EnumTrustedEmailStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumTrustedEmailStatusFromValue(v string) (*EnumTrustedEmailStatus, error) {
	ev := EnumTrustedEmailStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumTrustedEmailStatus: valid values are %v", v, AllowedEnumTrustedEmailStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumTrustedEmailStatus) IsValid() bool {
	for _, existing := range AllowedEnumTrustedEmailStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumTrustedEmailStatus value
func (v EnumTrustedEmailStatus) Ptr() *EnumTrustedEmailStatus {
	return &v
}

type NullableEnumTrustedEmailStatus struct {
	value *EnumTrustedEmailStatus
	isSet bool
}

func (v NullableEnumTrustedEmailStatus) Get() *EnumTrustedEmailStatus {
	return v.value
}

func (v *NullableEnumTrustedEmailStatus) Set(val *EnumTrustedEmailStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTrustedEmailStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTrustedEmailStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTrustedEmailStatus(val *EnumTrustedEmailStatus) *NullableEnumTrustedEmailStatus {
	return &NullableEnumTrustedEmailStatus{value: val, isSet: true}
}

func (v NullableEnumTrustedEmailStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTrustedEmailStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

