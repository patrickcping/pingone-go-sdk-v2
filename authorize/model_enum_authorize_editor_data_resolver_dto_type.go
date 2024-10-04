/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// EnumAuthorizeEditorDataResolverDTOType the model 'EnumAuthorizeEditorDataResolverDTOType'
type EnumAuthorizeEditorDataResolverDTOType string

// List of EnumAuthorizeEditorDataResolverDTOType
const (
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_ATTRIBUTE EnumAuthorizeEditorDataResolverDTOType = "ATTRIBUTE"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CONSTANT EnumAuthorizeEditorDataResolverDTOType = "CONSTANT"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CURRENT_REPETITION_VALUE EnumAuthorizeEditorDataResolverDTOType = "CURRENT_REPETITION_VALUE"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_CURRENT_USER_ID EnumAuthorizeEditorDataResolverDTOType = "CURRENT_USER_ID"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_REQUEST EnumAuthorizeEditorDataResolverDTOType = "REQUEST"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_SERVICE EnumAuthorizeEditorDataResolverDTOType = "SERVICE"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_SYSTEM EnumAuthorizeEditorDataResolverDTOType = "SYSTEM"
	ENUMAUTHORIZEEDITORDATARESOLVERDTOTYPE_USER EnumAuthorizeEditorDataResolverDTOType = "USER"
)

// All allowed values of EnumAuthorizeEditorDataResolverDTOType enum
var AllowedEnumAuthorizeEditorDataResolverDTOTypeEnumValues = []EnumAuthorizeEditorDataResolverDTOType{
	"ATTRIBUTE",
	"CONSTANT",
	"CURRENT_REPETITION_VALUE",
	"CURRENT_USER_ID",
	"REQUEST",
	"SERVICE",
	"SYSTEM",
	"USER",
}

func (v *EnumAuthorizeEditorDataResolverDTOType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataResolverDTOType(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataResolverDTOTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataResolverDTOType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataResolverDTOTypeFromValue returns a pointer to a valid EnumAuthorizeEditorDataResolverDTOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataResolverDTOTypeFromValue(v string) (*EnumAuthorizeEditorDataResolverDTOType, error) {
	ev := EnumAuthorizeEditorDataResolverDTOType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataResolverDTOType: valid values are %v", v, AllowedEnumAuthorizeEditorDataResolverDTOTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataResolverDTOType) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataResolverDTOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataResolverDTOType value
func (v EnumAuthorizeEditorDataResolverDTOType) Ptr() *EnumAuthorizeEditorDataResolverDTOType {
	return &v
}

type NullableEnumAuthorizeEditorDataResolverDTOType struct {
	value *EnumAuthorizeEditorDataResolverDTOType
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataResolverDTOType) Get() *EnumAuthorizeEditorDataResolverDTOType {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataResolverDTOType) Set(val *EnumAuthorizeEditorDataResolverDTOType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataResolverDTOType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataResolverDTOType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataResolverDTOType(val *EnumAuthorizeEditorDataResolverDTOType) *NullableEnumAuthorizeEditorDataResolverDTOType {
	return &NullableEnumAuthorizeEditorDataResolverDTOType{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataResolverDTOType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataResolverDTOType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

