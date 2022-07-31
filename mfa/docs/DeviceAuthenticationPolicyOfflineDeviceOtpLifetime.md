# DeviceAuthenticationPolicyOfflineDeviceOtpLifetime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to **int32** | The duration (number of time units) that the passcode is valid before it expires. | [optional] 
**TimeUnit** | Pointer to [**EnumTimeUnit**](EnumTimeUnit.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetime

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetime() *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime`

NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetime instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpLifetime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetimeWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetimeWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime`

NewDeviceAuthenticationPolicyOfflineDeviceOtpLifetimeWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpLifetime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) GetTimeUnit() EnumTimeUnit`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) GetTimeUnitOk() (*EnumTimeUnit, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) SetTimeUnit(v EnumTimeUnit)`

SetTimeUnit sets TimeUnit field to given value.

### HasTimeUnit

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpLifetime) HasTimeUnit() bool`

HasTimeUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


