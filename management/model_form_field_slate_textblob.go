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

// checks if the FormFieldSlateTextblob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldSlateTextblob{}

// FormFieldSlateTextblob struct for FormFieldSlateTextblob
type FormFieldSlateTextblob struct {
	Type     EnumFormFieldType       `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A string that specifies the field content.
	Content *string `json:"content,omitempty"`
}

// NewFormFieldSlateTextblob instantiates a new FormFieldSlateTextblob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldSlateTextblob(type_ EnumFormFieldType, position FormFieldCommonPosition) *FormFieldSlateTextblob {
	this := FormFieldSlateTextblob{}
	this.Type = type_
	this.Position = position
	return &this
}

// NewFormFieldSlateTextblobWithDefaults instantiates a new FormFieldSlateTextblob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldSlateTextblobWithDefaults() *FormFieldSlateTextblob {
	this := FormFieldSlateTextblob{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldSlateTextblob) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldSlateTextblob) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldSlateTextblob) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldSlateTextblob) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldSlateTextblob) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldSlateTextblob) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *FormFieldSlateTextblob) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldSlateTextblob) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *FormFieldSlateTextblob) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *FormFieldSlateTextblob) SetContent(v string) {
	o.Content = &v
}

func (o FormFieldSlateTextblob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldSlateTextblob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	return toSerialize, nil
}

type NullableFormFieldSlateTextblob struct {
	value *FormFieldSlateTextblob
	isSet bool
}

func (v NullableFormFieldSlateTextblob) Get() *FormFieldSlateTextblob {
	return v.value
}

func (v *NullableFormFieldSlateTextblob) Set(val *FormFieldSlateTextblob) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldSlateTextblob) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldSlateTextblob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldSlateTextblob(val *FormFieldSlateTextblob) *NullableFormFieldSlateTextblob {
	return &NullableFormFieldSlateTextblob{value: val, isSet: true}
}

func (v NullableFormFieldSlateTextblob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldSlateTextblob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
