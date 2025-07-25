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

// checks if the FormComponents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormComponents{}

// FormComponents struct for FormComponents
type FormComponents struct {
	Fields []FormField `json:"fields"`
}

// NewFormComponents instantiates a new FormComponents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormComponents(fields []FormField) *FormComponents {
	this := FormComponents{}
	this.Fields = fields
	return &this
}

// NewFormComponentsWithDefaults instantiates a new FormComponents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormComponentsWithDefaults() *FormComponents {
	this := FormComponents{}
	return &this
}

// GetFields returns the Fields field value
func (o *FormComponents) GetFields() []FormField {
	if o == nil {
		var ret []FormField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *FormComponents) GetFieldsOk() ([]FormField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *FormComponents) SetFields(v []FormField) {
	o.Fields = v
}

func (o FormComponents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormComponents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

type NullableFormComponents struct {
	value *FormComponents
	isSet bool
}

func (v NullableFormComponents) Get() *FormComponents {
	return v.value
}

func (v *NullableFormComponents) Set(val *FormComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableFormComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableFormComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormComponents(val *FormComponents) *NullableFormComponents {
	return &NullableFormComponents{value: val, isSet: true}
}

func (v NullableFormComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
