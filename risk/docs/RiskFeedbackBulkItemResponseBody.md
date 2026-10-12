# RiskFeedbackBulkItemResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A unique identifier that is stored in log files and always included in an error response. This value can be used to track the error received by the client, with server-side activity included for troubleshooting purposes. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Subject** | [**RiskFeedbackSubject**](RiskFeedbackSubject.md) |  | 
**Verdict** | [**EnumRiskFeedbackVerdict**](EnumRiskFeedbackVerdict.md) |  | 
**OccurredFrom** | **time.Time** | A string that specifies the start of the period the label refers to, in UTC (format &#x60;YYYY-MM-DDThh:mm:ss.SSSZ&#x60;, milliseconds required). This is a required property. For &#x60;TIME_RANGE&#x60; subjects this is the beginning of the labeled period; otherwise it is the event time. Must be at most 3 days before &#x60;occurredTo&#x60; for &#x60;TIME_RANGE&#x60; subjects. | 
**OccurredTo** | Pointer to **time.Time** | A string that specifies the end of the labeled period in UTC (format &#x60;YYYY-MM-DDThh:mm:ss.SSSZ&#x60;, milliseconds required). Required for &#x60;TIME_RANGE&#x60; subjects; omitted otherwise. Must be at most 3 days after &#x60;occurredFrom&#x60;. | [optional] 
**Trigger** | Pointer to [**EnumRiskFeedbackTrigger**](EnumRiskFeedbackTrigger.md) |  | [optional] 
**FlowType** | Pointer to [**EnumRiskFeedbackFlowType**](EnumRiskFeedbackFlowType.md) |  | [optional] 
**ConfirmedAt** | Pointer to **time.Time** | A string that specifies when the outcome was confirmed, in UTC (format &#x60;YYYY-MM-DDThh:mm:ss.SSSZ&#x60;, milliseconds required). | [optional] 
**Reason** | Pointer to **string** | A string that specifies why the label was assigned. Maximum size is 200 characters. | [optional] 
**Analyst** | Pointer to **string** | A string that identifies the analyst who confirmed the label. Maximum size is 50 characters. | [optional] 
**Confidence** | Pointer to [**EnumRiskFeedbackConfidence**](EnumRiskFeedbackConfidence.md) |  | [optional] 
**FraudType** | Pointer to **string** | A string that specifies the fraud category, e.g. &#x60;BOT_FARM&#x60;. Maximum size is 50 characters. | [optional] 
**Attributes** | Pointer to **map[string]string** | A flat map of up to 20 context attributes (key max 100 characters, value max 300 characters) attached to the label. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the feedback was created (format ISO-8061). | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the feedback was last updated (format ISO-8061). | [optional] [readonly] 
**Code** | Pointer to **string** | A general fault code which the client must handle to provide all exception handling routines and to localize messages for users. This code is common across all PingOne services and is human readable (such as a defined constant rather than a number). | [optional] 
**Message** | Pointer to **string** | A short description of the error. This message is intended to assist with debugging and is returned in English only. | [optional] 
**Details** | Pointer to [**[]P1ErrorDetailsInner**](P1ErrorDetailsInner.md) | Additional details about the error. Optional information to help resolve the error and to display to users. | [optional] 

## Methods

### NewRiskFeedbackBulkItemResponseBody

`func NewRiskFeedbackBulkItemResponseBody(subject RiskFeedbackSubject, verdict EnumRiskFeedbackVerdict, occurredFrom time.Time, ) *RiskFeedbackBulkItemResponseBody`

NewRiskFeedbackBulkItemResponseBody instantiates a new RiskFeedbackBulkItemResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackBulkItemResponseBodyWithDefaults

`func NewRiskFeedbackBulkItemResponseBodyWithDefaults() *RiskFeedbackBulkItemResponseBody`

NewRiskFeedbackBulkItemResponseBodyWithDefaults instantiates a new RiskFeedbackBulkItemResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskFeedbackBulkItemResponseBody) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskFeedbackBulkItemResponseBody) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskFeedbackBulkItemResponseBody) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskFeedbackBulkItemResponseBody) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskFeedbackBulkItemResponseBody) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskFeedbackBulkItemResponseBody) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskFeedbackBulkItemResponseBody) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskFeedbackBulkItemResponseBody) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskFeedbackBulkItemResponseBody) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskFeedbackBulkItemResponseBody) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskFeedbackBulkItemResponseBody) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskFeedbackBulkItemResponseBody) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSubject

`func (o *RiskFeedbackBulkItemResponseBody) GetSubject() RiskFeedbackSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskFeedbackBulkItemResponseBody) GetSubjectOk() (*RiskFeedbackSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskFeedbackBulkItemResponseBody) SetSubject(v RiskFeedbackSubject)`

SetSubject sets Subject field to given value.


### GetVerdict

`func (o *RiskFeedbackBulkItemResponseBody) GetVerdict() EnumRiskFeedbackVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *RiskFeedbackBulkItemResponseBody) GetVerdictOk() (*EnumRiskFeedbackVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *RiskFeedbackBulkItemResponseBody) SetVerdict(v EnumRiskFeedbackVerdict)`

SetVerdict sets Verdict field to given value.


### GetOccurredFrom

`func (o *RiskFeedbackBulkItemResponseBody) GetOccurredFrom() time.Time`

GetOccurredFrom returns the OccurredFrom field if non-nil, zero value otherwise.

### GetOccurredFromOk

`func (o *RiskFeedbackBulkItemResponseBody) GetOccurredFromOk() (*time.Time, bool)`

GetOccurredFromOk returns a tuple with the OccurredFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredFrom

`func (o *RiskFeedbackBulkItemResponseBody) SetOccurredFrom(v time.Time)`

SetOccurredFrom sets OccurredFrom field to given value.


### GetOccurredTo

`func (o *RiskFeedbackBulkItemResponseBody) GetOccurredTo() time.Time`

GetOccurredTo returns the OccurredTo field if non-nil, zero value otherwise.

### GetOccurredToOk

`func (o *RiskFeedbackBulkItemResponseBody) GetOccurredToOk() (*time.Time, bool)`

GetOccurredToOk returns a tuple with the OccurredTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTo

`func (o *RiskFeedbackBulkItemResponseBody) SetOccurredTo(v time.Time)`

SetOccurredTo sets OccurredTo field to given value.

### HasOccurredTo

`func (o *RiskFeedbackBulkItemResponseBody) HasOccurredTo() bool`

HasOccurredTo returns a boolean if a field has been set.

### GetTrigger

`func (o *RiskFeedbackBulkItemResponseBody) GetTrigger() EnumRiskFeedbackTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RiskFeedbackBulkItemResponseBody) GetTriggerOk() (*EnumRiskFeedbackTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RiskFeedbackBulkItemResponseBody) SetTrigger(v EnumRiskFeedbackTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *RiskFeedbackBulkItemResponseBody) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetFlowType

`func (o *RiskFeedbackBulkItemResponseBody) GetFlowType() EnumRiskFeedbackFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *RiskFeedbackBulkItemResponseBody) GetFlowTypeOk() (*EnumRiskFeedbackFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *RiskFeedbackBulkItemResponseBody) SetFlowType(v EnumRiskFeedbackFlowType)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *RiskFeedbackBulkItemResponseBody) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *RiskFeedbackBulkItemResponseBody) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *RiskFeedbackBulkItemResponseBody) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *RiskFeedbackBulkItemResponseBody) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *RiskFeedbackBulkItemResponseBody) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetReason

`func (o *RiskFeedbackBulkItemResponseBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskFeedbackBulkItemResponseBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskFeedbackBulkItemResponseBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskFeedbackBulkItemResponseBody) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAnalyst

`func (o *RiskFeedbackBulkItemResponseBody) GetAnalyst() string`

GetAnalyst returns the Analyst field if non-nil, zero value otherwise.

### GetAnalystOk

`func (o *RiskFeedbackBulkItemResponseBody) GetAnalystOk() (*string, bool)`

GetAnalystOk returns a tuple with the Analyst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyst

`func (o *RiskFeedbackBulkItemResponseBody) SetAnalyst(v string)`

SetAnalyst sets Analyst field to given value.

### HasAnalyst

`func (o *RiskFeedbackBulkItemResponseBody) HasAnalyst() bool`

HasAnalyst returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskFeedbackBulkItemResponseBody) GetConfidence() EnumRiskFeedbackConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskFeedbackBulkItemResponseBody) GetConfidenceOk() (*EnumRiskFeedbackConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskFeedbackBulkItemResponseBody) SetConfidence(v EnumRiskFeedbackConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskFeedbackBulkItemResponseBody) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetFraudType

`func (o *RiskFeedbackBulkItemResponseBody) GetFraudType() string`

GetFraudType returns the FraudType field if non-nil, zero value otherwise.

### GetFraudTypeOk

`func (o *RiskFeedbackBulkItemResponseBody) GetFraudTypeOk() (*string, bool)`

GetFraudTypeOk returns a tuple with the FraudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudType

`func (o *RiskFeedbackBulkItemResponseBody) SetFraudType(v string)`

SetFraudType sets FraudType field to given value.

### HasFraudType

`func (o *RiskFeedbackBulkItemResponseBody) HasFraudType() bool`

HasFraudType returns a boolean if a field has been set.

### GetAttributes

`func (o *RiskFeedbackBulkItemResponseBody) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RiskFeedbackBulkItemResponseBody) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RiskFeedbackBulkItemResponseBody) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RiskFeedbackBulkItemResponseBody) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskFeedbackBulkItemResponseBody) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskFeedbackBulkItemResponseBody) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskFeedbackBulkItemResponseBody) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskFeedbackBulkItemResponseBody) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskFeedbackBulkItemResponseBody) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskFeedbackBulkItemResponseBody) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskFeedbackBulkItemResponseBody) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskFeedbackBulkItemResponseBody) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCode

`func (o *RiskFeedbackBulkItemResponseBody) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RiskFeedbackBulkItemResponseBody) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RiskFeedbackBulkItemResponseBody) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RiskFeedbackBulkItemResponseBody) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *RiskFeedbackBulkItemResponseBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RiskFeedbackBulkItemResponseBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RiskFeedbackBulkItemResponseBody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RiskFeedbackBulkItemResponseBody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetails

`func (o *RiskFeedbackBulkItemResponseBody) GetDetails() []P1ErrorDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RiskFeedbackBulkItemResponseBody) GetDetailsOk() (*[]P1ErrorDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RiskFeedbackBulkItemResponseBody) SetDetails(v []P1ErrorDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RiskFeedbackBulkItemResponseBody) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


