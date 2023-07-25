# RiskPredictorDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detect** | [**EnumPredictorNewDeviceDetectType**](EnumPredictorNewDeviceDetectType.md) |  | 
**ActivationAt** | Pointer to **time.Time** | You can use the &#x60;activationAt&#x60; parameter to specify a date on which the learning process for the predictor should be restarted. This can be used in conjunction with the fallback setting (&#x60;default.result.level&#x60;) to force strong authentication when moving the predictor to production. The date should be in an RFC3339 format. Note that activation date uses UTC time. | [optional] 

## Methods

### NewRiskPredictorDeviceAllOf

`func NewRiskPredictorDeviceAllOf(detect EnumPredictorNewDeviceDetectType, ) *RiskPredictorDeviceAllOf`

NewRiskPredictorDeviceAllOf instantiates a new RiskPredictorDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorDeviceAllOfWithDefaults

`func NewRiskPredictorDeviceAllOfWithDefaults() *RiskPredictorDeviceAllOf`

NewRiskPredictorDeviceAllOfWithDefaults instantiates a new RiskPredictorDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetect

`func (o *RiskPredictorDeviceAllOf) GetDetect() EnumPredictorNewDeviceDetectType`

GetDetect returns the Detect field if non-nil, zero value otherwise.

### GetDetectOk

`func (o *RiskPredictorDeviceAllOf) GetDetectOk() (*EnumPredictorNewDeviceDetectType, bool)`

GetDetectOk returns a tuple with the Detect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetect

`func (o *RiskPredictorDeviceAllOf) SetDetect(v EnumPredictorNewDeviceDetectType)`

SetDetect sets Detect field to given value.


### GetActivationAt

`func (o *RiskPredictorDeviceAllOf) GetActivationAt() time.Time`

GetActivationAt returns the ActivationAt field if non-nil, zero value otherwise.

### GetActivationAtOk

`func (o *RiskPredictorDeviceAllOf) GetActivationAtOk() (*time.Time, bool)`

GetActivationAtOk returns a tuple with the ActivationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationAt

`func (o *RiskPredictorDeviceAllOf) SetActivationAt(v time.Time)`

SetActivationAt sets ActivationAt field to given value.

### HasActivationAt

`func (o *RiskPredictorDeviceAllOf) HasActivationAt() bool`

HasActivationAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


