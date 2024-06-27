# GroupMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to [**GroupMembershipEmbedded**](GroupMembershipEmbedded.md) |  | [optional] 
**Id** | **string** | ID of the group to assign | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | Pointer to **string** | A string that specifies the group name assigned to the user. | [optional] [readonly] 
**Type** | Pointer to [**EnumGroupMembershipType**](EnumGroupMembershipType.md) |  | [optional] 

## Methods

### NewGroupMembership

`func NewGroupMembership(id string, ) *GroupMembership`

NewGroupMembership instantiates a new GroupMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMembershipWithDefaults

`func NewGroupMembershipWithDefaults() *GroupMembership`

NewGroupMembershipWithDefaults instantiates a new GroupMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GroupMembership) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupMembership) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupMembership) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupMembership) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *GroupMembership) GetEmbedded() GroupMembershipEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *GroupMembership) GetEmbeddedOk() (*GroupMembershipEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *GroupMembership) SetEmbedded(v GroupMembershipEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *GroupMembership) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetId

`func (o *GroupMembership) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMembership) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMembership) SetId(v string)`

SetId sets Id field to given value.


### GetEnvironment

`func (o *GroupMembership) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GroupMembership) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GroupMembership) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GroupMembership) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *GroupMembership) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupMembership) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupMembership) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupMembership) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GroupMembership) GetType() EnumGroupMembershipType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupMembership) GetTypeOk() (*EnumGroupMembershipType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupMembership) SetType(v EnumGroupMembershipType)`

SetType sets Type field to given value.

### HasType

`func (o *GroupMembership) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


