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

// EnumFormRecaptchaV2Size A string that specifies the size of the reCAPTCHA widget.
type EnumFormRecaptchaV2Size string

// List of EnumFormRecaptchaV2Size
const (
	ENUMFORMRECAPTCHAV2SIZE_NORMAL  EnumFormRecaptchaV2Size = "NORMAL"
	ENUMFORMRECAPTCHAV2SIZE_COMPACT EnumFormRecaptchaV2Size = "COMPACT"
)

// All allowed values of EnumFormRecaptchaV2Size enum
var AllowedEnumFormRecaptchaV2SizeEnumValues = []EnumFormRecaptchaV2Size{
	"NORMAL",
	"COMPACT",
}

func (v *EnumFormRecaptchaV2Size) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFormRecaptchaV2Size(value)
	for _, existing := range AllowedEnumFormRecaptchaV2SizeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFormRecaptchaV2Size", value)
}

// NewEnumFormRecaptchaV2SizeFromValue returns a pointer to a valid EnumFormRecaptchaV2Size
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFormRecaptchaV2SizeFromValue(v string) (*EnumFormRecaptchaV2Size, error) {
	ev := EnumFormRecaptchaV2Size(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFormRecaptchaV2Size: valid values are %v", v, AllowedEnumFormRecaptchaV2SizeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFormRecaptchaV2Size) IsValid() bool {
	for _, existing := range AllowedEnumFormRecaptchaV2SizeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFormRecaptchaV2Size value
func (v EnumFormRecaptchaV2Size) Ptr() *EnumFormRecaptchaV2Size {
	return &v
}

type NullableEnumFormRecaptchaV2Size struct {
	value *EnumFormRecaptchaV2Size
	isSet bool
}

func (v NullableEnumFormRecaptchaV2Size) Get() *EnumFormRecaptchaV2Size {
	return v.value
}

func (v *NullableEnumFormRecaptchaV2Size) Set(val *EnumFormRecaptchaV2Size) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFormRecaptchaV2Size) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFormRecaptchaV2Size) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFormRecaptchaV2Size(val *EnumFormRecaptchaV2Size) *NullableEnumFormRecaptchaV2Size {
	return &NullableEnumFormRecaptchaV2Size{value: val, isSet: true}
}

func (v NullableEnumFormRecaptchaV2Size) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFormRecaptchaV2Size) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
