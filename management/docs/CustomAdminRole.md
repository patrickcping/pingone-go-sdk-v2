# CustomAdminRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**ApplicableTo** | Pointer to [**[]EnumCustomAdminRoleApplicableTo**](EnumCustomAdminRoleApplicableTo.md) | The scope types to which the role can be applied. | [optional] 
**CanAssign** | Pointer to [**[]CustomAdminRoleCanAssignInner**](CustomAdminRoleCanAssignInner.md) | A relationship that specifies if an actor is assigned the current custom role for a jurisdiction, then the actor can assign any of this set of roles to another actor for the same jurisdiction or sub-jurisdiction. This capability is dreived from the canBeAssignedBy property. | [optional] [readonly] 
**CanBeAssignedBy** | [**[]CustomAdminRoleCanAssignInner**](CustomAdminRoleCanAssignInner.md) | A relationship that determines whether a user assigned to one of this set of roles for a jurisdiction can assign the current custom role to another user for the same jurisdiction or sub-jurisdiction. | 
**Description** | Pointer to **string** | The description of the role. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | The role name. | 
**Permissions** | [**[]CustomAdminRolePermissionsInner**](CustomAdminRolePermissionsInner.md) | The set of permissions assigned to the role. | 
**Type** | Pointer to [**EnumCustomAdminRoleType**](EnumCustomAdminRoleType.md) |  | [optional] 

## Methods

### NewCustomAdminRole

`func NewCustomAdminRole(canBeAssignedBy []CustomAdminRoleCanAssignInner, name string, permissions []CustomAdminRolePermissionsInner, ) *CustomAdminRole`

NewCustomAdminRole instantiates a new CustomAdminRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAdminRoleWithDefaults

`func NewCustomAdminRoleWithDefaults() *CustomAdminRole`

NewCustomAdminRoleWithDefaults instantiates a new CustomAdminRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CustomAdminRole) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomAdminRole) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomAdminRole) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomAdminRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplicableTo

`func (o *CustomAdminRole) GetApplicableTo() []EnumCustomAdminRoleApplicableTo`

GetApplicableTo returns the ApplicableTo field if non-nil, zero value otherwise.

### GetApplicableToOk

`func (o *CustomAdminRole) GetApplicableToOk() (*[]EnumCustomAdminRoleApplicableTo, bool)`

GetApplicableToOk returns a tuple with the ApplicableTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableTo

`func (o *CustomAdminRole) SetApplicableTo(v []EnumCustomAdminRoleApplicableTo)`

SetApplicableTo sets ApplicableTo field to given value.

### HasApplicableTo

`func (o *CustomAdminRole) HasApplicableTo() bool`

HasApplicableTo returns a boolean if a field has been set.

### GetCanAssign

`func (o *CustomAdminRole) GetCanAssign() []CustomAdminRoleCanAssignInner`

GetCanAssign returns the CanAssign field if non-nil, zero value otherwise.

### GetCanAssignOk

`func (o *CustomAdminRole) GetCanAssignOk() (*[]CustomAdminRoleCanAssignInner, bool)`

GetCanAssignOk returns a tuple with the CanAssign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssign

`func (o *CustomAdminRole) SetCanAssign(v []CustomAdminRoleCanAssignInner)`

SetCanAssign sets CanAssign field to given value.

### HasCanAssign

`func (o *CustomAdminRole) HasCanAssign() bool`

HasCanAssign returns a boolean if a field has been set.

### GetCanBeAssignedBy

`func (o *CustomAdminRole) GetCanBeAssignedBy() []CustomAdminRoleCanAssignInner`

GetCanBeAssignedBy returns the CanBeAssignedBy field if non-nil, zero value otherwise.

### GetCanBeAssignedByOk

`func (o *CustomAdminRole) GetCanBeAssignedByOk() (*[]CustomAdminRoleCanAssignInner, bool)`

GetCanBeAssignedByOk returns a tuple with the CanBeAssignedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanBeAssignedBy

`func (o *CustomAdminRole) SetCanBeAssignedBy(v []CustomAdminRoleCanAssignInner)`

SetCanBeAssignedBy sets CanBeAssignedBy field to given value.


### GetDescription

`func (o *CustomAdminRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomAdminRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomAdminRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomAdminRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *CustomAdminRole) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CustomAdminRole) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CustomAdminRole) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CustomAdminRole) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CustomAdminRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomAdminRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomAdminRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomAdminRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomAdminRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAdminRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAdminRole) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CustomAdminRole) GetPermissions() []CustomAdminRolePermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CustomAdminRole) GetPermissionsOk() (*[]CustomAdminRolePermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CustomAdminRole) SetPermissions(v []CustomAdminRolePermissionsInner)`

SetPermissions sets Permissions field to given value.


### GetType

`func (o *CustomAdminRole) GetType() EnumCustomAdminRoleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomAdminRole) GetTypeOk() (*EnumCustomAdminRoleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomAdminRole) SetType(v EnumCustomAdminRoleType)`

SetType sets Type field to given value.

### HasType

`func (o *CustomAdminRole) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


