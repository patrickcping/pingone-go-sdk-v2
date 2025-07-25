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

// checks if the SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy{}

// SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy Details of the MFA policy that should be used. If it is omitted, the environment default MFA policy is used.
type SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy struct {
	// The ID of the MFA policy that should be used.
	Id string `json:"id"`
}

// NewSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy instantiates a new SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy(id string) *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy {
	this := SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicyWithDefaults instantiates a new SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicyWithDefaults() *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy {
	this := SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy struct {
	value *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy
	isSet bool
}

func (v NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) Get() *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy {
	return v.value
}

func (v *NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) Set(val *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy(val *SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) *NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy {
	return &NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
