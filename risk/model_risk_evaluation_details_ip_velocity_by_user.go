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

// checks if the RiskEvaluationDetailsIpVelocityByUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationDetailsIpVelocityByUser{}

// RiskEvaluationDetailsIpVelocityByUser struct for RiskEvaluationDetailsIpVelocityByUser
type RiskEvaluationDetailsIpVelocityByUser struct {
	Level *EnumRiskLevel `json:"level,omitempty"`
	// A string indicating the reason the user was flagged. For example \"More than 13 IPs were accessed by John during the last 1 hour.\"
	Reason *string `json:"reason,omitempty"`
	Threshold *RiskEvaluationDetailsIpVelocityByUserThreshold `json:"threshold,omitempty"`
	Velocity *RiskEvaluationDetailsIpVelocityByUserVelocity `json:"velocity,omitempty"`
}

// NewRiskEvaluationDetailsIpVelocityByUser instantiates a new RiskEvaluationDetailsIpVelocityByUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationDetailsIpVelocityByUser() *RiskEvaluationDetailsIpVelocityByUser {
	this := RiskEvaluationDetailsIpVelocityByUser{}
	return &this
}

// NewRiskEvaluationDetailsIpVelocityByUserWithDefaults instantiates a new RiskEvaluationDetailsIpVelocityByUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationDetailsIpVelocityByUserWithDefaults() *RiskEvaluationDetailsIpVelocityByUser {
	this := RiskEvaluationDetailsIpVelocityByUser{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetLevel() EnumRiskLevel {
	if o == nil || IsNil(o.Level) {
		var ret EnumRiskLevel
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetLevelOk() (*EnumRiskLevel, bool) {
	if o == nil || IsNil(o.Level) {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) HasLevel() bool {
	if o != nil && !IsNil(o.Level) {
		return true
	}

	return false
}

// SetLevel gets a reference to the given EnumRiskLevel and assigns it to the Level field.
func (o *RiskEvaluationDetailsIpVelocityByUser) SetLevel(v EnumRiskLevel) {
	o.Level = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RiskEvaluationDetailsIpVelocityByUser) SetReason(v string) {
	o.Reason = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetThreshold() RiskEvaluationDetailsIpVelocityByUserThreshold {
	if o == nil || IsNil(o.Threshold) {
		var ret RiskEvaluationDetailsIpVelocityByUserThreshold
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetThresholdOk() (*RiskEvaluationDetailsIpVelocityByUserThreshold, bool) {
	if o == nil || IsNil(o.Threshold) {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) HasThreshold() bool {
	if o != nil && !IsNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given RiskEvaluationDetailsIpVelocityByUserThreshold and assigns it to the Threshold field.
func (o *RiskEvaluationDetailsIpVelocityByUser) SetThreshold(v RiskEvaluationDetailsIpVelocityByUserThreshold) {
	o.Threshold = &v
}

// GetVelocity returns the Velocity field value if set, zero value otherwise.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetVelocity() RiskEvaluationDetailsIpVelocityByUserVelocity {
	if o == nil || IsNil(o.Velocity) {
		var ret RiskEvaluationDetailsIpVelocityByUserVelocity
		return ret
	}
	return *o.Velocity
}

// GetVelocityOk returns a tuple with the Velocity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) GetVelocityOk() (*RiskEvaluationDetailsIpVelocityByUserVelocity, bool) {
	if o == nil || IsNil(o.Velocity) {
		return nil, false
	}
	return o.Velocity, true
}

// HasVelocity returns a boolean if a field has been set.
func (o *RiskEvaluationDetailsIpVelocityByUser) HasVelocity() bool {
	if o != nil && !IsNil(o.Velocity) {
		return true
	}

	return false
}

// SetVelocity gets a reference to the given RiskEvaluationDetailsIpVelocityByUserVelocity and assigns it to the Velocity field.
func (o *RiskEvaluationDetailsIpVelocityByUser) SetVelocity(v RiskEvaluationDetailsIpVelocityByUserVelocity) {
	o.Velocity = &v
}

func (o RiskEvaluationDetailsIpVelocityByUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationDetailsIpVelocityByUser) ToMap() (map[string]interface{}, error) {
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

type NullableRiskEvaluationDetailsIpVelocityByUser struct {
	value *RiskEvaluationDetailsIpVelocityByUser
	isSet bool
}

func (v NullableRiskEvaluationDetailsIpVelocityByUser) Get() *RiskEvaluationDetailsIpVelocityByUser {
	return v.value
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUser) Set(val *RiskEvaluationDetailsIpVelocityByUser) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationDetailsIpVelocityByUser) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationDetailsIpVelocityByUser(val *RiskEvaluationDetailsIpVelocityByUser) *NullableRiskEvaluationDetailsIpVelocityByUser {
	return &NullableRiskEvaluationDetailsIpVelocityByUser{value: val, isSet: true}
}

func (v NullableRiskEvaluationDetailsIpVelocityByUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationDetailsIpVelocityByUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


