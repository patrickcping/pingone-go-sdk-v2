# FIDOPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | FIDO policy&#39;s UUID. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]interface{}** |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | The name to use for the FIDO policy. | 
**Description** | Pointer to **string** | Description of the FIDO policy. | [optional] 
**AttestationRequirements** | [**EnumFIDOAttestationRequirements**](EnumFIDOAttestationRequirements.md) |  | 
**AllowedAuthenticators** | Pointer to [**[]FIDOPolicyAllowedAuthenticatorsInner**](FIDOPolicyAllowedAuthenticatorsInner.md) | If &#x60;attestationRequirements&#x60; is set to &#x60;SPECIFIC&#x60;, this array is used to specify the authenticators that you want to allow. | [optional] 
**EnforceDuringAuthentication** | Pointer to **bool** | This parameter is relevant only if you have set &#x60;attestationRequirements&#x60; to &#x60;SPECIFIC&#x60; in order to restrict usage to only certain authenticators. If set to &#x60;true&#x60;, the policy will be applied both during registration and during each authentication attempt. If set to &#x60;false&#x60;, the policy is applied only during registration. Default is &#x60;false&#x60;. | [optional] [default to false]
**Default** | Pointer to **bool** | Whether this policy should serve as the default FIDO policy. | [optional] 
**ResidentKeyRequirement** | [**EnumFIDOResidentKeyRequirement**](EnumFIDOResidentKeyRequirement.md) |  | 

## Methods

### NewFIDOPolicy

`func NewFIDOPolicy(name string, attestationRequirements EnumFIDOAttestationRequirements, residentKeyRequirement EnumFIDOResidentKeyRequirement, ) *FIDOPolicy`

NewFIDOPolicy instantiates a new FIDOPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDOPolicyWithDefaults

`func NewFIDOPolicyWithDefaults() *FIDOPolicy`

NewFIDOPolicyWithDefaults instantiates a new FIDOPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FIDOPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FIDOPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FIDOPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FIDOPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FIDOPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FIDOPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FIDOPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FIDOPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FIDOPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FIDOPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FIDOPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FIDOPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEmbedded

`func (o *FIDOPolicy) GetEmbedded() map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *FIDOPolicy) GetEmbeddedOk() (*map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *FIDOPolicy) SetEmbedded(v map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *FIDOPolicy) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironment

`func (o *FIDOPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *FIDOPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *FIDOPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *FIDOPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *FIDOPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FIDOPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FIDOPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FIDOPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FIDOPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FIDOPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FIDOPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAttestationRequirements

`func (o *FIDOPolicy) GetAttestationRequirements() EnumFIDOAttestationRequirements`

GetAttestationRequirements returns the AttestationRequirements field if non-nil, zero value otherwise.

### GetAttestationRequirementsOk

`func (o *FIDOPolicy) GetAttestationRequirementsOk() (*EnumFIDOAttestationRequirements, bool)`

GetAttestationRequirementsOk returns a tuple with the AttestationRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttestationRequirements

`func (o *FIDOPolicy) SetAttestationRequirements(v EnumFIDOAttestationRequirements)`

SetAttestationRequirements sets AttestationRequirements field to given value.


### GetAllowedAuthenticators

`func (o *FIDOPolicy) GetAllowedAuthenticators() []FIDOPolicyAllowedAuthenticatorsInner`

GetAllowedAuthenticators returns the AllowedAuthenticators field if non-nil, zero value otherwise.

### GetAllowedAuthenticatorsOk

`func (o *FIDOPolicy) GetAllowedAuthenticatorsOk() (*[]FIDOPolicyAllowedAuthenticatorsInner, bool)`

GetAllowedAuthenticatorsOk returns a tuple with the AllowedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticators

`func (o *FIDOPolicy) SetAllowedAuthenticators(v []FIDOPolicyAllowedAuthenticatorsInner)`

SetAllowedAuthenticators sets AllowedAuthenticators field to given value.

### HasAllowedAuthenticators

`func (o *FIDOPolicy) HasAllowedAuthenticators() bool`

HasAllowedAuthenticators returns a boolean if a field has been set.

### GetEnforceDuringAuthentication

`func (o *FIDOPolicy) GetEnforceDuringAuthentication() bool`

GetEnforceDuringAuthentication returns the EnforceDuringAuthentication field if non-nil, zero value otherwise.

### GetEnforceDuringAuthenticationOk

`func (o *FIDOPolicy) GetEnforceDuringAuthenticationOk() (*bool, bool)`

GetEnforceDuringAuthenticationOk returns a tuple with the EnforceDuringAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceDuringAuthentication

`func (o *FIDOPolicy) SetEnforceDuringAuthentication(v bool)`

SetEnforceDuringAuthentication sets EnforceDuringAuthentication field to given value.

### HasEnforceDuringAuthentication

`func (o *FIDOPolicy) HasEnforceDuringAuthentication() bool`

HasEnforceDuringAuthentication returns a boolean if a field has been set.

### GetDefault

`func (o *FIDOPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *FIDOPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *FIDOPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *FIDOPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetResidentKeyRequirement

`func (o *FIDOPolicy) GetResidentKeyRequirement() EnumFIDOResidentKeyRequirement`

GetResidentKeyRequirement returns the ResidentKeyRequirement field if non-nil, zero value otherwise.

### GetResidentKeyRequirementOk

`func (o *FIDOPolicy) GetResidentKeyRequirementOk() (*EnumFIDOResidentKeyRequirement, bool)`

GetResidentKeyRequirementOk returns a tuple with the ResidentKeyRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidentKeyRequirement

`func (o *FIDOPolicy) SetResidentKeyRequirement(v EnumFIDOResidentKeyRequirement)`

SetResidentKeyRequirement sets ResidentKeyRequirement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


