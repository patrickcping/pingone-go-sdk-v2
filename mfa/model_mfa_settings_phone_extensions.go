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

// checks if the MFASettingsPhoneExtensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFASettingsPhoneExtensions{}

// MFASettingsPhoneExtensions Contains settings for phone extension support.
type MFASettingsPhoneExtensions struct {
	// Set to `true` to allow one-time passwords to be delivered via voice to phone numbers that include extensions. Set to `false` to disable support for phone numbers with extensions. By default, support for extensions is disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewMFASettingsPhoneExtensions instantiates a new MFASettingsPhoneExtensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFASettingsPhoneExtensions() *MFASettingsPhoneExtensions {
	this := MFASettingsPhoneExtensions{}
	return &this
}

// NewMFASettingsPhoneExtensionsWithDefaults instantiates a new MFASettingsPhoneExtensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFASettingsPhoneExtensionsWithDefaults() *MFASettingsPhoneExtensions {
	this := MFASettingsPhoneExtensions{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MFASettingsPhoneExtensions) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettingsPhoneExtensions) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MFASettingsPhoneExtensions) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MFASettingsPhoneExtensions) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o MFASettingsPhoneExtensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFASettingsPhoneExtensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableMFASettingsPhoneExtensions struct {
	value *MFASettingsPhoneExtensions
	isSet bool
}

func (v NullableMFASettingsPhoneExtensions) Get() *MFASettingsPhoneExtensions {
	return v.value
}

func (v *NullableMFASettingsPhoneExtensions) Set(val *MFASettingsPhoneExtensions) {
	v.value = val
	v.isSet = true
}

func (v NullableMFASettingsPhoneExtensions) IsSet() bool {
	return v.isSet
}

func (v *NullableMFASettingsPhoneExtensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFASettingsPhoneExtensions(val *MFASettingsPhoneExtensions) *NullableMFASettingsPhoneExtensions {
	return &NullableMFASettingsPhoneExtensions{value: val, isSet: true}
}

func (v NullableMFASettingsPhoneExtensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFASettingsPhoneExtensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
