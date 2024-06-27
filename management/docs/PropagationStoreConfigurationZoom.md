# PropagationStoreConfigurationZoom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**API_KEY** | Pointer to **string** | The client API key. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;JWT Bearer Token&#x60;, otherwise optional. | [optional] 
**API_SECRET** | Pointer to **string** | The client API secret. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;JWT Bearer Token&#x60;, otherwise optional. | [optional] 
**AUTHENTICATION_METHOD** | Pointer to [**EnumPropagationStoreTypeZoomAuthenticationMethod**](EnumPropagationStoreTypeZoomAuthenticationMethod.md) |  | [optional] [default to ENUMPROPAGATIONSTORETYPEZOOMAUTHENTICATIONMETHOD_JWT_BEARER_TOKEN]
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be removed (deprovisioned) following the action configured in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**OAUTH_ACCOUNT_ID** | Pointer to **string** | OAuth account identifier. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_CLIENT_ID** | Pointer to **string** | OAuth client identifier. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_CLIENT_SECRET** | Pointer to **string** | OAuth client secret. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_TOKEN_URL** | Pointer to **string** | OAuth token request endpoint. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**SCIM_URL** | **string** | The SCIM URL. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationZoom

`func NewPropagationStoreConfigurationZoom(sCIMURL string, ) *PropagationStoreConfigurationZoom`

NewPropagationStoreConfigurationZoom instantiates a new PropagationStoreConfigurationZoom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationZoomWithDefaults

`func NewPropagationStoreConfigurationZoomWithDefaults() *PropagationStoreConfigurationZoom`

NewPropagationStoreConfigurationZoomWithDefaults instantiates a new PropagationStoreConfigurationZoom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPI_KEY

`func (o *PropagationStoreConfigurationZoom) GetAPI_KEY() string`

GetAPI_KEY returns the API_KEY field if non-nil, zero value otherwise.

### GetAPI_KEYOk

`func (o *PropagationStoreConfigurationZoom) GetAPI_KEYOk() (*string, bool)`

GetAPI_KEYOk returns a tuple with the API_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_KEY

`func (o *PropagationStoreConfigurationZoom) SetAPI_KEY(v string)`

SetAPI_KEY sets API_KEY field to given value.

### HasAPI_KEY

`func (o *PropagationStoreConfigurationZoom) HasAPI_KEY() bool`

HasAPI_KEY returns a boolean if a field has been set.

### GetAPI_SECRET

`func (o *PropagationStoreConfigurationZoom) GetAPI_SECRET() string`

GetAPI_SECRET returns the API_SECRET field if non-nil, zero value otherwise.

### GetAPI_SECRETOk

`func (o *PropagationStoreConfigurationZoom) GetAPI_SECRETOk() (*string, bool)`

GetAPI_SECRETOk returns a tuple with the API_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_SECRET

`func (o *PropagationStoreConfigurationZoom) SetAPI_SECRET(v string)`

SetAPI_SECRET sets API_SECRET field to given value.

### HasAPI_SECRET

`func (o *PropagationStoreConfigurationZoom) HasAPI_SECRET() bool`

HasAPI_SECRET returns a boolean if a field has been set.

### GetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationZoom) GetAUTHENTICATION_METHOD() EnumPropagationStoreTypeZoomAuthenticationMethod`

GetAUTHENTICATION_METHOD returns the AUTHENTICATION_METHOD field if non-nil, zero value otherwise.

### GetAUTHENTICATION_METHODOk

`func (o *PropagationStoreConfigurationZoom) GetAUTHENTICATION_METHODOk() (*EnumPropagationStoreTypeZoomAuthenticationMethod, bool)`

GetAUTHENTICATION_METHODOk returns a tuple with the AUTHENTICATION_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationZoom) SetAUTHENTICATION_METHOD(v EnumPropagationStoreTypeZoomAuthenticationMethod)`

SetAUTHENTICATION_METHOD sets AUTHENTICATION_METHOD field to given value.

### HasAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationZoom) HasAUTHENTICATION_METHOD() bool`

HasAUTHENTICATION_METHOD returns a boolean if a field has been set.

### GetCREATE_USERS

`func (o *PropagationStoreConfigurationZoom) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationZoom) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationZoom) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationZoom) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationZoom) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfigurationZoom) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfigurationZoom) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfigurationZoom) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationZoom) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationZoom) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationZoom) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationZoom) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_ACCOUNT_ID() string`

GetOAUTH_ACCOUNT_ID returns the OAUTH_ACCOUNT_ID field if non-nil, zero value otherwise.

### GetOAUTH_ACCOUNT_IDOk

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_ACCOUNT_IDOk() (*string, bool)`

GetOAUTH_ACCOUNT_IDOk returns a tuple with the OAUTH_ACCOUNT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfigurationZoom) SetOAUTH_ACCOUNT_ID(v string)`

SetOAUTH_ACCOUNT_ID sets OAUTH_ACCOUNT_ID field to given value.

### HasOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfigurationZoom) HasOAUTH_ACCOUNT_ID() bool`

HasOAUTH_ACCOUNT_ID returns a boolean if a field has been set.

### GetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_CLIENT_ID() string`

GetOAUTH_CLIENT_ID returns the OAUTH_CLIENT_ID field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_IDOk

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_CLIENT_IDOk() (*string, bool)`

GetOAUTH_CLIENT_IDOk returns a tuple with the OAUTH_CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationZoom) SetOAUTH_CLIENT_ID(v string)`

SetOAUTH_CLIENT_ID sets OAUTH_CLIENT_ID field to given value.

### HasOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationZoom) HasOAUTH_CLIENT_ID() bool`

HasOAUTH_CLIENT_ID returns a boolean if a field has been set.

### GetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_CLIENT_SECRET() string`

GetOAUTH_CLIENT_SECRET returns the OAUTH_CLIENT_SECRET field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_SECRETOk

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_CLIENT_SECRETOk() (*string, bool)`

GetOAUTH_CLIENT_SECRETOk returns a tuple with the OAUTH_CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationZoom) SetOAUTH_CLIENT_SECRET(v string)`

SetOAUTH_CLIENT_SECRET sets OAUTH_CLIENT_SECRET field to given value.

### HasOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationZoom) HasOAUTH_CLIENT_SECRET() bool`

HasOAUTH_CLIENT_SECRET returns a boolean if a field has been set.

### GetOAUTH_TOKEN_URL

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_TOKEN_URL() string`

GetOAUTH_TOKEN_URL returns the OAUTH_TOKEN_URL field if non-nil, zero value otherwise.

### GetOAUTH_TOKEN_URLOk

`func (o *PropagationStoreConfigurationZoom) GetOAUTH_TOKEN_URLOk() (*string, bool)`

GetOAUTH_TOKEN_URLOk returns a tuple with the OAUTH_TOKEN_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_TOKEN_URL

`func (o *PropagationStoreConfigurationZoom) SetOAUTH_TOKEN_URL(v string)`

SetOAUTH_TOKEN_URL sets OAUTH_TOKEN_URL field to given value.

### HasOAUTH_TOKEN_URL

`func (o *PropagationStoreConfigurationZoom) HasOAUTH_TOKEN_URL() bool`

HasOAUTH_TOKEN_URL returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationZoom) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationZoom) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationZoom) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationZoom) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetSCIM_URL

`func (o *PropagationStoreConfigurationZoom) GetSCIM_URL() string`

GetSCIM_URL returns the SCIM_URL field if non-nil, zero value otherwise.

### GetSCIM_URLOk

`func (o *PropagationStoreConfigurationZoom) GetSCIM_URLOk() (*string, bool)`

GetSCIM_URLOk returns a tuple with the SCIM_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_URL

`func (o *PropagationStoreConfigurationZoom) SetSCIM_URL(v string)`

SetSCIM_URL sets SCIM_URL field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationZoom) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationZoom) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationZoom) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationZoom) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


