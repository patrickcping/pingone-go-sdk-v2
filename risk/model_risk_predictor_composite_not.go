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

// checks if the RiskPredictorCompositeNot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCompositeNot{}

// RiskPredictorCompositeNot struct for RiskPredictorCompositeNot
type RiskPredictorCompositeNot struct {
	Not RiskPredictorCompositeOr `json:"not"`
	Type *EnumPredictorCompositeConditionType `json:"type,omitempty"`
}

// NewRiskPredictorCompositeNot instantiates a new RiskPredictorCompositeNot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCompositeNot(not RiskPredictorCompositeOr) *RiskPredictorCompositeNot {
	this := RiskPredictorCompositeNot{}
	this.Not = not
	return &this
}

// NewRiskPredictorCompositeNotWithDefaults instantiates a new RiskPredictorCompositeNot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCompositeNotWithDefaults() *RiskPredictorCompositeNot {
	this := RiskPredictorCompositeNot{}
	return &this
}

// GetNot returns the Not field value
func (o *RiskPredictorCompositeNot) GetNot() RiskPredictorCompositeOr {
	if o == nil {
		var ret RiskPredictorCompositeOr
		return ret
	}

	return o.Not
}

// GetNotOk returns a tuple with the Not field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeNot) GetNotOk() (*RiskPredictorCompositeOr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Not, true
}

// SetNot sets field value
func (o *RiskPredictorCompositeNot) SetNot(v RiskPredictorCompositeOr) {
	o.Not = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPredictorCompositeNot) GetType() EnumPredictorCompositeConditionType {
	if o == nil || IsNil(o.Type) {
		var ret EnumPredictorCompositeConditionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeNot) GetTypeOk() (*EnumPredictorCompositeConditionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPredictorCompositeNot) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumPredictorCompositeConditionType and assigns it to the Type field.
func (o *RiskPredictorCompositeNot) SetType(v EnumPredictorCompositeConditionType) {
	o.Type = &v
}

func (o RiskPredictorCompositeNot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCompositeNot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["not"] = o.Not
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRiskPredictorCompositeNot struct {
	value *RiskPredictorCompositeNot
	isSet bool
}

func (v NullableRiskPredictorCompositeNot) Get() *RiskPredictorCompositeNot {
	return v.value
}

func (v *NullableRiskPredictorCompositeNot) Set(val *RiskPredictorCompositeNot) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeNot) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeNot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeNot(val *RiskPredictorCompositeNot) *NullableRiskPredictorCompositeNot {
	return &NullableRiskPredictorCompositeNot{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeNot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeNot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


