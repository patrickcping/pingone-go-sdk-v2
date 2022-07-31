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

// SignOnPolicyActionCommonConditionAnonymousNetwork struct for SignOnPolicyActionCommonConditionAnonymousNetwork
type SignOnPolicyActionCommonConditionAnonymousNetwork struct {
	AnonymousNetwork []string `json:"anonymousNetwork"`
	Valid string `json:"valid"`
}

// NewSignOnPolicyActionCommonConditionAnonymousNetwork instantiates a new SignOnPolicyActionCommonConditionAnonymousNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonConditionAnonymousNetwork(anonymousNetwork []string, valid string) *SignOnPolicyActionCommonConditionAnonymousNetwork {
	this := SignOnPolicyActionCommonConditionAnonymousNetwork{}
	this.AnonymousNetwork = anonymousNetwork
	this.Valid = valid
	return &this
}

// NewSignOnPolicyActionCommonConditionAnonymousNetworkWithDefaults instantiates a new SignOnPolicyActionCommonConditionAnonymousNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonConditionAnonymousNetworkWithDefaults() *SignOnPolicyActionCommonConditionAnonymousNetwork {
	this := SignOnPolicyActionCommonConditionAnonymousNetwork{}
	return &this
}

// GetAnonymousNetwork returns the AnonymousNetwork field value
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) GetAnonymousNetwork() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AnonymousNetwork
}

// GetAnonymousNetworkOk returns a tuple with the AnonymousNetwork field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) GetAnonymousNetworkOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnonymousNetwork, true
}

// SetAnonymousNetwork sets field value
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) SetAnonymousNetwork(v []string) {
	o.AnonymousNetwork = v
}

// GetValid returns the Valid field value
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) GetValid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) GetValidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *SignOnPolicyActionCommonConditionAnonymousNetwork) SetValid(v string) {
	o.Valid = v
}

func (o SignOnPolicyActionCommonConditionAnonymousNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["anonymousNetwork"] = o.AnonymousNetwork
	}
	if true {
		toSerialize["valid"] = o.Valid
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionCommonConditionAnonymousNetwork struct {
	value *SignOnPolicyActionCommonConditionAnonymousNetwork
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionAnonymousNetwork) Get() *SignOnPolicyActionCommonConditionAnonymousNetwork {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionAnonymousNetwork) Set(val *SignOnPolicyActionCommonConditionAnonymousNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionAnonymousNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionAnonymousNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionAnonymousNetwork(val *SignOnPolicyActionCommonConditionAnonymousNetwork) *NullableSignOnPolicyActionCommonConditionAnonymousNetwork {
	return &NullableSignOnPolicyActionCommonConditionAnonymousNetwork{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionAnonymousNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionAnonymousNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


