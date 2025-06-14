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

// EnumCSRExportHeader the model 'EnumCSRExportHeader'
type EnumCSRExportHeader string

// List of EnumCSRExportHeader
const (
	ENUMCSREXPORTHEADER_PKCS10     EnumCSRExportHeader = "application/pkcs10"
	ENUMCSREXPORTHEADER_X_PEM_FILE EnumCSRExportHeader = "application/x-pem-file"
)

// All allowed values of EnumCSRExportHeader enum
var AllowedEnumCSRExportHeaderEnumValues = []EnumCSRExportHeader{
	"application/pkcs10",
	"application/x-pem-file",
}

func (v *EnumCSRExportHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCSRExportHeader(value)
	for _, existing := range AllowedEnumCSRExportHeaderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumCSRExportHeader(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumCSRExportHeaderFromValue returns a pointer to a valid EnumCSRExportHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCSRExportHeaderFromValue(v string) (*EnumCSRExportHeader, error) {
	ev := EnumCSRExportHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCSRExportHeader: valid values are %v", v, AllowedEnumCSRExportHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCSRExportHeader) IsValid() bool {
	for _, existing := range AllowedEnumCSRExportHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCSRExportHeader value
func (v EnumCSRExportHeader) Ptr() *EnumCSRExportHeader {
	return &v
}

type NullableEnumCSRExportHeader struct {
	value *EnumCSRExportHeader
	isSet bool
}

func (v NullableEnumCSRExportHeader) Get() *EnumCSRExportHeader {
	return v.value
}

func (v *NullableEnumCSRExportHeader) Set(val *EnumCSRExportHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCSRExportHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCSRExportHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCSRExportHeader(val *EnumCSRExportHeader) *NullableEnumCSRExportHeader {
	return &NullableEnumCSRExportHeader{value: val, isSet: true}
}

func (v NullableEnumCSRExportHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCSRExportHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
