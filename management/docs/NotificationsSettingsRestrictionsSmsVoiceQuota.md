# NotificationsSettingsRestrictionsSmsVoiceQuota

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daily** | **int32** | The maximum number of SMS and voice notifications that can be sent per user per day. - &#x60;restrictions.smsVoiceQuota.daily&#x60; can be set to any value between 0 and 50. - Trial accounts have a default value of 30. - The daily counters are reset every night at midnight UTC.  | 

## Methods

### NewNotificationsSettingsRestrictionsSmsVoiceQuota

`func NewNotificationsSettingsRestrictionsSmsVoiceQuota(daily int32, ) *NotificationsSettingsRestrictionsSmsVoiceQuota`

NewNotificationsSettingsRestrictionsSmsVoiceQuota instantiates a new NotificationsSettingsRestrictionsSmsVoiceQuota object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsRestrictionsSmsVoiceQuotaWithDefaults

`func NewNotificationsSettingsRestrictionsSmsVoiceQuotaWithDefaults() *NotificationsSettingsRestrictionsSmsVoiceQuota`

NewNotificationsSettingsRestrictionsSmsVoiceQuotaWithDefaults instantiates a new NotificationsSettingsRestrictionsSmsVoiceQuota object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaily

`func (o *NotificationsSettingsRestrictionsSmsVoiceQuota) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *NotificationsSettingsRestrictionsSmsVoiceQuota) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *NotificationsSettingsRestrictionsSmsVoiceQuota) SetDaily(v int32)`

SetDaily sets Daily field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


