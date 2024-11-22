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

// checks if the DeviceAuthenticationPolicyOfflineDeviceOtp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyOfflineDeviceOtp{}

// DeviceAuthenticationPolicyOfflineDeviceOtp struct for DeviceAuthenticationPolicyOfflineDeviceOtp
type DeviceAuthenticationPolicyOfflineDeviceOtp struct {
	LifeTime DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime `json:"lifeTime"`
	Failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure `json:"failure"`
	// Used to specify the length of the OTP that is shown to users. Minimum length is `6` digits and maximum is `10` digits. If the parameter is not provided, the default is `6` digits.
	OtpLength *int32 `json:"otpLength,omitempty"`
}

type _DeviceAuthenticationPolicyOfflineDeviceOtp DeviceAuthenticationPolicyOfflineDeviceOtp

// NewDeviceAuthenticationPolicyOfflineDeviceOtp instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyOfflineDeviceOtp(lifeTime DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime, failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure) *DeviceAuthenticationPolicyOfflineDeviceOtp {
	this := DeviceAuthenticationPolicyOfflineDeviceOtp{}
	this.LifeTime = lifeTime
	this.Failure = failure
	var otpLength int32 = 6
	this.OtpLength = &otpLength
	return &this
}

// NewDeviceAuthenticationPolicyOfflineDeviceOtpWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyOfflineDeviceOtpWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtp {
	this := DeviceAuthenticationPolicyOfflineDeviceOtp{}
	var otpLength int32 = 6
	this.OtpLength = &otpLength
	return &this
}

// GetLifeTime returns the LifeTime field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetLifeTime() DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime
		return ret
	}

	return o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetLifeTimeOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifeTime, true
}

// SetLifeTime sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetLifeTime(v DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) {
	o.LifeTime = v
}

// GetFailure returns the Failure field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetFailure() DeviceAuthenticationPolicyOfflineDeviceOtpFailure {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtpFailure
		return ret
	}

	return o.Failure
}

// GetFailureOk returns a tuple with the Failure field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetFailureOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailure, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failure, true
}

// SetFailure sets field value
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetFailure(v DeviceAuthenticationPolicyOfflineDeviceOtpFailure) {
	o.Failure = v
}

// GetOtpLength returns the OtpLength field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetOtpLength() int32 {
	if o == nil || IsNil(o.OtpLength) {
		var ret int32
		return ret
	}
	return *o.OtpLength
}

// GetOtpLengthOk returns a tuple with the OtpLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetOtpLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.OtpLength) {
		return nil, false
	}
	return o.OtpLength, true
}

// HasOtpLength returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) HasOtpLength() bool {
	if o != nil && !IsNil(o.OtpLength) {
		return true
	}

	return false
}

// SetOtpLength gets a reference to the given int32 and assigns it to the OtpLength field.
func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetOtpLength(v int32) {
	o.OtpLength = &v
}

func (o DeviceAuthenticationPolicyOfflineDeviceOtp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyOfflineDeviceOtp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lifeTime"] = o.LifeTime
	toSerialize["failure"] = o.Failure
	if !IsNil(o.OtpLength) {
		toSerialize["otpLength"] = o.OtpLength
	}
	return toSerialize, nil
}

func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lifeTime",
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

	varDeviceAuthenticationPolicyOfflineDeviceOtp := _DeviceAuthenticationPolicyOfflineDeviceOtp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthenticationPolicyOfflineDeviceOtp)

	if err != nil {
		return err
	}

	*o = DeviceAuthenticationPolicyOfflineDeviceOtp(varDeviceAuthenticationPolicyOfflineDeviceOtp)

	return err
}

type NullableDeviceAuthenticationPolicyOfflineDeviceOtp struct {
	value *DeviceAuthenticationPolicyOfflineDeviceOtp
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtp) Get() *DeviceAuthenticationPolicyOfflineDeviceOtp {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtp) Set(val *DeviceAuthenticationPolicyOfflineDeviceOtp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyOfflineDeviceOtp(val *DeviceAuthenticationPolicyOfflineDeviceOtp) *NullableDeviceAuthenticationPolicyOfflineDeviceOtp {
	return &NullableDeviceAuthenticationPolicyOfflineDeviceOtp{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyOfflineDeviceOtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyOfflineDeviceOtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


