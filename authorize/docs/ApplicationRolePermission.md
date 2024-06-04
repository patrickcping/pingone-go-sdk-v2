# ApplicationRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the application resource permission to associate with this role. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Action** | Pointer to **string** | The action associated with this permission. | [optional] [readonly] 
**Resource** | Pointer to [**ApplicationRolePermissionResource**](ApplicationRolePermissionResource.md) |  | [optional] 

## Methods

### NewApplicationRolePermission

`func NewApplicationRolePermission(id string, ) *ApplicationRolePermission`

NewApplicationRolePermission instantiates a new ApplicationRolePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRolePermissionWithDefaults

`func NewApplicationRolePermissionWithDefaults() *ApplicationRolePermission`

NewApplicationRolePermissionWithDefaults instantiates a new ApplicationRolePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRolePermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRolePermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRolePermission) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *ApplicationRolePermission) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationRolePermission) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationRolePermission) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationRolePermission) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetKey

`func (o *ApplicationRolePermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationRolePermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationRolePermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApplicationRolePermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRolePermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRolePermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRolePermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRolePermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *ApplicationRolePermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApplicationRolePermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApplicationRolePermission) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ApplicationRolePermission) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *ApplicationRolePermission) GetResource() ApplicationRolePermissionResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplicationRolePermission) GetResourceOk() (*ApplicationRolePermissionResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplicationRolePermission) SetResource(v ApplicationRolePermissionResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApplicationRolePermission) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


