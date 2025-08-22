# DeviceAuthenticationPolicyPingIDAllOfMobilePush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies whether push notification is enabled or disabled for the policy. | 
**NumberMatching** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to require the authenticating user to select a number that was displayed to them on the accessing device. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingIDAllOfMobilePush

`func NewDeviceAuthenticationPolicyPingIDAllOfMobilePush(enabled bool, ) *DeviceAuthenticationPolicyPingIDAllOfMobilePush`

NewDeviceAuthenticationPolicyPingIDAllOfMobilePush instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobilePush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDAllOfMobilePushWithDefaults

`func NewDeviceAuthenticationPolicyPingIDAllOfMobilePushWithDefaults() *DeviceAuthenticationPolicyPingIDAllOfMobilePush`

NewDeviceAuthenticationPolicyPingIDAllOfMobilePushWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobilePush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNumberMatching

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) GetNumberMatching() bool`

GetNumberMatching returns the NumberMatching field if non-nil, zero value otherwise.

### GetNumberMatchingOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) GetNumberMatchingOk() (*bool, bool)`

GetNumberMatchingOk returns a tuple with the NumberMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberMatching

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) SetNumberMatching(v bool)`

SetNumberMatching sets NumberMatching field to given value.

### HasNumberMatching

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobilePush) HasNumberMatching() bool`

HasNumberMatching returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


