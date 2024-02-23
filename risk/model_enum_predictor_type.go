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

// EnumPredictorType An enum type that specifies the type of the predictor. Note that both the New Device and Suspicious Device predictors use the type DEVICE. To differentiate between them, you use the `detect` field.
type EnumPredictorType string

// List of EnumPredictorType
const (
	ENUMPREDICTORTYPE_ADVERSARY_IN_THE_MIDDLE EnumPredictorType = "ADVERSARY_IN_THE_MIDDLE"
	ENUMPREDICTORTYPE_ANONYMOUS_NETWORK EnumPredictorType = "ANONYMOUS_NETWORK"
	ENUMPREDICTORTYPE_BOT EnumPredictorType = "BOT"
	ENUMPREDICTORTYPE_COMPOSITE EnumPredictorType = "COMPOSITE"
	ENUMPREDICTORTYPE_GEO_VELOCITY EnumPredictorType = "GEO_VELOCITY"
	ENUMPREDICTORTYPE_IP_REPUTATION EnumPredictorType = "IP_REPUTATION"
	ENUMPREDICTORTYPE_MAP EnumPredictorType = "MAP"
	ENUMPREDICTORTYPE_DEVICE EnumPredictorType = "DEVICE"
	ENUMPREDICTORTYPE_USER_LOCATION_ANOMALY EnumPredictorType = "USER_LOCATION_ANOMALY"
	ENUMPREDICTORTYPE_USER_RISK_BEHAVIOR EnumPredictorType = "USER_RISK_BEHAVIOR"
	ENUMPREDICTORTYPE_VELOCITY EnumPredictorType = "VELOCITY"
)

// All allowed values of EnumPredictorType enum
var AllowedEnumPredictorTypeEnumValues = []EnumPredictorType{
	"ADVERSARY_IN_THE_MIDDLE",
	"ANONYMOUS_NETWORK",
	"BOT",
	"COMPOSITE",
	"GEO_VELOCITY",
	"IP_REPUTATION",
	"MAP",
	"DEVICE",
	"USER_LOCATION_ANOMALY",
	"USER_RISK_BEHAVIOR",
	"VELOCITY",
}

func (v *EnumPredictorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPredictorType(value)
	for _, existing := range AllowedEnumPredictorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumPredictorType(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumPredictorTypeFromValue returns a pointer to a valid EnumPredictorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPredictorTypeFromValue(v string) (*EnumPredictorType, error) {
	ev := EnumPredictorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPredictorType: valid values are %v", v, AllowedEnumPredictorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPredictorType) IsValid() bool {
	for _, existing := range AllowedEnumPredictorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPredictorType value
func (v EnumPredictorType) Ptr() *EnumPredictorType {
	return &v
}

type NullableEnumPredictorType struct {
	value *EnumPredictorType
	isSet bool
}

func (v NullableEnumPredictorType) Get() *EnumPredictorType {
	return v.value
}

func (v *NullableEnumPredictorType) Set(val *EnumPredictorType) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPredictorType) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPredictorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPredictorType(val *EnumPredictorType) *NullableEnumPredictorType {
	return &NullableEnumPredictorType{value: val, isSet: true}
}

func (v NullableEnumPredictorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPredictorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

