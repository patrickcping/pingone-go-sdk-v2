# RiskPredictorCompositeConditionBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Not** | [**RiskPredictorCompositeOr**](RiskPredictorCompositeOr.md) |  | 
**Type** | Pointer to [**EnumPredictorCompositeConditionType**](EnumPredictorCompositeConditionType.md) |  | [optional] 
**And** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 
**Or** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 

## Methods

### NewRiskPredictorCompositeConditionBase

`func NewRiskPredictorCompositeConditionBase(not RiskPredictorCompositeOr, and []RiskPredictorCompositeCondition, or []RiskPredictorCompositeCondition, ) *RiskPredictorCompositeConditionBase`

NewRiskPredictorCompositeConditionBase instantiates a new RiskPredictorCompositeConditionBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeConditionBaseWithDefaults

`func NewRiskPredictorCompositeConditionBaseWithDefaults() *RiskPredictorCompositeConditionBase`

NewRiskPredictorCompositeConditionBaseWithDefaults instantiates a new RiskPredictorCompositeConditionBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNot

`func (o *RiskPredictorCompositeConditionBase) GetNot() RiskPredictorCompositeOr`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *RiskPredictorCompositeConditionBase) GetNotOk() (*RiskPredictorCompositeOr, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *RiskPredictorCompositeConditionBase) SetNot(v RiskPredictorCompositeOr)`

SetNot sets Not field to given value.


### GetType

`func (o *RiskPredictorCompositeConditionBase) GetType() EnumPredictorCompositeConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCompositeConditionBase) GetTypeOk() (*EnumPredictorCompositeConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCompositeConditionBase) SetType(v EnumPredictorCompositeConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCompositeConditionBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAnd

`func (o *RiskPredictorCompositeConditionBase) GetAnd() []RiskPredictorCompositeCondition`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *RiskPredictorCompositeConditionBase) GetAndOk() (*[]RiskPredictorCompositeCondition, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *RiskPredictorCompositeConditionBase) SetAnd(v []RiskPredictorCompositeCondition)`

SetAnd sets And field to given value.


### GetOr

`func (o *RiskPredictorCompositeConditionBase) GetOr() []RiskPredictorCompositeCondition`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *RiskPredictorCompositeConditionBase) GetOrOk() (*[]RiskPredictorCompositeCondition, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *RiskPredictorCompositeConditionBase) SetOr(v []RiskPredictorCompositeCondition)`

SetOr sets Or field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


