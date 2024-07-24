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

// checks if the AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO{}

// AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO struct for AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO
type AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO struct {
	AuthorizeEditorDataAuthenticationDTO
}

// NewAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO instantiates a new AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO(type_ string) *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataAuthenticationsNoneAuthenticationDTOWithDefaults instantiates a new AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataAuthenticationsNoneAuthenticationDTOWithDefaults() *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO{}
	return &this
}

func (o AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataAuthenticationDTO, errAuthorizeEditorDataAuthenticationDTO := json.Marshal(o.AuthorizeEditorDataAuthenticationDTO)
	if errAuthorizeEditorDataAuthenticationDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataAuthenticationDTO
	}
	errAuthorizeEditorDataAuthenticationDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataAuthenticationDTO), &toSerialize)
	if errAuthorizeEditorDataAuthenticationDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataAuthenticationDTO
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO struct {
	value *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) Get() *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) Set(val *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO(val *AuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) *NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO {
	return &NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAuthenticationsNoneAuthenticationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

