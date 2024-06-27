# PropagationStoreConfigurationGoogleApps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APPLICATION_NAME** | **string** | Name of the application using the store. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be deprovisioned (removed) following action specified in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**DOMAIN** | **string** | The account&#39;s domain name. | 
**OAUTH_ACCESS_TOKEN** | **string** | OAuth 2 access token. | 
**OAUTH_CLIENT_ID** | **string** | GoogleApps identifier of the client associated with the propagation store. | 
**OAUTH_CLIENT_SECRET** | **string** | GoogleApps secret of the client associated with the propagation store. | 
**OAUTH_REFRESH_TOKEN** | **string** | OAuth 2 refresh token. | 
**PROVISION_DISABLED_USERS_PROV_OPT** | Pointer to **bool** | Whether or not disabled users can be provisioned. Defaults to &#x60;true&#x60; and, if used, must be set to &#x60;true&#x60;. | [optional] [default to true]
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationGoogleApps

`func NewPropagationStoreConfigurationGoogleApps(aPPLICATIONNAME string, dOMAIN string, oAUTHACCESSTOKEN string, oAUTHCLIENTID string, oAUTHCLIENTSECRET string, oAUTHREFRESHTOKEN string, ) *PropagationStoreConfigurationGoogleApps`

NewPropagationStoreConfigurationGoogleApps instantiates a new PropagationStoreConfigurationGoogleApps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationGoogleAppsWithDefaults

`func NewPropagationStoreConfigurationGoogleAppsWithDefaults() *PropagationStoreConfigurationGoogleApps`

NewPropagationStoreConfigurationGoogleAppsWithDefaults instantiates a new PropagationStoreConfigurationGoogleApps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPPLICATION_NAME

`func (o *PropagationStoreConfigurationGoogleApps) GetAPPLICATION_NAME() string`

GetAPPLICATION_NAME returns the APPLICATION_NAME field if non-nil, zero value otherwise.

### GetAPPLICATION_NAMEOk

`func (o *PropagationStoreConfigurationGoogleApps) GetAPPLICATION_NAMEOk() (*string, bool)`

GetAPPLICATION_NAMEOk returns a tuple with the APPLICATION_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPLICATION_NAME

`func (o *PropagationStoreConfigurationGoogleApps) SetAPPLICATION_NAME(v string)`

SetAPPLICATION_NAME sets APPLICATION_NAME field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationGoogleApps) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGoogleApps) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationGoogleApps) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGoogleApps) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationGoogleApps) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationGoogleApps) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetDOMAIN

`func (o *PropagationStoreConfigurationGoogleApps) GetDOMAIN() string`

GetDOMAIN returns the DOMAIN field if non-nil, zero value otherwise.

### GetDOMAINOk

`func (o *PropagationStoreConfigurationGoogleApps) GetDOMAINOk() (*string, bool)`

GetDOMAINOk returns a tuple with the DOMAIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOMAIN

`func (o *PropagationStoreConfigurationGoogleApps) SetDOMAIN(v string)`

SetDOMAIN sets DOMAIN field to given value.


### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationGoogleApps) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.


### GetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_CLIENT_ID() string`

GetOAUTH_CLIENT_ID returns the OAUTH_CLIENT_ID field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_IDOk

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_CLIENT_IDOk() (*string, bool)`

GetOAUTH_CLIENT_IDOk returns a tuple with the OAUTH_CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationGoogleApps) SetOAUTH_CLIENT_ID(v string)`

SetOAUTH_CLIENT_ID sets OAUTH_CLIENT_ID field to given value.


### GetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_CLIENT_SECRET() string`

GetOAUTH_CLIENT_SECRET returns the OAUTH_CLIENT_SECRET field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_SECRETOk

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_CLIENT_SECRETOk() (*string, bool)`

GetOAUTH_CLIENT_SECRETOk returns a tuple with the OAUTH_CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationGoogleApps) SetOAUTH_CLIENT_SECRET(v string)`

SetOAUTH_CLIENT_SECRET sets OAUTH_CLIENT_SECRET field to given value.


### GetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_REFRESH_TOKEN() string`

GetOAUTH_REFRESH_TOKEN returns the OAUTH_REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_REFRESH_TOKENOk

`func (o *PropagationStoreConfigurationGoogleApps) GetOAUTH_REFRESH_TOKENOk() (*string, bool)`

GetOAUTH_REFRESH_TOKENOk returns a tuple with the OAUTH_REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfigurationGoogleApps) SetOAUTH_REFRESH_TOKEN(v string)`

SetOAUTH_REFRESH_TOKEN sets OAUTH_REFRESH_TOKEN field to given value.


### GetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationGoogleApps) GetPROVISION_DISABLED_USERS_PROV_OPT() bool`

GetPROVISION_DISABLED_USERS_PROV_OPT returns the PROVISION_DISABLED_USERS_PROV_OPT field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERS_PROV_OPTOk

`func (o *PropagationStoreConfigurationGoogleApps) GetPROVISION_DISABLED_USERS_PROV_OPTOk() (*bool, bool)`

GetPROVISION_DISABLED_USERS_PROV_OPTOk returns a tuple with the PROVISION_DISABLED_USERS_PROV_OPT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationGoogleApps) SetPROVISION_DISABLED_USERS_PROV_OPT(v bool)`

SetPROVISION_DISABLED_USERS_PROV_OPT sets PROVISION_DISABLED_USERS_PROV_OPT field to given value.

### HasPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationGoogleApps) HasPROVISION_DISABLED_USERS_PROV_OPT() bool`

HasPROVISION_DISABLED_USERS_PROV_OPT returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationGoogleApps) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationGoogleApps) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationGoogleApps) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationGoogleApps) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationGoogleApps) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationGoogleApps) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


