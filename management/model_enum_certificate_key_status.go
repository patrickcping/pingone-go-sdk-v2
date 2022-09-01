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

// EnumCertificateKeyStatus Specifies the status of the key. Options are `VALID`, `EXPIRING`, `EXPIRED`, `NOT_YET_VALID`, and `REVOKED`.
type EnumCertificateKeyStatus string

// List of EnumCertificateKeyStatus
const (
	ENUMCERTIFICATEKEYSTATUS_VALID EnumCertificateKeyStatus = "VALID"
	ENUMCERTIFICATEKEYSTATUS_EXPIRING EnumCertificateKeyStatus = "EXPIRING"
	ENUMCERTIFICATEKEYSTATUS_EXPIRED EnumCertificateKeyStatus = "EXPIRED"
	ENUMCERTIFICATEKEYSTATUS_NOT_YET_VALID EnumCertificateKeyStatus = "NOT_YET_VALID"
	ENUMCERTIFICATEKEYSTATUS_REVOKED EnumCertificateKeyStatus = "REVOKED"
)

// All allowed values of EnumCertificateKeyStatus enum
var AllowedEnumCertificateKeyStatusEnumValues = []EnumCertificateKeyStatus{
	"VALID",
	"EXPIRING",
	"EXPIRED",
	"NOT_YET_VALID",
	"REVOKED",
}

func (v *EnumCertificateKeyStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCertificateKeyStatus(value)
	for _, existing := range AllowedEnumCertificateKeyStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCertificateKeyStatus", value)
}

// NewEnumCertificateKeyStatusFromValue returns a pointer to a valid EnumCertificateKeyStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCertificateKeyStatusFromValue(v string) (*EnumCertificateKeyStatus, error) {
	ev := EnumCertificateKeyStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCertificateKeyStatus: valid values are %v", v, AllowedEnumCertificateKeyStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCertificateKeyStatus) IsValid() bool {
	for _, existing := range AllowedEnumCertificateKeyStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCertificateKeyStatus value
func (v EnumCertificateKeyStatus) Ptr() *EnumCertificateKeyStatus {
	return &v
}

type NullableEnumCertificateKeyStatus struct {
	value *EnumCertificateKeyStatus
	isSet bool
}

func (v NullableEnumCertificateKeyStatus) Get() *EnumCertificateKeyStatus {
	return v.value
}

func (v *NullableEnumCertificateKeyStatus) Set(val *EnumCertificateKeyStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCertificateKeyStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCertificateKeyStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCertificateKeyStatus(val *EnumCertificateKeyStatus) *NullableEnumCertificateKeyStatus {
	return &NullableEnumCertificateKeyStatus{value: val, isSet: true}
}

func (v NullableEnumCertificateKeyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCertificateKeyStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

