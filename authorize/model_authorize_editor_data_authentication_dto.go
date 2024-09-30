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

// checks if the AuthorizeEditorDataAuthenticationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataAuthenticationDTO{}

// AuthorizeEditorDataAuthenticationDTO struct for AuthorizeEditorDataAuthenticationDTO
type AuthorizeEditorDataAuthenticationDTO struct {
	Type EnumAuthorizeEditorDataAuthenticationDTOType `json:"type"`
}

// NewAuthorizeEditorDataAuthenticationDTO instantiates a new AuthorizeEditorDataAuthenticationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataAuthenticationDTO(type_ EnumAuthorizeEditorDataAuthenticationDTOType) *AuthorizeEditorDataAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationDTO{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataAuthenticationDTOWithDefaults instantiates a new AuthorizeEditorDataAuthenticationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataAuthenticationDTOWithDefaults() *AuthorizeEditorDataAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationDTO{}
	return &this
}

// GetType returns the Type field value
func (o *AuthorizeEditorDataAuthenticationDTO) GetType() EnumAuthorizeEditorDataAuthenticationDTOType {
	if o == nil {
		var ret EnumAuthorizeEditorDataAuthenticationDTOType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAuthenticationDTO) GetTypeOk() (*EnumAuthorizeEditorDataAuthenticationDTOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizeEditorDataAuthenticationDTO) SetType(v EnumAuthorizeEditorDataAuthenticationDTOType) {
	o.Type = v
}

func (o AuthorizeEditorDataAuthenticationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataAuthenticationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAuthorizeEditorDataAuthenticationDTO struct {
	value *AuthorizeEditorDataAuthenticationDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAuthenticationDTO) Get() *AuthorizeEditorDataAuthenticationDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAuthenticationDTO) Set(val *AuthorizeEditorDataAuthenticationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAuthenticationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAuthenticationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAuthenticationDTO(val *AuthorizeEditorDataAuthenticationDTO) *NullableAuthorizeEditorDataAuthenticationDTO {
	return &NullableAuthorizeEditorDataAuthenticationDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAuthenticationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAuthenticationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


