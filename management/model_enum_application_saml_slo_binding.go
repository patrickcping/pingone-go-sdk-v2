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

// EnumApplicationSAMLSloBinding A string that specifies the binding protocol to be used for the logout response. Options are HTTP_REDIRECT or HTTP_POST. The default is HTTP_POST; existing configurations with no data default to HTTP_POST.
type EnumApplicationSAMLSloBinding string

// List of EnumApplicationSAMLSloBinding
const (
	ENUMAPPLICATIONSAMLSLOBINDING_REDIRECT EnumApplicationSAMLSloBinding = "HTTP_REDIRECT"
	ENUMAPPLICATIONSAMLSLOBINDING_POST     EnumApplicationSAMLSloBinding = "HTTP_POST"
)

// All allowed values of EnumApplicationSAMLSloBinding enum
var AllowedEnumApplicationSAMLSloBindingEnumValues = []EnumApplicationSAMLSloBinding{
	"HTTP_REDIRECT",
	"HTTP_POST",
}

func (v *EnumApplicationSAMLSloBinding) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationSAMLSloBinding(value)
	for _, existing := range AllowedEnumApplicationSAMLSloBindingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumApplicationSAMLSloBinding(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumApplicationSAMLSloBindingFromValue returns a pointer to a valid EnumApplicationSAMLSloBinding
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationSAMLSloBindingFromValue(v string) (*EnumApplicationSAMLSloBinding, error) {
	ev := EnumApplicationSAMLSloBinding(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationSAMLSloBinding: valid values are %v", v, AllowedEnumApplicationSAMLSloBindingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationSAMLSloBinding) IsValid() bool {
	for _, existing := range AllowedEnumApplicationSAMLSloBindingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationSAMLSloBinding value
func (v EnumApplicationSAMLSloBinding) Ptr() *EnumApplicationSAMLSloBinding {
	return &v
}

type NullableEnumApplicationSAMLSloBinding struct {
	value *EnumApplicationSAMLSloBinding
	isSet bool
}

func (v NullableEnumApplicationSAMLSloBinding) Get() *EnumApplicationSAMLSloBinding {
	return v.value
}

func (v *NullableEnumApplicationSAMLSloBinding) Set(val *EnumApplicationSAMLSloBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationSAMLSloBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationSAMLSloBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationSAMLSloBinding(val *EnumApplicationSAMLSloBinding) *NullableEnumApplicationSAMLSloBinding {
	return &NullableEnumApplicationSAMLSloBinding{value: val, isSet: true}
}

func (v NullableEnumApplicationSAMLSloBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationSAMLSloBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
