/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ObjectPopulation struct for ObjectPopulation
type ObjectPopulation struct {
	// A string that specifies the population associated with the object.
	Id *string `json:"id,omitempty"`
}

// NewObjectPopulation instantiates a new ObjectPopulation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPopulation() *ObjectPopulation {
	this := ObjectPopulation{}
	return &this
}

// NewObjectPopulationWithDefaults instantiates a new ObjectPopulation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPopulationWithDefaults() *ObjectPopulation {
	this := ObjectPopulation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectPopulation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPopulation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectPopulation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectPopulation) SetId(v string) {
	o.Id = &v
}

func (o ObjectPopulation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableObjectPopulation struct {
	value *ObjectPopulation
	isSet bool
}

func (v NullableObjectPopulation) Get() *ObjectPopulation {
	return v.value
}

func (v *NullableObjectPopulation) Set(val *ObjectPopulation) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPopulation) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPopulation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPopulation(val *ObjectPopulation) *NullableObjectPopulation {
	return &NullableObjectPopulation{value: val, isSet: true}
}

func (v NullableObjectPopulation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPopulation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


