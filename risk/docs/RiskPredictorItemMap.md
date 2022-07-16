# RiskPredictorItemMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | **string** |  | 
**IpRange** | Pointer to **[]string** | List of CIDRs to include | [optional] 
**Between** | Pointer to [**RiskPredictorItemMapBetween**](RiskPredictorItemMapBetween.md) |  | [optional] 
**List** | Pointer to **[]string** | An array that specifies the list of entities that represent the risk conditions. | [optional] 

## Methods

### NewRiskPredictorItemMap

`func NewRiskPredictorItemMap(contains string, ) *RiskPredictorItemMap`

NewRiskPredictorItemMap instantiates a new RiskPredictorItemMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorItemMapWithDefaults

`func NewRiskPredictorItemMapWithDefaults() *RiskPredictorItemMap`

NewRiskPredictorItemMapWithDefaults instantiates a new RiskPredictorItemMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *RiskPredictorItemMap) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPredictorItemMap) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPredictorItemMap) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *RiskPredictorItemMap) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPredictorItemMap) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPredictorItemMap) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *RiskPredictorItemMap) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetBetween

`func (o *RiskPredictorItemMap) GetBetween() RiskPredictorItemMapBetween`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *RiskPredictorItemMap) GetBetweenOk() (*RiskPredictorItemMapBetween, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *RiskPredictorItemMap) SetBetween(v RiskPredictorItemMapBetween)`

SetBetween sets Between field to given value.

### HasBetween

`func (o *RiskPredictorItemMap) HasBetween() bool`

HasBetween returns a boolean if a field has been set.

### GetList

`func (o *RiskPredictorItemMap) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *RiskPredictorItemMap) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *RiskPredictorItemMap) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *RiskPredictorItemMap) HasList() bool`

HasList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


