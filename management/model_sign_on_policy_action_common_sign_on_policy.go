/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the SignOnPolicyActionCommonSignOnPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionCommonSignOnPolicy{}

// SignOnPolicyActionCommonSignOnPolicy The sign-on policy associated with this sign-on policy assignment
type SignOnPolicyActionCommonSignOnPolicy struct {
	// A string that specifies the sign-on policy resource's unique identifier associated with this sign-on policy assignment.
	Id string `json:"id"`
}

// NewSignOnPolicyActionCommonSignOnPolicy instantiates a new SignOnPolicyActionCommonSignOnPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonSignOnPolicy(id string) *SignOnPolicyActionCommonSignOnPolicy {
	this := SignOnPolicyActionCommonSignOnPolicy{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionCommonSignOnPolicyWithDefaults instantiates a new SignOnPolicyActionCommonSignOnPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonSignOnPolicyWithDefaults() *SignOnPolicyActionCommonSignOnPolicy {
	this := SignOnPolicyActionCommonSignOnPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionCommonSignOnPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonSignOnPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionCommonSignOnPolicy) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionCommonSignOnPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionCommonSignOnPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSignOnPolicyActionCommonSignOnPolicy struct {
	value *SignOnPolicyActionCommonSignOnPolicy
	isSet bool
}

func (v NullableSignOnPolicyActionCommonSignOnPolicy) Get() *SignOnPolicyActionCommonSignOnPolicy {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonSignOnPolicy) Set(val *SignOnPolicyActionCommonSignOnPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonSignOnPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonSignOnPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonSignOnPolicy(val *SignOnPolicyActionCommonSignOnPolicy) *NullableSignOnPolicyActionCommonSignOnPolicy {
	return &NullableSignOnPolicyActionCommonSignOnPolicy{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonSignOnPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonSignOnPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


