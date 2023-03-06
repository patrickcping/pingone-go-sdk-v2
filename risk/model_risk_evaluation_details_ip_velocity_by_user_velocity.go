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

// checks if the RiskEvaluationDetailsIpVelocityByUserVelocity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationDetailsIpVelocityByUserVelocity{}

// RiskEvaluationDetailsIpVelocityByUserVelocity Integer values for the measures used to calculate velocity.
type RiskEvaluationDetailsIpVelocityByUserVelocity struct {
	// This is the distinct count for the previous seconds indicated by the during value.
	DistinctCount *int32 `json:"distinctCount,omitempty"`
	// The interval (in seconds) during which the velocity is calculated.
	During *int32 `json:"during,omitempty"`
}

// NewRiskEvaluationDetailsIpVelocityByUserVelocity instantiates a new RiskEvaluationDetailsIpVelocityByUserVelocity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsIpVelocityByUserVelocity() *RiskEvaluationDetailsIpVelocityByUserVelocity {
	this := RiskEvaluationDetailsIpVelocityByUserVelocity{}
	return &this
}

// NewRiskEvaluationDetailsIpVelocityByUserVelocityWithDefaults instantiates a new RiskEvaluationDetailsIpVelocityByUserVelocity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsIpVelocityByUserVelocityWithDefaults() *RiskEvaluationDetailsIpVelocityByUserVelocity {
	this := RiskEvaluationDetailsIpVelocityByUserVelocity{}
	return &this
}

// GetDistinctCount returns the DistinctCount field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) GetDistinctCount() int32 {
	if o == nil || IsNil(o.DistinctCount) {
		var ret int32
		return ret
	}
	return *o.DistinctCount
}

// GetDistinctCountOk returns a tuple with the DistinctCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) GetDistinctCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DistinctCount) {
		return nil, false
	}
	return o.DistinctCount, true
}

// HasDistinctCount returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) HasDistinctCount() bool {
	if o != nil && !IsNil(o.DistinctCount) {
		return true
	}

	return false
}

// SetDistinctCount gets a reference to the given int32 and assigns it to the DistinctCount field.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) SetDistinctCount(v int32) {
	o.DistinctCount = &v
}

// GetDuring returns the During field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) GetDuring() int32 {
	if o == nil || IsNil(o.During) {
		var ret int32
		return ret
	}
	return *o.During
}

// GetDuringOk returns a tuple with the During field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) GetDuringOk() (*int32, bool) {
	if o == nil || IsNil(o.During) {
		return nil, false
	}
	return o.During, true
}

// HasDuring returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) HasDuring() bool {
	if o != nil && !IsNil(o.During) {
		return true
	}

	return false
}

// SetDuring gets a reference to the given int32 and assigns it to the During field.
func (o *RiskEvaluationDetailsIpVelocityByUserVelocity) SetDuring(v int32) {
	o.During = &v
}

func (o RiskEvaluationDetailsIpVelocityByUserVelocity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationDetailsIpVelocityByUserVelocity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DistinctCount) {
		toSerialize["distinctCount"] = o.DistinctCount
	}
	if !IsNil(o.During) {
		toSerialize["during"] = o.During
	}
	return toSerialize, nil
}

type NullableRiskEvaluationDetailsIpVelocityByUserVelocity struct {
	value *RiskEvaluationDetailsIpVelocityByUserVelocity
	isSet bool
}

func (v NullableRiskEvaluationDetailsIpVelocityByUserVelocity) Get() *RiskEvaluationDetailsIpVelocityByUserVelocity {
	return v.value
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUserVelocity) Set(val *RiskEvaluationDetailsIpVelocityByUserVelocity) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsIpVelocityByUserVelocity) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUserVelocity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsIpVelocityByUserVelocity(val *RiskEvaluationDetailsIpVelocityByUserVelocity) *NullableRiskEvaluationDetailsIpVelocityByUserVelocity {
	return &NullableRiskEvaluationDetailsIpVelocityByUserVelocity{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsIpVelocityByUserVelocity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUserVelocity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


