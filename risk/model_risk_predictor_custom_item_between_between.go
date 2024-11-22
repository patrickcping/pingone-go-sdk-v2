/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RiskPredictorCustomItemBetweenBetween type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCustomItemBetweenBetween{}

// RiskPredictorCustomItemBetweenBetween Minimum and maximum boundaries
type RiskPredictorCustomItemBetweenBetween struct {
	MinScore float32 `json:"minScore"`
	MaxScore float32 `json:"maxScore"`
}

type _RiskPredictorCustomItemBetweenBetween RiskPredictorCustomItemBetweenBetween

// NewRiskPredictorCustomItemBetweenBetween instantiates a new RiskPredictorCustomItemBetweenBetween object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCustomItemBetweenBetween(minScore float32, maxScore float32) *RiskPredictorCustomItemBetweenBetween {
	this := RiskPredictorCustomItemBetweenBetween{}
	this.MinScore = minScore
	this.MaxScore = maxScore
	return &this
}

// NewRiskPredictorCustomItemBetweenBetweenWithDefaults instantiates a new RiskPredictorCustomItemBetweenBetween object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCustomItemBetweenBetweenWithDefaults() *RiskPredictorCustomItemBetweenBetween {
	this := RiskPredictorCustomItemBetweenBetween{}
	return &this
}

// GetMinScore returns the MinScore field value
func (o *RiskPredictorCustomItemBetweenBetween) GetMinScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinScore
}

// GetMinScoreOk returns a tuple with the MinScore field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustomItemBetweenBetween) GetMinScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinScore, true
}

// SetMinScore sets field value
func (o *RiskPredictorCustomItemBetweenBetween) SetMinScore(v float32) {
	o.MinScore = v
}

// GetMaxScore returns the MaxScore field value
func (o *RiskPredictorCustomItemBetweenBetween) GetMaxScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MaxScore
}

// GetMaxScoreOk returns a tuple with the MaxScore field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustomItemBetweenBetween) GetMaxScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxScore, true
}

// SetMaxScore sets field value
func (o *RiskPredictorCustomItemBetweenBetween) SetMaxScore(v float32) {
	o.MaxScore = v
}

func (o RiskPredictorCustomItemBetweenBetween) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCustomItemBetweenBetween) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["minScore"] = o.MinScore
	toSerialize["maxScore"] = o.MaxScore
	return toSerialize, nil
}

func (o *RiskPredictorCustomItemBetweenBetween) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"minScore",
		"maxScore",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskPredictorCustomItemBetweenBetween := _RiskPredictorCustomItemBetweenBetween{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPredictorCustomItemBetweenBetween)

	if err != nil {
		return err
	}

	*o = RiskPredictorCustomItemBetweenBetween(varRiskPredictorCustomItemBetweenBetween)

	return err
}

type NullableRiskPredictorCustomItemBetweenBetween struct {
	value *RiskPredictorCustomItemBetweenBetween
	isSet bool
}

func (v NullableRiskPredictorCustomItemBetweenBetween) Get() *RiskPredictorCustomItemBetweenBetween {
	return v.value
}

func (v *NullableRiskPredictorCustomItemBetweenBetween) Set(val *RiskPredictorCustomItemBetweenBetween) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCustomItemBetweenBetween) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCustomItemBetweenBetween) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCustomItemBetweenBetween(val *RiskPredictorCustomItemBetweenBetween) *NullableRiskPredictorCustomItemBetweenBetween {
	return &NullableRiskPredictorCustomItemBetweenBetween{value: val, isSet: true}
}

func (v NullableRiskPredictorCustomItemBetweenBetween) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCustomItemBetweenBetween) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


