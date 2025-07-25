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

// checks if the ResourceApplicationResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceApplicationResource{}

// ResourceApplicationResource struct for ResourceApplicationResource
type ResourceApplicationResource struct {
	// The application resource's description.
	Description *string `json:"description,omitempty"`
	// The resource's unique identifier.
	Id *string `json:"id,omitempty"`
	// The application resource name. The name value must be unique.
	Name   string                             `json:"name"`
	Parent *ResourceApplicationResourceParent `json:"parent,omitempty"`
}

// NewResourceApplicationResource instantiates a new ResourceApplicationResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceApplicationResource(name string) *ResourceApplicationResource {
	this := ResourceApplicationResource{}
	this.Name = name
	return &this
}

// NewResourceApplicationResourceWithDefaults instantiates a new ResourceApplicationResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceApplicationResourceWithDefaults() *ResourceApplicationResource {
	this := ResourceApplicationResource{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ResourceApplicationResource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ResourceApplicationResource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ResourceApplicationResource) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceApplicationResource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceApplicationResource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceApplicationResource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ResourceApplicationResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ResourceApplicationResource) SetName(v string) {
	o.Name = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ResourceApplicationResource) GetParent() ResourceApplicationResourceParent {
	if o == nil || IsNil(o.Parent) {
		var ret ResourceApplicationResourceParent
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceApplicationResource) GetParentOk() (*ResourceApplicationResourceParent, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ResourceApplicationResource) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given ResourceApplicationResourceParent and assigns it to the Parent field.
func (o *ResourceApplicationResource) SetParent(v ResourceApplicationResourceParent) {
	o.Parent = &v
}

func (o ResourceApplicationResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceApplicationResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

type NullableResourceApplicationResource struct {
	value *ResourceApplicationResource
	isSet bool
}

func (v NullableResourceApplicationResource) Get() *ResourceApplicationResource {
	return v.value
}

func (v *NullableResourceApplicationResource) Set(val *ResourceApplicationResource) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceApplicationResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceApplicationResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceApplicationResource(val *ResourceApplicationResource) *NullableResourceApplicationResource {
	return &NullableResourceApplicationResource{value: val, isSet: true}
}

func (v NullableResourceApplicationResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceApplicationResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
