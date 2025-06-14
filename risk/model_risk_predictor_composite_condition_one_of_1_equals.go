/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// RiskPredictorCompositeConditionOneOf1Equals - struct for RiskPredictorCompositeConditionOneOf1Equals
type RiskPredictorCompositeConditionOneOf1Equals struct {
	Int32  *int32
	String *string
}

// int32AsRiskPredictorCompositeConditionOneOf1Equals is a convenience function that returns int32 wrapped in RiskPredictorCompositeConditionOneOf1Equals
func Int32AsRiskPredictorCompositeConditionOneOf1Equals(v *int32) RiskPredictorCompositeConditionOneOf1Equals {
	return RiskPredictorCompositeConditionOneOf1Equals{
		Int32: v,
	}
}

// stringAsRiskPredictorCompositeConditionOneOf1Equals is a convenience function that returns string wrapped in RiskPredictorCompositeConditionOneOf1Equals
func StringAsRiskPredictorCompositeConditionOneOf1Equals(v *string) RiskPredictorCompositeConditionOneOf1Equals {
	return RiskPredictorCompositeConditionOneOf1Equals{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskPredictorCompositeConditionOneOf1Equals) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RiskPredictorCompositeConditionOneOf1Equals)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RiskPredictorCompositeConditionOneOf1Equals)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskPredictorCompositeConditionOneOf1Equals) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskPredictorCompositeConditionOneOf1Equals) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableRiskPredictorCompositeConditionOneOf1Equals struct {
	value *RiskPredictorCompositeConditionOneOf1Equals
	isSet bool
}

func (v NullableRiskPredictorCompositeConditionOneOf1Equals) Get() *RiskPredictorCompositeConditionOneOf1Equals {
	return v.value
}

func (v *NullableRiskPredictorCompositeConditionOneOf1Equals) Set(val *RiskPredictorCompositeConditionOneOf1Equals) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeConditionOneOf1Equals) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeConditionOneOf1Equals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeConditionOneOf1Equals(val *RiskPredictorCompositeConditionOneOf1Equals) *NullableRiskPredictorCompositeConditionOneOf1Equals {
	return &NullableRiskPredictorCompositeConditionOneOf1Equals{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeConditionOneOf1Equals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeConditionOneOf1Equals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
