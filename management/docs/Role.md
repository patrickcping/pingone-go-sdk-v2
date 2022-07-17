# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicableTo** | Pointer to **[]string** | A string that specifies the scope to which the role applies. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the role. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the ID of the role. | [optional] [readonly] 
**Name** | Pointer to [**EnumRoleName**](EnumRoleName.md) |  | [optional] 
**Permissions** | Pointer to [**[]RolePermissionsInner**](RolePermissionsInner.md) | A string that specifies the set of permissions assigned to the role. | [optional] [readonly] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicableTo

`func (o *Role) GetApplicableTo() []string`

GetApplicableTo returns the ApplicableTo field if non-nil, zero value otherwise.

### GetApplicableToOk

`func (o *Role) GetApplicableToOk() (*[]string, bool)`

GetApplicableToOk returns a tuple with the ApplicableTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableTo

`func (o *Role) SetApplicableTo(v []string)`

SetApplicableTo sets ApplicableTo field to given value.

### HasApplicableTo

`func (o *Role) HasApplicableTo() bool`

HasApplicableTo returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Role) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Role) GetName() EnumRoleName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*EnumRoleName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v EnumRoleName)`

SetName sets Name field to given value.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *Role) GetPermissions() []RolePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Role) GetPermissionsOk() (*[]RolePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Role) SetPermissions(v []RolePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Role) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


