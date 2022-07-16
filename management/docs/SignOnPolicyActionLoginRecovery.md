# SignOnPolicyActionLoginRecovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies the enabled/disabled state of the account recovery action. The default is disabled when creating a new policy. When enabled, it allows the use of the forgot password flow. | 

## Methods

### NewSignOnPolicyActionLoginRecovery

`func NewSignOnPolicyActionLoginRecovery(enabled bool, ) *SignOnPolicyActionLoginRecovery`

NewSignOnPolicyActionLoginRecovery instantiates a new SignOnPolicyActionLoginRecovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionLoginRecoveryWithDefaults

`func NewSignOnPolicyActionLoginRecoveryWithDefaults() *SignOnPolicyActionLoginRecovery`

NewSignOnPolicyActionLoginRecoveryWithDefaults instantiates a new SignOnPolicyActionLoginRecovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionLoginRecovery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionLoginRecovery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionLoginRecovery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


