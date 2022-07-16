# ApplicationAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**ApplicationAccessControlRole**](ApplicationAccessControlRole.md) |  | [optional] 
**Group** | Pointer to [**ApplicationAccessControlGroup**](ApplicationAccessControlGroup.md) |  | [optional] 

## Methods

### NewApplicationAccessControl

`func NewApplicationAccessControl() *ApplicationAccessControl`

NewApplicationAccessControl instantiates a new ApplicationAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAccessControlWithDefaults

`func NewApplicationAccessControlWithDefaults() *ApplicationAccessControl`

NewApplicationAccessControlWithDefaults instantiates a new ApplicationAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ApplicationAccessControl) GetRole() ApplicationAccessControlRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApplicationAccessControl) GetRoleOk() (*ApplicationAccessControlRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApplicationAccessControl) SetRole(v ApplicationAccessControlRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApplicationAccessControl) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetGroup

`func (o *ApplicationAccessControl) GetGroup() ApplicationAccessControlGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApplicationAccessControl) GetGroupOk() (*ApplicationAccessControlGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApplicationAccessControl) SetGroup(v ApplicationAccessControlGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApplicationAccessControl) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


