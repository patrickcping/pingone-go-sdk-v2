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

// SignOnPolicyActionLoginRegistrationPopulation struct for SignOnPolicyActionLoginRegistrationPopulation
type SignOnPolicyActionLoginRegistrationPopulation struct {
	// A string that specifies the population ID associated with the newly registered user.
	Id string `json:"id"`
}

// NewSignOnPolicyActionLoginRegistrationPopulation instantiates a new SignOnPolicyActionLoginRegistrationPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginRegistrationPopulation(id string) *SignOnPolicyActionLoginRegistrationPopulation {
	this := SignOnPolicyActionLoginRegistrationPopulation{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionLoginRegistrationPopulationWithDefaults instantiates a new SignOnPolicyActionLoginRegistrationPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginRegistrationPopulationWithDefaults() *SignOnPolicyActionLoginRegistrationPopulation {
	this := SignOnPolicyActionLoginRegistrationPopulation{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionLoginRegistrationPopulation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginRegistrationPopulation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionLoginRegistrationPopulation) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionLoginRegistrationPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionLoginRegistrationPopulation struct {
	value *SignOnPolicyActionLoginRegistrationPopulation
	isSet bool
}

func (v NullableSignOnPolicyActionLoginRegistrationPopulation) Get() *SignOnPolicyActionLoginRegistrationPopulation {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginRegistrationPopulation) Set(val *SignOnPolicyActionLoginRegistrationPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginRegistrationPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginRegistrationPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginRegistrationPopulation(val *SignOnPolicyActionLoginRegistrationPopulation) *NullableSignOnPolicyActionLoginRegistrationPopulation {
	return &NullableSignOnPolicyActionLoginRegistrationPopulation{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginRegistrationPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginRegistrationPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


