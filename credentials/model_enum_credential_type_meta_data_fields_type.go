/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"fmt"
)

// EnumCredentialTypeMetaDataFieldsType Specifies the type of data in the defined field within a credential type: `Alphanumeric Text` - static text string of letters, numbers, or punctuation set on fields.value `Issued Timestamp` - date and time the credential was issued `Directory Attribute` - any PingOne Directory standard or custom attribute. 
type EnumCredentialTypeMetaDataFieldsType string

// List of EnumCredentialTypeMetaDataFieldsType
const (
	ENUMCREDENTIALTYPEMETADATAFIELDSTYPE_ALPHANUMERIC_TEXT EnumCredentialTypeMetaDataFieldsType = "Alphanumeric Text"
	ENUMCREDENTIALTYPEMETADATAFIELDSTYPE_ISSUED_TIMESTAMP EnumCredentialTypeMetaDataFieldsType = "Issued Timestamp"
	ENUMCREDENTIALTYPEMETADATAFIELDSTYPE_DIRECTORY_ATTRIBUTE EnumCredentialTypeMetaDataFieldsType = "Directory Attribute"
)

// All allowed values of EnumCredentialTypeMetaDataFieldsType enum
var AllowedEnumCredentialTypeMetaDataFieldsTypeEnumValues = []EnumCredentialTypeMetaDataFieldsType{
	"Alphanumeric Text",
	"Issued Timestamp",
	"Directory Attribute",
}

func (v *EnumCredentialTypeMetaDataFieldsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCredentialTypeMetaDataFieldsType(value)
	for _, existing := range AllowedEnumCredentialTypeMetaDataFieldsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCredentialTypeMetaDataFieldsType", value)
}

// NewEnumCredentialTypeMetaDataFieldsTypeFromValue returns a pointer to a valid EnumCredentialTypeMetaDataFieldsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCredentialTypeMetaDataFieldsTypeFromValue(v string) (*EnumCredentialTypeMetaDataFieldsType, error) {
	ev := EnumCredentialTypeMetaDataFieldsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCredentialTypeMetaDataFieldsType: valid values are %v", v, AllowedEnumCredentialTypeMetaDataFieldsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCredentialTypeMetaDataFieldsType) IsValid() bool {
	for _, existing := range AllowedEnumCredentialTypeMetaDataFieldsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCredentialTypeMetaDataFieldsType value
func (v EnumCredentialTypeMetaDataFieldsType) Ptr() *EnumCredentialTypeMetaDataFieldsType {
	return &v
}

type NullableEnumCredentialTypeMetaDataFieldsType struct {
	value *EnumCredentialTypeMetaDataFieldsType
	isSet bool
}

func (v NullableEnumCredentialTypeMetaDataFieldsType) Get() *EnumCredentialTypeMetaDataFieldsType {
	return v.value
}

func (v *NullableEnumCredentialTypeMetaDataFieldsType) Set(val *EnumCredentialTypeMetaDataFieldsType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCredentialTypeMetaDataFieldsType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCredentialTypeMetaDataFieldsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCredentialTypeMetaDataFieldsType(val *EnumCredentialTypeMetaDataFieldsType) *NullableEnumCredentialTypeMetaDataFieldsType {
	return &NullableEnumCredentialTypeMetaDataFieldsType{value: val, isSet: true}
}

func (v NullableEnumCredentialTypeMetaDataFieldsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCredentialTypeMetaDataFieldsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

