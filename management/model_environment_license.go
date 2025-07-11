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

// checks if the EnvironmentLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentLicense{}

// EnvironmentLicense struct for EnvironmentLicense
type EnvironmentLicense struct {
	// A string that specifies the active license associated with this environment. This property is required only if your organization has more than one active license.
	Id string `json:"id"`
}

// NewEnvironmentLicense instantiates a new EnvironmentLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLicense(id string) *EnvironmentLicense {
	this := EnvironmentLicense{}
	this.Id = id
	return &this
}

// NewEnvironmentLicenseWithDefaults instantiates a new EnvironmentLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLicenseWithDefaults() *EnvironmentLicense {
	this := EnvironmentLicense{}
	return &this
}

// GetId returns the Id field value
func (o *EnvironmentLicense) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvironmentLicense) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvironmentLicense) SetId(v string) {
	o.Id = v
}

func (o EnvironmentLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableEnvironmentLicense struct {
	value *EnvironmentLicense
	isSet bool
}

func (v NullableEnvironmentLicense) Get() *EnvironmentLicense {
	return v.value
}

func (v *NullableEnvironmentLicense) Set(val *EnvironmentLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLicense(val *EnvironmentLicense) *NullableEnvironmentLicense {
	return &NullableEnvironmentLicense{value: val, isSet: true}
}

func (v NullableEnvironmentLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
