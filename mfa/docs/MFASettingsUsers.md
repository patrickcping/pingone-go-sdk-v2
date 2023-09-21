# MFASettingsUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaEnabled** | Pointer to **bool** | Set &#x60;mfaEnabled&#x60; to &#x60;true&#x60; if you want MFA to be enabled by default for new users. | [optional] 

## Methods

### NewMFASettingsUsers

`func NewMFASettingsUsers() *MFASettingsUsers`

NewMFASettingsUsers instantiates a new MFASettingsUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFASettingsUsersWithDefaults

`func NewMFASettingsUsersWithDefaults() *MFASettingsUsers`

NewMFASettingsUsersWithDefaults instantiates a new MFASettingsUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaEnabled

`func (o *MFASettingsUsers) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *MFASettingsUsers) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *MFASettingsUsers) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.

### HasMfaEnabled

`func (o *MFASettingsUsers) HasMfaEnabled() bool`

HasMfaEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


