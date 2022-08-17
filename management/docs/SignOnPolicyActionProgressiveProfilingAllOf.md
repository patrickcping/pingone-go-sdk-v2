# SignOnPolicyActionProgressiveProfilingAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | [**[]SignOnPolicyActionProgressiveProfilingAllOfAttributes**](SignOnPolicyActionProgressiveProfilingAllOfAttributes.md) |  | 
**PreventMultiplePromptsPerFlow** | **bool** | A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required. | 
**PromptIntervalSeconds** | **int32** | An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required. | 
**PromptText** | **string** | A string that specifies text to display to the user when prompting for attribute values. This property is required. | 

## Methods

### NewSignOnPolicyActionProgressiveProfilingAllOf

`func NewSignOnPolicyActionProgressiveProfilingAllOf(attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string, ) *SignOnPolicyActionProgressiveProfilingAllOf`

NewSignOnPolicyActionProgressiveProfilingAllOf instantiates a new SignOnPolicyActionProgressiveProfilingAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionProgressiveProfilingAllOfWithDefaults

`func NewSignOnPolicyActionProgressiveProfilingAllOfWithDefaults() *SignOnPolicyActionProgressiveProfilingAllOf`

NewSignOnPolicyActionProgressiveProfilingAllOfWithDefaults instantiates a new SignOnPolicyActionProgressiveProfilingAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetAttributes() []SignOnPolicyActionProgressiveProfilingAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetAttributesOk() (*[]SignOnPolicyActionProgressiveProfilingAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetAttributes(v []SignOnPolicyActionProgressiveProfilingAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPreventMultiplePromptsPerFlow() bool`

GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field if non-nil, zero value otherwise.

### GetPreventMultiplePromptsPerFlowOk

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPreventMultiplePromptsPerFlowOk() (*bool, bool)`

GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPreventMultiplePromptsPerFlow(v bool)`

SetPreventMultiplePromptsPerFlow sets PreventMultiplePromptsPerFlow field to given value.


### GetPromptIntervalSeconds

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptIntervalSeconds() int32`

GetPromptIntervalSeconds returns the PromptIntervalSeconds field if non-nil, zero value otherwise.

### GetPromptIntervalSecondsOk

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptIntervalSecondsOk() (*int32, bool)`

GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptIntervalSeconds

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPromptIntervalSeconds(v int32)`

SetPromptIntervalSeconds sets PromptIntervalSeconds field to given value.


### GetPromptText

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptText() string`

GetPromptText returns the PromptText field if non-nil, zero value otherwise.

### GetPromptTextOk

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) GetPromptTextOk() (*string, bool)`

GetPromptTextOk returns a tuple with the PromptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptText

`func (o *SignOnPolicyActionProgressiveProfilingAllOf) SetPromptText(v string)`

SetPromptText sets PromptText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


