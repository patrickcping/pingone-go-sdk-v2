# DeviceAuthenticationPolicyCommonRememberMeWebLifeTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeUnit** | Pointer to [**EnumTimeUnitRememberMeWebLifeTime**](EnumTimeUnitRememberMeWebLifeTime.md) |  | [optional] 
**Duration** | Pointer to **int32** | Used in conjunction with &#x60;timeUnit&#x60; to define the \&quot;remember me\&quot; period. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTime

`func NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTime() *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime`

NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTime instantiates a new DeviceAuthenticationPolicyCommonRememberMeWebLifeTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTimeWithDefaults

`func NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTimeWithDefaults() *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime`

NewDeviceAuthenticationPolicyCommonRememberMeWebLifeTimeWithDefaults instantiates a new DeviceAuthenticationPolicyCommonRememberMeWebLifeTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) GetTimeUnit() EnumTimeUnitRememberMeWebLifeTime`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) GetTimeUnitOk() (*EnumTimeUnitRememberMeWebLifeTime, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) SetTimeUnit(v EnumTimeUnitRememberMeWebLifeTime)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.

### GetDuration

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DeviceAuthenticationPolicyCommonRememberMeWebLifeTime) HasDuration() bool`

HasDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


