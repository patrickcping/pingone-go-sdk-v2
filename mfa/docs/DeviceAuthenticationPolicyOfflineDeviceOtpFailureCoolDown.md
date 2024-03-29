# DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The duration (number of time units) the user is blocked after reaching the maximum number of passcode failures. | 
**TimeUnit** | [**EnumTimeUnit**](EnumTimeUnit.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(duration int32, timeUnit EnumTimeUnit, ) *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown`

NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDownWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDownWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown`

NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDownWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


