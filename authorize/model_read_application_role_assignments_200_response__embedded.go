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

// checks if the ReadApplicationRoleAssignments200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadApplicationRoleAssignments200ResponseEmbedded{}

// ReadApplicationRoleAssignments200ResponseEmbedded struct for ReadApplicationRoleAssignments200ResponseEmbedded
type ReadApplicationRoleAssignments200ResponseEmbedded struct {
	Assignments []ApplicationRoleAssignment `json:"assignments,omitempty"`
}

// NewReadApplicationRoleAssignments200ResponseEmbedded instantiates a new ReadApplicationRoleAssignments200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadApplicationRoleAssignments200ResponseEmbedded() *ReadApplicationRoleAssignments200ResponseEmbedded {
	this := ReadApplicationRoleAssignments200ResponseEmbedded{}
	return &this
}

// NewReadApplicationRoleAssignments200ResponseEmbeddedWithDefaults instantiates a new ReadApplicationRoleAssignments200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadApplicationRoleAssignments200ResponseEmbeddedWithDefaults() *ReadApplicationRoleAssignments200ResponseEmbedded {
	this := ReadApplicationRoleAssignments200ResponseEmbedded{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *ReadApplicationRoleAssignments200ResponseEmbedded) GetAssignments() []ApplicationRoleAssignment {
	if o == nil || IsNil(o.Assignments) {
		var ret []ApplicationRoleAssignment
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadApplicationRoleAssignments200ResponseEmbedded) GetAssignmentsOk() ([]ApplicationRoleAssignment, bool) {
	if o == nil || IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *ReadApplicationRoleAssignments200ResponseEmbedded) HasAssignments() bool {
	if o != nil && !IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []ApplicationRoleAssignment and assigns it to the Assignments field.
func (o *ReadApplicationRoleAssignments200ResponseEmbedded) SetAssignments(v []ApplicationRoleAssignment) {
	o.Assignments = v
}

func (o ReadApplicationRoleAssignments200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadApplicationRoleAssignments200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	return toSerialize, nil
}

type NullableReadApplicationRoleAssignments200ResponseEmbedded struct {
	value *ReadApplicationRoleAssignments200ResponseEmbedded
	isSet bool
}

func (v NullableReadApplicationRoleAssignments200ResponseEmbedded) Get() *ReadApplicationRoleAssignments200ResponseEmbedded {
	return v.value
}

func (v *NullableReadApplicationRoleAssignments200ResponseEmbedded) Set(val *ReadApplicationRoleAssignments200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadApplicationRoleAssignments200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadApplicationRoleAssignments200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadApplicationRoleAssignments200ResponseEmbedded(val *ReadApplicationRoleAssignments200ResponseEmbedded) *NullableReadApplicationRoleAssignments200ResponseEmbedded {
	return &NullableReadApplicationRoleAssignments200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadApplicationRoleAssignments200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadApplicationRoleAssignments200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


