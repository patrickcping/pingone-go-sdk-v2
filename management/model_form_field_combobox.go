/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the FormFieldCombobox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldCombobox{}

// FormFieldCombobox struct for FormFieldCombobox
type FormFieldCombobox struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A string that specifies the label shown to the end user for this option.
	Label string `json:"label"`
	// A string that specifies the value of the field if this option is selected.
	Value string `json:"value"`
}

type _FormFieldCombobox FormFieldCombobox

// NewFormFieldCombobox instantiates a new FormFieldCombobox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldCombobox(type_ EnumFormFieldType, position FormFieldCommonPosition, label string, value string) *FormFieldCombobox {
	this := FormFieldCombobox{}
	this.Type = type_
	this.Position = position
	this.Label = label
	this.Value = value
	return &this
}

// NewFormFieldComboboxWithDefaults instantiates a new FormFieldCombobox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldComboboxWithDefaults() *FormFieldCombobox {
	this := FormFieldCombobox{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldCombobox) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldCombobox) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldCombobox) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldCombobox) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldCombobox) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldCombobox) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetLabel returns the Label field value
func (o *FormFieldCombobox) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FormFieldCombobox) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FormFieldCombobox) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *FormFieldCombobox) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *FormFieldCombobox) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *FormFieldCombobox) SetValue(v string) {
	o.Value = v
}

func (o FormFieldCombobox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldCombobox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *FormFieldCombobox) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"position",
		"label",
		"value",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFormFieldCombobox := _FormFieldCombobox{}

	err = json.Unmarshal(bytes, &varFormFieldCombobox)

	if err != nil {
		return err
	}

	*o = FormFieldCombobox(varFormFieldCombobox)

	return err
}

type NullableFormFieldCombobox struct {
	value *FormFieldCombobox
	isSet bool
}

func (v NullableFormFieldCombobox) Get() *FormFieldCombobox {
	return v.value
}

func (v *NullableFormFieldCombobox) Set(val *FormFieldCombobox) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldCombobox) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldCombobox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldCombobox(val *FormFieldCombobox) *NullableFormFieldCombobox {
	return &NullableFormFieldCombobox{value: val, isSet: true}
}

func (v NullableFormFieldCombobox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldCombobox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


