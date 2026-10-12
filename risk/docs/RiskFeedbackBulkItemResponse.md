# RiskFeedbackBulkItemResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**EnumRiskFeedbackBulkItemStatus**](EnumRiskFeedbackBulkItemStatus.md) |  | [optional] 
**Body** | Pointer to [**RiskFeedbackBulkItemResponseBody**](RiskFeedbackBulkItemResponseBody.md) |  | [optional] 

## Methods

### NewRiskFeedbackBulkItemResponse

`func NewRiskFeedbackBulkItemResponse() *RiskFeedbackBulkItemResponse`

NewRiskFeedbackBulkItemResponse instantiates a new RiskFeedbackBulkItemResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackBulkItemResponseWithDefaults

`func NewRiskFeedbackBulkItemResponseWithDefaults() *RiskFeedbackBulkItemResponse`

NewRiskFeedbackBulkItemResponseWithDefaults instantiates a new RiskFeedbackBulkItemResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RiskFeedbackBulkItemResponse) GetStatus() EnumRiskFeedbackBulkItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskFeedbackBulkItemResponse) GetStatusOk() (*EnumRiskFeedbackBulkItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskFeedbackBulkItemResponse) SetStatus(v EnumRiskFeedbackBulkItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RiskFeedbackBulkItemResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBody

`func (o *RiskFeedbackBulkItemResponse) GetBody() RiskFeedbackBulkItemResponseBody`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RiskFeedbackBulkItemResponse) GetBodyOk() (*RiskFeedbackBulkItemResponseBody, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RiskFeedbackBulkItemResponse) SetBody(v RiskFeedbackBulkItemResponseBody)`

SetBody sets Body field to given value.

### HasBody

`func (o *RiskFeedbackBulkItemResponse) HasBody() bool`

HasBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


