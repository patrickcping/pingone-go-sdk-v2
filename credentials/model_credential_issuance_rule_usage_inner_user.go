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

// checks if the CredentialIssuanceRuleUsageInnerUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleUsageInnerUser{}

// CredentialIssuanceRuleUsageInnerUser struct for CredentialIssuanceRuleUsageInnerUser
type CredentialIssuanceRuleUsageInnerUser struct {
	// A string representing the identifier (UUID) of the user identified by the filter on the credential issuance rule.
	Id string `json:"id"`
}

// NewCredentialIssuanceRuleUsageInnerUser instantiates a new CredentialIssuanceRuleUsageInnerUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleUsageInnerUser(id string) *CredentialIssuanceRuleUsageInnerUser {
	this := CredentialIssuanceRuleUsageInnerUser{}
	this.Id = id
	return &this
}

// NewCredentialIssuanceRuleUsageInnerUserWithDefaults instantiates a new CredentialIssuanceRuleUsageInnerUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleUsageInnerUserWithDefaults() *CredentialIssuanceRuleUsageInnerUser {
	this := CredentialIssuanceRuleUsageInnerUser{}
	return &this
}

// GetId returns the Id field value
func (o *CredentialIssuanceRuleUsageInnerUser) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleUsageInnerUser) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CredentialIssuanceRuleUsageInnerUser) SetId(v string) {
	o.Id = v
}

func (o CredentialIssuanceRuleUsageInnerUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleUsageInnerUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleUsageInnerUser struct {
	value *CredentialIssuanceRuleUsageInnerUser
	isSet bool
}

func (v NullableCredentialIssuanceRuleUsageInnerUser) Get() *CredentialIssuanceRuleUsageInnerUser {
	return v.value
}

func (v *NullableCredentialIssuanceRuleUsageInnerUser) Set(val *CredentialIssuanceRuleUsageInnerUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleUsageInnerUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleUsageInnerUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleUsageInnerUser(val *CredentialIssuanceRuleUsageInnerUser) *NullableCredentialIssuanceRuleUsageInnerUser {
	return &NullableCredentialIssuanceRuleUsageInnerUser{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleUsageInnerUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleUsageInnerUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


