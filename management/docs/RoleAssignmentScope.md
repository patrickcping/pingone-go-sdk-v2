# RoleAssignmentScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that specifies the role assignment scope ID. | 
**Type** | [**EnumRoleAssignmentScopeType**](EnumRoleAssignmentScopeType.md) |  | 

## Methods

### NewRoleAssignmentScope

`func NewRoleAssignmentScope(id string, type_ EnumRoleAssignmentScopeType, ) *RoleAssignmentScope`

NewRoleAssignmentScope instantiates a new RoleAssignmentScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentScopeWithDefaults

`func NewRoleAssignmentScopeWithDefaults() *RoleAssignmentScope`

NewRoleAssignmentScopeWithDefaults instantiates a new RoleAssignmentScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleAssignmentScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleAssignmentScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleAssignmentScope) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RoleAssignmentScope) GetType() EnumRoleAssignmentScopeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleAssignmentScope) GetTypeOk() (*EnumRoleAssignmentScopeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleAssignmentScope) SetType(v EnumRoleAssignmentScopeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


