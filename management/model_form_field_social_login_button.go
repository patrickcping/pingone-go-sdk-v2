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

// checks if the FormFieldSocialLoginButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormFieldSocialLoginButton{}

// FormFieldSocialLoginButton struct for FormFieldSocialLoginButton
type FormFieldSocialLoginButton struct {
	Type     EnumFormFieldType       `json:"type"`
	Position FormFieldCommonPosition `json:"position"`
	// A string that specifies an identifier for the field component.
	Key *string `json:"key,omitempty"`
	// A string that specifies the social login button label.
	Label   string                     `json:"label"`
	Styles  *FormStyles                `json:"styles,omitempty"`
	IdpType EnumFormSocialLoginIdpType `json:"idpType"`
	// A string that specifies the external identity provider name.
	IdpName string `json:"idpName"`
	// A string that specifies the external identity provider's ID.
	IdpId string `json:"idpId"`
	// A boolean that specifies whether the external identity provider is enabled.
	IdpEnabled bool `json:"idpEnabled"`
}

// NewFormFieldSocialLoginButton instantiates a new FormFieldSocialLoginButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormFieldSocialLoginButton(type_ EnumFormFieldType, position FormFieldCommonPosition, label string, idpType EnumFormSocialLoginIdpType, idpName string, idpId string, idpEnabled bool) *FormFieldSocialLoginButton {
	this := FormFieldSocialLoginButton{}
	this.Type = type_
	this.Position = position
	this.Label = label
	this.IdpType = idpType
	this.IdpName = idpName
	this.IdpId = idpId
	this.IdpEnabled = idpEnabled
	return &this
}

// NewFormFieldSocialLoginButtonWithDefaults instantiates a new FormFieldSocialLoginButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormFieldSocialLoginButtonWithDefaults() *FormFieldSocialLoginButton {
	this := FormFieldSocialLoginButton{}
	return &this
}

// GetType returns the Type field value
func (o *FormFieldSocialLoginButton) GetType() EnumFormFieldType {
	if o == nil {
		var ret EnumFormFieldType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetTypeOk() (*EnumFormFieldType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormFieldSocialLoginButton) SetType(v EnumFormFieldType) {
	o.Type = v
}

// GetPosition returns the Position field value
func (o *FormFieldSocialLoginButton) GetPosition() FormFieldCommonPosition {
	if o == nil {
		var ret FormFieldCommonPosition
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetPositionOk() (*FormFieldCommonPosition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *FormFieldSocialLoginButton) SetPosition(v FormFieldCommonPosition) {
	o.Position = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FormFieldSocialLoginButton) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FormFieldSocialLoginButton) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FormFieldSocialLoginButton) SetKey(v string) {
	o.Key = &v
}

// GetLabel returns the Label field value
func (o *FormFieldSocialLoginButton) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FormFieldSocialLoginButton) SetLabel(v string) {
	o.Label = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *FormFieldSocialLoginButton) GetStyles() FormStyles {
	if o == nil || IsNil(o.Styles) {
		var ret FormStyles
		return ret
	}
	return *o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetStylesOk() (*FormStyles, bool) {
	if o == nil || IsNil(o.Styles) {
		return nil, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *FormFieldSocialLoginButton) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given FormStyles and assigns it to the Styles field.
func (o *FormFieldSocialLoginButton) SetStyles(v FormStyles) {
	o.Styles = &v
}

// GetIdpType returns the IdpType field value
func (o *FormFieldSocialLoginButton) GetIdpType() EnumFormSocialLoginIdpType {
	if o == nil {
		var ret EnumFormSocialLoginIdpType
		return ret
	}

	return o.IdpType
}

// GetIdpTypeOk returns a tuple with the IdpType field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetIdpTypeOk() (*EnumFormSocialLoginIdpType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpType, true
}

// SetIdpType sets field value
func (o *FormFieldSocialLoginButton) SetIdpType(v EnumFormSocialLoginIdpType) {
	o.IdpType = v
}

// GetIdpName returns the IdpName field value
func (o *FormFieldSocialLoginButton) GetIdpName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpName
}

// GetIdpNameOk returns a tuple with the IdpName field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetIdpNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpName, true
}

// SetIdpName sets field value
func (o *FormFieldSocialLoginButton) SetIdpName(v string) {
	o.IdpName = v
}

// GetIdpId returns the IdpId field value
func (o *FormFieldSocialLoginButton) GetIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpId, true
}

// SetIdpId sets field value
func (o *FormFieldSocialLoginButton) SetIdpId(v string) {
	o.IdpId = v
}

// GetIdpEnabled returns the IdpEnabled field value
func (o *FormFieldSocialLoginButton) GetIdpEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IdpEnabled
}

// GetIdpEnabledOk returns a tuple with the IdpEnabled field value
// and a boolean to check if the value has been set.
func (o *FormFieldSocialLoginButton) GetIdpEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdpEnabled, true
}

// SetIdpEnabled sets field value
func (o *FormFieldSocialLoginButton) SetIdpEnabled(v bool) {
	o.IdpEnabled = v
}

func (o FormFieldSocialLoginButton) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormFieldSocialLoginButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["position"] = o.Position
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	toSerialize["idpType"] = o.IdpType
	toSerialize["idpName"] = o.IdpName
	toSerialize["idpId"] = o.IdpId
	toSerialize["idpEnabled"] = o.IdpEnabled
	return toSerialize, nil
}

type NullableFormFieldSocialLoginButton struct {
	value *FormFieldSocialLoginButton
	isSet bool
}

func (v NullableFormFieldSocialLoginButton) Get() *FormFieldSocialLoginButton {
	return v.value
}

func (v *NullableFormFieldSocialLoginButton) Set(val *FormFieldSocialLoginButton) {
	v.value = val
	v.isSet = true
}

func (v NullableFormFieldSocialLoginButton) IsSet() bool {
	return v.isSet
}

func (v *NullableFormFieldSocialLoginButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormFieldSocialLoginButton(val *FormFieldSocialLoginButton) *NullableFormFieldSocialLoginButton {
	return &NullableFormFieldSocialLoginButton{value: val, isSet: true}
}

func (v NullableFormFieldSocialLoginButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormFieldSocialLoginButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
