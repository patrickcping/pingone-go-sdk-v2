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

// EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType the model 'EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType'
type EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType string

// List of EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType
const (
	ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOTYPE_NONE EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType = "NONE"
	ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOTYPE_HTTP EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType = "HTTP"
	ENUMAUTHORIZEEDITORDATADEFINITIONSSERVICEDEFINITIONDTOTYPE_CONNECTOR EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType = "CONNECTOR"
)

// All allowed values of EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType enum
var AllowedEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeEnumValues = []EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType{
	"NONE",
	"HTTP",
	"CONNECTOR",
}

func (v *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeFromValue returns a pointer to a valid EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeFromValue(v string) (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType, error) {
	ev := EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType: valid values are %v", v, AllowedEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType value
func (v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) Ptr() *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType {
	return &v
}

type NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType struct {
	value *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) Get() *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) Set(val *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType(val *EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) *NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType {
	return &NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

