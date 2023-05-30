# RiskPolicySetTriggersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumRiskPolicySetTriggerType**](EnumRiskPolicySetTriggerType.md) |  | 
**PolicySet** | [**RiskPolicySetEvaluatedPredictorsInner**](RiskPolicySetEvaluatedPredictorsInner.md) |  | 
**ExpiresAt** | **time.Time** | The time the trigger expires (format ISO-8061). | 

## Methods

### NewRiskPolicySetTriggersInner

`func NewRiskPolicySetTriggersInner(type_ EnumRiskPolicySetTriggerType, policySet RiskPolicySetEvaluatedPredictorsInner, expiresAt time.Time, ) *RiskPolicySetTriggersInner`

NewRiskPolicySetTriggersInner instantiates a new RiskPolicySetTriggersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetTriggersInnerWithDefaults

`func NewRiskPolicySetTriggersInnerWithDefaults() *RiskPolicySetTriggersInner`

NewRiskPolicySetTriggersInnerWithDefaults instantiates a new RiskPolicySetTriggersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RiskPolicySetTriggersInner) GetType() EnumRiskPolicySetTriggerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPolicySetTriggersInner) GetTypeOk() (*EnumRiskPolicySetTriggerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPolicySetTriggersInner) SetType(v EnumRiskPolicySetTriggerType)`

SetType sets Type field to given value.


### GetPolicySet

`func (o *RiskPolicySetTriggersInner) GetPolicySet() RiskPolicySetEvaluatedPredictorsInner`

GetPolicySet returns the PolicySet field if non-nil, zero value otherwise.

### GetPolicySetOk

`func (o *RiskPolicySetTriggersInner) GetPolicySetOk() (*RiskPolicySetEvaluatedPredictorsInner, bool)`

GetPolicySetOk returns a tuple with the PolicySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicySet

`func (o *RiskPolicySetTriggersInner) SetPolicySet(v RiskPolicySetEvaluatedPredictorsInner)`

SetPolicySet sets PolicySet field to given value.


### GetExpiresAt

`func (o *RiskPolicySetTriggersInner) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RiskPolicySetTriggersInner) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RiskPolicySetTriggersInner) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


