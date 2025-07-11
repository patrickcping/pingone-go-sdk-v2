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

// checks if the RiskPolicyConditionAggregatedScoresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicyConditionAggregatedScoresInner{}

// RiskPolicyConditionAggregatedScoresInner struct for RiskPolicyConditionAggregatedScoresInner
type RiskPolicyConditionAggregatedScoresInner struct {
	// Text that identifies a specific risk predictor in the environment. It uses the form `${details.xxxxxxx.level}`, where the string between `details` and `level` is the compact name of the relevant predictor.
	Value string `json:"value"`
	// The score you want to assign to the predictor when it is determined that there is a high risk for the predictor. Value should be between 0 and 100. If it is determined that there is medium risk, the predictor will automatically be assigned a score equal to half of the value you specified for high risk.
	Score int32 `json:"score"`
}

// NewRiskPolicyConditionAggregatedScoresInner instantiates a new RiskPolicyConditionAggregatedScoresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicyConditionAggregatedScoresInner(value string, score int32) *RiskPolicyConditionAggregatedScoresInner {
	this := RiskPolicyConditionAggregatedScoresInner{}
	this.Value = value
	this.Score = score
	return &this
}

// NewRiskPolicyConditionAggregatedScoresInnerWithDefaults instantiates a new RiskPolicyConditionAggregatedScoresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicyConditionAggregatedScoresInnerWithDefaults() *RiskPolicyConditionAggregatedScoresInner {
	this := RiskPolicyConditionAggregatedScoresInner{}
	return &this
}

// GetValue returns the Value field value
func (o *RiskPolicyConditionAggregatedScoresInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *RiskPolicyConditionAggregatedScoresInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *RiskPolicyConditionAggregatedScoresInner) SetValue(v string) {
	o.Value = v
}

// GetScore returns the Score field value
func (o *RiskPolicyConditionAggregatedScoresInner) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *RiskPolicyConditionAggregatedScoresInner) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *RiskPolicyConditionAggregatedScoresInner) SetScore(v int32) {
	o.Score = v
}

func (o RiskPolicyConditionAggregatedScoresInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicyConditionAggregatedScoresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["score"] = o.Score
	return toSerialize, nil
}

type NullableRiskPolicyConditionAggregatedScoresInner struct {
	value *RiskPolicyConditionAggregatedScoresInner
	isSet bool
}

func (v NullableRiskPolicyConditionAggregatedScoresInner) Get() *RiskPolicyConditionAggregatedScoresInner {
	return v.value
}

func (v *NullableRiskPolicyConditionAggregatedScoresInner) Set(val *RiskPolicyConditionAggregatedScoresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicyConditionAggregatedScoresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicyConditionAggregatedScoresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicyConditionAggregatedScoresInner(val *RiskPolicyConditionAggregatedScoresInner) *NullableRiskPolicyConditionAggregatedScoresInner {
	return &NullableRiskPolicyConditionAggregatedScoresInner{value: val, isSet: true}
}

func (v NullableRiskPolicyConditionAggregatedScoresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicyConditionAggregatedScoresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
