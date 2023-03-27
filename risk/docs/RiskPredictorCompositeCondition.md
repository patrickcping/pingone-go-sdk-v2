# RiskPredictorCompositeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Not** | [**RiskPredictorCompositeOr**](RiskPredictorCompositeOr.md) |  | 
**Or** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 
**And** | [**[]RiskPredictorCompositeCondition**](RiskPredictorCompositeCondition.md) |  | 
**Equals** | Pointer to **map[string]interface{}** |  | [optional] 
**NotEquals** | Pointer to **map[string]interface{}** |  | [optional] 
**Greater** | Pointer to **int32** |  | [optional] 
**GreaterEquals** | Pointer to **int32** |  | [optional] 
**Lower** | Pointer to **int32** |  | [optional] 
**LowerEquals** | Pointer to **int32** |  | [optional] 
**Value** | **string** |  | 
**Type** | Pointer to [**EnumPredictorCompositeConditionType**](EnumPredictorCompositeConditionType.md) |  | [optional] 

## Methods

### NewRiskPredictorCompositeCondition

`func NewRiskPredictorCompositeCondition(not RiskPredictorCompositeOr, or []RiskPredictorCompositeCondition, and []RiskPredictorCompositeCondition, value string, ) *RiskPredictorCompositeCondition`

NewRiskPredictorCompositeCondition instantiates a new RiskPredictorCompositeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeConditionWithDefaults

`func NewRiskPredictorCompositeConditionWithDefaults() *RiskPredictorCompositeCondition`

NewRiskPredictorCompositeConditionWithDefaults instantiates a new RiskPredictorCompositeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNot

`func (o *RiskPredictorCompositeCondition) GetNot() RiskPredictorCompositeOr`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *RiskPredictorCompositeCondition) GetNotOk() (*RiskPredictorCompositeOr, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *RiskPredictorCompositeCondition) SetNot(v RiskPredictorCompositeOr)`

SetNot sets Not field to given value.


### GetOr

`func (o *RiskPredictorCompositeCondition) GetOr() []RiskPredictorCompositeCondition`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *RiskPredictorCompositeCondition) GetOrOk() (*[]RiskPredictorCompositeCondition, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *RiskPredictorCompositeCondition) SetOr(v []RiskPredictorCompositeCondition)`

SetOr sets Or field to given value.


### GetAnd

`func (o *RiskPredictorCompositeCondition) GetAnd() []RiskPredictorCompositeCondition`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *RiskPredictorCompositeCondition) GetAndOk() (*[]RiskPredictorCompositeCondition, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *RiskPredictorCompositeCondition) SetAnd(v []RiskPredictorCompositeCondition)`

SetAnd sets And field to given value.


### GetEquals

`func (o *RiskPredictorCompositeCondition) GetEquals() map[string]interface{}`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *RiskPredictorCompositeCondition) GetEqualsOk() (*map[string]interface{}, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *RiskPredictorCompositeCondition) SetEquals(v map[string]interface{})`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *RiskPredictorCompositeCondition) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetNotEquals

`func (o *RiskPredictorCompositeCondition) GetNotEquals() map[string]interface{}`

GetNotEquals returns the NotEquals field if non-nil, zero value otherwise.

### GetNotEqualsOk

`func (o *RiskPredictorCompositeCondition) GetNotEqualsOk() (*map[string]interface{}, bool)`

GetNotEqualsOk returns a tuple with the NotEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEquals

`func (o *RiskPredictorCompositeCondition) SetNotEquals(v map[string]interface{})`

SetNotEquals sets NotEquals field to given value.

### HasNotEquals

`func (o *RiskPredictorCompositeCondition) HasNotEquals() bool`

HasNotEquals returns a boolean if a field has been set.

### GetGreater

`func (o *RiskPredictorCompositeCondition) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *RiskPredictorCompositeCondition) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *RiskPredictorCompositeCondition) SetGreater(v int32)`

SetGreater sets Greater field to given value.

### HasGreater

`func (o *RiskPredictorCompositeCondition) HasGreater() bool`

HasGreater returns a boolean if a field has been set.

### GetGreaterEquals

`func (o *RiskPredictorCompositeCondition) GetGreaterEquals() int32`

GetGreaterEquals returns the GreaterEquals field if non-nil, zero value otherwise.

### GetGreaterEqualsOk

`func (o *RiskPredictorCompositeCondition) GetGreaterEqualsOk() (*int32, bool)`

GetGreaterEqualsOk returns a tuple with the GreaterEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreaterEquals

`func (o *RiskPredictorCompositeCondition) SetGreaterEquals(v int32)`

SetGreaterEquals sets GreaterEquals field to given value.

### HasGreaterEquals

`func (o *RiskPredictorCompositeCondition) HasGreaterEquals() bool`

HasGreaterEquals returns a boolean if a field has been set.

### GetLower

`func (o *RiskPredictorCompositeCondition) GetLower() int32`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *RiskPredictorCompositeCondition) GetLowerOk() (*int32, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *RiskPredictorCompositeCondition) SetLower(v int32)`

SetLower sets Lower field to given value.

### HasLower

`func (o *RiskPredictorCompositeCondition) HasLower() bool`

HasLower returns a boolean if a field has been set.

### GetLowerEquals

`func (o *RiskPredictorCompositeCondition) GetLowerEquals() int32`

GetLowerEquals returns the LowerEquals field if non-nil, zero value otherwise.

### GetLowerEqualsOk

`func (o *RiskPredictorCompositeCondition) GetLowerEqualsOk() (*int32, bool)`

GetLowerEqualsOk returns a tuple with the LowerEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerEquals

`func (o *RiskPredictorCompositeCondition) SetLowerEquals(v int32)`

SetLowerEquals sets LowerEquals field to given value.

### HasLowerEquals

`func (o *RiskPredictorCompositeCondition) HasLowerEquals() bool`

HasLowerEquals returns a boolean if a field has been set.

### GetValue

`func (o *RiskPredictorCompositeCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskPredictorCompositeCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskPredictorCompositeCondition) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *RiskPredictorCompositeCondition) GetType() EnumPredictorCompositeConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCompositeCondition) GetTypeOk() (*EnumPredictorCompositeConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCompositeCondition) SetType(v EnumPredictorCompositeConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCompositeCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


