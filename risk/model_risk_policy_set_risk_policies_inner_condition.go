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

// RiskPolicySetRiskPoliciesInnerCondition The condition logic that determines when a policy is evaluated to true and when it is evaluated to false.
type RiskPolicySetRiskPoliciesInnerCondition struct {
	Contains *string `json:"contains,omitempty"`
	IpRange []string `json:"ipRange,omitempty"`
	Value *string `json:"value,omitempty"`
	Equals *RiskPolicySetRiskPoliciesInnerConditionEquals `json:"equals,omitempty"`
	AggregatedWeights []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner `json:"aggregatedWeights,omitempty"`
	Between *RiskPolicySetRiskPoliciesInnerConditionBetween `json:"between,omitempty"`
}

// NewRiskPolicySetRiskPoliciesInnerCondition instantiates a new RiskPolicySetRiskPoliciesInnerCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySetRiskPoliciesInnerCondition() *RiskPolicySetRiskPoliciesInnerCondition {
	this := RiskPolicySetRiskPoliciesInnerCondition{}
	return &this
}

// NewRiskPolicySetRiskPoliciesInnerConditionWithDefaults instantiates a new RiskPolicySetRiskPoliciesInnerCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetRiskPoliciesInnerConditionWithDefaults() *RiskPolicySetRiskPoliciesInnerCondition {
	this := RiskPolicySetRiskPoliciesInnerCondition{}
	return &this
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetContains() string {
	if o == nil || o.Contains == nil {
		var ret string
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetContainsOk() (*string, bool) {
	if o == nil || o.Contains == nil {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasContains() bool {
	if o != nil && o.Contains != nil {
		return true
	}

	return false
}

// SetContains gets a reference to the given string and assigns it to the Contains field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetContains(v string) {
	o.Contains = &v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetIpRange() []string {
	if o == nil || o.IpRange == nil {
		var ret []string
		return ret
	}
	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetIpRangeOk() ([]string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given []string and assigns it to the IpRange field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetIpRange(v []string) {
	o.IpRange = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetValue(v string) {
	o.Value = &v
}

// GetEquals returns the Equals field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetEquals() RiskPolicySetRiskPoliciesInnerConditionEquals {
	if o == nil || o.Equals == nil {
		var ret RiskPolicySetRiskPoliciesInnerConditionEquals
		return ret
	}
	return *o.Equals
}

// GetEqualsOk returns a tuple with the Equals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetEqualsOk() (*RiskPolicySetRiskPoliciesInnerConditionEquals, bool) {
	if o == nil || o.Equals == nil {
		return nil, false
	}
	return o.Equals, true
}

// HasEquals returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasEquals() bool {
	if o != nil && o.Equals != nil {
		return true
	}

	return false
}

// SetEquals gets a reference to the given RiskPolicySetRiskPoliciesInnerConditionEquals and assigns it to the Equals field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetEquals(v RiskPolicySetRiskPoliciesInnerConditionEquals) {
	o.Equals = &v
}

// GetAggregatedWeights returns the AggregatedWeights field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetAggregatedWeights() []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner {
	if o == nil || o.AggregatedWeights == nil {
		var ret []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner
		return ret
	}
	return o.AggregatedWeights
}

// GetAggregatedWeightsOk returns a tuple with the AggregatedWeights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetAggregatedWeightsOk() ([]RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner, bool) {
	if o == nil || o.AggregatedWeights == nil {
		return nil, false
	}
	return o.AggregatedWeights, true
}

// HasAggregatedWeights returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasAggregatedWeights() bool {
	if o != nil && o.AggregatedWeights != nil {
		return true
	}

	return false
}

// SetAggregatedWeights gets a reference to the given []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner and assigns it to the AggregatedWeights field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetAggregatedWeights(v []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner) {
	o.AggregatedWeights = v
}

// GetBetween returns the Between field value if set, zero value otherwise.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetBetween() RiskPolicySetRiskPoliciesInnerConditionBetween {
	if o == nil || o.Between == nil {
		var ret RiskPolicySetRiskPoliciesInnerConditionBetween
		return ret
	}
	return *o.Between
}

// GetBetweenOk returns a tuple with the Between field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) GetBetweenOk() (*RiskPolicySetRiskPoliciesInnerConditionBetween, bool) {
	if o == nil || o.Between == nil {
		return nil, false
	}
	return o.Between, true
}

// HasBetween returns a boolean if a field has been set.
func (o *RiskPolicySetRiskPoliciesInnerCondition) HasBetween() bool {
	if o != nil && o.Between != nil {
		return true
	}

	return false
}

// SetBetween gets a reference to the given RiskPolicySetRiskPoliciesInnerConditionBetween and assigns it to the Between field.
func (o *RiskPolicySetRiskPoliciesInnerCondition) SetBetween(v RiskPolicySetRiskPoliciesInnerConditionBetween) {
	o.Between = &v
}

func (o RiskPolicySetRiskPoliciesInnerCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contains != nil {
		toSerialize["contains"] = o.Contains
	}
	if o.IpRange != nil {
		toSerialize["ipRange"] = o.IpRange
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Equals != nil {
		toSerialize["equals"] = o.Equals
	}
	if o.AggregatedWeights != nil {
		toSerialize["aggregatedWeights"] = o.AggregatedWeights
	}
	if o.Between != nil {
		toSerialize["between"] = o.Between
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPolicySetRiskPoliciesInnerCondition struct {
	value *RiskPolicySetRiskPoliciesInnerCondition
	isSet bool
}

func (v NullableRiskPolicySetRiskPoliciesInnerCondition) Get() *RiskPolicySetRiskPoliciesInnerCondition {
	return v.value
}

func (v *NullableRiskPolicySetRiskPoliciesInnerCondition) Set(val *RiskPolicySetRiskPoliciesInnerCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySetRiskPoliciesInnerCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySetRiskPoliciesInnerCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySetRiskPoliciesInnerCondition(val *RiskPolicySetRiskPoliciesInnerCondition) *NullableRiskPolicySetRiskPoliciesInnerCondition {
	return &NullableRiskPolicySetRiskPoliciesInnerCondition{value: val, isSet: true}
}

func (v NullableRiskPolicySetRiskPoliciesInnerCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySetRiskPoliciesInnerCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


