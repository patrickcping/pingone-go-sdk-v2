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

// checks if the AuthorizeEditorDataProcessorsChainProcessorDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataProcessorsChainProcessorDTO{}

// AuthorizeEditorDataProcessorsChainProcessorDTO struct for AuthorizeEditorDataProcessorsChainProcessorDTO
type AuthorizeEditorDataProcessorsChainProcessorDTO struct {
	AuthorizeEditorDataProcessorDTO
	Processors []AuthorizeEditorDataProcessorDTO `json:"processors"`
}

// NewAuthorizeEditorDataProcessorsChainProcessorDTO instantiates a new AuthorizeEditorDataProcessorsChainProcessorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataProcessorsChainProcessorDTO(processors []AuthorizeEditorDataProcessorDTO, name string, type_ string) *AuthorizeEditorDataProcessorsChainProcessorDTO {
	this := AuthorizeEditorDataProcessorsChainProcessorDTO{}
	this.Name = name
	this.Type = type_
	this.Processors = processors
	return &this
}

// NewAuthorizeEditorDataProcessorsChainProcessorDTOWithDefaults instantiates a new AuthorizeEditorDataProcessorsChainProcessorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataProcessorsChainProcessorDTOWithDefaults() *AuthorizeEditorDataProcessorsChainProcessorDTO {
	this := AuthorizeEditorDataProcessorsChainProcessorDTO{}
	return &this
}

// GetProcessors returns the Processors field value
func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetProcessors() []AuthorizeEditorDataProcessorDTO {
	if o == nil {
		var ret []AuthorizeEditorDataProcessorDTO
		return ret
	}

	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) GetProcessorsOk() ([]AuthorizeEditorDataProcessorDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Processors, true
}

// SetProcessors sets field value
func (o *AuthorizeEditorDataProcessorsChainProcessorDTO) SetProcessors(v []AuthorizeEditorDataProcessorDTO) {
	o.Processors = v
}

func (o AuthorizeEditorDataProcessorsChainProcessorDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataProcessorsChainProcessorDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataProcessorDTO, errAuthorizeEditorDataProcessorDTO := json.Marshal(o.AuthorizeEditorDataProcessorDTO)
	if errAuthorizeEditorDataProcessorDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataProcessorDTO
	}
	errAuthorizeEditorDataProcessorDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataProcessorDTO), &toSerialize)
	if errAuthorizeEditorDataProcessorDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataProcessorDTO
	}
	toSerialize["processors"] = o.Processors
	return toSerialize, nil
}

type NullableAuthorizeEditorDataProcessorsChainProcessorDTO struct {
	value *AuthorizeEditorDataProcessorsChainProcessorDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataProcessorsChainProcessorDTO) Get() *AuthorizeEditorDataProcessorsChainProcessorDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataProcessorsChainProcessorDTO) Set(val *AuthorizeEditorDataProcessorsChainProcessorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataProcessorsChainProcessorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataProcessorsChainProcessorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataProcessorsChainProcessorDTO(val *AuthorizeEditorDataProcessorsChainProcessorDTO) *NullableAuthorizeEditorDataProcessorsChainProcessorDTO {
	return &NullableAuthorizeEditorDataProcessorsChainProcessorDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataProcessorsChainProcessorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataProcessorsChainProcessorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

