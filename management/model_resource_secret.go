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

// checks if the ResourceSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSecret{}

// ResourceSecret struct for ResourceSecret
type ResourceSecret struct {
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// An auto-generated resource client secret. Possible characters are a-z, A-Z, 0-9, -, ., _, ~. The secret has a minimum length of 64 characters per SHA-512 requirements when using the HS512 algorithm to sign ID tokens using the secret as the key.
	Secret   *string                 `json:"secret,omitempty"`
	Previous *ResourceSecretPrevious `json:"previous,omitempty"`
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
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSecret) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ResourceSecret) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
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
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSecret) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ResourceSecret) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ResourceSecret) SetSecret(v string) {
	o.Secret = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ResourceSecret) GetPrevious() ResourceSecretPrevious {
	if o == nil || IsNil(o.Previous) {
		var ret ResourceSecretPrevious
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceSecret) GetPreviousOk() (*ResourceSecretPrevious, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ResourceSecret) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given ResourceSecretPrevious and assigns it to the Previous field.
func (o *ResourceSecret) SetPrevious(v ResourceSecretPrevious) {
	o.Previous = &v
}

func (o ResourceSecret) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	return toSerialize, nil
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
