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

// checks if the RiskEvaluationEventSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskEvaluationEventSession{}

// RiskEvaluationEventSession struct for RiskEvaluationEventSession
type RiskEvaluationEventSession struct {
	// A string that specifies a unique session ID associated with the event.
	Id string `json:"id"`
}

// NewRiskEvaluationEventSession instantiates a new RiskEvaluationEventSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskEvaluationEventSession(id string) *RiskEvaluationEventSession {
	this := RiskEvaluationEventSession{}
	this.Id = id
	return &this
}

// NewRiskEvaluationEventSessionWithDefaults instantiates a new RiskEvaluationEventSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskEvaluationEventSessionWithDefaults() *RiskEvaluationEventSession {
	this := RiskEvaluationEventSession{}
	return &this
}

// GetId returns the Id field value
func (o *RiskEvaluationEventSession) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskEvaluationEventSession) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskEvaluationEventSession) SetId(v string) {
	o.Id = v
}

func (o RiskEvaluationEventSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskEvaluationEventSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableRiskEvaluationEventSession struct {
	value *RiskEvaluationEventSession
	isSet bool
}

func (v NullableRiskEvaluationEventSession) Get() *RiskEvaluationEventSession {
	return v.value
}

func (v *NullableRiskEvaluationEventSession) Set(val *RiskEvaluationEventSession) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskEvaluationEventSession) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskEvaluationEventSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskEvaluationEventSession(val *RiskEvaluationEventSession) *NullableRiskEvaluationEventSession {
	return &NullableRiskEvaluationEventSession{value: val, isSet: true}
}

func (v NullableRiskEvaluationEventSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskEvaluationEventSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
