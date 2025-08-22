# DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The amount of time an authentication request notification has to reach the device before timing out. Minimum is 15 seconds, maximum is 75 seconds. | 
**TimeUnit** | [**EnumTimeUnitSeconds**](EnumTimeUnitSeconds.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout(duration int32, timeUnit EnumTimeUnitSeconds, ) *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout`

NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeoutWithDefaults

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeoutWithDefaults() *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout`

NewDeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeoutWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) GetTimeUnit() EnumTimeUnitSeconds`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) GetTimeUnitOk() (*EnumTimeUnitSeconds, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfigurationDeviceTimeout) SetTimeUnit(v EnumTimeUnitSeconds)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


