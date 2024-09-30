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

// checks if the AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO{}

// AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO struct for AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO
type AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO struct {
	AuthorizeEditorDataProcessorDTO
	Predicate AuthorizeEditorDataProcessorDTO `json:"predicate"`
}

// NewAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO instantiates a new AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO(predicate AuthorizeEditorDataProcessorDTO, name string, type_ EnumAuthorizeEditorDataProcessorDTOType) *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO {
	this := AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO{}
	this.Name = name
	this.Type = type_
	this.Predicate = predicate
	return &this
}

// NewAuthorizeEditorDataProcessorsCollectionFilterProcessorDTOWithDefaults instantiates a new AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataProcessorsCollectionFilterProcessorDTOWithDefaults() *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO {
	this := AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO{}
	return &this
}

// GetPredicate returns the Predicate field value
func (o *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) GetPredicate() AuthorizeEditorDataProcessorDTO {
	if o == nil {
		var ret AuthorizeEditorDataProcessorDTO
		return ret
	}

	return o.Predicate
}

// GetPredicateOk returns a tuple with the Predicate field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) GetPredicateOk() (*AuthorizeEditorDataProcessorDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Predicate, true
}

// SetPredicate sets field value
func (o *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) SetPredicate(v AuthorizeEditorDataProcessorDTO) {
	o.Predicate = v
}

func (o AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataProcessorDTO, errAuthorizeEditorDataProcessorDTO := json.Marshal(o.AuthorizeEditorDataProcessorDTO)
	if errAuthorizeEditorDataProcessorDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataProcessorDTO
	}
	errAuthorizeEditorDataProcessorDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataProcessorDTO), &toSerialize)
	if errAuthorizeEditorDataProcessorDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataProcessorDTO
	}
	toSerialize["predicate"] = o.Predicate
	return toSerialize, nil
}

type NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO struct {
	value *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) Get() *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) Set(val *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO(val *AuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) *NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO {
	return &NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataProcessorsCollectionFilterProcessorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


