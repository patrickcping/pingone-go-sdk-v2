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

// checks if the RiskEvaluationDetailsUserBasedRiskBehavior type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationDetailsUserBasedRiskBehavior{}

// RiskEvaluationDetailsUserBasedRiskBehavior struct for RiskEvaluationDetailsUserBasedRiskBehavior
type RiskEvaluationDetailsUserBasedRiskBehavior struct {
	// A string that specifies the risk score calculated for the user behavior for all transactions associated with the accessing device, and for the current authentication attempt. Options are LOW, MEDIUM, and HIGH. This information is available only if customers have agreed to data consent and the intelligence.allowDataConsent property in the PingOne license is set to true.
	Level *EnumRiskLevel `json:"level,omitempty"`
	// A string that describes the reason (or reasons) provided for the user behavior risk score classification (for example, the operating system or browser type used by the device, and country in which the accessing device is located). Each reason is classified as Unusual, to indicate how much it deviates from normal user behavior, and its effect in calculating the overall user behavior risk score.
	Reason *string `json:"reason,omitempty"`
}

// NewRiskEvaluationDetailsUserBasedRiskBehavior instantiates a new RiskEvaluationDetailsUserBasedRiskBehavior object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsUserBasedRiskBehavior() *RiskEvaluationDetailsUserBasedRiskBehavior {
	this := RiskEvaluationDetailsUserBasedRiskBehavior{}
	return &this
}

// NewRiskEvaluationDetailsUserBasedRiskBehaviorWithDefaults instantiates a new RiskEvaluationDetailsUserBasedRiskBehavior object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsUserBasedRiskBehaviorWithDefaults() *RiskEvaluationDetailsUserBasedRiskBehavior {
	this := RiskEvaluationDetailsUserBasedRiskBehavior{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetLevel() EnumRiskLevel {
	if o == nil || IsNil(o.Level) {
		var ret EnumRiskLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given EnumRiskLevel and assigns it to the Level field.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) SetLevel(v EnumRiskLevel) {
	o.Level = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RiskEvaluationDetailsUserBasedRiskBehavior) SetReason(v string) {
	o.Reason = &v
}

func (o RiskEvaluationDetailsUserBasedRiskBehavior) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationDetailsUserBasedRiskBehavior) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableRiskEvaluationDetailsUserBasedRiskBehavior struct {
	value *RiskEvaluationDetailsUserBasedRiskBehavior
	isSet bool
}

func (v NullableRiskEvaluationDetailsUserBasedRiskBehavior) Get() *RiskEvaluationDetailsUserBasedRiskBehavior {
	return v.value
}

func (v *NullableRiskEvaluationDetailsUserBasedRiskBehavior) Set(val *RiskEvaluationDetailsUserBasedRiskBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsUserBasedRiskBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsUserBasedRiskBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsUserBasedRiskBehavior(val *RiskEvaluationDetailsUserBasedRiskBehavior) *NullableRiskEvaluationDetailsUserBasedRiskBehavior {
	return &NullableRiskEvaluationDetailsUserBasedRiskBehavior{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsUserBasedRiskBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsUserBasedRiskBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


