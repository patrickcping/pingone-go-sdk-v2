/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the EnvironmentStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentStatus{}

// EnvironmentStatus struct for EnvironmentStatus
type EnvironmentStatus struct {
	Status EnumEnvironmentStatus `json:"status"`
}

// NewEnvironmentStatus instantiates a new EnvironmentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentStatus(status EnumEnvironmentStatus) *EnvironmentStatus {
	this := EnvironmentStatus{}
	this.Status = status
	return &this
}

// NewEnvironmentStatusWithDefaults instantiates a new EnvironmentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentStatusWithDefaults() *EnvironmentStatus {
	this := EnvironmentStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *EnvironmentStatus) GetStatus() EnumEnvironmentStatus {
	if o == nil {
		var ret EnumEnvironmentStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *EnvironmentStatus) GetStatusOk() (*EnumEnvironmentStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *EnvironmentStatus) SetStatus(v EnumEnvironmentStatus) {
	o.Status = v
}

func (o EnvironmentStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableEnvironmentStatus struct {
	value *EnvironmentStatus
	isSet bool
}

func (v NullableEnvironmentStatus) Get() *EnvironmentStatus {
	return v.value
}

func (v *NullableEnvironmentStatus) Set(val *EnvironmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentStatus(val *EnvironmentStatus) *NullableEnvironmentStatus {
	return &NullableEnvironmentStatus{value: val, isSet: true}
}

func (v NullableEnvironmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
