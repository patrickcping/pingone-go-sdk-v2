# RiskEvaluationDetailsUserBasedRiskBehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to **string** | A string that specifies the risk score calculated for the user behavior for all transactions associated with the accessing device, and for the current authentication attempt. Options are LOW, MEDIUM, and HIGH. This information is available only if customers have agreed to data consent and the intelligence.allowDataConsent property in the PingOne license is set to true. | [optional] 
**Reason** | Pointer to **string** | A string that describes the reason (or reasons) provided for the user behavior risk score classification (for example, the operating system or browser type used by the device, and country in which the accessing device is located). Each reason is classified as Unusual, to indicate how much it deviates from normal user behavior, and its effect in calculating the overall user behavior risk score. | [optional] 

## Methods

### NewRiskEvaluationDetailsUserBasedRiskBehavior

`func NewRiskEvaluationDetailsUserBasedRiskBehavior() *RiskEvaluationDetailsUserBasedRiskBehavior`

NewRiskEvaluationDetailsUserBasedRiskBehavior instantiates a new RiskEvaluationDetailsUserBasedRiskBehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsUserBasedRiskBehaviorWithDefaults

`func NewRiskEvaluationDetailsUserBasedRiskBehaviorWithDefaults() *RiskEvaluationDetailsUserBasedRiskBehavior`

NewRiskEvaluationDetailsUserBasedRiskBehaviorWithDefaults instantiates a new RiskEvaluationDetailsUserBasedRiskBehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetReason

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskEvaluationDetailsUserBasedRiskBehavior) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


