# RiskEvaluationDetailsUserVelocityByIpThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**High** | Pointer to **int32** | An integer indicating the value calculated for the high threshold. If the IP was accessed by more than the high number of users during the past hour, the IP is flagged as a HIGH userVelocityByIp.level. | [optional] 
**Medium** | Pointer to **int32** | An integer indicating the value calculated for the medium threshold. If the IP was accessed by more than the medium number of users during the past hour, the IP is flagged as a MEDIUM userVelocityByIp.level | [optional] 
**Source** | Pointer to [**EnumThresholdSource**](EnumThresholdSource.md) |  | [optional] 
**CalculatedAt** | Pointer to **string** | A date-time indicating the timestamp for the calculated threshold. | [optional] 
**ExpiresAt** | Pointer to **string** | A date-time indicating when the threshold will be recalculated. The recalculation will happen before this time. | [optional] 

## Methods

### NewRiskEvaluationDetailsUserVelocityByIpThreshold

`func NewRiskEvaluationDetailsUserVelocityByIpThreshold() *RiskEvaluationDetailsUserVelocityByIpThreshold`

NewRiskEvaluationDetailsUserVelocityByIpThreshold instantiates a new RiskEvaluationDetailsUserVelocityByIpThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsUserVelocityByIpThresholdWithDefaults

`func NewRiskEvaluationDetailsUserVelocityByIpThresholdWithDefaults() *RiskEvaluationDetailsUserVelocityByIpThreshold`

NewRiskEvaluationDetailsUserVelocityByIpThresholdWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIpThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHigh

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetMedium

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetSource

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetSource() EnumThresholdSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetSourceOk() (*EnumThresholdSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetSource(v EnumThresholdSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCalculatedAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetCalculatedAt() string`

GetCalculatedAt returns the CalculatedAt field if non-nil, zero value otherwise.

### GetCalculatedAtOk

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetCalculatedAtOk() (*string, bool)`

GetCalculatedAtOk returns a tuple with the CalculatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculatedAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetCalculatedAt(v string)`

SetCalculatedAt sets CalculatedAt field to given value.

### HasCalculatedAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasCalculatedAt() bool`

HasCalculatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *RiskEvaluationDetailsUserVelocityByIpThreshold) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


