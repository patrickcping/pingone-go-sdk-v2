# UserApplicationRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | **string** | Specifies the application role ID to assign to the specified user. | 
**Name** | Pointer to **string** | Specifies the name of the application role assigned to the user. | [optional] 
**Description** | Pointer to **string** | Specifies the description of the application role assigned to the user. | [optional] 

## Methods

### NewUserApplicationRoleAssignment

`func NewUserApplicationRoleAssignment(id string, ) *UserApplicationRoleAssignment`

NewUserApplicationRoleAssignment instantiates a new UserApplicationRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserApplicationRoleAssignmentWithDefaults

`func NewUserApplicationRoleAssignmentWithDefaults() *UserApplicationRoleAssignment`

NewUserApplicationRoleAssignmentWithDefaults instantiates a new UserApplicationRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *UserApplicationRoleAssignment) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UserApplicationRoleAssignment) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UserApplicationRoleAssignment) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UserApplicationRoleAssignment) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *UserApplicationRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserApplicationRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserApplicationRoleAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UserApplicationRoleAssignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserApplicationRoleAssignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserApplicationRoleAssignment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserApplicationRoleAssignment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UserApplicationRoleAssignment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserApplicationRoleAssignment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserApplicationRoleAssignment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserApplicationRoleAssignment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


