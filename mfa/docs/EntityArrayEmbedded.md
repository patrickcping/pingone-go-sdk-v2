# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushCredentials** | Pointer to [**[]EntityArrayEmbeddedPushCredentialsInner**](EntityArrayEmbeddedPushCredentialsInner.md) |  | [optional] 
**DeviceAuthenticationPolicies** | Pointer to [**[]DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPushCredentials

`func (o *EntityArrayEmbedded) GetPushCredentials() []EntityArrayEmbeddedPushCredentialsInner`

GetPushCredentials returns the PushCredentials field if non-nil, zero value otherwise.

### GetPushCredentialsOk

`func (o *EntityArrayEmbedded) GetPushCredentialsOk() (*[]EntityArrayEmbeddedPushCredentialsInner, bool)`

GetPushCredentialsOk returns a tuple with the PushCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCredentials

`func (o *EntityArrayEmbedded) SetPushCredentials(v []EntityArrayEmbeddedPushCredentialsInner)`

SetPushCredentials sets PushCredentials field to given value.

### HasPushCredentials

`func (o *EntityArrayEmbedded) HasPushCredentials() bool`

HasPushCredentials returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicies

`func (o *EntityArrayEmbedded) GetDeviceAuthenticationPolicies() []DeviceAuthenticationPolicy`

GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPoliciesOk

`func (o *EntityArrayEmbedded) GetDeviceAuthenticationPoliciesOk() (*[]DeviceAuthenticationPolicy, bool)`

GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicies

`func (o *EntityArrayEmbedded) SetDeviceAuthenticationPolicies(v []DeviceAuthenticationPolicy)`

SetDeviceAuthenticationPolicies sets DeviceAuthenticationPolicies field to given value.

### HasDeviceAuthenticationPolicies

`func (o *EntityArrayEmbedded) HasDeviceAuthenticationPolicies() bool`

HasDeviceAuthenticationPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


