/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ImageTargets struct for ImageTargets
type ImageTargets struct {
	Original *ImageTargetsOriginal `json:"original,omitempty"`
}

// NewImageTargets instantiates a new ImageTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageTargets() *ImageTargets {
	this := ImageTargets{}
	return &this
}

// NewImageTargetsWithDefaults instantiates a new ImageTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageTargetsWithDefaults() *ImageTargets {
	this := ImageTargets{}
	return &this
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *ImageTargets) GetOriginal() ImageTargetsOriginal {
	if o == nil || o.Original == nil {
		var ret ImageTargetsOriginal
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageTargets) GetOriginalOk() (*ImageTargetsOriginal, bool) {
	if o == nil || o.Original == nil {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *ImageTargets) HasOriginal() bool {
	if o != nil && o.Original != nil {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given ImageTargetsOriginal and assigns it to the Original field.
func (o *ImageTargets) SetOriginal(v ImageTargetsOriginal) {
	o.Original = &v
}

func (o ImageTargets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Original != nil {
		toSerialize["original"] = o.Original
	}
	return json.Marshal(toSerialize)
}

type NullableImageTargets struct {
	value *ImageTargets
	isSet bool
}

func (v NullableImageTargets) Get() *ImageTargets {
	return v.value
}

func (v *NullableImageTargets) Set(val *ImageTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableImageTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableImageTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageTargets(val *ImageTargets) *NullableImageTargets {
	return &NullableImageTargets{value: val, isSet: true}
}

func (v NullableImageTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


