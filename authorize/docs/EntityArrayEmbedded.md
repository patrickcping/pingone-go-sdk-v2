# EntityArrayEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecisionEndpoints** | Pointer to [**[]DecisionEndpoint**](DecisionEndpoint.md) |  | [optional] 
**ApiServers** | Pointer to [**[]APIServer**](APIServer.md) |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


