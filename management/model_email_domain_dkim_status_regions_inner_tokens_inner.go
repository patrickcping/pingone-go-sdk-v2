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

// EmailDomainDKIMStatusRegionsInnerTokensInner struct for EmailDomainDKIMStatusRegionsInnerTokensInner
type EmailDomainDKIMStatusRegionsInnerTokensInner struct {
	// Record name.
	Key *string `json:"key,omitempty"`
	// Record value.
	Value *string `json:"value,omitempty"`
}

// NewEmailDomainDKIMStatusRegionsInnerTokensInner instantiates a new EmailDomainDKIMStatusRegionsInnerTokensInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailDomainDKIMStatusRegionsInnerTokensInner() *EmailDomainDKIMStatusRegionsInnerTokensInner {
	this := EmailDomainDKIMStatusRegionsInnerTokensInner{}
	return &this
}

// NewEmailDomainDKIMStatusRegionsInnerTokensInnerWithDefaults instantiates a new EmailDomainDKIMStatusRegionsInnerTokensInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailDomainDKIMStatusRegionsInnerTokensInnerWithDefaults() *EmailDomainDKIMStatusRegionsInnerTokensInner {
	this := EmailDomainDKIMStatusRegionsInnerTokensInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EmailDomainDKIMStatusRegionsInnerTokensInner) SetValue(v string) {
	o.Value = &v
}

func (o EmailDomainDKIMStatusRegionsInnerTokensInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEmailDomainDKIMStatusRegionsInnerTokensInner struct {
	value *EmailDomainDKIMStatusRegionsInnerTokensInner
	isSet bool
}

func (v NullableEmailDomainDKIMStatusRegionsInnerTokensInner) Get() *EmailDomainDKIMStatusRegionsInnerTokensInner {
	return v.value
}

func (v *NullableEmailDomainDKIMStatusRegionsInnerTokensInner) Set(val *EmailDomainDKIMStatusRegionsInnerTokensInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailDomainDKIMStatusRegionsInnerTokensInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailDomainDKIMStatusRegionsInnerTokensInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailDomainDKIMStatusRegionsInnerTokensInner(val *EmailDomainDKIMStatusRegionsInnerTokensInner) *NullableEmailDomainDKIMStatusRegionsInnerTokensInner {
	return &NullableEmailDomainDKIMStatusRegionsInnerTokensInner{value: val, isSet: true}
}

func (v NullableEmailDomainDKIMStatusRegionsInnerTokensInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailDomainDKIMStatusRegionsInnerTokensInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

