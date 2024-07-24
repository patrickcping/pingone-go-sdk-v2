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

// checks if the AuthorizeEditorDataCacheSettingsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataCacheSettingsDTO{}

// AuthorizeEditorDataCacheSettingsDTO struct for AuthorizeEditorDataCacheSettingsDTO
type AuthorizeEditorDataCacheSettingsDTO struct {
	TtlSeconds *int32 `json:"ttlSeconds,omitempty"`
}

// NewAuthorizeEditorDataCacheSettingsDTO instantiates a new AuthorizeEditorDataCacheSettingsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataCacheSettingsDTO() *AuthorizeEditorDataCacheSettingsDTO {
	this := AuthorizeEditorDataCacheSettingsDTO{}
	return &this
}

// NewAuthorizeEditorDataCacheSettingsDTOWithDefaults instantiates a new AuthorizeEditorDataCacheSettingsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataCacheSettingsDTOWithDefaults() *AuthorizeEditorDataCacheSettingsDTO {
	this := AuthorizeEditorDataCacheSettingsDTO{}
	return &this
}

// GetTtlSeconds returns the TtlSeconds field value if set, zero value otherwise.
func (o *AuthorizeEditorDataCacheSettingsDTO) GetTtlSeconds() int32 {
	if o == nil || IsNil(o.TtlSeconds) {
		var ret int32
		return ret
	}
	return *o.TtlSeconds
}

// GetTtlSecondsOk returns a tuple with the TtlSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataCacheSettingsDTO) GetTtlSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.TtlSeconds) {
		return nil, false
	}
	return o.TtlSeconds, true
}

// HasTtlSeconds returns a boolean if a field has been set.
func (o *AuthorizeEditorDataCacheSettingsDTO) HasTtlSeconds() bool {
	if o != nil && !IsNil(o.TtlSeconds) {
		return true
	}

	return false
}

// SetTtlSeconds gets a reference to the given int32 and assigns it to the TtlSeconds field.
func (o *AuthorizeEditorDataCacheSettingsDTO) SetTtlSeconds(v int32) {
	o.TtlSeconds = &v
}

func (o AuthorizeEditorDataCacheSettingsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataCacheSettingsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TtlSeconds) {
		toSerialize["ttlSeconds"] = o.TtlSeconds
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataCacheSettingsDTO struct {
	value *AuthorizeEditorDataCacheSettingsDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataCacheSettingsDTO) Get() *AuthorizeEditorDataCacheSettingsDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataCacheSettingsDTO) Set(val *AuthorizeEditorDataCacheSettingsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataCacheSettingsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataCacheSettingsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataCacheSettingsDTO(val *AuthorizeEditorDataCacheSettingsDTO) *NullableAuthorizeEditorDataCacheSettingsDTO {
	return &NullableAuthorizeEditorDataCacheSettingsDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataCacheSettingsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataCacheSettingsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

