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

// checks if the ObjectOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectOrganization{}

// ObjectOrganization struct for ObjectOrganization
type ObjectOrganization struct {
	// A string that specifies the organization associated with the object.
	Id *string `json:"id,omitempty"`
}

// NewObjectOrganization instantiates a new ObjectOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectOrganization() *ObjectOrganization {
	this := ObjectOrganization{}
	return &this
}

// NewObjectOrganizationWithDefaults instantiates a new ObjectOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectOrganizationWithDefaults() *ObjectOrganization {
	this := ObjectOrganization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectOrganization) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectOrganization) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectOrganization) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectOrganization) SetId(v string) {
	o.Id = &v
}

func (o ObjectOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableObjectOrganization struct {
	value *ObjectOrganization
	isSet bool
}

func (v NullableObjectOrganization) Get() *ObjectOrganization {
	return v.value
}

func (v *NullableObjectOrganization) Set(val *ObjectOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectOrganization(val *ObjectOrganization) *NullableObjectOrganization {
	return &NullableObjectOrganization{value: val, isSet: true}
}

func (v NullableObjectOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


