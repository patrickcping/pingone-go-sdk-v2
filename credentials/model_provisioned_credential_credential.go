/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ProvisionedCredentialCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionedCredentialCredential{}

// ProvisionedCredentialCredential struct for ProvisionedCredentialCredential
type ProvisionedCredentialCredential struct {
	Id *string `json:"id,omitempty"`
}

// NewProvisionedCredentialCredential instantiates a new ProvisionedCredentialCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionedCredentialCredential() *ProvisionedCredentialCredential {
	this := ProvisionedCredentialCredential{}
	return &this
}

// NewProvisionedCredentialCredentialWithDefaults instantiates a new ProvisionedCredentialCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionedCredentialCredentialWithDefaults() *ProvisionedCredentialCredential {
	this := ProvisionedCredentialCredential{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProvisionedCredentialCredential) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredentialCredential) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProvisionedCredentialCredential) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProvisionedCredentialCredential) SetId(v string) {
	o.Id = &v
}

func (o ProvisionedCredentialCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionedCredentialCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableProvisionedCredentialCredential struct {
	value *ProvisionedCredentialCredential
	isSet bool
}

func (v NullableProvisionedCredentialCredential) Get() *ProvisionedCredentialCredential {
	return v.value
}

func (v *NullableProvisionedCredentialCredential) Set(val *ProvisionedCredentialCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionedCredentialCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionedCredentialCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionedCredentialCredential(val *ProvisionedCredentialCredential) *NullableProvisionedCredentialCredential {
	return &NullableProvisionedCredentialCredential{value: val, isSet: true}
}

func (v NullableProvisionedCredentialCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionedCredentialCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


