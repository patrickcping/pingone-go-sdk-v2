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

// checks if the ApplicationTemplateIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationTemplateIntegration{}

// ApplicationTemplateIntegration struct for ApplicationTemplateIntegration
type ApplicationTemplateIntegration struct {
	// The UUID of the integration in Integration Catalog.
	Id string `json:"id"`
}

// NewApplicationTemplateIntegration instantiates a new ApplicationTemplateIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationTemplateIntegration(id string) *ApplicationTemplateIntegration {
	this := ApplicationTemplateIntegration{}
	this.Id = id
	return &this
}

// NewApplicationTemplateIntegrationWithDefaults instantiates a new ApplicationTemplateIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationTemplateIntegrationWithDefaults() *ApplicationTemplateIntegration {
	this := ApplicationTemplateIntegration{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationTemplateIntegration) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationTemplateIntegration) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationTemplateIntegration) SetId(v string) {
	o.Id = v
}

func (o ApplicationTemplateIntegration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationTemplateIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableApplicationTemplateIntegration struct {
	value *ApplicationTemplateIntegration
	isSet bool
}

func (v NullableApplicationTemplateIntegration) Get() *ApplicationTemplateIntegration {
	return v.value
}

func (v *NullableApplicationTemplateIntegration) Set(val *ApplicationTemplateIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationTemplateIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationTemplateIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationTemplateIntegration(val *ApplicationTemplateIntegration) *NullableApplicationTemplateIntegration {
	return &NullableApplicationTemplateIntegration{value: val, isSet: true}
}

func (v NullableApplicationTemplateIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationTemplateIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
