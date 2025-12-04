# DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of consecutive push notifications that can be ignored or rejected by a user within a defined period before push notifications are blocked for the application. The minimum value is 1 and the maximum value is 50. If this parameter is not provided, the default value is 5. | [optional] [default to 5]
**LockDuration** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration.md) |  | [optional] 
**TimePeriod** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitTimePeriod**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitTimePeriod.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit() *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitWithDefaults() *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLockDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetLockDuration() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration`

GetLockDuration returns the LockDuration field if non-nil, zero value otherwise.

### GetLockDurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetLockDurationOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration, bool)`

GetLockDurationOk returns a tuple with the LockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) SetLockDuration(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitLockDuration)`

SetLockDuration sets LockDuration field to given value.

### HasLockDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) HasLockDuration() bool`

HasLockDuration returns a boolean if a field has been set.

### GetTimePeriod

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetTimePeriod() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitTimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) GetTimePeriodOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitTimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) SetTimePeriod(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimitTimePeriod)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


