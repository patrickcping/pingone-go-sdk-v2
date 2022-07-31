# DeviceAuthenticationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | Device authentication policy&#39;s UUID. | [optional] [readonly] 
**Name** | Pointer to **string** | Device authentication policy&#39;s name. | [optional] 
**Sms** | Pointer to [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | [optional] 
**Voice** | Pointer to [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | [optional] 
**Email** | Pointer to [**DeviceAuthenticationPolicyOfflineDevice**](DeviceAuthenticationPolicyOfflineDevice.md) |  | [optional] 
**Mobile** | Pointer to [**DeviceAuthenticationPolicyMobile**](DeviceAuthenticationPolicyMobile.md) |  | [optional] 
**Totp** | Pointer to [**DeviceAuthenticationPolicyTotp**](DeviceAuthenticationPolicyTotp.md) |  | [optional] 
**SecurityKey** | Pointer to [**DeviceAuthenticationPolicySecurityKey**](DeviceAuthenticationPolicySecurityKey.md) |  | [optional] 
**Platform** | Pointer to [**DeviceAuthenticationPolicyPlatform**](DeviceAuthenticationPolicyPlatform.md) |  | [optional] 
**Default** | Pointer to **bool** | The default policy for Flow Manager. | [optional] 
**ForSignOnPolicy** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewDeviceAuthenticationPolicy

`func NewDeviceAuthenticationPolicy() *DeviceAuthenticationPolicy`

NewDeviceAuthenticationPolicy instantiates a new DeviceAuthenticationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyWithDefaults

`func NewDeviceAuthenticationPolicyWithDefaults() *DeviceAuthenticationPolicy`

NewDeviceAuthenticationPolicyWithDefaults instantiates a new DeviceAuthenticationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DeviceAuthenticationPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DeviceAuthenticationPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DeviceAuthenticationPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DeviceAuthenticationPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *DeviceAuthenticationPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAuthenticationPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DeviceAuthenticationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAuthenticationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAuthenticationPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAuthenticationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSms

`func (o *DeviceAuthenticationPolicy) GetSms() DeviceAuthenticationPolicyOfflineDevice`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *DeviceAuthenticationPolicy) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *DeviceAuthenticationPolicy) SetSms(v DeviceAuthenticationPolicyOfflineDevice)`

SetSms sets Sms field to given value.

### HasSms

`func (o *DeviceAuthenticationPolicy) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *DeviceAuthenticationPolicy) GetVoice() DeviceAuthenticationPolicyOfflineDevice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *DeviceAuthenticationPolicy) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *DeviceAuthenticationPolicy) SetVoice(v DeviceAuthenticationPolicyOfflineDevice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *DeviceAuthenticationPolicy) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetEmail

`func (o *DeviceAuthenticationPolicy) GetEmail() DeviceAuthenticationPolicyOfflineDevice`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceAuthenticationPolicy) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceAuthenticationPolicy) SetEmail(v DeviceAuthenticationPolicyOfflineDevice)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DeviceAuthenticationPolicy) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetMobile

`func (o *DeviceAuthenticationPolicy) GetMobile() DeviceAuthenticationPolicyMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *DeviceAuthenticationPolicy) GetMobileOk() (*DeviceAuthenticationPolicyMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *DeviceAuthenticationPolicy) SetMobile(v DeviceAuthenticationPolicyMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *DeviceAuthenticationPolicy) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTotp

`func (o *DeviceAuthenticationPolicy) GetTotp() DeviceAuthenticationPolicyTotp`

GetTotp returns the Totp field if non-nil, zero value otherwise.

### GetTotpOk

`func (o *DeviceAuthenticationPolicy) GetTotpOk() (*DeviceAuthenticationPolicyTotp, bool)`

GetTotpOk returns a tuple with the Totp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotp

`func (o *DeviceAuthenticationPolicy) SetTotp(v DeviceAuthenticationPolicyTotp)`

SetTotp sets Totp field to given value.

### HasTotp

`func (o *DeviceAuthenticationPolicy) HasTotp() bool`

HasTotp returns a boolean if a field has been set.

### GetSecurityKey

`func (o *DeviceAuthenticationPolicy) GetSecurityKey() DeviceAuthenticationPolicySecurityKey`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *DeviceAuthenticationPolicy) GetSecurityKeyOk() (*DeviceAuthenticationPolicySecurityKey, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *DeviceAuthenticationPolicy) SetSecurityKey(v DeviceAuthenticationPolicySecurityKey)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *DeviceAuthenticationPolicy) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceAuthenticationPolicy) GetPlatform() DeviceAuthenticationPolicyPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceAuthenticationPolicy) GetPlatformOk() (*DeviceAuthenticationPolicyPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceAuthenticationPolicy) SetPlatform(v DeviceAuthenticationPolicyPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceAuthenticationPolicy) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDefault

`func (o *DeviceAuthenticationPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *DeviceAuthenticationPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *DeviceAuthenticationPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *DeviceAuthenticationPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetForSignOnPolicy

`func (o *DeviceAuthenticationPolicy) GetForSignOnPolicy() bool`

GetForSignOnPolicy returns the ForSignOnPolicy field if non-nil, zero value otherwise.

### GetForSignOnPolicyOk

`func (o *DeviceAuthenticationPolicy) GetForSignOnPolicyOk() (*bool, bool)`

GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSignOnPolicy

`func (o *DeviceAuthenticationPolicy) SetForSignOnPolicy(v bool)`

SetForSignOnPolicy sets ForSignOnPolicy field to given value.

### HasForSignOnPolicy

`func (o *DeviceAuthenticationPolicy) HasForSignOnPolicy() bool`

HasForSignOnPolicy returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeviceAuthenticationPolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceAuthenticationPolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceAuthenticationPolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceAuthenticationPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


