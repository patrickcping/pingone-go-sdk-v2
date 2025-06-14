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

// checks if the LicenseAdvancedServicesPingId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseAdvancedServicesPingId{}

// LicenseAdvancedServicesPingId struct for LicenseAdvancedServicesPingId
type LicenseAdvancedServicesPingId struct {
	Included *bool   `json:"included,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// NewLicenseAdvancedServicesPingId instantiates a new LicenseAdvancedServicesPingId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseAdvancedServicesPingId() *LicenseAdvancedServicesPingId {
	this := LicenseAdvancedServicesPingId{}
	return &this
}

// NewLicenseAdvancedServicesPingIdWithDefaults instantiates a new LicenseAdvancedServicesPingId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseAdvancedServicesPingIdWithDefaults() *LicenseAdvancedServicesPingId {
	this := LicenseAdvancedServicesPingId{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *LicenseAdvancedServicesPingId) GetIncluded() bool {
	if o == nil || IsNil(o.Included) {
		var ret bool
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAdvancedServicesPingId) GetIncludedOk() (*bool, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *LicenseAdvancedServicesPingId) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given bool and assigns it to the Included field.
func (o *LicenseAdvancedServicesPingId) SetIncluded(v bool) {
	o.Included = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LicenseAdvancedServicesPingId) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseAdvancedServicesPingId) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LicenseAdvancedServicesPingId) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LicenseAdvancedServicesPingId) SetType(v string) {
	o.Type = &v
}

func (o LicenseAdvancedServicesPingId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseAdvancedServicesPingId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableLicenseAdvancedServicesPingId struct {
	value *LicenseAdvancedServicesPingId
	isSet bool
}

func (v NullableLicenseAdvancedServicesPingId) Get() *LicenseAdvancedServicesPingId {
	return v.value
}

func (v *NullableLicenseAdvancedServicesPingId) Set(val *LicenseAdvancedServicesPingId) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseAdvancedServicesPingId) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseAdvancedServicesPingId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseAdvancedServicesPingId(val *LicenseAdvancedServicesPingId) *NullableLicenseAdvancedServicesPingId {
	return &NullableLicenseAdvancedServicesPingId{value: val, isSet: true}
}

func (v NullableLicenseAdvancedServicesPingId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseAdvancedServicesPingId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
