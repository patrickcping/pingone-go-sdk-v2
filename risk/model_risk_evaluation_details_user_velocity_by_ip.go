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

// checks if the RiskEvaluationDetailsUserVelocityByIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationDetailsUserVelocityByIp{}

// RiskEvaluationDetailsUserVelocityByIp struct for RiskEvaluationDetailsUserVelocityByIp
type RiskEvaluationDetailsUserVelocityByIp struct {
	// An enum indicating whether the calculated number of users per IP is LOW, MEDIUM, or HIGH.
	Level *EnumRiskLevel `json:"level,omitempty"`
	// A string indicating the reason the user was flagged. For example \"More than 250 users accessed IP address 1.1.1.1 during the last 1 hour.\"
	Reason *string `json:"reason,omitempty"`
	Threshold *RiskEvaluationDetailsUserVelocityByIpThreshold `json:"threshold,omitempty"`
	Velocity *RiskEvaluationDetailsUserVelocityByIpVelocity `json:"velocity,omitempty"`
}

// NewRiskEvaluationDetailsUserVelocityByIp instantiates a new RiskEvaluationDetailsUserVelocityByIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsUserVelocityByIp() *RiskEvaluationDetailsUserVelocityByIp {
	this := RiskEvaluationDetailsUserVelocityByIp{}
	return &this
}

// NewRiskEvaluationDetailsUserVelocityByIpWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsUserVelocityByIpWithDefaults() *RiskEvaluationDetailsUserVelocityByIp {
	this := RiskEvaluationDetailsUserVelocityByIp{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevel() EnumRiskLevel {
	if o == nil || IsNil(o.Level) {
		var ret EnumRiskLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given EnumRiskLevel and assigns it to the Level field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetLevel(v EnumRiskLevel) {
	o.Level = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetReason(v string) {
	o.Reason = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetThreshold() RiskEvaluationDetailsUserVelocityByIpThreshold {
	if o == nil || IsNil(o.Threshold) {
		var ret RiskEvaluationDetailsUserVelocityByIpThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetThresholdOk() (*RiskEvaluationDetailsUserVelocityByIpThreshold, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given RiskEvaluationDetailsUserVelocityByIpThreshold and assigns it to the Threshold field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetThreshold(v RiskEvaluationDetailsUserVelocityByIpThreshold) {
	o.Threshold = &v
}

// GetVelocity returns the Velocity field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocity() RiskEvaluationDetailsUserVelocityByIpVelocity {
	if o == nil || IsNil(o.Velocity) {
		var ret RiskEvaluationDetailsUserVelocityByIpVelocity
		return ret
	}
	return *o.Velocity
}

// GetVelocityOk returns a tuple with the Velocity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocityOk() (*RiskEvaluationDetailsUserVelocityByIpVelocity, bool) {
	if o == nil || IsNil(o.Velocity) {
		return nil, false
	}
	return o.Velocity, true
}

// HasVelocity returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsUserVelocityByIp) HasVelocity() bool {
	if o != nil && !IsNil(o.Velocity) {
		return true
	}

	return false
}

// SetVelocity gets a reference to the given RiskEvaluationDetailsUserVelocityByIpVelocity and assigns it to the Velocity field.
func (o *RiskEvaluationDetailsUserVelocityByIp) SetVelocity(v RiskEvaluationDetailsUserVelocityByIpVelocity) {
	o.Velocity = &v
}

func (o RiskEvaluationDetailsUserVelocityByIp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationDetailsUserVelocityByIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level) {
		toSerialize["level"] = o.Level
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !IsNil(o.Velocity) {
		toSerialize["velocity"] = o.Velocity
	}
	return toSerialize, nil
}

type NullableRiskEvaluationDetailsUserVelocityByIp struct {
	value *RiskEvaluationDetailsUserVelocityByIp
	isSet bool
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) Get() *RiskEvaluationDetailsUserVelocityByIp {
	return v.value
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) Set(val *RiskEvaluationDetailsUserVelocityByIp) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsUserVelocityByIp(val *RiskEvaluationDetailsUserVelocityByIp) *NullableRiskEvaluationDetailsUserVelocityByIp {
	return &NullableRiskEvaluationDetailsUserVelocityByIp{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsUserVelocityByIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsUserVelocityByIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


