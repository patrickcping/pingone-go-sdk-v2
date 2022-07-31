# SignOnPolicyActionCommonConditionGreater

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Greater** | **int32** | An integer that specifies the maximum number of seconds to wait since the last sign on before prompting for a new sign-on action. | 
**SecondsSince** | **string** | A string representing a condition variable.  For more information, see documenation at https://apidocs.pingidentity.com/pingone/platform/v1/api/#sign-on-policy-actions . | 

## Methods

### NewSignOnPolicyActionCommonConditionGreater

`func NewSignOnPolicyActionCommonConditionGreater(greater int32, secondsSince string, ) *SignOnPolicyActionCommonConditionGreater`

NewSignOnPolicyActionCommonConditionGreater instantiates a new SignOnPolicyActionCommonConditionGreater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionCommonConditionGreaterWithDefaults

`func NewSignOnPolicyActionCommonConditionGreaterWithDefaults() *SignOnPolicyActionCommonConditionGreater`

NewSignOnPolicyActionCommonConditionGreaterWithDefaults instantiates a new SignOnPolicyActionCommonConditionGreater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreater

`func (o *SignOnPolicyActionCommonConditionGreater) GetGreater() int32`

GetGreater returns the Greater field if non-nil, zero value otherwise.

### GetGreaterOk

`func (o *SignOnPolicyActionCommonConditionGreater) GetGreaterOk() (*int32, bool)`

GetGreaterOk returns a tuple with the Greater field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreater

`func (o *SignOnPolicyActionCommonConditionGreater) SetGreater(v int32)`

SetGreater sets Greater field to given value.


### GetSecondsSince

`func (o *SignOnPolicyActionCommonConditionGreater) GetSecondsSince() string`

GetSecondsSince returns the SecondsSince field if non-nil, zero value otherwise.

### GetSecondsSinceOk

`func (o *SignOnPolicyActionCommonConditionGreater) GetSecondsSinceOk() (*string, bool)`

GetSecondsSinceOk returns a tuple with the SecondsSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsSince

`func (o *SignOnPolicyActionCommonConditionGreater) SetSecondsSince(v string)`

SetSecondsSince sets SecondsSince field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


