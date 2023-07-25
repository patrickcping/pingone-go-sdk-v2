# RiskPolicyConditionAggregatedWeightsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | Text that identifies a specific risk predictor in the environment. It uses the form &#x60;${details.xxxxxxx.level}&#x60;, where the string between &#x60;details&#x60; and &#x60;level&#x60; is the compact name of the relevant predictor. | 
**Weight** | **int32** | The score you want to assign to the predictor when it is determined that there is a high risk for the predictor. Value should be between 0 and 100. If it is determined that there is medium risk, the predictor will automatically be assigned a score equal to half of the value you specified for high risk. | 

## Methods

### NewRiskPolicyConditionAggregatedWeightsInner

`func NewRiskPolicyConditionAggregatedWeightsInner(value string, weight int32, ) *RiskPolicyConditionAggregatedWeightsInner`

NewRiskPolicyConditionAggregatedWeightsInner instantiates a new RiskPolicyConditionAggregatedWeightsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyConditionAggregatedWeightsInnerWithDefaults

`func NewRiskPolicyConditionAggregatedWeightsInnerWithDefaults() *RiskPolicyConditionAggregatedWeightsInner`

NewRiskPolicyConditionAggregatedWeightsInnerWithDefaults instantiates a new RiskPolicyConditionAggregatedWeightsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RiskPolicyConditionAggregatedWeightsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPolicyConditionAggregatedWeightsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPolicyConditionAggregatedWeightsInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetWeight

`func (o *RiskPolicyConditionAggregatedWeightsInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RiskPolicyConditionAggregatedWeightsInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RiskPolicyConditionAggregatedWeightsInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


