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

// checks if the SubscriptionTlsClientAuthKeyPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionTlsClientAuthKeyPair{}

// SubscriptionTlsClientAuthKeyPair An onject that describes the TLS client authentication key pair to use for mTLS.
type SubscriptionTlsClientAuthKeyPair struct {
	// The ID of a key to be used for outbound mutual TLS (mTLS) authentication. This key is used as a client credential to authenticate the webhook. The key must have a `usageType` of `OUTBOUND_MTLS`. See [Certificate management](https://apidocs.pingidentity.com/pingone/platform/v1/api/#certificate-management) for information on creating a key. If this property is set, `verifyTlsCertificates` must be set to `true`.
	Id *string `json:"id,omitempty"`
}

// NewSubscriptionTlsClientAuthKeyPair instantiates a new SubscriptionTlsClientAuthKeyPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionTlsClientAuthKeyPair() *SubscriptionTlsClientAuthKeyPair {
	this := SubscriptionTlsClientAuthKeyPair{}
	return &this
}

// NewSubscriptionTlsClientAuthKeyPairWithDefaults instantiates a new SubscriptionTlsClientAuthKeyPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTlsClientAuthKeyPairWithDefaults() *SubscriptionTlsClientAuthKeyPair {
	this := SubscriptionTlsClientAuthKeyPair{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionTlsClientAuthKeyPair) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTlsClientAuthKeyPair) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionTlsClientAuthKeyPair) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubscriptionTlsClientAuthKeyPair) SetId(v string) {
	o.Id = &v
}

func (o SubscriptionTlsClientAuthKeyPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionTlsClientAuthKeyPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableSubscriptionTlsClientAuthKeyPair struct {
	value *SubscriptionTlsClientAuthKeyPair
	isSet bool
}

func (v NullableSubscriptionTlsClientAuthKeyPair) Get() *SubscriptionTlsClientAuthKeyPair {
	return v.value
}

func (v *NullableSubscriptionTlsClientAuthKeyPair) Set(val *SubscriptionTlsClientAuthKeyPair) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTlsClientAuthKeyPair) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTlsClientAuthKeyPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTlsClientAuthKeyPair(val *SubscriptionTlsClientAuthKeyPair) *NullableSubscriptionTlsClientAuthKeyPair {
	return &NullableSubscriptionTlsClientAuthKeyPair{value: val, isSet: true}
}

func (v NullableSubscriptionTlsClientAuthKeyPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionTlsClientAuthKeyPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
