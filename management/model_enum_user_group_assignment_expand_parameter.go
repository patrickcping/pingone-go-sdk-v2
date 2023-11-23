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

// EnumUserGroupAssignmentExpandParameter When this parameter is appended to the request, the results are expanded to include additional information about each group.
type EnumUserGroupAssignmentExpandParameter string

// List of EnumUserGroupAssignmentExpandParameter
const (
	ENUMUSERGROUPASSIGNMENTEXPANDPARAMETER_GROUP EnumUserGroupAssignmentExpandParameter = "group"
)

// All allowed values of EnumUserGroupAssignmentExpandParameter enum
var AllowedEnumUserGroupAssignmentExpandParameterEnumValues = []EnumUserGroupAssignmentExpandParameter{
	"group",
}

func (v *EnumUserGroupAssignmentExpandParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumUserGroupAssignmentExpandParameter(value)
	for _, existing := range AllowedEnumUserGroupAssignmentExpandParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumUserGroupAssignmentExpandParameter(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumUserGroupAssignmentExpandParameterFromValue returns a pointer to a valid EnumUserGroupAssignmentExpandParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumUserGroupAssignmentExpandParameterFromValue(v string) (*EnumUserGroupAssignmentExpandParameter, error) {
	ev := EnumUserGroupAssignmentExpandParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumUserGroupAssignmentExpandParameter: valid values are %v", v, AllowedEnumUserGroupAssignmentExpandParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumUserGroupAssignmentExpandParameter) IsValid() bool {
	for _, existing := range AllowedEnumUserGroupAssignmentExpandParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumUserGroupAssignmentExpandParameter value
func (v EnumUserGroupAssignmentExpandParameter) Ptr() *EnumUserGroupAssignmentExpandParameter {
	return &v
}

type NullableEnumUserGroupAssignmentExpandParameter struct {
	value *EnumUserGroupAssignmentExpandParameter
	isSet bool
}

func (v NullableEnumUserGroupAssignmentExpandParameter) Get() *EnumUserGroupAssignmentExpandParameter {
	return v.value
}

func (v *NullableEnumUserGroupAssignmentExpandParameter) Set(val *EnumUserGroupAssignmentExpandParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumUserGroupAssignmentExpandParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumUserGroupAssignmentExpandParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumUserGroupAssignmentExpandParameter(val *EnumUserGroupAssignmentExpandParameter) *NullableEnumUserGroupAssignmentExpandParameter {
	return &NullableEnumUserGroupAssignmentExpandParameter{value: val, isSet: true}
}

func (v NullableEnumUserGroupAssignmentExpandParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumUserGroupAssignmentExpandParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

