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

// EnumApplicationAccessControlGroupType A string that specifies the group type required to access the application. Options are ANY_GROUP (the actor must belong to at least one group listed in the accessControl.group.groups property) and ALL_GROUPS (the actor must belong to all groups listed in the accessControl.group.groups property).
type EnumApplicationAccessControlGroupType string

// List of EnumApplicationAccessControlGroupType
const (
	ENUMAPPLICATIONACCESSCONTROLGROUPTYPE_ANY_GROUP EnumApplicationAccessControlGroupType = "ANY_GROUP"
	ENUMAPPLICATIONACCESSCONTROLGROUPTYPE_ALL_GROUPS EnumApplicationAccessControlGroupType = "ALL_GROUPS"
)

// All allowed values of EnumApplicationAccessControlGroupType enum
var AllowedEnumApplicationAccessControlGroupTypeEnumValues = []EnumApplicationAccessControlGroupType{
	"ANY_GROUP",
	"ALL_GROUPS",
}

func (v *EnumApplicationAccessControlGroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationAccessControlGroupType(value)
	for _, existing := range AllowedEnumApplicationAccessControlGroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumApplicationAccessControlGroupType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumApplicationAccessControlGroupTypeFromValue returns a pointer to a valid EnumApplicationAccessControlGroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationAccessControlGroupTypeFromValue(v string) (*EnumApplicationAccessControlGroupType, error) {
	ev := EnumApplicationAccessControlGroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationAccessControlGroupType: valid values are %v", v, AllowedEnumApplicationAccessControlGroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationAccessControlGroupType) IsValid() bool {
	for _, existing := range AllowedEnumApplicationAccessControlGroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationAccessControlGroupType value
func (v EnumApplicationAccessControlGroupType) Ptr() *EnumApplicationAccessControlGroupType {
	return &v
}

type NullableEnumApplicationAccessControlGroupType struct {
	value *EnumApplicationAccessControlGroupType
	isSet bool
}

func (v NullableEnumApplicationAccessControlGroupType) Get() *EnumApplicationAccessControlGroupType {
	return v.value
}

func (v *NullableEnumApplicationAccessControlGroupType) Set(val *EnumApplicationAccessControlGroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationAccessControlGroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationAccessControlGroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationAccessControlGroupType(val *EnumApplicationAccessControlGroupType) *NullableEnumApplicationAccessControlGroupType {
	return &NullableEnumApplicationAccessControlGroupType{value: val, isSet: true}
}

func (v NullableEnumApplicationAccessControlGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationAccessControlGroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

