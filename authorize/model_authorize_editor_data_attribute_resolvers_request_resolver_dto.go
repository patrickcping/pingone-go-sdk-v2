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

// checks if the AuthorizeEditorDataAttributeResolversRequestResolverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataAttributeResolversRequestResolverDTO{}

// AuthorizeEditorDataAttributeResolversRequestResolverDTO struct for AuthorizeEditorDataAttributeResolversRequestResolverDTO
type AuthorizeEditorDataAttributeResolversRequestResolverDTO struct {
	AuthorizeEditorDataResolverDTO
}

// NewAuthorizeEditorDataAttributeResolversRequestResolverDTO instantiates a new AuthorizeEditorDataAttributeResolversRequestResolverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataAttributeResolversRequestResolverDTO(type_ string) *AuthorizeEditorDataAttributeResolversRequestResolverDTO {
	this := AuthorizeEditorDataAttributeResolversRequestResolverDTO{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataAttributeResolversRequestResolverDTOWithDefaults instantiates a new AuthorizeEditorDataAttributeResolversRequestResolverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataAttributeResolversRequestResolverDTOWithDefaults() *AuthorizeEditorDataAttributeResolversRequestResolverDTO {
	this := AuthorizeEditorDataAttributeResolversRequestResolverDTO{}
	return &this
}

func (o AuthorizeEditorDataAttributeResolversRequestResolverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataAttributeResolversRequestResolverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataResolverDTO, errAuthorizeEditorDataResolverDTO := json.Marshal(o.AuthorizeEditorDataResolverDTO)
	if errAuthorizeEditorDataResolverDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataResolverDTO
	}
	errAuthorizeEditorDataResolverDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataResolverDTO), &toSerialize)
	if errAuthorizeEditorDataResolverDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataResolverDTO
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO struct {
	value *AuthorizeEditorDataAttributeResolversRequestResolverDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) Get() *AuthorizeEditorDataAttributeResolversRequestResolverDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) Set(val *AuthorizeEditorDataAttributeResolversRequestResolverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAttributeResolversRequestResolverDTO(val *AuthorizeEditorDataAttributeResolversRequestResolverDTO) *NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO {
	return &NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAttributeResolversRequestResolverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

