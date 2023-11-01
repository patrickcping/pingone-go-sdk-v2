# PropagationStoreConfigurationAzureADSAMLV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Azure Active Directory client ID. | 
**ClientSecret** | **string** | The Azure Active Directory client secret. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be deprovisioned (removed) following action specified in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**PROVISION_DISABLED_USERS_PROV_OPT** | Pointer to **bool** | Whether or not disabled users can be provisioned. Defaults to &#x60;true&#x60; and, if used, must be set to &#x60;true&#x60;. | [optional] [default to true]
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**RemoveLicensesWhenSkuIdEmpty** | **bool** | Whether or not remove licenses from user when &#x60;skuId&#x60; is empty. | 
**TenantDomain** | **string** | The account&#39;s Azure Active Directory domain. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationAzureADSAMLV2

`func NewPropagationStoreConfigurationAzureADSAMLV2(clientId string, clientSecret string, removeLicensesWhenSkuIdEmpty bool, tenantDomain string, ) *PropagationStoreConfigurationAzureADSAMLV2`

NewPropagationStoreConfigurationAzureADSAMLV2 instantiates a new PropagationStoreConfigurationAzureADSAMLV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationAzureADSAMLV2WithDefaults

`func NewPropagationStoreConfigurationAzureADSAMLV2WithDefaults() *PropagationStoreConfigurationAzureADSAMLV2`

NewPropagationStoreConfigurationAzureADSAMLV2WithDefaults instantiates a new PropagationStoreConfigurationAzureADSAMLV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetPROVISION_DISABLED_USERS_PROV_OPT() bool`

GetPROVISION_DISABLED_USERS_PROV_OPT returns the PROVISION_DISABLED_USERS_PROV_OPT field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERS_PROV_OPTOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetPROVISION_DISABLED_USERS_PROV_OPTOk() (*bool, bool)`

GetPROVISION_DISABLED_USERS_PROV_OPTOk returns a tuple with the PROVISION_DISABLED_USERS_PROV_OPT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetPROVISION_DISABLED_USERS_PROV_OPT(v bool)`

SetPROVISION_DISABLED_USERS_PROV_OPT sets PROVISION_DISABLED_USERS_PROV_OPT field to given value.

### HasPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasPROVISION_DISABLED_USERS_PROV_OPT() bool`

HasPROVISION_DISABLED_USERS_PROV_OPT returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetRemoveLicensesWhenSkuIdEmpty

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetRemoveLicensesWhenSkuIdEmpty() bool`

GetRemoveLicensesWhenSkuIdEmpty returns the RemoveLicensesWhenSkuIdEmpty field if non-nil, zero value otherwise.

### GetRemoveLicensesWhenSkuIdEmptyOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetRemoveLicensesWhenSkuIdEmptyOk() (*bool, bool)`

GetRemoveLicensesWhenSkuIdEmptyOk returns a tuple with the RemoveLicensesWhenSkuIdEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveLicensesWhenSkuIdEmpty

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetRemoveLicensesWhenSkuIdEmpty(v bool)`

SetRemoveLicensesWhenSkuIdEmpty sets RemoveLicensesWhenSkuIdEmpty field to given value.


### GetTenantDomain

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetTenantDomain() string`

GetTenantDomain returns the TenantDomain field if non-nil, zero value otherwise.

### GetTenantDomainOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetTenantDomainOk() (*string, bool)`

GetTenantDomainOk returns a tuple with the TenantDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantDomain

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetTenantDomain(v string)`

SetTenantDomain sets TenantDomain field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationAzureADSAMLV2) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationAzureADSAMLV2) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


