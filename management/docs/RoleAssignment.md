# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Gateway** | Pointer to [**GatewayInstanceGateway**](GatewayInstanceGateway.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the user role assignment ID. | [optional] [readonly] 
**ReadOnly** | Pointer to **bool** | A boolean that specifies whether this role assignment can be deleted by the current actor. | [optional] [readonly] 
**Role** | [**RoleAssignmentRole**](RoleAssignmentRole.md) |  | 
**Scope** | [**RoleAssignmentScope**](RoleAssignmentScope.md) |  | 

## Methods

### NewRoleAssignment

`func NewRoleAssignment(role RoleAssignmentRole, scope RoleAssignmentScope, ) *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RoleAssignment) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleAssignment) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleAssignment) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleAssignment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *RoleAssignment) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RoleAssignment) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RoleAssignment) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RoleAssignment) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGateway

`func (o *RoleAssignment) GetGateway() GatewayInstanceGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RoleAssignment) GetGatewayOk() (*GatewayInstanceGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RoleAssignment) SetGateway(v GatewayInstanceGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *RoleAssignment) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetId

`func (o *RoleAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoleAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReadOnly

`func (o *RoleAssignment) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *RoleAssignment) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *RoleAssignment) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *RoleAssignment) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetRole

`func (o *RoleAssignment) GetRole() RoleAssignmentRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleAssignment) GetRoleOk() (*RoleAssignmentRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleAssignment) SetRole(v RoleAssignmentRole)`

SetRole sets Role field to given value.


### GetScope

`func (o *RoleAssignment) GetScope() RoleAssignmentScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RoleAssignment) GetScopeOk() (*RoleAssignmentScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RoleAssignment) SetScope(v RoleAssignmentScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


