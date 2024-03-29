/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package model

import (
	"encoding/json"
)

// P1ErrorDetailsInnerInnerError Additional details to help the client developer resolve the fault (primarily for UI validation reasons).
type P1ErrorDetailsInnerInnerError struct {
	// Errors that failed due to range violation. This attribute represents the minimum value of the range.
	RangeMinimumValue *int32 `json:"rangeMinimumValue,omitempty"`
	// The maximum range or value of an attribute.
	RangeMaximumValue *int32 `json:"rangeMaximumValue,omitempty"`
	// A regex pattern describing an acceptable input pattern.
	AllowedPattern *string `json:"allowedPattern,omitempty"`
	// A list describing acceptable values.
	AllowedValues []string `json:"allowedValues,omitempty"`
	// The maximum value allowed for the request.
	MaximumValue     *int32   `json:"maximumValue,omitempty"`
	ReferencedValues []string `json:"referencedValues,omitempty"`
}

// NewP1ErrorDetailsInnerInnerError instantiates a new P1ErrorDetailsInnerInnerError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewP1ErrorDetailsInnerInnerError() *P1ErrorDetailsInnerInnerError {
	this := P1ErrorDetailsInnerInnerError{}
	return &this
}

// NewP1ErrorDetailsInnerInnerErrorWithDefaults instantiates a new P1ErrorDetailsInnerInnerError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewP1ErrorDetailsInnerInnerErrorWithDefaults() *P1ErrorDetailsInnerInnerError {
	this := P1ErrorDetailsInnerInnerError{}
	return &this
}

// GetRangeMinimumValue returns the RangeMinimumValue field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetRangeMinimumValue() int32 {
	if o == nil || o.RangeMinimumValue == nil {
		var ret int32
		return ret
	}
	return *o.RangeMinimumValue
}

// GetRangeMinimumValueOk returns a tuple with the RangeMinimumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetRangeMinimumValueOk() (*int32, bool) {
	if o == nil || o.RangeMinimumValue == nil {
		return nil, false
	}
	return o.RangeMinimumValue, true
}

// HasRangeMinimumValue returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasRangeMinimumValue() bool {
	if o != nil && o.RangeMinimumValue != nil {
		return true
	}

	return false
}

// SetRangeMinimumValue gets a reference to the given int32 and assigns it to the RangeMinimumValue field.
func (o *P1ErrorDetailsInnerInnerError) SetRangeMinimumValue(v int32) {
	o.RangeMinimumValue = &v
}

// GetRangeMaximumValue returns the RangeMaximumValue field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetRangeMaximumValue() int32 {
	if o == nil || o.RangeMaximumValue == nil {
		var ret int32
		return ret
	}
	return *o.RangeMaximumValue
}

// GetRangeMaximumValueOk returns a tuple with the RangeMaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetRangeMaximumValueOk() (*int32, bool) {
	if o == nil || o.RangeMaximumValue == nil {
		return nil, false
	}
	return o.RangeMaximumValue, true
}

// HasRangeMaximumValue returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasRangeMaximumValue() bool {
	if o != nil && o.RangeMaximumValue != nil {
		return true
	}

	return false
}

// SetRangeMaximumValue gets a reference to the given int32 and assigns it to the RangeMaximumValue field.
func (o *P1ErrorDetailsInnerInnerError) SetRangeMaximumValue(v int32) {
	o.RangeMaximumValue = &v
}

// GetAllowedPattern returns the AllowedPattern field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetAllowedPattern() string {
	if o == nil || o.AllowedPattern == nil {
		var ret string
		return ret
	}
	return *o.AllowedPattern
}

// GetAllowedPatternOk returns a tuple with the AllowedPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetAllowedPatternOk() (*string, bool) {
	if o == nil || o.AllowedPattern == nil {
		return nil, false
	}
	return o.AllowedPattern, true
}

// HasAllowedPattern returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasAllowedPattern() bool {
	if o != nil && o.AllowedPattern != nil {
		return true
	}

	return false
}

// SetAllowedPattern gets a reference to the given string and assigns it to the AllowedPattern field.
func (o *P1ErrorDetailsInnerInnerError) SetAllowedPattern(v string) {
	o.AllowedPattern = &v
}

// GetAllowedValues returns the AllowedValues field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetAllowedValues() []string {
	if o == nil || o.AllowedValues == nil {
		var ret []string
		return ret
	}
	return o.AllowedValues
}

// GetAllowedValuesOk returns a tuple with the AllowedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetAllowedValuesOk() ([]string, bool) {
	if o == nil || o.AllowedValues == nil {
		return nil, false
	}
	return o.AllowedValues, true
}

// HasAllowedValues returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasAllowedValues() bool {
	if o != nil && o.AllowedValues != nil {
		return true
	}

	return false
}

// SetAllowedValues gets a reference to the given []string and assigns it to the AllowedValues field.
func (o *P1ErrorDetailsInnerInnerError) SetAllowedValues(v []string) {
	o.AllowedValues = v
}

// GetMaximumValue returns the MaximumValue field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetMaximumValue() int32 {
	if o == nil || o.MaximumValue == nil {
		var ret int32
		return ret
	}
	return *o.MaximumValue
}

// GetMaximumValueOk returns a tuple with the MaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetMaximumValueOk() (*int32, bool) {
	if o == nil || o.MaximumValue == nil {
		return nil, false
	}
	return o.MaximumValue, true
}

// HasMaximumValue returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasMaximumValue() bool {
	if o != nil && o.MaximumValue != nil {
		return true
	}

	return false
}

// SetMaximumValue gets a reference to the given int32 and assigns it to the MaximumValue field.
func (o *P1ErrorDetailsInnerInnerError) SetMaximumValue(v int32) {
	o.MaximumValue = &v
}

// GetReferencedValues returns the ReferencedValues field value if set, zero value otherwise.
func (o *P1ErrorDetailsInnerInnerError) GetReferencedValues() []string {
	if o == nil || o.ReferencedValues == nil {
		var ret []string
		return ret
	}
	return o.ReferencedValues
}

// GetReferencedValuesOk returns a tuple with the ReferencedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *P1ErrorDetailsInnerInnerError) GetReferencedValuesOk() ([]string, bool) {
	if o == nil || o.ReferencedValues == nil {
		return nil, false
	}
	return o.ReferencedValues, true
}

// HasReferencedValues returns a boolean if a field has been set.
func (o *P1ErrorDetailsInnerInnerError) HasReferencedValues() bool {
	if o != nil && o.ReferencedValues != nil {
		return true
	}

	return false
}

// SetReferencedValues gets a reference to the given []string and assigns it to the ReferencedValues field.
func (o *P1ErrorDetailsInnerInnerError) SetReferencedValues(v []string) {
	o.ReferencedValues = v
}

func (o P1ErrorDetailsInnerInnerError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RangeMinimumValue != nil {
		toSerialize["rangeMinimumValue"] = o.RangeMinimumValue
	}
	if o.RangeMaximumValue != nil {
		toSerialize["rangeMaximumValue"] = o.RangeMaximumValue
	}
	if o.AllowedPattern != nil {
		toSerialize["allowedPattern"] = o.AllowedPattern
	}
	if o.AllowedValues != nil {
		toSerialize["allowedValues"] = o.AllowedValues
	}
	if o.MaximumValue != nil {
		toSerialize["maximumValue"] = o.MaximumValue
	}
	if o.ReferencedValues != nil {
		toSerialize["referencedValues"] = o.ReferencedValues
	}
	return json.Marshal(toSerialize)
}

type NullableP1ErrorDetailsInnerInnerError struct {
	value *P1ErrorDetailsInnerInnerError
	isSet bool
}

func (v NullableP1ErrorDetailsInnerInnerError) Get() *P1ErrorDetailsInnerInnerError {
	return v.value
}

func (v *NullableP1ErrorDetailsInnerInnerError) Set(val *P1ErrorDetailsInnerInnerError) {
	v.value = val
	v.isSet = true
}

func (v NullableP1ErrorDetailsInnerInnerError) IsSet() bool {
	return v.isSet
}

func (v *NullableP1ErrorDetailsInnerInnerError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableP1ErrorDetailsInnerInnerError(val *P1ErrorDetailsInnerInnerError) *NullableP1ErrorDetailsInnerInnerError {
	return &NullableP1ErrorDetailsInnerInnerError{value: val, isSet: true}
}

func (v NullableP1ErrorDetailsInnerInnerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableP1ErrorDetailsInnerInnerError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
