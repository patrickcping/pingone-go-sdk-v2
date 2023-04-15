# CredentialDigitalWalletPairingSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Environment** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**User** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**DigitalWallet** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**Challenge** | Pointer to **string** |  | [optional] 
**QrUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCredentialDigitalWalletPairingSession

`func NewCredentialDigitalWalletPairingSession() *CredentialDigitalWalletPairingSession`

NewCredentialDigitalWalletPairingSession instantiates a new CredentialDigitalWalletPairingSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialDigitalWalletPairingSessionWithDefaults

`func NewCredentialDigitalWalletPairingSessionWithDefaults() *CredentialDigitalWalletPairingSession`

NewCredentialDigitalWalletPairingSessionWithDefaults instantiates a new CredentialDigitalWalletPairingSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CredentialDigitalWalletPairingSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CredentialDigitalWalletPairingSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CredentialDigitalWalletPairingSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CredentialDigitalWalletPairingSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialDigitalWalletPairingSession) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialDigitalWalletPairingSession) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialDigitalWalletPairingSession) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialDigitalWalletPairingSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CredentialDigitalWalletPairingSession) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CredentialDigitalWalletPairingSession) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CredentialDigitalWalletPairingSession) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CredentialDigitalWalletPairingSession) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialDigitalWalletPairingSession) GetEnvironment() CredentialDigitalWalletNotificationResultsInnerNotification`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialDigitalWalletPairingSession) GetEnvironmentOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialDigitalWalletPairingSession) SetEnvironment(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialDigitalWalletPairingSession) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUser

`func (o *CredentialDigitalWalletPairingSession) GetUser() CredentialDigitalWalletNotificationResultsInnerNotification`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CredentialDigitalWalletPairingSession) GetUserOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CredentialDigitalWalletPairingSession) SetUser(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetUser sets User field to given value.

### HasUser

`func (o *CredentialDigitalWalletPairingSession) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDigitalWallet

`func (o *CredentialDigitalWalletPairingSession) GetDigitalWallet() CredentialDigitalWalletNotificationResultsInnerNotification`

GetDigitalWallet returns the DigitalWallet field if non-nil, zero value otherwise.

### GetDigitalWalletOk

`func (o *CredentialDigitalWalletPairingSession) GetDigitalWalletOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetDigitalWalletOk returns a tuple with the DigitalWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigitalWallet

`func (o *CredentialDigitalWalletPairingSession) SetDigitalWallet(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetDigitalWallet sets DigitalWallet field to given value.

### HasDigitalWallet

`func (o *CredentialDigitalWalletPairingSession) HasDigitalWallet() bool`

HasDigitalWallet returns a boolean if a field has been set.

### GetChallenge

`func (o *CredentialDigitalWalletPairingSession) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *CredentialDigitalWalletPairingSession) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *CredentialDigitalWalletPairingSession) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *CredentialDigitalWalletPairingSession) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetQrUrl

`func (o *CredentialDigitalWalletPairingSession) GetQrUrl() string`

GetQrUrl returns the QrUrl field if non-nil, zero value otherwise.

### GetQrUrlOk

`func (o *CredentialDigitalWalletPairingSession) GetQrUrlOk() (*string, bool)`

GetQrUrlOk returns a tuple with the QrUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrUrl

`func (o *CredentialDigitalWalletPairingSession) SetQrUrl(v string)`

SetQrUrl sets QrUrl field to given value.

### HasQrUrl

`func (o *CredentialDigitalWalletPairingSession) HasQrUrl() bool`

HasQrUrl returns a boolean if a field has been set.

### GetStatus

`func (o *CredentialDigitalWalletPairingSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CredentialDigitalWalletPairingSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CredentialDigitalWalletPairingSession) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CredentialDigitalWalletPairingSession) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


