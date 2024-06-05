# ResourceApplicationPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action associated with this permission. | 
**Description** | Pointer to **string** | The resource&#39;s description. | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] 
**Key** | Pointer to **string** | The resource.name:action pair of the permission. | [optional] [readonly] 
**Resource** | Pointer to [**ResourceApplicationPermissionResource**](ResourceApplicationPermissionResource.md) |  | [optional] 
**ResourceServer** | Pointer to [**ResourceApplicationPermissionResourceServer**](ResourceApplicationPermissionResourceServer.md) |  | [optional] 

## Methods

### NewResourceApplicationPermission

`func NewResourceApplicationPermission(action string, ) *ResourceApplicationPermission`

NewResourceApplicationPermission instantiates a new ResourceApplicationPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceApplicationPermissionWithDefaults

`func NewResourceApplicationPermissionWithDefaults() *ResourceApplicationPermission`

NewResourceApplicationPermissionWithDefaults instantiates a new ResourceApplicationPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ResourceApplicationPermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResourceApplicationPermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResourceApplicationPermission) SetAction(v string)`

SetAction sets Action field to given value.


### GetDescription

`func (o *ResourceApplicationPermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceApplicationPermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceApplicationPermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceApplicationPermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceApplicationPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceApplicationPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceApplicationPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceApplicationPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ResourceApplicationPermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ResourceApplicationPermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ResourceApplicationPermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ResourceApplicationPermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetResource

`func (o *ResourceApplicationPermission) GetResource() ResourceApplicationPermissionResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceApplicationPermission) GetResourceOk() (*ResourceApplicationPermissionResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceApplicationPermission) SetResource(v ResourceApplicationPermissionResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceApplicationPermission) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceServer

`func (o *ResourceApplicationPermission) GetResourceServer() ResourceApplicationPermissionResourceServer`

GetResourceServer returns the ResourceServer field if non-nil, zero value otherwise.

### GetResourceServerOk

`func (o *ResourceApplicationPermission) GetResourceServerOk() (*ResourceApplicationPermissionResourceServer, bool)`

GetResourceServerOk returns a tuple with the ResourceServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceServer

`func (o *ResourceApplicationPermission) SetResourceServer(v ResourceApplicationPermissionResourceServer)`

SetResourceServer sets ResourceServer field to given value.

### HasResourceServer

`func (o *ResourceApplicationPermission) HasResourceServer() bool`

HasResourceServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


