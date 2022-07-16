# RiskEvaluationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to [**RiskEvaluationEventBrowser**](RiskEvaluationEventBrowser.md) |  | [optional] 
**CompletionStatus** | Pointer to **string** | A string that specifies the state of the transaction. Options are FAILED, IN_PROGRESS, and SUCCESS. If a value is not provided, the value defaults to IN_PROGRESS. The value of this property can be changed only if its current state is IN_PROGRESS. | [optional] 
**EvaluatedFactors** | Pointer to [**RiskEvaluationEventEvaluatedFactors**](RiskEvaluationEventEvaluatedFactors.md) |  | [optional] 
**Ip** | **string** | A string that specifies the origin IP address of the authentication flow. This is a required property. | 
**Flow** | Pointer to [**RiskEvaluationEventFlow**](RiskEvaluationEventFlow.md) |  | [optional] 
**Origin** | Pointer to **string** | A string that specifies the calling service. | [optional] 
**Session** | Pointer to [**RiskEvaluationEventSession**](RiskEvaluationEventSession.md) |  | [optional] 
**TargetResource** | Pointer to [**RiskEvaluationEventTargetResource**](RiskEvaluationEventTargetResource.md) |  | [optional] 
**User** | [**RiskEvaluationEventUser**](RiskEvaluationEventUser.md) |  | 
**SharingType** | Pointer to **string** | A string that specifies the device sharing type. Options are UNSPECIFIED, SHARED, and PRIVATE. | [optional] 

## Methods

### NewRiskEvaluationEvent

`func NewRiskEvaluationEvent(ip string, user RiskEvaluationEventUser, ) *RiskEvaluationEvent`

NewRiskEvaluationEvent instantiates a new RiskEvaluationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEvaluationEventWithDefaults

`func NewRiskEvaluationEventWithDefaults() *RiskEvaluationEvent`

NewRiskEvaluationEventWithDefaults instantiates a new RiskEvaluationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *RiskEvaluationEvent) GetBrowser() RiskEvaluationEventBrowser`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *RiskEvaluationEvent) GetBrowserOk() (*RiskEvaluationEventBrowser, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *RiskEvaluationEvent) SetBrowser(v RiskEvaluationEventBrowser)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *RiskEvaluationEvent) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### GetCompletionStatus

`func (o *RiskEvaluationEvent) GetCompletionStatus() string`

GetCompletionStatus returns the CompletionStatus field if non-nil, zero value otherwise.

### GetCompletionStatusOk

`func (o *RiskEvaluationEvent) GetCompletionStatusOk() (*string, bool)`

GetCompletionStatusOk returns a tuple with the CompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionStatus

`func (o *RiskEvaluationEvent) SetCompletionStatus(v string)`

SetCompletionStatus sets CompletionStatus field to given value.

### HasCompletionStatus

`func (o *RiskEvaluationEvent) HasCompletionStatus() bool`

HasCompletionStatus returns a boolean if a field has been set.

### GetEvaluatedFactors

`func (o *RiskEvaluationEvent) GetEvaluatedFactors() RiskEvaluationEventEvaluatedFactors`

GetEvaluatedFactors returns the EvaluatedFactors field if non-nil, zero value otherwise.

### GetEvaluatedFactorsOk

`func (o *RiskEvaluationEvent) GetEvaluatedFactorsOk() (*RiskEvaluationEventEvaluatedFactors, bool)`

GetEvaluatedFactorsOk returns a tuple with the EvaluatedFactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedFactors

`func (o *RiskEvaluationEvent) SetEvaluatedFactors(v RiskEvaluationEventEvaluatedFactors)`

SetEvaluatedFactors sets EvaluatedFactors field to given value.

### HasEvaluatedFactors

`func (o *RiskEvaluationEvent) HasEvaluatedFactors() bool`

HasEvaluatedFactors returns a boolean if a field has been set.

### GetIp

`func (o *RiskEvaluationEvent) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *RiskEvaluationEvent) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *RiskEvaluationEvent) SetIp(v string)`

SetIp sets Ip field to given value.


### GetFlow

`func (o *RiskEvaluationEvent) GetFlow() RiskEvaluationEventFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *RiskEvaluationEvent) GetFlowOk() (*RiskEvaluationEventFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *RiskEvaluationEvent) SetFlow(v RiskEvaluationEventFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *RiskEvaluationEvent) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetOrigin

`func (o *RiskEvaluationEvent) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *RiskEvaluationEvent) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *RiskEvaluationEvent) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *RiskEvaluationEvent) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetSession

`func (o *RiskEvaluationEvent) GetSession() RiskEvaluationEventSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *RiskEvaluationEvent) GetSessionOk() (*RiskEvaluationEventSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *RiskEvaluationEvent) SetSession(v RiskEvaluationEventSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *RiskEvaluationEvent) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetTargetResource

`func (o *RiskEvaluationEvent) GetTargetResource() RiskEvaluationEventTargetResource`

GetTargetResource returns the TargetResource field if non-nil, zero value otherwise.

### GetTargetResourceOk

`func (o *RiskEvaluationEvent) GetTargetResourceOk() (*RiskEvaluationEventTargetResource, bool)`

GetTargetResourceOk returns a tuple with the TargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResource

`func (o *RiskEvaluationEvent) SetTargetResource(v RiskEvaluationEventTargetResource)`

SetTargetResource sets TargetResource field to given value.

### HasTargetResource

`func (o *RiskEvaluationEvent) HasTargetResource() bool`

HasTargetResource returns a boolean if a field has been set.

### GetUser

`func (o *RiskEvaluationEvent) GetUser() RiskEvaluationEventUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RiskEvaluationEvent) GetUserOk() (*RiskEvaluationEventUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RiskEvaluationEvent) SetUser(v RiskEvaluationEventUser)`

SetUser sets User field to given value.


### GetSharingType

`func (o *RiskEvaluationEvent) GetSharingType() string`

GetSharingType returns the SharingType field if non-nil, zero value otherwise.

### GetSharingTypeOk

`func (o *RiskEvaluationEvent) GetSharingTypeOk() (*string, bool)`

GetSharingTypeOk returns a tuple with the SharingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingType

`func (o *RiskEvaluationEvent) SetSharingType(v string)`

SetSharingType sets SharingType field to given value.

### HasSharingType

`func (o *RiskEvaluationEvent) HasSharingType() bool`

HasSharingType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


