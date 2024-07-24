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

// checks if the AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO{}

// AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO struct for AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO
type AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO struct {
	AuthorizeEditorDataRulesEffectSettingsDTO
}

// NewAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO instantiates a new AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO(type_ string) *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO {
	this := AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO{}
	this.Type = type_
	return &this
}

// NewAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTOWithDefaults instantiates a new AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTOWithDefaults() *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO {
	this := AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO{}
	return &this
}

func (o AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataRulesEffectSettingsDTO, errAuthorizeEditorDataRulesEffectSettingsDTO := json.Marshal(o.AuthorizeEditorDataRulesEffectSettingsDTO)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	errAuthorizeEditorDataRulesEffectSettingsDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataRulesEffectSettingsDTO), &toSerialize)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO struct {
	value *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) Get() *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) Set(val *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO(val *AuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) *NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO {
	return &NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsUnconditionalDenyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

