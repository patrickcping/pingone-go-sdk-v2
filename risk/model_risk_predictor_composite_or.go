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

// checks if the RiskPredictorCompositeOr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCompositeOr{}

// RiskPredictorCompositeOr struct for RiskPredictorCompositeOr
type RiskPredictorCompositeOr struct {
	Or []RiskPredictorCompositeCondition `json:"or"`
	Type *EnumPredictorCompositeConditionType `json:"type,omitempty"`
}

type _RiskPredictorCompositeOr RiskPredictorCompositeOr

// NewRiskPredictorCompositeOr instantiates a new RiskPredictorCompositeOr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCompositeOr(or []RiskPredictorCompositeCondition) *RiskPredictorCompositeOr {
	this := RiskPredictorCompositeOr{}
	this.Or = or
	return &this
}

// NewRiskPredictorCompositeOrWithDefaults instantiates a new RiskPredictorCompositeOr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCompositeOrWithDefaults() *RiskPredictorCompositeOr {
	this := RiskPredictorCompositeOr{}
	return &this
}

// GetOr returns the Or field value
func (o *RiskPredictorCompositeOr) GetOr() []RiskPredictorCompositeCondition {
	if o == nil {
		var ret []RiskPredictorCompositeCondition
		return ret
	}

	return o.Or
}

// GetOrOk returns a tuple with the Or field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeOr) GetOrOk() ([]RiskPredictorCompositeCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Or, true
}

// SetOr sets field value
func (o *RiskPredictorCompositeOr) SetOr(v []RiskPredictorCompositeCondition) {
	o.Or = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RiskPredictorCompositeOr) GetType() EnumPredictorCompositeConditionType {
	if o == nil || IsNil(o.Type) {
		var ret EnumPredictorCompositeConditionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCompositeOr) GetTypeOk() (*EnumPredictorCompositeConditionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RiskPredictorCompositeOr) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumPredictorCompositeConditionType and assigns it to the Type field.
func (o *RiskPredictorCompositeOr) SetType(v EnumPredictorCompositeConditionType) {
	o.Type = &v
}

func (o RiskPredictorCompositeOr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCompositeOr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["or"] = o.Or
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *RiskPredictorCompositeOr) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"or",
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

	varRiskPredictorCompositeOr := _RiskPredictorCompositeOr{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPredictorCompositeOr)

	if err != nil {
		return err
	}

	*o = RiskPredictorCompositeOr(varRiskPredictorCompositeOr)

	return err
}

type NullableRiskPredictorCompositeOr struct {
	value *RiskPredictorCompositeOr
	isSet bool
}

func (v NullableRiskPredictorCompositeOr) Get() *RiskPredictorCompositeOr {
	return v.value
}

func (v *NullableRiskPredictorCompositeOr) Set(val *RiskPredictorCompositeOr) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeOr) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeOr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeOr(val *RiskPredictorCompositeOr) *NullableRiskPredictorCompositeOr {
	return &NullableRiskPredictorCompositeOr{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeOr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeOr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


