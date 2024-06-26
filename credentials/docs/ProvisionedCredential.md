# ProvisionedCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**ClaimReference** | Pointer to [**ProvisionedCredentialClaimReference**](ProvisionedCredentialClaimReference.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Credential** | Pointer to [**ProvisionedCredentialCredential**](ProvisionedCredentialCredential.md) |  | [optional] 
**DigitalWallet** | Pointer to [**ProvisionedCredentialCredential**](ProvisionedCredentialCredential.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**ExpiredAt** | Pointer to **time.Time** | A string that specifies the date that the provisioned credential expires. If this value is null, the provisioned credential never expires. | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** | A string that specifies the tatus of the provisioned credential. | [optional] [readonly] 
**User** | Pointer to [**ProvisionedCredentialUser**](ProvisionedCredentialUser.md) |  | [optional] 
**WalletActions** | Pointer to [**[]ProvisionedCredentialWalletActionsInner**](ProvisionedCredentialWalletActionsInner.md) |  | [optional] 

## Methods

### NewProvisionedCredential

`func NewProvisionedCredential() *ProvisionedCredential`

NewProvisionedCredential instantiates a new ProvisionedCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedCredentialWithDefaults

`func NewProvisionedCredentialWithDefaults() *ProvisionedCredential`

NewProvisionedCredentialWithDefaults instantiates a new ProvisionedCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ProvisionedCredential) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProvisionedCredential) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProvisionedCredential) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProvisionedCredential) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetClaimReference

`func (o *ProvisionedCredential) GetClaimReference() ProvisionedCredentialClaimReference`

GetClaimReference returns the ClaimReference field if non-nil, zero value otherwise.

### GetClaimReferenceOk

`func (o *ProvisionedCredential) GetClaimReferenceOk() (*ProvisionedCredentialClaimReference, bool)`

GetClaimReferenceOk returns a tuple with the ClaimReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimReference

`func (o *ProvisionedCredential) SetClaimReference(v ProvisionedCredentialClaimReference)`

SetClaimReference sets ClaimReference field to given value.

### HasClaimReference

`func (o *ProvisionedCredential) HasClaimReference() bool`

HasClaimReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProvisionedCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProvisionedCredential) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProvisionedCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProvisionedCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredential

`func (o *ProvisionedCredential) GetCredential() ProvisionedCredentialCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ProvisionedCredential) GetCredentialOk() (*ProvisionedCredentialCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ProvisionedCredential) SetCredential(v ProvisionedCredentialCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ProvisionedCredential) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetDigitalWallet

`func (o *ProvisionedCredential) GetDigitalWallet() ProvisionedCredentialCredential`

GetDigitalWallet returns the DigitalWallet field if non-nil, zero value otherwise.

### GetDigitalWalletOk

`func (o *ProvisionedCredential) GetDigitalWalletOk() (*ProvisionedCredentialCredential, bool)`

GetDigitalWalletOk returns a tuple with the DigitalWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallet

`func (o *ProvisionedCredential) SetDigitalWallet(v ProvisionedCredentialCredential)`

SetDigitalWallet sets DigitalWallet field to given value.

### HasDigitalWallet

`func (o *ProvisionedCredential) HasDigitalWallet() bool`

HasDigitalWallet returns a boolean if a field has been set.

### GetEnvironment

`func (o *ProvisionedCredential) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ProvisionedCredential) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ProvisionedCredential) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ProvisionedCredential) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpiredAt

`func (o *ProvisionedCredential) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ProvisionedCredential) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ProvisionedCredential) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *ProvisionedCredential) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### GetId

`func (o *ProvisionedCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionedCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionedCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionedCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ProvisionedCredential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProvisionedCredential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProvisionedCredential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProvisionedCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUser

`func (o *ProvisionedCredential) GetUser() ProvisionedCredentialUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProvisionedCredential) GetUserOk() (*ProvisionedCredentialUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProvisionedCredential) SetUser(v ProvisionedCredentialUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ProvisionedCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWalletActions

`func (o *ProvisionedCredential) GetWalletActions() []ProvisionedCredentialWalletActionsInner`

GetWalletActions returns the WalletActions field if non-nil, zero value otherwise.

### GetWalletActionsOk

`func (o *ProvisionedCredential) GetWalletActionsOk() (*[]ProvisionedCredentialWalletActionsInner, bool)`

GetWalletActionsOk returns a tuple with the WalletActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletActions

`func (o *ProvisionedCredential) SetWalletActions(v []ProvisionedCredentialWalletActionsInner)`

SetWalletActions sets WalletActions field to given value.

### HasWalletActions

`func (o *ProvisionedCredential) HasWalletActions() bool`

HasWalletActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


