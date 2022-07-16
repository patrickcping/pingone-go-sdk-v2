# RiskEvaluationDetailsUserVelocityByIpVelocity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistinctCount** | Pointer to **int32** | Relevant only for predictors having a measure value of \&quot;DISTINCT_COUNT\&quot;. This is the number of users that accessed the IP in the past during seconds. | [optional] 
**During** | Pointer to **int32** | The interval (in seconds) during which the velocity is calculated. | [optional] 

## Methods

### NewRiskEvaluationDetailsUserVelocityByIpVelocity

`func NewRiskEvaluationDetailsUserVelocityByIpVelocity() *RiskEvaluationDetailsUserVelocityByIpVelocity`

NewRiskEvaluationDetailsUserVelocityByIpVelocity instantiates a new RiskEvaluationDetailsUserVelocityByIpVelocity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsUserVelocityByIpVelocityWithDefaults

`func NewRiskEvaluationDetailsUserVelocityByIpVelocityWithDefaults() *RiskEvaluationDetailsUserVelocityByIpVelocity`

NewRiskEvaluationDetailsUserVelocityByIpVelocityWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIpVelocity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistinctCount

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) GetDistinctCount() int32`

GetDistinctCount returns the DistinctCount field if non-nil, zero value otherwise.

### GetDistinctCountOk

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) GetDistinctCountOk() (*int32, bool)`

GetDistinctCountOk returns a tuple with the DistinctCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinctCount

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) SetDistinctCount(v int32)`

SetDistinctCount sets DistinctCount field to given value.

### HasDistinctCount

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) HasDistinctCount() bool`

HasDistinctCount returns a boolean if a field has been set.

### GetDuring

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) GetDuring() int32`

GetDuring returns the During field if non-nil, zero value otherwise.

### GetDuringOk

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) GetDuringOk() (*int32, bool)`

GetDuringOk returns a tuple with the During field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuring

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) SetDuring(v int32)`

SetDuring sets During field to given value.

### HasDuring

`func (o *RiskEvaluationDetailsUserVelocityByIpVelocity) HasDuring() bool`

HasDuring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


