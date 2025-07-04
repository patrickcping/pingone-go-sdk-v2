/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"fmt"
)

// EnumTimeUnit A string that defines the time unit of a specified duration in `SECONDS` or `MINUTES`.
type EnumTimeUnit string

// List of EnumTimeUnit
const (
	ENUMTIMEUNIT_SECONDS EnumTimeUnit = "SECONDS"
	ENUMTIMEUNIT_MINUTES EnumTimeUnit = "MINUTES"
)

// All allowed values of EnumTimeUnit enum
var AllowedEnumTimeUnitEnumValues = []EnumTimeUnit{
	"SECONDS",
	"MINUTES",
}

func (v *EnumTimeUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumTimeUnit(value)
	for _, existing := range AllowedEnumTimeUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumTimeUnit(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumTimeUnitFromValue returns a pointer to a valid EnumTimeUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumTimeUnitFromValue(v string) (*EnumTimeUnit, error) {
	ev := EnumTimeUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumTimeUnit: valid values are %v", v, AllowedEnumTimeUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumTimeUnit) IsValid() bool {
	for _, existing := range AllowedEnumTimeUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumTimeUnit value
func (v EnumTimeUnit) Ptr() *EnumTimeUnit {
	return &v
}

type NullableEnumTimeUnit struct {
	value *EnumTimeUnit
	isSet bool
}

func (v NullableEnumTimeUnit) Get() *EnumTimeUnit {
	return v.value
}

func (v *NullableEnumTimeUnit) Set(val *EnumTimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumTimeUnit(val *EnumTimeUnit) *NullableEnumTimeUnit {
	return &NullableEnumTimeUnit{value: val, isSet: true}
}

func (v NullableEnumTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
