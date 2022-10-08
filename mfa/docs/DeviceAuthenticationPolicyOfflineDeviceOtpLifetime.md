# DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The duration (number of time units) that the passcode is valid before it expires. | 
**TimeUnit** | [**EnumTimeUnit**](EnumTimeUnit.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(duration int32, timeUnit EnumTimeUnit, ) *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime`

NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTimeWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTimeWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime`

NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTimeWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


