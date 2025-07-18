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

// RiskPredictorCompositeConditionBase - struct for RiskPredictorCompositeConditionBase
type RiskPredictorCompositeConditionBase struct {
	RiskPredictorCompositeAnd *RiskPredictorCompositeAnd
	RiskPredictorCompositeNot *RiskPredictorCompositeNot
	RiskPredictorCompositeOr  *RiskPredictorCompositeOr
}

// RiskPredictorCompositeAndAsRiskPredictorCompositeConditionBase is a convenience function that returns RiskPredictorCompositeAnd wrapped in RiskPredictorCompositeConditionBase
func RiskPredictorCompositeAndAsRiskPredictorCompositeConditionBase(v *RiskPredictorCompositeAnd) RiskPredictorCompositeConditionBase {
	return RiskPredictorCompositeConditionBase{
		RiskPredictorCompositeAnd: v,
	}
}

// RiskPredictorCompositeNotAsRiskPredictorCompositeConditionBase is a convenience function that returns RiskPredictorCompositeNot wrapped in RiskPredictorCompositeConditionBase
func RiskPredictorCompositeNotAsRiskPredictorCompositeConditionBase(v *RiskPredictorCompositeNot) RiskPredictorCompositeConditionBase {
	return RiskPredictorCompositeConditionBase{
		RiskPredictorCompositeNot: v,
	}
}

// RiskPredictorCompositeOrAsRiskPredictorCompositeConditionBase is a convenience function that returns RiskPredictorCompositeOr wrapped in RiskPredictorCompositeConditionBase
func RiskPredictorCompositeOrAsRiskPredictorCompositeConditionBase(v *RiskPredictorCompositeOr) RiskPredictorCompositeConditionBase {
	return RiskPredictorCompositeConditionBase{
		RiskPredictorCompositeOr: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskPredictorCompositeConditionBase) UnmarshalJSON(data []byte) error {
	
	match := 0
	var common map[string]interface{}

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.RiskPredictorCompositeAnd = nil
	dst.RiskPredictorCompositeNot = nil
	dst.RiskPredictorCompositeOr = nil

	if common["and"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeAnd); err != nil {
			return err
		}
		match++
	}

	if common["or"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeOr); err != nil {
			return err
		}
		match++
	}

	if common["not"] != nil {
		if err := json.Unmarshal(data, &dst.RiskPredictorCompositeNot); err != nil {
			return err
		}
		match++
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RiskPredictorCompositeAnd = nil
		dst.RiskPredictorCompositeNot = nil
		dst.RiskPredictorCompositeOr = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RiskPredictorCompositeConditionBase)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RiskPredictorCompositeConditionBase)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskPredictorCompositeConditionBase) MarshalJSON() ([]byte, error) {
	if src.RiskPredictorCompositeAnd != nil {
		return json.Marshal(&src.RiskPredictorCompositeAnd)
	}

	if src.RiskPredictorCompositeNot != nil {
		return json.Marshal(&src.RiskPredictorCompositeNot)
	}

	if src.RiskPredictorCompositeOr != nil {
		return json.Marshal(&src.RiskPredictorCompositeOr)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskPredictorCompositeConditionBase) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RiskPredictorCompositeAnd != nil {
		return obj.RiskPredictorCompositeAnd
	}

	if obj.RiskPredictorCompositeNot != nil {
		return obj.RiskPredictorCompositeNot
	}

	if obj.RiskPredictorCompositeOr != nil {
		return obj.RiskPredictorCompositeOr
	}

	// all schemas are nil
	return nil
}

type NullableRiskPredictorCompositeConditionBase struct {
	value *RiskPredictorCompositeConditionBase
	isSet bool
}

func (v NullableRiskPredictorCompositeConditionBase) Get() *RiskPredictorCompositeConditionBase {
	return v.value
}

func (v *NullableRiskPredictorCompositeConditionBase) Set(val *RiskPredictorCompositeConditionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCompositeConditionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCompositeConditionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCompositeConditionBase(val *RiskPredictorCompositeConditionBase) *NullableRiskPredictorCompositeConditionBase {
	return &NullableRiskPredictorCompositeConditionBase{value: val, isSet: true}
}

func (v NullableRiskPredictorCompositeConditionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCompositeConditionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
