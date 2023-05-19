# RiskPolicyConditionAggregatedScoresInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Text that identifies a specific risk predictor in the environment. It uses the form &#x60;${details.xxxxxxx.level}&#x60;, where the string between &#x60;details&#x60; and &#x60;level&#x60; is the compact name of the relevant predictor. | 
**Score** | **int32** | The score you want to assign to the predictor when it is determined that there is a high risk for the predictor. Value should be between 0 and 100. If it is determined that there is medium risk, the predictor will automatically be assigned a score equal to half of the value you specified for high risk. | 

## Methods

### NewRiskPolicyConditionAggregatedScoresInner

`func NewRiskPolicyConditionAggregatedScoresInner(value string, score int32, ) *RiskPolicyConditionAggregatedScoresInner`

NewRiskPolicyConditionAggregatedScoresInner instantiates a new RiskPolicyConditionAggregatedScoresInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyConditionAggregatedScoresInnerWithDefaults

`func NewRiskPolicyConditionAggregatedScoresInnerWithDefaults() *RiskPolicyConditionAggregatedScoresInner`

NewRiskPolicyConditionAggregatedScoresInnerWithDefaults instantiates a new RiskPolicyConditionAggregatedScoresInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RiskPolicyConditionAggregatedScoresInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPolicyConditionAggregatedScoresInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPolicyConditionAggregatedScoresInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetScore

`func (o *RiskPolicyConditionAggregatedScoresInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskPolicyConditionAggregatedScoresInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskPolicyConditionAggregatedScoresInner) SetScore(v int32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


