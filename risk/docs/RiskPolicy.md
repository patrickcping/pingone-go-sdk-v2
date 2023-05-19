# RiskPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**RiskPolicyCondition**](RiskPolicyCondition.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was first created (format ISO-8061). | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies a description for this risk policy. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies a name for this risk policy. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters. | 
**Priority** | Pointer to **int32** | An integer that specifies priority of the policy inside a risk policy set, designating which policy should run first. This is a read-only value. The priority is determined by the order in which policies are listed in the policy set. The first policy in the list is assigned priority 1 and is evaluated first. The next policy in the list is assigned priority 2 and so on. | [optional] 
**Result** | [**RiskPolicyResult**](RiskPolicyResult.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewRiskPolicy

`func NewRiskPolicy(condition RiskPolicyCondition, name string, result RiskPolicyResult, ) *RiskPolicy`

NewRiskPolicy instantiates a new RiskPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyWithDefaults

`func NewRiskPolicyWithDefaults() *RiskPolicy`

NewRiskPolicyWithDefaults instantiates a new RiskPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RiskPolicy) GetCondition() RiskPolicyCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPolicy) GetConditionOk() (*RiskPolicyCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPolicy) SetCondition(v RiskPolicyCondition)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *RiskPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RiskPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *RiskPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *RiskPolicy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RiskPolicy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RiskPolicy) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RiskPolicy) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResult

`func (o *RiskPolicy) GetResult() RiskPolicyResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskPolicy) GetResultOk() (*RiskPolicyResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskPolicy) SetResult(v RiskPolicyResult)`

SetResult sets Result field to given value.


### GetUpdatedAt

`func (o *RiskPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


