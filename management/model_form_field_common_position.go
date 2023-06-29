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

// checks if the FormFieldCommonPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldCommonPosition{}

// FormFieldCommonPosition struct for FormFieldCommonPosition
type FormFieldCommonPosition struct {
	// An integer that specifies the number of rows (maximum number is 50).
	Row int32 `json:"row"`
	// An integer that specifies the number of columns (min = 1; max = 4).
	Col int32 `json:"col"`
	// An integer that specifies the width of the field (in percentage).
	Width *int32 `json:"width,omitempty"`
}

// NewFormFieldCommonPosition instantiates a new FormFieldCommonPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldCommonPosition(row int32, col int32) *FormFieldCommonPosition {
	this := FormFieldCommonPosition{}
	this.Row = row
	this.Col = col
	return &this
}

// NewFormFieldCommonPositionWithDefaults instantiates a new FormFieldCommonPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldCommonPositionWithDefaults() *FormFieldCommonPosition {
	this := FormFieldCommonPosition{}
	return &this
}

// GetRow returns the Row field value
func (o *FormFieldCommonPosition) GetRow() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Row
}

// GetRowOk returns a tuple with the Row field value
// and a boolean to check if the value has been set.
func (o *FormFieldCommonPosition) GetRowOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Row, true
}

// SetRow sets field value
func (o *FormFieldCommonPosition) SetRow(v int32) {
	o.Row = v
}

// GetCol returns the Col field value
func (o *FormFieldCommonPosition) GetCol() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Col
}

// GetColOk returns a tuple with the Col field value
// and a boolean to check if the value has been set.
func (o *FormFieldCommonPosition) GetColOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Col, true
}

// SetCol sets field value
func (o *FormFieldCommonPosition) SetCol(v int32) {
	o.Col = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *FormFieldCommonPosition) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCommonPosition) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *FormFieldCommonPosition) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *FormFieldCommonPosition) SetWidth(v int32) {
	o.Width = &v
}

func (o FormFieldCommonPosition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldCommonPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["row"] = o.Row
	toSerialize["col"] = o.Col
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableFormFieldCommonPosition struct {
	value *FormFieldCommonPosition
	isSet bool
}

func (v NullableFormFieldCommonPosition) Get() *FormFieldCommonPosition {
	return v.value
}

func (v *NullableFormFieldCommonPosition) Set(val *FormFieldCommonPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldCommonPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldCommonPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldCommonPosition(val *FormFieldCommonPosition) *NullableFormFieldCommonPosition {
	return &NullableFormFieldCommonPosition{value: val, isSet: true}
}

func (v NullableFormFieldCommonPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldCommonPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


