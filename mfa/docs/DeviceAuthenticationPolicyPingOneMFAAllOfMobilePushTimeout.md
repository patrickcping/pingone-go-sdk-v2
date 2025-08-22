# DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The amount of time a user has to respond to a push notification before it expires. Minimum is 40 seconds and maximum is 150 seconds. If this parameter is not provided, the duration is set to 40 seconds. | [default to 40]
**TimeUnit** | [**EnumTimeUnitPushTimeout**](EnumTimeUnitPushTimeout.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout(duration int32, timeUnit EnumTimeUnitPushTimeout, ) *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeoutWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeoutWithDefaults() *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeoutWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) GetTimeUnit() EnumTimeUnitPushTimeout`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) GetTimeUnitOk() (*EnumTimeUnitPushTimeout, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout) SetTimeUnit(v EnumTimeUnitPushTimeout)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


