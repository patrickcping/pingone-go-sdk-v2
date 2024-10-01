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

// EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo the model 'EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo'
type EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo string

// List of EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo
const (
	ENUMAUTHORIZEEDITORDATASTATEMENTSSTATEMENTDTOAPPLIESTO_ANYTHING EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo = "ANYTHING"
	ENUMAUTHORIZEEDITORDATASTATEMENTSSTATEMENTDTOAPPLIESTO_PERMIT EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo = "PERMIT"
	ENUMAUTHORIZEEDITORDATASTATEMENTSSTATEMENTDTOAPPLIESTO_DENY EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo = "DENY"
	ENUMAUTHORIZEEDITORDATASTATEMENTSSTATEMENTDTOAPPLIESTO_PERMIT_OR_DENY EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo = "PERMIT_OR_DENY"
	ENUMAUTHORIZEEDITORDATASTATEMENTSSTATEMENTDTOAPPLIESTO_INDETERMINATE EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo = "INDETERMINATE"
)

// All allowed values of EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo enum
var AllowedEnumAuthorizeEditorDataStatementsStatementDTOAppliesToEnumValues = []EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo{
	"ANYTHING",
	"PERMIT",
	"DENY",
	"PERMIT_OR_DENY",
	"INDETERMINATE",
}

func (v *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataStatementsStatementDTOAppliesToEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataStatementsStatementDTOAppliesToFromValue returns a pointer to a valid EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataStatementsStatementDTOAppliesToFromValue(v string) (*EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo, error) {
	ev := EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo: valid values are %v", v, AllowedEnumAuthorizeEditorDataStatementsStatementDTOAppliesToEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataStatementsStatementDTOAppliesToEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo value
func (v EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) Ptr() *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo {
	return &v
}

type NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo struct {
	value *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) Get() *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) Set(val *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo(val *EnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) *NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo {
	return &NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataStatementsStatementDTOAppliesTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

