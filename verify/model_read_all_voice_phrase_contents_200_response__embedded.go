/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the ReadAllVoicePhraseContents200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadAllVoicePhraseContents200ResponseEmbedded{}

// ReadAllVoicePhraseContents200ResponseEmbedded struct for ReadAllVoicePhraseContents200ResponseEmbedded
type ReadAllVoicePhraseContents200ResponseEmbedded struct {
	Contents []VoicePhraseContents `json:"contents,omitempty"`
}

// NewReadAllVoicePhraseContents200ResponseEmbedded instantiates a new ReadAllVoicePhraseContents200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadAllVoicePhraseContents200ResponseEmbedded() *ReadAllVoicePhraseContents200ResponseEmbedded {
	this := ReadAllVoicePhraseContents200ResponseEmbedded{}
	return &this
}

// NewReadAllVoicePhraseContents200ResponseEmbeddedWithDefaults instantiates a new ReadAllVoicePhraseContents200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadAllVoicePhraseContents200ResponseEmbeddedWithDefaults() *ReadAllVoicePhraseContents200ResponseEmbedded {
	this := ReadAllVoicePhraseContents200ResponseEmbedded{}
	return &this
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *ReadAllVoicePhraseContents200ResponseEmbedded) GetContents() []VoicePhraseContents {
	if o == nil || IsNil(o.Contents) {
		var ret []VoicePhraseContents
		return ret
	}
	return o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadAllVoicePhraseContents200ResponseEmbedded) GetContentsOk() ([]VoicePhraseContents, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *ReadAllVoicePhraseContents200ResponseEmbedded) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given []VoicePhraseContents and assigns it to the Contents field.
func (o *ReadAllVoicePhraseContents200ResponseEmbedded) SetContents(v []VoicePhraseContents) {
	o.Contents = v
}

func (o ReadAllVoicePhraseContents200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadAllVoicePhraseContents200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return toSerialize, nil
}

type NullableReadAllVoicePhraseContents200ResponseEmbedded struct {
	value *ReadAllVoicePhraseContents200ResponseEmbedded
	isSet bool
}

func (v NullableReadAllVoicePhraseContents200ResponseEmbedded) Get() *ReadAllVoicePhraseContents200ResponseEmbedded {
	return v.value
}

func (v *NullableReadAllVoicePhraseContents200ResponseEmbedded) Set(val *ReadAllVoicePhraseContents200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableReadAllVoicePhraseContents200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableReadAllVoicePhraseContents200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadAllVoicePhraseContents200ResponseEmbedded(val *ReadAllVoicePhraseContents200ResponseEmbedded) *NullableReadAllVoicePhraseContents200ResponseEmbedded {
	return &NullableReadAllVoicePhraseContents200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableReadAllVoicePhraseContents200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadAllVoicePhraseContents200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


