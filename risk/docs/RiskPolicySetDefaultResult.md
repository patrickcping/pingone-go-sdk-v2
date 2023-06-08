# RiskPolicySetDefaultResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**EnumRiskPolicyResultLevel**](EnumRiskPolicyResultLevel.md) |  | 
**Type** | Pointer to [**EnumResultType**](EnumResultType.md) |  | [optional] 

## Methods

### NewRiskPolicySetDefaultResult

`func NewRiskPolicySetDefaultResult(level EnumRiskPolicyResultLevel, ) *RiskPolicySetDefaultResult`

NewRiskPolicySetDefaultResult instantiates a new RiskPolicySetDefaultResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetDefaultResultWithDefaults

`func NewRiskPolicySetDefaultResultWithDefaults() *RiskPolicySetDefaultResult`

NewRiskPolicySetDefaultResultWithDefaults instantiates a new RiskPolicySetDefaultResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskPolicySetDefaultResult) GetLevel() EnumRiskPolicyResultLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskPolicySetDefaultResult) GetLevelOk() (*EnumRiskPolicyResultLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskPolicySetDefaultResult) SetLevel(v EnumRiskPolicyResultLevel)`

SetLevel sets Level field to given value.


### GetType

`func (o *RiskPolicySetDefaultResult) GetType() EnumResultType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPolicySetDefaultResult) GetTypeOk() (*EnumResultType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPolicySetDefaultResult) SetType(v EnumResultType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPolicySetDefaultResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


