# ApplicationRolePermissionPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the permission resource associated with a specified role. | [optional] 
**Action** | Pointer to **string** | The action associated with this permission. | [optional] 
**Resource** | Pointer to [**ApplicationRolePermissionPermissionResource**](ApplicationRolePermissionPermissionResource.md) |  | [optional] 

## Methods

### NewApplicationRolePermissionPermission

`func NewApplicationRolePermissionPermission() *ApplicationRolePermissionPermission`

NewApplicationRolePermissionPermission instantiates a new ApplicationRolePermissionPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRolePermissionPermissionWithDefaults

`func NewApplicationRolePermissionPermissionWithDefaults() *ApplicationRolePermissionPermission`

NewApplicationRolePermissionPermissionWithDefaults instantiates a new ApplicationRolePermissionPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRolePermissionPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRolePermissionPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRolePermissionPermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRolePermissionPermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *ApplicationRolePermissionPermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApplicationRolePermissionPermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApplicationRolePermissionPermission) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ApplicationRolePermissionPermission) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResource

`func (o *ApplicationRolePermissionPermission) GetResource() ApplicationRolePermissionPermissionResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplicationRolePermissionPermission) GetResourceOk() (*ApplicationRolePermissionPermissionResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplicationRolePermissionPermission) SetResource(v ApplicationRolePermissionPermissionResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApplicationRolePermissionPermission) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


