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

// checks if the RiskPredictorVelocityAllOfEvery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorVelocityAllOfEvery{}

// RiskPredictorVelocityAllOfEvery struct for RiskPredictorVelocityAllOfEvery
type RiskPredictorVelocityAllOfEvery struct {
	Unit *EnumPredictorUnit `json:"unit,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	MinSample *int32 `json:"minSample,omitempty"`
}

// NewRiskPredictorVelocityAllOfEvery instantiates a new RiskPredictorVelocityAllOfEvery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorVelocityAllOfEvery() *RiskPredictorVelocityAllOfEvery {
	this := RiskPredictorVelocityAllOfEvery{}
	return &this
}

// NewRiskPredictorVelocityAllOfEveryWithDefaults instantiates a new RiskPredictorVelocityAllOfEvery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorVelocityAllOfEveryWithDefaults() *RiskPredictorVelocityAllOfEvery {
	this := RiskPredictorVelocityAllOfEvery{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *RiskPredictorVelocityAllOfEvery) GetUnit() EnumPredictorUnit {
	if o == nil || IsNil(o.Unit) {
		var ret EnumPredictorUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorVelocityAllOfEvery) GetUnitOk() (*EnumPredictorUnit, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *RiskPredictorVelocityAllOfEvery) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given EnumPredictorUnit and assigns it to the Unit field.
func (o *RiskPredictorVelocityAllOfEvery) SetUnit(v EnumPredictorUnit) {
	o.Unit = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *RiskPredictorVelocityAllOfEvery) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorVelocityAllOfEvery) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *RiskPredictorVelocityAllOfEvery) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *RiskPredictorVelocityAllOfEvery) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetMinSample returns the MinSample field value if set, zero value otherwise.
func (o *RiskPredictorVelocityAllOfEvery) GetMinSample() int32 {
	if o == nil || IsNil(o.MinSample) {
		var ret int32
		return ret
	}
	return *o.MinSample
}

// GetMinSampleOk returns a tuple with the MinSample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorVelocityAllOfEvery) GetMinSampleOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSample) {
		return nil, false
	}
	return o.MinSample, true
}

// HasMinSample returns a boolean if a field has been set.
func (o *RiskPredictorVelocityAllOfEvery) HasMinSample() bool {
	if o != nil && !IsNil(o.MinSample) {
		return true
	}

	return false
}

// SetMinSample gets a reference to the given int32 and assigns it to the MinSample field.
func (o *RiskPredictorVelocityAllOfEvery) SetMinSample(v int32) {
	o.MinSample = &v
}

func (o RiskPredictorVelocityAllOfEvery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorVelocityAllOfEvery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.MinSample) {
		toSerialize["minSample"] = o.MinSample
	}
	return toSerialize, nil
}

type NullableRiskPredictorVelocityAllOfEvery struct {
	value *RiskPredictorVelocityAllOfEvery
	isSet bool
}

func (v NullableRiskPredictorVelocityAllOfEvery) Get() *RiskPredictorVelocityAllOfEvery {
	return v.value
}

func (v *NullableRiskPredictorVelocityAllOfEvery) Set(val *RiskPredictorVelocityAllOfEvery) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorVelocityAllOfEvery) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorVelocityAllOfEvery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorVelocityAllOfEvery(val *RiskPredictorVelocityAllOfEvery) *NullableRiskPredictorVelocityAllOfEvery {
	return &NullableRiskPredictorVelocityAllOfEvery{value: val, isSet: true}
}

func (v NullableRiskPredictorVelocityAllOfEvery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorVelocityAllOfEvery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


