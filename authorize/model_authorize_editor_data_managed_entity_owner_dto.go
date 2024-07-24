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

// checks if the AuthorizeEditorDataManagedEntityOwnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataManagedEntityOwnerDTO{}

// AuthorizeEditorDataManagedEntityOwnerDTO struct for AuthorizeEditorDataManagedEntityOwnerDTO
type AuthorizeEditorDataManagedEntityOwnerDTO struct {
	Service AuthorizeEditorDataServiceObjectDTO `json:"service"`
}

// NewAuthorizeEditorDataManagedEntityOwnerDTO instantiates a new AuthorizeEditorDataManagedEntityOwnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataManagedEntityOwnerDTO(service AuthorizeEditorDataServiceObjectDTO) *AuthorizeEditorDataManagedEntityOwnerDTO {
	this := AuthorizeEditorDataManagedEntityOwnerDTO{}
	this.Service = service
	return &this
}

// NewAuthorizeEditorDataManagedEntityOwnerDTOWithDefaults instantiates a new AuthorizeEditorDataManagedEntityOwnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataManagedEntityOwnerDTOWithDefaults() *AuthorizeEditorDataManagedEntityOwnerDTO {
	this := AuthorizeEditorDataManagedEntityOwnerDTO{}
	return &this
}

// GetService returns the Service field value
func (o *AuthorizeEditorDataManagedEntityOwnerDTO) GetService() AuthorizeEditorDataServiceObjectDTO {
	if o == nil {
		var ret AuthorizeEditorDataServiceObjectDTO
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataManagedEntityOwnerDTO) GetServiceOk() (*AuthorizeEditorDataServiceObjectDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *AuthorizeEditorDataManagedEntityOwnerDTO) SetService(v AuthorizeEditorDataServiceObjectDTO) {
	o.Service = v
}

func (o AuthorizeEditorDataManagedEntityOwnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataManagedEntityOwnerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["service"] = o.Service
	return toSerialize, nil
}

type NullableAuthorizeEditorDataManagedEntityOwnerDTO struct {
	value *AuthorizeEditorDataManagedEntityOwnerDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataManagedEntityOwnerDTO) Get() *AuthorizeEditorDataManagedEntityOwnerDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataManagedEntityOwnerDTO) Set(val *AuthorizeEditorDataManagedEntityOwnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataManagedEntityOwnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataManagedEntityOwnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataManagedEntityOwnerDTO(val *AuthorizeEditorDataManagedEntityOwnerDTO) *NullableAuthorizeEditorDataManagedEntityOwnerDTO {
	return &NullableAuthorizeEditorDataManagedEntityOwnerDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataManagedEntityOwnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataManagedEntityOwnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

