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

// checks if the MFAPushCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFAPushCredential{}

// MFAPushCredential struct for MFAPushCredential
type MFAPushCredential struct {
	Type EnumMFAPushCredentialAttrType `json:"type"`
}

type _MFAPushCredential MFAPushCredential

// NewMFAPushCredential instantiates a new MFAPushCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAPushCredential(type_ EnumMFAPushCredentialAttrType) *MFAPushCredential {
	this := MFAPushCredential{}
	this.Type = type_
	return &this
}

// NewMFAPushCredentialWithDefaults instantiates a new MFAPushCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAPushCredentialWithDefaults() *MFAPushCredential {
	this := MFAPushCredential{}
	return &this
}

// GetType returns the Type field value
func (o *MFAPushCredential) GetType() EnumMFAPushCredentialAttrType {
	if o == nil {
		var ret EnumMFAPushCredentialAttrType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MFAPushCredential) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MFAPushCredential) SetType(v EnumMFAPushCredentialAttrType) {
	o.Type = v
}

func (o MFAPushCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFAPushCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *MFAPushCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varMFAPushCredential := _MFAPushCredential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMFAPushCredential)

	if err != nil {
		return err
	}

	*o = MFAPushCredential(varMFAPushCredential)

	return err
}

type NullableMFAPushCredential struct {
	value *MFAPushCredential
	isSet bool
}

func (v NullableMFAPushCredential) Get() *MFAPushCredential {
	return v.value
}

func (v *NullableMFAPushCredential) Set(val *MFAPushCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredential(val *MFAPushCredential) *NullableMFAPushCredential {
	return &NullableMFAPushCredential{value: val, isSet: true}
}

func (v NullableMFAPushCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


