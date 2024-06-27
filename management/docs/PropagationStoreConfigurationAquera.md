# PropagationStoreConfigurationAquera

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCESS_TOKEN** | Pointer to **string** | A string specifying the access token for account authentication. | [optional] 
**AUTHENTICATION_METHOD** | [**EnumPropagationStoreTypeAqueraAuthenticationMethod**](EnumPropagationStoreTypeAqueraAuthenticationMethod.md) |  | 
**BASIC_AUTH_PASSWORD** | **string** | The password for account authentication. | 
**BASIC_AUTH_USER** | **string** | The user name for account authentication. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**GROUP_NAME_SOURCE** | Pointer to [**EnumPropagationStoreTypeAqueraGroupSourceName**](EnumPropagationStoreTypeAqueraGroupSourceName.md) |  | [optional] 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**SCIM_URL** | **string** | The SCIM URL. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationAquera

`func NewPropagationStoreConfigurationAquera(aUTHENTICATIONMETHOD EnumPropagationStoreTypeAqueraAuthenticationMethod, bASICAUTHPASSWORD string, bASICAUTHUSER string, sCIMURL string, ) *PropagationStoreConfigurationAquera`

NewPropagationStoreConfigurationAquera instantiates a new PropagationStoreConfigurationAquera object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationAqueraWithDefaults

`func NewPropagationStoreConfigurationAqueraWithDefaults() *PropagationStoreConfigurationAquera`

NewPropagationStoreConfigurationAqueraWithDefaults instantiates a new PropagationStoreConfigurationAquera object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCESS_TOKEN

`func (o *PropagationStoreConfigurationAquera) GetACCESS_TOKEN() string`

GetACCESS_TOKEN returns the ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetACCESS_TOKENOk

`func (o *PropagationStoreConfigurationAquera) GetACCESS_TOKENOk() (*string, bool)`

GetACCESS_TOKENOk returns a tuple with the ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_TOKEN

`func (o *PropagationStoreConfigurationAquera) SetACCESS_TOKEN(v string)`

SetACCESS_TOKEN sets ACCESS_TOKEN field to given value.

### HasACCESS_TOKEN

`func (o *PropagationStoreConfigurationAquera) HasACCESS_TOKEN() bool`

HasACCESS_TOKEN returns a boolean if a field has been set.

### GetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationAquera) GetAUTHENTICATION_METHOD() EnumPropagationStoreTypeAqueraAuthenticationMethod`

GetAUTHENTICATION_METHOD returns the AUTHENTICATION_METHOD field if non-nil, zero value otherwise.

### GetAUTHENTICATION_METHODOk

`func (o *PropagationStoreConfigurationAquera) GetAUTHENTICATION_METHODOk() (*EnumPropagationStoreTypeAqueraAuthenticationMethod, bool)`

GetAUTHENTICATION_METHODOk returns a tuple with the AUTHENTICATION_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationAquera) SetAUTHENTICATION_METHOD(v EnumPropagationStoreTypeAqueraAuthenticationMethod)`

SetAUTHENTICATION_METHOD sets AUTHENTICATION_METHOD field to given value.


### GetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfigurationAquera) GetBASIC_AUTH_PASSWORD() string`

GetBASIC_AUTH_PASSWORD returns the BASIC_AUTH_PASSWORD field if non-nil, zero value otherwise.

### GetBASIC_AUTH_PASSWORDOk

`func (o *PropagationStoreConfigurationAquera) GetBASIC_AUTH_PASSWORDOk() (*string, bool)`

GetBASIC_AUTH_PASSWORDOk returns a tuple with the BASIC_AUTH_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfigurationAquera) SetBASIC_AUTH_PASSWORD(v string)`

SetBASIC_AUTH_PASSWORD sets BASIC_AUTH_PASSWORD field to given value.


### GetBASIC_AUTH_USER

`func (o *PropagationStoreConfigurationAquera) GetBASIC_AUTH_USER() string`

GetBASIC_AUTH_USER returns the BASIC_AUTH_USER field if non-nil, zero value otherwise.

### GetBASIC_AUTH_USEROk

`func (o *PropagationStoreConfigurationAquera) GetBASIC_AUTH_USEROk() (*string, bool)`

GetBASIC_AUTH_USEROk returns a tuple with the BASIC_AUTH_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_USER

`func (o *PropagationStoreConfigurationAquera) SetBASIC_AUTH_USER(v string)`

SetBASIC_AUTH_USER sets BASIC_AUTH_USER field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationAquera) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationAquera) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationAquera) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationAquera) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationAquera) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationAquera) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationAquera) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationAquera) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationAquera) GetGROUP_NAME_SOURCE() EnumPropagationStoreTypeAqueraGroupSourceName`

GetGROUP_NAME_SOURCE returns the GROUP_NAME_SOURCE field if non-nil, zero value otherwise.

### GetGROUP_NAME_SOURCEOk

`func (o *PropagationStoreConfigurationAquera) GetGROUP_NAME_SOURCEOk() (*EnumPropagationStoreTypeAqueraGroupSourceName, bool)`

GetGROUP_NAME_SOURCEOk returns a tuple with the GROUP_NAME_SOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationAquera) SetGROUP_NAME_SOURCE(v EnumPropagationStoreTypeAqueraGroupSourceName)`

SetGROUP_NAME_SOURCE sets GROUP_NAME_SOURCE field to given value.

### HasGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationAquera) HasGROUP_NAME_SOURCE() bool`

HasGROUP_NAME_SOURCE returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationAquera) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationAquera) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationAquera) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationAquera) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetSCIM_URL

`func (o *PropagationStoreConfigurationAquera) GetSCIM_URL() string`

GetSCIM_URL returns the SCIM_URL field if non-nil, zero value otherwise.

### GetSCIM_URLOk

`func (o *PropagationStoreConfigurationAquera) GetSCIM_URLOk() (*string, bool)`

GetSCIM_URLOk returns a tuple with the SCIM_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_URL

`func (o *PropagationStoreConfigurationAquera) SetSCIM_URL(v string)`

SetSCIM_URL sets SCIM_URL field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationAquera) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationAquera) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationAquera) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationAquera) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


