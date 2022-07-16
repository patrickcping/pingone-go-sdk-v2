# RiskEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string that specifies the risk evaluation result type. Options are VALUE. | [optional] [readonly] 
**Level** | Pointer to **string** | A string that specifies the risk evaluation result level. Options are HIGH, MEDIUM, and LOW. | [optional] [readonly] 
**Value** | Pointer to **string** | A string that specifies any custom attribute the administrator defines. | [optional] [readonly] 

## Methods

### NewRiskEvaluationResult

`func NewRiskEvaluationResult() *RiskEvaluationResult`

NewRiskEvaluationResult instantiates a new RiskEvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationResultWithDefaults

`func NewRiskEvaluationResultWithDefaults() *RiskEvaluationResult`

NewRiskEvaluationResultWithDefaults instantiates a new RiskEvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RiskEvaluationResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskEvaluationResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskEvaluationResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskEvaluationResult) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLevel

`func (o *RiskEvaluationResult) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationResult) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationResult) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationResult) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetValue

`func (o *RiskEvaluationResult) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskEvaluationResult) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskEvaluationResult) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskEvaluationResult) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


