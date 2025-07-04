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

// checks if the DeviceAuthenticationPolicyMobileOtpWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobileOtpWindow{}

// DeviceAuthenticationPolicyMobileOtpWindow struct for DeviceAuthenticationPolicyMobileOtpWindow
type DeviceAuthenticationPolicyMobileOtpWindow struct {
	StepSize DeviceAuthenticationPolicyMobileOtpWindowStepSize `json:"stepSize"`
}

// NewDeviceAuthenticationPolicyMobileOtpWindow instantiates a new DeviceAuthenticationPolicyMobileOtpWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileOtpWindow(stepSize DeviceAuthenticationPolicyMobileOtpWindowStepSize) *DeviceAuthenticationPolicyMobileOtpWindow {
	this := DeviceAuthenticationPolicyMobileOtpWindow{}
	this.StepSize = stepSize
	return &this
}

// NewDeviceAuthenticationPolicyMobileOtpWindowWithDefaults instantiates a new DeviceAuthenticationPolicyMobileOtpWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileOtpWindowWithDefaults() *DeviceAuthenticationPolicyMobileOtpWindow {
	this := DeviceAuthenticationPolicyMobileOtpWindow{}
	return &this
}

// GetStepSize returns the StepSize field value
func (o *DeviceAuthenticationPolicyMobileOtpWindow) GetStepSize() DeviceAuthenticationPolicyMobileOtpWindowStepSize {
	if o == nil {
		var ret DeviceAuthenticationPolicyMobileOtpWindowStepSize
		return ret
	}

	return o.StepSize
}

// GetStepSizeOk returns a tuple with the StepSize field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtpWindow) GetStepSizeOk() (*DeviceAuthenticationPolicyMobileOtpWindowStepSize, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepSize, true
}

// SetStepSize sets field value
func (o *DeviceAuthenticationPolicyMobileOtpWindow) SetStepSize(v DeviceAuthenticationPolicyMobileOtpWindowStepSize) {
	o.StepSize = v
}

func (o DeviceAuthenticationPolicyMobileOtpWindow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobileOtpWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["stepSize"] = o.StepSize
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyMobileOtpWindow struct {
	value *DeviceAuthenticationPolicyMobileOtpWindow
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) Get() *DeviceAuthenticationPolicyMobileOtpWindow {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) Set(val *DeviceAuthenticationPolicyMobileOtpWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileOtpWindow(val *DeviceAuthenticationPolicyMobileOtpWindow) *NullableDeviceAuthenticationPolicyMobileOtpWindow {
	return &NullableDeviceAuthenticationPolicyMobileOtpWindow{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileOtpWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileOtpWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
