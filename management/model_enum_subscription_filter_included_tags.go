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

// EnumSubscriptionFilterIncludedTags A tags that events must have to be monitored by a subscription. Currently, the available tags are `adminIdentityEvent`. Identifies the event as the action of an administrator on other administrators.
type EnumSubscriptionFilterIncludedTags string

// List of EnumSubscriptionFilterIncludedTags
const (
	ENUMSUBSCRIPTIONFILTERINCLUDEDTAGS_ADMIN_IDENTITY_EVENT EnumSubscriptionFilterIncludedTags = "adminIdentityEvent"
)

// All allowed values of EnumSubscriptionFilterIncludedTags enum
var AllowedEnumSubscriptionFilterIncludedTagsEnumValues = []EnumSubscriptionFilterIncludedTags{
	"adminIdentityEvent",
}

func (v *EnumSubscriptionFilterIncludedTags) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumSubscriptionFilterIncludedTags(value)
	for _, existing := range AllowedEnumSubscriptionFilterIncludedTagsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumSubscriptionFilterIncludedTags", value)
}

// NewEnumSubscriptionFilterIncludedTagsFromValue returns a pointer to a valid EnumSubscriptionFilterIncludedTags
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumSubscriptionFilterIncludedTagsFromValue(v string) (*EnumSubscriptionFilterIncludedTags, error) {
	ev := EnumSubscriptionFilterIncludedTags(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumSubscriptionFilterIncludedTags: valid values are %v", v, AllowedEnumSubscriptionFilterIncludedTagsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumSubscriptionFilterIncludedTags) IsValid() bool {
	for _, existing := range AllowedEnumSubscriptionFilterIncludedTagsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumSubscriptionFilterIncludedTags value
func (v EnumSubscriptionFilterIncludedTags) Ptr() *EnumSubscriptionFilterIncludedTags {
	return &v
}

type NullableEnumSubscriptionFilterIncludedTags struct {
	value *EnumSubscriptionFilterIncludedTags
	isSet bool
}

func (v NullableEnumSubscriptionFilterIncludedTags) Get() *EnumSubscriptionFilterIncludedTags {
	return v.value
}

func (v *NullableEnumSubscriptionFilterIncludedTags) Set(val *EnumSubscriptionFilterIncludedTags) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumSubscriptionFilterIncludedTags) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumSubscriptionFilterIncludedTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumSubscriptionFilterIncludedTags(val *EnumSubscriptionFilterIncludedTags) *NullableEnumSubscriptionFilterIncludedTags {
	return &NullableEnumSubscriptionFilterIncludedTags{value: val, isSet: true}
}

func (v NullableEnumSubscriptionFilterIncludedTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumSubscriptionFilterIncludedTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

