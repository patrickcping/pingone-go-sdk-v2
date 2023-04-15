# CredentialDigitalWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | A string that specifies the date and time the credential digital wallet was created. | [optional] 
**DigitalWalletApplication** | Pointer to [**CredentialDigitalWalletDigitalWalletApplication**](CredentialDigitalWalletDigitalWalletApplication.md) |  | [optional] 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the identifier (UUID) associated with the credential digital wallet app. | [optional] 
**Notification** | Pointer to [**CredentialDigitalWalletNotification**](CredentialDigitalWalletNotification.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** | A string that specifies the date and time the credential digital wallet was last updated; can be null. | [optional] 
**User** | Pointer to [**CredentialDigitalWalletUser**](CredentialDigitalWalletUser.md) |  | [optional] 
**Status** | Pointer to **string** | A string that specifies the status of the wallet. | [optional] 
**PairingSession** | Pointer to [**CredentialDigitalWalletPairingSession**](CredentialDigitalWalletPairingSession.md) |  | [optional] 

## Methods

### NewCredentialDigitalWallet

`func NewCredentialDigitalWallet() *CredentialDigitalWallet`

NewCredentialDigitalWallet instantiates a new CredentialDigitalWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialDigitalWalletWithDefaults

`func NewCredentialDigitalWalletWithDefaults() *CredentialDigitalWallet`

NewCredentialDigitalWalletWithDefaults instantiates a new CredentialDigitalWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *CredentialDigitalWallet) GetApplication() CredentialDigitalWalletApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *CredentialDigitalWallet) GetApplicationOk() (*CredentialDigitalWalletApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *CredentialDigitalWallet) SetApplication(v CredentialDigitalWalletApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *CredentialDigitalWallet) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialDigitalWallet) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialDigitalWallet) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialDigitalWallet) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialDigitalWallet) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDigitalWalletApplication

`func (o *CredentialDigitalWallet) GetDigitalWalletApplication() CredentialDigitalWalletDigitalWalletApplication`

GetDigitalWalletApplication returns the DigitalWalletApplication field if non-nil, zero value otherwise.

### GetDigitalWalletApplicationOk

`func (o *CredentialDigitalWallet) GetDigitalWalletApplicationOk() (*CredentialDigitalWalletDigitalWalletApplication, bool)`

GetDigitalWalletApplicationOk returns a tuple with the DigitalWalletApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWalletApplication

`func (o *CredentialDigitalWallet) SetDigitalWalletApplication(v CredentialDigitalWalletDigitalWalletApplication)`

SetDigitalWalletApplication sets DigitalWalletApplication field to given value.

### HasDigitalWalletApplication

`func (o *CredentialDigitalWallet) HasDigitalWalletApplication() bool`

HasDigitalWalletApplication returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialDigitalWallet) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialDigitalWallet) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialDigitalWallet) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialDigitalWallet) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CredentialDigitalWallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialDigitalWallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialDigitalWallet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialDigitalWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotification

`func (o *CredentialDigitalWallet) GetNotification() CredentialDigitalWalletNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *CredentialDigitalWallet) GetNotificationOk() (*CredentialDigitalWalletNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *CredentialDigitalWallet) SetNotification(v CredentialDigitalWalletNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *CredentialDigitalWallet) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CredentialDigitalWallet) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialDigitalWallet) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialDigitalWallet) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialDigitalWallet) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *CredentialDigitalWallet) GetUser() CredentialDigitalWalletUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CredentialDigitalWallet) GetUserOk() (*CredentialDigitalWalletUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CredentialDigitalWallet) SetUser(v CredentialDigitalWalletUser)`

SetUser sets User field to given value.

### HasUser

`func (o *CredentialDigitalWallet) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetStatus

`func (o *CredentialDigitalWallet) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CredentialDigitalWallet) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CredentialDigitalWallet) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CredentialDigitalWallet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPairingSession

`func (o *CredentialDigitalWallet) GetPairingSession() CredentialDigitalWalletPairingSession`

GetPairingSession returns the PairingSession field if non-nil, zero value otherwise.

### GetPairingSessionOk

`func (o *CredentialDigitalWallet) GetPairingSessionOk() (*CredentialDigitalWalletPairingSession, bool)`

GetPairingSessionOk returns a tuple with the PairingSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingSession

`func (o *CredentialDigitalWallet) SetPairingSession(v CredentialDigitalWalletPairingSession)`

SetPairingSession sets PairingSession field to given value.

### HasPairingSession

`func (o *CredentialDigitalWallet) HasPairingSession() bool`

HasPairingSession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


