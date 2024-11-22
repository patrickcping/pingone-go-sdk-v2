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

// checks if the DeviceAuthenticationPolicyOfflineDevice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyOfflineDevice{}

// DeviceAuthenticationPolicyOfflineDevice struct for DeviceAuthenticationPolicyOfflineDevice
type DeviceAuthenticationPolicyOfflineDevice struct {
	// Enabled or disabled in the policy.
	Enabled bool `json:"enabled"`
	// You can set `pairingDisabled` to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices.
	PairingDisabled *bool `json:"pairingDisabled,omitempty"`
	Otp DeviceAuthenticationPolicyOfflineDeviceOtp `json:"otp"`
	// Set to `true` if you want to allow users to provide nicknames for devices during pairing.
	PromptForNicknameOnPairing *bool `json:"promptForNicknameOnPairing,omitempty"`
}

type _DeviceAuthenticationPolicyOfflineDevice DeviceAuthenticationPolicyOfflineDevice

// NewDeviceAuthenticationPolicyOfflineDevice instantiates a new DeviceAuthenticationPolicyOfflineDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyOfflineDevice(enabled bool, otp DeviceAuthenticationPolicyOfflineDeviceOtp) *DeviceAuthenticationPolicyOfflineDevice {
	this := DeviceAuthenticationPolicyOfflineDevice{}
	this.Enabled = enabled
	this.Otp = otp
	return &this
}

// NewDeviceAuthenticationPolicyOfflineDeviceWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyOfflineDeviceWithDefaults() *DeviceAuthenticationPolicyOfflineDevice {
	this := DeviceAuthenticationPolicyOfflineDevice{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyOfflineDevice) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyOfflineDevice) SetEnabled(v bool) {
	o.Enabled = v
}

// GetPairingDisabled returns the PairingDisabled field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetPairingDisabled() bool {
	if o == nil || IsNil(o.PairingDisabled) {
		var ret bool
		return ret
	}
	return *o.PairingDisabled
}

// GetPairingDisabledOk returns a tuple with the PairingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetPairingDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PairingDisabled) {
		return nil, false
	}
	return o.PairingDisabled, true
}

// HasPairingDisabled returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) HasPairingDisabled() bool {
	if o != nil && !IsNil(o.PairingDisabled) {
		return true
	}

	return false
}

// SetPairingDisabled gets a reference to the given bool and assigns it to the PairingDisabled field.
func (o *DeviceAuthenticationPolicyOfflineDevice) SetPairingDisabled(v bool) {
	o.PairingDisabled = &v
}

// GetOtp returns the Otp field value
func (o *DeviceAuthenticationPolicyOfflineDevice) GetOtp() DeviceAuthenticationPolicyOfflineDeviceOtp {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDeviceOtp
		return ret
	}

	return o.Otp
}

// GetOtpOk returns a tuple with the Otp field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetOtpOk() (*DeviceAuthenticationPolicyOfflineDeviceOtp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Otp, true
}

// SetOtp sets field value
func (o *DeviceAuthenticationPolicyOfflineDevice) SetOtp(v DeviceAuthenticationPolicyOfflineDeviceOtp) {
	o.Otp = v
}

// GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetPromptForNicknameOnPairing() bool {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		var ret bool
		return ret
	}
	return *o.PromptForNicknameOnPairing
}

// GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) GetPromptForNicknameOnPairingOk() (*bool, bool) {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		return nil, false
	}
	return o.PromptForNicknameOnPairing, true
}

// HasPromptForNicknameOnPairing returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyOfflineDevice) HasPromptForNicknameOnPairing() bool {
	if o != nil && !IsNil(o.PromptForNicknameOnPairing) {
		return true
	}

	return false
}

// SetPromptForNicknameOnPairing gets a reference to the given bool and assigns it to the PromptForNicknameOnPairing field.
func (o *DeviceAuthenticationPolicyOfflineDevice) SetPromptForNicknameOnPairing(v bool) {
	o.PromptForNicknameOnPairing = &v
}

func (o DeviceAuthenticationPolicyOfflineDevice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyOfflineDevice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.PairingDisabled) {
		toSerialize["pairingDisabled"] = o.PairingDisabled
	}
	toSerialize["otp"] = o.Otp
	if !IsNil(o.PromptForNicknameOnPairing) {
		toSerialize["promptForNicknameOnPairing"] = o.PromptForNicknameOnPairing
	}
	return toSerialize, nil
}

func (o *DeviceAuthenticationPolicyOfflineDevice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
		"otp",
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

	varDeviceAuthenticationPolicyOfflineDevice := _DeviceAuthenticationPolicyOfflineDevice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceAuthenticationPolicyOfflineDevice)

	if err != nil {
		return err
	}

	*o = DeviceAuthenticationPolicyOfflineDevice(varDeviceAuthenticationPolicyOfflineDevice)

	return err
}

type NullableDeviceAuthenticationPolicyOfflineDevice struct {
	value *DeviceAuthenticationPolicyOfflineDevice
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyOfflineDevice) Get() *DeviceAuthenticationPolicyOfflineDevice {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyOfflineDevice) Set(val *DeviceAuthenticationPolicyOfflineDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyOfflineDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyOfflineDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyOfflineDevice(val *DeviceAuthenticationPolicyOfflineDevice) *NullableDeviceAuthenticationPolicyOfflineDevice {
	return &NullableDeviceAuthenticationPolicyOfflineDevice{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyOfflineDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyOfflineDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


