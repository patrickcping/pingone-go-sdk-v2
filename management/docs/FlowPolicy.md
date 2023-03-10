# FlowPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Application** | Pointer to [**FlowPolicyApplication**](FlowPolicyApplication.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Trigger** | Pointer to [**FlowPolicyTrigger**](FlowPolicyTrigger.md) |  | [optional] 

## Methods

### NewFlowPolicy

`func NewFlowPolicy() *FlowPolicy`

NewFlowPolicy instantiates a new FlowPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowPolicyWithDefaults

`func NewFlowPolicyWithDefaults() *FlowPolicy`

NewFlowPolicyWithDefaults instantiates a new FlowPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FlowPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FlowPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *FlowPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FlowPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FlowPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FlowPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetApplication

`func (o *FlowPolicy) GetApplication() FlowPolicyApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *FlowPolicy) GetApplicationOk() (*FlowPolicyApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *FlowPolicy) SetApplication(v FlowPolicyApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *FlowPolicy) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEnabled

`func (o *FlowPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlowPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlowPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlowPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *FlowPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrigger

`func (o *FlowPolicy) GetTrigger() FlowPolicyTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *FlowPolicy) GetTriggerOk() (*FlowPolicyTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *FlowPolicy) SetTrigger(v FlowPolicyTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *FlowPolicy) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


