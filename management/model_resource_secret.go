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

// ResourceSecret struct for ResourceSecret
type ResourceSecret struct {
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// An auto-generated resource client secret. Possible characters are a-z, A-Z, 0-9, -, ., _, ~. The secret has a minimum length of 64 characters per SHA-512 requirements when using the HS512 algorithm to sign ID tokens using the secret as the key.
	Secret *string `json:"secret,omitempty"`
}

// NewResourceSecret instantiates a new ResourceSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSecret() *ResourceSecret {
	this := ResourceSecret{}
	return &this
}

// NewResourceSecretWithDefaults instantiates a new ResourceSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSecretWithDefaults() *ResourceSecret {
	this := ResourceSecret{}
	return &this
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ResourceSecret) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSecret) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ResourceSecret) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ResourceSecret) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ResourceSecret) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSecret) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ResourceSecret) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ResourceSecret) SetSecret(v string) {
	o.Secret = &v
}

func (o ResourceSecret) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	return json.Marshal(toSerialize)
}

type NullableResourceSecret struct {
	value *ResourceSecret
	isSet bool
}

func (v NullableResourceSecret) Get() *ResourceSecret {
	return v.value
}

func (v *NullableResourceSecret) Set(val *ResourceSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSecret(val *ResourceSecret) *NullableResourceSecret {
	return &NullableResourceSecret{value: val, isSet: true}
}

func (v NullableResourceSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


