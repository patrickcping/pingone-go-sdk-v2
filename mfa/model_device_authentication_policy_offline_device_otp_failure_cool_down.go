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

// checks if the DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown{}

// DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown struct for DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown
type DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown struct {
	// The duration (number of time units) the user is blocked after reaching the maximum number of passcode failures.
	Duration int32 `json:"duration"`
	TimeUnit EnumTimeUnit `json:"timeUnit"`
}

type _DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown

// NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(duration int32, timeUnit EnumTimeUnit) *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown {
	this := DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown{}
	this.Duration = duration
	this.TimeUnit = timeUnit
	return &this
}

// NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDownWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDownWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown {
	this := DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown{}
	return &this
}

// GetDuration returns the Duration field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) SetDuration(v int32) {
	o.Duration = v
}

// GetTimeUnit returns the TimeUnit field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetTimeUnit() EnumTimeUnit {
	if o == nil {
		var ret EnumTimeUnit
		return ret
	}

	return o.TimeUnit
}

// GetTimeUnitOk returns a tuple with the TimeUnit field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetTimeUnitOk() (*EnumTimeUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeUnit, true
}

// SetTimeUnit sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) SetTimeUnit(v EnumTimeUnit) {
	o.TimeUnit = v
}

func (o DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["duration"] = o.Duration
	toSerialize["timeUnit"] = o.TimeUnit
	return toSerialize, nil
}

func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) UnmarshalJSON(data []byte) (err error) {
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

	varDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown := _DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown)

	if err != nil {
		return err
	}

	*o = DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(varDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown)

	return err
}

type NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown struct {
	value *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) Get() *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) Set(val *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(val *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown {
	return &NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


