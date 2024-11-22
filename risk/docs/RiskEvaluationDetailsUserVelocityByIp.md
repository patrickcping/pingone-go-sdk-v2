# RiskEvaluationDetailsUserVelocityByIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**EnumRiskLevel**](EnumRiskLevel.md) | An enum indicating whether the calculated number of users per IP is LOW, MEDIUM, or HIGH. | [optional] 
**Reason** | Pointer to **string** | A string indicating the reason the user was flagged. For example \&quot;More than 250 users accessed IP address 1.1.1.1 during the last 1 hour.\&quot; | [optional] 
**Threshold** | Pointer to [**RiskEvaluationDetailsUserVelocityByIpThreshold**](RiskEvaluationDetailsUserVelocityByIpThreshold.md) |  | [optional] 
**Velocity** | Pointer to [**RiskEvaluationDetailsUserVelocityByIpVelocity**](RiskEvaluationDetailsUserVelocityByIpVelocity.md) |  | [optional] 

## Methods

### NewRiskEvaluationDetailsUserVelocityByIp

`func NewRiskEvaluationDetailsUserVelocityByIp() *RiskEvaluationDetailsUserVelocityByIp`

NewRiskEvaluationDetailsUserVelocityByIp instantiates a new RiskEvaluationDetailsUserVelocityByIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsUserVelocityByIpWithDefaults

`func NewRiskEvaluationDetailsUserVelocityByIpWithDefaults() *RiskEvaluationDetailsUserVelocityByIp`

NewRiskEvaluationDetailsUserVelocityByIpWithDefaults instantiates a new RiskEvaluationDetailsUserVelocityByIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevel() EnumRiskLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetLevelOk() (*EnumRiskLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationDetailsUserVelocityByIp) SetLevel(v EnumRiskLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationDetailsUserVelocityByIp) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetReason

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskEvaluationDetailsUserVelocityByIp) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskEvaluationDetailsUserVelocityByIp) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetThreshold

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetThreshold() RiskEvaluationDetailsUserVelocityByIpThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetThresholdOk() (*RiskEvaluationDetailsUserVelocityByIpThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RiskEvaluationDetailsUserVelocityByIp) SetThreshold(v RiskEvaluationDetailsUserVelocityByIpThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RiskEvaluationDetailsUserVelocityByIp) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetVelocity

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocity() RiskEvaluationDetailsUserVelocityByIpVelocity`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *RiskEvaluationDetailsUserVelocityByIp) GetVelocityOk() (*RiskEvaluationDetailsUserVelocityByIpVelocity, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *RiskEvaluationDetailsUserVelocityByIp) SetVelocity(v RiskEvaluationDetailsUserVelocityByIpVelocity)`

SetVelocity sets Velocity field to given value.

### HasVelocity

`func (o *RiskEvaluationDetailsUserVelocityByIp) HasVelocity() bool`

HasVelocity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


