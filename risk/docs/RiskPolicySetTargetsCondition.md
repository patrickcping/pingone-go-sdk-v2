# RiskPolicySetTargetsCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]RiskPolicySetTargetsConditionAndInner**](RiskPolicySetTargetsConditionAndInner.md) | Array of the conditions for applying the policy. It is mandatory to include in the array the relevant transaction types. The relevant user groups and applications are optional. Each element in the array contains the target transaction types/user groups/applications and the event attribute that is checked. | [optional] 

## Methods

### NewRiskPolicySetTargetsCondition

`func NewRiskPolicySetTargetsCondition() *RiskPolicySetTargetsCondition`

NewRiskPolicySetTargetsCondition instantiates a new RiskPolicySetTargetsCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetTargetsConditionWithDefaults

`func NewRiskPolicySetTargetsConditionWithDefaults() *RiskPolicySetTargetsCondition`

NewRiskPolicySetTargetsConditionWithDefaults instantiates a new RiskPolicySetTargetsCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *RiskPolicySetTargetsCondition) GetAnd() []RiskPolicySetTargetsConditionAndInner`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *RiskPolicySetTargetsCondition) GetAndOk() (*[]RiskPolicySetTargetsConditionAndInner, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *RiskPolicySetTargetsCondition) SetAnd(v []RiskPolicySetTargetsConditionAndInner)`

SetAnd sets And field to given value.

### HasAnd

`func (o *RiskPolicySetTargetsCondition) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


