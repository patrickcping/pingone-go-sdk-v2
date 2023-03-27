# RiskPredictorCustomItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contains** | **string** |  | 
**IpRange** | **[]string** | List of CIDRs to include | 
**Between** | [**RiskPredictorCustomItemBetweenBetween**](RiskPredictorCustomItemBetweenBetween.md) |  | 
**List** | **[]string** | An array that specifies the list of entities that represent the risk conditions. | 

## Methods

### NewRiskPredictorCustomItem

`func NewRiskPredictorCustomItem(contains string, ipRange []string, between RiskPredictorCustomItemBetweenBetween, list []string, ) *RiskPredictorCustomItem`

NewRiskPredictorCustomItem instantiates a new RiskPredictorCustomItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorCustomItemWithDefaults

`func NewRiskPredictorCustomItemWithDefaults() *RiskPredictorCustomItem`

NewRiskPredictorCustomItemWithDefaults instantiates a new RiskPredictorCustomItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContains

`func (o *RiskPredictorCustomItem) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPredictorCustomItem) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPredictorCustomItem) SetContains(v string)`

SetContains sets Contains field to given value.


### GetIpRange

`func (o *RiskPredictorCustomItem) GetIpRange() []string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *RiskPredictorCustomItem) GetIpRangeOk() (*[]string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *RiskPredictorCustomItem) SetIpRange(v []string)`

SetIpRange sets IpRange field to given value.


### GetBetween

`func (o *RiskPredictorCustomItem) GetBetween() RiskPredictorCustomItemBetweenBetween`

GetBetween returns the Between field if non-nil, zero value otherwise.

### GetBetweenOk

`func (o *RiskPredictorCustomItem) GetBetweenOk() (*RiskPredictorCustomItemBetweenBetween, bool)`

GetBetweenOk returns a tuple with the Between field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetween

`func (o *RiskPredictorCustomItem) SetBetween(v RiskPredictorCustomItemBetweenBetween)`

SetBetween sets Between field to given value.


### GetList

`func (o *RiskPredictorCustomItem) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *RiskPredictorCustomItem) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *RiskPredictorCustomItem) SetList(v []string)`

SetList sets List field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


