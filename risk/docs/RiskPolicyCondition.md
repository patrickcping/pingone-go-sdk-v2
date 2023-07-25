# RiskPolicyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EnumRiskPolicyConditionType**](EnumRiskPolicyConditionType.md) |  | [optional] 
**Contains** | Pointer to **string** |  | [optional] 
**IpRange** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Equals** | Pointer to [**RiskPolicyConditionEquals**](RiskPolicyConditionEquals.md) |  | [optional] 
**AggregatedWeights** | Pointer to [**[]RiskPolicyConditionAggregatedWeightsInner**](RiskPolicyConditionAggregatedWeightsInner.md) | Required for weights-based policies. The elements in the array are &#x60;value&#x60;-&#x60;weight&#x60; pairs, representing a weighting for the weighted average calculation that should be assigned to a specific predictor when it is determined that there is a high risk for the predictor. | [optional] 
**AggregatedScores** | Pointer to [**[]RiskPolicyConditionAggregatedScoresInner**](RiskPolicyConditionAggregatedScoresInner.md) | Required for score-based policies. The elements in the array are &#x60;value&#x60;-&#x60;score&#x60; pairs, representing the score that should be assigned to a specific predictor when it is determined that there is a high risk for the predictor. | [optional] 
**Between** | Pointer to [**RiskPolicyConditionBetween**](RiskPolicyConditionBetween.md) |  | [optional] 

## Methods

### NewRiskPolicyCondition

`func NewRiskPolicyCondition() *RiskPolicyCondition`

NewRiskPolicyCondition instantiates a new RiskPolicyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicyConditionWithDefaults

`func NewRiskPolicyConditionWithDefaults() *RiskPolicyCondition`

NewRiskPolicyConditionWithDefaults instantiates a new RiskPolicyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RiskPolicyCondition) GetType() EnumRiskPolicyConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPolicyCondition) GetTypeOk() (*EnumRiskPolicyConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPolicyCondition) SetType(v EnumRiskPolicyConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPolicyCondition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContains

`func (o *RiskPolicyCondition) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPolicyCondition) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPolicyCondition) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *RiskPolicyCondition) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetIpRange

`func (o *RiskPolicyCondition) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPolicyCondition) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPolicyCondition) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *RiskPolicyCondition) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetValue

`func (o *RiskPolicyCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPolicyCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPolicyCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskPolicyCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEquals

`func (o *RiskPolicyCondition) GetEquals() RiskPolicyConditionEquals`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *RiskPolicyCondition) GetEqualsOk() (*RiskPolicyConditionEquals, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *RiskPolicyCondition) SetEquals(v RiskPolicyConditionEquals)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *RiskPolicyCondition) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetAggregatedWeights

`func (o *RiskPolicyCondition) GetAggregatedWeights() []RiskPolicyConditionAggregatedWeightsInner`

GetAggregatedWeights returns the AggregatedWeights field if non-nil, zero value otherwise.

### GetAggregatedWeightsOk

`func (o *RiskPolicyCondition) GetAggregatedWeightsOk() (*[]RiskPolicyConditionAggregatedWeightsInner, bool)`

GetAggregatedWeightsOk returns a tuple with the AggregatedWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedWeights

`func (o *RiskPolicyCondition) SetAggregatedWeights(v []RiskPolicyConditionAggregatedWeightsInner)`

SetAggregatedWeights sets AggregatedWeights field to given value.

### HasAggregatedWeights

`func (o *RiskPolicyCondition) HasAggregatedWeights() bool`

HasAggregatedWeights returns a boolean if a field has been set.

### GetAggregatedScores

`func (o *RiskPolicyCondition) GetAggregatedScores() []RiskPolicyConditionAggregatedScoresInner`

GetAggregatedScores returns the AggregatedScores field if non-nil, zero value otherwise.

### GetAggregatedScoresOk

`func (o *RiskPolicyCondition) GetAggregatedScoresOk() (*[]RiskPolicyConditionAggregatedScoresInner, bool)`

GetAggregatedScoresOk returns a tuple with the AggregatedScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedScores

`func (o *RiskPolicyCondition) SetAggregatedScores(v []RiskPolicyConditionAggregatedScoresInner)`

SetAggregatedScores sets AggregatedScores field to given value.

### HasAggregatedScores

`func (o *RiskPolicyCondition) HasAggregatedScores() bool`

HasAggregatedScores returns a boolean if a field has been set.

### GetBetween

`func (o *RiskPolicyCondition) GetBetween() RiskPolicyConditionBetween`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *RiskPolicyCondition) GetBetweenOk() (*RiskPolicyConditionBetween, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *RiskPolicyCondition) SetBetween(v RiskPolicyConditionBetween)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *RiskPolicyCondition) HasBetween() bool`

HasBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


