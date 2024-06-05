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

// EnumCertificateKeyEncryptionAlgorithm To replace
type EnumCertificateKeyEncryptionAlgorithm string

// List of EnumCertificateKeyEncryptionAlgorithm
const (
	ENUMCERTIFICATEKEYENCRYPTIONALGORITHM_AES_256 EnumCertificateKeyEncryptionAlgorithm = "AES_256"
)

// All allowed values of EnumCertificateKeyEncryptionAlgorithm enum
var AllowedEnumCertificateKeyEncryptionAlgorithmEnumValues = []EnumCertificateKeyEncryptionAlgorithm{
	"AES_256",
}

func (v *EnumCertificateKeyEncryptionAlgorithm) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCertificateKeyEncryptionAlgorithm(value)
	for _, existing := range AllowedEnumCertificateKeyEncryptionAlgorithmEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumCertificateKeyEncryptionAlgorithm(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumCertificateKeyEncryptionAlgorithmFromValue returns a pointer to a valid EnumCertificateKeyEncryptionAlgorithm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCertificateKeyEncryptionAlgorithmFromValue(v string) (*EnumCertificateKeyEncryptionAlgorithm, error) {
	ev := EnumCertificateKeyEncryptionAlgorithm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCertificateKeyEncryptionAlgorithm: valid values are %v", v, AllowedEnumCertificateKeyEncryptionAlgorithmEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCertificateKeyEncryptionAlgorithm) IsValid() bool {
	for _, existing := range AllowedEnumCertificateKeyEncryptionAlgorithmEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCertificateKeyEncryptionAlgorithm value
func (v EnumCertificateKeyEncryptionAlgorithm) Ptr() *EnumCertificateKeyEncryptionAlgorithm {
	return &v
}

type NullableEnumCertificateKeyEncryptionAlgorithm struct {
	value *EnumCertificateKeyEncryptionAlgorithm
	isSet bool
}

func (v NullableEnumCertificateKeyEncryptionAlgorithm) Get() *EnumCertificateKeyEncryptionAlgorithm {
	return v.value
}

func (v *NullableEnumCertificateKeyEncryptionAlgorithm) Set(val *EnumCertificateKeyEncryptionAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCertificateKeyEncryptionAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCertificateKeyEncryptionAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCertificateKeyEncryptionAlgorithm(val *EnumCertificateKeyEncryptionAlgorithm) *NullableEnumCertificateKeyEncryptionAlgorithm {
	return &NullableEnumCertificateKeyEncryptionAlgorithm{value: val, isSet: true}
}

func (v NullableEnumCertificateKeyEncryptionAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCertificateKeyEncryptionAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
