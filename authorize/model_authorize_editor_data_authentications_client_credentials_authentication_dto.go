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

// checks if the AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO{}

// AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO struct for AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO
type AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO struct {
	AuthorizeEditorDataAuthenticationDTO
	TokenEndpoint string `json:"tokenEndpoint"`
	ClientId string `json:"clientId"`
	ClientSecret AuthorizeEditorDataReferenceObjectDTO `json:"clientSecret"`
	Scope string `json:"scope"`
}

// NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO instantiates a new AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO(tokenEndpoint string, clientId string, clientSecret AuthorizeEditorDataReferenceObjectDTO, scope string, type_ string) *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO{}
	this.Type = type_
	this.TokenEndpoint = tokenEndpoint
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.Scope = scope
	return &this
}

// NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTOWithDefaults instantiates a new AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTOWithDefaults() *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO {
	this := AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO{}
	return &this
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetClientId returns the ClientId field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientSecret() AuthorizeEditorDataReferenceObjectDTO {
	if o == nil {
		var ret AuthorizeEditorDataReferenceObjectDTO
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetClientSecretOk() (*AuthorizeEditorDataReferenceObjectDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetClientSecret(v AuthorizeEditorDataReferenceObjectDTO) {
	o.ClientSecret = v
}

// GetScope returns the Scope field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) SetScope(v string) {
	o.Scope = v
}

func (o AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedAuthorizeEditorDataAuthenticationDTO, errAuthorizeEditorDataAuthenticationDTO := json.Marshal(o.AuthorizeEditorDataAuthenticationDTO)
	if errAuthorizeEditorDataAuthenticationDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataAuthenticationDTO
	}
	errAuthorizeEditorDataAuthenticationDTO = json.Unmarshal([]byte(serializedAuthorizeEditorDataAuthenticationDTO), &toSerialize)
	if errAuthorizeEditorDataAuthenticationDTO != nil {
		return map[string]interface{}{}, errAuthorizeEditorDataAuthenticationDTO
	}
	toSerialize["tokenEndpoint"] = o.TokenEndpoint
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["scope"] = o.Scope
	return toSerialize, nil
}

type NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO struct {
	value *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) Get() *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) Set(val *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO(val *AuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) *NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO {
	return &NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAuthenticationsClientCredentialsAuthenticationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

