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

// EnumAdministratorSecurityMfaStatus This applies only to the specified environment, and can be either `OPT_IN` (indicating MFA is to be used for administrator sign-ons), or `OPT_OUT` (indicating MFA is not to be used for administrator sign-ons). This currently defaults to `OPT_OUT`.
type EnumAdministratorSecurityMfaStatus string

// List of EnumAdministratorSecurityMfaStatus
const (
	ENUMADMINISTRATORSECURITYMFASTATUS_IN EnumAdministratorSecurityMfaStatus = "OPT_IN"
	ENUMADMINISTRATORSECURITYMFASTATUS_OUT EnumAdministratorSecurityMfaStatus = "OPT_OUT"
)

// All allowed values of EnumAdministratorSecurityMfaStatus enum
var AllowedEnumAdministratorSecurityMfaStatusEnumValues = []EnumAdministratorSecurityMfaStatus{
	"OPT_IN",
	"OPT_OUT",
}

func (v *EnumAdministratorSecurityMfaStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAdministratorSecurityMfaStatus(value)
	for _, existing := range AllowedEnumAdministratorSecurityMfaStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAdministratorSecurityMfaStatus(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAdministratorSecurityMfaStatusFromValue returns a pointer to a valid EnumAdministratorSecurityMfaStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAdministratorSecurityMfaStatusFromValue(v string) (*EnumAdministratorSecurityMfaStatus, error) {
	ev := EnumAdministratorSecurityMfaStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAdministratorSecurityMfaStatus: valid values are %v", v, AllowedEnumAdministratorSecurityMfaStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAdministratorSecurityMfaStatus) IsValid() bool {
	for _, existing := range AllowedEnumAdministratorSecurityMfaStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAdministratorSecurityMfaStatus value
func (v EnumAdministratorSecurityMfaStatus) Ptr() *EnumAdministratorSecurityMfaStatus {
	return &v
}

type NullableEnumAdministratorSecurityMfaStatus struct {
	value *EnumAdministratorSecurityMfaStatus
	isSet bool
}

func (v NullableEnumAdministratorSecurityMfaStatus) Get() *EnumAdministratorSecurityMfaStatus {
	return v.value
}

func (v *NullableEnumAdministratorSecurityMfaStatus) Set(val *EnumAdministratorSecurityMfaStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAdministratorSecurityMfaStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAdministratorSecurityMfaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAdministratorSecurityMfaStatus(val *EnumAdministratorSecurityMfaStatus) *NullableEnumAdministratorSecurityMfaStatus {
	return &NullableEnumAdministratorSecurityMfaStatus{value: val, isSet: true}
}

func (v NullableEnumAdministratorSecurityMfaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAdministratorSecurityMfaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
