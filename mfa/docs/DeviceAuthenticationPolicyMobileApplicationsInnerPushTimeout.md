# DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The amount of time a user has to respond to a push notification before it expires. Minimum is 40 seconds and maximum is 150 seconds. If this parameter is not provided, the duration is set to 40 seconds. | [default to 40]
**TimeUnit** | [**EnumTimeUnitPushTimeout**](EnumTimeUnitPushTimeout.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout(duration int32, timeUnit EnumTimeUnitPushTimeout, ) *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout`

NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeoutWithDefaults

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeoutWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout`

NewDeviceAuthenticationPolicyMobileApplicationsInnerPushTimeoutWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetTimeUnit() EnumTimeUnitPushTimeout`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) GetTimeUnitOk() (*EnumTimeUnitPushTimeout, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) SetTimeUnit(v EnumTimeUnitPushTimeout)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


