# RiskFeedbackSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumRiskFeedbackSubjectType**](EnumRiskFeedbackSubjectType.md) |  | 
**Value** | Pointer to **string** | A string that specifies the identifier of the labeled entity. Required for every subject type except &#x60;TIME_RANGE&#x60; (which carries no value). For &#x60;RISK_EVALUATION_ID&#x60; it is the risk evaluation resource ID; for &#x60;EXTERNAL_TRANSACTION_ID&#x60; the caller&#39;s transaction ID; for &#x60;PINGONE_USER_ID&#x60; the PingOne user ID; for &#x60;EXTERNAL_USER_ID&#x60; the caller&#39;s user identifier; for &#x60;IP_ADDRESS&#x60; the IP address of the event. | [optional] 

## Methods

### NewRiskFeedbackSubject

`func NewRiskFeedbackSubject(type_ EnumRiskFeedbackSubjectType, ) *RiskFeedbackSubject`

NewRiskFeedbackSubject instantiates a new RiskFeedbackSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskFeedbackSubjectWithDefaults

`func NewRiskFeedbackSubjectWithDefaults() *RiskFeedbackSubject`

NewRiskFeedbackSubjectWithDefaults instantiates a new RiskFeedbackSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RiskFeedbackSubject) GetType() EnumRiskFeedbackSubjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskFeedbackSubject) GetTypeOk() (*EnumRiskFeedbackSubjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskFeedbackSubject) SetType(v EnumRiskFeedbackSubjectType)`

SetType sets Type field to given value.


### GetValue

`func (o *RiskFeedbackSubject) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RiskFeedbackSubject) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RiskFeedbackSubject) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RiskFeedbackSubject) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


