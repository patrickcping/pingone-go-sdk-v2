# RiskPredictorUserLocationAnomalyAllOfRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | **int32** | A value to apply to the distance radius.  Minimum of 10 miles (16 km) and maximum of 100 miles (160 km) | 
**Unit** | [**EnumDistanceUnit**](EnumDistanceUnit.md) |  | 

## Methods

### NewRiskPredictorUserLocationAnomalyAllOfRadius

`func NewRiskPredictorUserLocationAnomalyAllOfRadius(distance int32, unit EnumDistanceUnit, ) *RiskPredictorUserLocationAnomalyAllOfRadius`

NewRiskPredictorUserLocationAnomalyAllOfRadius instantiates a new RiskPredictorUserLocationAnomalyAllOfRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorUserLocationAnomalyAllOfRadiusWithDefaults

`func NewRiskPredictorUserLocationAnomalyAllOfRadiusWithDefaults() *RiskPredictorUserLocationAnomalyAllOfRadius`

NewRiskPredictorUserLocationAnomalyAllOfRadiusWithDefaults instantiates a new RiskPredictorUserLocationAnomalyAllOfRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistance

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) SetDistance(v int32)`

SetDistance sets Distance field to given value.


### GetUnit

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetUnit() EnumDistanceUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) GetUnitOk() (*EnumDistanceUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *RiskPredictorUserLocationAnomalyAllOfRadius) SetUnit(v EnumDistanceUnit)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


