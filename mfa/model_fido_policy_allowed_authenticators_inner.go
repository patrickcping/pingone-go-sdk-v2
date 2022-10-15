/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// FIDOPolicyAllowedAuthenticatorsInner struct for FIDOPolicyAllowedAuthenticatorsInner
type FIDOPolicyAllowedAuthenticatorsInner struct {
	// The identifier of the authenticator to allow.
	Id string `json:"id"`
}

// NewFIDOPolicyAllowedAuthenticatorsInner instantiates a new FIDOPolicyAllowedAuthenticatorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDOPolicyAllowedAuthenticatorsInner(id string) *FIDOPolicyAllowedAuthenticatorsInner {
	this := FIDOPolicyAllowedAuthenticatorsInner{}
	this.Id = id
	return &this
}

// NewFIDOPolicyAllowedAuthenticatorsInnerWithDefaults instantiates a new FIDOPolicyAllowedAuthenticatorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDOPolicyAllowedAuthenticatorsInnerWithDefaults() *FIDOPolicyAllowedAuthenticatorsInner {
	this := FIDOPolicyAllowedAuthenticatorsInner{}
	return &this
}

// GetId returns the Id field value
func (o *FIDOPolicyAllowedAuthenticatorsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FIDOPolicyAllowedAuthenticatorsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FIDOPolicyAllowedAuthenticatorsInner) SetId(v string) {
	o.Id = v
}

func (o FIDOPolicyAllowedAuthenticatorsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableFIDOPolicyAllowedAuthenticatorsInner struct {
	value *FIDOPolicyAllowedAuthenticatorsInner
	isSet bool
}

func (v NullableFIDOPolicyAllowedAuthenticatorsInner) Get() *FIDOPolicyAllowedAuthenticatorsInner {
	return v.value
}

func (v *NullableFIDOPolicyAllowedAuthenticatorsInner) Set(val *FIDOPolicyAllowedAuthenticatorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDOPolicyAllowedAuthenticatorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDOPolicyAllowedAuthenticatorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDOPolicyAllowedAuthenticatorsInner(val *FIDOPolicyAllowedAuthenticatorsInner) *NullableFIDOPolicyAllowedAuthenticatorsInner {
	return &NullableFIDOPolicyAllowedAuthenticatorsInner{value: val, isSet: true}
}

func (v NullableFIDOPolicyAllowedAuthenticatorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDOPolicyAllowedAuthenticatorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


