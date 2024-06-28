# IdentityPropagationPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the group. Search all groups for a specific group ID with a SCIM filter on GET /environments/{environmentID}/groups. Retrieve all the group IDs associated with a user with GET /environments/{environmentID}/users/{userID}?include&#x3D;memberOfGroupIDs. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | Unique name of the propagation plan | 
**Status** | Pointer to [**EnumIdentityPropagationPlanStatus**](EnumIdentityPropagationPlanStatus.md) |  | [optional] 

## Methods

### NewIdentityPropagationPlan

`func NewIdentityPropagationPlan(name string, ) *IdentityPropagationPlan`

NewIdentityPropagationPlan instantiates a new IdentityPropagationPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityPropagationPlanWithDefaults

`func NewIdentityPropagationPlanWithDefaults() *IdentityPropagationPlan`

NewIdentityPropagationPlanWithDefaults instantiates a new IdentityPropagationPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IdentityPropagationPlan) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityPropagationPlan) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityPropagationPlan) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityPropagationPlan) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *IdentityPropagationPlan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityPropagationPlan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityPropagationPlan) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityPropagationPlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *IdentityPropagationPlan) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityPropagationPlan) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityPropagationPlan) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityPropagationPlan) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *IdentityPropagationPlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityPropagationPlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityPropagationPlan) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *IdentityPropagationPlan) GetStatus() EnumIdentityPropagationPlanStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityPropagationPlan) GetStatusOk() (*EnumIdentityPropagationPlanStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityPropagationPlan) SetStatus(v EnumIdentityPropagationPlanStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityPropagationPlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


