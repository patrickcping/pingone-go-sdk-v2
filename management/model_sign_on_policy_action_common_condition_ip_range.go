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

// checks if the SignOnPolicyActionCommonConditionIPRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionCommonConditionIPRange{}

// SignOnPolicyActionCommonConditionIPRange struct for SignOnPolicyActionCommonConditionIPRange
type SignOnPolicyActionCommonConditionIPRange struct {
	Contains string   `json:"contains"`
	IpRange  []string `json:"ipRange"`
}

// NewSignOnPolicyActionCommonConditionIPRange instantiates a new SignOnPolicyActionCommonConditionIPRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionCommonConditionIPRange(contains string, ipRange []string) *SignOnPolicyActionCommonConditionIPRange {
	this := SignOnPolicyActionCommonConditionIPRange{}
	this.Contains = contains
	this.IpRange = ipRange
	return &this
}

// NewSignOnPolicyActionCommonConditionIPRangeWithDefaults instantiates a new SignOnPolicyActionCommonConditionIPRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionCommonConditionIPRangeWithDefaults() *SignOnPolicyActionCommonConditionIPRange {
	this := SignOnPolicyActionCommonConditionIPRange{}
	return &this
}

// GetContains returns the Contains field value
func (o *SignOnPolicyActionCommonConditionIPRange) GetContains() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contains
}

// GetContainsOk returns a tuple with the Contains field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionIPRange) GetContainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contains, true
}

// SetContains sets field value
func (o *SignOnPolicyActionCommonConditionIPRange) SetContains(v string) {
	o.Contains = v
}

// GetIpRange returns the IpRange field value
func (o *SignOnPolicyActionCommonConditionIPRange) GetIpRange() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionCommonConditionIPRange) GetIpRangeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpRange, true
}

// SetIpRange sets field value
func (o *SignOnPolicyActionCommonConditionIPRange) SetIpRange(v []string) {
	o.IpRange = v
}

func (o SignOnPolicyActionCommonConditionIPRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionCommonConditionIPRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contains"] = o.Contains
	toSerialize["ipRange"] = o.IpRange
	return toSerialize, nil
}

type NullableSignOnPolicyActionCommonConditionIPRange struct {
	value *SignOnPolicyActionCommonConditionIPRange
	isSet bool
}

func (v NullableSignOnPolicyActionCommonConditionIPRange) Get() *SignOnPolicyActionCommonConditionIPRange {
	return v.value
}

func (v *NullableSignOnPolicyActionCommonConditionIPRange) Set(val *SignOnPolicyActionCommonConditionIPRange) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionCommonConditionIPRange) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionCommonConditionIPRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionCommonConditionIPRange(val *SignOnPolicyActionCommonConditionIPRange) *NullableSignOnPolicyActionCommonConditionIPRange {
	return &NullableSignOnPolicyActionCommonConditionIPRange{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionCommonConditionIPRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionCommonConditionIPRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
