# DeviceAuthenticationPolicyPostResponse

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
**Totp** | [**DeviceAuthenticationPolicyCommonTotp**](DeviceAuthenticationPolicyCommonTotp.md) |  | 
**Default** | **bool** | A boolean that specifies whether the policy is the default for the environment. | 
**ForSignOnPolicy** | **bool** |  | [readonly] 
**IgnoreUserLock** | Pointer to **bool** | When applying an MFA policy, PingOne ordinarily checks if a user account is locked, and if so, prevents the user from authenticating. Set &#x60;ignoreUserLock&#x60; to &#x60;true&#x60; if you want PingOne to skip this account lock check. | [optional] 
**NotificationsPolicy** | Pointer to [**DeviceAuthenticationPolicyCommonNotificationsPolicy**](DeviceAuthenticationPolicyCommonNotificationsPolicy.md) |  | [optional] 
**RememberMe** | Pointer to [**DeviceAuthenticationPolicyCommonRememberMe**](DeviceAuthenticationPolicyCommonRememberMe.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Mobile** | [**DeviceAuthenticationPolicyPingIDAllOfMobile**](DeviceAuthenticationPolicyPingIDAllOfMobile.md) |  | 
**Whatsapp** | Pointer to [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | [optional] 
**Desktop** | Pointer to [**DeviceAuthenticationPolicyPingIDDevice**](DeviceAuthenticationPolicyPingIDDevice.md) |  | [optional] 
**Yubikey** | Pointer to [**DeviceAuthenticationPolicyPingIDDevice**](DeviceAuthenticationPolicyPingIDDevice.md) |  | [optional] 
**OathToken** | Pointer to [**DeviceAuthenticationPolicyPingIDDevice**](DeviceAuthenticationPolicyPingIDDevice.md) |  | [optional] 
**Embedded** | Pointer to [**EntityArrayEmbedded**](EntityArrayEmbedded.md) |  | [optional] 
**Count** | Pointer to **float32** |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPostResponse

`func NewDeviceAuthenticationPolicyPostResponse(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, totp DeviceAuthenticationPolicyCommonTotp, default_ bool, forSignOnPolicy bool, mobile DeviceAuthenticationPolicyPingIDAllOfMobile, ) *DeviceAuthenticationPolicyPostResponse`

NewDeviceAuthenticationPolicyPostResponse instantiates a new DeviceAuthenticationPolicyPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPostResponseWithDefaults

`func NewDeviceAuthenticationPolicyPostResponseWithDefaults() *DeviceAuthenticationPolicyPostResponse`

NewDeviceAuthenticationPolicyPostResponseWithDefaults instantiates a new DeviceAuthenticationPolicyPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DeviceAuthenticationPolicyPostResponse) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAuthenticationPolicyPostResponse) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAuthenticationPolicyPostResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *DeviceAuthenticationPolicyPostResponse) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceAuthenticationPolicyPostResponse) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceAuthenticationPolicyPostResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DeviceAuthenticationPolicyPostResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyPostResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAuthenticationPolicyPostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceAuthenticationPolicyPostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthenticationPolicyPostResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPostResponse) GetNewDeviceNotification() EnumMFADevicePolicyNewDeviceNotification`

GetNewDeviceNotification returns the NewDeviceNotification field if non-nil, zero value otherwise.

### GetNewDeviceNotificationOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetNewDeviceNotificationOk() (*EnumMFADevicePolicyNewDeviceNotification, bool)`

GetNewDeviceNotificationOk returns a tuple with the NewDeviceNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPostResponse) SetNewDeviceNotification(v EnumMFADevicePolicyNewDeviceNotification)`

SetNewDeviceNotification sets NewDeviceNotification field to given value.

### HasNewDeviceNotification

`func (o *DeviceAuthenticationPolicyPostResponse) HasNewDeviceNotification() bool`

HasNewDeviceNotification returns a boolean if a field has been set.

### GetAuthentication

`func (o *DeviceAuthenticationPolicyPostResponse) GetAuthentication() DeviceAuthenticationPolicyCommonAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetAuthenticationOk() (*DeviceAuthenticationPolicyCommonAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DeviceAuthenticationPolicyPostResponse) SetAuthentication(v DeviceAuthenticationPolicyCommonAuthentication)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DeviceAuthenticationPolicyPostResponse) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetSms

`func (o *DeviceAuthenticationPolicyPostResponse) GetSms() DeviceAuthenticationPolicyOfflineDevice`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *DeviceAuthenticationPolicyPostResponse) SetSms(v DeviceAuthenticationPolicyOfflineDevice)`

SetSms sets Sms field to given value.


### GetVoice

`func (o *DeviceAuthenticationPolicyPostResponse) GetVoice() DeviceAuthenticationPolicyOfflineDevice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *DeviceAuthenticationPolicyPostResponse) SetVoice(v DeviceAuthenticationPolicyOfflineDevice)`

SetVoice sets Voice field to given value.


### GetEmail

`func (o *DeviceAuthenticationPolicyPostResponse) GetEmail() DeviceAuthenticationPolicyOfflineDevice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceAuthenticationPolicyPostResponse) SetEmail(v DeviceAuthenticationPolicyOfflineDevice)`

SetEmail sets Email field to given value.


### GetFido2

`func (o *DeviceAuthenticationPolicyPostResponse) GetFido2() DeviceAuthenticationPolicyCommonFido2`

GetFido2 returns the Fido2 field if non-nil, zero value otherwise.

### GetFido2Ok

`func (o *DeviceAuthenticationPolicyPostResponse) GetFido2Ok() (*DeviceAuthenticationPolicyCommonFido2, bool)`

GetFido2Ok returns a tuple with the Fido2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2

`func (o *DeviceAuthenticationPolicyPostResponse) SetFido2(v DeviceAuthenticationPolicyCommonFido2)`

SetFido2 sets Fido2 field to given value.

### HasFido2

`func (o *DeviceAuthenticationPolicyPostResponse) HasFido2() bool`

HasFido2 returns a boolean if a field has been set.

### GetTotp

`func (o *DeviceAuthenticationPolicyPostResponse) GetTotp() DeviceAuthenticationPolicyCommonTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetTotpOk() (*DeviceAuthenticationPolicyCommonTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicyPostResponse) SetTotp(v DeviceAuthenticationPolicyCommonTotp)`

SetTotp sets Totp field to given value.


### GetDefault

`func (o *DeviceAuthenticationPolicyPostResponse) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeviceAuthenticationPolicyPostResponse) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPostResponse) GetForSignOnPolicy() bool`

GetForSignOnPolicy returns the ForSignOnPolicy field if non-nil, zero value otherwise.

### GetForSignOnPolicyOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetForSignOnPolicyOk() (*bool, bool)`

GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSignOnPolicy

`func (o *DeviceAuthenticationPolicyPostResponse) SetForSignOnPolicy(v bool)`

SetForSignOnPolicy sets ForSignOnPolicy field to given value.


### GetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPostResponse) GetIgnoreUserLock() bool`

GetIgnoreUserLock returns the IgnoreUserLock field if non-nil, zero value otherwise.

### GetIgnoreUserLockOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetIgnoreUserLockOk() (*bool, bool)`

GetIgnoreUserLockOk returns a tuple with the IgnoreUserLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPostResponse) SetIgnoreUserLock(v bool)`

SetIgnoreUserLock sets IgnoreUserLock field to given value.

### HasIgnoreUserLock

`func (o *DeviceAuthenticationPolicyPostResponse) HasIgnoreUserLock() bool`

HasIgnoreUserLock returns a boolean if a field has been set.

### GetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPostResponse) GetNotificationsPolicy() DeviceAuthenticationPolicyCommonNotificationsPolicy`

GetNotificationsPolicy returns the NotificationsPolicy field if non-nil, zero value otherwise.

### GetNotificationsPolicyOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetNotificationsPolicyOk() (*DeviceAuthenticationPolicyCommonNotificationsPolicy, bool)`

GetNotificationsPolicyOk returns a tuple with the NotificationsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPostResponse) SetNotificationsPolicy(v DeviceAuthenticationPolicyCommonNotificationsPolicy)`

SetNotificationsPolicy sets NotificationsPolicy field to given value.

### HasNotificationsPolicy

`func (o *DeviceAuthenticationPolicyPostResponse) HasNotificationsPolicy() bool`

HasNotificationsPolicy returns a boolean if a field has been set.

### GetRememberMe

`func (o *DeviceAuthenticationPolicyPostResponse) GetRememberMe() DeviceAuthenticationPolicyCommonRememberMe`

GetRememberMe returns the RememberMe field if non-nil, zero value otherwise.

### GetRememberMeOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetRememberMeOk() (*DeviceAuthenticationPolicyCommonRememberMe, bool)`

GetRememberMeOk returns a tuple with the RememberMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMe

`func (o *DeviceAuthenticationPolicyPostResponse) SetRememberMe(v DeviceAuthenticationPolicyCommonRememberMe)`

SetRememberMe sets RememberMe field to given value.

### HasRememberMe

`func (o *DeviceAuthenticationPolicyPostResponse) HasRememberMe() bool`

HasRememberMe returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeviceAuthenticationPolicyPostResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthenticationPolicyPostResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthenticationPolicyPostResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicyPostResponse) GetMobile() DeviceAuthenticationPolicyPingIDAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetMobileOk() (*DeviceAuthenticationPolicyPingIDAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicyPostResponse) SetMobile(v DeviceAuthenticationPolicyPingIDAllOfMobile)`

SetMobile sets Mobile field to given value.


### GetWhatsapp

`func (o *DeviceAuthenticationPolicyPostResponse) GetWhatsapp() DeviceAuthenticationPolicyOfflineDevice`

GetWhatsapp returns the Whatsapp field if non-nil, zero value otherwise.

### GetWhatsappOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetWhatsappOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetWhatsappOk returns a tuple with the Whatsapp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsapp

`func (o *DeviceAuthenticationPolicyPostResponse) SetWhatsapp(v DeviceAuthenticationPolicyOfflineDevice)`

SetWhatsapp sets Whatsapp field to given value.

### HasWhatsapp

`func (o *DeviceAuthenticationPolicyPostResponse) HasWhatsapp() bool`

HasWhatsapp returns a boolean if a field has been set.

### GetDesktop

`func (o *DeviceAuthenticationPolicyPostResponse) GetDesktop() DeviceAuthenticationPolicyPingIDDevice`

GetDesktop returns the Desktop field if non-nil, zero value otherwise.

### GetDesktopOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetDesktopOk() (*DeviceAuthenticationPolicyPingIDDevice, bool)`

GetDesktopOk returns a tuple with the Desktop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktop

`func (o *DeviceAuthenticationPolicyPostResponse) SetDesktop(v DeviceAuthenticationPolicyPingIDDevice)`

SetDesktop sets Desktop field to given value.

### HasDesktop

`func (o *DeviceAuthenticationPolicyPostResponse) HasDesktop() bool`

HasDesktop returns a boolean if a field has been set.

### GetYubikey

`func (o *DeviceAuthenticationPolicyPostResponse) GetYubikey() DeviceAuthenticationPolicyPingIDDevice`

GetYubikey returns the Yubikey field if non-nil, zero value otherwise.

### GetYubikeyOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetYubikeyOk() (*DeviceAuthenticationPolicyPingIDDevice, bool)`

GetYubikeyOk returns a tuple with the Yubikey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikey

`func (o *DeviceAuthenticationPolicyPostResponse) SetYubikey(v DeviceAuthenticationPolicyPingIDDevice)`

SetYubikey sets Yubikey field to given value.

### HasYubikey

`func (o *DeviceAuthenticationPolicyPostResponse) HasYubikey() bool`

HasYubikey returns a boolean if a field has been set.

### GetOathToken

`func (o *DeviceAuthenticationPolicyPostResponse) GetOathToken() DeviceAuthenticationPolicyPingIDDevice`

GetOathToken returns the OathToken field if non-nil, zero value otherwise.

### GetOathTokenOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetOathTokenOk() (*DeviceAuthenticationPolicyPingIDDevice, bool)`

GetOathTokenOk returns a tuple with the OathToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOathToken

`func (o *DeviceAuthenticationPolicyPostResponse) SetOathToken(v DeviceAuthenticationPolicyPingIDDevice)`

SetOathToken sets OathToken field to given value.

### HasOathToken

`func (o *DeviceAuthenticationPolicyPostResponse) HasOathToken() bool`

HasOathToken returns a boolean if a field has been set.

### GetEmbedded

`func (o *DeviceAuthenticationPolicyPostResponse) GetEmbedded() EntityArrayEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetEmbeddedOk() (*EntityArrayEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DeviceAuthenticationPolicyPostResponse) SetEmbedded(v EntityArrayEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DeviceAuthenticationPolicyPostResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetCount

`func (o *DeviceAuthenticationPolicyPostResponse) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceAuthenticationPolicyPostResponse) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeviceAuthenticationPolicyPostResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSize

`func (o *DeviceAuthenticationPolicyPostResponse) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DeviceAuthenticationPolicyPostResponse) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DeviceAuthenticationPolicyPostResponse) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DeviceAuthenticationPolicyPostResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


