# MFASettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Authentication** | Pointer to [**MFASettingsAuthentication**](MFASettingsAuthentication.md) |  | [optional] 
**Lockout** | Pointer to [**MFASettingsLockout**](MFASettingsLockout.md) |  | [optional] 
**Pairing** | [**MFASettingsPairing**](MFASettingsPairing.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewMFASettings

`func NewMFASettings(pairing MFASettingsPairing, ) *MFASettings`

NewMFASettings instantiates a new MFASettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFASettingsWithDefaults

`func NewMFASettingsWithDefaults() *MFASettings`

NewMFASettingsWithDefaults instantiates a new MFASettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *MFASettings) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *MFASettings) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *MFASettings) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *MFASettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetAuthentication

`func (o *MFASettings) GetAuthentication() MFASettingsAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *MFASettings) GetAuthenticationOk() (*MFASettingsAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *MFASettings) SetAuthentication(v MFASettingsAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *MFASettings) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetLockout

`func (o *MFASettings) GetLockout() MFASettingsLockout`

GetLockout returns the Lockout field if non-nil, zero value otherwise.

### GetLockoutOk

`func (o *MFASettings) GetLockoutOk() (*MFASettingsLockout, bool)`

GetLockoutOk returns a tuple with the Lockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockout

`func (o *MFASettings) SetLockout(v MFASettingsLockout)`

SetLockout sets Lockout field to given value.

### HasLockout

`func (o *MFASettings) HasLockout() bool`

HasLockout returns a boolean if a field has been set.

### GetPairing

`func (o *MFASettings) GetPairing() MFASettingsPairing`

GetPairing returns the Pairing field if non-nil, zero value otherwise.

### GetPairingOk

`func (o *MFASettings) GetPairingOk() (*MFASettingsPairing, bool)`

GetPairingOk returns a tuple with the Pairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairing

`func (o *MFASettings) SetPairing(v MFASettingsPairing)`

SetPairing sets Pairing field to given value.


### GetUpdatedAt

`func (o *MFASettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MFASettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MFASettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MFASettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


