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

// checks if the CredentialIssuanceRuleUsageDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleUsageDetails{}

// CredentialIssuanceRuleUsageDetails struct for CredentialIssuanceRuleUsageDetails
type CredentialIssuanceRuleUsageDetails struct {
	Issued  []CredentialIssuanceRuleUsageInner `json:"issued,omitempty"`
	Revoked []CredentialIssuanceRuleUsageInner `json:"revoked,omitempty"`
	Updated []CredentialIssuanceRuleUsageInner `json:"updated,omitempty"`
}

// NewCredentialIssuanceRuleUsageDetails instantiates a new CredentialIssuanceRuleUsageDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleUsageDetails() *CredentialIssuanceRuleUsageDetails {
	this := CredentialIssuanceRuleUsageDetails{}
	return &this
}

// NewCredentialIssuanceRuleUsageDetailsWithDefaults instantiates a new CredentialIssuanceRuleUsageDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleUsageDetailsWithDefaults() *CredentialIssuanceRuleUsageDetails {
	this := CredentialIssuanceRuleUsageDetails{}
	return &this
}

// GetIssued returns the Issued field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleUsageDetails) GetIssued() []CredentialIssuanceRuleUsageInner {
	if o == nil || IsNil(o.Issued) {
		var ret []CredentialIssuanceRuleUsageInner
		return ret
	}
	return o.Issued
}

// GetIssuedOk returns a tuple with the Issued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleUsageDetails) GetIssuedOk() ([]CredentialIssuanceRuleUsageInner, bool) {
	if o == nil || IsNil(o.Issued) {
		return nil, false
	}
	return o.Issued, true
}

// HasIssued returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleUsageDetails) HasIssued() bool {
	if o != nil && !IsNil(o.Issued) {
		return true
	}

	return false
}

// SetIssued gets a reference to the given []CredentialIssuanceRuleUsageInner and assigns it to the Issued field.
func (o *CredentialIssuanceRuleUsageDetails) SetIssued(v []CredentialIssuanceRuleUsageInner) {
	o.Issued = v
}

// GetRevoked returns the Revoked field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleUsageDetails) GetRevoked() []CredentialIssuanceRuleUsageInner {
	if o == nil || IsNil(o.Revoked) {
		var ret []CredentialIssuanceRuleUsageInner
		return ret
	}
	return o.Revoked
}

// GetRevokedOk returns a tuple with the Revoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleUsageDetails) GetRevokedOk() ([]CredentialIssuanceRuleUsageInner, bool) {
	if o == nil || IsNil(o.Revoked) {
		return nil, false
	}
	return o.Revoked, true
}

// HasRevoked returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleUsageDetails) HasRevoked() bool {
	if o != nil && !IsNil(o.Revoked) {
		return true
	}

	return false
}

// SetRevoked gets a reference to the given []CredentialIssuanceRuleUsageInner and assigns it to the Revoked field.
func (o *CredentialIssuanceRuleUsageDetails) SetRevoked(v []CredentialIssuanceRuleUsageInner) {
	o.Revoked = v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleUsageDetails) GetUpdated() []CredentialIssuanceRuleUsageInner {
	if o == nil || IsNil(o.Updated) {
		var ret []CredentialIssuanceRuleUsageInner
		return ret
	}
	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleUsageDetails) GetUpdatedOk() ([]CredentialIssuanceRuleUsageInner, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleUsageDetails) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given []CredentialIssuanceRuleUsageInner and assigns it to the Updated field.
func (o *CredentialIssuanceRuleUsageDetails) SetUpdated(v []CredentialIssuanceRuleUsageInner) {
	o.Updated = v
}

func (o CredentialIssuanceRuleUsageDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleUsageDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issued) {
		toSerialize["issued"] = o.Issued
	}
	if !IsNil(o.Revoked) {
		toSerialize["revoked"] = o.Revoked
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleUsageDetails struct {
	value *CredentialIssuanceRuleUsageDetails
	isSet bool
}

func (v NullableCredentialIssuanceRuleUsageDetails) Get() *CredentialIssuanceRuleUsageDetails {
	return v.value
}

func (v *NullableCredentialIssuanceRuleUsageDetails) Set(val *CredentialIssuanceRuleUsageDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleUsageDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleUsageDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleUsageDetails(val *CredentialIssuanceRuleUsageDetails) *NullableCredentialIssuanceRuleUsageDetails {
	return &NullableCredentialIssuanceRuleUsageDetails{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleUsageDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleUsageDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
