# MFASettingsLockout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureCount** | **int32** | An integer that defines the maximum number of incorrect authentication attempts before the account is locked. | 
**DurationSeconds** | **int32** | An integer that defines the number of seconds to keep the account in a locked state. | 

## Methods

### NewMFASettingsLockout

`func NewMFASettingsLockout(failureCount int32, durationSeconds int32, ) *MFASettingsLockout`

NewMFASettingsLockout instantiates a new MFASettingsLockout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFASettingsLockoutWithDefaults

`func NewMFASettingsLockoutWithDefaults() *MFASettingsLockout`

NewMFASettingsLockoutWithDefaults instantiates a new MFASettingsLockout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureCount

`func (o *MFASettingsLockout) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *MFASettingsLockout) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *MFASettingsLockout) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.


### GetDurationSeconds

`func (o *MFASettingsLockout) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *MFASettingsLockout) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *MFASettingsLockout) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


