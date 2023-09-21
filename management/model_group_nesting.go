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

// checks if the GroupNesting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupNesting{}

// GroupNesting struct for GroupNesting
type GroupNesting struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// ID of the group to nest
	Id string `json:"id"`
	// The type of the group nesting
	Type *string `json:"type,omitempty"`
}

// NewGroupNesting instantiates a new GroupNesting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupNesting(id string) *GroupNesting {
	this := GroupNesting{}
	this.Id = id
	return &this
}

// NewGroupNestingWithDefaults instantiates a new GroupNesting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupNestingWithDefaults() *GroupNesting {
	this := GroupNesting{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GroupNesting) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNesting) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GroupNesting) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *GroupNesting) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetId returns the Id field value
func (o *GroupNesting) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupNesting) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupNesting) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupNesting) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupNesting) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupNesting) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupNesting) SetType(v string) {
	o.Type = &v
}

func (o GroupNesting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupNesting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGroupNesting struct {
	value *GroupNesting
	isSet bool
}

func (v NullableGroupNesting) Get() *GroupNesting {
	return v.value
}

func (v *NullableGroupNesting) Set(val *GroupNesting) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupNesting) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupNesting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupNesting(val *GroupNesting) *NullableGroupNesting {
	return &NullableGroupNesting{value: val, isSet: true}
}

func (v NullableGroupNesting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupNesting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


