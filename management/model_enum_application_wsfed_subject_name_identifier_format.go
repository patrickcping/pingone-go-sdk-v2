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

// EnumApplicationWSFEDSubjectNameIdentifierFormat The format to use for the SubjectNameIdentifier element.
type EnumApplicationWSFEDSubjectNameIdentifierFormat string

// List of EnumApplicationWSFEDSubjectNameIdentifierFormat
const (
	ENUMAPPLICATIONWSFEDSUBJECTNAMEIDENTIFIERFORMAT_UNSPECIFIED   EnumApplicationWSFEDSubjectNameIdentifierFormat = "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"
	ENUMAPPLICATIONWSFEDSUBJECTNAMEIDENTIFIERFORMAT_EMAIL_ADDRESS EnumApplicationWSFEDSubjectNameIdentifierFormat = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
)

// All allowed values of EnumApplicationWSFEDSubjectNameIdentifierFormat enum
var AllowedEnumApplicationWSFEDSubjectNameIdentifierFormatEnumValues = []EnumApplicationWSFEDSubjectNameIdentifierFormat{
	"urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified",
	"urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress",
}

func (v *EnumApplicationWSFEDSubjectNameIdentifierFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumApplicationWSFEDSubjectNameIdentifierFormat(value)
	for _, existing := range AllowedEnumApplicationWSFEDSubjectNameIdentifierFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumApplicationWSFEDSubjectNameIdentifierFormat(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumApplicationWSFEDSubjectNameIdentifierFormatFromValue returns a pointer to a valid EnumApplicationWSFEDSubjectNameIdentifierFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumApplicationWSFEDSubjectNameIdentifierFormatFromValue(v string) (*EnumApplicationWSFEDSubjectNameIdentifierFormat, error) {
	ev := EnumApplicationWSFEDSubjectNameIdentifierFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumApplicationWSFEDSubjectNameIdentifierFormat: valid values are %v", v, AllowedEnumApplicationWSFEDSubjectNameIdentifierFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumApplicationWSFEDSubjectNameIdentifierFormat) IsValid() bool {
	for _, existing := range AllowedEnumApplicationWSFEDSubjectNameIdentifierFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumApplicationWSFEDSubjectNameIdentifierFormat value
func (v EnumApplicationWSFEDSubjectNameIdentifierFormat) Ptr() *EnumApplicationWSFEDSubjectNameIdentifierFormat {
	return &v
}

type NullableEnumApplicationWSFEDSubjectNameIdentifierFormat struct {
	value *EnumApplicationWSFEDSubjectNameIdentifierFormat
	isSet bool
}

func (v NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) Get() *EnumApplicationWSFEDSubjectNameIdentifierFormat {
	return v.value
}

func (v *NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) Set(val *EnumApplicationWSFEDSubjectNameIdentifierFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumApplicationWSFEDSubjectNameIdentifierFormat(val *EnumApplicationWSFEDSubjectNameIdentifierFormat) *NullableEnumApplicationWSFEDSubjectNameIdentifierFormat {
	return &NullableEnumApplicationWSFEDSubjectNameIdentifierFormat{value: val, isSet: true}
}

func (v NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumApplicationWSFEDSubjectNameIdentifierFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
