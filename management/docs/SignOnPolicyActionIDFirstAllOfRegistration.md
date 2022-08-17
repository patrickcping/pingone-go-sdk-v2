# SignOnPolicyActionIDFirstAllOfRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies the enabled/disabled state of the policy action. The default is disabled when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option. | 
**External** | Pointer to [**SignOnPolicyActionLoginAllOfRegistrationExternal**](SignOnPolicyActionLoginAllOfRegistrationExternal.md) |  | [optional] 
**Population** | Pointer to [**SignOnPolicyActionLoginAllOfRegistrationPopulation**](SignOnPolicyActionLoginAllOfRegistrationPopulation.md) |  | [optional] 
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]

## Methods

### NewSignOnPolicyActionIDFirstAllOfRegistration

`func NewSignOnPolicyActionIDFirstAllOfRegistration(enabled bool, ) *SignOnPolicyActionIDFirstAllOfRegistration`

NewSignOnPolicyActionIDFirstAllOfRegistration instantiates a new SignOnPolicyActionIDFirstAllOfRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDFirstAllOfRegistrationWithDefaults

`func NewSignOnPolicyActionIDFirstAllOfRegistrationWithDefaults() *SignOnPolicyActionIDFirstAllOfRegistration`

NewSignOnPolicyActionIDFirstAllOfRegistrationWithDefaults instantiates a new SignOnPolicyActionIDFirstAllOfRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExternal

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetExternal() SignOnPolicyActionLoginAllOfRegistrationExternal`

GetExternal returns the External field if non-nil, zero value otherwise.

### GetExternalOk

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetExternalOk() (*SignOnPolicyActionLoginAllOfRegistrationExternal, bool)`

GetExternalOk returns a tuple with the External field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternal

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) SetExternal(v SignOnPolicyActionLoginAllOfRegistrationExternal)`

SetExternal sets External field to given value.

### HasExternal

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) HasExternal() bool`

HasExternal returns a boolean if a field has been set.

### GetPopulation

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetPopulation() SignOnPolicyActionLoginAllOfRegistrationPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetPopulationOk() (*SignOnPolicyActionLoginAllOfRegistrationPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) SetPopulation(v SignOnPolicyActionLoginAllOfRegistrationPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.

### GetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetConfirmIdentityProviderAttributes() bool`

GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field if non-nil, zero value otherwise.

### GetConfirmIdentityProviderAttributesOk

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) GetConfirmIdentityProviderAttributesOk() (*bool, bool)`

GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) SetConfirmIdentityProviderAttributes(v bool)`

SetConfirmIdentityProviderAttributes sets ConfirmIdentityProviderAttributes field to given value.

### HasConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOfRegistration) HasConfirmIdentityProviderAttributes() bool`

HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


