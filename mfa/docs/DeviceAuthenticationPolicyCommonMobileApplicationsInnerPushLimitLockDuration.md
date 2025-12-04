# DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The length of time that push notifications should be blocked for the application if the defined limit has been reached. The minimum value is 1 minute and the maximum value is 120 minutes. If this parameter is not provided, the default value is 30 minutes. | [default to 30]
**TimeUnit** | [**EnumTimeUnit**](EnumTimeUnit.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration(duration int32, timeUnit EnumTimeUnit, ) *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDurationWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDurationWithDefaults() *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDurationWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


