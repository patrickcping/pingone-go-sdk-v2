# RiskFeedbackBulkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**RiskFeedbackBulkSummary**](RiskFeedbackBulkSummary.md) |  | [optional] 
**Responses** | Pointer to [**[]RiskFeedbackBulkItemResponse**](RiskFeedbackBulkItemResponse.md) | An array of per-item results, order-aligned with the request &#x60;items&#x60; array. | [optional] 
**Errors** | Pointer to [**[]RiskFeedbackBulkResponseErrorsInner**](RiskFeedbackBulkResponseErrorsInner.md) | A reference list of the failed per-item response elements (e.g. {\&quot;$ref\&quot;: \&quot;#/responses/1\&quot;}). | [optional] 

## Methods

### NewRiskFeedbackBulkResponse

`func NewRiskFeedbackBulkResponse() *RiskFeedbackBulkResponse`

NewRiskFeedbackBulkResponse instantiates a new RiskFeedbackBulkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackBulkResponseWithDefaults

`func NewRiskFeedbackBulkResponseWithDefaults() *RiskFeedbackBulkResponse`

NewRiskFeedbackBulkResponseWithDefaults instantiates a new RiskFeedbackBulkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *RiskFeedbackBulkResponse) GetSummary() RiskFeedbackBulkSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RiskFeedbackBulkResponse) GetSummaryOk() (*RiskFeedbackBulkSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RiskFeedbackBulkResponse) SetSummary(v RiskFeedbackBulkSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RiskFeedbackBulkResponse) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetResponses

`func (o *RiskFeedbackBulkResponse) GetResponses() []RiskFeedbackBulkItemResponse`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *RiskFeedbackBulkResponse) GetResponsesOk() (*[]RiskFeedbackBulkItemResponse, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *RiskFeedbackBulkResponse) SetResponses(v []RiskFeedbackBulkItemResponse)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *RiskFeedbackBulkResponse) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetErrors

`func (o *RiskFeedbackBulkResponse) GetErrors() []RiskFeedbackBulkResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RiskFeedbackBulkResponse) GetErrorsOk() (*[]RiskFeedbackBulkResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RiskFeedbackBulkResponse) SetErrors(v []RiskFeedbackBulkResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RiskFeedbackBulkResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


