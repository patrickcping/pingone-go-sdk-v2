# EntityArrayEmbeddedRolesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**ApplicableTo** | Pointer to [**[]EnumRoleAssignmentScopeType**](EnumRoleAssignmentScopeType.md) | A set of strings that specifies the scopes to which the role applies. | [optional] [readonly] 
**Description** | Pointer to **string** | Specifies the description of the application role assigned to the user. | [optional] 
**Id** | **string** | Specifies the application role ID to assign to the specified user. | 
**Name** | Pointer to **string** | Specifies the name of the application role assigned to the user. | [optional] 
**Permissions** | Pointer to [**[]RolePermissionsInner**](RolePermissionsInner.md) | A set of permissions assigned to the role. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 

## Methods

### NewEntityArrayEmbeddedRolesInner

`func NewEntityArrayEmbeddedRolesInner(id string, ) *EntityArrayEmbeddedRolesInner`

NewEntityArrayEmbeddedRolesInner instantiates a new EntityArrayEmbeddedRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedRolesInnerWithDefaults

`func NewEntityArrayEmbeddedRolesInnerWithDefaults() *EntityArrayEmbeddedRolesInner`

NewEntityArrayEmbeddedRolesInnerWithDefaults instantiates a new EntityArrayEmbeddedRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntityArrayEmbeddedRolesInner) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntityArrayEmbeddedRolesInner) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntityArrayEmbeddedRolesInner) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntityArrayEmbeddedRolesInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplicableTo

`func (o *EntityArrayEmbeddedRolesInner) GetApplicableTo() []EnumRoleAssignmentScopeType`

GetApplicableTo returns the ApplicableTo field if non-nil, zero value otherwise.

### GetApplicableToOk

`func (o *EntityArrayEmbeddedRolesInner) GetApplicableToOk() (*[]EnumRoleAssignmentScopeType, bool)`

GetApplicableToOk returns a tuple with the ApplicableTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableTo

`func (o *EntityArrayEmbeddedRolesInner) SetApplicableTo(v []EnumRoleAssignmentScopeType)`

SetApplicableTo sets ApplicableTo field to given value.

### HasApplicableTo

`func (o *EntityArrayEmbeddedRolesInner) HasApplicableTo() bool`

HasApplicableTo returns a boolean if a field has been set.

### GetDescription

`func (o *EntityArrayEmbeddedRolesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedRolesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedRolesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedRolesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EntityArrayEmbeddedRolesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedRolesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedRolesInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntityArrayEmbeddedRolesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedRolesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedRolesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityArrayEmbeddedRolesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *EntityArrayEmbeddedRolesInner) GetPermissions() []RolePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EntityArrayEmbeddedRolesInner) GetPermissionsOk() (*[]RolePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EntityArrayEmbeddedRolesInner) SetPermissions(v []RolePermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EntityArrayEmbeddedRolesInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetEnvironment

`func (o *EntityArrayEmbeddedRolesInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedRolesInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedRolesInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedRolesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


