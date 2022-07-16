# RiskEvaluation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | The time the resource was created (format ISO-8061). | [optional] [readonly] 
**Details** | Pointer to [**RiskEvaluationDetails**](RiskEvaluationDetails.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Event** | [**RiskEvaluationEvent**](RiskEvaluationEvent.md) |  | 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**RiskPolicySet** | Pointer to [**RiskEvaluationRiskPolicySet**](RiskEvaluationRiskPolicySet.md) |  | [optional] 
**Result** | Pointer to [**RiskEvaluationResult**](RiskEvaluationResult.md) |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewRiskEvaluation

`func NewRiskEvaluation(event RiskEvaluationEvent, ) *RiskEvaluation`

NewRiskEvaluation instantiates a new RiskEvaluation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationWithDefaults

`func NewRiskEvaluationWithDefaults() *RiskEvaluation`

NewRiskEvaluationWithDefaults instantiates a new RiskEvaluation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *RiskEvaluation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskEvaluation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskEvaluation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskEvaluation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDetails

`func (o *RiskEvaluation) GetDetails() RiskEvaluationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RiskEvaluation) GetDetailsOk() (*RiskEvaluationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RiskEvaluation) SetDetails(v RiskEvaluationDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RiskEvaluation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskEvaluation) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskEvaluation) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskEvaluation) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskEvaluation) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEvent

`func (o *RiskEvaluation) GetEvent() RiskEvaluationEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RiskEvaluation) GetEventOk() (*RiskEvaluationEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RiskEvaluation) SetEvent(v RiskEvaluationEvent)`

SetEvent sets Event field to given value.


### GetId

`func (o *RiskEvaluation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskEvaluation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskEvaluation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskEvaluation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRiskPolicySet

`func (o *RiskEvaluation) GetRiskPolicySet() RiskEvaluationRiskPolicySet`

GetRiskPolicySet returns the RiskPolicySet field if non-nil, zero value otherwise.

### GetRiskPolicySetOk

`func (o *RiskEvaluation) GetRiskPolicySetOk() (*RiskEvaluationRiskPolicySet, bool)`

GetRiskPolicySetOk returns a tuple with the RiskPolicySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskPolicySet

`func (o *RiskEvaluation) SetRiskPolicySet(v RiskEvaluationRiskPolicySet)`

SetRiskPolicySet sets RiskPolicySet field to given value.

### HasRiskPolicySet

`func (o *RiskEvaluation) HasRiskPolicySet() bool`

HasRiskPolicySet returns a boolean if a field has been set.

### GetResult

`func (o *RiskEvaluation) GetResult() RiskEvaluationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RiskEvaluation) GetResultOk() (*RiskEvaluationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RiskEvaluation) SetResult(v RiskEvaluationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *RiskEvaluation) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskEvaluation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskEvaluation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskEvaluation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskEvaluation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


