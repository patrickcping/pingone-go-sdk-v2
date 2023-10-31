# PropagationStoreConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCESS_TOKEN** | Pointer to **string** | A string specifying the access token for account authentication. | [optional] 
**AUTHENTICATION_METHOD** | [**EnumPropagationStoreTypeZoomAuthenticationMethod**](EnumPropagationStoreTypeZoomAuthenticationMethod.md) |  | [default to ENUMPROPAGATIONSTORETYPEZOOMAUTHENTICATIONMETHOD_JWT_BEARER_TOKEN]
**BASIC_AUTH_PASSWORD** | **string** | The password for account authentication. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;Basic Authentication&#x60;, otherwise optional. | 
**BASIC_AUTH_USER** | **string** | The user name for account authentication. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;Basic Authentication&#x60;, otherwise optional. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**GROUP_NAME_SOURCE** | Pointer to [**EnumPropagationStoreTypeSCIMGroupNameSource**](EnumPropagationStoreTypeSCIMGroupNameSource.md) |  | [optional] 
**REMOVE_ACTION** | Pointer to [**EnumPropagationStoreTypeRemoveActionDisableDelete**](EnumPropagationStoreTypeRemoveActionDisableDelete.md) |  | [optional] 
**SCIM_URL** | **string** | The SCIM URL. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 
**ClientId** | **string** | The Azure Active Directory client ID. | 
**ClientSecret** | **string** | The Azure Active Directory client secret. | 
**DEPROVISION_USERS** | Pointer to **bool** | Whether or not users are allowed to be removed (deprovisioned) following the action configured in &#x60;REMOVE_ACTION&#x60;. | [optional] 
**PROVISION_DISABLED_USERS_PROV_OPT** | Pointer to **bool** | Whether or not disabled users can be provisioned. | [optional] 
**RemoveLicensesWhenSkuIdEmpty** | **bool** | Whether or not remove licenses from user when &#x60;skuId&#x60; is empty. | 
**TenantDomain** | **string** | The account&#39;s Azure Active Directory domain. | 
**BASE_URL** | **string** | Base URL of the target propagation store. | 
**OAUTH_ACCESS_TOKEN** | **string** | OAuth 2 access token. | 
**APPLICATION_NAME** | **string** | Name of the application using the store. | 
**DOMAIN** | **string** | The account&#39;s domain name. | 
**OAUTH_CLIENT_ID** | **string** | OAuth client identifier. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | 
**OAUTH_CLIENT_SECRET** | **string** | OAuth client secret. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | 
**OAUTH_REFRESH_TOKEN** | **string** | OAuth refresh token for account authentication. | 
**ATTRIBUTE_METADATA** | Pointer to **string** | User-defined attribute metadata. | [optional] 
**CLIENT_ID** | **string** | Unique identifier (UUID) of the Salesforce client associated with the propagation store. | 
**CLIENT_SECRET** | **string** | The Salesforce client secret. | 
**DELETE_USERS** | Pointer to **bool** | Whether or not users are allowed to be deleted. | [optional] 
**ENVIRONMENT_ID** | **string** | Unique identifier (UUID) of the PingOne environment associated with the propagation store. | 
**GATEWAY_BASE_URL** | **string** | Base URL of the gateway. | 
**GATEWAY_ID** | **string** | Identifier of the gateway to which the connector connects. | 
**LDAP_TYPE** | [**EnumPropagationStoreTypeLDAPGatewayLDAPType**](EnumPropagationStoreTypeLDAPGatewayLDAPType.md) |  | 
**OAUTH_URL** | **string** | OAuth token request endpoint. | 
**DEFAULT_AUTH_METHOD** | Pointer to [**EnumPropagationStoreTypePingOneDefaultAuthMethod**](EnumPropagationStoreTypePingOneDefaultAuthMethod.md) |  | [optional] 
**MFA_USER_DEVICE_MANAGEMENT** | Pointer to **string** | How to manage MFA user devices when synchronizing. Options are either Merge with devices in PingOne or Overwrite devices in PingOne. | [optional] 
**REGION** | [**EnumPropagationStoreTypePingOneRegion**](EnumPropagationStoreTypePingOneRegion.md) |  | 
**ACCOUNT_ID** | Pointer to **string** | The Salesforce account ID. | [optional] 
**ENABLE_COMMUNITIES** | **bool** | Whether or not to enable Salesforce communities. | 
**FREEZE_USER_FLAG** | **bool** | Whether the user account is frozen. | 
**PERMISSION_SET_MANAGEMENT** | **string** | The permission sets to be merged with Salesforce. | 
**PROFILE_ID** | Pointer to **string** | The Salesforce profile ID. | [optional] 
**PROVISION_DISABLED_USERS** | Pointer to **bool** | Whether or not disabled users can be provisioned. | [optional] 
**SALESFORCE_DOMAIN** | **string** | The account&#39;s salesforce.com domain. | 
**RECORD_TYPE** | [**EnumPropagationStoreTypeSalesforceContactsRecordType**](EnumPropagationStoreTypeSalesforceContactsRecordType.md) |  | 
**AUTHORIZATION_TYPE** | **string** | The authorization header type. | 
**GROUPS_RESOURCE** | Pointer to **string** | API endpoint path to the group entity. | [optional] 
**OAUTH_TOKEN_REQUEST** | Pointer to **string** | OAuth token request endpoint. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth 2 Bearer Token&#x60;, otherwise optional. | [optional] 
**SCIM_VERSION** | **string** | The SCIM version. | 
**UNIQUE_USER_IDENTIFIER** | [**EnumPropagationStoreTypeSlackUniqueUserIdentifier**](EnumPropagationStoreTypeSlackUniqueUserIdentifier.md) |  | 
**USER_FILTER** | **string** | A string that specifies a SCIM filter expression. | 
**USERS_RESOURCE** | **string** | API endpoint path to the user entity. | 
**AdministratorPassword** | **string** | Password for the administrator. | 
**AdministratorUsername** | **string** | Username for the administrator. | 
**ServiceNowUrl** | **string** | The URL for the ServiceNow account. | 
**ExcludeContingentWorkers** | Pointer to **bool** | Whether or not contingent workers are excluded. | [optional] 
**ExcludeEmployees** | Pointer to **bool** | Whether or not employees are excluded. | [optional] 
**ExcludeInactiveWorkers** | Pointer to **bool** | Whether or not inactive workers are excluded. | [optional] 
**Host** | Pointer to **string** | The Workday API host. | [optional] 
**Password** | **string** | The password for account authentication. | 
**TenantId** | **string** | The Workday tenant ID. | 
**Username** | **string** | The user name for account authentication. | 
**API_KEY** | Pointer to **string** | The client API key. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;JWT Bearer Token&#x60;, otherwise optional. | [optional] 
**API_SECRET** | Pointer to **string** | The client API secret. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;JWT Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_ACCOUNT_ID** | Pointer to **string** | OAuth account identifier. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 
**OAUTH_TOKEN_URL** | Pointer to **string** | OAuth token request endpoint. Required when &#x60;AUTHENTICATION_METHOD&#x60; is &#x60;OAuth Bearer Token&#x60;, otherwise optional. | [optional] 

## Methods

### NewPropagationStoreConfiguration

`func NewPropagationStoreConfiguration(aUTHENTICATIONMETHOD EnumPropagationStoreTypeZoomAuthenticationMethod, bASICAUTHPASSWORD string, bASICAUTHUSER string, sCIMURL string, clientId string, clientSecret string, removeLicensesWhenSkuIdEmpty bool, tenantDomain string, bASEURL string, oAUTHACCESSTOKEN string, aPPLICATIONNAME string, dOMAIN string, oAUTHCLIENTID string, oAUTHCLIENTSECRET string, oAUTHREFRESHTOKEN string, cLIENTID string, cLIENTSECRET string, eNVIRONMENTID string, gATEWAYBASEURL string, gATEWAYID string, lDAPTYPE EnumPropagationStoreTypeLDAPGatewayLDAPType, oAUTHURL string, rEGION EnumPropagationStoreTypePingOneRegion, eNABLECOMMUNITIES bool, fREEZEUSERFLAG bool, pERMISSIONSETMANAGEMENT string, sALESFORCEDOMAIN string, rECORDTYPE EnumPropagationStoreTypeSalesforceContactsRecordType, aUTHORIZATIONTYPE string, sCIMVERSION string, uNIQUEUSERIDENTIFIER EnumPropagationStoreTypeSlackUniqueUserIdentifier, uSERFILTER string, uSERSRESOURCE string, administratorPassword string, administratorUsername string, serviceNowUrl string, password string, tenantId string, username string, ) *PropagationStoreConfiguration`

NewPropagationStoreConfiguration instantiates a new PropagationStoreConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationWithDefaults

`func NewPropagationStoreConfigurationWithDefaults() *PropagationStoreConfiguration`

NewPropagationStoreConfigurationWithDefaults instantiates a new PropagationStoreConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCESS_TOKEN

`func (o *PropagationStoreConfiguration) GetACCESS_TOKEN() string`

GetACCESS_TOKEN returns the ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetACCESS_TOKENOk

`func (o *PropagationStoreConfiguration) GetACCESS_TOKENOk() (*string, bool)`

GetACCESS_TOKENOk returns a tuple with the ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCESS_TOKEN

`func (o *PropagationStoreConfiguration) SetACCESS_TOKEN(v string)`

SetACCESS_TOKEN sets ACCESS_TOKEN field to given value.

### HasACCESS_TOKEN

`func (o *PropagationStoreConfiguration) HasACCESS_TOKEN() bool`

HasACCESS_TOKEN returns a boolean if a field has been set.

### GetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfiguration) GetAUTHENTICATION_METHOD() EnumPropagationStoreTypeZoomAuthenticationMethod`

GetAUTHENTICATION_METHOD returns the AUTHENTICATION_METHOD field if non-nil, zero value otherwise.

### GetAUTHENTICATION_METHODOk

`func (o *PropagationStoreConfiguration) GetAUTHENTICATION_METHODOk() (*EnumPropagationStoreTypeZoomAuthenticationMethod, bool)`

GetAUTHENTICATION_METHODOk returns a tuple with the AUTHENTICATION_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHENTICATION_METHOD

`func (o *PropagationStoreConfiguration) SetAUTHENTICATION_METHOD(v EnumPropagationStoreTypeZoomAuthenticationMethod)`

SetAUTHENTICATION_METHOD sets AUTHENTICATION_METHOD field to given value.


### GetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfiguration) GetBASIC_AUTH_PASSWORD() string`

GetBASIC_AUTH_PASSWORD returns the BASIC_AUTH_PASSWORD field if non-nil, zero value otherwise.

### GetBASIC_AUTH_PASSWORDOk

`func (o *PropagationStoreConfiguration) GetBASIC_AUTH_PASSWORDOk() (*string, bool)`

GetBASIC_AUTH_PASSWORDOk returns a tuple with the BASIC_AUTH_PASSWORD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_PASSWORD

`func (o *PropagationStoreConfiguration) SetBASIC_AUTH_PASSWORD(v string)`

SetBASIC_AUTH_PASSWORD sets BASIC_AUTH_PASSWORD field to given value.


### GetBASIC_AUTH_USER

`func (o *PropagationStoreConfiguration) GetBASIC_AUTH_USER() string`

GetBASIC_AUTH_USER returns the BASIC_AUTH_USER field if non-nil, zero value otherwise.

### GetBASIC_AUTH_USEROk

`func (o *PropagationStoreConfiguration) GetBASIC_AUTH_USEROk() (*string, bool)`

GetBASIC_AUTH_USEROk returns a tuple with the BASIC_AUTH_USER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASIC_AUTH_USER

`func (o *PropagationStoreConfiguration) SetBASIC_AUTH_USER(v string)`

SetBASIC_AUTH_USER sets BASIC_AUTH_USER field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfiguration) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfiguration) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfiguration) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfiguration) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfiguration) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfiguration) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfiguration) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfiguration) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfiguration) GetGROUP_NAME_SOURCE() EnumPropagationStoreTypeSCIMGroupNameSource`

GetGROUP_NAME_SOURCE returns the GROUP_NAME_SOURCE field if non-nil, zero value otherwise.

### GetGROUP_NAME_SOURCEOk

`func (o *PropagationStoreConfiguration) GetGROUP_NAME_SOURCEOk() (*EnumPropagationStoreTypeSCIMGroupNameSource, bool)`

GetGROUP_NAME_SOURCEOk returns a tuple with the GROUP_NAME_SOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUP_NAME_SOURCE

`func (o *PropagationStoreConfiguration) SetGROUP_NAME_SOURCE(v EnumPropagationStoreTypeSCIMGroupNameSource)`

SetGROUP_NAME_SOURCE sets GROUP_NAME_SOURCE field to given value.

### HasGROUP_NAME_SOURCE

`func (o *PropagationStoreConfiguration) HasGROUP_NAME_SOURCE() bool`

HasGROUP_NAME_SOURCE returns a boolean if a field has been set.

### GetREMOVE_ACTION

`func (o *PropagationStoreConfiguration) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisableDelete`

GetREMOVE_ACTION returns the REMOVE_ACTION field if non-nil, zero value otherwise.

### GetREMOVE_ACTIONOk

`func (o *PropagationStoreConfiguration) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisableDelete, bool)`

GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREMOVE_ACTION

`func (o *PropagationStoreConfiguration) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisableDelete)`

SetREMOVE_ACTION sets REMOVE_ACTION field to given value.

### HasREMOVE_ACTION

`func (o *PropagationStoreConfiguration) HasREMOVE_ACTION() bool`

HasREMOVE_ACTION returns a boolean if a field has been set.

### GetSCIM_URL

`func (o *PropagationStoreConfiguration) GetSCIM_URL() string`

GetSCIM_URL returns the SCIM_URL field if non-nil, zero value otherwise.

### GetSCIM_URLOk

`func (o *PropagationStoreConfiguration) GetSCIM_URLOk() (*string, bool)`

GetSCIM_URLOk returns a tuple with the SCIM_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_URL

`func (o *PropagationStoreConfiguration) SetSCIM_URL(v string)`

SetSCIM_URL sets SCIM_URL field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfiguration) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfiguration) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfiguration) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfiguration) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.

### GetClientId

`func (o *PropagationStoreConfiguration) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PropagationStoreConfiguration) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PropagationStoreConfiguration) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PropagationStoreConfiguration) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PropagationStoreConfiguration) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PropagationStoreConfiguration) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDEPROVISION_USERS

`func (o *PropagationStoreConfiguration) GetDEPROVISION_USERS() bool`

GetDEPROVISION_USERS returns the DEPROVISION_USERS field if non-nil, zero value otherwise.

### GetDEPROVISION_USERSOk

`func (o *PropagationStoreConfiguration) GetDEPROVISION_USERSOk() (*bool, bool)`

GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEPROVISION_USERS

`func (o *PropagationStoreConfiguration) SetDEPROVISION_USERS(v bool)`

SetDEPROVISION_USERS sets DEPROVISION_USERS field to given value.

### HasDEPROVISION_USERS

`func (o *PropagationStoreConfiguration) HasDEPROVISION_USERS() bool`

HasDEPROVISION_USERS returns a boolean if a field has been set.

### GetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfiguration) GetPROVISION_DISABLED_USERS_PROV_OPT() bool`

GetPROVISION_DISABLED_USERS_PROV_OPT returns the PROVISION_DISABLED_USERS_PROV_OPT field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERS_PROV_OPTOk

`func (o *PropagationStoreConfiguration) GetPROVISION_DISABLED_USERS_PROV_OPTOk() (*bool, bool)`

GetPROVISION_DISABLED_USERS_PROV_OPTOk returns a tuple with the PROVISION_DISABLED_USERS_PROV_OPT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfiguration) SetPROVISION_DISABLED_USERS_PROV_OPT(v bool)`

SetPROVISION_DISABLED_USERS_PROV_OPT sets PROVISION_DISABLED_USERS_PROV_OPT field to given value.

### HasPROVISION_DISABLED_USERS_PROV_OPT

`func (o *PropagationStoreConfiguration) HasPROVISION_DISABLED_USERS_PROV_OPT() bool`

HasPROVISION_DISABLED_USERS_PROV_OPT returns a boolean if a field has been set.

### GetRemoveLicensesWhenSkuIdEmpty

`func (o *PropagationStoreConfiguration) GetRemoveLicensesWhenSkuIdEmpty() bool`

GetRemoveLicensesWhenSkuIdEmpty returns the RemoveLicensesWhenSkuIdEmpty field if non-nil, zero value otherwise.

### GetRemoveLicensesWhenSkuIdEmptyOk

`func (o *PropagationStoreConfiguration) GetRemoveLicensesWhenSkuIdEmptyOk() (*bool, bool)`

GetRemoveLicensesWhenSkuIdEmptyOk returns a tuple with the RemoveLicensesWhenSkuIdEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveLicensesWhenSkuIdEmpty

`func (o *PropagationStoreConfiguration) SetRemoveLicensesWhenSkuIdEmpty(v bool)`

SetRemoveLicensesWhenSkuIdEmpty sets RemoveLicensesWhenSkuIdEmpty field to given value.


### GetTenantDomain

`func (o *PropagationStoreConfiguration) GetTenantDomain() string`

GetTenantDomain returns the TenantDomain field if non-nil, zero value otherwise.

### GetTenantDomainOk

`func (o *PropagationStoreConfiguration) GetTenantDomainOk() (*string, bool)`

GetTenantDomainOk returns a tuple with the TenantDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantDomain

`func (o *PropagationStoreConfiguration) SetTenantDomain(v string)`

SetTenantDomain sets TenantDomain field to given value.


### GetBASE_URL

`func (o *PropagationStoreConfiguration) GetBASE_URL() string`

GetBASE_URL returns the BASE_URL field if non-nil, zero value otherwise.

### GetBASE_URLOk

`func (o *PropagationStoreConfiguration) GetBASE_URLOk() (*string, bool)`

GetBASE_URLOk returns a tuple with the BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBASE_URL

`func (o *PropagationStoreConfiguration) SetBASE_URL(v string)`

SetBASE_URL sets BASE_URL field to given value.


### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfiguration) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfiguration) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfiguration) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.


### GetAPPLICATION_NAME

`func (o *PropagationStoreConfiguration) GetAPPLICATION_NAME() string`

GetAPPLICATION_NAME returns the APPLICATION_NAME field if non-nil, zero value otherwise.

### GetAPPLICATION_NAMEOk

`func (o *PropagationStoreConfiguration) GetAPPLICATION_NAMEOk() (*string, bool)`

GetAPPLICATION_NAMEOk returns a tuple with the APPLICATION_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPPLICATION_NAME

`func (o *PropagationStoreConfiguration) SetAPPLICATION_NAME(v string)`

SetAPPLICATION_NAME sets APPLICATION_NAME field to given value.


### GetDOMAIN

`func (o *PropagationStoreConfiguration) GetDOMAIN() string`

GetDOMAIN returns the DOMAIN field if non-nil, zero value otherwise.

### GetDOMAINOk

`func (o *PropagationStoreConfiguration) GetDOMAINOk() (*string, bool)`

GetDOMAINOk returns a tuple with the DOMAIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDOMAIN

`func (o *PropagationStoreConfiguration) SetDOMAIN(v string)`

SetDOMAIN sets DOMAIN field to given value.


### GetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfiguration) GetOAUTH_CLIENT_ID() string`

GetOAUTH_CLIENT_ID returns the OAUTH_CLIENT_ID field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_IDOk

`func (o *PropagationStoreConfiguration) GetOAUTH_CLIENT_IDOk() (*string, bool)`

GetOAUTH_CLIENT_IDOk returns a tuple with the OAUTH_CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_ID

`func (o *PropagationStoreConfiguration) SetOAUTH_CLIENT_ID(v string)`

SetOAUTH_CLIENT_ID sets OAUTH_CLIENT_ID field to given value.


### GetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfiguration) GetOAUTH_CLIENT_SECRET() string`

GetOAUTH_CLIENT_SECRET returns the OAUTH_CLIENT_SECRET field if non-nil, zero value otherwise.

### GetOAUTH_CLIENT_SECRETOk

`func (o *PropagationStoreConfiguration) GetOAUTH_CLIENT_SECRETOk() (*string, bool)`

GetOAUTH_CLIENT_SECRETOk returns a tuple with the OAUTH_CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_CLIENT_SECRET

`func (o *PropagationStoreConfiguration) SetOAUTH_CLIENT_SECRET(v string)`

SetOAUTH_CLIENT_SECRET sets OAUTH_CLIENT_SECRET field to given value.


### GetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfiguration) GetOAUTH_REFRESH_TOKEN() string`

GetOAUTH_REFRESH_TOKEN returns the OAUTH_REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_REFRESH_TOKENOk

`func (o *PropagationStoreConfiguration) GetOAUTH_REFRESH_TOKENOk() (*string, bool)`

GetOAUTH_REFRESH_TOKENOk returns a tuple with the OAUTH_REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfiguration) SetOAUTH_REFRESH_TOKEN(v string)`

SetOAUTH_REFRESH_TOKEN sets OAUTH_REFRESH_TOKEN field to given value.


### GetATTRIBUTE_METADATA

`func (o *PropagationStoreConfiguration) GetATTRIBUTE_METADATA() string`

GetATTRIBUTE_METADATA returns the ATTRIBUTE_METADATA field if non-nil, zero value otherwise.

### GetATTRIBUTE_METADATAOk

`func (o *PropagationStoreConfiguration) GetATTRIBUTE_METADATAOk() (*string, bool)`

GetATTRIBUTE_METADATAOk returns a tuple with the ATTRIBUTE_METADATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATTRIBUTE_METADATA

`func (o *PropagationStoreConfiguration) SetATTRIBUTE_METADATA(v string)`

SetATTRIBUTE_METADATA sets ATTRIBUTE_METADATA field to given value.

### HasATTRIBUTE_METADATA

`func (o *PropagationStoreConfiguration) HasATTRIBUTE_METADATA() bool`

HasATTRIBUTE_METADATA returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *PropagationStoreConfiguration) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *PropagationStoreConfiguration) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *PropagationStoreConfiguration) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *PropagationStoreConfiguration) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *PropagationStoreConfiguration) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *PropagationStoreConfiguration) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetDELETE_USERS

`func (o *PropagationStoreConfiguration) GetDELETE_USERS() bool`

GetDELETE_USERS returns the DELETE_USERS field if non-nil, zero value otherwise.

### GetDELETE_USERSOk

`func (o *PropagationStoreConfiguration) GetDELETE_USERSOk() (*bool, bool)`

GetDELETE_USERSOk returns a tuple with the DELETE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE_USERS

`func (o *PropagationStoreConfiguration) SetDELETE_USERS(v bool)`

SetDELETE_USERS sets DELETE_USERS field to given value.

### HasDELETE_USERS

`func (o *PropagationStoreConfiguration) HasDELETE_USERS() bool`

HasDELETE_USERS returns a boolean if a field has been set.

### GetENVIRONMENT_ID

`func (o *PropagationStoreConfiguration) GetENVIRONMENT_ID() string`

GetENVIRONMENT_ID returns the ENVIRONMENT_ID field if non-nil, zero value otherwise.

### GetENVIRONMENT_IDOk

`func (o *PropagationStoreConfiguration) GetENVIRONMENT_IDOk() (*string, bool)`

GetENVIRONMENT_IDOk returns a tuple with the ENVIRONMENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENVIRONMENT_ID

`func (o *PropagationStoreConfiguration) SetENVIRONMENT_ID(v string)`

SetENVIRONMENT_ID sets ENVIRONMENT_ID field to given value.


### GetGATEWAY_BASE_URL

`func (o *PropagationStoreConfiguration) GetGATEWAY_BASE_URL() string`

GetGATEWAY_BASE_URL returns the GATEWAY_BASE_URL field if non-nil, zero value otherwise.

### GetGATEWAY_BASE_URLOk

`func (o *PropagationStoreConfiguration) GetGATEWAY_BASE_URLOk() (*string, bool)`

GetGATEWAY_BASE_URLOk returns a tuple with the GATEWAY_BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGATEWAY_BASE_URL

`func (o *PropagationStoreConfiguration) SetGATEWAY_BASE_URL(v string)`

SetGATEWAY_BASE_URL sets GATEWAY_BASE_URL field to given value.


### GetGATEWAY_ID

`func (o *PropagationStoreConfiguration) GetGATEWAY_ID() string`

GetGATEWAY_ID returns the GATEWAY_ID field if non-nil, zero value otherwise.

### GetGATEWAY_IDOk

`func (o *PropagationStoreConfiguration) GetGATEWAY_IDOk() (*string, bool)`

GetGATEWAY_IDOk returns a tuple with the GATEWAY_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGATEWAY_ID

`func (o *PropagationStoreConfiguration) SetGATEWAY_ID(v string)`

SetGATEWAY_ID sets GATEWAY_ID field to given value.


### GetLDAP_TYPE

`func (o *PropagationStoreConfiguration) GetLDAP_TYPE() EnumPropagationStoreTypeLDAPGatewayLDAPType`

GetLDAP_TYPE returns the LDAP_TYPE field if non-nil, zero value otherwise.

### GetLDAP_TYPEOk

`func (o *PropagationStoreConfiguration) GetLDAP_TYPEOk() (*EnumPropagationStoreTypeLDAPGatewayLDAPType, bool)`

GetLDAP_TYPEOk returns a tuple with the LDAP_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_TYPE

`func (o *PropagationStoreConfiguration) SetLDAP_TYPE(v EnumPropagationStoreTypeLDAPGatewayLDAPType)`

SetLDAP_TYPE sets LDAP_TYPE field to given value.


### GetOAUTH_URL

`func (o *PropagationStoreConfiguration) GetOAUTH_URL() string`

GetOAUTH_URL returns the OAUTH_URL field if non-nil, zero value otherwise.

### GetOAUTH_URLOk

`func (o *PropagationStoreConfiguration) GetOAUTH_URLOk() (*string, bool)`

GetOAUTH_URLOk returns a tuple with the OAUTH_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_URL

`func (o *PropagationStoreConfiguration) SetOAUTH_URL(v string)`

SetOAUTH_URL sets OAUTH_URL field to given value.


### GetDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfiguration) GetDEFAULT_AUTH_METHOD() EnumPropagationStoreTypePingOneDefaultAuthMethod`

GetDEFAULT_AUTH_METHOD returns the DEFAULT_AUTH_METHOD field if non-nil, zero value otherwise.

### GetDEFAULT_AUTH_METHODOk

`func (o *PropagationStoreConfiguration) GetDEFAULT_AUTH_METHODOk() (*EnumPropagationStoreTypePingOneDefaultAuthMethod, bool)`

GetDEFAULT_AUTH_METHODOk returns a tuple with the DEFAULT_AUTH_METHOD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfiguration) SetDEFAULT_AUTH_METHOD(v EnumPropagationStoreTypePingOneDefaultAuthMethod)`

SetDEFAULT_AUTH_METHOD sets DEFAULT_AUTH_METHOD field to given value.

### HasDEFAULT_AUTH_METHOD

`func (o *PropagationStoreConfiguration) HasDEFAULT_AUTH_METHOD() bool`

HasDEFAULT_AUTH_METHOD returns a boolean if a field has been set.

### GetMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfiguration) GetMFA_USER_DEVICE_MANAGEMENT() string`

GetMFA_USER_DEVICE_MANAGEMENT returns the MFA_USER_DEVICE_MANAGEMENT field if non-nil, zero value otherwise.

### GetMFA_USER_DEVICE_MANAGEMENTOk

`func (o *PropagationStoreConfiguration) GetMFA_USER_DEVICE_MANAGEMENTOk() (*string, bool)`

GetMFA_USER_DEVICE_MANAGEMENTOk returns a tuple with the MFA_USER_DEVICE_MANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfiguration) SetMFA_USER_DEVICE_MANAGEMENT(v string)`

SetMFA_USER_DEVICE_MANAGEMENT sets MFA_USER_DEVICE_MANAGEMENT field to given value.

### HasMFA_USER_DEVICE_MANAGEMENT

`func (o *PropagationStoreConfiguration) HasMFA_USER_DEVICE_MANAGEMENT() bool`

HasMFA_USER_DEVICE_MANAGEMENT returns a boolean if a field has been set.

### GetREGION

`func (o *PropagationStoreConfiguration) GetREGION() EnumPropagationStoreTypePingOneRegion`

GetREGION returns the REGION field if non-nil, zero value otherwise.

### GetREGIONOk

`func (o *PropagationStoreConfiguration) GetREGIONOk() (*EnumPropagationStoreTypePingOneRegion, bool)`

GetREGIONOk returns a tuple with the REGION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetREGION

`func (o *PropagationStoreConfiguration) SetREGION(v EnumPropagationStoreTypePingOneRegion)`

SetREGION sets REGION field to given value.


### GetACCOUNT_ID

`func (o *PropagationStoreConfiguration) GetACCOUNT_ID() string`

GetACCOUNT_ID returns the ACCOUNT_ID field if non-nil, zero value otherwise.

### GetACCOUNT_IDOk

`func (o *PropagationStoreConfiguration) GetACCOUNT_IDOk() (*string, bool)`

GetACCOUNT_IDOk returns a tuple with the ACCOUNT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ID

`func (o *PropagationStoreConfiguration) SetACCOUNT_ID(v string)`

SetACCOUNT_ID sets ACCOUNT_ID field to given value.

### HasACCOUNT_ID

`func (o *PropagationStoreConfiguration) HasACCOUNT_ID() bool`

HasACCOUNT_ID returns a boolean if a field has been set.

### GetENABLE_COMMUNITIES

`func (o *PropagationStoreConfiguration) GetENABLE_COMMUNITIES() bool`

GetENABLE_COMMUNITIES returns the ENABLE_COMMUNITIES field if non-nil, zero value otherwise.

### GetENABLE_COMMUNITIESOk

`func (o *PropagationStoreConfiguration) GetENABLE_COMMUNITIESOk() (*bool, bool)`

GetENABLE_COMMUNITIESOk returns a tuple with the ENABLE_COMMUNITIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_COMMUNITIES

`func (o *PropagationStoreConfiguration) SetENABLE_COMMUNITIES(v bool)`

SetENABLE_COMMUNITIES sets ENABLE_COMMUNITIES field to given value.


### GetFREEZE_USER_FLAG

`func (o *PropagationStoreConfiguration) GetFREEZE_USER_FLAG() bool`

GetFREEZE_USER_FLAG returns the FREEZE_USER_FLAG field if non-nil, zero value otherwise.

### GetFREEZE_USER_FLAGOk

`func (o *PropagationStoreConfiguration) GetFREEZE_USER_FLAGOk() (*bool, bool)`

GetFREEZE_USER_FLAGOk returns a tuple with the FREEZE_USER_FLAG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFREEZE_USER_FLAG

`func (o *PropagationStoreConfiguration) SetFREEZE_USER_FLAG(v bool)`

SetFREEZE_USER_FLAG sets FREEZE_USER_FLAG field to given value.


### GetPERMISSION_SET_MANAGEMENT

`func (o *PropagationStoreConfiguration) GetPERMISSION_SET_MANAGEMENT() string`

GetPERMISSION_SET_MANAGEMENT returns the PERMISSION_SET_MANAGEMENT field if non-nil, zero value otherwise.

### GetPERMISSION_SET_MANAGEMENTOk

`func (o *PropagationStoreConfiguration) GetPERMISSION_SET_MANAGEMENTOk() (*string, bool)`

GetPERMISSION_SET_MANAGEMENTOk returns a tuple with the PERMISSION_SET_MANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPERMISSION_SET_MANAGEMENT

`func (o *PropagationStoreConfiguration) SetPERMISSION_SET_MANAGEMENT(v string)`

SetPERMISSION_SET_MANAGEMENT sets PERMISSION_SET_MANAGEMENT field to given value.


### GetPROFILE_ID

`func (o *PropagationStoreConfiguration) GetPROFILE_ID() string`

GetPROFILE_ID returns the PROFILE_ID field if non-nil, zero value otherwise.

### GetPROFILE_IDOk

`func (o *PropagationStoreConfiguration) GetPROFILE_IDOk() (*string, bool)`

GetPROFILE_IDOk returns a tuple with the PROFILE_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROFILE_ID

`func (o *PropagationStoreConfiguration) SetPROFILE_ID(v string)`

SetPROFILE_ID sets PROFILE_ID field to given value.

### HasPROFILE_ID

`func (o *PropagationStoreConfiguration) HasPROFILE_ID() bool`

HasPROFILE_ID returns a boolean if a field has been set.

### GetPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfiguration) GetPROVISION_DISABLED_USERS() bool`

GetPROVISION_DISABLED_USERS returns the PROVISION_DISABLED_USERS field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERSOk

`func (o *PropagationStoreConfiguration) GetPROVISION_DISABLED_USERSOk() (*bool, bool)`

GetPROVISION_DISABLED_USERSOk returns a tuple with the PROVISION_DISABLED_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfiguration) SetPROVISION_DISABLED_USERS(v bool)`

SetPROVISION_DISABLED_USERS sets PROVISION_DISABLED_USERS field to given value.

### HasPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfiguration) HasPROVISION_DISABLED_USERS() bool`

HasPROVISION_DISABLED_USERS returns a boolean if a field has been set.

### GetSALESFORCE_DOMAIN

`func (o *PropagationStoreConfiguration) GetSALESFORCE_DOMAIN() string`

GetSALESFORCE_DOMAIN returns the SALESFORCE_DOMAIN field if non-nil, zero value otherwise.

### GetSALESFORCE_DOMAINOk

`func (o *PropagationStoreConfiguration) GetSALESFORCE_DOMAINOk() (*string, bool)`

GetSALESFORCE_DOMAINOk returns a tuple with the SALESFORCE_DOMAIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSALESFORCE_DOMAIN

`func (o *PropagationStoreConfiguration) SetSALESFORCE_DOMAIN(v string)`

SetSALESFORCE_DOMAIN sets SALESFORCE_DOMAIN field to given value.


### GetRECORD_TYPE

`func (o *PropagationStoreConfiguration) GetRECORD_TYPE() EnumPropagationStoreTypeSalesforceContactsRecordType`

GetRECORD_TYPE returns the RECORD_TYPE field if non-nil, zero value otherwise.

### GetRECORD_TYPEOk

`func (o *PropagationStoreConfiguration) GetRECORD_TYPEOk() (*EnumPropagationStoreTypeSalesforceContactsRecordType, bool)`

GetRECORD_TYPEOk returns a tuple with the RECORD_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRECORD_TYPE

`func (o *PropagationStoreConfiguration) SetRECORD_TYPE(v EnumPropagationStoreTypeSalesforceContactsRecordType)`

SetRECORD_TYPE sets RECORD_TYPE field to given value.


### GetAUTHORIZATION_TYPE

`func (o *PropagationStoreConfiguration) GetAUTHORIZATION_TYPE() string`

GetAUTHORIZATION_TYPE returns the AUTHORIZATION_TYPE field if non-nil, zero value otherwise.

### GetAUTHORIZATION_TYPEOk

`func (o *PropagationStoreConfiguration) GetAUTHORIZATION_TYPEOk() (*string, bool)`

GetAUTHORIZATION_TYPEOk returns a tuple with the AUTHORIZATION_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTHORIZATION_TYPE

`func (o *PropagationStoreConfiguration) SetAUTHORIZATION_TYPE(v string)`

SetAUTHORIZATION_TYPE sets AUTHORIZATION_TYPE field to given value.


### GetGROUPS_RESOURCE

`func (o *PropagationStoreConfiguration) GetGROUPS_RESOURCE() string`

GetGROUPS_RESOURCE returns the GROUPS_RESOURCE field if non-nil, zero value otherwise.

### GetGROUPS_RESOURCEOk

`func (o *PropagationStoreConfiguration) GetGROUPS_RESOURCEOk() (*string, bool)`

GetGROUPS_RESOURCEOk returns a tuple with the GROUPS_RESOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGROUPS_RESOURCE

`func (o *PropagationStoreConfiguration) SetGROUPS_RESOURCE(v string)`

SetGROUPS_RESOURCE sets GROUPS_RESOURCE field to given value.

### HasGROUPS_RESOURCE

`func (o *PropagationStoreConfiguration) HasGROUPS_RESOURCE() bool`

HasGROUPS_RESOURCE returns a boolean if a field has been set.

### GetOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfiguration) GetOAUTH_TOKEN_REQUEST() string`

GetOAUTH_TOKEN_REQUEST returns the OAUTH_TOKEN_REQUEST field if non-nil, zero value otherwise.

### GetOAUTH_TOKEN_REQUESTOk

`func (o *PropagationStoreConfiguration) GetOAUTH_TOKEN_REQUESTOk() (*string, bool)`

GetOAUTH_TOKEN_REQUESTOk returns a tuple with the OAUTH_TOKEN_REQUEST field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfiguration) SetOAUTH_TOKEN_REQUEST(v string)`

SetOAUTH_TOKEN_REQUEST sets OAUTH_TOKEN_REQUEST field to given value.

### HasOAUTH_TOKEN_REQUEST

`func (o *PropagationStoreConfiguration) HasOAUTH_TOKEN_REQUEST() bool`

HasOAUTH_TOKEN_REQUEST returns a boolean if a field has been set.

### GetSCIM_VERSION

`func (o *PropagationStoreConfiguration) GetSCIM_VERSION() string`

GetSCIM_VERSION returns the SCIM_VERSION field if non-nil, zero value otherwise.

### GetSCIM_VERSIONOk

`func (o *PropagationStoreConfiguration) GetSCIM_VERSIONOk() (*string, bool)`

GetSCIM_VERSIONOk returns a tuple with the SCIM_VERSION field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSCIM_VERSION

`func (o *PropagationStoreConfiguration) SetSCIM_VERSION(v string)`

SetSCIM_VERSION sets SCIM_VERSION field to given value.


### GetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfiguration) GetUNIQUE_USER_IDENTIFIER() EnumPropagationStoreTypeSlackUniqueUserIdentifier`

GetUNIQUE_USER_IDENTIFIER returns the UNIQUE_USER_IDENTIFIER field if non-nil, zero value otherwise.

### GetUNIQUE_USER_IDENTIFIEROk

`func (o *PropagationStoreConfiguration) GetUNIQUE_USER_IDENTIFIEROk() (*EnumPropagationStoreTypeSlackUniqueUserIdentifier, bool)`

GetUNIQUE_USER_IDENTIFIEROk returns a tuple with the UNIQUE_USER_IDENTIFIER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNIQUE_USER_IDENTIFIER

`func (o *PropagationStoreConfiguration) SetUNIQUE_USER_IDENTIFIER(v EnumPropagationStoreTypeSlackUniqueUserIdentifier)`

SetUNIQUE_USER_IDENTIFIER sets UNIQUE_USER_IDENTIFIER field to given value.


### GetUSER_FILTER

`func (o *PropagationStoreConfiguration) GetUSER_FILTER() string`

GetUSER_FILTER returns the USER_FILTER field if non-nil, zero value otherwise.

### GetUSER_FILTEROk

`func (o *PropagationStoreConfiguration) GetUSER_FILTEROk() (*string, bool)`

GetUSER_FILTEROk returns a tuple with the USER_FILTER field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_FILTER

`func (o *PropagationStoreConfiguration) SetUSER_FILTER(v string)`

SetUSER_FILTER sets USER_FILTER field to given value.


### GetUSERS_RESOURCE

`func (o *PropagationStoreConfiguration) GetUSERS_RESOURCE() string`

GetUSERS_RESOURCE returns the USERS_RESOURCE field if non-nil, zero value otherwise.

### GetUSERS_RESOURCEOk

`func (o *PropagationStoreConfiguration) GetUSERS_RESOURCEOk() (*string, bool)`

GetUSERS_RESOURCEOk returns a tuple with the USERS_RESOURCE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSERS_RESOURCE

`func (o *PropagationStoreConfiguration) SetUSERS_RESOURCE(v string)`

SetUSERS_RESOURCE sets USERS_RESOURCE field to given value.


### GetAdministratorPassword

`func (o *PropagationStoreConfiguration) GetAdministratorPassword() string`

GetAdministratorPassword returns the AdministratorPassword field if non-nil, zero value otherwise.

### GetAdministratorPasswordOk

`func (o *PropagationStoreConfiguration) GetAdministratorPasswordOk() (*string, bool)`

GetAdministratorPasswordOk returns a tuple with the AdministratorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorPassword

`func (o *PropagationStoreConfiguration) SetAdministratorPassword(v string)`

SetAdministratorPassword sets AdministratorPassword field to given value.


### GetAdministratorUsername

`func (o *PropagationStoreConfiguration) GetAdministratorUsername() string`

GetAdministratorUsername returns the AdministratorUsername field if non-nil, zero value otherwise.

### GetAdministratorUsernameOk

`func (o *PropagationStoreConfiguration) GetAdministratorUsernameOk() (*string, bool)`

GetAdministratorUsernameOk returns a tuple with the AdministratorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorUsername

`func (o *PropagationStoreConfiguration) SetAdministratorUsername(v string)`

SetAdministratorUsername sets AdministratorUsername field to given value.


### GetServiceNowUrl

`func (o *PropagationStoreConfiguration) GetServiceNowUrl() string`

GetServiceNowUrl returns the ServiceNowUrl field if non-nil, zero value otherwise.

### GetServiceNowUrlOk

`func (o *PropagationStoreConfiguration) GetServiceNowUrlOk() (*string, bool)`

GetServiceNowUrlOk returns a tuple with the ServiceNowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNowUrl

`func (o *PropagationStoreConfiguration) SetServiceNowUrl(v string)`

SetServiceNowUrl sets ServiceNowUrl field to given value.


### GetExcludeContingentWorkers

`func (o *PropagationStoreConfiguration) GetExcludeContingentWorkers() bool`

GetExcludeContingentWorkers returns the ExcludeContingentWorkers field if non-nil, zero value otherwise.

### GetExcludeContingentWorkersOk

`func (o *PropagationStoreConfiguration) GetExcludeContingentWorkersOk() (*bool, bool)`

GetExcludeContingentWorkersOk returns a tuple with the ExcludeContingentWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContingentWorkers

`func (o *PropagationStoreConfiguration) SetExcludeContingentWorkers(v bool)`

SetExcludeContingentWorkers sets ExcludeContingentWorkers field to given value.

### HasExcludeContingentWorkers

`func (o *PropagationStoreConfiguration) HasExcludeContingentWorkers() bool`

HasExcludeContingentWorkers returns a boolean if a field has been set.

### GetExcludeEmployees

`func (o *PropagationStoreConfiguration) GetExcludeEmployees() bool`

GetExcludeEmployees returns the ExcludeEmployees field if non-nil, zero value otherwise.

### GetExcludeEmployeesOk

`func (o *PropagationStoreConfiguration) GetExcludeEmployeesOk() (*bool, bool)`

GetExcludeEmployeesOk returns a tuple with the ExcludeEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeEmployees

`func (o *PropagationStoreConfiguration) SetExcludeEmployees(v bool)`

SetExcludeEmployees sets ExcludeEmployees field to given value.

### HasExcludeEmployees

`func (o *PropagationStoreConfiguration) HasExcludeEmployees() bool`

HasExcludeEmployees returns a boolean if a field has been set.

### GetExcludeInactiveWorkers

`func (o *PropagationStoreConfiguration) GetExcludeInactiveWorkers() bool`

GetExcludeInactiveWorkers returns the ExcludeInactiveWorkers field if non-nil, zero value otherwise.

### GetExcludeInactiveWorkersOk

`func (o *PropagationStoreConfiguration) GetExcludeInactiveWorkersOk() (*bool, bool)`

GetExcludeInactiveWorkersOk returns a tuple with the ExcludeInactiveWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeInactiveWorkers

`func (o *PropagationStoreConfiguration) SetExcludeInactiveWorkers(v bool)`

SetExcludeInactiveWorkers sets ExcludeInactiveWorkers field to given value.

### HasExcludeInactiveWorkers

`func (o *PropagationStoreConfiguration) HasExcludeInactiveWorkers() bool`

HasExcludeInactiveWorkers returns a boolean if a field has been set.

### GetHost

`func (o *PropagationStoreConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PropagationStoreConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PropagationStoreConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PropagationStoreConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPassword

`func (o *PropagationStoreConfiguration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PropagationStoreConfiguration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PropagationStoreConfiguration) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetTenantId

`func (o *PropagationStoreConfiguration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *PropagationStoreConfiguration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *PropagationStoreConfiguration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUsername

`func (o *PropagationStoreConfiguration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PropagationStoreConfiguration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PropagationStoreConfiguration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAPI_KEY

`func (o *PropagationStoreConfiguration) GetAPI_KEY() string`

GetAPI_KEY returns the API_KEY field if non-nil, zero value otherwise.

### GetAPI_KEYOk

`func (o *PropagationStoreConfiguration) GetAPI_KEYOk() (*string, bool)`

GetAPI_KEYOk returns a tuple with the API_KEY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_KEY

`func (o *PropagationStoreConfiguration) SetAPI_KEY(v string)`

SetAPI_KEY sets API_KEY field to given value.

### HasAPI_KEY

`func (o *PropagationStoreConfiguration) HasAPI_KEY() bool`

HasAPI_KEY returns a boolean if a field has been set.

### GetAPI_SECRET

`func (o *PropagationStoreConfiguration) GetAPI_SECRET() string`

GetAPI_SECRET returns the API_SECRET field if non-nil, zero value otherwise.

### GetAPI_SECRETOk

`func (o *PropagationStoreConfiguration) GetAPI_SECRETOk() (*string, bool)`

GetAPI_SECRETOk returns a tuple with the API_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPI_SECRET

`func (o *PropagationStoreConfiguration) SetAPI_SECRET(v string)`

SetAPI_SECRET sets API_SECRET field to given value.

### HasAPI_SECRET

`func (o *PropagationStoreConfiguration) HasAPI_SECRET() bool`

HasAPI_SECRET returns a boolean if a field has been set.

### GetOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfiguration) GetOAUTH_ACCOUNT_ID() string`

GetOAUTH_ACCOUNT_ID returns the OAUTH_ACCOUNT_ID field if non-nil, zero value otherwise.

### GetOAUTH_ACCOUNT_IDOk

`func (o *PropagationStoreConfiguration) GetOAUTH_ACCOUNT_IDOk() (*string, bool)`

GetOAUTH_ACCOUNT_IDOk returns a tuple with the OAUTH_ACCOUNT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfiguration) SetOAUTH_ACCOUNT_ID(v string)`

SetOAUTH_ACCOUNT_ID sets OAUTH_ACCOUNT_ID field to given value.

### HasOAUTH_ACCOUNT_ID

`func (o *PropagationStoreConfiguration) HasOAUTH_ACCOUNT_ID() bool`

HasOAUTH_ACCOUNT_ID returns a boolean if a field has been set.

### GetOAUTH_TOKEN_URL

`func (o *PropagationStoreConfiguration) GetOAUTH_TOKEN_URL() string`

GetOAUTH_TOKEN_URL returns the OAUTH_TOKEN_URL field if non-nil, zero value otherwise.

### GetOAUTH_TOKEN_URLOk

`func (o *PropagationStoreConfiguration) GetOAUTH_TOKEN_URLOk() (*string, bool)`

GetOAUTH_TOKEN_URLOk returns a tuple with the OAUTH_TOKEN_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_TOKEN_URL

`func (o *PropagationStoreConfiguration) SetOAUTH_TOKEN_URL(v string)`

SetOAUTH_TOKEN_URL sets OAUTH_TOKEN_URL field to given value.

### HasOAUTH_TOKEN_URL

`func (o *PropagationStoreConfiguration) HasOAUTH_TOKEN_URL() bool`

HasOAUTH_TOKEN_URL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


