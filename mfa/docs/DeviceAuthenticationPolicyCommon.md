# DeviceAuthenticationPolicyCommon

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

## Methods

### NewDeviceAuthenticationPolicyCommon

`func NewDeviceAuthenticationPolicyCommon(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyCommonMobile, totp DeviceAuthenticationPolicyCommonTotp, default_ bool, forSignOnPolicy bool, ) *DeviceAuthenticationPolicyCommon`

NewDeviceAuthenticationPolicyCommon instantiates a new DeviceAuthenticationPolicyCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonWithDefaults

`func NewDeviceAuthenticationPolicyCommonWithDefaults() *DeviceAuthenticationPolicyCommon`

NewDeviceAuthenticationPolicyCommonWithDefaults instantiates a new DeviceAuthenticationPolicyCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeviceAuthenticationPolicyCommon) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAuthenticationPolicyCommon) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAuthenticationPolicyCommon) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAuthenticationPolicyCommon) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeviceAuthenticationPolicyCommon) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceAuthenticationPolicyCommon) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceAuthenticationPolicyCommon) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceAuthenticationPolicyCommon) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DeviceAuthenticationPolicyCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAuthenticationPolicyCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceAuthenticationPolicyCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthenticationPolicyCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthenticationPolicyCommon) SetName(v string)`

SetName sets Name field to given value.


### GetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyCommon) GetNewDeviceNotification() EnumMFADevicePolicyNewDeviceNotification`

GetNewDeviceNotification returns the NewDeviceNotification field if non-nil, zero value otherwise.

### GetNewDeviceNotificationOk

`func (o *DeviceAuthenticationPolicyCommon) GetNewDeviceNotificationOk() (*EnumMFADevicePolicyNewDeviceNotification, bool)`

GetNewDeviceNotificationOk returns a tuple with the NewDeviceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyCommon) SetNewDeviceNotification(v EnumMFADevicePolicyNewDeviceNotification)`

SetNewDeviceNotification sets NewDeviceNotification field to given value.

### HasNewDeviceNotification

`func (o *DeviceAuthenticationPolicyCommon) HasNewDeviceNotification() bool`

HasNewDeviceNotification returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeviceAuthenticationPolicyCommon) GetAuthentication() DeviceAuthenticationPolicyCommonAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeviceAuthenticationPolicyCommon) GetAuthenticationOk() (*DeviceAuthenticationPolicyCommonAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeviceAuthenticationPolicyCommon) SetAuthentication(v DeviceAuthenticationPolicyCommonAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeviceAuthenticationPolicyCommon) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSms

`func (o *DeviceAuthenticationPolicyCommon) GetSms() DeviceAuthenticationPolicyOfflineDevice`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *DeviceAuthenticationPolicyCommon) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *DeviceAuthenticationPolicyCommon) SetSms(v DeviceAuthenticationPolicyOfflineDevice)`

SetSms sets Sms field to given value.


### GetVoice

`func (o *DeviceAuthenticationPolicyCommon) GetVoice() DeviceAuthenticationPolicyOfflineDevice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *DeviceAuthenticationPolicyCommon) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *DeviceAuthenticationPolicyCommon) SetVoice(v DeviceAuthenticationPolicyOfflineDevice)`

SetVoice sets Voice field to given value.


### GetEmail

`func (o *DeviceAuthenticationPolicyCommon) GetEmail() DeviceAuthenticationPolicyOfflineDevice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceAuthenticationPolicyCommon) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceAuthenticationPolicyCommon) SetEmail(v DeviceAuthenticationPolicyOfflineDevice)`

SetEmail sets Email field to given value.


### GetFido2

`func (o *DeviceAuthenticationPolicyCommon) GetFido2() DeviceAuthenticationPolicyCommonFido2`

GetFido2 returns the Fido2 field if non-nil, zero value otherwise.

### GetFido2Ok

`func (o *DeviceAuthenticationPolicyCommon) GetFido2Ok() (*DeviceAuthenticationPolicyCommonFido2, bool)`

GetFido2Ok returns a tuple with the Fido2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2

`func (o *DeviceAuthenticationPolicyCommon) SetFido2(v DeviceAuthenticationPolicyCommonFido2)`

SetFido2 sets Fido2 field to given value.

### HasFido2

`func (o *DeviceAuthenticationPolicyCommon) HasFido2() bool`

HasFido2 returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicyCommon) GetMobile() DeviceAuthenticationPolicyCommonMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicyCommon) GetMobileOk() (*DeviceAuthenticationPolicyCommonMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicyCommon) SetMobile(v DeviceAuthenticationPolicyCommonMobile)`

SetMobile sets Mobile field to given value.


### GetTotp

`func (o *DeviceAuthenticationPolicyCommon) GetTotp() DeviceAuthenticationPolicyCommonTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicyCommon) GetTotpOk() (*DeviceAuthenticationPolicyCommonTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicyCommon) SetTotp(v DeviceAuthenticationPolicyCommonTotp)`

SetTotp sets Totp field to given value.


### GetDefault

`func (o *DeviceAuthenticationPolicyCommon) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeviceAuthenticationPolicyCommon) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeviceAuthenticationPolicyCommon) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyCommon) GetForSignOnPolicy() bool`

GetForSignOnPolicy returns the ForSignOnPolicy field if non-nil, zero value otherwise.

### GetForSignOnPolicyOk

`func (o *DeviceAuthenticationPolicyCommon) GetForSignOnPolicyOk() (*bool, bool)`

GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyCommon) SetForSignOnPolicy(v bool)`

SetForSignOnPolicy sets ForSignOnPolicy field to given value.


### GetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyCommon) GetIgnoreUserLock() bool`

GetIgnoreUserLock returns the IgnoreUserLock field if non-nil, zero value otherwise.

### GetIgnoreUserLockOk

`func (o *DeviceAuthenticationPolicyCommon) GetIgnoreUserLockOk() (*bool, bool)`

GetIgnoreUserLockOk returns a tuple with the IgnoreUserLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyCommon) SetIgnoreUserLock(v bool)`

SetIgnoreUserLock sets IgnoreUserLock field to given value.

### HasIgnoreUserLock

`func (o *DeviceAuthenticationPolicyCommon) HasIgnoreUserLock() bool`

HasIgnoreUserLock returns a boolean if a field has been set.

### GetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyCommon) GetNotificationsPolicy() DeviceAuthenticationPolicyCommonNotificationsPolicy`

GetNotificationsPolicy returns the NotificationsPolicy field if non-nil, zero value otherwise.

### GetNotificationsPolicyOk

`func (o *DeviceAuthenticationPolicyCommon) GetNotificationsPolicyOk() (*DeviceAuthenticationPolicyCommonNotificationsPolicy, bool)`

GetNotificationsPolicyOk returns a tuple with the NotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyCommon) SetNotificationsPolicy(v DeviceAuthenticationPolicyCommonNotificationsPolicy)`

SetNotificationsPolicy sets NotificationsPolicy field to given value.

### HasNotificationsPolicy

`func (o *DeviceAuthenticationPolicyCommon) HasNotificationsPolicy() bool`

HasNotificationsPolicy returns a boolean if a field has been set.

### GetOathToken

`func (o *DeviceAuthenticationPolicyCommon) GetOathToken() DeviceAuthenticationPolicyOathToken`

GetOathToken returns the OathToken field if non-nil, zero value otherwise.

### GetOathTokenOk

`func (o *DeviceAuthenticationPolicyCommon) GetOathTokenOk() (*DeviceAuthenticationPolicyOathToken, bool)`

GetOathTokenOk returns a tuple with the OathToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOathToken

`func (o *DeviceAuthenticationPolicyCommon) SetOathToken(v DeviceAuthenticationPolicyOathToken)`

SetOathToken sets OathToken field to given value.

### HasOathToken

`func (o *DeviceAuthenticationPolicyCommon) HasOathToken() bool`

HasOathToken returns a boolean if a field has been set.

### GetRememberMe

`func (o *DeviceAuthenticationPolicyCommon) GetRememberMe() DeviceAuthenticationPolicyCommonRememberMe`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *DeviceAuthenticationPolicyCommon) GetRememberMeOk() (*DeviceAuthenticationPolicyCommonRememberMe, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *DeviceAuthenticationPolicyCommon) SetRememberMe(v DeviceAuthenticationPolicyCommonRememberMe)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *DeviceAuthenticationPolicyCommon) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeviceAuthenticationPolicyCommon) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthenticationPolicyCommon) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthenticationPolicyCommon) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthenticationPolicyCommon) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


