/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the DeviceAuthenticationPolicyMobileApplicationsInnerPush type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobileApplicationsInnerPush{}

// DeviceAuthenticationPolicyMobileApplicationsInnerPush struct for DeviceAuthenticationPolicyMobileApplicationsInnerPush
type DeviceAuthenticationPolicyMobileApplicationsInnerPush struct {
	// Specifies whether push notification is enabled or disabled for the policy.
	Enabled bool `json:"enabled"`
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPush instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPush object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPush(enabled bool) *DeviceAuthenticationPolicyMobileApplicationsInnerPush {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPush{}
	this.Enabled = enabled
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPushWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPush object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerPush {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPush{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPush) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPush) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPush) SetEnabled(v bool) {
	o.Enabled = v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerPush) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerPush) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInnerPush
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) Get() *DeviceAuthenticationPolicyMobileApplicationsInnerPush {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) Set(val *DeviceAuthenticationPolicyMobileApplicationsInnerPush) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInnerPush(val *DeviceAuthenticationPolicyMobileApplicationsInnerPush) *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPush) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
