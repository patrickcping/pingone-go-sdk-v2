# RiskPredictorNewDeviceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationAt** | Pointer to **time.Time** | You can use the &#x60;activationAt&#x60; parameter to specify a date on which the learning process for the predictor should be restarted. This can be used in conjunction with the fallback setting (&#x60;default.result.level&#x60;) to force strong authentication when moving the predictor to production. The date should be in an RFC3339 format. Note that activation date uses UTC time. | [optional] 

## Methods

### NewRiskPredictorNewDeviceAllOf

`func NewRiskPredictorNewDeviceAllOf() *RiskPredictorNewDeviceAllOf`

NewRiskPredictorNewDeviceAllOf instantiates a new RiskPredictorNewDeviceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPredictorNewDeviceAllOfWithDefaults

`func NewRiskPredictorNewDeviceAllOfWithDefaults() *RiskPredictorNewDeviceAllOf`

NewRiskPredictorNewDeviceAllOfWithDefaults instantiates a new RiskPredictorNewDeviceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationAt

`func (o *RiskPredictorNewDeviceAllOf) GetActivationAt() time.Time`

GetActivationAt returns the ActivationAt field if non-nil, zero value otherwise.

### GetActivationAtOk

`func (o *RiskPredictorNewDeviceAllOf) GetActivationAtOk() (*time.Time, bool)`

GetActivationAtOk returns a tuple with the ActivationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationAt

`func (o *RiskPredictorNewDeviceAllOf) SetActivationAt(v time.Time)`

SetActivationAt sets ActivationAt field to given value.

### HasActivationAt

`func (o *RiskPredictorNewDeviceAllOf) HasActivationAt() bool`

HasActivationAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

