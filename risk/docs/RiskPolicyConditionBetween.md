# RiskPolicyConditionBetween

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinScore** | **int32** | Required for policies of type &#x60;AGGREGATED_SCORES&#x60; or &#x60;AGGREGATED_WEIGHTS&#x60;. The beginning of the risk score range that will be translated into the specified risk level (&#x60;MEDIUM&#x60; or &#x60;HIGH&#x60;). Must be between &#x60;0&#x60; and &#x60;1000&#x60;. | 
**MaxScore** | **int32** | Required for policies of type &#x60;AGGREGATED_SCORES&#x60; or &#x60;AGGREGATED_WEIGHTS&#x60;. The end of the risk score range that will be translated into the specified risk level (&#x60;MEDIUM&#x60; or &#x60;HIGH&#x60;). Must be between &#x60;0&#x60; and &#x60;1000&#x60;. | 

## Methods

### NewRiskPolicyConditionBetween

`func NewRiskPolicyConditionBetween(minScore int32, maxScore int32, ) *RiskPolicyConditionBetween`

NewRiskPolicyConditionBetween instantiates a new RiskPolicyConditionBetween object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyConditionBetweenWithDefaults

`func NewRiskPolicyConditionBetweenWithDefaults() *RiskPolicyConditionBetween`

NewRiskPolicyConditionBetweenWithDefaults instantiates a new RiskPolicyConditionBetween object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinScore

`func (o *RiskPolicyConditionBetween) GetMinScore() int32`

GetMinScore returns the MinScore field if non-nil, zero value otherwise.

### GetMinScoreOk

`func (o *RiskPolicyConditionBetween) GetMinScoreOk() (*int32, bool)`

GetMinScoreOk returns a tuple with the MinScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinScore

`func (o *RiskPolicyConditionBetween) SetMinScore(v int32)`

SetMinScore sets MinScore field to given value.


### GetMaxScore

`func (o *RiskPolicyConditionBetween) GetMaxScore() int32`

GetMaxScore returns the MaxScore field if non-nil, zero value otherwise.

### GetMaxScoreOk

`func (o *RiskPolicyConditionBetween) GetMaxScoreOk() (*int32, bool)`

GetMaxScoreOk returns a tuple with the MaxScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxScore

`func (o *RiskPolicyConditionBetween) SetMaxScore(v int32)`

SetMaxScore sets MaxScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


