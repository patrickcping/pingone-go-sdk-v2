# PropagationStoreConfigurationPingOne

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CLIENT_ID** | Pointer to **string** | Unique identifier (UUID) of the PingOne client associated with the propagation store. | [optional] 
**CLIENT_SECRET** | Pointer to **string** | The PingOne client secret. | [optional] 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEFAULT_AUTH_METHOD** | Pointer to [**EnumPropagationStoreTypePingOneDefaultAuthMethod**](EnumPropagationStoreTypePingOneDefaultAuthMethod.md) |  | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**ENVIRONMENT_ID** | **string** | Unique identifier (UUID) of the PingOne environment associated with the propagation store. | 
**MFA_USER_DEVICE_MANAGEMENT** | Pointer to **string** | How to manage MFA user devices when synchronizing. Options are either Merge with devices in PingOne or Overwrite devices in PingOne. | [optional] 
**PROVISION_DISABLED_USERS_PROV_OPT** | Pointer to **bool** | Whether or not disabled users can be provisioned. | [optional] 
**REGION** | [**EnumPropagationStoreTypePingOneRegion**](EnumPropagationStoreTypePingOneRegion.md) |  | 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationPingOne

`func NewPropagationStoreConfigurationPingOne(eNVIRONMENTID string, rEGION EnumPropagationStoreTypePingOneRegion, ) *PropagationStoreConfigurationPingOne`

NewPropagationStoreConfigurationPingOne instantiates a new PropagationStoreConfigurationPingOne object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationPingOneWithDefaults

`func NewPropagationStoreConfigurationPingOneWithDefaults() *PropagationStoreConfigurationPingOne`

NewPropagationStoreConfigurationPingOneWithDefaults instantiates a new PropagationStoreConfigurationPingOne object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCLIENT_ID

`func (o *PropagationStoreConfigurationPingOne) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *PropagationStoreConfigurationPingOne) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *PropagationStoreConfigurationPingOne) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.

### HasCLIENT_ID

`func (o *PropagationStoreConfigurationPingOne) HasCLIENT_ID() bool`

HasCLIENT_ID returns a boolean if a field has been set.

### GetCLIENT_SECRET

`func (o *PropagationStoreConfigurationPingOne) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *PropagationStoreConfigurationPingOne) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *PropagationStoreConfigurationPingOne) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.

### HasCLIENT_SECRET

`func (o *PropagationStoreConfigurationPingOne) HasCLIENT_SECRET() bool`

HasCLIENT_SECRET returns a boolean if a field has been set.

### GetCREATE_USERS

`func (o *PropagationStoreConfigurationPingOne) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationPingOne) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationPingOne) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationPingOne) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfigurationPingOne) GetDEFAULT_AUTH_METHOD() EnumPropagationStoreTypePingOneDefaultAuthMethod`

GetDEFAULT_AUTH_METHOD returns the DEFAULT_AUTH_METHOD field if non-nil, zero value otherwise.

### GetDEFAULT_AUTH_METHODOk

`func (o *PropagationStoreConfigurationPingOne) GetDEFAULT_AUTH_METHODOk() (*EnumPropagationStoreTypePingOneDefaultAuthMethod, bool)`

GetDEFAULT_AUTH_METHODOk returns a tuple with the DEFAULT_AUTH_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfigurationPingOne) SetDEFAULT_AUTH_METHOD(v EnumPropagationStoreTypePingOneDefaultAuthMethod)`

SetDEFAULT_AUTH_METHOD sets DEFAULT_AUTH_METHOD field to given value.

### HasDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfigurationPingOne) HasDEFAULT_AUTH_METHOD() bool`

HasDEFAULT_AUTH_METHOD returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationPingOne) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationPingOne) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationPingOne) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationPingOne) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetENVIRONMENT_ID

`func (o *PropagationStoreConfigurationPingOne) GetENVIRONMENT_ID() string`

GetENVIRONMENT_ID returns the ENVIRONMENT_ID field if non-nil, zero value otherwise.

### GetENVIRONMENT_IDOk

`func (o *PropagationStoreConfigurationPingOne) GetENVIRONMENT_IDOk() (*string, bool)`

GetENVIRONMENT_IDOk returns a tuple with the ENVIRONMENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENVIRONMENT_ID

`func (o *PropagationStoreConfigurationPingOne) SetENVIRONMENT_ID(v string)`

SetENVIRONMENT_ID sets ENVIRONMENT_ID field to given value.


### GetMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfigurationPingOne) GetMFA_USER_DEVICE_MANAGEMENT() string`

GetMFA_USER_DEVICE_MANAGEMENT returns the MFA_USER_DEVICE_MANAGEMENT field if non-nil, zero value otherwise.

### GetMFA_USER_DEVICE_MANAGEMENTOk

`func (o *PropagationStoreConfigurationPingOne) GetMFA_USER_DEVICE_MANAGEMENTOk() (*string, bool)`

GetMFA_USER_DEVICE_MANAGEMENTOk returns a tuple with the MFA_USER_DEVICE_MANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfigurationPingOne) SetMFA_USER_DEVICE_MANAGEMENT(v string)`

SetMFA_USER_DEVICE_MANAGEMENT sets MFA_USER_DEVICE_MANAGEMENT field to given value.

### HasMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfigurationPingOne) HasMFA_USER_DEVICE_MANAGEMENT() bool`

HasMFA_USER_DEVICE_MANAGEMENT returns a boolean if a field has been set.

### GetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationPingOne) GetPROVISION_DISABLED_USERS_PROV_OPT() bool`

GetPROVISION_DISABLED_USERS_PROV_OPT returns the PROVISION_DISABLED_USERS_PROV_OPT field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERS_PROV_OPTOk

`func (o *PropagationStoreConfigurationPingOne) GetPROVISION_DISABLED_USERS_PROV_OPTOk() (*bool, bool)`

GetPROVISION_DISABLED_USERS_PROV_OPTOk returns a tuple with the PROVISION_DISABLED_USERS_PROV_OPT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationPingOne) SetPROVISION_DISABLED_USERS_PROV_OPT(v bool)`

SetPROVISION_DISABLED_USERS_PROV_OPT sets PROVISION_DISABLED_USERS_PROV_OPT field to given value.

### HasPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfigurationPingOne) HasPROVISION_DISABLED_USERS_PROV_OPT() bool`

HasPROVISION_DISABLED_USERS_PROV_OPT returns a boolean if a field has been set.

### GetREGION

`func (o *PropagationStoreConfigurationPingOne) GetREGION() EnumPropagationStoreTypePingOneRegion`

GetREGION returns the REGION field if non-nil, zero value otherwise.

### GetREGIONOk

`func (o *PropagationStoreConfigurationPingOne) GetREGIONOk() (*EnumPropagationStoreTypePingOneRegion, bool)`

GetREGIONOk returns a tuple with the REGION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREGION

`func (o *PropagationStoreConfigurationPingOne) SetREGION(v EnumPropagationStoreTypePingOneRegion)`

SetREGION sets REGION field to given value.


### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationPingOne) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationPingOne) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationPingOne) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationPingOne) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationPingOne) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationPingOne) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationPingOne) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationPingOne) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


