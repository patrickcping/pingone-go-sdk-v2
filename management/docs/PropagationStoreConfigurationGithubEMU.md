# PropagationStoreConfigurationGithubEMU

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BASE_URL** | **string** | Base URL of the target propagation store. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be deprovisioned (removed) following action specified in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**OAUTH_ACCESS_TOKEN** | **string** | OAuth 2 access token. | 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationGithubEMU

`func NewPropagationStoreConfigurationGithubEMU(bASEURL string, oAUTHACCESSTOKEN string, ) *PropagationStoreConfigurationGithubEMU`

NewPropagationStoreConfigurationGithubEMU instantiates a new PropagationStoreConfigurationGithubEMU object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationGithubEMUWithDefaults

`func NewPropagationStoreConfigurationGithubEMUWithDefaults() *PropagationStoreConfigurationGithubEMU`

NewPropagationStoreConfigurationGithubEMUWithDefaults instantiates a new PropagationStoreConfigurationGithubEMU object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBASE_URL

`func (o *PropagationStoreConfigurationGithubEMU) GetBASE_URL() string`

GetBASE_URL returns the BASE_URL field if non-nil, zero value otherwise.

### GetBASE_URLOk

`func (o *PropagationStoreConfigurationGithubEMU) GetBASE_URLOk() (*string, bool)`

GetBASE_URLOk returns a tuple with the BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE_URL

`func (o *PropagationStoreConfigurationGithubEMU) SetBASE_URL(v string)`

SetBASE_URL sets BASE_URL field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationGithubEMU) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGithubEMU) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationGithubEMU) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGithubEMU) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGithubEMU) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationGithubEMU) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfigurationGithubEMU) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationGithubEMU) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.


### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationGithubEMU) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationGithubEMU) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationGithubEMU) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationGithubEMU) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationGithubEMU) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationGithubEMU) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


