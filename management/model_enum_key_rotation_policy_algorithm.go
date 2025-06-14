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

// EnumKeyRotationPolicyAlgorithm The algorithm this KRP applies to generated `KrpKeys`. `RSA` is currently the only supported value.
type EnumKeyRotationPolicyAlgorithm string

// List of EnumKeyRotationPolicyAlgorithm
const (
	ENUMKEYROTATIONPOLICYALGORITHM_RSA EnumKeyRotationPolicyAlgorithm = "RSA"
)

// All allowed values of EnumKeyRotationPolicyAlgorithm enum
var AllowedEnumKeyRotationPolicyAlgorithmEnumValues = []EnumKeyRotationPolicyAlgorithm{
	"RSA",
}

func (v *EnumKeyRotationPolicyAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumKeyRotationPolicyAlgorithm(value)
	for _, existing := range AllowedEnumKeyRotationPolicyAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumKeyRotationPolicyAlgorithm(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumKeyRotationPolicyAlgorithmFromValue returns a pointer to a valid EnumKeyRotationPolicyAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumKeyRotationPolicyAlgorithmFromValue(v string) (*EnumKeyRotationPolicyAlgorithm, error) {
	ev := EnumKeyRotationPolicyAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumKeyRotationPolicyAlgorithm: valid values are %v", v, AllowedEnumKeyRotationPolicyAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumKeyRotationPolicyAlgorithm) IsValid() bool {
	for _, existing := range AllowedEnumKeyRotationPolicyAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumKeyRotationPolicyAlgorithm value
func (v EnumKeyRotationPolicyAlgorithm) Ptr() *EnumKeyRotationPolicyAlgorithm {
	return &v
}

type NullableEnumKeyRotationPolicyAlgorithm struct {
	value *EnumKeyRotationPolicyAlgorithm
	isSet bool
}

func (v NullableEnumKeyRotationPolicyAlgorithm) Get() *EnumKeyRotationPolicyAlgorithm {
	return v.value
}

func (v *NullableEnumKeyRotationPolicyAlgorithm) Set(val *EnumKeyRotationPolicyAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumKeyRotationPolicyAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumKeyRotationPolicyAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumKeyRotationPolicyAlgorithm(val *EnumKeyRotationPolicyAlgorithm) *NullableEnumKeyRotationPolicyAlgorithm {
	return &NullableEnumKeyRotationPolicyAlgorithm{value: val, isSet: true}
}

func (v NullableEnumKeyRotationPolicyAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumKeyRotationPolicyAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
