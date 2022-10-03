# LicenseMfa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPushNotification** | Pointer to **bool** | A read-only boolean that specifies whether push notifications are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value. | [optional] 
**AllowNotificationOutsideWhitelist** | Pointer to **bool** | A read-only boolean that specifies whether the license supports sending notifications outside of the environment&#39;s whitelist. | [optional] 
**AllowFido2Devices** | Pointer to **bool** | A read-only boolean that specifies whether FIDO2 devices are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value. | [optional] 
**AllowVoiceOtp** | Pointer to **bool** |  | [optional] 
**AllowEmailOtp** | Pointer to **bool** |  | [optional] 
**AllowSmsOtp** | Pointer to **bool** |  | [optional] 
**AllowTotp** | Pointer to **bool** |  | [optional] 

## Methods

### NewLicenseMfa

`func NewLicenseMfa() *LicenseMfa`

NewLicenseMfa instantiates a new LicenseMfa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseMfaWithDefaults

`func NewLicenseMfaWithDefaults() *LicenseMfa`

NewLicenseMfaWithDefaults instantiates a new LicenseMfa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPushNotification

`func (o *LicenseMfa) GetAllowPushNotification() bool`

GetAllowPushNotification returns the AllowPushNotification field if non-nil, zero value otherwise.

### GetAllowPushNotificationOk

`func (o *LicenseMfa) GetAllowPushNotificationOk() (*bool, bool)`

GetAllowPushNotificationOk returns a tuple with the AllowPushNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPushNotification

`func (o *LicenseMfa) SetAllowPushNotification(v bool)`

SetAllowPushNotification sets AllowPushNotification field to given value.

### HasAllowPushNotification

`func (o *LicenseMfa) HasAllowPushNotification() bool`

HasAllowPushNotification returns a boolean if a field has been set.

### GetAllowNotificationOutsideWhitelist

`func (o *LicenseMfa) GetAllowNotificationOutsideWhitelist() bool`

GetAllowNotificationOutsideWhitelist returns the AllowNotificationOutsideWhitelist field if non-nil, zero value otherwise.

### GetAllowNotificationOutsideWhitelistOk

`func (o *LicenseMfa) GetAllowNotificationOutsideWhitelistOk() (*bool, bool)`

GetAllowNotificationOutsideWhitelistOk returns a tuple with the AllowNotificationOutsideWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNotificationOutsideWhitelist

`func (o *LicenseMfa) SetAllowNotificationOutsideWhitelist(v bool)`

SetAllowNotificationOutsideWhitelist sets AllowNotificationOutsideWhitelist field to given value.

### HasAllowNotificationOutsideWhitelist

`func (o *LicenseMfa) HasAllowNotificationOutsideWhitelist() bool`

HasAllowNotificationOutsideWhitelist returns a boolean if a field has been set.

### GetAllowFido2Devices

`func (o *LicenseMfa) GetAllowFido2Devices() bool`

GetAllowFido2Devices returns the AllowFido2Devices field if non-nil, zero value otherwise.

### GetAllowFido2DevicesOk

`func (o *LicenseMfa) GetAllowFido2DevicesOk() (*bool, bool)`

GetAllowFido2DevicesOk returns a tuple with the AllowFido2Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFido2Devices

`func (o *LicenseMfa) SetAllowFido2Devices(v bool)`

SetAllowFido2Devices sets AllowFido2Devices field to given value.

### HasAllowFido2Devices

`func (o *LicenseMfa) HasAllowFido2Devices() bool`

HasAllowFido2Devices returns a boolean if a field has been set.

### GetAllowVoiceOtp

`func (o *LicenseMfa) GetAllowVoiceOtp() bool`

GetAllowVoiceOtp returns the AllowVoiceOtp field if non-nil, zero value otherwise.

### GetAllowVoiceOtpOk

`func (o *LicenseMfa) GetAllowVoiceOtpOk() (*bool, bool)`

GetAllowVoiceOtpOk returns a tuple with the AllowVoiceOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVoiceOtp

`func (o *LicenseMfa) SetAllowVoiceOtp(v bool)`

SetAllowVoiceOtp sets AllowVoiceOtp field to given value.

### HasAllowVoiceOtp

`func (o *LicenseMfa) HasAllowVoiceOtp() bool`

HasAllowVoiceOtp returns a boolean if a field has been set.

### GetAllowEmailOtp

`func (o *LicenseMfa) GetAllowEmailOtp() bool`

GetAllowEmailOtp returns the AllowEmailOtp field if non-nil, zero value otherwise.

### GetAllowEmailOtpOk

`func (o *LicenseMfa) GetAllowEmailOtpOk() (*bool, bool)`

GetAllowEmailOtpOk returns a tuple with the AllowEmailOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailOtp

`func (o *LicenseMfa) SetAllowEmailOtp(v bool)`

SetAllowEmailOtp sets AllowEmailOtp field to given value.

### HasAllowEmailOtp

`func (o *LicenseMfa) HasAllowEmailOtp() bool`

HasAllowEmailOtp returns a boolean if a field has been set.

### GetAllowSmsOtp

`func (o *LicenseMfa) GetAllowSmsOtp() bool`

GetAllowSmsOtp returns the AllowSmsOtp field if non-nil, zero value otherwise.

### GetAllowSmsOtpOk

`func (o *LicenseMfa) GetAllowSmsOtpOk() (*bool, bool)`

GetAllowSmsOtpOk returns a tuple with the AllowSmsOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSmsOtp

`func (o *LicenseMfa) SetAllowSmsOtp(v bool)`

SetAllowSmsOtp sets AllowSmsOtp field to given value.

### HasAllowSmsOtp

`func (o *LicenseMfa) HasAllowSmsOtp() bool`

HasAllowSmsOtp returns a boolean if a field has been set.

### GetAllowTotp

`func (o *LicenseMfa) GetAllowTotp() bool`

GetAllowTotp returns the AllowTotp field if non-nil, zero value otherwise.

### GetAllowTotpOk

`func (o *LicenseMfa) GetAllowTotpOk() (*bool, bool)`

GetAllowTotpOk returns a tuple with the AllowTotp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTotp

`func (o *LicenseMfa) SetAllowTotp(v bool)`

SetAllowTotp sets AllowTotp field to given value.

### HasAllowTotp

`func (o *LicenseMfa) HasAllowTotp() bool`

HasAllowTotp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


