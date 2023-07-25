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

// checks if the FormFieldCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldCommon{}

// FormFieldCommon struct for FormFieldCommon
type FormFieldCommon struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
}

// NewFormFieldCommon instantiates a new FormFieldCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldCommon(type_ EnumFormFieldType, position FormFieldCommonPosition) *FormFieldCommon {
	this := FormFieldCommon{}
	this.Type = type_
	this.Position = position
	return &this
}

// NewFormFieldCommonWithDefaults instantiates a new FormFieldCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldCommonWithDefaults() *FormFieldCommon {
	this := FormFieldCommon{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldCommon) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldCommon) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldCommon) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldCommon) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldCommon) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldCommon) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

func (o FormFieldCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

type NullableFormFieldCommon struct {
	value *FormFieldCommon
	isSet bool
}

func (v NullableFormFieldCommon) Get() *FormFieldCommon {
	return v.value
}

func (v *NullableFormFieldCommon) Set(val *FormFieldCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldCommon(val *FormFieldCommon) *NullableFormFieldCommon {
	return &NullableFormFieldCommon{value: val, isSet: true}
}

func (v NullableFormFieldCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


