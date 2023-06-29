/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"fmt"
)

// EnumCredentialIssuanceRuleNotificationMethod Specifies the method for notifying the user; can be SMS, EMAIL or both.
type EnumCredentialIssuanceRuleNotificationMethod string

// List of EnumCredentialIssuanceRuleNotificationMethod
const (
	ENUMCREDENTIALISSUANCERULENOTIFICATIONMETHOD_SMS EnumCredentialIssuanceRuleNotificationMethod = "SMS"
	ENUMCREDENTIALISSUANCERULENOTIFICATIONMETHOD_EMAIL EnumCredentialIssuanceRuleNotificationMethod = "EMAIL"
)

// All allowed values of EnumCredentialIssuanceRuleNotificationMethod enum
var AllowedEnumCredentialIssuanceRuleNotificationMethodEnumValues = []EnumCredentialIssuanceRuleNotificationMethod{
	"SMS",
	"EMAIL",
}

func (v *EnumCredentialIssuanceRuleNotificationMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCredentialIssuanceRuleNotificationMethod(value)
	for _, existing := range AllowedEnumCredentialIssuanceRuleNotificationMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCredentialIssuanceRuleNotificationMethod", value)
}

// NewEnumCredentialIssuanceRuleNotificationMethodFromValue returns a pointer to a valid EnumCredentialIssuanceRuleNotificationMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCredentialIssuanceRuleNotificationMethodFromValue(v string) (*EnumCredentialIssuanceRuleNotificationMethod, error) {
	ev := EnumCredentialIssuanceRuleNotificationMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCredentialIssuanceRuleNotificationMethod: valid values are %v", v, AllowedEnumCredentialIssuanceRuleNotificationMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCredentialIssuanceRuleNotificationMethod) IsValid() bool {
	for _, existing := range AllowedEnumCredentialIssuanceRuleNotificationMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCredentialIssuanceRuleNotificationMethod value
func (v EnumCredentialIssuanceRuleNotificationMethod) Ptr() *EnumCredentialIssuanceRuleNotificationMethod {
	return &v
}

type NullableEnumCredentialIssuanceRuleNotificationMethod struct {
	value *EnumCredentialIssuanceRuleNotificationMethod
	isSet bool
}

func (v NullableEnumCredentialIssuanceRuleNotificationMethod) Get() *EnumCredentialIssuanceRuleNotificationMethod {
	return v.value
}

func (v *NullableEnumCredentialIssuanceRuleNotificationMethod) Set(val *EnumCredentialIssuanceRuleNotificationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCredentialIssuanceRuleNotificationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCredentialIssuanceRuleNotificationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCredentialIssuanceRuleNotificationMethod(val *EnumCredentialIssuanceRuleNotificationMethod) *NullableEnumCredentialIssuanceRuleNotificationMethod {
	return &NullableEnumCredentialIssuanceRuleNotificationMethod{value: val, isSet: true}
}

func (v NullableEnumCredentialIssuanceRuleNotificationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCredentialIssuanceRuleNotificationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

