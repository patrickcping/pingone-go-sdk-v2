# RiskEvaluationDetailsIpAddressReputation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**EnumRiskLevel**](EnumRiskLevel.md) | A string that specifies the risk level of the evaluated IP address. Options are LOW, MEDIUM, and HIGH. If the score is less than 55, the level is LOW; if the score is greater than 77, the level is HIGH; if the score is between 55 and 77, the level is MEDIUM. Note that these guidelines could change based on data analytics and product consideration. If the ipAddressReputation.score is unknown, NULL is returned. | [optional] 
**Score** | Pointer to **int32** | An integer that represents the calculated score of the IP address involved in the transaction. Scores range between 0 and 100. A score of 0 indicates a non-risky IP address; a score of 100 indicates a high-risk IP address. If the IP address reputation score is not available for the specific IP address, NULL is returned. | [optional] 

## Methods

### NewRiskEvaluationDetailsIpAddressReputation

`func NewRiskEvaluationDetailsIpAddressReputation() *RiskEvaluationDetailsIpAddressReputation`

NewRiskEvaluationDetailsIpAddressReputation instantiates a new RiskEvaluationDetailsIpAddressReputation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsIpAddressReputationWithDefaults

`func NewRiskEvaluationDetailsIpAddressReputationWithDefaults() *RiskEvaluationDetailsIpAddressReputation`

NewRiskEvaluationDetailsIpAddressReputationWithDefaults instantiates a new RiskEvaluationDetailsIpAddressReputation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskEvaluationDetailsIpAddressReputation) GetLevel() EnumRiskLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationDetailsIpAddressReputation) GetLevelOk() (*EnumRiskLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationDetailsIpAddressReputation) SetLevel(v EnumRiskLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationDetailsIpAddressReputation) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetScore

`func (o *RiskEvaluationDetailsIpAddressReputation) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RiskEvaluationDetailsIpAddressReputation) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RiskEvaluationDetailsIpAddressReputation) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RiskEvaluationDetailsIpAddressReputation) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


