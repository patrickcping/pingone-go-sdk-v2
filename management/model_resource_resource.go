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

// ResourceResource struct for ResourceResource
type ResourceResource struct {
	Id *string `json:"id,omitempty"`
}

// NewResourceResource instantiates a new ResourceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceResource() *ResourceResource {
	this := ResourceResource{}
	return &this
}

// NewResourceResourceWithDefaults instantiates a new ResourceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceResourceWithDefaults() *ResourceResource {
	this := ResourceResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceResource) SetId(v string) {
	o.Id = &v
}

func (o ResourceResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableResourceResource struct {
	value *ResourceResource
	isSet bool
}

func (v NullableResourceResource) Get() *ResourceResource {
	return v.value
}

func (v *NullableResourceResource) Set(val *ResourceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceResource(val *ResourceResource) *NullableResourceResource {
	return &NullableResourceResource{value: val, isSet: true}
}

func (v NullableResourceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


