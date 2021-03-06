/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// MFAPushCredentialAPNSAllOf struct for MFAPushCredentialAPNSAllOf
type MFAPushCredentialAPNSAllOf struct {
	// A string that Apple uses as an identifier to identify teams.
	TeamId string `json:"teamId"`
	// A string that Apple uses as the authentication token signing key to securely connect to APNS. This is a p8 file with a private key format.
	Token string `json:"token"`
}

// NewMFAPushCredentialAPNSAllOf instantiates a new MFAPushCredentialAPNSAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredentialAPNSAllOf(teamId string, token string) *MFAPushCredentialAPNSAllOf {
	this := MFAPushCredentialAPNSAllOf{}
	this.TeamId = teamId
	this.Token = token
	return &this
}

// NewMFAPushCredentialAPNSAllOfWithDefaults instantiates a new MFAPushCredentialAPNSAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialAPNSAllOfWithDefaults() *MFAPushCredentialAPNSAllOf {
	this := MFAPushCredentialAPNSAllOf{}
	return &this
}

// GetTeamId returns the TeamId field value
func (o *MFAPushCredentialAPNSAllOf) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialAPNSAllOf) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *MFAPushCredentialAPNSAllOf) SetTeamId(v string) {
	o.TeamId = v
}

// GetToken returns the Token field value
func (o *MFAPushCredentialAPNSAllOf) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredentialAPNSAllOf) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *MFAPushCredentialAPNSAllOf) SetToken(v string) {
	o.Token = v
}

func (o MFAPushCredentialAPNSAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["teamId"] = o.TeamId
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableMFAPushCredentialAPNSAllOf struct {
	value *MFAPushCredentialAPNSAllOf
	isSet bool
}

func (v NullableMFAPushCredentialAPNSAllOf) Get() *MFAPushCredentialAPNSAllOf {
	return v.value
}

func (v *NullableMFAPushCredentialAPNSAllOf) Set(val *MFAPushCredentialAPNSAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialAPNSAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialAPNSAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialAPNSAllOf(val *MFAPushCredentialAPNSAllOf) *NullableMFAPushCredentialAPNSAllOf {
	return &NullableMFAPushCredentialAPNSAllOf{value: val, isSet: true}
}

func (v NullableMFAPushCredentialAPNSAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialAPNSAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


