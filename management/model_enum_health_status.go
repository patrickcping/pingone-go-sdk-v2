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

// EnumHealthStatus An enumeration that specifies whether or not the gateway is in a healthy state.
type EnumHealthStatus string

// List of EnumHealthStatus
const (
	ENUMHEALTHSTATUS_HEALTHY EnumHealthStatus = "HEALTHY"
	ENUMHEALTHSTATUS_DEGRADED EnumHealthStatus = "DEGRADED"
	ENUMHEALTHSTATUS_UNHEALTHY EnumHealthStatus = "UNHEALTHY"
)

// All allowed values of EnumHealthStatus enum
var AllowedEnumHealthStatusEnumValues = []EnumHealthStatus{
	"HEALTHY",
	"DEGRADED",
	"UNHEALTHY",
}

func (v *EnumHealthStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumHealthStatus(value)
	for _, existing := range AllowedEnumHealthStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumHealthStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumHealthStatusFromValue returns a pointer to a valid EnumHealthStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumHealthStatusFromValue(v string) (*EnumHealthStatus, error) {
	ev := EnumHealthStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumHealthStatus: valid values are %v", v, AllowedEnumHealthStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumHealthStatus) IsValid() bool {
	for _, existing := range AllowedEnumHealthStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumHealthStatus value
func (v EnumHealthStatus) Ptr() *EnumHealthStatus {
	return &v
}

type NullableEnumHealthStatus struct {
	value *EnumHealthStatus
	isSet bool
}

func (v NullableEnumHealthStatus) Get() *EnumHealthStatus {
	return v.value
}

func (v *NullableEnumHealthStatus) Set(val *EnumHealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumHealthStatus(val *EnumHealthStatus) *NullableEnumHealthStatus {
	return &NullableEnumHealthStatus{value: val, isSet: true}
}

func (v NullableEnumHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

