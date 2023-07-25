# RiskPredictorCommonDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | **int32** | An integer type. This specifies the weight assigned to the risk predictor in a new policy by default. | 
**Score** | Pointer to **int32** |  | [optional] 
**Evaluated** | Pointer to **bool** |  | [optional] 
**Result** | Pointer to [**RiskPredictorCommonDefaultResult**](RiskPredictorCommonDefaultResult.md) |  | [optional] 

## Methods

### NewRiskPredictorCommonDefault

`func NewRiskPredictorCommonDefault(weight int32, ) *RiskPredictorCommonDefault`

NewRiskPredictorCommonDefault instantiates a new RiskPredictorCommonDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCommonDefaultWithDefaults

`func NewRiskPredictorCommonDefaultWithDefaults() *RiskPredictorCommonDefault`

NewRiskPredictorCommonDefaultWithDefaults instantiates a new RiskPredictorCommonDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *RiskPredictorCommonDefault) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *RiskPredictorCommonDefault) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *RiskPredictorCommonDefault) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetScore

`func (o *RiskPredictorCommonDefault) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskPredictorCommonDefault) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskPredictorCommonDefault) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskPredictorCommonDefault) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetEvaluated

`func (o *RiskPredictorCommonDefault) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *RiskPredictorCommonDefault) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *RiskPredictorCommonDefault) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.

### HasEvaluated

`func (o *RiskPredictorCommonDefault) HasEvaluated() bool`

HasEvaluated returns a boolean if a field has been set.

### GetResult

`func (o *RiskPredictorCommonDefault) GetResult() RiskPredictorCommonDefaultResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskPredictorCommonDefault) GetResultOk() (*RiskPredictorCommonDefaultResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskPredictorCommonDefault) SetResult(v RiskPredictorCommonDefaultResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *RiskPredictorCommonDefault) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


