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

// checks if the FormFieldText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldText{}

// FormFieldText struct for FormFieldText
type FormFieldText struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A boolean that specifies whether the linked directory attribute is disabled.
	AttributeDisabled *bool `json:"attributeDisabled,omitempty"`
	// A string that specifies an identifier for the field component.
	Key string `json:"key"`
	// A string of escaped JSON that is designed to store a series of text and translatable keys.
	Label string `json:"label"`
	LabelMode *EnumFormElementLabelMode `json:"labelMode,omitempty"`
	// A boolean that specifies whether the field is required.
	Required *bool `json:"required,omitempty"`
	// A boolean that specifies whether the end user can type an entry that is not in a predefined list.
	OtherOptionEnabled *bool `json:"otherOptionEnabled,omitempty"`
	// A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list.
	OtherOptionKey *string `json:"otherOptionKey,omitempty"`
	// A string that specifies the label for a custom or \"other\" choice in a list.
	OtherOptionLabel *string `json:"otherOptionLabel,omitempty"`
	// A string that specifies the label for the other option in drop-down controls.
	OtherOptionInputLabel *string `json:"otherOptionInputLabel,omitempty"`
	// A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute.
	OtherOptionAttributeDisabled *bool `json:"otherOptionAttributeDisabled,omitempty"`
	Layout *EnumFormElementLayout `json:"layout,omitempty"`
	// An array of objects (label/value pairs) that specifies the unique list of options. This is a required property when the type is `RADIO`, `CHECKBOX`, or `DROPDOWN`.
	Options []FormElementOption `json:"options,omitempty"`
	Validation FormElementValidation `json:"validation"`
}

// NewFormFieldText instantiates a new FormFieldText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldText(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, validation FormElementValidation) *FormFieldText {
	this := FormFieldText{}
	this.Type = type_
	this.Position = position
	this.Key = key
	this.Label = label
	this.Validation = validation
	return &this
}

// NewFormFieldTextWithDefaults instantiates a new FormFieldText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldTextWithDefaults() *FormFieldText {
	this := FormFieldText{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldText) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldText) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldText) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldText) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetAttributeDisabled returns the AttributeDisabled field value if set, zero value otherwise.
func (o *FormFieldText) GetAttributeDisabled() bool {
	if o == nil || IsNil(o.AttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.AttributeDisabled
}

// GetAttributeDisabledOk returns a tuple with the AttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AttributeDisabled) {
		return nil, false
	}
	return o.AttributeDisabled, true
}

// HasAttributeDisabled returns a boolean if a field has been set.
func (o *FormFieldText) HasAttributeDisabled() bool {
	if o != nil && !IsNil(o.AttributeDisabled) {
		return true
	}

	return false
}

// SetAttributeDisabled gets a reference to the given bool and assigns it to the AttributeDisabled field.
func (o *FormFieldText) SetAttributeDisabled(v bool) {
	o.AttributeDisabled = &v
}

// GetKey returns the Key field value
func (o *FormFieldText) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FormFieldText) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *FormFieldText) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FormFieldText) SetLabel(v string) {
	o.Label = v
}

// GetLabelMode returns the LabelMode field value if set, zero value otherwise.
func (o *FormFieldText) GetLabelMode() EnumFormElementLabelMode {
	if o == nil || IsNil(o.LabelMode) {
		var ret EnumFormElementLabelMode
		return ret
	}
	return *o.LabelMode
}

// GetLabelModeOk returns a tuple with the LabelMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetLabelModeOk() (*EnumFormElementLabelMode, bool) {
	if o == nil || IsNil(o.LabelMode) {
		return nil, false
	}
	return o.LabelMode, true
}

// HasLabelMode returns a boolean if a field has been set.
func (o *FormFieldText) HasLabelMode() bool {
	if o != nil && !IsNil(o.LabelMode) {
		return true
	}

	return false
}

// SetLabelMode gets a reference to the given EnumFormElementLabelMode and assigns it to the LabelMode field.
func (o *FormFieldText) SetLabelMode(v EnumFormElementLabelMode) {
	o.LabelMode = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FormFieldText) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FormFieldText) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FormFieldText) SetRequired(v bool) {
	o.Required = &v
}

// GetOtherOptionEnabled returns the OtherOptionEnabled field value if set, zero value otherwise.
func (o *FormFieldText) GetOtherOptionEnabled() bool {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionEnabled
}

// GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOtherOptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		return nil, false
	}
	return o.OtherOptionEnabled, true
}

// HasOtherOptionEnabled returns a boolean if a field has been set.
func (o *FormFieldText) HasOtherOptionEnabled() bool {
	if o != nil && !IsNil(o.OtherOptionEnabled) {
		return true
	}

	return false
}

// SetOtherOptionEnabled gets a reference to the given bool and assigns it to the OtherOptionEnabled field.
func (o *FormFieldText) SetOtherOptionEnabled(v bool) {
	o.OtherOptionEnabled = &v
}

// GetOtherOptionKey returns the OtherOptionKey field value if set, zero value otherwise.
func (o *FormFieldText) GetOtherOptionKey() string {
	if o == nil || IsNil(o.OtherOptionKey) {
		var ret string
		return ret
	}
	return *o.OtherOptionKey
}

// GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOtherOptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionKey) {
		return nil, false
	}
	return o.OtherOptionKey, true
}

// HasOtherOptionKey returns a boolean if a field has been set.
func (o *FormFieldText) HasOtherOptionKey() bool {
	if o != nil && !IsNil(o.OtherOptionKey) {
		return true
	}

	return false
}

// SetOtherOptionKey gets a reference to the given string and assigns it to the OtherOptionKey field.
func (o *FormFieldText) SetOtherOptionKey(v string) {
	o.OtherOptionKey = &v
}

// GetOtherOptionLabel returns the OtherOptionLabel field value if set, zero value otherwise.
func (o *FormFieldText) GetOtherOptionLabel() string {
	if o == nil || IsNil(o.OtherOptionLabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionLabel
}

// GetOtherOptionLabelOk returns a tuple with the OtherOptionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOtherOptionLabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionLabel) {
		return nil, false
	}
	return o.OtherOptionLabel, true
}

// HasOtherOptionLabel returns a boolean if a field has been set.
func (o *FormFieldText) HasOtherOptionLabel() bool {
	if o != nil && !IsNil(o.OtherOptionLabel) {
		return true
	}

	return false
}

// SetOtherOptionLabel gets a reference to the given string and assigns it to the OtherOptionLabel field.
func (o *FormFieldText) SetOtherOptionLabel(v string) {
	o.OtherOptionLabel = &v
}

// GetOtherOptionInputLabel returns the OtherOptionInputLabel field value if set, zero value otherwise.
func (o *FormFieldText) GetOtherOptionInputLabel() string {
	if o == nil || IsNil(o.OtherOptionInputLabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionInputLabel
}

// GetOtherOptionInputLabelOk returns a tuple with the OtherOptionInputLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOtherOptionInputLabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionInputLabel) {
		return nil, false
	}
	return o.OtherOptionInputLabel, true
}

// HasOtherOptionInputLabel returns a boolean if a field has been set.
func (o *FormFieldText) HasOtherOptionInputLabel() bool {
	if o != nil && !IsNil(o.OtherOptionInputLabel) {
		return true
	}

	return false
}

// SetOtherOptionInputLabel gets a reference to the given string and assigns it to the OtherOptionInputLabel field.
func (o *FormFieldText) SetOtherOptionInputLabel(v string) {
	o.OtherOptionInputLabel = &v
}

// GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field value if set, zero value otherwise.
func (o *FormFieldText) GetOtherOptionAttributeDisabled() bool {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionAttributeDisabled
}

// GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOtherOptionAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		return nil, false
	}
	return o.OtherOptionAttributeDisabled, true
}

// HasOtherOptionAttributeDisabled returns a boolean if a field has been set.
func (o *FormFieldText) HasOtherOptionAttributeDisabled() bool {
	if o != nil && !IsNil(o.OtherOptionAttributeDisabled) {
		return true
	}

	return false
}

// SetOtherOptionAttributeDisabled gets a reference to the given bool and assigns it to the OtherOptionAttributeDisabled field.
func (o *FormFieldText) SetOtherOptionAttributeDisabled(v bool) {
	o.OtherOptionAttributeDisabled = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *FormFieldText) GetLayout() EnumFormElementLayout {
	if o == nil || IsNil(o.Layout) {
		var ret EnumFormElementLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetLayoutOk() (*EnumFormElementLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *FormFieldText) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given EnumFormElementLayout and assigns it to the Layout field.
func (o *FormFieldText) SetLayout(v EnumFormElementLayout) {
	o.Layout = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *FormFieldText) GetOptions() []FormElementOption {
	if o == nil || IsNil(o.Options) {
		var ret []FormElementOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetOptionsOk() ([]FormElementOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FormFieldText) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []FormElementOption and assigns it to the Options field.
func (o *FormFieldText) SetOptions(v []FormElementOption) {
	o.Options = v
}

// GetValidation returns the Validation field value
func (o *FormFieldText) GetValidation() FormElementValidation {
	if o == nil {
		var ret FormElementValidation
		return ret
	}

	return o.Validation
}

// GetValidationOk returns a tuple with the Validation field value
// and a boolean to check if the value has been set.
func (o *FormFieldText) GetValidationOk() (*FormElementValidation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validation, true
}

// SetValidation sets field value
func (o *FormFieldText) SetValidation(v FormElementValidation) {
	o.Validation = v
}

func (o FormFieldText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	if !IsNil(o.AttributeDisabled) {
		toSerialize["attributeDisabled"] = o.AttributeDisabled
	}
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if !IsNil(o.LabelMode) {
		toSerialize["labelMode"] = o.LabelMode
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.OtherOptionEnabled) {
		toSerialize["otherOptionEnabled"] = o.OtherOptionEnabled
	}
	if !IsNil(o.OtherOptionKey) {
		toSerialize["otherOptionKey"] = o.OtherOptionKey
	}
	if !IsNil(o.OtherOptionLabel) {
		toSerialize["otherOptionLabel"] = o.OtherOptionLabel
	}
	if !IsNil(o.OtherOptionInputLabel) {
		toSerialize["otherOptionInputLabel"] = o.OtherOptionInputLabel
	}
	if !IsNil(o.OtherOptionAttributeDisabled) {
		toSerialize["otherOptionAttributeDisabled"] = o.OtherOptionAttributeDisabled
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["validation"] = o.Validation
	return toSerialize, nil
}

type NullableFormFieldText struct {
	value *FormFieldText
	isSet bool
}

func (v NullableFormFieldText) Get() *FormFieldText {
	return v.value
}

func (v *NullableFormFieldText) Set(val *FormFieldText) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldText) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldText(val *FormFieldText) *NullableFormFieldText {
	return &NullableFormFieldText{value: val, isSet: true}
}

func (v NullableFormFieldText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


