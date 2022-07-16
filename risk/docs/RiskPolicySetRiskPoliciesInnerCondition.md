# RiskPolicySetRiskPoliciesInnerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | Pointer to **string** |  | [optional] 
**IpRange** | Pointer to **[]string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Equals** | Pointer to [**RiskPolicySetRiskPoliciesInnerConditionEquals**](RiskPolicySetRiskPoliciesInnerConditionEquals.md) |  | [optional] 
**AggregatedWeights** | Pointer to [**[]RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner**](RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner.md) |  | [optional] 
**Between** | Pointer to [**RiskPolicySetRiskPoliciesInnerConditionBetween**](RiskPolicySetRiskPoliciesInnerConditionBetween.md) |  | [optional] 

## Methods

### NewRiskPolicySetRiskPoliciesInnerCondition

`func NewRiskPolicySetRiskPoliciesInnerCondition() *RiskPolicySetRiskPoliciesInnerCondition`

NewRiskPolicySetRiskPoliciesInnerCondition instantiates a new RiskPolicySetRiskPoliciesInnerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetRiskPoliciesInnerConditionWithDefaults

`func NewRiskPolicySetRiskPoliciesInnerConditionWithDefaults() *RiskPolicySetRiskPoliciesInnerCondition`

NewRiskPolicySetRiskPoliciesInnerConditionWithDefaults instantiates a new RiskPolicySetRiskPoliciesInnerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetIpRange

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetValue

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetEquals

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetEquals() RiskPolicySetRiskPoliciesInnerConditionEquals`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetEqualsOk() (*RiskPolicySetRiskPoliciesInnerConditionEquals, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetEquals(v RiskPolicySetRiskPoliciesInnerConditionEquals)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetAggregatedWeights

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetAggregatedWeights() []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner`

GetAggregatedWeights returns the AggregatedWeights field if non-nil, zero value otherwise.

### GetAggregatedWeightsOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetAggregatedWeightsOk() (*[]RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner, bool)`

GetAggregatedWeightsOk returns a tuple with the AggregatedWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedWeights

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetAggregatedWeights(v []RiskPolicySetRiskPoliciesInnerConditionAggregatedWeightsInner)`

SetAggregatedWeights sets AggregatedWeights field to given value.

### HasAggregatedWeights

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasAggregatedWeights() bool`

HasAggregatedWeights returns a boolean if a field has been set.

### GetBetween

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetBetween() RiskPolicySetRiskPoliciesInnerConditionBetween`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *RiskPolicySetRiskPoliciesInnerCondition) GetBetweenOk() (*RiskPolicySetRiskPoliciesInnerConditionBetween, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *RiskPolicySetRiskPoliciesInnerCondition) SetBetween(v RiskPolicySetRiskPoliciesInnerConditionBetween)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *RiskPolicySetRiskPoliciesInnerCondition) HasBetween() bool`

HasBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


