# RiskEvaluationDetailsUserRiskBehavior

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | Pointer to [**EnumRiskLevel**](EnumRiskLevel.md) |  | [optional] 
**Reason** | Pointer to **string** | A string that describes the reason (or reasons) provided for the user behavior risk score classification within the organization (for example, the operating system or browser type used by the device, and country in which the accessing device is located). Each reason is classified as Unusual or Very Unusual, to indicate how much it deviates from normal user behavior, and its effect in calculating the overall user behavior risk score. | [optional] 

## Methods

### NewRiskEvaluationDetailsUserRiskBehavior

`func NewRiskEvaluationDetailsUserRiskBehavior() *RiskEvaluationDetailsUserRiskBehavior`

NewRiskEvaluationDetailsUserRiskBehavior instantiates a new RiskEvaluationDetailsUserRiskBehavior object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationDetailsUserRiskBehaviorWithDefaults

`func NewRiskEvaluationDetailsUserRiskBehaviorWithDefaults() *RiskEvaluationDetailsUserRiskBehavior`

NewRiskEvaluationDetailsUserRiskBehaviorWithDefaults instantiates a new RiskEvaluationDetailsUserRiskBehavior object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *RiskEvaluationDetailsUserRiskBehavior) GetLevel() EnumRiskLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *RiskEvaluationDetailsUserRiskBehavior) GetLevelOk() (*EnumRiskLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *RiskEvaluationDetailsUserRiskBehavior) SetLevel(v EnumRiskLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *RiskEvaluationDetailsUserRiskBehavior) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetReason

`func (o *RiskEvaluationDetailsUserRiskBehavior) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskEvaluationDetailsUserRiskBehavior) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskEvaluationDetailsUserRiskBehavior) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskEvaluationDetailsUserRiskBehavior) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


