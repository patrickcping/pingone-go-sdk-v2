/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2021-10-17
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// checks if the RiskPolicyConditionAggregatedWeightsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicyConditionAggregatedWeightsInner{}

// RiskPolicyConditionAggregatedWeightsInner struct for RiskPolicyConditionAggregatedWeightsInner
type RiskPolicyConditionAggregatedWeightsInner struct {
	// Text that identifies a specific risk predictor in the environment. It uses the form `${details.xxxxxxx.level}`, where the string between `details` and `level` is the compact name of the relevant predictor.
	Value string `json:"value"`
	// The score you want to assign to the predictor when it is determined that there is a high risk for the predictor. Value should be between 0 and 100. If it is determined that there is medium risk, the predictor will automatically be assigned a score equal to half of the value you specified for high risk.
	Weight int32 `json:"weight"`
}

// NewRiskPolicyConditionAggregatedWeightsInner instantiates a new RiskPolicyConditionAggregatedWeightsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicyConditionAggregatedWeightsInner(value string, weight int32) *RiskPolicyConditionAggregatedWeightsInner {
	this := RiskPolicyConditionAggregatedWeightsInner{}
	this.Value = value
	this.Weight = weight
	return &this
}

// NewRiskPolicyConditionAggregatedWeightsInnerWithDefaults instantiates a new RiskPolicyConditionAggregatedWeightsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicyConditionAggregatedWeightsInnerWithDefaults() *RiskPolicyConditionAggregatedWeightsInner {
	this := RiskPolicyConditionAggregatedWeightsInner{}
	return &this
}

// GetValue returns the Value field value
func (o *RiskPolicyConditionAggregatedWeightsInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RiskPolicyConditionAggregatedWeightsInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RiskPolicyConditionAggregatedWeightsInner) SetValue(v string) {
	o.Value = v
}

// GetWeight returns the Weight field value
func (o *RiskPolicyConditionAggregatedWeightsInner) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *RiskPolicyConditionAggregatedWeightsInner) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *RiskPolicyConditionAggregatedWeightsInner) SetWeight(v int32) {
	o.Weight = v
}

func (o RiskPolicyConditionAggregatedWeightsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicyConditionAggregatedWeightsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableRiskPolicyConditionAggregatedWeightsInner struct {
	value *RiskPolicyConditionAggregatedWeightsInner
	isSet bool
}

func (v NullableRiskPolicyConditionAggregatedWeightsInner) Get() *RiskPolicyConditionAggregatedWeightsInner {
	return v.value
}

func (v *NullableRiskPolicyConditionAggregatedWeightsInner) Set(val *RiskPolicyConditionAggregatedWeightsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicyConditionAggregatedWeightsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicyConditionAggregatedWeightsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicyConditionAggregatedWeightsInner(val *RiskPolicyConditionAggregatedWeightsInner) *NullableRiskPolicyConditionAggregatedWeightsInner {
	return &NullableRiskPolicyConditionAggregatedWeightsInner{value: val, isSet: true}
}

func (v NullableRiskPolicyConditionAggregatedWeightsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicyConditionAggregatedWeightsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

