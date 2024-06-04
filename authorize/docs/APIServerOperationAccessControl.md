# APIServerOperationAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**APIServerOperationAccessControlGroup**](APIServerOperationAccessControlGroup.md) |  | [optional] 
**Permission** | Pointer to [**APIServerOperationAccessControlPermission**](APIServerOperationAccessControlPermission.md) |  | [optional] 
**Scope** | Pointer to [**APIServerOperationAccessControlScope**](APIServerOperationAccessControlScope.md) |  | [optional] 

## Methods

### NewAPIServerOperationAccessControl

`func NewAPIServerOperationAccessControl() *APIServerOperationAccessControl`

NewAPIServerOperationAccessControl instantiates a new APIServerOperationAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationAccessControlWithDefaults

`func NewAPIServerOperationAccessControlWithDefaults() *APIServerOperationAccessControl`

NewAPIServerOperationAccessControlWithDefaults instantiates a new APIServerOperationAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *APIServerOperationAccessControl) GetGroup() APIServerOperationAccessControlGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *APIServerOperationAccessControl) GetGroupOk() (*APIServerOperationAccessControlGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *APIServerOperationAccessControl) SetGroup(v APIServerOperationAccessControlGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *APIServerOperationAccessControl) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPermission

`func (o *APIServerOperationAccessControl) GetPermission() APIServerOperationAccessControlPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *APIServerOperationAccessControl) GetPermissionOk() (*APIServerOperationAccessControlPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *APIServerOperationAccessControl) SetPermission(v APIServerOperationAccessControlPermission)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *APIServerOperationAccessControl) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetScope

`func (o *APIServerOperationAccessControl) GetScope() APIServerOperationAccessControlScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *APIServerOperationAccessControl) GetScopeOk() (*APIServerOperationAccessControlScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *APIServerOperationAccessControl) SetScope(v APIServerOperationAccessControlScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *APIServerOperationAccessControl) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


