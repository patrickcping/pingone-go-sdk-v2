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

// checks if the AuthorizeEditorDataConditionsOrConditionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataConditionsOrConditionDTO{}

// AuthorizeEditorDataConditionsOrConditionDTO struct for AuthorizeEditorDataConditionsOrConditionDTO
type AuthorizeEditorDataConditionsOrConditionDTO struct {
	Type EnumAuthorizeEditorDataConditionDTOType `json:"type"`
	Conditions []AuthorizeEditorDataConditionDTO `json:"conditions"`
}

// NewAuthorizeEditorDataConditionsOrConditionDTO instantiates a new AuthorizeEditorDataConditionsOrConditionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataConditionsOrConditionDTO(type_ EnumAuthorizeEditorDataConditionDTOType, conditions []AuthorizeEditorDataConditionDTO) *AuthorizeEditorDataConditionsOrConditionDTO {
	this := AuthorizeEditorDataConditionsOrConditionDTO{}
	this.Type = type_
	this.Conditions = conditions
	return &this
}

// NewAuthorizeEditorDataConditionsOrConditionDTOWithDefaults instantiates a new AuthorizeEditorDataConditionsOrConditionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataConditionsOrConditionDTOWithDefaults() *AuthorizeEditorDataConditionsOrConditionDTO {
	this := AuthorizeEditorDataConditionsOrConditionDTO{}
	return &this
}

// GetType returns the Type field value
func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetType() EnumAuthorizeEditorDataConditionDTOType {
	if o == nil {
		var ret EnumAuthorizeEditorDataConditionDTOType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetTypeOk() (*EnumAuthorizeEditorDataConditionDTOType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthorizeEditorDataConditionsOrConditionDTO) SetType(v EnumAuthorizeEditorDataConditionDTOType) {
	o.Type = v
}

// GetConditions returns the Conditions field value
func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetConditions() []AuthorizeEditorDataConditionDTO {
	if o == nil {
		var ret []AuthorizeEditorDataConditionDTO
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConditionsOrConditionDTO) GetConditionsOk() ([]AuthorizeEditorDataConditionDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *AuthorizeEditorDataConditionsOrConditionDTO) SetConditions(v []AuthorizeEditorDataConditionDTO) {
	o.Conditions = v
}

func (o AuthorizeEditorDataConditionsOrConditionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataConditionsOrConditionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["conditions"] = o.Conditions
	return toSerialize, nil
}

type NullableAuthorizeEditorDataConditionsOrConditionDTO struct {
	value *AuthorizeEditorDataConditionsOrConditionDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataConditionsOrConditionDTO) Get() *AuthorizeEditorDataConditionsOrConditionDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataConditionsOrConditionDTO) Set(val *AuthorizeEditorDataConditionsOrConditionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataConditionsOrConditionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataConditionsOrConditionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataConditionsOrConditionDTO(val *AuthorizeEditorDataConditionsOrConditionDTO) *NullableAuthorizeEditorDataConditionsOrConditionDTO {
	return &NullableAuthorizeEditorDataConditionsOrConditionDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataConditionsOrConditionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataConditionsOrConditionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


