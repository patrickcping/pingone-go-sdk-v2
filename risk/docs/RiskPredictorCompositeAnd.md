# RiskPredictorCompositeAnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 
**Type** | Pointer to [**EnumPredictorCompositeConditionType**](EnumPredictorCompositeConditionType.md) |  | [optional] 

## Methods

### NewRiskPredictorCompositeAnd

`func NewRiskPredictorCompositeAnd(and []RiskPredictorCompositeCondition, ) *RiskPredictorCompositeAnd`

NewRiskPredictorCompositeAnd instantiates a new RiskPredictorCompositeAnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeAndWithDefaults

`func NewRiskPredictorCompositeAndWithDefaults() *RiskPredictorCompositeAnd`

NewRiskPredictorCompositeAndWithDefaults instantiates a new RiskPredictorCompositeAnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *RiskPredictorCompositeAnd) GetAnd() []RiskPredictorCompositeCondition`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *RiskPredictorCompositeAnd) GetAndOk() (*[]RiskPredictorCompositeCondition, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *RiskPredictorCompositeAnd) SetAnd(v []RiskPredictorCompositeCondition)`

SetAnd sets And field to given value.


### GetType

`func (o *RiskPredictorCompositeAnd) GetType() EnumPredictorCompositeConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCompositeAnd) GetTypeOk() (*EnumPredictorCompositeConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCompositeAnd) SetType(v EnumPredictorCompositeConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCompositeAnd) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


