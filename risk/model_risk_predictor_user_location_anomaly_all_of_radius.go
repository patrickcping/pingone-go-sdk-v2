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

// checks if the RiskPredictorUserLocationAnomalyAllOfRadius type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorUserLocationAnomalyAllOfRadius{}

// RiskPredictorUserLocationAnomalyAllOfRadius struct for RiskPredictorUserLocationAnomalyAllOfRadius
type RiskPredictorUserLocationAnomalyAllOfRadius struct {
	// A value to apply to the distance radius.  Minimum of 10 miles (16 km) and maximum of 100 miles (160 km)
	Distance int32 `json:"distance"`
	Unit EnumDistanceUnit `json:"unit"`
}

// NewRiskPredictorUserLocationAnomalyAllOfRadius instantiates a new RiskPredictorUserLocationAnomalyAllOfRadius object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorUserLocationAnomalyAllOfRadius(distance int32, unit EnumDistanceUnit) *RiskPredictorUserLocationAnomalyAllOfRadius {
	this := RiskPredictorUserLocationAnomalyAllOfRadius{}
	this.Distance = distance
	this.Unit = unit
	return &this
}

// NewRiskPredictorUserLocationAnomalyAllOfRadiusWithDefaults instantiates a new RiskPredictorUserLocationAnomalyAllOfRadius object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorUserLocationAnomalyAllOfRadiusWithDefaults() *RiskPredictorUserLocationAnomalyAllOfRadius {
	this := RiskPredictorUserLocationAnomalyAllOfRadius{}
	return &this
}

// GetDistance returns the Distance field value
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetDistanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distance, true
}

// SetDistance sets field value
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) SetDistance(v int32) {
	o.Distance = v
}

// GetUnit returns the Unit field value
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetUnit() EnumDistanceUnit {
	if o == nil {
		var ret EnumDistanceUnit
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetUnitOk() (*EnumDistanceUnit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *RiskPredictorUserLocationAnomalyAllOfRadius) SetUnit(v EnumDistanceUnit) {
	o.Unit = v
}

func (o RiskPredictorUserLocationAnomalyAllOfRadius) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorUserLocationAnomalyAllOfRadius) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["distance"] = o.Distance
	toSerialize["unit"] = o.Unit
	return toSerialize, nil
}

type NullableRiskPredictorUserLocationAnomalyAllOfRadius struct {
	value *RiskPredictorUserLocationAnomalyAllOfRadius
	isSet bool
}

func (v NullableRiskPredictorUserLocationAnomalyAllOfRadius) Get() *RiskPredictorUserLocationAnomalyAllOfRadius {
	return v.value
}

func (v *NullableRiskPredictorUserLocationAnomalyAllOfRadius) Set(val *RiskPredictorUserLocationAnomalyAllOfRadius) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorUserLocationAnomalyAllOfRadius) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorUserLocationAnomalyAllOfRadius) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorUserLocationAnomalyAllOfRadius(val *RiskPredictorUserLocationAnomalyAllOfRadius) *NullableRiskPredictorUserLocationAnomalyAllOfRadius {
	return &NullableRiskPredictorUserLocationAnomalyAllOfRadius{value: val, isSet: true}
}

func (v NullableRiskPredictorUserLocationAnomalyAllOfRadius) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorUserLocationAnomalyAllOfRadius) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


