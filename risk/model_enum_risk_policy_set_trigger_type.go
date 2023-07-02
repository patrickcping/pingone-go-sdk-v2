/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// EnumRiskPolicySetTriggerType A string that specifies the trigger type for the policy set. The only option is `POLICY_SET_STAGING`.
type EnumRiskPolicySetTriggerType string

// List of EnumRiskPolicySetTriggerType
const (
	ENUMRISKPOLICYSETTRIGGERTYPE_POLICY_SET_STAGING EnumRiskPolicySetTriggerType = "POLICY_SET_STAGING"
)

// All allowed values of EnumRiskPolicySetTriggerType enum
var AllowedEnumRiskPolicySetTriggerTypeEnumValues = []EnumRiskPolicySetTriggerType{
	"POLICY_SET_STAGING",
}

func (v *EnumRiskPolicySetTriggerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumRiskPolicySetTriggerType(value)
	for _, existing := range AllowedEnumRiskPolicySetTriggerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumRiskPolicySetTriggerType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumRiskPolicySetTriggerTypeFromValue returns a pointer to a valid EnumRiskPolicySetTriggerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumRiskPolicySetTriggerTypeFromValue(v string) (*EnumRiskPolicySetTriggerType, error) {
	ev := EnumRiskPolicySetTriggerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumRiskPolicySetTriggerType: valid values are %v", v, AllowedEnumRiskPolicySetTriggerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumRiskPolicySetTriggerType) IsValid() bool {
	for _, existing := range AllowedEnumRiskPolicySetTriggerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumRiskPolicySetTriggerType value
func (v EnumRiskPolicySetTriggerType) Ptr() *EnumRiskPolicySetTriggerType {
	return &v
}

type NullableEnumRiskPolicySetTriggerType struct {
	value *EnumRiskPolicySetTriggerType
	isSet bool
}

func (v NullableEnumRiskPolicySetTriggerType) Get() *EnumRiskPolicySetTriggerType {
	return v.value
}

func (v *NullableEnumRiskPolicySetTriggerType) Set(val *EnumRiskPolicySetTriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumRiskPolicySetTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumRiskPolicySetTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumRiskPolicySetTriggerType(val *EnumRiskPolicySetTriggerType) *NullableEnumRiskPolicySetTriggerType {
	return &NullableEnumRiskPolicySetTriggerType{value: val, isSet: true}
}

func (v NullableEnumRiskPolicySetTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumRiskPolicySetTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

