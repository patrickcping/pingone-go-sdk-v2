# DeviceAuthenticationPolicyCommonRememberMeWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Set to &#x60;true&#x60; if you want the MFA policy to include a \&quot;remember me\&quot; option. | 
**LifeTime** | [**DeviceAuthenticationPolicyCommonRememberMeWebLifeTime**](DeviceAuthenticationPolicyCommonRememberMeWebLifeTime.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyCommonRememberMeWeb

`func NewDeviceAuthenticationPolicyCommonRememberMeWeb(enabled bool, lifeTime DeviceAuthenticationPolicyCommonRememberMeWebLifeTime, ) *DeviceAuthenticationPolicyCommonRememberMeWeb`

NewDeviceAuthenticationPolicyCommonRememberMeWeb instantiates a new DeviceAuthenticationPolicyCommonRememberMeWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonRememberMeWebWithDefaults

`func NewDeviceAuthenticationPolicyCommonRememberMeWebWithDefaults() *DeviceAuthenticationPolicyCommonRememberMeWeb`

NewDeviceAuthenticationPolicyCommonRememberMeWebWithDefaults instantiates a new DeviceAuthenticationPolicyCommonRememberMeWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLifeTime

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) GetLifeTime() DeviceAuthenticationPolicyCommonRememberMeWebLifeTime`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) GetLifeTimeOk() (*DeviceAuthenticationPolicyCommonRememberMeWebLifeTime, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *DeviceAuthenticationPolicyCommonRememberMeWeb) SetLifeTime(v DeviceAuthenticationPolicyCommonRememberMeWebLifeTime)`

SetLifeTime sets LifeTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


