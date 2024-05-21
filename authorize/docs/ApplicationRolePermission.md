# ApplicationRolePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the application resource permission to associate with this role. | 
**Permission** | Pointer to [**ApplicationRolePermissionPermission**](ApplicationRolePermissionPermission.md) |  | [optional] 

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


### GetPermission

`func (o *ApplicationRolePermission) GetPermission() ApplicationRolePermissionPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ApplicationRolePermission) GetPermissionOk() (*ApplicationRolePermissionPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ApplicationRolePermission) SetPermission(v ApplicationRolePermissionPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ApplicationRolePermission) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


