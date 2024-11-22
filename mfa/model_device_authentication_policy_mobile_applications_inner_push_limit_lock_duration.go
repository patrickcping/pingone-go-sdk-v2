/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration{}

// DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration struct for DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration
type DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration struct {
	// The length of time that push notifications should be blocked for the application if the defined limit has been reached. The minimum value is 1 minute and the maximum value is 120 minutes. If this parameter is not provided, the default value is 30 minutes.
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnit `json:"timeUnit"`
}

type _DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration(duration int32, timeUnit EnumTimeUnit) *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDurationWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDurationWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration {
	this := DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration{}
	var duration int32 = 30
	this.Duration = duration
	return &this
}

// GetDuration returns the Duration field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) GetTimeUnit() EnumTimeUnit {
	if o == nil {
		var ret EnumTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) GetTimeUnitOk() (*EnumTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) SetTimeUnit(v EnumTimeUnit) {
	o.TimeUnit = v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"duration",
		"timeUnit",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration := _DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration)

	if err != nil {
		return err
	}

	*o = DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration(varDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration)

	return err
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) Get() *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) Set(val *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration(val *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInnerPushLimitLockDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


