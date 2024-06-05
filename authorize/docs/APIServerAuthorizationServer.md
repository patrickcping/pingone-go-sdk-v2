# APIServerAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**APIServerAuthorizationServerResource**](APIServerAuthorizationServerResource.md) |  | [optional] 
**Type** | [**EnumAPIServerAuthorizationServerType**](EnumAPIServerAuthorizationServerType.md) |  | 

## Methods

### NewAPIServerAuthorizationServer

`func NewAPIServerAuthorizationServer(type_ EnumAPIServerAuthorizationServerType, ) *APIServerAuthorizationServer`

NewAPIServerAuthorizationServer instantiates a new APIServerAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerAuthorizationServerWithDefaults

`func NewAPIServerAuthorizationServerWithDefaults() *APIServerAuthorizationServer`

NewAPIServerAuthorizationServerWithDefaults instantiates a new APIServerAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *APIServerAuthorizationServer) GetResource() APIServerAuthorizationServerResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *APIServerAuthorizationServer) GetResourceOk() (*APIServerAuthorizationServerResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *APIServerAuthorizationServer) SetResource(v APIServerAuthorizationServerResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *APIServerAuthorizationServer) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetType

`func (o *APIServerAuthorizationServer) GetType() EnumAPIServerAuthorizationServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *APIServerAuthorizationServer) GetTypeOk() (*EnumAPIServerAuthorizationServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *APIServerAuthorizationServer) SetType(v EnumAPIServerAuthorizationServerType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


