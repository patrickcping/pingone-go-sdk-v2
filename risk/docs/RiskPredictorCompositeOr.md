# RiskPredictorCompositeOr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Or** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 
**Type** | Pointer to [**EnumPredictorCompositeConditionType**](EnumPredictorCompositeConditionType.md) |  | [optional] 

## Methods

### NewRiskPredictorCompositeOr

`func NewRiskPredictorCompositeOr(or []RiskPredictorCompositeCondition, ) *RiskPredictorCompositeOr`

NewRiskPredictorCompositeOr instantiates a new RiskPredictorCompositeOr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeOrWithDefaults

`func NewRiskPredictorCompositeOrWithDefaults() *RiskPredictorCompositeOr`

NewRiskPredictorCompositeOrWithDefaults instantiates a new RiskPredictorCompositeOr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOr

`func (o *RiskPredictorCompositeOr) GetOr() []RiskPredictorCompositeCondition`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *RiskPredictorCompositeOr) GetOrOk() (*[]RiskPredictorCompositeCondition, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *RiskPredictorCompositeOr) SetOr(v []RiskPredictorCompositeCondition)`

SetOr sets Or field to given value.


### GetType

`func (o *RiskPredictorCompositeOr) GetType() EnumPredictorCompositeConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCompositeOr) GetTypeOk() (*EnumPredictorCompositeConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCompositeOr) SetType(v EnumPredictorCompositeConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCompositeOr) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


