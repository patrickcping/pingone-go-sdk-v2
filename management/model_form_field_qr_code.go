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

// checks if the FormFieldQrCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldQrCode{}

// FormFieldQrCode struct for FormFieldQrCode
type FormFieldQrCode struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
}

// NewFormFieldQrCode instantiates a new FormFieldQrCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldQrCode(type_ EnumFormFieldType, position FormFieldCommonPosition) *FormFieldQrCode {
	this := FormFieldQrCode{}
	this.Type = type_
	this.Position = position
	return &this
}

// NewFormFieldQrCodeWithDefaults instantiates a new FormFieldQrCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldQrCodeWithDefaults() *FormFieldQrCode {
	this := FormFieldQrCode{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldQrCode) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldQrCode) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldQrCode) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldQrCode) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldQrCode) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldQrCode) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

func (o FormFieldQrCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldQrCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	return toSerialize, nil
}

type NullableFormFieldQrCode struct {
	value *FormFieldQrCode
	isSet bool
}

func (v NullableFormFieldQrCode) Get() *FormFieldQrCode {
	return v.value
}

func (v *NullableFormFieldQrCode) Set(val *FormFieldQrCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldQrCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldQrCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldQrCode(val *FormFieldQrCode) *NullableFormFieldQrCode {
	return &NullableFormFieldQrCode{value: val, isSet: true}
}

func (v NullableFormFieldQrCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldQrCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


