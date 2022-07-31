# SignOnPolicyActionCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Condition** | Pointer to [**SignOnPolicyActionCommonCondition**](SignOnPolicyActionCommonCondition.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy assignment resource’s unique identifier. | [optional] [readonly] 
**Priority** | **int32** | An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property. | 
**SignOnPolicy** | Pointer to [**SignOnPolicyActionCommonSignOnPolicy**](SignOnPolicyActionCommonSignOnPolicy.md) |  | [optional] 
**Type** | [**EnumSignOnPolicyType**](EnumSignOnPolicyType.md) |  | 

## Methods

### NewSignOnPolicyActionCommon

`func NewSignOnPolicyActionCommon(priority int32, type_ EnumSignOnPolicyType, ) *SignOnPolicyActionCommon`

NewSignOnPolicyActionCommon instantiates a new SignOnPolicyActionCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonWithDefaults

`func NewSignOnPolicyActionCommonWithDefaults() *SignOnPolicyActionCommon`

NewSignOnPolicyActionCommonWithDefaults instantiates a new SignOnPolicyActionCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionCommon) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionCommon) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionCommon) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionCommon) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionCommon) GetCondition() SignOnPolicyActionCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionCommon) GetConditionOk() (*SignOnPolicyActionCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionCommon) SetCondition(v SignOnPolicyActionCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionCommon) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionCommon) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionCommon) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionCommon) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionCommon) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionCommon) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionCommon) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionCommon) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionCommon) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionCommon) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionCommon) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionCommon) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionCommon) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionCommon) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionCommon) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


