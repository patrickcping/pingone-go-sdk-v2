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

// EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod Should be set to EMAIL.
type EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod string

// List of EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod
const (
	ENUMNOTIFICATIONSSETTINGSEMAILDELIVERYSETTINGSCUSTOMREQUESTSDELIVERYMETHOD_EMAIL EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod = "EMAIL"
)

// All allowed values of EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod enum
var AllowedEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodEnumValues = []EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod{
	"EMAIL",
}

func (v *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod(value)
	for _, existing := range AllowedEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodFromValue returns a pointer to a valid EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodFromValue(v string) (*EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod, error) {
	ev := EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod: valid values are %v", v, AllowedEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod value
func (v EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) Ptr() *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod {
	return &v
}

type NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod struct {
	value *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod
	isSet bool
}

func (v NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) Get() *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod {
	return v.value
}

func (v *NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) Set(val *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod(val *EnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) *NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod {
	return &NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod{value: val, isSet: true}
}

func (v NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumNotificationsSettingsEmailDeliverySettingsCustomRequestsDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
