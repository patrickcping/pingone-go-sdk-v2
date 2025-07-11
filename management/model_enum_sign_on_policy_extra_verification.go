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

// EnumSignOnPolicyExtraVerification Specifies the level of further verification when deviceAuthorization is enabled. The PingOne platform performs an extra verification check by sending a silent push notification to the customer native application, and receives a confirmation in return.  `extraVerification` can be one of the following levels: `disabled` (default): The PingOne platform does not perform the extra verification check. `permissive`: The PingOne platform performs the extra verification check. Upon timeout or failure to get a response from the native app, the MFA step is treated as successfully completed. `restrictive`: The PingOne platform performs the extra verification check.The PingOne platform performs the extra verification check. Upon timeout or failure to get a response from the native app, the MFA step is treated as failed.
type EnumSignOnPolicyExtraVerification string

// List of EnumSignOnPolicyExtraVerification
const (
	ENUMSIGNONPOLICYEXTRAVERIFICATION_DISABLED    EnumSignOnPolicyExtraVerification = "disabled"
	ENUMSIGNONPOLICYEXTRAVERIFICATION_PERMISSIVE  EnumSignOnPolicyExtraVerification = "permissive"
	ENUMSIGNONPOLICYEXTRAVERIFICATION_RESTRICTIVE EnumSignOnPolicyExtraVerification = "restrictive"
)

// All allowed values of EnumSignOnPolicyExtraVerification enum
var AllowedEnumSignOnPolicyExtraVerificationEnumValues = []EnumSignOnPolicyExtraVerification{
	"disabled",
	"permissive",
	"restrictive",
}

func (v *EnumSignOnPolicyExtraVerification) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSignOnPolicyExtraVerification(value)
	for _, existing := range AllowedEnumSignOnPolicyExtraVerificationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumSignOnPolicyExtraVerification(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumSignOnPolicyExtraVerificationFromValue returns a pointer to a valid EnumSignOnPolicyExtraVerification
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSignOnPolicyExtraVerificationFromValue(v string) (*EnumSignOnPolicyExtraVerification, error) {
	ev := EnumSignOnPolicyExtraVerification(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSignOnPolicyExtraVerification: valid values are %v", v, AllowedEnumSignOnPolicyExtraVerificationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSignOnPolicyExtraVerification) IsValid() bool {
	for _, existing := range AllowedEnumSignOnPolicyExtraVerificationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSignOnPolicyExtraVerification value
func (v EnumSignOnPolicyExtraVerification) Ptr() *EnumSignOnPolicyExtraVerification {
	return &v
}

type NullableEnumSignOnPolicyExtraVerification struct {
	value *EnumSignOnPolicyExtraVerification
	isSet bool
}

func (v NullableEnumSignOnPolicyExtraVerification) Get() *EnumSignOnPolicyExtraVerification {
	return v.value
}

func (v *NullableEnumSignOnPolicyExtraVerification) Set(val *EnumSignOnPolicyExtraVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSignOnPolicyExtraVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSignOnPolicyExtraVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSignOnPolicyExtraVerification(val *EnumSignOnPolicyExtraVerification) *NullableEnumSignOnPolicyExtraVerification {
	return &NullableEnumSignOnPolicyExtraVerification{value: val, isSet: true}
}

func (v NullableEnumSignOnPolicyExtraVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSignOnPolicyExtraVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
