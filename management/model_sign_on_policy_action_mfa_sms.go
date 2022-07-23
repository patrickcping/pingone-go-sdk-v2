/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionMFASms Specifies MFA through SMS text messaging options.
type SignOnPolicyActionMFASms struct {
	// A boolean that specifies the enabled/disabled state of the MFA through SMS action. The default is disabled when creating a new policy. When enabled, it allows users to receive the one-time password and authenticate through SMS text message.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewSignOnPolicyActionMFASms instantiates a new SignOnPolicyActionMFASms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFASms() *SignOnPolicyActionMFASms {
	this := SignOnPolicyActionMFASms{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewSignOnPolicyActionMFASmsWithDefaults instantiates a new SignOnPolicyActionMFASms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFASmsWithDefaults() *SignOnPolicyActionMFASms {
	this := SignOnPolicyActionMFASms{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SignOnPolicyActionMFASms) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFASms) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SignOnPolicyActionMFASms) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SignOnPolicyActionMFASms) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o SignOnPolicyActionMFASms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionMFASms struct {
	value *SignOnPolicyActionMFASms
	isSet bool
}

func (v NullableSignOnPolicyActionMFASms) Get() *SignOnPolicyActionMFASms {
	return v.value
}

func (v *NullableSignOnPolicyActionMFASms) Set(val *SignOnPolicyActionMFASms) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFASms) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFASms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFASms(val *SignOnPolicyActionMFASms) *NullableSignOnPolicyActionMFASms {
	return &NullableSignOnPolicyActionMFASms{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFASms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFASms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


