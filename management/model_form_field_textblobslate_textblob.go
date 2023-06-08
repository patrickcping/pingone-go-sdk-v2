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

// checks if the FormFieldTextblobslateTextblob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldTextblobslateTextblob{}

// FormFieldTextblobslateTextblob struct for FormFieldTextblobslateTextblob
type FormFieldTextblobslateTextblob struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
}

// NewFormFieldTextblobslateTextblob instantiates a new FormFieldTextblobslateTextblob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldTextblobslateTextblob(type_ EnumFormFieldType, position FormFieldCommonPosition) *FormFieldTextblobslateTextblob {
	this := FormFieldTextblobslateTextblob{}
	this.Type = type_
	this.Position = position
	return &this
}

// NewFormFieldTextblobslateTextblobWithDefaults instantiates a new FormFieldTextblobslateTextblob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldTextblobslateTextblobWithDefaults() *FormFieldTextblobslateTextblob {
	this := FormFieldTextblobslateTextblob{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldTextblobslateTextblob) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldTextblobslateTextblob) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldTextblobslateTextblob) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldTextblobslateTextblob) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldTextblobslateTextblob) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldTextblobslateTextblob) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

func (o FormFieldTextblobslateTextblob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldTextblobslateTextblob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

type NullableFormFieldTextblobslateTextblob struct {
	value *FormFieldTextblobslateTextblob
	isSet bool
}

func (v NullableFormFieldTextblobslateTextblob) Get() *FormFieldTextblobslateTextblob {
	return v.value
}

func (v *NullableFormFieldTextblobslateTextblob) Set(val *FormFieldTextblobslateTextblob) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldTextblobslateTextblob) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldTextblobslateTextblob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldTextblobslateTextblob(val *FormFieldTextblobslateTextblob) *NullableFormFieldTextblobslateTextblob {
	return &NullableFormFieldTextblobslateTextblob{value: val, isSet: true}
}

func (v NullableFormFieldTextblobslateTextblob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldTextblobslateTextblob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

