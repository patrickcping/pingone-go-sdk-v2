# PropagationStoreConfigurationSCIM

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AUTHENTICATION_METHOD** | [**EnumPropagationStoreTypeSCIMAuthenticationMethod**](EnumPropagationStoreTypeSCIMAuthenticationMethod.md) |  | 
**AUTHORIZATION_TYPE** | **string** | The authorization header type. | 
**BASIC_AUTH_PASSWORD** | Pointer to **string** | The password for account authentication. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;Basic Authentication&#x60;, otherwise optional. | [optional] 
**BASIC_AUTH_USER** | Pointer to **string** | The user name for account authentication. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;Basic Authentication&#x60;, otherwise optional. | [optional] 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**GROUP_NAME_SOURCE** | Pointer to [**EnumPropagationStoreTypeSCIMGroupNameSource**](EnumPropagationStoreTypeSCIMGroupNameSource.md) |  | [optional] 
**GROUPS_RESOURCE** | Pointer to **string** | API endpoint path to the group entity. | [optional] 
**OAUTH_ACCESS_TOKEN** | Pointer to **string** | OAuth access token for account authentication. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth 2 Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_CLIENT_ID** | Pointer to **string** | OAuth client ID. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth 2 Client Credentials&#x60;, otherwise optional. | [optional] 
**OAUTH_CLIENT_SECRET** | Pointer to **string** | OAuth client secret. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth 2 Client Credentials&#x60;, otherwise optional. | [optional] 
**OAUTH_TOKEN_REQUEST** | Pointer to **string** | OAuth token request endpoint. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth 2 Bearer Token&#x60;, otherwise optional. | [optional] 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**SCIM_URL** | **string** | The SCIM URL. | 
**SCIM_VERSION** | **string** | The SCIM version. | 
**UNIQUE_USER_IDENTIFIER** | [**EnumPropagationStoreTypeSCIMUniqueUserIdentifier**](EnumPropagationStoreTypeSCIMUniqueUserIdentifier.md) |  | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 
**USER_FILTER** | **string** | A string that specifies a SCIM filter expression. | 
**USERS_RESOURCE** | **string** | API endpoint path to the user entity. | 

## Methods

### NewPropagationStoreConfigurationSCIM

`func NewPropagationStoreConfigurationSCIM(aUTHENTICATIONMETHOD EnumPropagationStoreTypeSCIMAuthenticationMethod, aUTHORIZATIONTYPE string, sCIMURL string, sCIMVERSION string, uNIQUEUSERIDENTIFIER EnumPropagationStoreTypeSCIMUniqueUserIdentifier, uSERFILTER string, uSERSRESOURCE string, ) *PropagationStoreConfigurationSCIM`

NewPropagationStoreConfigurationSCIM instantiates a new PropagationStoreConfigurationSCIM object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationSCIMWithDefaults

`func NewPropagationStoreConfigurationSCIMWithDefaults() *PropagationStoreConfigurationSCIM`

NewPropagationStoreConfigurationSCIMWithDefaults instantiates a new PropagationStoreConfigurationSCIM object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationSCIM) GetAUTHENTICATION_METHOD() EnumPropagationStoreTypeSCIMAuthenticationMethod`

GetAUTHENTICATION_METHOD returns the AUTHENTICATION_METHOD field if non-nil, zero value otherwise.

### GetAUTHENTICATION_METHODOk

`func (o *PropagationStoreConfigurationSCIM) GetAUTHENTICATION_METHODOk() (*EnumPropagationStoreTypeSCIMAuthenticationMethod, bool)`

GetAUTHENTICATION_METHODOk returns a tuple with the AUTHENTICATION_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfigurationSCIM) SetAUTHENTICATION_METHOD(v EnumPropagationStoreTypeSCIMAuthenticationMethod)`

SetAUTHENTICATION_METHOD sets AUTHENTICATION_METHOD field to given value.


### GetAUTHORIZATION_TYPE

`func (o *PropagationStoreConfigurationSCIM) GetAUTHORIZATION_TYPE() string`

GetAUTHORIZATION_TYPE returns the AUTHORIZATION_TYPE field if non-nil, zero value otherwise.

### GetAUTHORIZATION_TYPEOk

`func (o *PropagationStoreConfigurationSCIM) GetAUTHORIZATION_TYPEOk() (*string, bool)`

GetAUTHORIZATION_TYPEOk returns a tuple with the AUTHORIZATION_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHORIZATION_TYPE

`func (o *PropagationStoreConfigurationSCIM) SetAUTHORIZATION_TYPE(v string)`

SetAUTHORIZATION_TYPE sets AUTHORIZATION_TYPE field to given value.


### GetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_PASSWORD() string`

GetBASIC_AUTH_PASSWORD returns the BASIC_AUTH_PASSWORD field if non-nil, zero value otherwise.

### GetBASIC_AUTH_PASSWORDOk

`func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_PASSWORDOk() (*string, bool)`

GetBASIC_AUTH_PASSWORDOk returns a tuple with the BASIC_AUTH_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfigurationSCIM) SetBASIC_AUTH_PASSWORD(v string)`

SetBASIC_AUTH_PASSWORD sets BASIC_AUTH_PASSWORD field to given value.

### HasBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfigurationSCIM) HasBASIC_AUTH_PASSWORD() bool`

HasBASIC_AUTH_PASSWORD returns a boolean if a field has been set.

### GetBASIC_AUTH_USER

`func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_USER() string`

GetBASIC_AUTH_USER returns the BASIC_AUTH_USER field if non-nil, zero value otherwise.

### GetBASIC_AUTH_USEROk

`func (o *PropagationStoreConfigurationSCIM) GetBASIC_AUTH_USEROk() (*string, bool)`

GetBASIC_AUTH_USEROk returns a tuple with the BASIC_AUTH_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_USER

`func (o *PropagationStoreConfigurationSCIM) SetBASIC_AUTH_USER(v string)`

SetBASIC_AUTH_USER sets BASIC_AUTH_USER field to given value.

### HasBASIC_AUTH_USER

`func (o *PropagationStoreConfigurationSCIM) HasBASIC_AUTH_USER() bool`

HasBASIC_AUTH_USER returns a boolean if a field has been set.

### GetCREATE_USERS

`func (o *PropagationStoreConfigurationSCIM) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationSCIM) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationSCIM) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationSCIM) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationSCIM) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationSCIM) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationSCIM) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationSCIM) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationSCIM) GetGROUP_NAME_SOURCE() EnumPropagationStoreTypeSCIMGroupNameSource`

GetGROUP_NAME_SOURCE returns the GROUP_NAME_SOURCE field if non-nil, zero value otherwise.

### GetGROUP_NAME_SOURCEOk

`func (o *PropagationStoreConfigurationSCIM) GetGROUP_NAME_SOURCEOk() (*EnumPropagationStoreTypeSCIMGroupNameSource, bool)`

GetGROUP_NAME_SOURCEOk returns a tuple with the GROUP_NAME_SOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationSCIM) SetGROUP_NAME_SOURCE(v EnumPropagationStoreTypeSCIMGroupNameSource)`

SetGROUP_NAME_SOURCE sets GROUP_NAME_SOURCE field to given value.

### HasGROUP_NAME_SOURCE

`func (o *PropagationStoreConfigurationSCIM) HasGROUP_NAME_SOURCE() bool`

HasGROUP_NAME_SOURCE returns a boolean if a field has been set.

### GetGROUPS_RESOURCE

`func (o *PropagationStoreConfigurationSCIM) GetGROUPS_RESOURCE() string`

GetGROUPS_RESOURCE returns the GROUPS_RESOURCE field if non-nil, zero value otherwise.

### GetGROUPS_RESOURCEOk

`func (o *PropagationStoreConfigurationSCIM) GetGROUPS_RESOURCEOk() (*string, bool)`

GetGROUPS_RESOURCEOk returns a tuple with the GROUPS_RESOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUPS_RESOURCE

`func (o *PropagationStoreConfigurationSCIM) SetGROUPS_RESOURCE(v string)`

SetGROUPS_RESOURCE sets GROUPS_RESOURCE field to given value.

### HasGROUPS_RESOURCE

`func (o *PropagationStoreConfigurationSCIM) HasGROUPS_RESOURCE() bool`

HasGROUPS_RESOURCE returns a boolean if a field has been set.

### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSCIM) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.

### HasOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSCIM) HasOAUTH_ACCESS_TOKEN() bool`

HasOAUTH_ACCESS_TOKEN returns a boolean if a field has been set.

### GetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_ID() string`

GetOAUTH_CLIENT_ID returns the OAUTH_CLIENT_ID field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_IDOk

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_IDOk() (*string, bool)`

GetOAUTH_CLIENT_IDOk returns a tuple with the OAUTH_CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationSCIM) SetOAUTH_CLIENT_ID(v string)`

SetOAUTH_CLIENT_ID sets OAUTH_CLIENT_ID field to given value.

### HasOAUTH_CLIENT_ID

`func (o *PropagationStoreConfigurationSCIM) HasOAUTH_CLIENT_ID() bool`

HasOAUTH_CLIENT_ID returns a boolean if a field has been set.

### GetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_SECRET() string`

GetOAUTH_CLIENT_SECRET returns the OAUTH_CLIENT_SECRET field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_SECRETOk

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_CLIENT_SECRETOk() (*string, bool)`

GetOAUTH_CLIENT_SECRETOk returns a tuple with the OAUTH_CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationSCIM) SetOAUTH_CLIENT_SECRET(v string)`

SetOAUTH_CLIENT_SECRET sets OAUTH_CLIENT_SECRET field to given value.

### HasOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfigurationSCIM) HasOAUTH_CLIENT_SECRET() bool`

HasOAUTH_CLIENT_SECRET returns a boolean if a field has been set.

### GetOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_TOKEN_REQUEST() string`

GetOAUTH_TOKEN_REQUEST returns the OAUTH_TOKEN_REQUEST field if non-nil, zero value otherwise.

### GetOAUTH_TOKEN_REQUESTOk

`func (o *PropagationStoreConfigurationSCIM) GetOAUTH_TOKEN_REQUESTOk() (*string, bool)`

GetOAUTH_TOKEN_REQUESTOk returns a tuple with the OAUTH_TOKEN_REQUEST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfigurationSCIM) SetOAUTH_TOKEN_REQUEST(v string)`

SetOAUTH_TOKEN_REQUEST sets OAUTH_TOKEN_REQUEST field to given value.

### HasOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfigurationSCIM) HasOAUTH_TOKEN_REQUEST() bool`

HasOAUTH_TOKEN_REQUEST returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfigurationSCIM) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfigurationSCIM) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfigurationSCIM) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfigurationSCIM) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetSCIM_URL

`func (o *PropagationStoreConfigurationSCIM) GetSCIM_URL() string`

GetSCIM_URL returns the SCIM_URL field if non-nil, zero value otherwise.

### GetSCIM_URLOk

`func (o *PropagationStoreConfigurationSCIM) GetSCIM_URLOk() (*string, bool)`

GetSCIM_URLOk returns a tuple with the SCIM_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_URL

`func (o *PropagationStoreConfigurationSCIM) SetSCIM_URL(v string)`

SetSCIM_URL sets SCIM_URL field to given value.


### GetSCIM_VERSION

`func (o *PropagationStoreConfigurationSCIM) GetSCIM_VERSION() string`

GetSCIM_VERSION returns the SCIM_VERSION field if non-nil, zero value otherwise.

### GetSCIM_VERSIONOk

`func (o *PropagationStoreConfigurationSCIM) GetSCIM_VERSIONOk() (*string, bool)`

GetSCIM_VERSIONOk returns a tuple with the SCIM_VERSION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_VERSION

`func (o *PropagationStoreConfigurationSCIM) SetSCIM_VERSION(v string)`

SetSCIM_VERSION sets SCIM_VERSION field to given value.


### GetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfigurationSCIM) GetUNIQUE_USER_IDENTIFIER() EnumPropagationStoreTypeSCIMUniqueUserIdentifier`

GetUNIQUE_USER_IDENTIFIER returns the UNIQUE_USER_IDENTIFIER field if non-nil, zero value otherwise.

### GetUNIQUE_USER_IDENTIFIEROk

`func (o *PropagationStoreConfigurationSCIM) GetUNIQUE_USER_IDENTIFIEROk() (*EnumPropagationStoreTypeSCIMUniqueUserIdentifier, bool)`

GetUNIQUE_USER_IDENTIFIEROk returns a tuple with the UNIQUE_USER_IDENTIFIER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfigurationSCIM) SetUNIQUE_USER_IDENTIFIER(v EnumPropagationStoreTypeSCIMUniqueUserIdentifier)`

SetUNIQUE_USER_IDENTIFIER sets UNIQUE_USER_IDENTIFIER field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationSCIM) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationSCIM) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationSCIM) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationSCIM) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.

### GetUSER_FILTER

`func (o *PropagationStoreConfigurationSCIM) GetUSER_FILTER() string`

GetUSER_FILTER returns the USER_FILTER field if non-nil, zero value otherwise.

### GetUSER_FILTEROk

`func (o *PropagationStoreConfigurationSCIM) GetUSER_FILTEROk() (*string, bool)`

GetUSER_FILTEROk returns a tuple with the USER_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_FILTER

`func (o *PropagationStoreConfigurationSCIM) SetUSER_FILTER(v string)`

SetUSER_FILTER sets USER_FILTER field to given value.


### GetUSERS_RESOURCE

`func (o *PropagationStoreConfigurationSCIM) GetUSERS_RESOURCE() string`

GetUSERS_RESOURCE returns the USERS_RESOURCE field if non-nil, zero value otherwise.

### GetUSERS_RESOURCEOk

`func (o *PropagationStoreConfigurationSCIM) GetUSERS_RESOURCEOk() (*string, bool)`

GetUSERS_RESOURCEOk returns a tuple with the USERS_RESOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERS_RESOURCE

`func (o *PropagationStoreConfigurationSCIM) SetUSERS_RESOURCE(v string)`

SetUSERS_RESOURCE sets USERS_RESOURCE field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


