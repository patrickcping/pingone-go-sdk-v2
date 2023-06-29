/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// EnumFIDO2PolicyUserVerificationOption Can be one of the following values: - `REQUIRED` - only FIDO devices supporting user verification can be used - `DISCOURAGED` - user verification is not required, even when supported by the FIDO device. In cases where user verification is required by the FIDO device itself, this setting does not override the device setting. - `PREFERRED` - user verification is required if the user's FIDO device supports it, but is not required if the user's device does not support it. For usernameless flows, only FIDO devices supporting user verification can be used, regardless of the value you set for `userVerification.option`. 
type EnumFIDO2PolicyUserVerificationOption string

// List of EnumFIDO2PolicyUserVerificationOption
const (
	ENUMFIDO2POLICYUSERVERIFICATIONOPTION_REQUIRED EnumFIDO2PolicyUserVerificationOption = "REQUIRED"
	ENUMFIDO2POLICYUSERVERIFICATIONOPTION_DISCOURAGED EnumFIDO2PolicyUserVerificationOption = "DISCOURAGED"
	ENUMFIDO2POLICYUSERVERIFICATIONOPTION_PREFERRED EnumFIDO2PolicyUserVerificationOption = "PREFERRED"
)

// All allowed values of EnumFIDO2PolicyUserVerificationOption enum
var AllowedEnumFIDO2PolicyUserVerificationOptionEnumValues = []EnumFIDO2PolicyUserVerificationOption{
	"REQUIRED",
	"DISCOURAGED",
	"PREFERRED",
}

func (v *EnumFIDO2PolicyUserVerificationOption) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumFIDO2PolicyUserVerificationOption(value)
	for _, existing := range AllowedEnumFIDO2PolicyUserVerificationOptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumFIDO2PolicyUserVerificationOption", value)
}

// NewEnumFIDO2PolicyUserVerificationOptionFromValue returns a pointer to a valid EnumFIDO2PolicyUserVerificationOption
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumFIDO2PolicyUserVerificationOptionFromValue(v string) (*EnumFIDO2PolicyUserVerificationOption, error) {
	ev := EnumFIDO2PolicyUserVerificationOption(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumFIDO2PolicyUserVerificationOption: valid values are %v", v, AllowedEnumFIDO2PolicyUserVerificationOptionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumFIDO2PolicyUserVerificationOption) IsValid() bool {
	for _, existing := range AllowedEnumFIDO2PolicyUserVerificationOptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumFIDO2PolicyUserVerificationOption value
func (v EnumFIDO2PolicyUserVerificationOption) Ptr() *EnumFIDO2PolicyUserVerificationOption {
	return &v
}

type NullableEnumFIDO2PolicyUserVerificationOption struct {
	value *EnumFIDO2PolicyUserVerificationOption
	isSet bool
}

func (v NullableEnumFIDO2PolicyUserVerificationOption) Get() *EnumFIDO2PolicyUserVerificationOption {
	return v.value
}

func (v *NullableEnumFIDO2PolicyUserVerificationOption) Set(val *EnumFIDO2PolicyUserVerificationOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumFIDO2PolicyUserVerificationOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumFIDO2PolicyUserVerificationOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumFIDO2PolicyUserVerificationOption(val *EnumFIDO2PolicyUserVerificationOption) *NullableEnumFIDO2PolicyUserVerificationOption {
	return &NullableEnumFIDO2PolicyUserVerificationOption{value: val, isSet: true}
}

func (v NullableEnumFIDO2PolicyUserVerificationOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumFIDO2PolicyUserVerificationOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

