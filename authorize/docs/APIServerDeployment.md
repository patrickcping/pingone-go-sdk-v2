# APIServerDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessControl** | Pointer to [**APIServerDeploymentAccessControl**](APIServerDeploymentAccessControl.md) |  | [optional] 
**AuthorizationVersion** | Pointer to [**APIServerDeploymentAuthorizationVersion**](APIServerDeploymentAuthorizationVersion.md) |  | [optional] 
**DecisionEndpoint** | Pointer to [**APIServerDeploymentDecisionEndpoint**](APIServerDeploymentDecisionEndpoint.md) |  | [optional] 
**DeployedAt** | Pointer to **time.Time** | The time of most recent successful deployment. Null if the API service has never been successfully deployed. | [optional] [readonly] 
**Policy** | Pointer to [**APIServerDeploymentPolicy**](APIServerDeploymentPolicy.md) |  | [optional] 
**Status** | Pointer to [**APIServerDeploymentStatus**](APIServerDeploymentStatus.md) |  | [optional] 

## Methods

### NewAPIServerDeployment

`func NewAPIServerDeployment() *APIServerDeployment`

NewAPIServerDeployment instantiates a new APIServerDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerDeploymentWithDefaults

`func NewAPIServerDeploymentWithDefaults() *APIServerDeployment`

NewAPIServerDeploymentWithDefaults instantiates a new APIServerDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *APIServerDeployment) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *APIServerDeployment) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *APIServerDeployment) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *APIServerDeployment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessControl

`func (o *APIServerDeployment) GetAccessControl() APIServerDeploymentAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *APIServerDeployment) GetAccessControlOk() (*APIServerDeploymentAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *APIServerDeployment) SetAccessControl(v APIServerDeploymentAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *APIServerDeployment) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetAuthorizationVersion

`func (o *APIServerDeployment) GetAuthorizationVersion() APIServerDeploymentAuthorizationVersion`

GetAuthorizationVersion returns the AuthorizationVersion field if non-nil, zero value otherwise.

### GetAuthorizationVersionOk

`func (o *APIServerDeployment) GetAuthorizationVersionOk() (*APIServerDeploymentAuthorizationVersion, bool)`

GetAuthorizationVersionOk returns a tuple with the AuthorizationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationVersion

`func (o *APIServerDeployment) SetAuthorizationVersion(v APIServerDeploymentAuthorizationVersion)`

SetAuthorizationVersion sets AuthorizationVersion field to given value.

### HasAuthorizationVersion

`func (o *APIServerDeployment) HasAuthorizationVersion() bool`

HasAuthorizationVersion returns a boolean if a field has been set.

### GetDecisionEndpoint

`func (o *APIServerDeployment) GetDecisionEndpoint() APIServerDeploymentDecisionEndpoint`

GetDecisionEndpoint returns the DecisionEndpoint field if non-nil, zero value otherwise.

### GetDecisionEndpointOk

`func (o *APIServerDeployment) GetDecisionEndpointOk() (*APIServerDeploymentDecisionEndpoint, bool)`

GetDecisionEndpointOk returns a tuple with the DecisionEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionEndpoint

`func (o *APIServerDeployment) SetDecisionEndpoint(v APIServerDeploymentDecisionEndpoint)`

SetDecisionEndpoint sets DecisionEndpoint field to given value.

### HasDecisionEndpoint

`func (o *APIServerDeployment) HasDecisionEndpoint() bool`

HasDecisionEndpoint returns a boolean if a field has been set.

### GetDeployedAt

`func (o *APIServerDeployment) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *APIServerDeployment) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *APIServerDeployment) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *APIServerDeployment) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetPolicy

`func (o *APIServerDeployment) GetPolicy() APIServerDeploymentPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *APIServerDeployment) GetPolicyOk() (*APIServerDeploymentPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *APIServerDeployment) SetPolicy(v APIServerDeploymentPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *APIServerDeployment) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetStatus

`func (o *APIServerDeployment) GetStatus() APIServerDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIServerDeployment) GetStatusOk() (*APIServerDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIServerDeployment) SetStatus(v APIServerDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *APIServerDeployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


