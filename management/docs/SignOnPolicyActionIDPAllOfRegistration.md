# SignOnPolicyActionIDPAllOfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]
**Enabled** | **bool** | A boolean that specifies the enabled/disabled state of the policy action. The property is disabled by default when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option. | 
**Population** | Pointer to [**SignOnPolicyActionLoginAllOfRegistrationPopulation**](SignOnPolicyActionLoginAllOfRegistrationPopulation.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionIDPAllOfRegistration

`func NewSignOnPolicyActionIDPAllOfRegistration(enabled bool, ) *SignOnPolicyActionIDPAllOfRegistration`

NewSignOnPolicyActionIDPAllOfRegistration instantiates a new SignOnPolicyActionIDPAllOfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDPAllOfRegistrationWithDefaults

`func NewSignOnPolicyActionIDPAllOfRegistrationWithDefaults() *SignOnPolicyActionIDPAllOfRegistration`

NewSignOnPolicyActionIDPAllOfRegistrationWithDefaults instantiates a new SignOnPolicyActionIDPAllOfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetConfirmIdentityProviderAttributes() bool`

GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field if non-nil, zero value otherwise.

### GetConfirmIdentityProviderAttributesOk

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetConfirmIdentityProviderAttributesOk() (*bool, bool)`

GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDPAllOfRegistration) SetConfirmIdentityProviderAttributes(v bool)`

SetConfirmIdentityProviderAttributes sets ConfirmIdentityProviderAttributes field to given value.

### HasConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDPAllOfRegistration) HasConfirmIdentityProviderAttributes() bool`

HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.

### GetEnabled

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionIDPAllOfRegistration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPopulation

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetPopulation() SignOnPolicyActionLoginAllOfRegistrationPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *SignOnPolicyActionIDPAllOfRegistration) GetPopulationOk() (*SignOnPolicyActionLoginAllOfRegistrationPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *SignOnPolicyActionIDPAllOfRegistration) SetPopulation(v SignOnPolicyActionLoginAllOfRegistrationPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *SignOnPolicyActionIDPAllOfRegistration) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


