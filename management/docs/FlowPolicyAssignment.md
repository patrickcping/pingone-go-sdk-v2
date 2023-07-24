# FlowPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the flow policy assignment resource&#39;s unique identifier. | [optional] [readonly] 
**Application** | Pointer to [**ObjectApplication**](ObjectApplication.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**FlowPolicy** | [**FlowPolicyAssignmentFlowPolicy**](FlowPolicyAssignmentFlowPolicy.md) |  | 
**Priority** | **int32** | The order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. | 

## Methods

### NewFlowPolicyAssignment

`func NewFlowPolicyAssignment(flowPolicy FlowPolicyAssignmentFlowPolicy, priority int32, ) *FlowPolicyAssignment`

NewFlowPolicyAssignment instantiates a new FlowPolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowPolicyAssignmentWithDefaults

`func NewFlowPolicyAssignmentWithDefaults() *FlowPolicyAssignment`

NewFlowPolicyAssignmentWithDefaults instantiates a new FlowPolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *FlowPolicyAssignment) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlowPolicyAssignment) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlowPolicyAssignment) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlowPolicyAssignment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *FlowPolicyAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowPolicyAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowPolicyAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowPolicyAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApplication

`func (o *FlowPolicyAssignment) GetApplication() ObjectApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *FlowPolicyAssignment) GetApplicationOk() (*ObjectApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *FlowPolicyAssignment) SetApplication(v ObjectApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *FlowPolicyAssignment) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEnvironment

`func (o *FlowPolicyAssignment) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FlowPolicyAssignment) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FlowPolicyAssignment) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FlowPolicyAssignment) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFlowPolicy

`func (o *FlowPolicyAssignment) GetFlowPolicy() FlowPolicyAssignmentFlowPolicy`

GetFlowPolicy returns the FlowPolicy field if non-nil, zero value otherwise.

### GetFlowPolicyOk

`func (o *FlowPolicyAssignment) GetFlowPolicyOk() (*FlowPolicyAssignmentFlowPolicy, bool)`

GetFlowPolicyOk returns a tuple with the FlowPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowPolicy

`func (o *FlowPolicyAssignment) SetFlowPolicy(v FlowPolicyAssignmentFlowPolicy)`

SetFlowPolicy sets FlowPolicy field to given value.


### GetPriority

`func (o *FlowPolicyAssignment) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *FlowPolicyAssignment) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *FlowPolicyAssignment) SetPriority(v int32)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


