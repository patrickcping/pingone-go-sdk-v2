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

// EnumGroupSourceType External groups only. Set during user creation/update. Supported values are `GATEWAY`, `IDP`
type EnumGroupSourceType string

// List of EnumGroupSourceType
const (
	ENUMGROUPSOURCETYPE_GATEWAY EnumGroupSourceType = "GATEWAY"
	ENUMGROUPSOURCETYPE_IDP     EnumGroupSourceType = "IDP"
)

// All allowed values of EnumGroupSourceType enum
var AllowedEnumGroupSourceTypeEnumValues = []EnumGroupSourceType{
	"GATEWAY",
	"IDP",
}

func (v *EnumGroupSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumGroupSourceType(value)
	for _, existing := range AllowedEnumGroupSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumGroupSourceType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumGroupSourceTypeFromValue returns a pointer to a valid EnumGroupSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumGroupSourceTypeFromValue(v string) (*EnumGroupSourceType, error) {
	ev := EnumGroupSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumGroupSourceType: valid values are %v", v, AllowedEnumGroupSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumGroupSourceType) IsValid() bool {
	for _, existing := range AllowedEnumGroupSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumGroupSourceType value
func (v EnumGroupSourceType) Ptr() *EnumGroupSourceType {
	return &v
}

type NullableEnumGroupSourceType struct {
	value *EnumGroupSourceType
	isSet bool
}

func (v NullableEnumGroupSourceType) Get() *EnumGroupSourceType {
	return v.value
}

func (v *NullableEnumGroupSourceType) Set(val *EnumGroupSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumGroupSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumGroupSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumGroupSourceType(val *EnumGroupSourceType) *NullableEnumGroupSourceType {
	return &NullableEnumGroupSourceType{value: val, isSet: true}
}

func (v NullableEnumGroupSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumGroupSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
