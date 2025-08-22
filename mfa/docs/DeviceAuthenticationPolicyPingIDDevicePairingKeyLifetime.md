# DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | For \&quot;desktop\&quot; only. The amount of time the pairing key is valid. Can be expressed in minutes or hours. Minimum is one minute, maximum is 48 hours. If the &#x60;pairingKeyLifetime&#x60; object is not provided, then 48 hours is used. | 
**TimeUnit** | [**EnumTimeUnit**](EnumTimeUnit.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime

`func NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime(duration int32, timeUnit EnumTimeUnit, ) *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime`

NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime instantiates a new DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetimeWithDefaults

`func NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetimeWithDefaults() *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime`

NewDeviceAuthenticationPolicyPingIDDevicePairingKeyLifetimeWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetTimeUnit

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


