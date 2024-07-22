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

// checks if the IntegrationVersionCommonIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionCommonIntegration{}

// IntegrationVersionCommonIntegration The parent integration of the integration version.
type IntegrationVersionCommonIntegration struct {
	// The platform-generated ID of the parent integration.
	Id string `json:"id"`
}

// NewIntegrationVersionCommonIntegration instantiates a new IntegrationVersionCommonIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionCommonIntegration(id string) *IntegrationVersionCommonIntegration {
	this := IntegrationVersionCommonIntegration{}
	this.Id = id
	return &this
}

// NewIntegrationVersionCommonIntegrationWithDefaults instantiates a new IntegrationVersionCommonIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionCommonIntegrationWithDefaults() *IntegrationVersionCommonIntegration {
	this := IntegrationVersionCommonIntegration{}
	return &this
}

// GetId returns the Id field value
func (o *IntegrationVersionCommonIntegration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommonIntegration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IntegrationVersionCommonIntegration) SetId(v string) {
	o.Id = v
}

func (o IntegrationVersionCommonIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionCommonIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableIntegrationVersionCommonIntegration struct {
	value *IntegrationVersionCommonIntegration
	isSet bool
}

func (v NullableIntegrationVersionCommonIntegration) Get() *IntegrationVersionCommonIntegration {
	return v.value
}

func (v *NullableIntegrationVersionCommonIntegration) Set(val *IntegrationVersionCommonIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionCommonIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionCommonIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionCommonIntegration(val *IntegrationVersionCommonIntegration) *NullableIntegrationVersionCommonIntegration {
	return &NullableIntegrationVersionCommonIntegration{value: val, isSet: true}
}

func (v NullableIntegrationVersionCommonIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionCommonIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

