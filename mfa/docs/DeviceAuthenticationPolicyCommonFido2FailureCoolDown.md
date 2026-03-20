# DeviceAuthenticationPolicyCommonFido2FailureCoolDown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | Used in conjunction with &#x60;timeUnit&#x60; to specify how long the user should be blocked after reaching the maximum number of failures, before they can try authenticating again. Minimum period is two minutes, maximum is 30 minutes. | [optional] 
**TimeUnit** | Pointer to [**EnumTimeUnit**](EnumTimeUnit.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonFido2FailureCoolDown

`func NewDeviceAuthenticationPolicyCommonFido2FailureCoolDown() *DeviceAuthenticationPolicyCommonFido2FailureCoolDown`

NewDeviceAuthenticationPolicyCommonFido2FailureCoolDown instantiates a new DeviceAuthenticationPolicyCommonFido2FailureCoolDown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonFido2FailureCoolDownWithDefaults

`func NewDeviceAuthenticationPolicyCommonFido2FailureCoolDownWithDefaults() *DeviceAuthenticationPolicyCommonFido2FailureCoolDown`

NewDeviceAuthenticationPolicyCommonFido2FailureCoolDownWithDefaults instantiates a new DeviceAuthenticationPolicyCommonFido2FailureCoolDown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *DeviceAuthenticationPolicyCommonFido2FailureCoolDown) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


