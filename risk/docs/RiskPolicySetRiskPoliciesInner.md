# RiskPolicySetRiskPoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**RiskPolicySetRiskPoliciesInnerCondition**](RiskPolicySetRiskPoliciesInnerCondition.md) |  | 
**CreatedAt** | Pointer to **string** | The time the resource was first created (format ISO-8061). | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies a description for this risk policy. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies a name for this risk policy. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters. | 
**Priority** | Pointer to **int32** | An integer that specifies priority of the policy inside a risk policy set, designating which policy should run first. This is a read-only value. The priority is determined by the order in which policies are listed in the policy set. The first policy in the list is assigned priority 1 and is evaluated first. The next policy in the list is assigned priority 2 and so on. | [optional] 
**Result** | [**RiskPolicyResult**](RiskPolicyResult.md) |  | 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewRiskPolicySetRiskPoliciesInner

`func NewRiskPolicySetRiskPoliciesInner(condition RiskPolicySetRiskPoliciesInnerCondition, name string, result RiskPolicyResult, ) *RiskPolicySetRiskPoliciesInner`

NewRiskPolicySetRiskPoliciesInner instantiates a new RiskPolicySetRiskPoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetRiskPoliciesInnerWithDefaults

`func NewRiskPolicySetRiskPoliciesInnerWithDefaults() *RiskPolicySetRiskPoliciesInner`

NewRiskPolicySetRiskPoliciesInnerWithDefaults instantiates a new RiskPolicySetRiskPoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RiskPolicySetRiskPoliciesInner) GetCondition() RiskPolicySetRiskPoliciesInnerCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPolicySetRiskPoliciesInner) GetConditionOk() (*RiskPolicySetRiskPoliciesInnerCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPolicySetRiskPoliciesInner) SetCondition(v RiskPolicySetRiskPoliciesInnerCondition)`

SetCondition sets Condition field to given value.


### GetCreatedAt

`func (o *RiskPolicySetRiskPoliciesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskPolicySetRiskPoliciesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskPolicySetRiskPoliciesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskPolicySetRiskPoliciesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *RiskPolicySetRiskPoliciesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskPolicySetRiskPoliciesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskPolicySetRiskPoliciesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskPolicySetRiskPoliciesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskPolicySetRiskPoliciesInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskPolicySetRiskPoliciesInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskPolicySetRiskPoliciesInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskPolicySetRiskPoliciesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *RiskPolicySetRiskPoliciesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskPolicySetRiskPoliciesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskPolicySetRiskPoliciesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskPolicySetRiskPoliciesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskPolicySetRiskPoliciesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskPolicySetRiskPoliciesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskPolicySetRiskPoliciesInner) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *RiskPolicySetRiskPoliciesInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RiskPolicySetRiskPoliciesInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RiskPolicySetRiskPoliciesInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RiskPolicySetRiskPoliciesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetResult

`func (o *RiskPolicySetRiskPoliciesInner) GetResult() RiskPolicyResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskPolicySetRiskPoliciesInner) GetResultOk() (*RiskPolicyResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskPolicySetRiskPoliciesInner) SetResult(v RiskPolicyResult)`

SetResult sets Result field to given value.


### GetUpdatedAt

`func (o *RiskPolicySetRiskPoliciesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskPolicySetRiskPoliciesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskPolicySetRiskPoliciesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskPolicySetRiskPoliciesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


