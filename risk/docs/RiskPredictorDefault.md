# RiskPredictorDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **int32** | An integer type. This specifies the weight assigned to the risk predictor in a new policy by default. | 
**Result** | [**RiskPredictorDefaultResult**](RiskPredictorDefaultResult.md) |  | 

## Methods

### NewRiskPredictorDefault

`func NewRiskPredictorDefault(weight int32, result RiskPredictorDefaultResult, ) *RiskPredictorDefault`

NewRiskPredictorDefault instantiates a new RiskPredictorDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorDefaultWithDefaults

`func NewRiskPredictorDefaultWithDefaults() *RiskPredictorDefault`

NewRiskPredictorDefaultWithDefaults instantiates a new RiskPredictorDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *RiskPredictorDefault) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RiskPredictorDefault) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RiskPredictorDefault) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetResult

`func (o *RiskPredictorDefault) GetResult() RiskPredictorDefaultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskPredictorDefault) GetResultOk() (*RiskPredictorDefaultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskPredictorDefault) SetResult(v RiskPredictorDefaultResult)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


