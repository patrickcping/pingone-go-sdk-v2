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

// EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf the model 'EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf'
type EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf string

// List of EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf
const (
	ENUMAUTHORIZEEDITORDATASTATEMENTSREFERENCEABLESTATEMENTDTOAPPLIESIF_ANYTHING EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf = "ANYTHING"
	ENUMAUTHORIZEEDITORDATASTATEMENTSREFERENCEABLESTATEMENTDTOAPPLIESIF_FINAL_DECISION_MATCHES EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf = "FINAL_DECISION_MATCHES"
	ENUMAUTHORIZEEDITORDATASTATEMENTSREFERENCEABLESTATEMENTDTOAPPLIESIF_PATH_MATCHES EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf = "PATH_MATCHES"
)

// All allowed values of EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf enum
var AllowedEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfEnumValues = []EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf{
	"ANYTHING",
	"FINAL_DECISION_MATCHES",
	"PATH_MATCHES",
}

func (v *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfFromValue returns a pointer to a valid EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfFromValue(v string) (*EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf, error) {
	ev := EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf: valid values are %v", v, AllowedEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf value
func (v EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) Ptr() *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf {
	return &v
}

type NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf struct {
	value *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) Get() *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) Set(val *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf(val *EnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) *NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf {
	return &NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataStatementsReferenceableStatementDTOAppliesIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

