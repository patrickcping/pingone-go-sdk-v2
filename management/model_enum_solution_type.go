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

// EnumSolutionType The solution type selected when creating the environment. Ignored on PUT operations. The following values are supported: `CIAM_TRIAL`. The Customer trial experience. Indicates the Customer solution type, and the Solution Designer was selected. `WF_TRIAL`. The Workforce trial experience. Indicates the Workforce solution type, and the Solution Designer was selected. `CUSTOMER`. Indicates the Customer solution type was selected. This solution type uses PingOne MFA, rather than PingID. `WORKFORCE`. Indicates the Workforce solution type was selected. This solution type uses PingID, rather than PingOne MFA.
type EnumSolutionType string

// List of EnumSolutionType
const (
	ENUMSOLUTIONTYPE_WORKFORCE  EnumSolutionType = "WORKFORCE"
	ENUMSOLUTIONTYPE_CUSTOMER   EnumSolutionType = "CUSTOMER"
	ENUMSOLUTIONTYPE_CIAM_TRIAL EnumSolutionType = "CIAM_TRIAL"
	ENUMSOLUTIONTYPE_WF_TRIAL   EnumSolutionType = "WF_TRIAL"
)

// All allowed values of EnumSolutionType enum
var AllowedEnumSolutionTypeEnumValues = []EnumSolutionType{
	"WORKFORCE",
	"CUSTOMER",
	"CIAM_TRIAL",
	"WF_TRIAL",
}

func (v *EnumSolutionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSolutionType(value)
	for _, existing := range AllowedEnumSolutionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumSolutionType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumSolutionTypeFromValue returns a pointer to a valid EnumSolutionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSolutionTypeFromValue(v string) (*EnumSolutionType, error) {
	ev := EnumSolutionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSolutionType: valid values are %v", v, AllowedEnumSolutionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSolutionType) IsValid() bool {
	for _, existing := range AllowedEnumSolutionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSolutionType value
func (v EnumSolutionType) Ptr() *EnumSolutionType {
	return &v
}

type NullableEnumSolutionType struct {
	value *EnumSolutionType
	isSet bool
}

func (v NullableEnumSolutionType) Get() *EnumSolutionType {
	return v.value
}

func (v *NullableEnumSolutionType) Set(val *EnumSolutionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSolutionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSolutionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSolutionType(val *EnumSolutionType) *NullableEnumSolutionType {
	return &NullableEnumSolutionType{value: val, isSet: true}
}

func (v NullableEnumSolutionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSolutionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
