# SignOnPolicyActionLoginAllOfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies the enabled/disabled state of the policy action. The default is disabled when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option. | 
**External** | Pointer to [**SignOnPolicyActionLoginAllOfRegistrationExternal**](SignOnPolicyActionLoginAllOfRegistrationExternal.md) |  | [optional] 
**Population** | Pointer to [**SignOnPolicyActionLoginAllOfRegistrationPopulation**](SignOnPolicyActionLoginAllOfRegistrationPopulation.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionLoginAllOfRegistration

`func NewSignOnPolicyActionLoginAllOfRegistration(enabled bool, ) *SignOnPolicyActionLoginAllOfRegistration`

NewSignOnPolicyActionLoginAllOfRegistration instantiates a new SignOnPolicyActionLoginAllOfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionLoginAllOfRegistrationWithDefaults

`func NewSignOnPolicyActionLoginAllOfRegistrationWithDefaults() *SignOnPolicyActionLoginAllOfRegistration`

NewSignOnPolicyActionLoginAllOfRegistrationWithDefaults instantiates a new SignOnPolicyActionLoginAllOfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionLoginAllOfRegistration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExternal

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetExternal() SignOnPolicyActionLoginAllOfRegistrationExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetExternalOk() (*SignOnPolicyActionLoginAllOfRegistrationExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *SignOnPolicyActionLoginAllOfRegistration) SetExternal(v SignOnPolicyActionLoginAllOfRegistrationExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *SignOnPolicyActionLoginAllOfRegistration) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetPopulation

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetPopulation() SignOnPolicyActionLoginAllOfRegistrationPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *SignOnPolicyActionLoginAllOfRegistration) GetPopulationOk() (*SignOnPolicyActionLoginAllOfRegistrationPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *SignOnPolicyActionLoginAllOfRegistration) SetPopulation(v SignOnPolicyActionLoginAllOfRegistrationPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *SignOnPolicyActionLoginAllOfRegistration) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


