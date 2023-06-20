# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PushCredentials** | Pointer to [**[]MFAPushCredentialResponse**](MFAPushCredentialResponse.md) |  | [optional] 
**DeviceAuthenticationPolicies** | Pointer to [**[]DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md) |  | [optional] 
**FidoDevicesMetadata** | Pointer to **[]map[string]interface{}** |  | [optional] 
**FidoPolicies** | Pointer to [**[]FIDOPolicy**](FIDOPolicy.md) |  | [optional] 

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

`func (o *EntityArrayEmbedded) GetPushCredentials() []MFAPushCredentialResponse`

GetPushCredentials returns the PushCredentials field if non-nil, zero value otherwise.

### GetPushCredentialsOk

`func (o *EntityArrayEmbedded) GetPushCredentialsOk() (*[]MFAPushCredentialResponse, bool)`

GetPushCredentialsOk returns a tuple with the PushCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushCredentials

`func (o *EntityArrayEmbedded) SetPushCredentials(v []MFAPushCredentialResponse)`

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

### GetFidoDevicesMetadata

`func (o *EntityArrayEmbedded) GetFidoDevicesMetadata() []map[string]interface{}`

GetFidoDevicesMetadata returns the FidoDevicesMetadata field if non-nil, zero value otherwise.

### GetFidoDevicesMetadataOk

`func (o *EntityArrayEmbedded) GetFidoDevicesMetadataOk() (*[]map[string]interface{}, bool)`

GetFidoDevicesMetadataOk returns a tuple with the FidoDevicesMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFidoDevicesMetadata

`func (o *EntityArrayEmbedded) SetFidoDevicesMetadata(v []map[string]interface{})`

SetFidoDevicesMetadata sets FidoDevicesMetadata field to given value.

### HasFidoDevicesMetadata

`func (o *EntityArrayEmbedded) HasFidoDevicesMetadata() bool`

HasFidoDevicesMetadata returns a boolean if a field has been set.

### GetFidoPolicies

`func (o *EntityArrayEmbedded) GetFidoPolicies() []FIDOPolicy`

GetFidoPolicies returns the FidoPolicies field if non-nil, zero value otherwise.

### GetFidoPoliciesOk

`func (o *EntityArrayEmbedded) GetFidoPoliciesOk() (*[]FIDOPolicy, bool)`

GetFidoPoliciesOk returns a tuple with the FidoPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFidoPolicies

`func (o *EntityArrayEmbedded) SetFidoPolicies(v []FIDOPolicy)`

SetFidoPolicies sets FidoPolicies field to given value.

### HasFidoPolicies

`func (o *EntityArrayEmbedded) HasFidoPolicies() bool`

HasFidoPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


