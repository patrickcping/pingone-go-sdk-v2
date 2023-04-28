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

// checks if the RiskPredictorCommonConditionScoresInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCommonConditionScoresInner{}

// RiskPredictorCommonConditionScoresInner struct for RiskPredictorCommonConditionScoresInner
type RiskPredictorCommonConditionScoresInner struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewRiskPredictorCommonConditionScoresInner instantiates a new RiskPredictorCommonConditionScoresInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCommonConditionScoresInner() *RiskPredictorCommonConditionScoresInner {
	this := RiskPredictorCommonConditionScoresInner{}
	return &this
}

// NewRiskPredictorCommonConditionScoresInnerWithDefaults instantiates a new RiskPredictorCommonConditionScoresInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCommonConditionScoresInnerWithDefaults() *RiskPredictorCommonConditionScoresInner {
	this := RiskPredictorCommonConditionScoresInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RiskPredictorCommonConditionScoresInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonConditionScoresInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RiskPredictorCommonConditionScoresInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RiskPredictorCommonConditionScoresInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskPredictorCommonConditionScoresInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommonConditionScoresInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskPredictorCommonConditionScoresInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskPredictorCommonConditionScoresInner) SetValue(v string) {
	o.Value = &v
}

func (o RiskPredictorCommonConditionScoresInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCommonConditionScoresInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableRiskPredictorCommonConditionScoresInner struct {
	value *RiskPredictorCommonConditionScoresInner
	isSet bool
}

func (v NullableRiskPredictorCommonConditionScoresInner) Get() *RiskPredictorCommonConditionScoresInner {
	return v.value
}

func (v *NullableRiskPredictorCommonConditionScoresInner) Set(val *RiskPredictorCommonConditionScoresInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCommonConditionScoresInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCommonConditionScoresInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCommonConditionScoresInner(val *RiskPredictorCommonConditionScoresInner) *NullableRiskPredictorCommonConditionScoresInner {
	return &NullableRiskPredictorCommonConditionScoresInner{value: val, isSet: true}
}

func (v NullableRiskPredictorCommonConditionScoresInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCommonConditionScoresInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

