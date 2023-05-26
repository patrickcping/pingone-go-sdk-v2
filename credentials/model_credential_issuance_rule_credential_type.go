/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the CredentialIssuanceRuleCredentialType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleCredentialType{}

// CredentialIssuanceRuleCredentialType struct for CredentialIssuanceRuleCredentialType
type CredentialIssuanceRuleCredentialType struct {
	// A string that specifies the iIdentifier (UUID) of the credential type with which this credential issuance rule is associated.
	Id *string `json:"id,omitempty"`
}

// NewCredentialIssuanceRuleCredentialType instantiates a new CredentialIssuanceRuleCredentialType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleCredentialType() *CredentialIssuanceRuleCredentialType {
	this := CredentialIssuanceRuleCredentialType{}
	return &this
}

// NewCredentialIssuanceRuleCredentialTypeWithDefaults instantiates a new CredentialIssuanceRuleCredentialType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleCredentialTypeWithDefaults() *CredentialIssuanceRuleCredentialType {
	this := CredentialIssuanceRuleCredentialType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleCredentialType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleCredentialType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleCredentialType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialIssuanceRuleCredentialType) SetId(v string) {
	o.Id = &v
}

func (o CredentialIssuanceRuleCredentialType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleCredentialType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleCredentialType struct {
	value *CredentialIssuanceRuleCredentialType
	isSet bool
}

func (v NullableCredentialIssuanceRuleCredentialType) Get() *CredentialIssuanceRuleCredentialType {
	return v.value
}

func (v *NullableCredentialIssuanceRuleCredentialType) Set(val *CredentialIssuanceRuleCredentialType) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleCredentialType) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleCredentialType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleCredentialType(val *CredentialIssuanceRuleCredentialType) *NullableCredentialIssuanceRuleCredentialType {
	return &NullableCredentialIssuanceRuleCredentialType{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleCredentialType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleCredentialType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

