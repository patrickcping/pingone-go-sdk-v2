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

// checks if the DeviceAuthenticationPolicyMobile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobile{}

// DeviceAuthenticationPolicyMobile Mobile device authentication policy settings.
type DeviceAuthenticationPolicyMobile struct {
	// A boolean that specifies whether the method is enabled or disabled in the policy.
	Enabled      bool                                                `json:"enabled"`
	Otp          DeviceAuthenticationPolicyMobileOtp                 `json:"otp"`
	Applications []DeviceAuthenticationPolicyMobileApplicationsInner `json:"applications,omitempty"`
	// Set to `true` if you want to allow users to provide nicknames for devices during pairing.
	PromptForNicknameOnPairing *bool `json:"promptForNicknameOnPairing,omitempty"`
}

// NewDeviceAuthenticationPolicyMobile instantiates a new DeviceAuthenticationPolicyMobile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobile(enabled bool, otp DeviceAuthenticationPolicyMobileOtp) *DeviceAuthenticationPolicyMobile {
	this := DeviceAuthenticationPolicyMobile{}
	this.Enabled = enabled
	this.Otp = otp
	return &this
}

// NewDeviceAuthenticationPolicyMobileWithDefaults instantiates a new DeviceAuthenticationPolicyMobile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileWithDefaults() *DeviceAuthenticationPolicyMobile {
	this := DeviceAuthenticationPolicyMobile{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *DeviceAuthenticationPolicyMobile) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobile) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DeviceAuthenticationPolicyMobile) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOtp returns the Otp field value
func (o *DeviceAuthenticationPolicyMobile) GetOtp() DeviceAuthenticationPolicyMobileOtp {
	if o == nil {
		var ret DeviceAuthenticationPolicyMobileOtp
		return ret
	}

	return o.Otp
}

// GetOtpOk returns a tuple with the Otp field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobile) GetOtpOk() (*DeviceAuthenticationPolicyMobileOtp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Otp, true
}

// SetOtp sets field value
func (o *DeviceAuthenticationPolicyMobile) SetOtp(v DeviceAuthenticationPolicyMobileOtp) {
	o.Otp = v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobile) GetApplications() []DeviceAuthenticationPolicyMobileApplicationsInner {
	if o == nil || IsNil(o.Applications) {
		var ret []DeviceAuthenticationPolicyMobileApplicationsInner
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobile) GetApplicationsOk() ([]DeviceAuthenticationPolicyMobileApplicationsInner, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobile) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []DeviceAuthenticationPolicyMobileApplicationsInner and assigns it to the Applications field.
func (o *DeviceAuthenticationPolicyMobile) SetApplications(v []DeviceAuthenticationPolicyMobileApplicationsInner) {
	o.Applications = v
}

// GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobile) GetPromptForNicknameOnPairing() bool {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		var ret bool
		return ret
	}
	return *o.PromptForNicknameOnPairing
}

// GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobile) GetPromptForNicknameOnPairingOk() (*bool, bool) {
	if o == nil || IsNil(o.PromptForNicknameOnPairing) {
		return nil, false
	}
	return o.PromptForNicknameOnPairing, true
}

// HasPromptForNicknameOnPairing returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobile) HasPromptForNicknameOnPairing() bool {
	if o != nil && !IsNil(o.PromptForNicknameOnPairing) {
		return true
	}

	return false
}

// SetPromptForNicknameOnPairing gets a reference to the given bool and assigns it to the PromptForNicknameOnPairing field.
func (o *DeviceAuthenticationPolicyMobile) SetPromptForNicknameOnPairing(v bool) {
	o.PromptForNicknameOnPairing = &v
}

func (o DeviceAuthenticationPolicyMobile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	toSerialize["otp"] = o.Otp
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.PromptForNicknameOnPairing) {
		toSerialize["promptForNicknameOnPairing"] = o.PromptForNicknameOnPairing
	}
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyMobile struct {
	value *DeviceAuthenticationPolicyMobile
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobile) Get() *DeviceAuthenticationPolicyMobile {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobile) Set(val *DeviceAuthenticationPolicyMobile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobile(val *DeviceAuthenticationPolicyMobile) *NullableDeviceAuthenticationPolicyMobile {
	return &NullableDeviceAuthenticationPolicyMobile{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
