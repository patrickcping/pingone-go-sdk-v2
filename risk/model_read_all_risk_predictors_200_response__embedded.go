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

// checks if the ReadAllRiskPredictors200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllRiskPredictors200ResponseEmbedded{}

// ReadAllRiskPredictors200ResponseEmbedded struct for ReadAllRiskPredictors200ResponseEmbedded
type ReadAllRiskPredictors200ResponseEmbedded struct {
	RiskPredictors []RiskPredictor `json:"riskPredictors,omitempty"`
}

// NewReadAllRiskPredictors200ResponseEmbedded instantiates a new ReadAllRiskPredictors200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllRiskPredictors200ResponseEmbedded() *ReadAllRiskPredictors200ResponseEmbedded {
	this := ReadAllRiskPredictors200ResponseEmbedded{}
	return &this
}

// NewReadAllRiskPredictors200ResponseEmbeddedWithDefaults instantiates a new ReadAllRiskPredictors200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllRiskPredictors200ResponseEmbeddedWithDefaults() *ReadAllRiskPredictors200ResponseEmbedded {
	this := ReadAllRiskPredictors200ResponseEmbedded{}
	return &this
}

// GetRiskPredictors returns the RiskPredictors field value if set, zero value otherwise.
func (o *ReadAllRiskPredictors200ResponseEmbedded) GetRiskPredictors() []RiskPredictor {
	if o == nil || IsNil(o.RiskPredictors) {
		var ret []RiskPredictor
		return ret
	}
	return o.RiskPredictors
}

// GetRiskPredictorsOk returns a tuple with the RiskPredictors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllRiskPredictors200ResponseEmbedded) GetRiskPredictorsOk() ([]RiskPredictor, bool) {
	if o == nil || IsNil(o.RiskPredictors) {
		return nil, false
	}
	return o.RiskPredictors, true
}

// HasRiskPredictors returns a boolean if a field has been set.
func (o *ReadAllRiskPredictors200ResponseEmbedded) HasRiskPredictors() bool {
	if o != nil && !IsNil(o.RiskPredictors) {
		return true
	}

	return false
}

// SetRiskPredictors gets a reference to the given []RiskPredictor and assigns it to the RiskPredictors field.
func (o *ReadAllRiskPredictors200ResponseEmbedded) SetRiskPredictors(v []RiskPredictor) {
	o.RiskPredictors = v
}

func (o ReadAllRiskPredictors200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllRiskPredictors200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RiskPredictors) {
		toSerialize["riskPredictors"] = o.RiskPredictors
	}
	return toSerialize, nil
}

type NullableReadAllRiskPredictors200ResponseEmbedded struct {
	value *ReadAllRiskPredictors200ResponseEmbedded
	isSet bool
}

func (v NullableReadAllRiskPredictors200ResponseEmbedded) Get() *ReadAllRiskPredictors200ResponseEmbedded {
	return v.value
}

func (v *NullableReadAllRiskPredictors200ResponseEmbedded) Set(val *ReadAllRiskPredictors200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllRiskPredictors200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllRiskPredictors200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllRiskPredictors200ResponseEmbedded(val *ReadAllRiskPredictors200ResponseEmbedded) *NullableReadAllRiskPredictors200ResponseEmbedded {
	return &NullableReadAllRiskPredictors200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadAllRiskPredictors200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllRiskPredictors200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


