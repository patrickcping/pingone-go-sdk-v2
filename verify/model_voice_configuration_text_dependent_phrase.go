/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VoiceConfigurationTextDependentPhrase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoiceConfigurationTextDependentPhrase{}

// VoiceConfigurationTextDependentPhrase struct for VoiceConfigurationTextDependentPhrase
type VoiceConfigurationTextDependentPhrase struct {
	Id string `json:"id"`
}

type _VoiceConfigurationTextDependentPhrase VoiceConfigurationTextDependentPhrase

// NewVoiceConfigurationTextDependentPhrase instantiates a new VoiceConfigurationTextDependentPhrase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceConfigurationTextDependentPhrase(id string) *VoiceConfigurationTextDependentPhrase {
	this := VoiceConfigurationTextDependentPhrase{}
	this.Id = id
	return &this
}

// NewVoiceConfigurationTextDependentPhraseWithDefaults instantiates a new VoiceConfigurationTextDependentPhrase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceConfigurationTextDependentPhraseWithDefaults() *VoiceConfigurationTextDependentPhrase {
	this := VoiceConfigurationTextDependentPhrase{}
	return &this
}

// GetId returns the Id field value
func (o *VoiceConfigurationTextDependentPhrase) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VoiceConfigurationTextDependentPhrase) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VoiceConfigurationTextDependentPhrase) SetId(v string) {
	o.Id = v
}

func (o VoiceConfigurationTextDependentPhrase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceConfigurationTextDependentPhrase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *VoiceConfigurationTextDependentPhrase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varVoiceConfigurationTextDependentPhrase := _VoiceConfigurationTextDependentPhrase{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVoiceConfigurationTextDependentPhrase)

	if err != nil {
		return err
	}

	*o = VoiceConfigurationTextDependentPhrase(varVoiceConfigurationTextDependentPhrase)

	return err
}

type NullableVoiceConfigurationTextDependentPhrase struct {
	value *VoiceConfigurationTextDependentPhrase
	isSet bool
}

func (v NullableVoiceConfigurationTextDependentPhrase) Get() *VoiceConfigurationTextDependentPhrase {
	return v.value
}

func (v *NullableVoiceConfigurationTextDependentPhrase) Set(val *VoiceConfigurationTextDependentPhrase) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceConfigurationTextDependentPhrase) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceConfigurationTextDependentPhrase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceConfigurationTextDependentPhrase(val *VoiceConfigurationTextDependentPhrase) *NullableVoiceConfigurationTextDependentPhrase {
	return &NullableVoiceConfigurationTextDependentPhrase{value: val, isSet: true}
}

func (v NullableVoiceConfigurationTextDependentPhrase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceConfigurationTextDependentPhrase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


