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

// checks if the DeviceAuthenticationPolicyMobileOtp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobileOtp{}

// DeviceAuthenticationPolicyMobileOtp struct for DeviceAuthenticationPolicyMobileOtp
type DeviceAuthenticationPolicyMobileOtp struct {
	Failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure `json:"failure"`
	Window *DeviceAuthenticationPolicyMobileOtpWindow `json:"window,omitempty"`
}

type _DeviceAuthenticationPolicyMobileOtp DeviceAuthenticationPolicyMobileOtp

// NewDeviceAuthenticationPolicyMobileOtp instantiates a new DeviceAuthenticationPolicyMobileOtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileOtp(failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure) *DeviceAuthenticationPolicyMobileOtp {
	this := DeviceAuthenticationPolicyMobileOtp{}
	this.Failure = failure
	return &this
}

// NewDeviceAuthenticationPolicyMobileOtpWithDefaults instantiates a new DeviceAuthenticationPolicyMobileOtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileOtpWithDefaults() *DeviceAuthenticationPolicyMobileOtp {
	this := DeviceAuthenticationPolicyMobileOtp{}
	return &this
}

// GetFailure returns the Failure field value
func (o *DeviceAuthenticationPolicyMobileOtp) GetFailure() DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtpFailure
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtp) GetFailureOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *DeviceAuthenticationPolicyMobileOtp) SetFailure(v DeviceAuthenticationPolicyOfflineDeviceOtpFailure) {
	o.Failure = v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileOtp) GetWindow() DeviceAuthenticationPolicyMobileOtpWindow {
	if o == nil || IsNil(o.Window) {
		var ret DeviceAuthenticationPolicyMobileOtpWindow
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileOtp) GetWindowOk() (*DeviceAuthenticationPolicyMobileOtpWindow, bool) {
	if o == nil || IsNil(o.Window) {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileOtp) HasWindow() bool {
	if o != nil && !IsNil(o.Window) {
		return true
	}

	return false
}

// SetWindow gets a reference to the given DeviceAuthenticationPolicyMobileOtpWindow and assigns it to the Window field.
func (o *DeviceAuthenticationPolicyMobileOtp) SetWindow(v DeviceAuthenticationPolicyMobileOtpWindow) {
	o.Window = &v
}

func (o DeviceAuthenticationPolicyMobileOtp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobileOtp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failure"] = o.Failure
	if !IsNil(o.Window) {
		toSerialize["window"] = o.Window
	}
	return toSerialize, nil
}

func (o *DeviceAuthenticationPolicyMobileOtp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"failure",
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

	varDeviceAuthenticationPolicyMobileOtp := _DeviceAuthenticationPolicyMobileOtp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthenticationPolicyMobileOtp)

	if err != nil {
		return err
	}

	*o = DeviceAuthenticationPolicyMobileOtp(varDeviceAuthenticationPolicyMobileOtp)

	return err
}

type NullableDeviceAuthenticationPolicyMobileOtp struct {
	value *DeviceAuthenticationPolicyMobileOtp
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileOtp) Get() *DeviceAuthenticationPolicyMobileOtp {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileOtp) Set(val *DeviceAuthenticationPolicyMobileOtp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileOtp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileOtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileOtp(val *DeviceAuthenticationPolicyMobileOtp) *NullableDeviceAuthenticationPolicyMobileOtp {
	return &NullableDeviceAuthenticationPolicyMobileOtp{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileOtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileOtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


