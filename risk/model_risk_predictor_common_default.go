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

// checks if the RiskPredictorCommonDefault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCommonDefault{}

// RiskPredictorCommonDefault Contains the default values used for a new risk predictor.
type RiskPredictorCommonDefault struct {
	// An integer type. This specifies the weight assigned to the risk predictor in a new policy by default.
	Weight    int32                             `json:"weight"`
	Score     *int32                            `json:"score,omitempty"`
	Evaluated *bool                             `json:"evaluated,omitempty"`
	Result    *RiskPredictorCommonDefaultResult `json:"result,omitempty"`
}

// NewRiskPredictorCommonDefault instantiates a new RiskPredictorCommonDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCommonDefault(weight int32) *RiskPredictorCommonDefault {
	this := RiskPredictorCommonDefault{}
	this.Weight = weight
	return &this
}

// NewRiskPredictorCommonDefaultWithDefaults instantiates a new RiskPredictorCommonDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCommonDefaultWithDefaults() *RiskPredictorCommonDefault {
	this := RiskPredictorCommonDefault{}
	return &this
}

// GetWeight returns the Weight field value
func (o *RiskPredictorCommonDefault) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonDefault) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *RiskPredictorCommonDefault) SetWeight(v int32) {
	o.Weight = v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *RiskPredictorCommonDefault) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonDefault) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *RiskPredictorCommonDefault) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *RiskPredictorCommonDefault) SetScore(v int32) {
	o.Score = &v
}

// GetEvaluated returns the Evaluated field value if set, zero value otherwise.
func (o *RiskPredictorCommonDefault) GetEvaluated() bool {
	if o == nil || IsNil(o.Evaluated) {
		var ret bool
		return ret
	}
	return *o.Evaluated
}

// GetEvaluatedOk returns a tuple with the Evaluated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonDefault) GetEvaluatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Evaluated) {
		return nil, false
	}
	return o.Evaluated, true
}

// HasEvaluated returns a boolean if a field has been set.
func (o *RiskPredictorCommonDefault) HasEvaluated() bool {
	if o != nil && !IsNil(o.Evaluated) {
		return true
	}

	return false
}

// SetEvaluated gets a reference to the given bool and assigns it to the Evaluated field.
func (o *RiskPredictorCommonDefault) SetEvaluated(v bool) {
	o.Evaluated = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RiskPredictorCommonDefault) GetResult() RiskPredictorCommonDefaultResult {
	if o == nil || IsNil(o.Result) {
		var ret RiskPredictorCommonDefaultResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonDefault) GetResultOk() (*RiskPredictorCommonDefaultResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RiskPredictorCommonDefault) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given RiskPredictorCommonDefaultResult and assigns it to the Result field.
func (o *RiskPredictorCommonDefault) SetResult(v RiskPredictorCommonDefaultResult) {
	o.Result = &v
}

func (o RiskPredictorCommonDefault) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCommonDefault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["weight"] = o.Weight
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Evaluated) {
		toSerialize["evaluated"] = o.Evaluated
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableRiskPredictorCommonDefault struct {
	value *RiskPredictorCommonDefault
	isSet bool
}

func (v NullableRiskPredictorCommonDefault) Get() *RiskPredictorCommonDefault {
	return v.value
}

func (v *NullableRiskPredictorCommonDefault) Set(val *RiskPredictorCommonDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCommonDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCommonDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCommonDefault(val *RiskPredictorCommonDefault) *NullableRiskPredictorCommonDefault {
	return &NullableRiskPredictorCommonDefault{value: val, isSet: true}
}

func (v NullableRiskPredictorCommonDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCommonDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
