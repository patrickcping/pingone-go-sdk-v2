/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the AuthorizeEditorDataConditionsComparandDTOCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataConditionsComparandDTOCommon{}

// AuthorizeEditorDataConditionsComparandDTOCommon struct for AuthorizeEditorDataConditionsComparandDTOCommon
type AuthorizeEditorDataConditionsComparandDTOCommon struct {
	Type string `json:"type"`
}

// NewAuthorizeEditorDataConditionsComparandDTOCommon instantiates a new AuthorizeEditorDataConditionsComparandDTOCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataConditionsComparandDTOCommon(type_ string) *AuthorizeEditorDataConditionsComparandDTOCommon {
	this := AuthorizeEditorDataConditionsComparandDTOCommon{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataConditionsComparandDTOCommonWithDefaults instantiates a new AuthorizeEditorDataConditionsComparandDTOCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataConditionsComparandDTOCommonWithDefaults() *AuthorizeEditorDataConditionsComparandDTOCommon {
	this := AuthorizeEditorDataConditionsComparandDTOCommon{}
	return &this
}

// GetType returns the Type field value
func (o *AuthorizeEditorDataConditionsComparandDTOCommon) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConditionsComparandDTOCommon) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizeEditorDataConditionsComparandDTOCommon) SetType(v string) {
	o.Type = v
}

func (o AuthorizeEditorDataConditionsComparandDTOCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataConditionsComparandDTOCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthorizeEditorDataConditionsComparandDTOCommon struct {
	value *AuthorizeEditorDataConditionsComparandDTOCommon
	isSet bool
}

func (v NullableAuthorizeEditorDataConditionsComparandDTOCommon) Get() *AuthorizeEditorDataConditionsComparandDTOCommon {
	return v.value
}

func (v *NullableAuthorizeEditorDataConditionsComparandDTOCommon) Set(val *AuthorizeEditorDataConditionsComparandDTOCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataConditionsComparandDTOCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataConditionsComparandDTOCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataConditionsComparandDTOCommon(val *AuthorizeEditorDataConditionsComparandDTOCommon) *NullableAuthorizeEditorDataConditionsComparandDTOCommon {
	return &NullableAuthorizeEditorDataConditionsComparandDTOCommon{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataConditionsComparandDTOCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataConditionsComparandDTOCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


