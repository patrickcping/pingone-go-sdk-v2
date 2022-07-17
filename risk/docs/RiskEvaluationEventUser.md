# RiskEvaluationEventUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that specifies the ID of the user associated with the event (maximum size 1024 characters). This is a required property. | 
**Name** | Pointer to **string** | A string that specifies the name of the user associated with the event (maximum size 1024 characters). | [optional] 
**Type** | [**EnumUserType**](EnumUserType.md) |  | 
**Groups** | Pointer to [**[]RiskEvaluationEventUserGroupsInner**](RiskEvaluationEventUserGroupsInner.md) | An array of group names. | [optional] 

## Methods

### NewRiskEvaluationEventUser

`func NewRiskEvaluationEventUser(id string, type_ EnumUserType, ) *RiskEvaluationEventUser`

NewRiskEvaluationEventUser instantiates a new RiskEvaluationEventUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationEventUserWithDefaults

`func NewRiskEvaluationEventUserWithDefaults() *RiskEvaluationEventUser`

NewRiskEvaluationEventUserWithDefaults instantiates a new RiskEvaluationEventUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskEvaluationEventUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskEvaluationEventUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskEvaluationEventUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RiskEvaluationEventUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskEvaluationEventUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskEvaluationEventUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskEvaluationEventUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RiskEvaluationEventUser) GetType() EnumUserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskEvaluationEventUser) GetTypeOk() (*EnumUserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskEvaluationEventUser) SetType(v EnumUserType)`

SetType sets Type field to given value.


### GetGroups

`func (o *RiskEvaluationEventUser) GetGroups() []RiskEvaluationEventUserGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RiskEvaluationEventUser) GetGroupsOk() (*[]RiskEvaluationEventUserGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RiskEvaluationEventUser) SetGroups(v []RiskEvaluationEventUserGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RiskEvaluationEventUser) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


