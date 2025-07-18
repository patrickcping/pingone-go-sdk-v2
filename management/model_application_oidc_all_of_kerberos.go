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

// checks if the ApplicationOIDCAllOfKerberos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOIDCAllOfKerberos{}

// ApplicationOIDCAllOfKerberos Object containing Kerberos settings
type ApplicationOIDCAllOfKerberos struct {
	Key ApplicationOIDCAllOfKerberosKey `json:"key"`
}

// NewApplicationOIDCAllOfKerberos instantiates a new ApplicationOIDCAllOfKerberos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOfKerberos(key ApplicationOIDCAllOfKerberosKey) *ApplicationOIDCAllOfKerberos {
	this := ApplicationOIDCAllOfKerberos{}
	this.Key = key
	return &this
}

// NewApplicationOIDCAllOfKerberosWithDefaults instantiates a new ApplicationOIDCAllOfKerberos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfKerberosWithDefaults() *ApplicationOIDCAllOfKerberos {
	this := ApplicationOIDCAllOfKerberos{}
	return &this
}

// GetKey returns the Key field value
func (o *ApplicationOIDCAllOfKerberos) GetKey() ApplicationOIDCAllOfKerberosKey {
	if o == nil {
		var ret ApplicationOIDCAllOfKerberosKey
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfKerberos) GetKeyOk() (*ApplicationOIDCAllOfKerberosKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApplicationOIDCAllOfKerberos) SetKey(v ApplicationOIDCAllOfKerberosKey) {
	o.Key = v
}

func (o ApplicationOIDCAllOfKerberos) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOIDCAllOfKerberos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableApplicationOIDCAllOfKerberos struct {
	value *ApplicationOIDCAllOfKerberos
	isSet bool
}

func (v NullableApplicationOIDCAllOfKerberos) Get() *ApplicationOIDCAllOfKerberos {
	return v.value
}

func (v *NullableApplicationOIDCAllOfKerberos) Set(val *ApplicationOIDCAllOfKerberos) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOfKerberos) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOfKerberos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOfKerberos(val *ApplicationOIDCAllOfKerberos) *NullableApplicationOIDCAllOfKerberos {
	return &NullableApplicationOIDCAllOfKerberos{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOfKerberos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOfKerberos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
