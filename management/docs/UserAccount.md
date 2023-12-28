# UserAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanAuthenticate** | **bool** | A boolean that specifies whether the user can authenticate. If the value is set to &#x60;false&#x60;, the account is locked or the user is disabled, and unless specified otherwise by the administrative configuration, the user will be unable to authenticate. | 
**LockedAt** | Pointer to **time.Time** | A string that specifies the date-time the specified user account was locked. This property might be absent if the account is unlocked or if the account was locked out automatically by failed password attempts. | [optional] [readonly] 
**SecondsUntilUnlock** | Pointer to **int32** | An integer that specifies the number of seconds until the user&#39;s account is unlocked. This property is absent if the account is unlocked, or if it will not automatically unlock (and must be unlocked by an administrator). | [optional] [readonly] 
**Status** | [**EnumUserStatus**](EnumUserStatus.md) |  | 
**UnlockAt** | Pointer to **time.Time** | A string that specifies the time the specified user account will be unlocked. This property is absent if the account is unlocked, or if it will not automatically unlock (and must be unlocked by an administrator). | [optional] 

## Methods

### NewUserAccount

`func NewUserAccount(canAuthenticate bool, status EnumUserStatus, ) *UserAccount`

NewUserAccount instantiates a new UserAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccountWithDefaults

`func NewUserAccountWithDefaults() *UserAccount`

NewUserAccountWithDefaults instantiates a new UserAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanAuthenticate

`func (o *UserAccount) GetCanAuthenticate() bool`

GetCanAuthenticate returns the CanAuthenticate field if non-nil, zero value otherwise.

### GetCanAuthenticateOk

`func (o *UserAccount) GetCanAuthenticateOk() (*bool, bool)`

GetCanAuthenticateOk returns a tuple with the CanAuthenticate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAuthenticate

`func (o *UserAccount) SetCanAuthenticate(v bool)`

SetCanAuthenticate sets CanAuthenticate field to given value.


### GetLockedAt

`func (o *UserAccount) GetLockedAt() time.Time`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *UserAccount) GetLockedAtOk() (*time.Time, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *UserAccount) SetLockedAt(v time.Time)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *UserAccount) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetSecondsUntilUnlock

`func (o *UserAccount) GetSecondsUntilUnlock() int32`

GetSecondsUntilUnlock returns the SecondsUntilUnlock field if non-nil, zero value otherwise.

### GetSecondsUntilUnlockOk

`func (o *UserAccount) GetSecondsUntilUnlockOk() (*int32, bool)`

GetSecondsUntilUnlockOk returns a tuple with the SecondsUntilUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsUntilUnlock

`func (o *UserAccount) SetSecondsUntilUnlock(v int32)`

SetSecondsUntilUnlock sets SecondsUntilUnlock field to given value.

### HasSecondsUntilUnlock

`func (o *UserAccount) HasSecondsUntilUnlock() bool`

HasSecondsUntilUnlock returns a boolean if a field has been set.

### GetStatus

`func (o *UserAccount) GetStatus() EnumUserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserAccount) GetStatusOk() (*EnumUserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserAccount) SetStatus(v EnumUserStatus)`

SetStatus sets Status field to given value.


### GetUnlockAt

`func (o *UserAccount) GetUnlockAt() time.Time`

GetUnlockAt returns the UnlockAt field if non-nil, zero value otherwise.

### GetUnlockAtOk

`func (o *UserAccount) GetUnlockAtOk() (*time.Time, bool)`

GetUnlockAtOk returns a tuple with the UnlockAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockAt

`func (o *UserAccount) SetUnlockAt(v time.Time)`

SetUnlockAt sets UnlockAt field to given value.

### HasUnlockAt

`func (o *UserAccount) HasUnlockAt() bool`

HasUnlockAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


