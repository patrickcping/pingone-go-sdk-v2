/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// IdentityProviderLinkedInAllOf struct for IdentityProviderLinkedInAllOf
type IdentityProviderLinkedInAllOf struct {
	// A string that specifies the application ID from LinkedIn. This is a required property.
	ClientId string `json:"clientId"`
	// A string that specifies the application secret from LinkedIn. This is a required property.
	ClientSecret string `json:"clientSecret"`
}

// NewIdentityProviderLinkedInAllOf instantiates a new IdentityProviderLinkedInAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderLinkedInAllOf(clientId string, clientSecret string) *IdentityProviderLinkedInAllOf {
	this := IdentityProviderLinkedInAllOf{}
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	return &this
}

// NewIdentityProviderLinkedInAllOfWithDefaults instantiates a new IdentityProviderLinkedInAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderLinkedInAllOfWithDefaults() *IdentityProviderLinkedInAllOf {
	this := IdentityProviderLinkedInAllOf{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *IdentityProviderLinkedInAllOf) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinkedInAllOf) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *IdentityProviderLinkedInAllOf) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *IdentityProviderLinkedInAllOf) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinkedInAllOf) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *IdentityProviderLinkedInAllOf) SetClientSecret(v string) {
	o.ClientSecret = v
}

func (o IdentityProviderLinkedInAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if true {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityProviderLinkedInAllOf struct {
	value *IdentityProviderLinkedInAllOf
	isSet bool
}

func (v NullableIdentityProviderLinkedInAllOf) Get() *IdentityProviderLinkedInAllOf {
	return v.value
}

func (v *NullableIdentityProviderLinkedInAllOf) Set(val *IdentityProviderLinkedInAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderLinkedInAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderLinkedInAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderLinkedInAllOf(val *IdentityProviderLinkedInAllOf) *NullableIdentityProviderLinkedInAllOf {
	return &NullableIdentityProviderLinkedInAllOf{value: val, isSet: true}
}

func (v NullableIdentityProviderLinkedInAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderLinkedInAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

