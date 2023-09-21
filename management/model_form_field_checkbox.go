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

// checks if the FormFieldCheckbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldCheckbox{}

// FormFieldCheckbox struct for FormFieldCheckbox
type FormFieldCheckbox struct {
	Type EnumFormFieldType `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A boolean that specifies whether the linked directory attribute is disabled.
	AttributeDisabled *bool `json:"attributeDisabled,omitempty"`
	// A string that specifies an identifier for the field component.
	Key string `json:"key"`
	// A string of escaped JSON that is designed to store a series of text and translatable keys.
	Label *string `json:"label,omitempty"`
	LabelMode *EnumFormElementLabelMode `json:"labelMode,omitempty"`
	// A boolean that specifies whether the field is required.
	Required bool `json:"required"`
	// A boolean that specifies whether the end user can type an entry that is not in a predefined list.
	OtherOptionEnabled *bool `json:"otherOptionEnabled,omitempty"`
	// A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list.
	OtherOptionKey *string `json:"otherOptionKey,omitempty"`
	// A string that specifies the label for a custom or \"other\" choice in a list.
	OtherOptionlabel *string `json:"otherOptionlabel,omitempty"`
	// A string that specifies the label for the other option in drop-down controls.
	OtherOptionInputlabel *string `json:"otherOptionInputlabel,omitempty"`
	// A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute.
	OtherOptionAttributeDisabled *bool `json:"otherOptionAttributeDisabled,omitempty"`
	Layout EnumFormElementLayout `json:"layout"`
	// An array of strings that specifies the unique list of options. This is a required property when the type is `RADIO`, `CHECKBOX`, or `DROPDOWN`.
	Options []string `json:"options"`
	Validation *FormElementValidation `json:"validation,omitempty"`
}

// NewFormFieldCheckbox instantiates a new FormFieldCheckbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldCheckbox(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, required bool, layout EnumFormElementLayout, options []string) *FormFieldCheckbox {
	this := FormFieldCheckbox{}
	this.Type = type_
	this.Position = position
	this.Key = key
	this.Required = required
	this.Layout = layout
	this.Options = options
	return &this
}

// NewFormFieldCheckboxWithDefaults instantiates a new FormFieldCheckbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldCheckboxWithDefaults() *FormFieldCheckbox {
	this := FormFieldCheckbox{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldCheckbox) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldCheckbox) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldCheckbox) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldCheckbox) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetAttributeDisabled returns the AttributeDisabled field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetAttributeDisabled() bool {
	if o == nil || IsNil(o.AttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.AttributeDisabled
}

// GetAttributeDisabledOk returns a tuple with the AttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AttributeDisabled) {
		return nil, false
	}
	return o.AttributeDisabled, true
}

// HasAttributeDisabled returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasAttributeDisabled() bool {
	if o != nil && !IsNil(o.AttributeDisabled) {
		return true
	}

	return false
}

// SetAttributeDisabled gets a reference to the given bool and assigns it to the AttributeDisabled field.
func (o *FormFieldCheckbox) SetAttributeDisabled(v bool) {
	o.AttributeDisabled = &v
}

// GetKey returns the Key field value
func (o *FormFieldCheckbox) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *FormFieldCheckbox) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FormFieldCheckbox) SetLabel(v string) {
	o.Label = &v
}

// GetLabelMode returns the LabelMode field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetLabelMode() EnumFormElementLabelMode {
	if o == nil || IsNil(o.LabelMode) {
		var ret EnumFormElementLabelMode
		return ret
	}
	return *o.LabelMode
}

// GetLabelModeOk returns a tuple with the LabelMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetLabelModeOk() (*EnumFormElementLabelMode, bool) {
	if o == nil || IsNil(o.LabelMode) {
		return nil, false
	}
	return o.LabelMode, true
}

// HasLabelMode returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasLabelMode() bool {
	if o != nil && !IsNil(o.LabelMode) {
		return true
	}

	return false
}

// SetLabelMode gets a reference to the given EnumFormElementLabelMode and assigns it to the LabelMode field.
func (o *FormFieldCheckbox) SetLabelMode(v EnumFormElementLabelMode) {
	o.LabelMode = &v
}

// GetRequired returns the Required field value
func (o *FormFieldCheckbox) GetRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *FormFieldCheckbox) SetRequired(v bool) {
	o.Required = v
}

// GetOtherOptionEnabled returns the OtherOptionEnabled field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetOtherOptionEnabled() bool {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionEnabled
}

// GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOtherOptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionEnabled) {
		return nil, false
	}
	return o.OtherOptionEnabled, true
}

// HasOtherOptionEnabled returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasOtherOptionEnabled() bool {
	if o != nil && !IsNil(o.OtherOptionEnabled) {
		return true
	}

	return false
}

// SetOtherOptionEnabled gets a reference to the given bool and assigns it to the OtherOptionEnabled field.
func (o *FormFieldCheckbox) SetOtherOptionEnabled(v bool) {
	o.OtherOptionEnabled = &v
}

// GetOtherOptionKey returns the OtherOptionKey field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetOtherOptionKey() string {
	if o == nil || IsNil(o.OtherOptionKey) {
		var ret string
		return ret
	}
	return *o.OtherOptionKey
}

// GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOtherOptionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionKey) {
		return nil, false
	}
	return o.OtherOptionKey, true
}

// HasOtherOptionKey returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasOtherOptionKey() bool {
	if o != nil && !IsNil(o.OtherOptionKey) {
		return true
	}

	return false
}

// SetOtherOptionKey gets a reference to the given string and assigns it to the OtherOptionKey field.
func (o *FormFieldCheckbox) SetOtherOptionKey(v string) {
	o.OtherOptionKey = &v
}

// GetOtherOptionlabel returns the OtherOptionlabel field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetOtherOptionlabel() string {
	if o == nil || IsNil(o.OtherOptionlabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionlabel
}

// GetOtherOptionlabelOk returns a tuple with the OtherOptionlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOtherOptionlabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionlabel) {
		return nil, false
	}
	return o.OtherOptionlabel, true
}

// HasOtherOptionlabel returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasOtherOptionlabel() bool {
	if o != nil && !IsNil(o.OtherOptionlabel) {
		return true
	}

	return false
}

// SetOtherOptionlabel gets a reference to the given string and assigns it to the OtherOptionlabel field.
func (o *FormFieldCheckbox) SetOtherOptionlabel(v string) {
	o.OtherOptionlabel = &v
}

// GetOtherOptionInputlabel returns the OtherOptionInputlabel field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetOtherOptionInputlabel() string {
	if o == nil || IsNil(o.OtherOptionInputlabel) {
		var ret string
		return ret
	}
	return *o.OtherOptionInputlabel
}

// GetOtherOptionInputlabelOk returns a tuple with the OtherOptionInputlabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOtherOptionInputlabelOk() (*string, bool) {
	if o == nil || IsNil(o.OtherOptionInputlabel) {
		return nil, false
	}
	return o.OtherOptionInputlabel, true
}

// HasOtherOptionInputlabel returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasOtherOptionInputlabel() bool {
	if o != nil && !IsNil(o.OtherOptionInputlabel) {
		return true
	}

	return false
}

// SetOtherOptionInputlabel gets a reference to the given string and assigns it to the OtherOptionInputlabel field.
func (o *FormFieldCheckbox) SetOtherOptionInputlabel(v string) {
	o.OtherOptionInputlabel = &v
}

// GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetOtherOptionAttributeDisabled() bool {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		var ret bool
		return ret
	}
	return *o.OtherOptionAttributeDisabled
}

// GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOtherOptionAttributeDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OtherOptionAttributeDisabled) {
		return nil, false
	}
	return o.OtherOptionAttributeDisabled, true
}

// HasOtherOptionAttributeDisabled returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasOtherOptionAttributeDisabled() bool {
	if o != nil && !IsNil(o.OtherOptionAttributeDisabled) {
		return true
	}

	return false
}

// SetOtherOptionAttributeDisabled gets a reference to the given bool and assigns it to the OtherOptionAttributeDisabled field.
func (o *FormFieldCheckbox) SetOtherOptionAttributeDisabled(v bool) {
	o.OtherOptionAttributeDisabled = &v
}

// GetLayout returns the Layout field value
func (o *FormFieldCheckbox) GetLayout() EnumFormElementLayout {
	if o == nil {
		var ret EnumFormElementLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetLayoutOk() (*EnumFormElementLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *FormFieldCheckbox) SetLayout(v EnumFormElementLayout) {
	o.Layout = v
}

// GetOptions returns the Options field value
func (o *FormFieldCheckbox) GetOptions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetOptionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *FormFieldCheckbox) SetOptions(v []string) {
	o.Options = v
}

// GetValidation returns the Validation field value if set, zero value otherwise.
func (o *FormFieldCheckbox) GetValidation() FormElementValidation {
	if o == nil || IsNil(o.Validation) {
		var ret FormElementValidation
		return ret
	}
	return *o.Validation
}

// GetValidationOk returns a tuple with the Validation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldCheckbox) GetValidationOk() (*FormElementValidation, bool) {
	if o == nil || IsNil(o.Validation) {
		return nil, false
	}
	return o.Validation, true
}

// HasValidation returns a boolean if a field has been set.
func (o *FormFieldCheckbox) HasValidation() bool {
	if o != nil && !IsNil(o.Validation) {
		return true
	}

	return false
}

// SetValidation gets a reference to the given FormElementValidation and assigns it to the Validation field.
func (o *FormFieldCheckbox) SetValidation(v FormElementValidation) {
	o.Validation = &v
}

func (o FormFieldCheckbox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldCheckbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	if !IsNil(o.AttributeDisabled) {
		toSerialize["attributeDisabled"] = o.AttributeDisabled
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LabelMode) {
		toSerialize["labelMode"] = o.LabelMode
	}
	toSerialize["required"] = o.Required
	if !IsNil(o.OtherOptionEnabled) {
		toSerialize["otherOptionEnabled"] = o.OtherOptionEnabled
	}
	if !IsNil(o.OtherOptionKey) {
		toSerialize["otherOptionKey"] = o.OtherOptionKey
	}
	if !IsNil(o.OtherOptionlabel) {
		toSerialize["otherOptionlabel"] = o.OtherOptionlabel
	}
	if !IsNil(o.OtherOptionInputlabel) {
		toSerialize["otherOptionInputlabel"] = o.OtherOptionInputlabel
	}
	if !IsNil(o.OtherOptionAttributeDisabled) {
		toSerialize["otherOptionAttributeDisabled"] = o.OtherOptionAttributeDisabled
	}
	toSerialize["layout"] = o.Layout
	toSerialize["options"] = o.Options
	if !IsNil(o.Validation) {
		toSerialize["validation"] = o.Validation
	}
	return toSerialize, nil
}

type NullableFormFieldCheckbox struct {
	value *FormFieldCheckbox
	isSet bool
}

func (v NullableFormFieldCheckbox) Get() *FormFieldCheckbox {
	return v.value
}

func (v *NullableFormFieldCheckbox) Set(val *FormFieldCheckbox) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldCheckbox) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldCheckbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldCheckbox(val *FormFieldCheckbox) *NullableFormFieldCheckbox {
	return &NullableFormFieldCheckbox{value: val, isSet: true}
}

func (v NullableFormFieldCheckbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldCheckbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


