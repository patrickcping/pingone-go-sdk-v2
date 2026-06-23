# RiskPolicySetTargetsConditionAndInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to **[]string** | Used to list the strings for the target transaction types, user groups, or applications, for example, &#x60;[\&quot;AUTHENTICATION\&quot;, \&quot;AUTHORIZATION\&quot;]&#x60;. For transaction types, should contain one or more of the following: &#x60;REGISTRATION&#x60;, &#x60;AUTHENTICATION&#x60;, &#x60;ACCESS&#x60;, &#x60;AUTHORIZATION&#x60;, &#x60;TRANSACTION&#x60;. For user groups, should contain the name of one or more user groups. For applications, should contain the PingOne ID of one or more applications. | [optional] 
**Contains** | Pointer to **string** | The event attribute that is checked for the values specified in the &#x60;list&#x60; parameter. For transaction type, this should be set to &#x60;${event.flow.type}&#x60;. For user groups, this should be set to &#x60;${event.user.groups}&#x60;. For applications, this should be set to &#x60;${event.targetResource.id}&#x60;. | [optional] 
**Type** | Pointer to [**EnumRiskPolicySetTargetsConditionType**](EnumRiskPolicySetTargetsConditionType.md) |  | [optional] 

## Methods

### NewRiskPolicySetTargetsConditionAndInner

`func NewRiskPolicySetTargetsConditionAndInner() *RiskPolicySetTargetsConditionAndInner`

NewRiskPolicySetTargetsConditionAndInner instantiates a new RiskPolicySetTargetsConditionAndInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskPolicySetTargetsConditionAndInnerWithDefaults

`func NewRiskPolicySetTargetsConditionAndInnerWithDefaults() *RiskPolicySetTargetsConditionAndInner`

NewRiskPolicySetTargetsConditionAndInnerWithDefaults instantiates a new RiskPolicySetTargetsConditionAndInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *RiskPolicySetTargetsConditionAndInner) GetList() []string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *RiskPolicySetTargetsConditionAndInner) GetListOk() (*[]string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *RiskPolicySetTargetsConditionAndInner) SetList(v []string)`

SetList sets List field to given value.

### HasList

`func (o *RiskPolicySetTargetsConditionAndInner) HasList() bool`

HasList returns a boolean if a field has been set.

### GetContains

`func (o *RiskPolicySetTargetsConditionAndInner) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *RiskPolicySetTargetsConditionAndInner) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *RiskPolicySetTargetsConditionAndInner) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *RiskPolicySetTargetsConditionAndInner) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetType

`func (o *RiskPolicySetTargetsConditionAndInner) GetType() EnumRiskPolicySetTargetsConditionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskPolicySetTargetsConditionAndInner) GetTypeOk() (*EnumRiskPolicySetTargetsConditionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskPolicySetTargetsConditionAndInner) SetType(v EnumRiskPolicySetTargetsConditionType)`

SetType sets Type field to given value.

### HasType

`func (o *RiskPolicySetTargetsConditionAndInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


