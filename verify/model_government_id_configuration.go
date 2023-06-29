/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the GovernmentIdConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GovernmentIdConfiguration{}

// GovernmentIdConfiguration struct for GovernmentIdConfiguration
type GovernmentIdConfiguration struct {
	Verify EnumVerify `json:"verify"`
}

// NewGovernmentIdConfiguration instantiates a new GovernmentIdConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernmentIdConfiguration(verify EnumVerify) *GovernmentIdConfiguration {
	this := GovernmentIdConfiguration{}
	this.Verify = verify
	return &this
}

// NewGovernmentIdConfigurationWithDefaults instantiates a new GovernmentIdConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernmentIdConfigurationWithDefaults() *GovernmentIdConfiguration {
	this := GovernmentIdConfiguration{}
	return &this
}

// GetVerify returns the Verify field value
func (o *GovernmentIdConfiguration) GetVerify() EnumVerify {
	if o == nil {
		var ret EnumVerify
		return ret
	}

	return o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value
// and a boolean to check if the value has been set.
func (o *GovernmentIdConfiguration) GetVerifyOk() (*EnumVerify, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verify, true
}

// SetVerify sets field value
func (o *GovernmentIdConfiguration) SetVerify(v EnumVerify) {
	o.Verify = v
}

func (o GovernmentIdConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GovernmentIdConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verify"] = o.Verify
	return toSerialize, nil
}

type NullableGovernmentIdConfiguration struct {
	value *GovernmentIdConfiguration
	isSet bool
}

func (v NullableGovernmentIdConfiguration) Get() *GovernmentIdConfiguration {
	return v.value
}

func (v *NullableGovernmentIdConfiguration) Set(val *GovernmentIdConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernmentIdConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernmentIdConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernmentIdConfiguration(val *GovernmentIdConfiguration) *NullableGovernmentIdConfiguration {
	return &NullableGovernmentIdConfiguration{value: val, isSet: true}
}

func (v NullableGovernmentIdConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernmentIdConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


