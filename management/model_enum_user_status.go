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

// EnumUserStatus A string that specifies the account locked state.
type EnumUserStatus string

// List of EnumUserStatus
const (
	ENUMUSERSTATUS_LOCKED EnumUserStatus = "LOCKED"
	ENUMUSERSTATUS_OK EnumUserStatus = "OK"
)

// All allowed values of EnumUserStatus enum
var AllowedEnumUserStatusEnumValues = []EnumUserStatus{
	"LOCKED",
	"OK",
}

func (v *EnumUserStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserStatus(value)
	for _, existing := range AllowedEnumUserStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserStatusFromValue returns a pointer to a valid EnumUserStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserStatusFromValue(v string) (*EnumUserStatus, error) {
	ev := EnumUserStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserStatus: valid values are %v", v, AllowedEnumUserStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserStatus) IsValid() bool {
	for _, existing := range AllowedEnumUserStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserStatus value
func (v EnumUserStatus) Ptr() *EnumUserStatus {
	return &v
}

type NullableEnumUserStatus struct {
	value *EnumUserStatus
	isSet bool
}

func (v NullableEnumUserStatus) Get() *EnumUserStatus {
	return v.value
}

func (v *NullableEnumUserStatus) Set(val *EnumUserStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserStatus(val *EnumUserStatus) *NullableEnumUserStatus {
	return &NullableEnumUserStatus{value: val, isSet: true}
}

func (v NullableEnumUserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

