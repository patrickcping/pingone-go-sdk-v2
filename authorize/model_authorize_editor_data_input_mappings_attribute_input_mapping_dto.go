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

// checks if the AuthorizeEditorDataInputMappingsAttributeInputMappingDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataInputMappingsAttributeInputMappingDTO{}

// AuthorizeEditorDataInputMappingsAttributeInputMappingDTO struct for AuthorizeEditorDataInputMappingsAttributeInputMappingDTO
type AuthorizeEditorDataInputMappingsAttributeInputMappingDTO struct {
	AuthorizeEditorDataInputMappingDTO
	Value AuthorizeEditorDataReferenceObjectDTO `json:"value"`
}

// NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTO instantiates a new AuthorizeEditorDataInputMappingsAttributeInputMappingDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTO(value AuthorizeEditorDataReferenceObjectDTO, property string, type_ string) *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO {
	this := AuthorizeEditorDataInputMappingsAttributeInputMappingDTO{}
	this.Property = property
	this.Type = type_
	this.Value = value
	return &this
}

// NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTOWithDefaults instantiates a new AuthorizeEditorDataInputMappingsAttributeInputMappingDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTOWithDefaults() *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO {
	this := AuthorizeEditorDataInputMappingsAttributeInputMappingDTO{}
	return &this
}

// GetValue returns the Value field value
func (o *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) GetValue() AuthorizeEditorDataReferenceObjectDTO {
	if o == nil {
		var ret AuthorizeEditorDataReferenceObjectDTO
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) GetValueOk() (*AuthorizeEditorDataReferenceObjectDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) SetValue(v AuthorizeEditorDataReferenceObjectDTO) {
	o.Value = v
}

func (o AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataInputMappingDTO, errAuthorizeEditorDataInputMappingDTO := json.Marshal(o.AuthorizeEditorDataInputMappingDTO)
	if errAuthorizeEditorDataInputMappingDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataInputMappingDTO
	}
	errAuthorizeEditorDataInputMappingDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataInputMappingDTO), &toSerialize)
	if errAuthorizeEditorDataInputMappingDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataInputMappingDTO
	}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO struct {
	value *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) Get() *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) Set(val *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO(val *AuthorizeEditorDataInputMappingsAttributeInputMappingDTO) *NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO {
	return &NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataInputMappingsAttributeInputMappingDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

