# DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of consecutive push notifications that can be ignored or rejected by a user within a defined period before push notifications are blocked for the application. The minimum value is 1 and the maximum value is 50. If this parameter is not provided, the default value is 5. | [optional] [default to 5]
**LockDuration** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitLockDuration**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitLockDuration.md) |  | [optional] 
**TimePeriod** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitTimePeriod**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitTimePeriod.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit() *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitWithDefaults() *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLockDuration

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetLockDuration() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitLockDuration`

GetLockDuration returns the LockDuration field if non-nil, zero value otherwise.

### GetLockDurationOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetLockDurationOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitLockDuration, bool)`

GetLockDurationOk returns a tuple with the LockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockDuration

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) SetLockDuration(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitLockDuration)`

SetLockDuration sets LockDuration field to given value.

### HasLockDuration

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) HasLockDuration() bool`

HasLockDuration returns a boolean if a field has been set.

### GetTimePeriod

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetTimePeriod() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitTimePeriod`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) GetTimePeriodOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitTimePeriod, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) SetTimePeriod(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimitTimePeriod)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


