# DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The total amount of time an authentication request notification has to be handled by the user before timing out. This includes both the time until the notification is displayed to the user and the time the user takes to respond. Minimum is 30 seconds, maximum is 90 seconds. totalTimeout.duration must exceed deviceTimeout.duration by at least 15 seconds. | 
**TimeUnit** | [**EnumTimeUnitSeconds**](EnumTimeUnitSeconds.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout(duration int32, timeUnit EnumTimeUnitSeconds, ) *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeoutWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeoutWithDefaults() *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeoutWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) GetTimeUnit() EnumTimeUnitSeconds`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) GetTimeUnitOk() (*EnumTimeUnitSeconds, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfigurationTotalTimeout) SetTimeUnit(v EnumTimeUnitSeconds)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


