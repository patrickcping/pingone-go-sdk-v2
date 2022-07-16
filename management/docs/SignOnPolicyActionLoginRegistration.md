# SignOnPolicyActionLoginRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies the enabled/disabled state of the policy action. The default is disabled when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option. | 
**External** | Pointer to [**SignOnPolicyActionLoginRegistrationExternal**](SignOnPolicyActionLoginRegistrationExternal.md) |  | [optional] 
**Population** | Pointer to [**SignOnPolicyActionLoginRegistrationPopulation**](SignOnPolicyActionLoginRegistrationPopulation.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionLoginRegistration

`func NewSignOnPolicyActionLoginRegistration(enabled bool, ) *SignOnPolicyActionLoginRegistration`

NewSignOnPolicyActionLoginRegistration instantiates a new SignOnPolicyActionLoginRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionLoginRegistrationWithDefaults

`func NewSignOnPolicyActionLoginRegistrationWithDefaults() *SignOnPolicyActionLoginRegistration`

NewSignOnPolicyActionLoginRegistrationWithDefaults instantiates a new SignOnPolicyActionLoginRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionLoginRegistration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionLoginRegistration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionLoginRegistration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExternal

`func (o *SignOnPolicyActionLoginRegistration) GetExternal() SignOnPolicyActionLoginRegistrationExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *SignOnPolicyActionLoginRegistration) GetExternalOk() (*SignOnPolicyActionLoginRegistrationExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *SignOnPolicyActionLoginRegistration) SetExternal(v SignOnPolicyActionLoginRegistrationExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *SignOnPolicyActionLoginRegistration) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetPopulation

`func (o *SignOnPolicyActionLoginRegistration) GetPopulation() SignOnPolicyActionLoginRegistrationPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *SignOnPolicyActionLoginRegistration) GetPopulationOk() (*SignOnPolicyActionLoginRegistrationPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *SignOnPolicyActionLoginRegistration) SetPopulation(v SignOnPolicyActionLoginRegistrationPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *SignOnPolicyActionLoginRegistration) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


