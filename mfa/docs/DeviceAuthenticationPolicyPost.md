# DeviceAuthenticationPolicyPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | Device authentication policy&#39;s UUID. | [optional] [readonly] 
**Name** | **string** | Device authentication policy&#39;s name. | 
**NewDeviceNotification** | Pointer to [**EnumMFADevicePolicyNewDeviceNotification**](EnumMFADevicePolicyNewDeviceNotification.md) |  | [optional] 
**Authentication** | Pointer to [**DeviceAuthenticationPolicyCommonAuthentication**](DeviceAuthenticationPolicyCommonAuthentication.md) |  | [optional] 
**Sms** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Voice** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Email** | [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | 
**Fido2** | Pointer to [**DeviceAuthenticationPolicyCommonFido2**](DeviceAuthenticationPolicyCommonFido2.md) |  | [optional] 
**Mobile** | [**DeviceAuthenticationPolicyCommonMobile**](DeviceAuthenticationPolicyCommonMobile.md) |  | 
**Totp** | [**DeviceAuthenticationPolicyCommonTotp**](DeviceAuthenticationPolicyCommonTotp.md) |  | 
**Default** | **bool** | A boolean that specifies whether the policy is the default for the environment. | 
**ForSignOnPolicy** | **bool** |  | [readonly] 
**IgnoreUserLock** | Pointer to **bool** | When applying an MFA policy, PingOne ordinarily checks if a user account is locked, and if so, prevents the user from authenticating. Set &#x60;ignoreUserLock&#x60; to &#x60;true&#x60; if you want PingOne to skip this account lock check. | [optional] 
**NotificationsPolicy** | Pointer to [**DeviceAuthenticationPolicyCommonNotificationsPolicy**](DeviceAuthenticationPolicyCommonNotificationsPolicy.md) |  | [optional] 
**OathToken** | Pointer to [**DeviceAuthenticationPolicyOathToken**](DeviceAuthenticationPolicyOathToken.md) |  | [optional] 
**RememberMe** | Pointer to [**DeviceAuthenticationPolicyCommonRememberMe**](DeviceAuthenticationPolicyCommonRememberMe.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Whatsapp** | Pointer to [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | [optional] 
**Desktop** | Pointer to [**DeviceAuthenticationPolicyPingIDDevice**](DeviceAuthenticationPolicyPingIDDevice.md) |  | [optional] 
**Yubikey** | Pointer to [**DeviceAuthenticationPolicyPingIDDevice**](DeviceAuthenticationPolicyPingIDDevice.md) |  | [optional] 
**MigrationData** | [**[]DeviceAuthenticationPolicyMigrateData**](DeviceAuthenticationPolicyMigrateData.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyPost

`func NewDeviceAuthenticationPolicyPost(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyCommonMobile, totp DeviceAuthenticationPolicyCommonTotp, default_ bool, forSignOnPolicy bool, migrationData []DeviceAuthenticationPolicyMigrateData, ) *DeviceAuthenticationPolicyPost`

NewDeviceAuthenticationPolicyPost instantiates a new DeviceAuthenticationPolicyPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPostWithDefaults

`func NewDeviceAuthenticationPolicyPostWithDefaults() *DeviceAuthenticationPolicyPost`

NewDeviceAuthenticationPolicyPostWithDefaults instantiates a new DeviceAuthenticationPolicyPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeviceAuthenticationPolicyPost) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAuthenticationPolicyPost) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAuthenticationPolicyPost) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAuthenticationPolicyPost) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

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


### GetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPost) GetNewDeviceNotification() EnumMFADevicePolicyNewDeviceNotification`

GetNewDeviceNotification returns the NewDeviceNotification field if non-nil, zero value otherwise.

### GetNewDeviceNotificationOk

`func (o *DeviceAuthenticationPolicyPost) GetNewDeviceNotificationOk() (*EnumMFADevicePolicyNewDeviceNotification, bool)`

GetNewDeviceNotificationOk returns a tuple with the NewDeviceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPost) SetNewDeviceNotification(v EnumMFADevicePolicyNewDeviceNotification)`

SetNewDeviceNotification sets NewDeviceNotification field to given value.

### HasNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPost) HasNewDeviceNotification() bool`

HasNewDeviceNotification returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeviceAuthenticationPolicyPost) GetAuthentication() DeviceAuthenticationPolicyCommonAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeviceAuthenticationPolicyPost) GetAuthenticationOk() (*DeviceAuthenticationPolicyCommonAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeviceAuthenticationPolicyPost) SetAuthentication(v DeviceAuthenticationPolicyCommonAuthentication)`

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

`func (o *DeviceAuthenticationPolicyPost) GetFido2() DeviceAuthenticationPolicyCommonFido2`

GetFido2 returns the Fido2 field if non-nil, zero value otherwise.

### GetFido2Ok

`func (o *DeviceAuthenticationPolicyPost) GetFido2Ok() (*DeviceAuthenticationPolicyCommonFido2, bool)`

GetFido2Ok returns a tuple with the Fido2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2

`func (o *DeviceAuthenticationPolicyPost) SetFido2(v DeviceAuthenticationPolicyCommonFido2)`

SetFido2 sets Fido2 field to given value.

### HasFido2

`func (o *DeviceAuthenticationPolicyPost) HasFido2() bool`

HasFido2 returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicyPost) GetMobile() DeviceAuthenticationPolicyCommonMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicyPost) GetMobileOk() (*DeviceAuthenticationPolicyCommonMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicyPost) SetMobile(v DeviceAuthenticationPolicyCommonMobile)`

SetMobile sets Mobile field to given value.


### GetTotp

`func (o *DeviceAuthenticationPolicyPost) GetTotp() DeviceAuthenticationPolicyCommonTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicyPost) GetTotpOk() (*DeviceAuthenticationPolicyCommonTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicyPost) SetTotp(v DeviceAuthenticationPolicyCommonTotp)`

SetTotp sets Totp field to given value.


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


### GetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPost) GetIgnoreUserLock() bool`

GetIgnoreUserLock returns the IgnoreUserLock field if non-nil, zero value otherwise.

### GetIgnoreUserLockOk

`func (o *DeviceAuthenticationPolicyPost) GetIgnoreUserLockOk() (*bool, bool)`

GetIgnoreUserLockOk returns a tuple with the IgnoreUserLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPost) SetIgnoreUserLock(v bool)`

SetIgnoreUserLock sets IgnoreUserLock field to given value.

### HasIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPost) HasIgnoreUserLock() bool`

HasIgnoreUserLock returns a boolean if a field has been set.

### GetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPost) GetNotificationsPolicy() DeviceAuthenticationPolicyCommonNotificationsPolicy`

GetNotificationsPolicy returns the NotificationsPolicy field if non-nil, zero value otherwise.

### GetNotificationsPolicyOk

`func (o *DeviceAuthenticationPolicyPost) GetNotificationsPolicyOk() (*DeviceAuthenticationPolicyCommonNotificationsPolicy, bool)`

GetNotificationsPolicyOk returns a tuple with the NotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPost) SetNotificationsPolicy(v DeviceAuthenticationPolicyCommonNotificationsPolicy)`

SetNotificationsPolicy sets NotificationsPolicy field to given value.

### HasNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPost) HasNotificationsPolicy() bool`

HasNotificationsPolicy returns a boolean if a field has been set.

### GetOathToken

`func (o *DeviceAuthenticationPolicyPost) GetOathToken() DeviceAuthenticationPolicyOathToken`

GetOathToken returns the OathToken field if non-nil, zero value otherwise.

### GetOathTokenOk

`func (o *DeviceAuthenticationPolicyPost) GetOathTokenOk() (*DeviceAuthenticationPolicyOathToken, bool)`

GetOathTokenOk returns a tuple with the OathToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOathToken

`func (o *DeviceAuthenticationPolicyPost) SetOathToken(v DeviceAuthenticationPolicyOathToken)`

SetOathToken sets OathToken field to given value.

### HasOathToken

`func (o *DeviceAuthenticationPolicyPost) HasOathToken() bool`

HasOathToken returns a boolean if a field has been set.

### GetRememberMe

`func (o *DeviceAuthenticationPolicyPost) GetRememberMe() DeviceAuthenticationPolicyCommonRememberMe`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *DeviceAuthenticationPolicyPost) GetRememberMeOk() (*DeviceAuthenticationPolicyCommonRememberMe, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *DeviceAuthenticationPolicyPost) SetRememberMe(v DeviceAuthenticationPolicyCommonRememberMe)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *DeviceAuthenticationPolicyPost) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

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

### GetWhatsapp

`func (o *DeviceAuthenticationPolicyPost) GetWhatsapp() DeviceAuthenticationPolicyOfflineDevice`

GetWhatsapp returns the Whatsapp field if non-nil, zero value otherwise.

### GetWhatsappOk

`func (o *DeviceAuthenticationPolicyPost) GetWhatsappOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetWhatsappOk returns a tuple with the Whatsapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsapp

`func (o *DeviceAuthenticationPolicyPost) SetWhatsapp(v DeviceAuthenticationPolicyOfflineDevice)`

SetWhatsapp sets Whatsapp field to given value.

### HasWhatsapp

`func (o *DeviceAuthenticationPolicyPost) HasWhatsapp() bool`

HasWhatsapp returns a boolean if a field has been set.

### GetDesktop

`func (o *DeviceAuthenticationPolicyPost) GetDesktop() DeviceAuthenticationPolicyPingIDDevice`

GetDesktop returns the Desktop field if non-nil, zero value otherwise.

### GetDesktopOk

`func (o *DeviceAuthenticationPolicyPost) GetDesktopOk() (*DeviceAuthenticationPolicyPingIDDevice, bool)`

GetDesktopOk returns a tuple with the Desktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktop

`func (o *DeviceAuthenticationPolicyPost) SetDesktop(v DeviceAuthenticationPolicyPingIDDevice)`

SetDesktop sets Desktop field to given value.

### HasDesktop

`func (o *DeviceAuthenticationPolicyPost) HasDesktop() bool`

HasDesktop returns a boolean if a field has been set.

### GetYubikey

`func (o *DeviceAuthenticationPolicyPost) GetYubikey() DeviceAuthenticationPolicyPingIDDevice`

GetYubikey returns the Yubikey field if non-nil, zero value otherwise.

### GetYubikeyOk

`func (o *DeviceAuthenticationPolicyPost) GetYubikeyOk() (*DeviceAuthenticationPolicyPingIDDevice, bool)`

GetYubikeyOk returns a tuple with the Yubikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikey

`func (o *DeviceAuthenticationPolicyPost) SetYubikey(v DeviceAuthenticationPolicyPingIDDevice)`

SetYubikey sets Yubikey field to given value.

### HasYubikey

`func (o *DeviceAuthenticationPolicyPost) HasYubikey() bool`

HasYubikey returns a boolean if a field has been set.

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


