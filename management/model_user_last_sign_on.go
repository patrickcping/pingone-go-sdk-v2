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

// UserLastSignOn struct for UserLastSignOn
type UserLastSignOn struct {
	// The time of the last successful login of the user through the PingOne flow API.
	At *string `json:"at,omitempty"`
	// The IP address of the last successful login of the user through the PingOne flow API.
	RemoteIp *string `json:"remoteIp,omitempty"`
}

// NewUserLastSignOn instantiates a new UserLastSignOn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLastSignOn() *UserLastSignOn {
	this := UserLastSignOn{}
	return &this
}

// NewUserLastSignOnWithDefaults instantiates a new UserLastSignOn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLastSignOnWithDefaults() *UserLastSignOn {
	this := UserLastSignOn{}
	return &this
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *UserLastSignOn) GetAt() string {
	if o == nil || o.At == nil {
		var ret string
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLastSignOn) GetAtOk() (*string, bool) {
	if o == nil || o.At == nil {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *UserLastSignOn) HasAt() bool {
	if o != nil && o.At != nil {
		return true
	}

	return false
}

// SetAt gets a reference to the given string and assigns it to the At field.
func (o *UserLastSignOn) SetAt(v string) {
	o.At = &v
}

// GetRemoteIp returns the RemoteIp field value if set, zero value otherwise.
func (o *UserLastSignOn) GetRemoteIp() string {
	if o == nil || o.RemoteIp == nil {
		var ret string
		return ret
	}
	return *o.RemoteIp
}

// GetRemoteIpOk returns a tuple with the RemoteIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLastSignOn) GetRemoteIpOk() (*string, bool) {
	if o == nil || o.RemoteIp == nil {
		return nil, false
	}
	return o.RemoteIp, true
}

// HasRemoteIp returns a boolean if a field has been set.
func (o *UserLastSignOn) HasRemoteIp() bool {
	if o != nil && o.RemoteIp != nil {
		return true
	}

	return false
}

// SetRemoteIp gets a reference to the given string and assigns it to the RemoteIp field.
func (o *UserLastSignOn) SetRemoteIp(v string) {
	o.RemoteIp = &v
}

func (o UserLastSignOn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.At != nil {
		toSerialize["at"] = o.At
	}
	if o.RemoteIp != nil {
		toSerialize["remoteIp"] = o.RemoteIp
	}
	return json.Marshal(toSerialize)
}

type NullableUserLastSignOn struct {
	value *UserLastSignOn
	isSet bool
}

func (v NullableUserLastSignOn) Get() *UserLastSignOn {
	return v.value
}

func (v *NullableUserLastSignOn) Set(val *UserLastSignOn) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLastSignOn) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLastSignOn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLastSignOn(val *UserLastSignOn) *NullableUserLastSignOn {
	return &NullableUserLastSignOn{value: val, isSet: true}
}

func (v NullableUserLastSignOn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLastSignOn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


