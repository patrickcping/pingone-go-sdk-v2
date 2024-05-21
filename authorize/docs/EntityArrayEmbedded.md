# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServers** | Pointer to [**[]APIServer**](APIServer.md) |  | [optional] 
**Assignments** | Pointer to [**[]ApplicationRoleAssignment**](ApplicationRoleAssignment.md) |  | [optional] 
**DecisionEndpoints** | Pointer to [**[]DecisionEndpoint**](DecisionEndpoint.md) |  | [optional] 
**Permissions** | Pointer to [**[]EntityArrayEmbeddedPermissionsInner**](EntityArrayEmbeddedPermissionsInner.md) |  | [optional] 
**Resources** | Pointer to [**[]ApplicationResource**](ApplicationResource.md) |  | [optional] 
**Roles** | Pointer to [**[]ApplicationRole**](ApplicationRole.md) |  | [optional] 

## Methods

### NewEntityArrayEmbedded

`func NewEntityArrayEmbedded() *EntityArrayEmbedded`

NewEntityArrayEmbedded instantiates a new EntityArrayEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedWithDefaults

`func NewEntityArrayEmbeddedWithDefaults() *EntityArrayEmbedded`

NewEntityArrayEmbeddedWithDefaults instantiates a new EntityArrayEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServers

`func (o *EntityArrayEmbedded) GetApiServers() []APIServer`

GetApiServers returns the ApiServers field if non-nil, zero value otherwise.

### GetApiServersOk

`func (o *EntityArrayEmbedded) GetApiServersOk() (*[]APIServer, bool)`

GetApiServersOk returns a tuple with the ApiServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServers

`func (o *EntityArrayEmbedded) SetApiServers(v []APIServer)`

SetApiServers sets ApiServers field to given value.

### HasApiServers

`func (o *EntityArrayEmbedded) HasApiServers() bool`

HasApiServers returns a boolean if a field has been set.

### GetAssignments

`func (o *EntityArrayEmbedded) GetAssignments() []ApplicationRoleAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *EntityArrayEmbedded) GetAssignmentsOk() (*[]ApplicationRoleAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *EntityArrayEmbedded) SetAssignments(v []ApplicationRoleAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *EntityArrayEmbedded) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetDecisionEndpoints

`func (o *EntityArrayEmbedded) GetDecisionEndpoints() []DecisionEndpoint`

GetDecisionEndpoints returns the DecisionEndpoints field if non-nil, zero value otherwise.

### GetDecisionEndpointsOk

`func (o *EntityArrayEmbedded) GetDecisionEndpointsOk() (*[]DecisionEndpoint, bool)`

GetDecisionEndpointsOk returns a tuple with the DecisionEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionEndpoints

`func (o *EntityArrayEmbedded) SetDecisionEndpoints(v []DecisionEndpoint)`

SetDecisionEndpoints sets DecisionEndpoints field to given value.

### HasDecisionEndpoints

`func (o *EntityArrayEmbedded) HasDecisionEndpoints() bool`

HasDecisionEndpoints returns a boolean if a field has been set.

### GetPermissions

`func (o *EntityArrayEmbedded) GetPermissions() []EntityArrayEmbeddedPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *EntityArrayEmbedded) GetPermissionsOk() (*[]EntityArrayEmbeddedPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *EntityArrayEmbedded) SetPermissions(v []EntityArrayEmbeddedPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *EntityArrayEmbedded) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetResources

`func (o *EntityArrayEmbedded) GetResources() []ApplicationResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EntityArrayEmbedded) GetResourcesOk() (*[]ApplicationResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EntityArrayEmbedded) SetResources(v []ApplicationResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EntityArrayEmbedded) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRoles

`func (o *EntityArrayEmbedded) GetRoles() []ApplicationRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *EntityArrayEmbedded) GetRolesOk() (*[]ApplicationRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *EntityArrayEmbedded) SetRoles(v []ApplicationRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *EntityArrayEmbedded) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


