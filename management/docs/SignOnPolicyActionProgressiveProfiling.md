# SignOnPolicyActionProgressiveProfiling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Condition** | Pointer to [**SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy assignment resource’s unique identifier. | [optional] [readonly] 
**Priority** | **int32** | An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property. | 
**SignOnPolicy** | Pointer to [**SignOnPolicyActionCommonSignOnPolicy**](SignOnPolicyActionCommonSignOnPolicy.md) |  | [optional] 
**Type** | [**EnumSignOnPolicyType**](EnumSignOnPolicyType.md) |  | 
**Attributes** | [**[]SignOnPolicyActionProgressiveProfilingAllOfAttributes**](SignOnPolicyActionProgressiveProfilingAllOfAttributes.md) |  | 
**PreventMultiplePromptsPerFlow** | **bool** | A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required. | 
**PromptIntervalSeconds** | **int32** | An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required. | 
**PromptText** | **string** | A string that specifies text to display to the user when prompting for attribute values. This property is required. | 

## Methods

### NewSignOnPolicyActionProgressiveProfiling

`func NewSignOnPolicyActionProgressiveProfiling(priority int32, type_ EnumSignOnPolicyType, attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string, ) *SignOnPolicyActionProgressiveProfiling`

NewSignOnPolicyActionProgressiveProfiling instantiates a new SignOnPolicyActionProgressiveProfiling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionProgressiveProfilingWithDefaults

`func NewSignOnPolicyActionProgressiveProfilingWithDefaults() *SignOnPolicyActionProgressiveProfiling`

NewSignOnPolicyActionProgressiveProfilingWithDefaults instantiates a new SignOnPolicyActionProgressiveProfiling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionProgressiveProfiling) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionProgressiveProfiling) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionProgressiveProfiling) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionProgressiveProfiling) GetCondition() SignOnPolicyActionCommonConditionOrOrInner`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionProgressiveProfiling) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionProgressiveProfiling) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionProgressiveProfiling) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionProgressiveProfiling) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionProgressiveProfiling) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionProgressiveProfiling) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionProgressiveProfiling) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionProgressiveProfiling) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionProgressiveProfiling) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionProgressiveProfiling) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionProgressiveProfiling) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionProgressiveProfiling) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionProgressiveProfiling) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionProgressiveProfiling) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionProgressiveProfiling) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SignOnPolicyActionProgressiveProfiling) GetAttributes() []SignOnPolicyActionProgressiveProfilingAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetAttributesOk() (*[]SignOnPolicyActionProgressiveProfilingAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SignOnPolicyActionProgressiveProfiling) SetAttributes(v []SignOnPolicyActionProgressiveProfilingAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyActionProgressiveProfiling) GetPreventMultiplePromptsPerFlow() bool`

GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field if non-nil, zero value otherwise.

### GetPreventMultiplePromptsPerFlowOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetPreventMultiplePromptsPerFlowOk() (*bool, bool)`

GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyActionProgressiveProfiling) SetPreventMultiplePromptsPerFlow(v bool)`

SetPreventMultiplePromptsPerFlow sets PreventMultiplePromptsPerFlow field to given value.


### GetPromptIntervalSeconds

`func (o *SignOnPolicyActionProgressiveProfiling) GetPromptIntervalSeconds() int32`

GetPromptIntervalSeconds returns the PromptIntervalSeconds field if non-nil, zero value otherwise.

### GetPromptIntervalSecondsOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetPromptIntervalSecondsOk() (*int32, bool)`

GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptIntervalSeconds

`func (o *SignOnPolicyActionProgressiveProfiling) SetPromptIntervalSeconds(v int32)`

SetPromptIntervalSeconds sets PromptIntervalSeconds field to given value.


### GetPromptText

`func (o *SignOnPolicyActionProgressiveProfiling) GetPromptText() string`

GetPromptText returns the PromptText field if non-nil, zero value otherwise.

### GetPromptTextOk

`func (o *SignOnPolicyActionProgressiveProfiling) GetPromptTextOk() (*string, bool)`

GetPromptTextOk returns a tuple with the PromptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptText

`func (o *SignOnPolicyActionProgressiveProfiling) SetPromptText(v string)`

SetPromptText sets PromptText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


