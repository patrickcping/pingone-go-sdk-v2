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

// checks if the AuthorizeEditorDataAttributeResolversAttributeResolverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataAttributeResolversAttributeResolverDTO{}

// AuthorizeEditorDataAttributeResolversAttributeResolverDTO struct for AuthorizeEditorDataAttributeResolversAttributeResolverDTO
type AuthorizeEditorDataAttributeResolversAttributeResolverDTO struct {
	AuthorizeEditorDataResolverDTO
	Value AuthorizeEditorDataReferenceObjectDTO `json:"value"`
}

// NewAuthorizeEditorDataAttributeResolversAttributeResolverDTO instantiates a new AuthorizeEditorDataAttributeResolversAttributeResolverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataAttributeResolversAttributeResolverDTO(value AuthorizeEditorDataReferenceObjectDTO, type_ string) *AuthorizeEditorDataAttributeResolversAttributeResolverDTO {
	this := AuthorizeEditorDataAttributeResolversAttributeResolverDTO{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewAuthorizeEditorDataAttributeResolversAttributeResolverDTOWithDefaults instantiates a new AuthorizeEditorDataAttributeResolversAttributeResolverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataAttributeResolversAttributeResolverDTOWithDefaults() *AuthorizeEditorDataAttributeResolversAttributeResolverDTO {
	this := AuthorizeEditorDataAttributeResolversAttributeResolverDTO{}
	return &this
}

// GetValue returns the Value field value
func (o *AuthorizeEditorDataAttributeResolversAttributeResolverDTO) GetValue() AuthorizeEditorDataReferenceObjectDTO {
	if o == nil {
		var ret AuthorizeEditorDataReferenceObjectDTO
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAttributeResolversAttributeResolverDTO) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AuthorizeEditorDataAttributeResolversAttributeResolverDTO) SetValue(v AuthorizeEditorDataReferenceObjectDTO) {
	o.Value = v
}

func (o AuthorizeEditorDataAttributeResolversAttributeResolverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataAttributeResolversAttributeResolverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataResolverDTO, errAuthorizeEditorDataResolverDTO := json.Marshal(o.AuthorizeEditorDataResolverDTO)
	if errAuthorizeEditorDataResolverDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataResolverDTO
	}
	errAuthorizeEditorDataResolverDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataResolverDTO), &toSerialize)
	if errAuthorizeEditorDataResolverDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataResolverDTO
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO struct {
	value *AuthorizeEditorDataAttributeResolversAttributeResolverDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) Get() *AuthorizeEditorDataAttributeResolversAttributeResolverDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) Set(val *AuthorizeEditorDataAttributeResolversAttributeResolverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO(val *AuthorizeEditorDataAttributeResolversAttributeResolverDTO) *NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO {
	return &NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAttributeResolversAttributeResolverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

