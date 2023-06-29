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

// checks if the FormElementPasswordVerify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormElementPasswordVerify{}

// FormElementPasswordVerify struct for FormElementPasswordVerify
type FormElementPasswordVerify struct {
	// A string that when a second field for verifies password is used, this poperty specifies the field label for that verify field.
	LabelPasswordVerify *string `json:"labelPasswordVerify,omitempty"`
}

// NewFormElementPasswordVerify instantiates a new FormElementPasswordVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormElementPasswordVerify() *FormElementPasswordVerify {
	this := FormElementPasswordVerify{}
	return &this
}

// NewFormElementPasswordVerifyWithDefaults instantiates a new FormElementPasswordVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormElementPasswordVerifyWithDefaults() *FormElementPasswordVerify {
	this := FormElementPasswordVerify{}
	return &this
}

// GetLabelPasswordVerify returns the LabelPasswordVerify field value if set, zero value otherwise.
func (o *FormElementPasswordVerify) GetLabelPasswordVerify() string {
	if o == nil || IsNil(o.LabelPasswordVerify) {
		var ret string
		return ret
	}
	return *o.LabelPasswordVerify
}

// GetLabelPasswordVerifyOk returns a tuple with the LabelPasswordVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormElementPasswordVerify) GetLabelPasswordVerifyOk() (*string, bool) {
	if o == nil || IsNil(o.LabelPasswordVerify) {
		return nil, false
	}
	return o.LabelPasswordVerify, true
}

// HasLabelPasswordVerify returns a boolean if a field has been set.
func (o *FormElementPasswordVerify) HasLabelPasswordVerify() bool {
	if o != nil && !IsNil(o.LabelPasswordVerify) {
		return true
	}

	return false
}

// SetLabelPasswordVerify gets a reference to the given string and assigns it to the LabelPasswordVerify field.
func (o *FormElementPasswordVerify) SetLabelPasswordVerify(v string) {
	o.LabelPasswordVerify = &v
}

func (o FormElementPasswordVerify) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormElementPasswordVerify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelPasswordVerify) {
		toSerialize["labelPasswordVerify"] = o.LabelPasswordVerify
	}
	return toSerialize, nil
}

type NullableFormElementPasswordVerify struct {
	value *FormElementPasswordVerify
	isSet bool
}

func (v NullableFormElementPasswordVerify) Get() *FormElementPasswordVerify {
	return v.value
}

func (v *NullableFormElementPasswordVerify) Set(val *FormElementPasswordVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableFormElementPasswordVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableFormElementPasswordVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormElementPasswordVerify(val *FormElementPasswordVerify) *NullableFormElementPasswordVerify {
	return &NullableFormElementPasswordVerify{value: val, isSet: true}
}

func (v NullableFormElementPasswordVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormElementPasswordVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


