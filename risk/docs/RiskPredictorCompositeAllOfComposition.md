# RiskPredictorCompositeAllOfComposition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | [**NullableOneOfRiskPredictorCompositeNotRiskPredictorCompositeAndRiskPredictorCompositeOr**](oneOf&lt;RiskPredictorCompositeNot,RiskPredictorCompositeAnd,RiskPredictorCompositeOr&gt;.md) |  | 
**Level** | [**EnumRiskLevel**](EnumRiskLevel.md) |  | 

## Methods

### NewRiskPredictorCompositeAllOfComposition

`func NewRiskPredictorCompositeAllOfComposition(condition NullableOneOfRiskPredictorCompositeNotRiskPredictorCompositeAndRiskPredictorCompositeOr, level EnumRiskLevel, ) *RiskPredictorCompositeAllOfComposition`

NewRiskPredictorCompositeAllOfComposition instantiates a new RiskPredictorCompositeAllOfComposition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeAllOfCompositionWithDefaults

`func NewRiskPredictorCompositeAllOfCompositionWithDefaults() *RiskPredictorCompositeAllOfComposition`

NewRiskPredictorCompositeAllOfCompositionWithDefaults instantiates a new RiskPredictorCompositeAllOfComposition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RiskPredictorCompositeAllOfComposition) GetCondition() OneOfRiskPredictorCompositeNotRiskPredictorCompositeAndRiskPredictorCompositeOr`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RiskPredictorCompositeAllOfComposition) GetConditionOk() (*OneOfRiskPredictorCompositeNotRiskPredictorCompositeAndRiskPredictorCompositeOr, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RiskPredictorCompositeAllOfComposition) SetCondition(v OneOfRiskPredictorCompositeNotRiskPredictorCompositeAndRiskPredictorCompositeOr)`

SetCondition sets Condition field to given value.


### SetConditionNil

`func (o *RiskPredictorCompositeAllOfComposition) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *RiskPredictorCompositeAllOfComposition) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetLevel

`func (o *RiskPredictorCompositeAllOfComposition) GetLevel() EnumRiskLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskPredictorCompositeAllOfComposition) GetLevelOk() (*EnumRiskLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskPredictorCompositeAllOfComposition) SetLevel(v EnumRiskLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


