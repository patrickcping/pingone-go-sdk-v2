/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// checks if the RiskPredictorCustomItemIPRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCustomItemIPRange{}

// RiskPredictorCustomItemIPRange The mapping of risk levels for the IP ranges specified.
type RiskPredictorCustomItemIPRange struct {
	Contains string `json:"contains"`
	Type *string `json:"type,omitempty"`
	// List of CIDRs to include
	IpRange []string `json:"ipRange"`
}

// NewRiskPredictorCustomItemIPRange instantiates a new RiskPredictorCustomItemIPRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCustomItemIPRange(contains string, ipRange []string) *RiskPredictorCustomItemIPRange {
	this := RiskPredictorCustomItemIPRange{}
	this.Contains = contains
	this.IpRange = ipRange
	return &this
}

// NewRiskPredictorCustomItemIPRangeWithDefaults instantiates a new RiskPredictorCustomItemIPRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCustomItemIPRangeWithDefaults() *RiskPredictorCustomItemIPRange {
	this := RiskPredictorCustomItemIPRange{}
	return &this
}

// GetContains returns the Contains field value
func (o *RiskPredictorCustomItemIPRange) GetContains() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contains
}

// GetContainsOk returns a tuple with the Contains field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustomItemIPRange) GetContainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contains, true
}

// SetContains sets field value
func (o *RiskPredictorCustomItemIPRange) SetContains(v string) {
	o.Contains = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPredictorCustomItemIPRange) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustomItemIPRange) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPredictorCustomItemIPRange) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RiskPredictorCustomItemIPRange) SetType(v string) {
	o.Type = &v
}

// GetIpRange returns the IpRange field value
func (o *RiskPredictorCustomItemIPRange) GetIpRange() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustomItemIPRange) GetIpRangeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpRange, true
}

// SetIpRange sets field value
func (o *RiskPredictorCustomItemIPRange) SetIpRange(v []string) {
	o.IpRange = v
}

func (o RiskPredictorCustomItemIPRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCustomItemIPRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contains"] = o.Contains
	// skip: type is readOnly
	toSerialize["ipRange"] = o.IpRange
	return toSerialize, nil
}

type NullableRiskPredictorCustomItemIPRange struct {
	value *RiskPredictorCustomItemIPRange
	isSet bool
}

func (v NullableRiskPredictorCustomItemIPRange) Get() *RiskPredictorCustomItemIPRange {
	return v.value
}

func (v *NullableRiskPredictorCustomItemIPRange) Set(val *RiskPredictorCustomItemIPRange) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCustomItemIPRange) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCustomItemIPRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCustomItemIPRange(val *RiskPredictorCustomItemIPRange) *NullableRiskPredictorCustomItemIPRange {
	return &NullableRiskPredictorCustomItemIPRange{value: val, isSet: true}
}

func (v NullableRiskPredictorCustomItemIPRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCustomItemIPRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


