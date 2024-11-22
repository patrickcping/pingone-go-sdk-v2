/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FIDO2PolicyUserDisplayNameAttributesAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDO2PolicyUserDisplayNameAttributesAttributesInner{}

// FIDO2PolicyUserDisplayNameAttributesAttributesInner struct for FIDO2PolicyUserDisplayNameAttributesAttributesInner
type FIDO2PolicyUserDisplayNameAttributesAttributesInner struct {
	// The name of the attribute to use for the user display name.
	Name string `json:"name"`
	// The sub-attributes to use for the user display name.
	SubAttributes []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner `json:"subAttributes,omitempty"`
}

type _FIDO2PolicyUserDisplayNameAttributesAttributesInner FIDO2PolicyUserDisplayNameAttributesAttributesInner

// NewFIDO2PolicyUserDisplayNameAttributesAttributesInner instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDO2PolicyUserDisplayNameAttributesAttributesInner(name string) *FIDO2PolicyUserDisplayNameAttributesAttributesInner {
	this := FIDO2PolicyUserDisplayNameAttributesAttributesInner{}
	this.Name = name
	return &this
}

// NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerWithDefaults instantiates a new FIDO2PolicyUserDisplayNameAttributesAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDO2PolicyUserDisplayNameAttributesAttributesInnerWithDefaults() *FIDO2PolicyUserDisplayNameAttributesAttributesInner {
	this := FIDO2PolicyUserDisplayNameAttributesAttributesInner{}
	return &this
}

// GetName returns the Name field value
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) SetName(v string) {
	o.Name = v
}

// GetSubAttributes returns the SubAttributes field value if set, zero value otherwise.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetSubAttributes() []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner {
	if o == nil || IsNil(o.SubAttributes) {
		var ret []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner
		return ret
	}
	return o.SubAttributes
}

// GetSubAttributesOk returns a tuple with the SubAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) GetSubAttributesOk() ([]FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner, bool) {
	if o == nil || IsNil(o.SubAttributes) {
		return nil, false
	}
	return o.SubAttributes, true
}

// HasSubAttributes returns a boolean if a field has been set.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) HasSubAttributes() bool {
	if o != nil && !IsNil(o.SubAttributes) {
		return true
	}

	return false
}

// SetSubAttributes gets a reference to the given []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner and assigns it to the SubAttributes field.
func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) SetSubAttributes(v []FIDO2PolicyUserDisplayNameAttributesAttributesInnerSubAttributesInner) {
	o.SubAttributes = v
}

func (o FIDO2PolicyUserDisplayNameAttributesAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDO2PolicyUserDisplayNameAttributesAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.SubAttributes) {
		toSerialize["subAttributes"] = o.SubAttributes
	}
	return toSerialize, nil
}

func (o *FIDO2PolicyUserDisplayNameAttributesAttributesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varFIDO2PolicyUserDisplayNameAttributesAttributesInner := _FIDO2PolicyUserDisplayNameAttributesAttributesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFIDO2PolicyUserDisplayNameAttributesAttributesInner)

	if err != nil {
		return err
	}

	*o = FIDO2PolicyUserDisplayNameAttributesAttributesInner(varFIDO2PolicyUserDisplayNameAttributesAttributesInner)

	return err
}

type NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner struct {
	value *FIDO2PolicyUserDisplayNameAttributesAttributesInner
	isSet bool
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) Get() *FIDO2PolicyUserDisplayNameAttributesAttributesInner {
	return v.value
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) Set(val *FIDO2PolicyUserDisplayNameAttributesAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDO2PolicyUserDisplayNameAttributesAttributesInner(val *FIDO2PolicyUserDisplayNameAttributesAttributesInner) *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner {
	return &NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner{value: val, isSet: true}
}

func (v NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributesAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


