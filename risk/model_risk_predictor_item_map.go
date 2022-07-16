/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
)

// RiskPredictorItemMap struct for RiskPredictorItemMap
type RiskPredictorItemMap struct {
	Contains string `json:"contains"`
	// List of CIDRs to include
	IpRange []string `json:"ipRange,omitempty"`
	Between *RiskPredictorItemMapBetween `json:"between,omitempty"`
	// An array that specifies the list of entities that represent the risk conditions.
	List []string `json:"list,omitempty"`
}

// NewRiskPredictorItemMap instantiates a new RiskPredictorItemMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorItemMap(contains string) *RiskPredictorItemMap {
	this := RiskPredictorItemMap{}
	this.Contains = contains
	return &this
}

// NewRiskPredictorItemMapWithDefaults instantiates a new RiskPredictorItemMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorItemMapWithDefaults() *RiskPredictorItemMap {
	this := RiskPredictorItemMap{}
	return &this
}

// GetContains returns the Contains field value
func (o *RiskPredictorItemMap) GetContains() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contains
}

// GetContainsOk returns a tuple with the Contains field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMap) GetContainsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contains, true
}

// SetContains sets field value
func (o *RiskPredictorItemMap) SetContains(v string) {
	o.Contains = v
}

// GetIpRange returns the IpRange field value if set, zero value otherwise.
func (o *RiskPredictorItemMap) GetIpRange() []string {
	if o == nil || o.IpRange == nil {
		var ret []string
		return ret
	}
	return o.IpRange
}

// GetIpRangeOk returns a tuple with the IpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMap) GetIpRangeOk() ([]string, bool) {
	if o == nil || o.IpRange == nil {
		return nil, false
	}
	return o.IpRange, true
}

// HasIpRange returns a boolean if a field has been set.
func (o *RiskPredictorItemMap) HasIpRange() bool {
	if o != nil && o.IpRange != nil {
		return true
	}

	return false
}

// SetIpRange gets a reference to the given []string and assigns it to the IpRange field.
func (o *RiskPredictorItemMap) SetIpRange(v []string) {
	o.IpRange = v
}

// GetBetween returns the Between field value if set, zero value otherwise.
func (o *RiskPredictorItemMap) GetBetween() RiskPredictorItemMapBetween {
	if o == nil || o.Between == nil {
		var ret RiskPredictorItemMapBetween
		return ret
	}
	return *o.Between
}

// GetBetweenOk returns a tuple with the Between field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMap) GetBetweenOk() (*RiskPredictorItemMapBetween, bool) {
	if o == nil || o.Between == nil {
		return nil, false
	}
	return o.Between, true
}

// HasBetween returns a boolean if a field has been set.
func (o *RiskPredictorItemMap) HasBetween() bool {
	if o != nil && o.Between != nil {
		return true
	}

	return false
}

// SetBetween gets a reference to the given RiskPredictorItemMapBetween and assigns it to the Between field.
func (o *RiskPredictorItemMap) SetBetween(v RiskPredictorItemMapBetween) {
	o.Between = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *RiskPredictorItemMap) GetList() []string {
	if o == nil || o.List == nil {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorItemMap) GetListOk() ([]string, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *RiskPredictorItemMap) HasList() bool {
	if o != nil && o.List != nil {
		return true
	}

	return false
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *RiskPredictorItemMap) SetList(v []string) {
	o.List = v
}

func (o RiskPredictorItemMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contains"] = o.Contains
	}
	if o.IpRange != nil {
		toSerialize["ipRange"] = o.IpRange
	}
	if o.Between != nil {
		toSerialize["between"] = o.Between
	}
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	return json.Marshal(toSerialize)
}

type NullableRiskPredictorItemMap struct {
	value *RiskPredictorItemMap
	isSet bool
}

func (v NullableRiskPredictorItemMap) Get() *RiskPredictorItemMap {
	return v.value
}

func (v *NullableRiskPredictorItemMap) Set(val *RiskPredictorItemMap) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorItemMap) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorItemMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorItemMap(val *RiskPredictorItemMap) *NullableRiskPredictorItemMap {
	return &NullableRiskPredictorItemMap{value: val, isSet: true}
}

func (v NullableRiskPredictorItemMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorItemMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


