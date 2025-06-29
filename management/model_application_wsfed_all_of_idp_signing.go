/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the ApplicationWSFEDAllOfIdpSigning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationWSFEDAllOfIdpSigning{}

// ApplicationWSFEDAllOfIdpSigning Contains the information about the signing of requests by the identity provider (IdP).
type ApplicationWSFEDAllOfIdpSigning struct {
	Algorithm EnumApplicationWSFEDIDPSigningAlgorithm `json:"algorithm"`
	Key       ApplicationWSFEDAllOfIdpSigningKey      `json:"key"`
}

// NewApplicationWSFEDAllOfIdpSigning instantiates a new ApplicationWSFEDAllOfIdpSigning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationWSFEDAllOfIdpSigning(algorithm EnumApplicationWSFEDIDPSigningAlgorithm, key ApplicationWSFEDAllOfIdpSigningKey) *ApplicationWSFEDAllOfIdpSigning {
	this := ApplicationWSFEDAllOfIdpSigning{}
	this.Algorithm = algorithm
	this.Key = key
	return &this
}

// NewApplicationWSFEDAllOfIdpSigningWithDefaults instantiates a new ApplicationWSFEDAllOfIdpSigning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWSFEDAllOfIdpSigningWithDefaults() *ApplicationWSFEDAllOfIdpSigning {
	this := ApplicationWSFEDAllOfIdpSigning{}
	return &this
}

// GetAlgorithm returns the Algorithm field value
func (o *ApplicationWSFEDAllOfIdpSigning) GetAlgorithm() EnumApplicationWSFEDIDPSigningAlgorithm {
	if o == nil {
		var ret EnumApplicationWSFEDIDPSigningAlgorithm
		return ret
	}

	return o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value
// and a boolean to check if the value has been set.
func (o *ApplicationWSFEDAllOfIdpSigning) GetAlgorithmOk() (*EnumApplicationWSFEDIDPSigningAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Algorithm, true
}

// SetAlgorithm sets field value
func (o *ApplicationWSFEDAllOfIdpSigning) SetAlgorithm(v EnumApplicationWSFEDIDPSigningAlgorithm) {
	o.Algorithm = v
}

// GetKey returns the Key field value
func (o *ApplicationWSFEDAllOfIdpSigning) GetKey() ApplicationWSFEDAllOfIdpSigningKey {
	if o == nil {
		var ret ApplicationWSFEDAllOfIdpSigningKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApplicationWSFEDAllOfIdpSigning) GetKeyOk() (*ApplicationWSFEDAllOfIdpSigningKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApplicationWSFEDAllOfIdpSigning) SetKey(v ApplicationWSFEDAllOfIdpSigningKey) {
	o.Key = v
}

func (o ApplicationWSFEDAllOfIdpSigning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationWSFEDAllOfIdpSigning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["algorithm"] = o.Algorithm
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableApplicationWSFEDAllOfIdpSigning struct {
	value *ApplicationWSFEDAllOfIdpSigning
	isSet bool
}

func (v NullableApplicationWSFEDAllOfIdpSigning) Get() *ApplicationWSFEDAllOfIdpSigning {
	return v.value
}

func (v *NullableApplicationWSFEDAllOfIdpSigning) Set(val *ApplicationWSFEDAllOfIdpSigning) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationWSFEDAllOfIdpSigning) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationWSFEDAllOfIdpSigning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationWSFEDAllOfIdpSigning(val *ApplicationWSFEDAllOfIdpSigning) *NullableApplicationWSFEDAllOfIdpSigning {
	return &NullableApplicationWSFEDAllOfIdpSigning{value: val, isSet: true}
}

func (v NullableApplicationWSFEDAllOfIdpSigning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationWSFEDAllOfIdpSigning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
