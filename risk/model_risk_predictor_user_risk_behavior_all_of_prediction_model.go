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

// checks if the RiskPredictorUserRiskBehaviorAllOfPredictionModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorUserRiskBehaviorAllOfPredictionModel{}

// RiskPredictorUserRiskBehaviorAllOfPredictionModel struct for RiskPredictorUserRiskBehaviorAllOfPredictionModel
type RiskPredictorUserRiskBehaviorAllOfPredictionModel struct {
	Name EnumUserRiskBehaviorRiskModel `json:"name"`
}

type _RiskPredictorUserRiskBehaviorAllOfPredictionModel RiskPredictorUserRiskBehaviorAllOfPredictionModel

// NewRiskPredictorUserRiskBehaviorAllOfPredictionModel instantiates a new RiskPredictorUserRiskBehaviorAllOfPredictionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorUserRiskBehaviorAllOfPredictionModel(name EnumUserRiskBehaviorRiskModel) *RiskPredictorUserRiskBehaviorAllOfPredictionModel {
	this := RiskPredictorUserRiskBehaviorAllOfPredictionModel{}
	this.Name = name
	return &this
}

// NewRiskPredictorUserRiskBehaviorAllOfPredictionModelWithDefaults instantiates a new RiskPredictorUserRiskBehaviorAllOfPredictionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorUserRiskBehaviorAllOfPredictionModelWithDefaults() *RiskPredictorUserRiskBehaviorAllOfPredictionModel {
	this := RiskPredictorUserRiskBehaviorAllOfPredictionModel{}
	return &this
}

// GetName returns the Name field value
func (o *RiskPredictorUserRiskBehaviorAllOfPredictionModel) GetName() EnumUserRiskBehaviorRiskModel {
	if o == nil {
		var ret EnumUserRiskBehaviorRiskModel
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorUserRiskBehaviorAllOfPredictionModel) GetNameOk() (*EnumUserRiskBehaviorRiskModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorUserRiskBehaviorAllOfPredictionModel) SetName(v EnumUserRiskBehaviorRiskModel) {
	o.Name = v
}

func (o RiskPredictorUserRiskBehaviorAllOfPredictionModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorUserRiskBehaviorAllOfPredictionModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *RiskPredictorUserRiskBehaviorAllOfPredictionModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varRiskPredictorUserRiskBehaviorAllOfPredictionModel := _RiskPredictorUserRiskBehaviorAllOfPredictionModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPredictorUserRiskBehaviorAllOfPredictionModel)

	if err != nil {
		return err
	}

	*o = RiskPredictorUserRiskBehaviorAllOfPredictionModel(varRiskPredictorUserRiskBehaviorAllOfPredictionModel)

	return err
}

type NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel struct {
	value *RiskPredictorUserRiskBehaviorAllOfPredictionModel
	isSet bool
}

func (v NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) Get() *RiskPredictorUserRiskBehaviorAllOfPredictionModel {
	return v.value
}

func (v *NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) Set(val *RiskPredictorUserRiskBehaviorAllOfPredictionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorUserRiskBehaviorAllOfPredictionModel(val *RiskPredictorUserRiskBehaviorAllOfPredictionModel) *NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel {
	return &NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel{value: val, isSet: true}
}

func (v NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorUserRiskBehaviorAllOfPredictionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


