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

// checks if the EntityArrayEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityArrayEmbedded{}

// EntityArrayEmbedded struct for EntityArrayEmbedded
type EntityArrayEmbedded struct {
	Items []EntityArrayEmbeddedItemsInner `json:"items,omitempty"`
	IssuanceRules []CredentialIssuanceRule `json:"issuanceRules,omitempty"`
	StagedChanges []CredentialIssuanceRuleStagedChange `json:"stagedChanges,omitempty"`
	DigitalWalletApplications []DigitalWalletApplication `json:"digitalWalletApplications,omitempty"`
	DigitalWallets []CredentialDigitalWallet `json:"digitalWallets,omitempty"`
	ProvisionedCredentials []ProvisionedCredential `json:"provisionedCredentials,omitempty"`
}

// NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityArrayEmbedded() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded {
	this := EntityArrayEmbedded{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetItems() []EntityArrayEmbeddedItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []EntityArrayEmbeddedItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetItemsOk() ([]EntityArrayEmbeddedItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EntityArrayEmbeddedItemsInner and assigns it to the Items field.
func (o *EntityArrayEmbedded) SetItems(v []EntityArrayEmbeddedItemsInner) {
	o.Items = v
}

// GetIssuanceRules returns the IssuanceRules field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetIssuanceRules() []CredentialIssuanceRule {
	if o == nil || IsNil(o.IssuanceRules) {
		var ret []CredentialIssuanceRule
		return ret
	}
	return o.IssuanceRules
}

// GetIssuanceRulesOk returns a tuple with the IssuanceRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetIssuanceRulesOk() ([]CredentialIssuanceRule, bool) {
	if o == nil || IsNil(o.IssuanceRules) {
		return nil, false
	}
	return o.IssuanceRules, true
}

// HasIssuanceRules returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasIssuanceRules() bool {
	if o != nil && !IsNil(o.IssuanceRules) {
		return true
	}

	return false
}

// SetIssuanceRules gets a reference to the given []CredentialIssuanceRule and assigns it to the IssuanceRules field.
func (o *EntityArrayEmbedded) SetIssuanceRules(v []CredentialIssuanceRule) {
	o.IssuanceRules = v
}

// GetStagedChanges returns the StagedChanges field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetStagedChanges() []CredentialIssuanceRuleStagedChange {
	if o == nil || IsNil(o.StagedChanges) {
		var ret []CredentialIssuanceRuleStagedChange
		return ret
	}
	return o.StagedChanges
}

// GetStagedChangesOk returns a tuple with the StagedChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetStagedChangesOk() ([]CredentialIssuanceRuleStagedChange, bool) {
	if o == nil || IsNil(o.StagedChanges) {
		return nil, false
	}
	return o.StagedChanges, true
}

// HasStagedChanges returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasStagedChanges() bool {
	if o != nil && !IsNil(o.StagedChanges) {
		return true
	}

	return false
}

// SetStagedChanges gets a reference to the given []CredentialIssuanceRuleStagedChange and assigns it to the StagedChanges field.
func (o *EntityArrayEmbedded) SetStagedChanges(v []CredentialIssuanceRuleStagedChange) {
	o.StagedChanges = v
}

// GetDigitalWalletApplications returns the DigitalWalletApplications field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetDigitalWalletApplications() []DigitalWalletApplication {
	if o == nil || IsNil(o.DigitalWalletApplications) {
		var ret []DigitalWalletApplication
		return ret
	}
	return o.DigitalWalletApplications
}

// GetDigitalWalletApplicationsOk returns a tuple with the DigitalWalletApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetDigitalWalletApplicationsOk() ([]DigitalWalletApplication, bool) {
	if o == nil || IsNil(o.DigitalWalletApplications) {
		return nil, false
	}
	return o.DigitalWalletApplications, true
}

// HasDigitalWalletApplications returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasDigitalWalletApplications() bool {
	if o != nil && !IsNil(o.DigitalWalletApplications) {
		return true
	}

	return false
}

// SetDigitalWalletApplications gets a reference to the given []DigitalWalletApplication and assigns it to the DigitalWalletApplications field.
func (o *EntityArrayEmbedded) SetDigitalWalletApplications(v []DigitalWalletApplication) {
	o.DigitalWalletApplications = v
}

// GetDigitalWallets returns the DigitalWallets field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetDigitalWallets() []CredentialDigitalWallet {
	if o == nil || IsNil(o.DigitalWallets) {
		var ret []CredentialDigitalWallet
		return ret
	}
	return o.DigitalWallets
}

// GetDigitalWalletsOk returns a tuple with the DigitalWallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetDigitalWalletsOk() ([]CredentialDigitalWallet, bool) {
	if o == nil || IsNil(o.DigitalWallets) {
		return nil, false
	}
	return o.DigitalWallets, true
}

// HasDigitalWallets returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasDigitalWallets() bool {
	if o != nil && !IsNil(o.DigitalWallets) {
		return true
	}

	return false
}

// SetDigitalWallets gets a reference to the given []CredentialDigitalWallet and assigns it to the DigitalWallets field.
func (o *EntityArrayEmbedded) SetDigitalWallets(v []CredentialDigitalWallet) {
	o.DigitalWallets = v
}

// GetProvisionedCredentials returns the ProvisionedCredentials field value if set, zero value otherwise.
func (o *EntityArrayEmbedded) GetProvisionedCredentials() []ProvisionedCredential {
	if o == nil || IsNil(o.ProvisionedCredentials) {
		var ret []ProvisionedCredential
		return ret
	}
	return o.ProvisionedCredentials
}

// GetProvisionedCredentialsOk returns a tuple with the ProvisionedCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityArrayEmbedded) GetProvisionedCredentialsOk() ([]ProvisionedCredential, bool) {
	if o == nil || IsNil(o.ProvisionedCredentials) {
		return nil, false
	}
	return o.ProvisionedCredentials, true
}

// HasProvisionedCredentials returns a boolean if a field has been set.
func (o *EntityArrayEmbedded) HasProvisionedCredentials() bool {
	if o != nil && !IsNil(o.ProvisionedCredentials) {
		return true
	}

	return false
}

// SetProvisionedCredentials gets a reference to the given []ProvisionedCredential and assigns it to the ProvisionedCredentials field.
func (o *EntityArrayEmbedded) SetProvisionedCredentials(v []ProvisionedCredential) {
	o.ProvisionedCredentials = v
}

func (o EntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityArrayEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.IssuanceRules) {
		toSerialize["issuanceRules"] = o.IssuanceRules
	}
	if !IsNil(o.StagedChanges) {
		toSerialize["stagedChanges"] = o.StagedChanges
	}
	if !IsNil(o.DigitalWalletApplications) {
		toSerialize["digitalWalletApplications"] = o.DigitalWalletApplications
	}
	if !IsNil(o.DigitalWallets) {
		toSerialize["digitalWallets"] = o.DigitalWallets
	}
	if !IsNil(o.ProvisionedCredentials) {
		toSerialize["provisionedCredentials"] = o.ProvisionedCredentials
	}
	return toSerialize, nil
}

type NullableEntityArrayEmbedded struct {
	value *EntityArrayEmbedded
	isSet bool
}

func (v NullableEntityArrayEmbedded) Get() *EntityArrayEmbedded {
	return v.value
}

func (v *NullableEntityArrayEmbedded) Set(val *EntityArrayEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbedded(val *EntityArrayEmbedded) *NullableEntityArrayEmbedded {
	return &NullableEntityArrayEmbedded{value: val, isSet: true}
}

func (v NullableEntityArrayEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


