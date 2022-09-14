# SignOnPolicyActionPingIDWinLoginPasswordless

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Condition** | Pointer to [**SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy assignment resource’s unique identifier. | [optional] [readonly] 
**Priority** | **int32** | An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property. | 
**SignOnPolicy** | Pointer to [**SignOnPolicyActionCommonSignOnPolicy**](SignOnPolicyActionCommonSignOnPolicy.md) |  | [optional] 
**Type** | [**EnumSignOnPolicyType**](EnumSignOnPolicyType.md) |  | 
**UniqueUserAttribute** | [**SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute**](SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute.md) |  | 
**OfflineMode** | [**SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode**](SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode.md) |  | 

## Methods

### NewSignOnPolicyActionPingIDWinLoginPasswordless

`func NewSignOnPolicyActionPingIDWinLoginPasswordless(priority int32, type_ EnumSignOnPolicyType, uniqueUserAttribute SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute, offlineMode SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode, ) *SignOnPolicyActionPingIDWinLoginPasswordless`

NewSignOnPolicyActionPingIDWinLoginPasswordless instantiates a new SignOnPolicyActionPingIDWinLoginPasswordless object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionPingIDWinLoginPasswordlessWithDefaults

`func NewSignOnPolicyActionPingIDWinLoginPasswordlessWithDefaults() *SignOnPolicyActionPingIDWinLoginPasswordless`

NewSignOnPolicyActionPingIDWinLoginPasswordlessWithDefaults instantiates a new SignOnPolicyActionPingIDWinLoginPasswordless object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetCondition() SignOnPolicyActionCommonConditionOrOrInner`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


### GetUniqueUserAttribute

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetUniqueUserAttribute() SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute`

GetUniqueUserAttribute returns the UniqueUserAttribute field if non-nil, zero value otherwise.

### GetUniqueUserAttributeOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetUniqueUserAttributeOk() (*SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute, bool)`

GetUniqueUserAttributeOk returns a tuple with the UniqueUserAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUserAttribute

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetUniqueUserAttribute(v SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute)`

SetUniqueUserAttribute sets UniqueUserAttribute field to given value.


### GetOfflineMode

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetOfflineMode() SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode`

GetOfflineMode returns the OfflineMode field if non-nil, zero value otherwise.

### GetOfflineModeOk

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetOfflineModeOk() (*SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode, bool)`

GetOfflineModeOk returns a tuple with the OfflineMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfflineMode

`func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetOfflineMode(v SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode)`

SetOfflineMode sets OfflineMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


