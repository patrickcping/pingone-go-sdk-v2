# UserCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** |  | [optional] 
**CredentialType** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**Environment** | Pointer to [**Environment**](Environment.md) |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Notification** | Pointer to [**CredentialDigitalWalletNotification**](CredentialDigitalWalletNotification.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**CredentialDigitalWalletNotificationResultsInnerNotification**](CredentialDigitalWalletNotificationResultsInnerNotification.md) |  | [optional] 
**Userdata** | Pointer to [**UserCredentialUserdata**](UserCredentialUserdata.md) |  | [optional] 

## Methods

### NewUserCredential

`func NewUserCredential() *UserCredential`

NewUserCredential instantiates a new UserCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialWithDefaults

`func NewUserCredentialWithDefaults() *UserCredential`

NewUserCredentialWithDefaults instantiates a new UserCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *UserCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialType

`func (o *UserCredential) GetCredentialType() CredentialDigitalWalletNotificationResultsInnerNotification`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *UserCredential) GetCredentialTypeOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *UserCredential) SetCredentialType(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *UserCredential) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetEnvironment

`func (o *UserCredential) GetEnvironment() Environment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UserCredential) GetEnvironmentOk() (*Environment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UserCredential) SetEnvironment(v Environment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UserCredential) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpiresAt

`func (o *UserCredential) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserCredential) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserCredential) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserCredential) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *UserCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotification

`func (o *UserCredential) GetNotification() CredentialDigitalWalletNotification`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *UserCredential) GetNotificationOk() (*CredentialDigitalWalletNotification, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *UserCredential) SetNotification(v CredentialDigitalWalletNotification)`

SetNotification sets Notification field to given value.

### HasNotification

`func (o *UserCredential) HasNotification() bool`

HasNotification returns a boolean if a field has been set.

### GetStatus

`func (o *UserCredential) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCredential) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCredential) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCredential) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserCredential) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserCredential) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserCredential) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserCredential) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *UserCredential) GetUser() CredentialDigitalWalletNotificationResultsInnerNotification`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserCredential) GetUserOk() (*CredentialDigitalWalletNotificationResultsInnerNotification, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserCredential) SetUser(v CredentialDigitalWalletNotificationResultsInnerNotification)`

SetUser sets User field to given value.

### HasUser

`func (o *UserCredential) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserdata

`func (o *UserCredential) GetUserdata() UserCredentialUserdata`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *UserCredential) GetUserdataOk() (*UserCredentialUserdata, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *UserCredential) SetUserdata(v UserCredentialUserdata)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *UserCredential) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


