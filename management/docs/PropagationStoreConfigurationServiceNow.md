# PropagationStoreConfigurationServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorPassword** | **string** | Password for the administrator. | 
**AdministratorUsername** | **string** | Username for the administrator. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **string** | Whether or not users are allowed to be deprovisioned (removed) following action specified in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not new users are allowed to be disabled. | [optional] 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisable**](EnumPropagationStoreTypeRemoveActionDisable.md) |  | [optional] 
**ServiceNowUrl** | **string** | The URL for the ServiceNow account. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationServiceNow

`func NewPropagationStoreConfigurationServiceNow(administratorPassword string, administratorUsername string, serviceNowUrl string, ) *PropagationStoreConfigurationServiceNow`

NewPropagationStoreConfigurationServiceNow instantiates a new PropagationStoreConfigurationServiceNow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationServiceNowWithDefaults

`func NewPropagationStoreConfigurationServiceNowWithDefaults() *PropagationStoreConfigurationServiceNow`

NewPropagationStoreConfigurationServiceNowWithDefaults instantiates a new PropagationStoreConfigurationServiceNow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdministratorPassword

`func (o *PropagationStoreConfigurationServiceNow) GetAdministratorPassword() string`

GetAdministratorPassword returns the AdministratorPassword field if non-nil, zero value otherwise.

### GetAdministratorPasswordOk

`func (o *PropagationStoreConfigurationServiceNow) GetAdministratorPasswordOk() (*string, bool)`

GetAdministratorPasswordOk returns a tuple with the AdministratorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPassword

`func (o *PropagationStoreConfigurationServiceNow) SetAdministratorPassword(v string)`

SetAdministratorPassword sets AdministratorPassword field to given value.


### GetAdministratorUsername

`func (o *PropagationStoreConfigurationServiceNow) GetAdministratorUsername() string`

GetAdministratorUsername returns the AdministratorUsername field if non-nil, zero value otherwise.

### GetAdministratorUsernameOk

`func (o *PropagationStoreConfigurationServiceNow) GetAdministratorUsernameOk() (*string, bool)`

GetAdministratorUsernameOk returns a tuple with the AdministratorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorUsername

`func (o *PropagationStoreConfigurationServiceNow) SetAdministratorUsername(v string)`

SetAdministratorUsername sets AdministratorUsername field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationServiceNow) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationServiceNow) GetDEPROVISION_USERS() string`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationServiceNow) GetDEPROVISION_USERSOk() (*string, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationServiceNow) SetDEPROVISION_USERS(v string)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationServiceNow) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationServiceNow) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationServiceNow) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationServiceNow) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationServiceNow) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationServiceNow) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisable`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationServiceNow) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisable, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationServiceNow) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisable)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationServiceNow) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetServiceNowUrl

`func (o *PropagationStoreConfigurationServiceNow) GetServiceNowUrl() string`

GetServiceNowUrl returns the ServiceNowUrl field if non-nil, zero value otherwise.

### GetServiceNowUrlOk

`func (o *PropagationStoreConfigurationServiceNow) GetServiceNowUrlOk() (*string, bool)`

GetServiceNowUrlOk returns a tuple with the ServiceNowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowUrl

`func (o *PropagationStoreConfigurationServiceNow) SetServiceNowUrl(v string)`

SetServiceNowUrl sets ServiceNowUrl field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationServiceNow) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationServiceNow) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


