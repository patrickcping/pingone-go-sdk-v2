# DeviceAuthenticationPolicyPingOneMFA

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

## Methods

### NewDeviceAuthenticationPolicyPingOneMFA

`func NewDeviceAuthenticationPolicyPingOneMFA(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyCommonMobile, totp DeviceAuthenticationPolicyCommonTotp, default_ bool, forSignOnPolicy bool, ) *DeviceAuthenticationPolicyPingOneMFA`

NewDeviceAuthenticationPolicyPingOneMFA instantiates a new DeviceAuthenticationPolicyPingOneMFA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAWithDefaults() *DeviceAuthenticationPolicyPingOneMFA`

NewDeviceAuthenticationPolicyPingOneMFAWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetName(v string)`

SetName sets Name field to given value.


### GetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetNewDeviceNotification() EnumMFADevicePolicyNewDeviceNotification`

GetNewDeviceNotification returns the NewDeviceNotification field if non-nil, zero value otherwise.

### GetNewDeviceNotificationOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetNewDeviceNotificationOk() (*EnumMFADevicePolicyNewDeviceNotification, bool)`

GetNewDeviceNotificationOk returns a tuple with the NewDeviceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetNewDeviceNotification(v EnumMFADevicePolicyNewDeviceNotification)`

SetNewDeviceNotification sets NewDeviceNotification field to given value.

### HasNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasNewDeviceNotification() bool`

HasNewDeviceNotification returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetAuthentication() DeviceAuthenticationPolicyCommonAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetAuthenticationOk() (*DeviceAuthenticationPolicyCommonAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetAuthentication(v DeviceAuthenticationPolicyCommonAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSms

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetSms() DeviceAuthenticationPolicyOfflineDevice`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetSms(v DeviceAuthenticationPolicyOfflineDevice)`

SetSms sets Sms field to given value.


### GetVoice

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetVoice() DeviceAuthenticationPolicyOfflineDevice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetVoice(v DeviceAuthenticationPolicyOfflineDevice)`

SetVoice sets Voice field to given value.


### GetEmail

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetEmail() DeviceAuthenticationPolicyOfflineDevice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetEmail(v DeviceAuthenticationPolicyOfflineDevice)`

SetEmail sets Email field to given value.


### GetFido2

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetFido2() DeviceAuthenticationPolicyCommonFido2`

GetFido2 returns the Fido2 field if non-nil, zero value otherwise.

### GetFido2Ok

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetFido2Ok() (*DeviceAuthenticationPolicyCommonFido2, bool)`

GetFido2Ok returns a tuple with the Fido2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetFido2(v DeviceAuthenticationPolicyCommonFido2)`

SetFido2 sets Fido2 field to given value.

### HasFido2

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasFido2() bool`

HasFido2 returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetMobile() DeviceAuthenticationPolicyCommonMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetMobileOk() (*DeviceAuthenticationPolicyCommonMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetMobile(v DeviceAuthenticationPolicyCommonMobile)`

SetMobile sets Mobile field to given value.


### GetTotp

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetTotp() DeviceAuthenticationPolicyCommonTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetTotpOk() (*DeviceAuthenticationPolicyCommonTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetTotp(v DeviceAuthenticationPolicyCommonTotp)`

SetTotp sets Totp field to given value.


### GetDefault

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetForSignOnPolicy() bool`

GetForSignOnPolicy returns the ForSignOnPolicy field if non-nil, zero value otherwise.

### GetForSignOnPolicyOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetForSignOnPolicyOk() (*bool, bool)`

GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetForSignOnPolicy(v bool)`

SetForSignOnPolicy sets ForSignOnPolicy field to given value.


### GetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetIgnoreUserLock() bool`

GetIgnoreUserLock returns the IgnoreUserLock field if non-nil, zero value otherwise.

### GetIgnoreUserLockOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetIgnoreUserLockOk() (*bool, bool)`

GetIgnoreUserLockOk returns a tuple with the IgnoreUserLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetIgnoreUserLock(v bool)`

SetIgnoreUserLock sets IgnoreUserLock field to given value.

### HasIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasIgnoreUserLock() bool`

HasIgnoreUserLock returns a boolean if a field has been set.

### GetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetNotificationsPolicy() DeviceAuthenticationPolicyCommonNotificationsPolicy`

GetNotificationsPolicy returns the NotificationsPolicy field if non-nil, zero value otherwise.

### GetNotificationsPolicyOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetNotificationsPolicyOk() (*DeviceAuthenticationPolicyCommonNotificationsPolicy, bool)`

GetNotificationsPolicyOk returns a tuple with the NotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetNotificationsPolicy(v DeviceAuthenticationPolicyCommonNotificationsPolicy)`

SetNotificationsPolicy sets NotificationsPolicy field to given value.

### HasNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasNotificationsPolicy() bool`

HasNotificationsPolicy returns a boolean if a field has been set.

### GetOathToken

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetOathToken() DeviceAuthenticationPolicyOathToken`

GetOathToken returns the OathToken field if non-nil, zero value otherwise.

### GetOathTokenOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetOathTokenOk() (*DeviceAuthenticationPolicyOathToken, bool)`

GetOathTokenOk returns a tuple with the OathToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOathToken

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetOathToken(v DeviceAuthenticationPolicyOathToken)`

SetOathToken sets OathToken field to given value.

### HasOathToken

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasOathToken() bool`

HasOathToken returns a boolean if a field has been set.

### GetRememberMe

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetRememberMe() DeviceAuthenticationPolicyCommonRememberMe`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetRememberMeOk() (*DeviceAuthenticationPolicyCommonRememberMe, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetRememberMe(v DeviceAuthenticationPolicyCommonRememberMe)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWhatsapp

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetWhatsapp() DeviceAuthenticationPolicyOfflineDevice`

GetWhatsapp returns the Whatsapp field if non-nil, zero value otherwise.

### GetWhatsappOk

`func (o *DeviceAuthenticationPolicyPingOneMFA) GetWhatsappOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetWhatsappOk returns a tuple with the Whatsapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsapp

`func (o *DeviceAuthenticationPolicyPingOneMFA) SetWhatsapp(v DeviceAuthenticationPolicyOfflineDevice)`

SetWhatsapp sets Whatsapp field to given value.

### HasWhatsapp

`func (o *DeviceAuthenticationPolicyPingOneMFA) HasWhatsapp() bool`

HasWhatsapp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


