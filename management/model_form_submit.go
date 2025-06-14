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

// checks if the FormSubmit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormSubmit{}

// FormSubmit struct for FormSubmit
type FormSubmit struct {
	// A string that specifies an identifier for the field component.
	Key *string `json:"key,omitempty"`
	// A string that specifies the button label.
	Label  string      `json:"label"`
	Styles *FormStyles `json:"styles,omitempty"`
}

// NewFormSubmit instantiates a new FormSubmit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormSubmit(label string) *FormSubmit {
	this := FormSubmit{}
	this.Label = label
	return &this
}

// NewFormSubmitWithDefaults instantiates a new FormSubmit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormSubmitWithDefaults() *FormSubmit {
	this := FormSubmit{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FormSubmit) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormSubmit) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FormSubmit) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FormSubmit) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value
func (o *FormSubmit) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FormSubmit) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FormSubmit) SetLabel(v string) {
	o.Label = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *FormSubmit) GetStyles() FormStyles {
	if o == nil || IsNil(o.Styles) {
		var ret FormStyles
		return ret
	}
	return *o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormSubmit) GetStylesOk() (*FormStyles, bool) {
	if o == nil || IsNil(o.Styles) {
		return nil, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *FormSubmit) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given FormStyles and assigns it to the Styles field.
func (o *FormSubmit) SetStyles(v FormStyles) {
	o.Styles = &v
}

func (o FormSubmit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormSubmit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	return toSerialize, nil
}

type NullableFormSubmit struct {
	value *FormSubmit
	isSet bool
}

func (v NullableFormSubmit) Get() *FormSubmit {
	return v.value
}

func (v *NullableFormSubmit) Set(val *FormSubmit) {
	v.value = val
	v.isSet = true
}

func (v NullableFormSubmit) IsSet() bool {
	return v.isSet
}

func (v *NullableFormSubmit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormSubmit(val *FormSubmit) *NullableFormSubmit {
	return &NullableFormSubmit{value: val, isSet: true}
}

func (v NullableFormSubmit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormSubmit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
