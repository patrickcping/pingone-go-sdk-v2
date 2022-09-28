# LicenseMfa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPushNotification** | Pointer to **bool** | A read-only boolean that specifies whether push notifications are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value. | [optional] 
**AllowMfaNotificationsOutsideWhitelist** | Pointer to **bool** | A read-only boolean that specifies whether the license supports sending notifications outside of the environment&#39;s whitelist. | [optional] 
**AllowFido2Devices** | Pointer to **bool** | A read-only boolean that specifies whether FIDO2 devices are allowed. For TRIAL (unpaid) licenses, the default value is true. For other license package types, adoption of the feature determines the default value. | [optional] 

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

### GetAllowMfaNotificationsOutsideWhitelist

`func (o *LicenseMfa) GetAllowMfaNotificationsOutsideWhitelist() bool`

GetAllowMfaNotificationsOutsideWhitelist returns the AllowMfaNotificationsOutsideWhitelist field if non-nil, zero value otherwise.

### GetAllowMfaNotificationsOutsideWhitelistOk

`func (o *LicenseMfa) GetAllowMfaNotificationsOutsideWhitelistOk() (*bool, bool)`

GetAllowMfaNotificationsOutsideWhitelistOk returns a tuple with the AllowMfaNotificationsOutsideWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMfaNotificationsOutsideWhitelist

`func (o *LicenseMfa) SetAllowMfaNotificationsOutsideWhitelist(v bool)`

SetAllowMfaNotificationsOutsideWhitelist sets AllowMfaNotificationsOutsideWhitelist field to given value.

### HasAllowMfaNotificationsOutsideWhitelist

`func (o *LicenseMfa) HasAllowMfaNotificationsOutsideWhitelist() bool`

HasAllowMfaNotificationsOutsideWhitelist returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


