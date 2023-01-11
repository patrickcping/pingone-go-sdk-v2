/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout struct for DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout
type DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout struct {
	// The amount of time a user has to respond to a push notification before it expires. Minimum is 40 seconds and maximum is 150 seconds. If this parameter is not provided, the duration is set to 40 seconds.
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnitPushTimeout `json:"timeUnit"`
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout(duration int32, timeUnit EnumTimeUnitPushTimeout) *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeoutWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeoutWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout{}
	var duration int32 = 40
	this.Duration = duration
	return &this
}

// GetDuration returns the Duration field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetDurationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetTimeUnit() EnumTimeUnitPushTimeout {
	if o == nil {
		var ret EnumTimeUnitPushTimeout
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetTimeUnitOk() (*EnumTimeUnitPushTimeout, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) SetTimeUnit(v EnumTimeUnitPushTimeout) {
	o.TimeUnit = v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["timeUnit"] = o.TimeUnit
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) Get() *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) Set(val *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout(val *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


