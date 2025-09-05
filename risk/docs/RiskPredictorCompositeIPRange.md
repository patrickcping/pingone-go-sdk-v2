# RiskPredictorCompositeIPRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | Pointer to **string** |  | [optional] 
**NotContains** | Pointer to **string** |  | [optional] 
**Equals** | Pointer to **string** |  | [optional] 
**NotEquals** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**EnumPredictorCompositeConditionType**](EnumPredictorCompositeConditionType.md) |  | [optional] 
**IpRange** | **[]string** | List of CIDRs to include | 
**Ip** | Pointer to **string** |  | [optional] 

## Methods

### NewRiskPredictorCompositeIPRange

`func NewRiskPredictorCompositeIPRange(ipRange []string, ) *RiskPredictorCompositeIPRange`

NewRiskPredictorCompositeIPRange instantiates a new RiskPredictorCompositeIPRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCompositeIPRangeWithDefaults

`func NewRiskPredictorCompositeIPRangeWithDefaults() *RiskPredictorCompositeIPRange`

NewRiskPredictorCompositeIPRangeWithDefaults instantiates a new RiskPredictorCompositeIPRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *RiskPredictorCompositeIPRange) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPredictorCompositeIPRange) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPredictorCompositeIPRange) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *RiskPredictorCompositeIPRange) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetNotContains

`func (o *RiskPredictorCompositeIPRange) GetNotContains() string`

GetNotContains returns the NotContains field if non-nil, zero value otherwise.

### GetNotContainsOk

`func (o *RiskPredictorCompositeIPRange) GetNotContainsOk() (*string, bool)`

GetNotContainsOk returns a tuple with the NotContains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotContains

`func (o *RiskPredictorCompositeIPRange) SetNotContains(v string)`

SetNotContains sets NotContains field to given value.

### HasNotContains

`func (o *RiskPredictorCompositeIPRange) HasNotContains() bool`

HasNotContains returns a boolean if a field has been set.

### GetEquals

`func (o *RiskPredictorCompositeIPRange) GetEquals() string`

GetEquals returns the Equals field if non-nil, zero value otherwise.

### GetEqualsOk

`func (o *RiskPredictorCompositeIPRange) GetEqualsOk() (*string, bool)`

GetEqualsOk returns a tuple with the Equals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquals

`func (o *RiskPredictorCompositeIPRange) SetEquals(v string)`

SetEquals sets Equals field to given value.

### HasEquals

`func (o *RiskPredictorCompositeIPRange) HasEquals() bool`

HasEquals returns a boolean if a field has been set.

### GetNotEquals

`func (o *RiskPredictorCompositeIPRange) GetNotEquals() string`

GetNotEquals returns the NotEquals field if non-nil, zero value otherwise.

### GetNotEqualsOk

`func (o *RiskPredictorCompositeIPRange) GetNotEqualsOk() (*string, bool)`

GetNotEqualsOk returns a tuple with the NotEquals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotEquals

`func (o *RiskPredictorCompositeIPRange) SetNotEquals(v string)`

SetNotEquals sets NotEquals field to given value.

### HasNotEquals

`func (o *RiskPredictorCompositeIPRange) HasNotEquals() bool`

HasNotEquals returns a boolean if a field has been set.

### GetType

`func (o *RiskPredictorCompositeIPRange) GetType() EnumPredictorCompositeConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCompositeIPRange) GetTypeOk() (*EnumPredictorCompositeConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCompositeIPRange) SetType(v EnumPredictorCompositeConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCompositeIPRange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpRange

`func (o *RiskPredictorCompositeIPRange) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPredictorCompositeIPRange) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPredictorCompositeIPRange) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetIp

`func (o *RiskPredictorCompositeIPRange) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RiskPredictorCompositeIPRange) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RiskPredictorCompositeIPRange) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *RiskPredictorCompositeIPRange) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


