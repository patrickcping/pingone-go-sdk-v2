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

// checks if the AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO{}

// AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO struct for AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO
type AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO struct {
	AuthorizeEditorDataRulesEffectSettingsDTO
	Condition AuthorizeEditorDataConditionDTO `json:"condition"`
}

// NewAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO instantiates a new AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO(condition AuthorizeEditorDataConditionDTO, type_ string) *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO {
	this := AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO{}
	this.Type = type_
	this.Condition = condition
	return &this
}

// NewAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTOWithDefaults instantiates a new AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTOWithDefaults() *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO {
	this := AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO{}
	return &this
}

// GetCondition returns the Condition field value
func (o *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) GetCondition() AuthorizeEditorDataConditionDTO {
	if o == nil {
		var ret AuthorizeEditorDataConditionDTO
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) SetCondition(v AuthorizeEditorDataConditionDTO) {
	o.Condition = v
}

func (o AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataRulesEffectSettingsDTO, errAuthorizeEditorDataRulesEffectSettingsDTO := json.Marshal(o.AuthorizeEditorDataRulesEffectSettingsDTO)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	errAuthorizeEditorDataRulesEffectSettingsDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataRulesEffectSettingsDTO), &toSerialize)
	if errAuthorizeEditorDataRulesEffectSettingsDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataRulesEffectSettingsDTO
	}
	toSerialize["condition"] = o.Condition
	return toSerialize, nil
}

type NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO struct {
	value *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) Get() *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) Set(val *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO(val *AuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) *NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO {
	return &NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRuleseffectSettingsConditionalPermitElseDenyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

