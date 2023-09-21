/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the ObjectIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectIssuer{}

// ObjectIssuer struct for ObjectIssuer
type ObjectIssuer struct {
	// A string that specifies the identifier (UUID) of the credential issuer.
	Id *string `json:"id,omitempty"`
}

// NewObjectIssuer instantiates a new ObjectIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectIssuer() *ObjectIssuer {
	this := ObjectIssuer{}
	return &this
}

// NewObjectIssuerWithDefaults instantiates a new ObjectIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIssuerWithDefaults() *ObjectIssuer {
	this := ObjectIssuer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObjectIssuer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectIssuer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObjectIssuer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObjectIssuer) SetId(v string) {
	o.Id = &v
}

func (o ObjectIssuer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableObjectIssuer struct {
	value *ObjectIssuer
	isSet bool
}

func (v NullableObjectIssuer) Get() *ObjectIssuer {
	return v.value
}

func (v *NullableObjectIssuer) Set(val *ObjectIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectIssuer(val *ObjectIssuer) *NullableObjectIssuer {
	return &NullableObjectIssuer{value: val, isSet: true}
}

func (v NullableObjectIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


