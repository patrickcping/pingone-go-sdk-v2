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

// checks if the DeviceAuthenticationPolicyAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyAuthentication{}

// DeviceAuthenticationPolicyAuthentication An object that contains the device selection settings.
type DeviceAuthenticationPolicyAuthentication struct {
	DeviceSelection EnumMFADevicePolicySelection `json:"deviceSelection"`
}

// NewDeviceAuthenticationPolicyAuthentication instantiates a new DeviceAuthenticationPolicyAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyAuthentication(deviceSelection EnumMFADevicePolicySelection) *DeviceAuthenticationPolicyAuthentication {
	this := DeviceAuthenticationPolicyAuthentication{}
	this.DeviceSelection = deviceSelection
	return &this
}

// NewDeviceAuthenticationPolicyAuthenticationWithDefaults instantiates a new DeviceAuthenticationPolicyAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyAuthenticationWithDefaults() *DeviceAuthenticationPolicyAuthentication {
	this := DeviceAuthenticationPolicyAuthentication{}
	return &this
}

// GetDeviceSelection returns the DeviceSelection field value
func (o *DeviceAuthenticationPolicyAuthentication) GetDeviceSelection() EnumMFADevicePolicySelection {
	if o == nil {
		var ret EnumMFADevicePolicySelection
		return ret
	}

	return o.DeviceSelection
}

// GetDeviceSelectionOk returns a tuple with the DeviceSelection field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyAuthentication) GetDeviceSelectionOk() (*EnumMFADevicePolicySelection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceSelection, true
}

// SetDeviceSelection sets field value
func (o *DeviceAuthenticationPolicyAuthentication) SetDeviceSelection(v EnumMFADevicePolicySelection) {
	o.DeviceSelection = v
}

func (o DeviceAuthenticationPolicyAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceSelection"] = o.DeviceSelection
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyAuthentication struct {
	value *DeviceAuthenticationPolicyAuthentication
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyAuthentication) Get() *DeviceAuthenticationPolicyAuthentication {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyAuthentication) Set(val *DeviceAuthenticationPolicyAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyAuthentication(val *DeviceAuthenticationPolicyAuthentication) *NullableDeviceAuthenticationPolicyAuthentication {
	return &NullableDeviceAuthenticationPolicyAuthentication{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


