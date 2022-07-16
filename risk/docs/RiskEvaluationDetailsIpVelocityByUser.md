# RiskEvaluationDetailsIpVelocityByUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | An enum indicating whether the number of distinct IPs for the user in the past hour is flagged as LOW, MEDIUM, or HIGH. | [optional] 
**Reason** | Pointer to **string** | A string indicating the reason the user was flagged. For example \&quot;More than 13 IPs were accessed by John during the last 1 hour.\&quot; | [optional] 
**Threshold** | Pointer to [**RiskEvaluationDetailsIpVelocityByUserThreshold**](RiskEvaluationDetailsIpVelocityByUserThreshold.md) |  | [optional] 
**Velocity** | Pointer to [**RiskEvaluationDetailsIpVelocityByUserVelocity**](RiskEvaluationDetailsIpVelocityByUserVelocity.md) |  | [optional] 

## Methods

### NewRiskEvaluationDetailsIpVelocityByUser

`func NewRiskEvaluationDetailsIpVelocityByUser() *RiskEvaluationDetailsIpVelocityByUser`

NewRiskEvaluationDetailsIpVelocityByUser instantiates a new RiskEvaluationDetailsIpVelocityByUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsIpVelocityByUserWithDefaults

`func NewRiskEvaluationDetailsIpVelocityByUserWithDefaults() *RiskEvaluationDetailsIpVelocityByUser`

NewRiskEvaluationDetailsIpVelocityByUserWithDefaults instantiates a new RiskEvaluationDetailsIpVelocityByUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationDetailsIpVelocityByUser) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationDetailsIpVelocityByUser) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetReason

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskEvaluationDetailsIpVelocityByUser) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskEvaluationDetailsIpVelocityByUser) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetThreshold

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetThreshold() RiskEvaluationDetailsIpVelocityByUserThreshold`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetThresholdOk() (*RiskEvaluationDetailsIpVelocityByUserThreshold, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RiskEvaluationDetailsIpVelocityByUser) SetThreshold(v RiskEvaluationDetailsIpVelocityByUserThreshold)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *RiskEvaluationDetailsIpVelocityByUser) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetVelocity

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetVelocity() RiskEvaluationDetailsIpVelocityByUserVelocity`

GetVelocity returns the Velocity field if non-nil, zero value otherwise.

### GetVelocityOk

`func (o *RiskEvaluationDetailsIpVelocityByUser) GetVelocityOk() (*RiskEvaluationDetailsIpVelocityByUserVelocity, bool)`

GetVelocityOk returns a tuple with the Velocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVelocity

`func (o *RiskEvaluationDetailsIpVelocityByUser) SetVelocity(v RiskEvaluationDetailsIpVelocityByUserVelocity)`

SetVelocity sets Velocity field to given value.

### HasVelocity

`func (o *RiskEvaluationDetailsIpVelocityByUser) HasVelocity() bool`

HasVelocity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


