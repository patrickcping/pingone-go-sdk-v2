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

// RiskPredictorCustomItem - struct for RiskPredictorCustomItem
type RiskPredictorCustomItem struct {
	RiskPredictorCustomItemBetween *RiskPredictorCustomItemBetween
	RiskPredictorCustomItemIPRange *RiskPredictorCustomItemIPRange
	RiskPredictorCustomItemList    *RiskPredictorCustomItemList
}

// RiskPredictorCustomItemBetweenAsRiskPredictorCustomItem is a convenience function that returns RiskPredictorCustomItemBetween wrapped in RiskPredictorCustomItem
func RiskPredictorCustomItemBetweenAsRiskPredictorCustomItem(v *RiskPredictorCustomItemBetween) RiskPredictorCustomItem {
	return RiskPredictorCustomItem{
		RiskPredictorCustomItemBetween: v,
	}
}

// RiskPredictorCustomItemIPRangeAsRiskPredictorCustomItem is a convenience function that returns RiskPredictorCustomItemIPRange wrapped in RiskPredictorCustomItem
func RiskPredictorCustomItemIPRangeAsRiskPredictorCustomItem(v *RiskPredictorCustomItemIPRange) RiskPredictorCustomItem {
	return RiskPredictorCustomItem{
		RiskPredictorCustomItemIPRange: v,
	}
}

// RiskPredictorCustomItemListAsRiskPredictorCustomItem is a convenience function that returns RiskPredictorCustomItemList wrapped in RiskPredictorCustomItem
func RiskPredictorCustomItemListAsRiskPredictorCustomItem(v *RiskPredictorCustomItemList) RiskPredictorCustomItem {
	return RiskPredictorCustomItem{
		RiskPredictorCustomItemList: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskPredictorCustomItem) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RiskPredictorCustomItemBetween
	err = newStrictDecoder(data).Decode(&dst.RiskPredictorCustomItemBetween)
	if err == nil {
		jsonRiskPredictorCustomItemBetween, _ := json.Marshal(dst.RiskPredictorCustomItemBetween)
		if string(jsonRiskPredictorCustomItemBetween) == "{}" { // empty struct
			dst.RiskPredictorCustomItemBetween = nil
		} else {
			match++
		}
	} else {
		dst.RiskPredictorCustomItemBetween = nil
	}

	// try to unmarshal data into RiskPredictorCustomItemIPRange
	err = newStrictDecoder(data).Decode(&dst.RiskPredictorCustomItemIPRange)
	if err == nil {
		jsonRiskPredictorCustomItemIPRange, _ := json.Marshal(dst.RiskPredictorCustomItemIPRange)
		if string(jsonRiskPredictorCustomItemIPRange) == "{}" { // empty struct
			dst.RiskPredictorCustomItemIPRange = nil
		} else {
			match++
		}
	} else {
		dst.RiskPredictorCustomItemIPRange = nil
	}

	// try to unmarshal data into RiskPredictorCustomItemList
	err = newStrictDecoder(data).Decode(&dst.RiskPredictorCustomItemList)
	if err == nil {
		jsonRiskPredictorCustomItemList, _ := json.Marshal(dst.RiskPredictorCustomItemList)
		if string(jsonRiskPredictorCustomItemList) == "{}" { // empty struct
			dst.RiskPredictorCustomItemList = nil
		} else {
			match++
		}
	} else {
		dst.RiskPredictorCustomItemList = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RiskPredictorCustomItemBetween = nil
		dst.RiskPredictorCustomItemIPRange = nil
		dst.RiskPredictorCustomItemList = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RiskPredictorCustomItem)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RiskPredictorCustomItem)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskPredictorCustomItem) MarshalJSON() ([]byte, error) {
	if src.RiskPredictorCustomItemBetween != nil {
		return json.Marshal(&src.RiskPredictorCustomItemBetween)
	}

	if src.RiskPredictorCustomItemIPRange != nil {
		return json.Marshal(&src.RiskPredictorCustomItemIPRange)
	}

	if src.RiskPredictorCustomItemList != nil {
		return json.Marshal(&src.RiskPredictorCustomItemList)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskPredictorCustomItem) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RiskPredictorCustomItemBetween != nil {
		return obj.RiskPredictorCustomItemBetween
	}

	if obj.RiskPredictorCustomItemIPRange != nil {
		return obj.RiskPredictorCustomItemIPRange
	}

	if obj.RiskPredictorCustomItemList != nil {
		return obj.RiskPredictorCustomItemList
	}

	// all schemas are nil
	return nil
}

type NullableRiskPredictorCustomItem struct {
	value *RiskPredictorCustomItem
	isSet bool
}

func (v NullableRiskPredictorCustomItem) Get() *RiskPredictorCustomItem {
	return v.value
}

func (v *NullableRiskPredictorCustomItem) Set(val *RiskPredictorCustomItem) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCustomItem) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCustomItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCustomItem(val *RiskPredictorCustomItem) *NullableRiskPredictorCustomItem {
	return &NullableRiskPredictorCustomItem{value: val, isSet: true}
}

func (v NullableRiskPredictorCustomItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCustomItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
