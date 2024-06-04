# ResourceApplicationPermissionResourceServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID for the associated application resource server. | [optional] 
**Environment** | Pointer to [**ResourceApplicationPermissionResourceServerEnvironment**](ResourceApplicationPermissionResourceServerEnvironment.md) |  | [optional] 

## Methods

### NewResourceApplicationPermissionResourceServer

`func NewResourceApplicationPermissionResourceServer() *ResourceApplicationPermissionResourceServer`

NewResourceApplicationPermissionResourceServer instantiates a new ResourceApplicationPermissionResourceServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceApplicationPermissionResourceServerWithDefaults

`func NewResourceApplicationPermissionResourceServerWithDefaults() *ResourceApplicationPermissionResourceServer`

NewResourceApplicationPermissionResourceServerWithDefaults instantiates a new ResourceApplicationPermissionResourceServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceApplicationPermissionResourceServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceApplicationPermissionResourceServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceApplicationPermissionResourceServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceApplicationPermissionResourceServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *ResourceApplicationPermissionResourceServer) GetEnvironment() ResourceApplicationPermissionResourceServerEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceApplicationPermissionResourceServer) GetEnvironmentOk() (*ResourceApplicationPermissionResourceServerEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceApplicationPermissionResourceServer) SetEnvironment(v ResourceApplicationPermissionResourceServerEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceApplicationPermissionResourceServer) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


