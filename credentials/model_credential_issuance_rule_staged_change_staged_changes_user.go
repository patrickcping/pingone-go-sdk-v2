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

// checks if the CredentialIssuanceRuleStagedChangeStagedChangesUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleStagedChangeStagedChangesUser{}

// CredentialIssuanceRuleStagedChangeStagedChangesUser struct for CredentialIssuanceRuleStagedChangeStagedChangesUser
type CredentialIssuanceRuleStagedChangeStagedChangesUser struct {
	// A string that specifies the identifier (UUID) of the user identified by the filter on the credential issuance rule.
	Id string `json:"id"`
}

// NewCredentialIssuanceRuleStagedChangeStagedChangesUser instantiates a new CredentialIssuanceRuleStagedChangeStagedChangesUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleStagedChangeStagedChangesUser(id string) *CredentialIssuanceRuleStagedChangeStagedChangesUser {
	this := CredentialIssuanceRuleStagedChangeStagedChangesUser{}
	this.Id = id
	return &this
}

// NewCredentialIssuanceRuleStagedChangeStagedChangesUserWithDefaults instantiates a new CredentialIssuanceRuleStagedChangeStagedChangesUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleStagedChangeStagedChangesUserWithDefaults() *CredentialIssuanceRuleStagedChangeStagedChangesUser {
	this := CredentialIssuanceRuleStagedChangeStagedChangesUser{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialIssuanceRuleStagedChangeStagedChangesUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChangesUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialIssuanceRuleStagedChangeStagedChangesUser) SetId(v string) {
	o.Id = v
}

func (o CredentialIssuanceRuleStagedChangeStagedChangesUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleStagedChangeStagedChangesUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleStagedChangeStagedChangesUser struct {
	value *CredentialIssuanceRuleStagedChangeStagedChangesUser
	isSet bool
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) Get() *CredentialIssuanceRuleStagedChangeStagedChangesUser {
	return v.value
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) Set(val *CredentialIssuanceRuleStagedChangeStagedChangesUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleStagedChangeStagedChangesUser(val *CredentialIssuanceRuleStagedChangeStagedChangesUser) *NullableCredentialIssuanceRuleStagedChangeStagedChangesUser {
	return &NullableCredentialIssuanceRuleStagedChangeStagedChangesUser{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChangesUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
