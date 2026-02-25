# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]EntityArrayEmbeddedItemsInner**](EntityArrayEmbeddedItemsInner.md) |  | [optional] 
**IssuanceRules** | Pointer to [**[]CredentialIssuanceRule**](CredentialIssuanceRule.md) |  | [optional] 
**StagedChanges** | Pointer to [**[]CredentialIssuanceRuleStagedChange**](CredentialIssuanceRuleStagedChange.md) |  | [optional] 
**DigitalWalletApplications** | Pointer to [**[]DigitalWalletApplication**](DigitalWalletApplication.md) |  | [optional] 
**DigitalWallets** | Pointer to [**[]CredentialDigitalWallet**](CredentialDigitalWallet.md) |  | [optional] 
**ProvisionedCredentials** | Pointer to [**[]ProvisionedCredential**](ProvisionedCredential.md) |  | [optional] 
**Versions** | Pointer to [**[]CredentialTypeVersion**](CredentialTypeVersion.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EntityArrayEmbedded) GetItems() []EntityArrayEmbeddedItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EntityArrayEmbedded) GetItemsOk() (*[]EntityArrayEmbeddedItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EntityArrayEmbedded) SetItems(v []EntityArrayEmbeddedItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *EntityArrayEmbedded) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetIssuanceRules

`func (o *EntityArrayEmbedded) GetIssuanceRules() []CredentialIssuanceRule`

GetIssuanceRules returns the IssuanceRules field if non-nil, zero value otherwise.

### GetIssuanceRulesOk

`func (o *EntityArrayEmbedded) GetIssuanceRulesOk() (*[]CredentialIssuanceRule, bool)`

GetIssuanceRulesOk returns a tuple with the IssuanceRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceRules

`func (o *EntityArrayEmbedded) SetIssuanceRules(v []CredentialIssuanceRule)`

SetIssuanceRules sets IssuanceRules field to given value.

### HasIssuanceRules

`func (o *EntityArrayEmbedded) HasIssuanceRules() bool`

HasIssuanceRules returns a boolean if a field has been set.

### GetStagedChanges

`func (o *EntityArrayEmbedded) GetStagedChanges() []CredentialIssuanceRuleStagedChange`

GetStagedChanges returns the StagedChanges field if non-nil, zero value otherwise.

### GetStagedChangesOk

`func (o *EntityArrayEmbedded) GetStagedChangesOk() (*[]CredentialIssuanceRuleStagedChange, bool)`

GetStagedChangesOk returns a tuple with the StagedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagedChanges

`func (o *EntityArrayEmbedded) SetStagedChanges(v []CredentialIssuanceRuleStagedChange)`

SetStagedChanges sets StagedChanges field to given value.

### HasStagedChanges

`func (o *EntityArrayEmbedded) HasStagedChanges() bool`

HasStagedChanges returns a boolean if a field has been set.

### GetDigitalWalletApplications

`func (o *EntityArrayEmbedded) GetDigitalWalletApplications() []DigitalWalletApplication`

GetDigitalWalletApplications returns the DigitalWalletApplications field if non-nil, zero value otherwise.

### GetDigitalWalletApplicationsOk

`func (o *EntityArrayEmbedded) GetDigitalWalletApplicationsOk() (*[]DigitalWalletApplication, bool)`

GetDigitalWalletApplicationsOk returns a tuple with the DigitalWalletApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletApplications

`func (o *EntityArrayEmbedded) SetDigitalWalletApplications(v []DigitalWalletApplication)`

SetDigitalWalletApplications sets DigitalWalletApplications field to given value.

### HasDigitalWalletApplications

`func (o *EntityArrayEmbedded) HasDigitalWalletApplications() bool`

HasDigitalWalletApplications returns a boolean if a field has been set.

### GetDigitalWallets

`func (o *EntityArrayEmbedded) GetDigitalWallets() []CredentialDigitalWallet`

GetDigitalWallets returns the DigitalWallets field if non-nil, zero value otherwise.

### GetDigitalWalletsOk

`func (o *EntityArrayEmbedded) GetDigitalWalletsOk() (*[]CredentialDigitalWallet, bool)`

GetDigitalWalletsOk returns a tuple with the DigitalWallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallets

`func (o *EntityArrayEmbedded) SetDigitalWallets(v []CredentialDigitalWallet)`

SetDigitalWallets sets DigitalWallets field to given value.

### HasDigitalWallets

`func (o *EntityArrayEmbedded) HasDigitalWallets() bool`

HasDigitalWallets returns a boolean if a field has been set.

### GetProvisionedCredentials

`func (o *EntityArrayEmbedded) GetProvisionedCredentials() []ProvisionedCredential`

GetProvisionedCredentials returns the ProvisionedCredentials field if non-nil, zero value otherwise.

### GetProvisionedCredentialsOk

`func (o *EntityArrayEmbedded) GetProvisionedCredentialsOk() (*[]ProvisionedCredential, bool)`

GetProvisionedCredentialsOk returns a tuple with the ProvisionedCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedCredentials

`func (o *EntityArrayEmbedded) SetProvisionedCredentials(v []ProvisionedCredential)`

SetProvisionedCredentials sets ProvisionedCredentials field to given value.

### HasProvisionedCredentials

`func (o *EntityArrayEmbedded) HasProvisionedCredentials() bool`

HasProvisionedCredentials returns a boolean if a field has been set.

### GetVersions

`func (o *EntityArrayEmbedded) GetVersions() []CredentialTypeVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *EntityArrayEmbedded) GetVersionsOk() (*[]CredentialTypeVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *EntityArrayEmbedded) SetVersions(v []CredentialTypeVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *EntityArrayEmbedded) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


