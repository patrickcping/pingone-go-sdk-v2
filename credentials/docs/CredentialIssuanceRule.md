# CredentialIssuanceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automation** | [**CredentialIssuanceRuleAutomation**](CredentialIssuanceRuleAutomation.md) |  | 
**CreatedAt** | Pointer to **string** | A string that specifies the date and time the credential issuance rule was created. | [optional] 
**CredentialType** | Pointer to [**CredentialIssuanceRuleCredentialType**](CredentialIssuanceRuleCredentialType.md) |  | [optional] 
**DigitalWalletApplication** | Pointer to [**CredentialIssuanceRuleDigitalWalletApplication**](CredentialIssuanceRuleDigitalWalletApplication.md) |  | [optional] 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**Filter** | Pointer to [**CredentialIssuanceRuleFilter**](CredentialIssuanceRuleFilter.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) of the credential issuance rule. | [optional] 
**Notification** | Pointer to [**CredentialIssuanceRuleNotification**](CredentialIssuanceRuleNotification.md) |  | [optional] 
**Status** | [**EnumCredentialIssuanceRuleStatus**](EnumCredentialIssuanceRuleStatus.md) |  | 
**UpdatedAt** | Pointer to **string** | A string that specifies the date and time the credential issuance rule was last updated; can be null. | [optional] 

## Methods

### NewCredentialIssuanceRule

`func NewCredentialIssuanceRule(automation CredentialIssuanceRuleAutomation, status EnumCredentialIssuanceRuleStatus, ) *CredentialIssuanceRule`

NewCredentialIssuanceRule instantiates a new CredentialIssuanceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleWithDefaults

`func NewCredentialIssuanceRuleWithDefaults() *CredentialIssuanceRule`

NewCredentialIssuanceRuleWithDefaults instantiates a new CredentialIssuanceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomation

`func (o *CredentialIssuanceRule) GetAutomation() CredentialIssuanceRuleAutomation`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *CredentialIssuanceRule) GetAutomationOk() (*CredentialIssuanceRuleAutomation, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *CredentialIssuanceRule) SetAutomation(v CredentialIssuanceRuleAutomation)`

SetAutomation sets Automation field to given value.


### GetCreatedAt

`func (o *CredentialIssuanceRule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialIssuanceRule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialIssuanceRule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialIssuanceRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialType

`func (o *CredentialIssuanceRule) GetCredentialType() CredentialIssuanceRuleCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CredentialIssuanceRule) GetCredentialTypeOk() (*CredentialIssuanceRuleCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CredentialIssuanceRule) SetCredentialType(v CredentialIssuanceRuleCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *CredentialIssuanceRule) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetDigitalWalletApplication

`func (o *CredentialIssuanceRule) GetDigitalWalletApplication() CredentialIssuanceRuleDigitalWalletApplication`

GetDigitalWalletApplication returns the DigitalWalletApplication field if non-nil, zero value otherwise.

### GetDigitalWalletApplicationOk

`func (o *CredentialIssuanceRule) GetDigitalWalletApplicationOk() (*CredentialIssuanceRuleDigitalWalletApplication, bool)`

GetDigitalWalletApplicationOk returns a tuple with the DigitalWalletApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletApplication

`func (o *CredentialIssuanceRule) SetDigitalWalletApplication(v CredentialIssuanceRuleDigitalWalletApplication)`

SetDigitalWalletApplication sets DigitalWalletApplication field to given value.

### HasDigitalWalletApplication

`func (o *CredentialIssuanceRule) HasDigitalWalletApplication() bool`

HasDigitalWalletApplication returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialIssuanceRule) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialIssuanceRule) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialIssuanceRule) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialIssuanceRule) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFilter

`func (o *CredentialIssuanceRule) GetFilter() CredentialIssuanceRuleFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CredentialIssuanceRule) GetFilterOk() (*CredentialIssuanceRuleFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CredentialIssuanceRule) SetFilter(v CredentialIssuanceRuleFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CredentialIssuanceRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *CredentialIssuanceRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialIssuanceRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialIssuanceRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialIssuanceRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotification

`func (o *CredentialIssuanceRule) GetNotification() CredentialIssuanceRuleNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *CredentialIssuanceRule) GetNotificationOk() (*CredentialIssuanceRuleNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *CredentialIssuanceRule) SetNotification(v CredentialIssuanceRuleNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *CredentialIssuanceRule) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *CredentialIssuanceRule) GetStatus() EnumCredentialIssuanceRuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CredentialIssuanceRule) GetStatusOk() (*EnumCredentialIssuanceRuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CredentialIssuanceRule) SetStatus(v EnumCredentialIssuanceRuleStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *CredentialIssuanceRule) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialIssuanceRule) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialIssuanceRule) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialIssuanceRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


