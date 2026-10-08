# RiskFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the feedback resource&#39;s unique identifier. | [optional] [readonly] 
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

## Methods

### NewRiskFeedback

`func NewRiskFeedback(subject RiskFeedbackSubject, verdict EnumRiskFeedbackVerdict, occurredFrom time.Time, ) *RiskFeedback`

NewRiskFeedback instantiates a new RiskFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackWithDefaults

`func NewRiskFeedbackWithDefaults() *RiskFeedback`

NewRiskFeedbackWithDefaults instantiates a new RiskFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskFeedback) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskFeedback) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskFeedback) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskFeedback) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *RiskFeedback) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskFeedback) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskFeedback) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskFeedback) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *RiskFeedback) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *RiskFeedback) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *RiskFeedback) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *RiskFeedback) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSubject

`func (o *RiskFeedback) GetSubject() RiskFeedbackSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiskFeedback) GetSubjectOk() (*RiskFeedbackSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiskFeedback) SetSubject(v RiskFeedbackSubject)`

SetSubject sets Subject field to given value.


### GetVerdict

`func (o *RiskFeedback) GetVerdict() EnumRiskFeedbackVerdict`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *RiskFeedback) GetVerdictOk() (*EnumRiskFeedbackVerdict, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *RiskFeedback) SetVerdict(v EnumRiskFeedbackVerdict)`

SetVerdict sets Verdict field to given value.


### GetOccurredFrom

`func (o *RiskFeedback) GetOccurredFrom() time.Time`

GetOccurredFrom returns the OccurredFrom field if non-nil, zero value otherwise.

### GetOccurredFromOk

`func (o *RiskFeedback) GetOccurredFromOk() (*time.Time, bool)`

GetOccurredFromOk returns a tuple with the OccurredFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredFrom

`func (o *RiskFeedback) SetOccurredFrom(v time.Time)`

SetOccurredFrom sets OccurredFrom field to given value.


### GetOccurredTo

`func (o *RiskFeedback) GetOccurredTo() time.Time`

GetOccurredTo returns the OccurredTo field if non-nil, zero value otherwise.

### GetOccurredToOk

`func (o *RiskFeedback) GetOccurredToOk() (*time.Time, bool)`

GetOccurredToOk returns a tuple with the OccurredTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTo

`func (o *RiskFeedback) SetOccurredTo(v time.Time)`

SetOccurredTo sets OccurredTo field to given value.

### HasOccurredTo

`func (o *RiskFeedback) HasOccurredTo() bool`

HasOccurredTo returns a boolean if a field has been set.

### GetTrigger

`func (o *RiskFeedback) GetTrigger() EnumRiskFeedbackTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *RiskFeedback) GetTriggerOk() (*EnumRiskFeedbackTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *RiskFeedback) SetTrigger(v EnumRiskFeedbackTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *RiskFeedback) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetFlowType

`func (o *RiskFeedback) GetFlowType() EnumRiskFeedbackFlowType`

GetFlowType returns the FlowType field if non-nil, zero value otherwise.

### GetFlowTypeOk

`func (o *RiskFeedback) GetFlowTypeOk() (*EnumRiskFeedbackFlowType, bool)`

GetFlowTypeOk returns a tuple with the FlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowType

`func (o *RiskFeedback) SetFlowType(v EnumRiskFeedbackFlowType)`

SetFlowType sets FlowType field to given value.

### HasFlowType

`func (o *RiskFeedback) HasFlowType() bool`

HasFlowType returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *RiskFeedback) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *RiskFeedback) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *RiskFeedback) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *RiskFeedback) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetReason

`func (o *RiskFeedback) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RiskFeedback) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RiskFeedback) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RiskFeedback) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAnalyst

`func (o *RiskFeedback) GetAnalyst() string`

GetAnalyst returns the Analyst field if non-nil, zero value otherwise.

### GetAnalystOk

`func (o *RiskFeedback) GetAnalystOk() (*string, bool)`

GetAnalystOk returns a tuple with the Analyst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyst

`func (o *RiskFeedback) SetAnalyst(v string)`

SetAnalyst sets Analyst field to given value.

### HasAnalyst

`func (o *RiskFeedback) HasAnalyst() bool`

HasAnalyst returns a boolean if a field has been set.

### GetConfidence

`func (o *RiskFeedback) GetConfidence() EnumRiskFeedbackConfidence`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RiskFeedback) GetConfidenceOk() (*EnumRiskFeedbackConfidence, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RiskFeedback) SetConfidence(v EnumRiskFeedbackConfidence)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RiskFeedback) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetFraudType

`func (o *RiskFeedback) GetFraudType() string`

GetFraudType returns the FraudType field if non-nil, zero value otherwise.

### GetFraudTypeOk

`func (o *RiskFeedback) GetFraudTypeOk() (*string, bool)`

GetFraudTypeOk returns a tuple with the FraudType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraudType

`func (o *RiskFeedback) SetFraudType(v string)`

SetFraudType sets FraudType field to given value.

### HasFraudType

`func (o *RiskFeedback) HasFraudType() bool`

HasFraudType returns a boolean if a field has been set.

### GetAttributes

`func (o *RiskFeedback) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RiskFeedback) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RiskFeedback) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RiskFeedback) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RiskFeedback) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RiskFeedback) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RiskFeedback) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RiskFeedback) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RiskFeedback) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RiskFeedback) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RiskFeedback) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RiskFeedback) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


