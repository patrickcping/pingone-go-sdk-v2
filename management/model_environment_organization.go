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

// EnvironmentOrganization struct for EnvironmentOrganization
type EnvironmentOrganization struct {
	// A string that specifies the organization resource’s unique identifier associated with the environment.
	Id *string `json:"id,omitempty"`
}

// NewEnvironmentOrganization instantiates a new EnvironmentOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentOrganization() *EnvironmentOrganization {
	this := EnvironmentOrganization{}
	return &this
}

// NewEnvironmentOrganizationWithDefaults instantiates a new EnvironmentOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentOrganizationWithDefaults() *EnvironmentOrganization {
	this := EnvironmentOrganization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentOrganization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentOrganization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentOrganization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentOrganization) SetId(v string) {
	o.Id = &v
}

func (o EnvironmentOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentOrganization struct {
	value *EnvironmentOrganization
	isSet bool
}

func (v NullableEnvironmentOrganization) Get() *EnvironmentOrganization {
	return v.value
}

func (v *NullableEnvironmentOrganization) Set(val *EnvironmentOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentOrganization(val *EnvironmentOrganization) *NullableEnvironmentOrganization {
	return &NullableEnvironmentOrganization{value: val, isSet: true}
}

func (v NullableEnvironmentOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


