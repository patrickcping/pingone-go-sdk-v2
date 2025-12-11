# DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | The amount of time an issued pairing key can be used until it expires. Minimum is 1 minute and maximum is 48 hours. If this parameter is not provided, the duration is set to 10 minutes. | [default to 10]
**TimeUnit** | [**EnumTimeUnitPairingKeyLifetime**](EnumTimeUnitPairingKeyLifetime.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime(duration int32, timeUnit EnumTimeUnitPairingKeyLifetime, ) *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetimeWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetimeWithDefaults() *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetimeWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) GetTimeUnit() EnumTimeUnitPairingKeyLifetime`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) GetTimeUnitOk() (*EnumTimeUnitPairingKeyLifetime, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime) SetTimeUnit(v EnumTimeUnitPairingKeyLifetime)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


