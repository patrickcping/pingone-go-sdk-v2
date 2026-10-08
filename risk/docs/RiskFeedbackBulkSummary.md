# RiskFeedbackBulkSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requested** | Pointer to **int32** | The number of items the request contained. | [optional] 
**Errors** | Pointer to **int32** | The number of items that failed validation. | [optional] 
**Successful** | Pointer to **int32** | The number of items that were persisted. | [optional] 
**Skipped** | Pointer to **int32** | The number of items that were not processed. | [optional] 

## Methods

### NewRiskFeedbackBulkSummary

`func NewRiskFeedbackBulkSummary() *RiskFeedbackBulkSummary`

NewRiskFeedbackBulkSummary instantiates a new RiskFeedbackBulkSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackBulkSummaryWithDefaults

`func NewRiskFeedbackBulkSummaryWithDefaults() *RiskFeedbackBulkSummary`

NewRiskFeedbackBulkSummaryWithDefaults instantiates a new RiskFeedbackBulkSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequested

`func (o *RiskFeedbackBulkSummary) GetRequested() int32`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RiskFeedbackBulkSummary) GetRequestedOk() (*int32, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RiskFeedbackBulkSummary) SetRequested(v int32)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *RiskFeedbackBulkSummary) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetErrors

`func (o *RiskFeedbackBulkSummary) GetErrors() int32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RiskFeedbackBulkSummary) GetErrorsOk() (*int32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RiskFeedbackBulkSummary) SetErrors(v int32)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RiskFeedbackBulkSummary) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetSuccessful

`func (o *RiskFeedbackBulkSummary) GetSuccessful() int32`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *RiskFeedbackBulkSummary) GetSuccessfulOk() (*int32, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *RiskFeedbackBulkSummary) SetSuccessful(v int32)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *RiskFeedbackBulkSummary) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetSkipped

`func (o *RiskFeedbackBulkSummary) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *RiskFeedbackBulkSummary) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *RiskFeedbackBulkSummary) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *RiskFeedbackBulkSummary) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


