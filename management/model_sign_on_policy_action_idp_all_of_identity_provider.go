/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SignOnPolicyActionIDPAllOfIdentityProvider A reference to the external identity provider that is used to authenticate the user. This property is required.
type SignOnPolicyActionIDPAllOfIdentityProvider struct {
	// A string that specifies the ID of the external identity provider to which the user is redirected for sign-on. This property is required.
	Id string `json:"id"`
}

// NewSignOnPolicyActionIDPAllOfIdentityProvider instantiates a new SignOnPolicyActionIDPAllOfIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionIDPAllOfIdentityProvider(id string) *SignOnPolicyActionIDPAllOfIdentityProvider {
	this := SignOnPolicyActionIDPAllOfIdentityProvider{}
	this.Id = id
	return &this
}

// NewSignOnPolicyActionIDPAllOfIdentityProviderWithDefaults instantiates a new SignOnPolicyActionIDPAllOfIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionIDPAllOfIdentityProviderWithDefaults() *SignOnPolicyActionIDPAllOfIdentityProvider {
	this := SignOnPolicyActionIDPAllOfIdentityProvider{}
	return &this
}

// GetId returns the Id field value
func (o *SignOnPolicyActionIDPAllOfIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionIDPAllOfIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SignOnPolicyActionIDPAllOfIdentityProvider) SetId(v string) {
	o.Id = v
}

func (o SignOnPolicyActionIDPAllOfIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSignOnPolicyActionIDPAllOfIdentityProvider struct {
	value *SignOnPolicyActionIDPAllOfIdentityProvider
	isSet bool
}

func (v NullableSignOnPolicyActionIDPAllOfIdentityProvider) Get() *SignOnPolicyActionIDPAllOfIdentityProvider {
	return v.value
}

func (v *NullableSignOnPolicyActionIDPAllOfIdentityProvider) Set(val *SignOnPolicyActionIDPAllOfIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionIDPAllOfIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionIDPAllOfIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionIDPAllOfIdentityProvider(val *SignOnPolicyActionIDPAllOfIdentityProvider) *NullableSignOnPolicyActionIDPAllOfIdentityProvider {
	return &NullableSignOnPolicyActionIDPAllOfIdentityProvider{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionIDPAllOfIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionIDPAllOfIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


