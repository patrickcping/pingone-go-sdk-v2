/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumCSRResponseImportHeader the model 'EnumCSRResponseImportHeader'
type EnumCSRResponseImportHeader string

// List of EnumCSRResponseImportHeader
const (
	ENUMCSRRESPONSEIMPORTHEADER_APPLICATION_X_PEM_FILE EnumCSRResponseImportHeader = "application/x-pem-file"
)

// All allowed values of EnumCSRResponseImportHeader enum
var AllowedEnumCSRResponseImportHeaderEnumValues = []EnumCSRResponseImportHeader{
	"application/x-pem-file",
}

func (v *EnumCSRResponseImportHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCSRResponseImportHeader(value)
	for _, existing := range AllowedEnumCSRResponseImportHeaderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCSRResponseImportHeader", value)
}

// NewEnumCSRResponseImportHeaderFromValue returns a pointer to a valid EnumCSRResponseImportHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCSRResponseImportHeaderFromValue(v string) (*EnumCSRResponseImportHeader, error) {
	ev := EnumCSRResponseImportHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCSRResponseImportHeader: valid values are %v", v, AllowedEnumCSRResponseImportHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCSRResponseImportHeader) IsValid() bool {
	for _, existing := range AllowedEnumCSRResponseImportHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCSRResponseImportHeader value
func (v EnumCSRResponseImportHeader) Ptr() *EnumCSRResponseImportHeader {
	return &v
}

type NullableEnumCSRResponseImportHeader struct {
	value *EnumCSRResponseImportHeader
	isSet bool
}

func (v NullableEnumCSRResponseImportHeader) Get() *EnumCSRResponseImportHeader {
	return v.value
}

func (v *NullableEnumCSRResponseImportHeader) Set(val *EnumCSRResponseImportHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCSRResponseImportHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCSRResponseImportHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCSRResponseImportHeader(val *EnumCSRResponseImportHeader) *NullableEnumCSRResponseImportHeader {
	return &NullableEnumCSRResponseImportHeader{value: val, isSet: true}
}

func (v NullableEnumCSRResponseImportHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCSRResponseImportHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

