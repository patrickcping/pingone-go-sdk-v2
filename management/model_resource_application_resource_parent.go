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

// checks if the ResourceApplicationResourceParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceApplicationResourceParent{}

// ResourceApplicationResourceParent The application resource's parent.
type ResourceApplicationResourceParent struct {
	Type *EnumResourceApplicationResourceType `json:"type,omitempty"`
	// The application resource's parent ID.
	Id *string `json:"id,omitempty"`
}

// NewResourceApplicationResourceParent instantiates a new ResourceApplicationResourceParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceApplicationResourceParent() *ResourceApplicationResourceParent {
	this := ResourceApplicationResourceParent{}
	return &this
}

// NewResourceApplicationResourceParentWithDefaults instantiates a new ResourceApplicationResourceParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceApplicationResourceParentWithDefaults() *ResourceApplicationResourceParent {
	this := ResourceApplicationResourceParent{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResourceApplicationResourceParent) GetType() EnumResourceApplicationResourceType {
	if o == nil || IsNil(o.Type) {
		var ret EnumResourceApplicationResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResourceParent) GetTypeOk() (*EnumResourceApplicationResourceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResourceApplicationResourceParent) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumResourceApplicationResourceType and assigns it to the Type field.
func (o *ResourceApplicationResourceParent) SetType(v EnumResourceApplicationResourceType) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceApplicationResourceParent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResourceParent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceApplicationResourceParent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceApplicationResourceParent) SetId(v string) {
	o.Id = &v
}

func (o ResourceApplicationResourceParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceApplicationResourceParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableResourceApplicationResourceParent struct {
	value *ResourceApplicationResourceParent
	isSet bool
}

func (v NullableResourceApplicationResourceParent) Get() *ResourceApplicationResourceParent {
	return v.value
}

func (v *NullableResourceApplicationResourceParent) Set(val *ResourceApplicationResourceParent) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceApplicationResourceParent) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceApplicationResourceParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceApplicationResourceParent(val *ResourceApplicationResourceParent) *NullableResourceApplicationResourceParent {
	return &NullableResourceApplicationResourceParent{value: val, isSet: true}
}

func (v NullableResourceApplicationResourceParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceApplicationResourceParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


