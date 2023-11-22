/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the ApplicationSAMLAllOfIdpSigningKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSAMLAllOfIdpSigningKey{}

// ApplicationSAMLAllOfIdpSigningKey struct for ApplicationSAMLAllOfIdpSigningKey
type ApplicationSAMLAllOfIdpSigningKey struct {
	// A string that specifies the certificate to be used by the identity provider to sign assertions and responses. If this property is omitted, the default signing certificate for the environment is used.
	Id string `json:"id"`
}

type _ApplicationSAMLAllOfIdpSigningKey ApplicationSAMLAllOfIdpSigningKey

// NewApplicationSAMLAllOfIdpSigningKey instantiates a new ApplicationSAMLAllOfIdpSigningKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAMLAllOfIdpSigningKey(id string) *ApplicationSAMLAllOfIdpSigningKey {
	this := ApplicationSAMLAllOfIdpSigningKey{}
	this.Id = id
	return &this
}

// NewApplicationSAMLAllOfIdpSigningKeyWithDefaults instantiates a new ApplicationSAMLAllOfIdpSigningKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLAllOfIdpSigningKeyWithDefaults() *ApplicationSAMLAllOfIdpSigningKey {
	this := ApplicationSAMLAllOfIdpSigningKey{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationSAMLAllOfIdpSigningKey) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOfIdpSigningKey) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationSAMLAllOfIdpSigningKey) SetId(v string) {
	o.Id = v
}

func (o ApplicationSAMLAllOfIdpSigningKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSAMLAllOfIdpSigningKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *ApplicationSAMLAllOfIdpSigningKey) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationSAMLAllOfIdpSigningKey := _ApplicationSAMLAllOfIdpSigningKey{}

	err = json.Unmarshal(bytes, &varApplicationSAMLAllOfIdpSigningKey)

	if err != nil {
		return err
	}

	*o = ApplicationSAMLAllOfIdpSigningKey(varApplicationSAMLAllOfIdpSigningKey)

	return err
}

type NullableApplicationSAMLAllOfIdpSigningKey struct {
	value *ApplicationSAMLAllOfIdpSigningKey
	isSet bool
}

func (v NullableApplicationSAMLAllOfIdpSigningKey) Get() *ApplicationSAMLAllOfIdpSigningKey {
	return v.value
}

func (v *NullableApplicationSAMLAllOfIdpSigningKey) Set(val *ApplicationSAMLAllOfIdpSigningKey) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAMLAllOfIdpSigningKey) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAMLAllOfIdpSigningKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAMLAllOfIdpSigningKey(val *ApplicationSAMLAllOfIdpSigningKey) *NullableApplicationSAMLAllOfIdpSigningKey {
	return &NullableApplicationSAMLAllOfIdpSigningKey{value: val, isSet: true}
}

func (v NullableApplicationSAMLAllOfIdpSigningKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAMLAllOfIdpSigningKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


