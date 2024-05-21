# ApplicationRoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the API server operation. This is randomly generated when the operation is created. | [optional] 
**Role** | Pointer to **map[string]interface{}** | The role associated with the role assignment. | [optional] 
**Subject** | Pointer to **map[string]interface{}** | The user associated with the role assignment. | [optional] 

## Methods

### NewApplicationRoleAssignment

`func NewApplicationRoleAssignment() *ApplicationRoleAssignment`

NewApplicationRoleAssignment instantiates a new ApplicationRoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRoleAssignmentWithDefaults

`func NewApplicationRoleAssignmentWithDefaults() *ApplicationRoleAssignment`

NewApplicationRoleAssignmentWithDefaults instantiates a new ApplicationRoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationRoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *ApplicationRoleAssignment) GetRole() map[string]interface{}`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApplicationRoleAssignment) GetRoleOk() (*map[string]interface{}, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApplicationRoleAssignment) SetRole(v map[string]interface{})`

SetRole sets Role field to given value.

### HasRole

`func (o *ApplicationRoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSubject

`func (o *ApplicationRoleAssignment) GetSubject() map[string]interface{}`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ApplicationRoleAssignment) GetSubjectOk() (*map[string]interface{}, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ApplicationRoleAssignment) SetSubject(v map[string]interface{})`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ApplicationRoleAssignment) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


