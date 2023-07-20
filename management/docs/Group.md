# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | The unique identifier for the group. Search all groups for a specific group ID with a SCIM filter on GET /environments/{environmentID}/groups. Retrieve all the group IDs associated with a user with GET /environments/{environmentID}/users/{userID}?include&#x3D;memberOfGroupIDs. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Population** | Pointer to [**GroupPopulation**](GroupPopulation.md) |  | [optional] 
**Name** | **string** | The group name. A group name can be reused across populations, but the same user cannot belong to multiple groups with the same group name. Population groups cannot share a group name with an environment group. Search all groups for a specific group name with a SCIM filter on GET /environments/{environmentID}/groups. Retrieve all the group names associated with a user with GET /environments/{environmentID}/users/{userID}?include&#x3D;memberOfGroupNames. Use this operation to determine group membership in attribute mappings for claims and assertions. | 
**UserFilter** | Pointer to **string** | A SCIM filter that determines which users are dynamically added to the group. For more information, see Adding users to a group and Removing users from a group. | [optional] 
**Description** | Pointer to **string** | The group description. | [optional] 
**ExternalId** | Pointer to **string** | A user-defined identifier for the group. Use this propertry to syncronize a group in PingOne with the same group in an external system. PingOne does not directly use this property. Search all groups for a specific external ID with a SCIM filter on GET /environments/{environmentID}/groups | [optional] 
**CustomData** | Pointer to **map[string]interface{}** | Optional User-defined custom data. | [optional] 
**DirectMemberCounts** | Pointer to [**GroupDirectMemberCounts**](GroupDirectMemberCounts.md) |  | [optional] 
**TotalMemberCounts** | Pointer to [**GroupTotalMemberCounts**](GroupTotalMemberCounts.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup(name string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Group) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Group) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Group) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Group) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *Group) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Group) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Group) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Group) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetPopulation

`func (o *Group) GetPopulation() GroupPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *Group) GetPopulationOk() (*GroupPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *Group) SetPopulation(v GroupPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *Group) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.


### GetUserFilter

`func (o *Group) GetUserFilter() string`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *Group) GetUserFilterOk() (*string, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *Group) SetUserFilter(v string)`

SetUserFilter sets UserFilter field to given value.

### HasUserFilter

`func (o *Group) HasUserFilter() bool`

HasUserFilter returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalId

`func (o *Group) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Group) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Group) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Group) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetCustomData

`func (o *Group) GetCustomData() map[string]interface{}`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *Group) GetCustomDataOk() (*map[string]interface{}, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *Group) SetCustomData(v map[string]interface{})`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *Group) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### GetDirectMemberCounts

`func (o *Group) GetDirectMemberCounts() GroupDirectMemberCounts`

GetDirectMemberCounts returns the DirectMemberCounts field if non-nil, zero value otherwise.

### GetDirectMemberCountsOk

`func (o *Group) GetDirectMemberCountsOk() (*GroupDirectMemberCounts, bool)`

GetDirectMemberCountsOk returns a tuple with the DirectMemberCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMemberCounts

`func (o *Group) SetDirectMemberCounts(v GroupDirectMemberCounts)`

SetDirectMemberCounts sets DirectMemberCounts field to given value.

### HasDirectMemberCounts

`func (o *Group) HasDirectMemberCounts() bool`

HasDirectMemberCounts returns a boolean if a field has been set.

### GetTotalMemberCounts

`func (o *Group) GetTotalMemberCounts() GroupTotalMemberCounts`

GetTotalMemberCounts returns the TotalMemberCounts field if non-nil, zero value otherwise.

### GetTotalMemberCountsOk

`func (o *Group) GetTotalMemberCountsOk() (*GroupTotalMemberCounts, bool)`

GetTotalMemberCountsOk returns a tuple with the TotalMemberCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemberCounts

`func (o *Group) SetTotalMemberCounts(v GroupTotalMemberCounts)`

SetTotalMemberCounts sets TotalMemberCounts field to given value.

### HasTotalMemberCounts

`func (o *Group) HasTotalMemberCounts() bool`

HasTotalMemberCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


