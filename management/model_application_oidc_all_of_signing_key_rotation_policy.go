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

// checks if the ApplicationOIDCAllOfSigningKeyRotationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOIDCAllOfSigningKeyRotationPolicy{}

// ApplicationOIDCAllOfSigningKeyRotationPolicy Contains the Key Rotation Policy (KRP) ID. This property is required if `signing` is set.
type ApplicationOIDCAllOfSigningKeyRotationPolicy struct {
	// Reference to a KRP ID from certificate management. This property is required if `signing` is set.
	Id string `json:"id"`
}

// NewApplicationOIDCAllOfSigningKeyRotationPolicy instantiates a new ApplicationOIDCAllOfSigningKeyRotationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOfSigningKeyRotationPolicy(id string) *ApplicationOIDCAllOfSigningKeyRotationPolicy {
	this := ApplicationOIDCAllOfSigningKeyRotationPolicy{}
	this.Id = id
	return &this
}

// NewApplicationOIDCAllOfSigningKeyRotationPolicyWithDefaults instantiates a new ApplicationOIDCAllOfSigningKeyRotationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfSigningKeyRotationPolicyWithDefaults() *ApplicationOIDCAllOfSigningKeyRotationPolicy {
	this := ApplicationOIDCAllOfSigningKeyRotationPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationOIDCAllOfSigningKeyRotationPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfSigningKeyRotationPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationOIDCAllOfSigningKeyRotationPolicy) SetId(v string) {
	o.Id = v
}

func (o ApplicationOIDCAllOfSigningKeyRotationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOIDCAllOfSigningKeyRotationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableApplicationOIDCAllOfSigningKeyRotationPolicy struct {
	value *ApplicationOIDCAllOfSigningKeyRotationPolicy
	isSet bool
}

func (v NullableApplicationOIDCAllOfSigningKeyRotationPolicy) Get() *ApplicationOIDCAllOfSigningKeyRotationPolicy {
	return v.value
}

func (v *NullableApplicationOIDCAllOfSigningKeyRotationPolicy) Set(val *ApplicationOIDCAllOfSigningKeyRotationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOfSigningKeyRotationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOfSigningKeyRotationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOfSigningKeyRotationPolicy(val *ApplicationOIDCAllOfSigningKeyRotationPolicy) *NullableApplicationOIDCAllOfSigningKeyRotationPolicy {
	return &NullableApplicationOIDCAllOfSigningKeyRotationPolicy{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOfSigningKeyRotationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOfSigningKeyRotationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


