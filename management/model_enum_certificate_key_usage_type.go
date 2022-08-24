/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// EnumCertificateKeyUsageType Specifies how the certificate is used. Options are ENCRYPTION and SIGNING.
type EnumCertificateKeyUsageType string

// List of EnumCertificateKeyUsageType
const (
	ENUMCERTIFICATEKEYUSAGETYPE_ENCRYPTION EnumCertificateKeyUsageType = "ENCRYPTION"
	ENUMCERTIFICATEKEYUSAGETYPE_SIGNING EnumCertificateKeyUsageType = "SIGNING"
)

// All allowed values of EnumCertificateKeyUsageType enum
var AllowedEnumCertificateKeyUsageTypeEnumValues = []EnumCertificateKeyUsageType{
	"ENCRYPTION",
	"SIGNING",
}

func (v *EnumCertificateKeyUsageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCertificateKeyUsageType(value)
	for _, existing := range AllowedEnumCertificateKeyUsageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCertificateKeyUsageType", value)
}

// NewEnumCertificateKeyUsageTypeFromValue returns a pointer to a valid EnumCertificateKeyUsageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCertificateKeyUsageTypeFromValue(v string) (*EnumCertificateKeyUsageType, error) {
	ev := EnumCertificateKeyUsageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCertificateKeyUsageType: valid values are %v", v, AllowedEnumCertificateKeyUsageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCertificateKeyUsageType) IsValid() bool {
	for _, existing := range AllowedEnumCertificateKeyUsageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCertificateKeyUsageType value
func (v EnumCertificateKeyUsageType) Ptr() *EnumCertificateKeyUsageType {
	return &v
}

type NullableEnumCertificateKeyUsageType struct {
	value *EnumCertificateKeyUsageType
	isSet bool
}

func (v NullableEnumCertificateKeyUsageType) Get() *EnumCertificateKeyUsageType {
	return v.value
}

func (v *NullableEnumCertificateKeyUsageType) Set(val *EnumCertificateKeyUsageType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCertificateKeyUsageType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCertificateKeyUsageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCertificateKeyUsageType(val *EnumCertificateKeyUsageType) *NullableEnumCertificateKeyUsageType {
	return &NullableEnumCertificateKeyUsageType{value: val, isSet: true}
}

func (v NullableEnumCertificateKeyUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCertificateKeyUsageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
