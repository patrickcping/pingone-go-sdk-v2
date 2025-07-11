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

// EnumUpdateStatus An enumeration that specifies one of the following values: AT_LATEST: The gateway instance's version is at or after the supported version marked latest. UPGRADE_AVAILABLE: The gateway instance's version is at the supported version that is marked recommended but there is a later supported version marked recommended. UPGRADE_RECOMMENDED: The gateway instance's version is at a known version but the version is not marked as recommended or latest. The version has greater than 30 days support. UPGRADE_REQUIRED: The gateway instance's version is at a known version but the version is not marked as recommended or latest. The version has support ending within the next month. NOT_SUPPORTED: The gateway instance's version is not known or supported.
type EnumUpdateStatus string

// List of EnumUpdateStatus
const (
	ENUMUPDATESTATUS_AT_LATEST           EnumUpdateStatus = "AT_LATEST"
	ENUMUPDATESTATUS_UPGRADE_AVAILABLE   EnumUpdateStatus = "UPGRADE_AVAILABLE"
	ENUMUPDATESTATUS_UPGRADE_RECOMMENDED EnumUpdateStatus = "UPGRADE_RECOMMENDED"
	ENUMUPDATESTATUS_UPGRADE_REQUIRED    EnumUpdateStatus = "UPGRADE_REQUIRED"
	ENUMUPDATESTATUS_NOT_SUPPORTED       EnumUpdateStatus = "NOT_SUPPORTED"
)

// All allowed values of EnumUpdateStatus enum
var AllowedEnumUpdateStatusEnumValues = []EnumUpdateStatus{
	"AT_LATEST",
	"UPGRADE_AVAILABLE",
	"UPGRADE_RECOMMENDED",
	"UPGRADE_REQUIRED",
	"NOT_SUPPORTED",
}

func (v *EnumUpdateStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUpdateStatus(value)
	for _, existing := range AllowedEnumUpdateStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUpdateStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUpdateStatusFromValue returns a pointer to a valid EnumUpdateStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUpdateStatusFromValue(v string) (*EnumUpdateStatus, error) {
	ev := EnumUpdateStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUpdateStatus: valid values are %v", v, AllowedEnumUpdateStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUpdateStatus) IsValid() bool {
	for _, existing := range AllowedEnumUpdateStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUpdateStatus value
func (v EnumUpdateStatus) Ptr() *EnumUpdateStatus {
	return &v
}

type NullableEnumUpdateStatus struct {
	value *EnumUpdateStatus
	isSet bool
}

func (v NullableEnumUpdateStatus) Get() *EnumUpdateStatus {
	return v.value
}

func (v *NullableEnumUpdateStatus) Set(val *EnumUpdateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUpdateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUpdateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUpdateStatus(val *EnumUpdateStatus) *NullableEnumUpdateStatus {
	return &NullableEnumUpdateStatus{value: val, isSet: true}
}

func (v NullableEnumUpdateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUpdateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
