# SignOnPolicyActionLoginAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceLockoutForIdentityProviders** | Pointer to **bool** | A boolean that if set to true and if the user&#39;s account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented. | [optional] 
**NewUserProvisioning** | Pointer to [**SignOnPolicyActionLoginAllOfNewUserProvisioning**](SignOnPolicyActionLoginAllOfNewUserProvisioning.md) |  | [optional] 
**Recovery** | Pointer to [**SignOnPolicyActionLoginAllOfRecovery**](SignOnPolicyActionLoginAllOfRecovery.md) |  | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionLoginAllOfRegistration**](SignOnPolicyActionLoginAllOfRegistration.md) |  | [optional] 
**SocialProviders** | Pointer to [**[]SignOnPolicyActionLoginAllOfSocialProviders**](SignOnPolicyActionLoginAllOfSocialProviders.md) | An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow. | [optional] 

## Methods

### NewSignOnPolicyActionLoginAllOf

`func NewSignOnPolicyActionLoginAllOf() *SignOnPolicyActionLoginAllOf`

NewSignOnPolicyActionLoginAllOf instantiates a new SignOnPolicyActionLoginAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionLoginAllOfWithDefaults

`func NewSignOnPolicyActionLoginAllOfWithDefaults() *SignOnPolicyActionLoginAllOf`

NewSignOnPolicyActionLoginAllOfWithDefaults instantiates a new SignOnPolicyActionLoginAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionLoginAllOf) GetEnforceLockoutForIdentityProviders() bool`

GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field if non-nil, zero value otherwise.

### GetEnforceLockoutForIdentityProvidersOk

`func (o *SignOnPolicyActionLoginAllOf) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool)`

GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionLoginAllOf) SetEnforceLockoutForIdentityProviders(v bool)`

SetEnforceLockoutForIdentityProviders sets EnforceLockoutForIdentityProviders field to given value.

### HasEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionLoginAllOf) HasEnforceLockoutForIdentityProviders() bool`

HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.

### GetNewUserProvisioning

`func (o *SignOnPolicyActionLoginAllOf) GetNewUserProvisioning() SignOnPolicyActionLoginAllOfNewUserProvisioning`

GetNewUserProvisioning returns the NewUserProvisioning field if non-nil, zero value otherwise.

### GetNewUserProvisioningOk

`func (o *SignOnPolicyActionLoginAllOf) GetNewUserProvisioningOk() (*SignOnPolicyActionLoginAllOfNewUserProvisioning, bool)`

GetNewUserProvisioningOk returns a tuple with the NewUserProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserProvisioning

`func (o *SignOnPolicyActionLoginAllOf) SetNewUserProvisioning(v SignOnPolicyActionLoginAllOfNewUserProvisioning)`

SetNewUserProvisioning sets NewUserProvisioning field to given value.

### HasNewUserProvisioning

`func (o *SignOnPolicyActionLoginAllOf) HasNewUserProvisioning() bool`

HasNewUserProvisioning returns a boolean if a field has been set.

### GetRecovery

`func (o *SignOnPolicyActionLoginAllOf) GetRecovery() SignOnPolicyActionLoginAllOfRecovery`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *SignOnPolicyActionLoginAllOf) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *SignOnPolicyActionLoginAllOf) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *SignOnPolicyActionLoginAllOf) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionLoginAllOf) GetRegistration() SignOnPolicyActionLoginAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionLoginAllOf) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionLoginAllOf) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionLoginAllOf) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSocialProviders

`func (o *SignOnPolicyActionLoginAllOf) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *SignOnPolicyActionLoginAllOf) GetSocialProvidersOk() (*[]SignOnPolicyActionLoginAllOfSocialProviders, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *SignOnPolicyActionLoginAllOf) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *SignOnPolicyActionLoginAllOf) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


