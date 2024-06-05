/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the APIServerPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerPolicy{}

// APIServerPolicy struct for APIServerPolicy
type APIServerPolicy struct {
	// The ID of the root policy.
	Id string `json:"id"`
}

// NewAPIServerPolicy instantiates a new APIServerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerPolicy(id string) *APIServerPolicy {
	this := APIServerPolicy{}
	this.Id = id
	return &this
}

// NewAPIServerPolicyWithDefaults instantiates a new APIServerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerPolicyWithDefaults() *APIServerPolicy {
	this := APIServerPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *APIServerPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIServerPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIServerPolicy) SetId(v string) {
	o.Id = v
}

func (o APIServerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAPIServerPolicy struct {
	value *APIServerPolicy
	isSet bool
}

func (v NullableAPIServerPolicy) Get() *APIServerPolicy {
	return v.value
}

func (v *NullableAPIServerPolicy) Set(val *APIServerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerPolicy(val *APIServerPolicy) *NullableAPIServerPolicy {
	return &NullableAPIServerPolicy{value: val, isSet: true}
}

func (v NullableAPIServerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

