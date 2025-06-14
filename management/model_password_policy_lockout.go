/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the PasswordPolicyLockout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicyLockout{}

// PasswordPolicyLockout Settings to control the user's lockout on unsuccessful authentication attempts
type PasswordPolicyLockout struct {
	// The length of time before a password is automatically moved out of the lock out state. The value must be a positive, non-zero integer.
	DurationSeconds *int32 `json:"durationSeconds,omitempty"`
	// The number of tries before a password is placed in the lockout state. The value must be a positive, non-zero integer.
	FailureCount *int32 `json:"failureCount,omitempty"`
}

// NewPasswordPolicyLockout instantiates a new PasswordPolicyLockout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicyLockout() *PasswordPolicyLockout {
	this := PasswordPolicyLockout{}
	return &this
}

// NewPasswordPolicyLockoutWithDefaults instantiates a new PasswordPolicyLockout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyLockoutWithDefaults() *PasswordPolicyLockout {
	this := PasswordPolicyLockout{}
	return &this
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *PasswordPolicyLockout) GetDurationSeconds() int32 {
	if o == nil || IsNil(o.DurationSeconds) {
		var ret int32
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyLockout) GetDurationSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.DurationSeconds) {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *PasswordPolicyLockout) HasDurationSeconds() bool {
	if o != nil && !IsNil(o.DurationSeconds) {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given int32 and assigns it to the DurationSeconds field.
func (o *PasswordPolicyLockout) SetDurationSeconds(v int32) {
	o.DurationSeconds = &v
}

// GetFailureCount returns the FailureCount field value if set, zero value otherwise.
func (o *PasswordPolicyLockout) GetFailureCount() int32 {
	if o == nil || IsNil(o.FailureCount) {
		var ret int32
		return ret
	}
	return *o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicyLockout) GetFailureCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FailureCount) {
		return nil, false
	}
	return o.FailureCount, true
}

// HasFailureCount returns a boolean if a field has been set.
func (o *PasswordPolicyLockout) HasFailureCount() bool {
	if o != nil && !IsNil(o.FailureCount) {
		return true
	}

	return false
}

// SetFailureCount gets a reference to the given int32 and assigns it to the FailureCount field.
func (o *PasswordPolicyLockout) SetFailureCount(v int32) {
	o.FailureCount = &v
}

func (o PasswordPolicyLockout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicyLockout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DurationSeconds) {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	if !IsNil(o.FailureCount) {
		toSerialize["failureCount"] = o.FailureCount
	}
	return toSerialize, nil
}

type NullablePasswordPolicyLockout struct {
	value *PasswordPolicyLockout
	isSet bool
}

func (v NullablePasswordPolicyLockout) Get() *PasswordPolicyLockout {
	return v.value
}

func (v *NullablePasswordPolicyLockout) Set(val *PasswordPolicyLockout) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicyLockout) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicyLockout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicyLockout(val *PasswordPolicyLockout) *NullablePasswordPolicyLockout {
	return &NullablePasswordPolicyLockout{value: val, isSet: true}
}

func (v NullablePasswordPolicyLockout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicyLockout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
