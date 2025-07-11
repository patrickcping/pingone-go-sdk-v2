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

// checks if the LicenseReplacesLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseReplacesLicense{}

// LicenseReplacesLicense A read-only object that specifies the license ID of the license that is replaced by this license.
type LicenseReplacesLicense struct {
	// A read-only string that specifies the license ID of the license that is replaced by this license.
	Id *string `json:"id,omitempty"`
}

// NewLicenseReplacesLicense instantiates a new LicenseReplacesLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseReplacesLicense() *LicenseReplacesLicense {
	this := LicenseReplacesLicense{}
	return &this
}

// NewLicenseReplacesLicenseWithDefaults instantiates a new LicenseReplacesLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseReplacesLicenseWithDefaults() *LicenseReplacesLicense {
	this := LicenseReplacesLicense{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LicenseReplacesLicense) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseReplacesLicense) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LicenseReplacesLicense) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LicenseReplacesLicense) SetId(v string) {
	o.Id = &v
}

func (o LicenseReplacesLicense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseReplacesLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableLicenseReplacesLicense struct {
	value *LicenseReplacesLicense
	isSet bool
}

func (v NullableLicenseReplacesLicense) Get() *LicenseReplacesLicense {
	return v.value
}

func (v *NullableLicenseReplacesLicense) Set(val *LicenseReplacesLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseReplacesLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseReplacesLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseReplacesLicense(val *LicenseReplacesLicense) *NullableLicenseReplacesLicense {
	return &NullableLicenseReplacesLicense{value: val, isSet: true}
}

func (v NullableLicenseReplacesLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseReplacesLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
