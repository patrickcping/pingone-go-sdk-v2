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

// checks if the FIDO2PolicyUserDisplayNameAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDO2PolicyUserDisplayNameAttributes{}

// FIDO2PolicyUserDisplayNameAttributes Used to specify the string associated with the users's account that is displayed during registration and authentication.
type FIDO2PolicyUserDisplayNameAttributes struct {
	// List of strings associated with the users's account that can be displayed during registration and authentication. Each object in the array is a name:value pair where the first part is \"name\" and the second is the name of the user attribute, for example, `{\"name\": \"username\"}`, `{\"name\": \"email\"}`. If you want to use the \"name\" attribute for the user, you must also specify the \"subAttributes\", which can be either the \"given\" and \"family\" user attributes or the \"formatted\" user attribute. For example, `{\"name\": “name”, “subAttributes”:[{“name”: “given”}, {“name”: “family”}]}, {\"name\": \"email\"}` or `{\"name\": “name”, “subAttributes”:[{“name”: “formatted”}]}, {\"name\": \"email\"}`. - The content of the list should reflect the preferred order. - If the first attribute is empty for the user, PingOne will continue through the list until a non-empty attribute is found. - You can specify any user attribute (including custom attributes) that meet the following criteria: attribute type must be String, validation cannot be set to enumerated values. - The array must contain the user attribute \"username\" - to ensure that there is at least one non-empty attribute. - You can have a maximum of six user attributes in the list. 
	Attributes []FIDO2PolicyUserDisplayNameAttributesAttributesInner `json:"attributes"`
}

type _FIDO2PolicyUserDisplayNameAttributes FIDO2PolicyUserDisplayNameAttributes

// NewFIDO2PolicyUserDisplayNameAttributes instantiates a new FIDO2PolicyUserDisplayNameAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDO2PolicyUserDisplayNameAttributes(attributes []FIDO2PolicyUserDisplayNameAttributesAttributesInner) *FIDO2PolicyUserDisplayNameAttributes {
	this := FIDO2PolicyUserDisplayNameAttributes{}
	this.Attributes = attributes
	return &this
}

// NewFIDO2PolicyUserDisplayNameAttributesWithDefaults instantiates a new FIDO2PolicyUserDisplayNameAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDO2PolicyUserDisplayNameAttributesWithDefaults() *FIDO2PolicyUserDisplayNameAttributes {
	this := FIDO2PolicyUserDisplayNameAttributes{}
	return &this
}

// GetAttributes returns the Attributes field value
func (o *FIDO2PolicyUserDisplayNameAttributes) GetAttributes() []FIDO2PolicyUserDisplayNameAttributesAttributesInner {
	if o == nil {
		var ret []FIDO2PolicyUserDisplayNameAttributesAttributesInner
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyUserDisplayNameAttributes) GetAttributesOk() ([]FIDO2PolicyUserDisplayNameAttributesAttributesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *FIDO2PolicyUserDisplayNameAttributes) SetAttributes(v []FIDO2PolicyUserDisplayNameAttributesAttributesInner) {
	o.Attributes = v
}

func (o FIDO2PolicyUserDisplayNameAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDO2PolicyUserDisplayNameAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *FIDO2PolicyUserDisplayNameAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributes",
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

	varFIDO2PolicyUserDisplayNameAttributes := _FIDO2PolicyUserDisplayNameAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFIDO2PolicyUserDisplayNameAttributes)

	if err != nil {
		return err
	}

	*o = FIDO2PolicyUserDisplayNameAttributes(varFIDO2PolicyUserDisplayNameAttributes)

	return err
}

type NullableFIDO2PolicyUserDisplayNameAttributes struct {
	value *FIDO2PolicyUserDisplayNameAttributes
	isSet bool
}

func (v NullableFIDO2PolicyUserDisplayNameAttributes) Get() *FIDO2PolicyUserDisplayNameAttributes {
	return v.value
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributes) Set(val *FIDO2PolicyUserDisplayNameAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDO2PolicyUserDisplayNameAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDO2PolicyUserDisplayNameAttributes(val *FIDO2PolicyUserDisplayNameAttributes) *NullableFIDO2PolicyUserDisplayNameAttributes {
	return &NullableFIDO2PolicyUserDisplayNameAttributes{value: val, isSet: true}
}

func (v NullableFIDO2PolicyUserDisplayNameAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDO2PolicyUserDisplayNameAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


