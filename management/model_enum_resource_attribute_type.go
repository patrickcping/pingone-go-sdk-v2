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

// EnumResourceAttributeType A string that specifies the type of resource attribute. Options are: CORE: The claim is required and cannot not be removed. CUSTOM: The claim is not a CORE attribute. All created attributes are of this type. PREDEFINED: A designation for predefined OIDC resource attributes such as given_name. These attributes cannot be removed; however, they can be modified. 
type EnumResourceAttributeType string

// List of EnumResourceAttributeType
const (
	ENUMRESOURCEATTRIBUTETYPE_CORE EnumResourceAttributeType = "CORE"
	ENUMRESOURCEATTRIBUTETYPE_CUSTOM EnumResourceAttributeType = "CUSTOM"
	ENUMRESOURCEATTRIBUTETYPE_PREDEFINED EnumResourceAttributeType = "PREDEFINED"
)

// All allowed values of EnumResourceAttributeType enum
var AllowedEnumResourceAttributeTypeEnumValues = []EnumResourceAttributeType{
	"CORE",
	"CUSTOM",
	"PREDEFINED",
}

func (v *EnumResourceAttributeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumResourceAttributeType(value)
	for _, existing := range AllowedEnumResourceAttributeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumResourceAttributeType", value)
}

// NewEnumResourceAttributeTypeFromValue returns a pointer to a valid EnumResourceAttributeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumResourceAttributeTypeFromValue(v string) (*EnumResourceAttributeType, error) {
	ev := EnumResourceAttributeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumResourceAttributeType: valid values are %v", v, AllowedEnumResourceAttributeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumResourceAttributeType) IsValid() bool {
	for _, existing := range AllowedEnumResourceAttributeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumResourceAttributeType value
func (v EnumResourceAttributeType) Ptr() *EnumResourceAttributeType {
	return &v
}

type NullableEnumResourceAttributeType struct {
	value *EnumResourceAttributeType
	isSet bool
}

func (v NullableEnumResourceAttributeType) Get() *EnumResourceAttributeType {
	return v.value
}

func (v *NullableEnumResourceAttributeType) Set(val *EnumResourceAttributeType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumResourceAttributeType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumResourceAttributeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumResourceAttributeType(val *EnumResourceAttributeType) *NullableEnumResourceAttributeType {
	return &NullableEnumResourceAttributeType{value: val, isSet: true}
}

func (v NullableEnumResourceAttributeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumResourceAttributeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

