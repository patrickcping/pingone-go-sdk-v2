# APIServerOperationsValueAccessControlGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]APIServerOperationsValueAccessControlGroupGroupsInner**](APIServerOperationsValueAccessControlGroupGroupsInner.md) | An array that specifies the list of groups that define the access requirements for the operation. The end user must be a member of one or more of these groups to gain access to the operation. This is a required property if operations.value.accessControl.group is set. The ID must reference a group that exists at the time the data is persisted. There is no referential integrity between a group and this configuration. If a group is subsequently deleted, the access control configuration will continue to reference that group. | 

## Methods

### NewAPIServerOperationsValueAccessControlGroup

`func NewAPIServerOperationsValueAccessControlGroup(groups []APIServerOperationsValueAccessControlGroupGroupsInner, ) *APIServerOperationsValueAccessControlGroup`

NewAPIServerOperationsValueAccessControlGroup instantiates a new APIServerOperationsValueAccessControlGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationsValueAccessControlGroupWithDefaults

`func NewAPIServerOperationsValueAccessControlGroupWithDefaults() *APIServerOperationsValueAccessControlGroup`

NewAPIServerOperationsValueAccessControlGroupWithDefaults instantiates a new APIServerOperationsValueAccessControlGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *APIServerOperationsValueAccessControlGroup) GetGroups() []APIServerOperationsValueAccessControlGroupGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *APIServerOperationsValueAccessControlGroup) GetGroupsOk() (*[]APIServerOperationsValueAccessControlGroupGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *APIServerOperationsValueAccessControlGroup) SetGroups(v []APIServerOperationsValueAccessControlGroupGroupsInner)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


