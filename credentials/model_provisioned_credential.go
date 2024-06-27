/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the ProvisionedCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionedCredential{}

// ProvisionedCredential struct for ProvisionedCredential
type ProvisionedCredential struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	ClaimReference *ProvisionedCredentialClaimReference `json:"claimReference,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Credential *ProvisionedCredentialCredential `json:"credential,omitempty"`
	DigitalWallet *ProvisionedCredentialCredential `json:"digitalWallet,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the date that the provisioned credential expires. If this value is null, the provisioned credential never expires.
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
	Id *string `json:"id,omitempty"`
	// A string that specifies the tatus of the provisioned credential.
	Status *string `json:"status,omitempty"`
	User *ProvisionedCredentialUser `json:"user,omitempty"`
	WalletActions []ProvisionedCredentialWalletActionsInner `json:"walletActions,omitempty"`
}

// NewProvisionedCredential instantiates a new ProvisionedCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionedCredential() *ProvisionedCredential {
	this := ProvisionedCredential{}
	return &this
}

// NewProvisionedCredentialWithDefaults instantiates a new ProvisionedCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionedCredentialWithDefaults() *ProvisionedCredential {
	this := ProvisionedCredential{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *ProvisionedCredential) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetClaimReference returns the ClaimReference field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetClaimReference() ProvisionedCredentialClaimReference {
	if o == nil || IsNil(o.ClaimReference) {
		var ret ProvisionedCredentialClaimReference
		return ret
	}
	return *o.ClaimReference
}

// GetClaimReferenceOk returns a tuple with the ClaimReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetClaimReferenceOk() (*ProvisionedCredentialClaimReference, bool) {
	if o == nil || IsNil(o.ClaimReference) {
		return nil, false
	}
	return o.ClaimReference, true
}

// HasClaimReference returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasClaimReference() bool {
	if o != nil && !IsNil(o.ClaimReference) {
		return true
	}

	return false
}

// SetClaimReference gets a reference to the given ProvisionedCredentialClaimReference and assigns it to the ClaimReference field.
func (o *ProvisionedCredential) SetClaimReference(v ProvisionedCredentialClaimReference) {
	o.ClaimReference = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ProvisionedCredential) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredential returns the Credential field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetCredential() ProvisionedCredentialCredential {
	if o == nil || IsNil(o.Credential) {
		var ret ProvisionedCredentialCredential
		return ret
	}
	return *o.Credential
}

// GetCredentialOk returns a tuple with the Credential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetCredentialOk() (*ProvisionedCredentialCredential, bool) {
	if o == nil || IsNil(o.Credential) {
		return nil, false
	}
	return o.Credential, true
}

// HasCredential returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasCredential() bool {
	if o != nil && !IsNil(o.Credential) {
		return true
	}

	return false
}

// SetCredential gets a reference to the given ProvisionedCredentialCredential and assigns it to the Credential field.
func (o *ProvisionedCredential) SetCredential(v ProvisionedCredentialCredential) {
	o.Credential = &v
}

// GetDigitalWallet returns the DigitalWallet field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetDigitalWallet() ProvisionedCredentialCredential {
	if o == nil || IsNil(o.DigitalWallet) {
		var ret ProvisionedCredentialCredential
		return ret
	}
	return *o.DigitalWallet
}

// GetDigitalWalletOk returns a tuple with the DigitalWallet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetDigitalWalletOk() (*ProvisionedCredentialCredential, bool) {
	if o == nil || IsNil(o.DigitalWallet) {
		return nil, false
	}
	return o.DigitalWallet, true
}

// HasDigitalWallet returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasDigitalWallet() bool {
	if o != nil && !IsNil(o.DigitalWallet) {
		return true
	}

	return false
}

// SetDigitalWallet gets a reference to the given ProvisionedCredentialCredential and assigns it to the DigitalWallet field.
func (o *ProvisionedCredential) SetDigitalWallet(v ProvisionedCredentialCredential) {
	o.DigitalWallet = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *ProvisionedCredential) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetExpiredAt() time.Time {
	if o == nil || IsNil(o.ExpiredAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiredAt
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetExpiredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiredAt) {
		return nil, false
	}
	return o.ExpiredAt, true
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasExpiredAt() bool {
	if o != nil && !IsNil(o.ExpiredAt) {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given time.Time and assigns it to the ExpiredAt field.
func (o *ProvisionedCredential) SetExpiredAt(v time.Time) {
	o.ExpiredAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ProvisionedCredential) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ProvisionedCredential) SetStatus(v string) {
	o.Status = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetUser() ProvisionedCredentialUser {
	if o == nil || IsNil(o.User) {
		var ret ProvisionedCredentialUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetUserOk() (*ProvisionedCredentialUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ProvisionedCredentialUser and assigns it to the User field.
func (o *ProvisionedCredential) SetUser(v ProvisionedCredentialUser) {
	o.User = &v
}

// GetWalletActions returns the WalletActions field value if set, zero value otherwise.
func (o *ProvisionedCredential) GetWalletActions() []ProvisionedCredentialWalletActionsInner {
	if o == nil || IsNil(o.WalletActions) {
		var ret []ProvisionedCredentialWalletActionsInner
		return ret
	}
	return o.WalletActions
}

// GetWalletActionsOk returns a tuple with the WalletActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionedCredential) GetWalletActionsOk() ([]ProvisionedCredentialWalletActionsInner, bool) {
	if o == nil || IsNil(o.WalletActions) {
		return nil, false
	}
	return o.WalletActions, true
}

// HasWalletActions returns a boolean if a field has been set.
func (o *ProvisionedCredential) HasWalletActions() bool {
	if o != nil && !IsNil(o.WalletActions) {
		return true
	}

	return false
}

// SetWalletActions gets a reference to the given []ProvisionedCredentialWalletActionsInner and assigns it to the WalletActions field.
func (o *ProvisionedCredential) SetWalletActions(v []ProvisionedCredentialWalletActionsInner) {
	o.WalletActions = v
}

func (o ProvisionedCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionedCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.ClaimReference) {
		toSerialize["claimReference"] = o.ClaimReference
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Credential) {
		toSerialize["credential"] = o.Credential
	}
	if !IsNil(o.DigitalWallet) {
		toSerialize["digitalWallet"] = o.DigitalWallet
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.ExpiredAt) {
		toSerialize["expiredAt"] = o.ExpiredAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.WalletActions) {
		toSerialize["walletActions"] = o.WalletActions
	}
	return toSerialize, nil
}

type NullableProvisionedCredential struct {
	value *ProvisionedCredential
	isSet bool
}

func (v NullableProvisionedCredential) Get() *ProvisionedCredential {
	return v.value
}

func (v *NullableProvisionedCredential) Set(val *ProvisionedCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionedCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionedCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionedCredential(val *ProvisionedCredential) *NullableProvisionedCredential {
	return &NullableProvisionedCredential{value: val, isSet: true}
}

func (v NullableProvisionedCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionedCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


