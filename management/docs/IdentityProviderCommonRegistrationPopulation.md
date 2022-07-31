# IdentityProviderCommonRegistrationPopulation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An external IdP to use as authoritative. Setting this attribute gives management of linked users to the IdP and also triggers just-in-time provisioning of new users. These users are created in the population indicated with registration.population.id. | [optional] 

## Methods

### NewIdentityProviderCommonRegistrationPopulation

`func NewIdentityProviderCommonRegistrationPopulation() *IdentityProviderCommonRegistrationPopulation`

NewIdentityProviderCommonRegistrationPopulation instantiates a new IdentityProviderCommonRegistrationPopulation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderCommonRegistrationPopulationWithDefaults

`func NewIdentityProviderCommonRegistrationPopulationWithDefaults() *IdentityProviderCommonRegistrationPopulation`

NewIdentityProviderCommonRegistrationPopulationWithDefaults instantiates a new IdentityProviderCommonRegistrationPopulation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentityProviderCommonRegistrationPopulation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderCommonRegistrationPopulation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderCommonRegistrationPopulation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderCommonRegistrationPopulation) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


