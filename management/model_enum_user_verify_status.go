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

// EnumUserVerifyStatus Indicates whether ID verification can be done for the user. This value can be NOT_INITIATED (the initial value), ENABLED, or DISABLED. If the user verification status is DISABLED, a new verification status cannot be created for that user until the status is changed to ENABLED.
type EnumUserVerifyStatus string

// List of EnumUserVerifyStatus
const (
	ENUMUSERVERIFYSTATUS_NOT_INITIATED EnumUserVerifyStatus = "NOT_INITIATED"
	ENUMUSERVERIFYSTATUS_ENABLED       EnumUserVerifyStatus = "ENABLED"
	ENUMUSERVERIFYSTATUS_DISABLED      EnumUserVerifyStatus = "DISABLED"
)

// All allowed values of EnumUserVerifyStatus enum
var AllowedEnumUserVerifyStatusEnumValues = []EnumUserVerifyStatus{
	"NOT_INITIATED",
	"ENABLED",
	"DISABLED",
}

func (v *EnumUserVerifyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserVerifyStatus(value)
	for _, existing := range AllowedEnumUserVerifyStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserVerifyStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserVerifyStatusFromValue returns a pointer to a valid EnumUserVerifyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserVerifyStatusFromValue(v string) (*EnumUserVerifyStatus, error) {
	ev := EnumUserVerifyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserVerifyStatus: valid values are %v", v, AllowedEnumUserVerifyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserVerifyStatus) IsValid() bool {
	for _, existing := range AllowedEnumUserVerifyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserVerifyStatus value
func (v EnumUserVerifyStatus) Ptr() *EnumUserVerifyStatus {
	return &v
}

type NullableEnumUserVerifyStatus struct {
	value *EnumUserVerifyStatus
	isSet bool
}

func (v NullableEnumUserVerifyStatus) Get() *EnumUserVerifyStatus {
	return v.value
}

func (v *NullableEnumUserVerifyStatus) Set(val *EnumUserVerifyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserVerifyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserVerifyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserVerifyStatus(val *EnumUserVerifyStatus) *NullableEnumUserVerifyStatus {
	return &NullableEnumUserVerifyStatus{value: val, isSet: true}
}

func (v NullableEnumUserVerifyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserVerifyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
