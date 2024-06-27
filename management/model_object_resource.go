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

// checks if the ObjectResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectResource{}

// ObjectResource struct for ObjectResource
type ObjectResource struct {
	// A string that specifies the resource associated with the object.
	Id *string `json:"id,omitempty"`
}

// NewObjectResource instantiates a new ObjectResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectResource() *ObjectResource {
	this := ObjectResource{}
	return &this
}

// NewObjectResourceWithDefaults instantiates a new ObjectResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectResourceWithDefaults() *ObjectResource {
	this := ObjectResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectResource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectResource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectResource) SetId(v string) {
	o.Id = &v
}

func (o ObjectResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableObjectResource struct {
	value *ObjectResource
	isSet bool
}

func (v NullableObjectResource) Get() *ObjectResource {
	return v.value
}

func (v *NullableObjectResource) Set(val *ObjectResource) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectResource) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectResource(val *ObjectResource) *NullableObjectResource {
	return &NullableObjectResource{value: val, isSet: true}
}

func (v NullableObjectResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


