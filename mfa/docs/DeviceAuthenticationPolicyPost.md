# DeviceAuthenticationPolicyPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | Device authentication policy&#39;s UUID. | [optional] [readonly] 
**Name** | **string** | Device authentication policy&#39;s name. | 
**Authentication** | Pointer to [**DeviceAuthenticationPolicyAuthentication**](DeviceAuthenticationPolicyAuthentication.md) |  | [optional] 
**Sms** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Voice** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Email** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Fido2** | Pointer to [**DeviceAuthenticationPolicyFido2**](DeviceAuthenticationPolicyFido2.md) |  | [optional] 
**Mobile** | [**DeviceAuthenticationPolicyMobile**](DeviceAuthenticationPolicyMobile.md) |  | 
**Totp** | [**DeviceAuthenticationPolicyTotp**](DeviceAuthenticationPolicyTotp.md) |  | 
**SecurityKey** | Pointer to [**DeviceAuthenticationPolicyFIDODevice**](DeviceAuthenticationPolicyFIDODevice.md) |  | [optional] 
**Platform** | Pointer to [**DeviceAuthenticationPolicyFIDODevice**](DeviceAuthenticationPolicyFIDODevice.md) |  | [optional] 
**Default** | **bool** | A boolean that specifies whether the policy is the default for the environment. | 
**ForSignOnPolicy** | **bool** |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**MigrationData** | [**[]DeviceAuthenticationPolicyMigrateData**](DeviceAuthenticationPolicyMigrateData.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyPost

`func NewDeviceAuthenticationPolicyPost(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyMobile, totp DeviceAuthenticationPolicyTotp, default_ bool, forSignOnPolicy bool, migrationData []DeviceAuthenticationPolicyMigrateData, ) *DeviceAuthenticationPolicyPost`

NewDeviceAuthenticationPolicyPost instantiates a new DeviceAuthenticationPolicyPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPostWithDefaults

`func NewDeviceAuthenticationPolicyPostWithDefaults() *DeviceAuthenticationPolicyPost`

NewDeviceAuthenticationPolicyPostWithDefaults instantiates a new DeviceAuthenticationPolicyPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DeviceAuthenticationPolicyPost) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceAuthenticationPolicyPost) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceAuthenticationPolicyPost) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceAuthenticationPolicyPost) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DeviceAuthenticationPolicyPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAuthenticationPolicyPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceAuthenticationPolicyPost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthenticationPolicyPost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthenticationPolicyPost) SetName(v string)`

SetName sets Name field to given value.


### GetAuthentication

`func (o *DeviceAuthenticationPolicyPost) GetAuthentication() DeviceAuthenticationPolicyAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeviceAuthenticationPolicyPost) GetAuthenticationOk() (*DeviceAuthenticationPolicyAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeviceAuthenticationPolicyPost) SetAuthentication(v DeviceAuthenticationPolicyAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeviceAuthenticationPolicyPost) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSms

`func (o *DeviceAuthenticationPolicyPost) GetSms() DeviceAuthenticationPolicyOfflineDevice`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *DeviceAuthenticationPolicyPost) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *DeviceAuthenticationPolicyPost) SetSms(v DeviceAuthenticationPolicyOfflineDevice)`

SetSms sets Sms field to given value.


### GetVoice

`func (o *DeviceAuthenticationPolicyPost) GetVoice() DeviceAuthenticationPolicyOfflineDevice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *DeviceAuthenticationPolicyPost) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *DeviceAuthenticationPolicyPost) SetVoice(v DeviceAuthenticationPolicyOfflineDevice)`

SetVoice sets Voice field to given value.


### GetEmail

`func (o *DeviceAuthenticationPolicyPost) GetEmail() DeviceAuthenticationPolicyOfflineDevice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceAuthenticationPolicyPost) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceAuthenticationPolicyPost) SetEmail(v DeviceAuthenticationPolicyOfflineDevice)`

SetEmail sets Email field to given value.


### GetFido2

`func (o *DeviceAuthenticationPolicyPost) GetFido2() DeviceAuthenticationPolicyFido2`

GetFido2 returns the Fido2 field if non-nil, zero value otherwise.

### GetFido2Ok

`func (o *DeviceAuthenticationPolicyPost) GetFido2Ok() (*DeviceAuthenticationPolicyFido2, bool)`

GetFido2Ok returns a tuple with the Fido2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2

`func (o *DeviceAuthenticationPolicyPost) SetFido2(v DeviceAuthenticationPolicyFido2)`

SetFido2 sets Fido2 field to given value.

### HasFido2

`func (o *DeviceAuthenticationPolicyPost) HasFido2() bool`

HasFido2 returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicyPost) GetMobile() DeviceAuthenticationPolicyMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicyPost) GetMobileOk() (*DeviceAuthenticationPolicyMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicyPost) SetMobile(v DeviceAuthenticationPolicyMobile)`

SetMobile sets Mobile field to given value.


### GetTotp

`func (o *DeviceAuthenticationPolicyPost) GetTotp() DeviceAuthenticationPolicyTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicyPost) GetTotpOk() (*DeviceAuthenticationPolicyTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicyPost) SetTotp(v DeviceAuthenticationPolicyTotp)`

SetTotp sets Totp field to given value.


### GetSecurityKey

`func (o *DeviceAuthenticationPolicyPost) GetSecurityKey() DeviceAuthenticationPolicyFIDODevice`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *DeviceAuthenticationPolicyPost) GetSecurityKeyOk() (*DeviceAuthenticationPolicyFIDODevice, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *DeviceAuthenticationPolicyPost) SetSecurityKey(v DeviceAuthenticationPolicyFIDODevice)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *DeviceAuthenticationPolicyPost) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceAuthenticationPolicyPost) GetPlatform() DeviceAuthenticationPolicyFIDODevice`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceAuthenticationPolicyPost) GetPlatformOk() (*DeviceAuthenticationPolicyFIDODevice, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceAuthenticationPolicyPost) SetPlatform(v DeviceAuthenticationPolicyFIDODevice)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceAuthenticationPolicyPost) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDefault

`func (o *DeviceAuthenticationPolicyPost) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeviceAuthenticationPolicyPost) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeviceAuthenticationPolicyPost) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPost) GetForSignOnPolicy() bool`

GetForSignOnPolicy returns the ForSignOnPolicy field if non-nil, zero value otherwise.

### GetForSignOnPolicyOk

`func (o *DeviceAuthenticationPolicyPost) GetForSignOnPolicyOk() (*bool, bool)`

GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPost) SetForSignOnPolicy(v bool)`

SetForSignOnPolicy sets ForSignOnPolicy field to given value.


### GetUpdatedAt

`func (o *DeviceAuthenticationPolicyPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthenticationPolicyPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthenticationPolicyPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthenticationPolicyPost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMigrationData

`func (o *DeviceAuthenticationPolicyPost) GetMigrationData() []DeviceAuthenticationPolicyMigrateData`

GetMigrationData returns the MigrationData field if non-nil, zero value otherwise.

### GetMigrationDataOk

`func (o *DeviceAuthenticationPolicyPost) GetMigrationDataOk() (*[]DeviceAuthenticationPolicyMigrateData, bool)`

GetMigrationDataOk returns a tuple with the MigrationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationData

`func (o *DeviceAuthenticationPolicyPost) SetMigrationData(v []DeviceAuthenticationPolicyMigrateData)`

SetMigrationData sets MigrationData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


