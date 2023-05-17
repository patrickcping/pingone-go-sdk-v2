# RiskPredictorCustomItemIPRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [readonly] 
**IpRange** | **[]string** | List of CIDRs to include | 

## Methods

### NewRiskPredictorCustomItemIPRange

`func NewRiskPredictorCustomItemIPRange(contains string, ipRange []string, ) *RiskPredictorCustomItemIPRange`

NewRiskPredictorCustomItemIPRange instantiates a new RiskPredictorCustomItemIPRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCustomItemIPRangeWithDefaults

`func NewRiskPredictorCustomItemIPRangeWithDefaults() *RiskPredictorCustomItemIPRange`

NewRiskPredictorCustomItemIPRangeWithDefaults instantiates a new RiskPredictorCustomItemIPRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *RiskPredictorCustomItemIPRange) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPredictorCustomItemIPRange) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPredictorCustomItemIPRange) SetContains(v string)`

SetContains sets Contains field to given value.


### GetType

`func (o *RiskPredictorCustomItemIPRange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPredictorCustomItemIPRange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPredictorCustomItemIPRange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPredictorCustomItemIPRange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpRange

`func (o *RiskPredictorCustomItemIPRange) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPredictorCustomItemIPRange) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPredictorCustomItemIPRange) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


