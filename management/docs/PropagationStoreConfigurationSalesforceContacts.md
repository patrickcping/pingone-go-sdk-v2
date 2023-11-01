# PropagationStoreConfigurationSalesforceContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ACCOUNT_ID** | Pointer to **string** | The Salesforce account ID. | [optional] 
**CLIENT_ID** | **string** | Unique identifier (UUID) of the Salesforce client associated with the propagation store. | 
**CLIENT_SECRET** | **string** | The Salesforce client secret. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DISABLE_USERS** | Pointer to **bool** | Whether or not users are allowed to be disabled. | [optional] 
**ENABLE_COMMUNITIES** | Pointer to **bool** | Whether or not to enable Salesforce communities. | [optional] 
**FREEZE_USER_FLAG** | Pointer to **bool** | Whether the user account is frozen. | [optional] 
**OAUTH_ACCESS_TOKEN** | **string** | OAuth access token for account authentication. | 
**OAUTH_REFRESH_TOKEN** | **string** | OAuth refresh token for account authentication. | 
**PERMISSION_SET_MANAGEMENT** | Pointer to **string** | The permission sets to be merged with Salesforce. | [optional] 
**PROFILE_ID** | Pointer to **string** | The Salesforce profile ID. | [optional] 
**PROVISION_DISABLED_USERS** | Pointer to **bool** | Whether or not disabled users can be provisioned. | [optional] 
**RECORD_TYPE** | [**EnumPropagationStoreTypeSalesforceContactsRecordType**](EnumPropagationStoreTypeSalesforceContactsRecordType.md) |  | 
**SALESFORCE_DOMAIN** | **string** | The account&#39;s salesforce.com domain. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationSalesforceContacts

`func NewPropagationStoreConfigurationSalesforceContacts(cLIENTID string, cLIENTSECRET string, oAUTHACCESSTOKEN string, oAUTHREFRESHTOKEN string, rECORDTYPE EnumPropagationStoreTypeSalesforceContactsRecordType, sALESFORCEDOMAIN string, ) *PropagationStoreConfigurationSalesforceContacts`

NewPropagationStoreConfigurationSalesforceContacts instantiates a new PropagationStoreConfigurationSalesforceContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationSalesforceContactsWithDefaults

`func NewPropagationStoreConfigurationSalesforceContactsWithDefaults() *PropagationStoreConfigurationSalesforceContacts`

NewPropagationStoreConfigurationSalesforceContactsWithDefaults instantiates a new PropagationStoreConfigurationSalesforceContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetACCOUNT_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) GetACCOUNT_ID() string`

GetACCOUNT_ID returns the ACCOUNT_ID field if non-nil, zero value otherwise.

### GetACCOUNT_IDOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetACCOUNT_IDOk() (*string, bool)`

GetACCOUNT_IDOk returns a tuple with the ACCOUNT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetACCOUNT_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) SetACCOUNT_ID(v string)`

SetACCOUNT_ID sets ACCOUNT_ID field to given value.

### HasACCOUNT_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) HasACCOUNT_ID() bool`

HasACCOUNT_ID returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *PropagationStoreConfigurationSalesforceContacts) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDISABLE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) GetDISABLE_USERS() bool`

GetDISABLE_USERS returns the DISABLE_USERS field if non-nil, zero value otherwise.

### GetDISABLE_USERSOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetDISABLE_USERSOk() (*bool, bool)`

GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDISABLE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) SetDISABLE_USERS(v bool)`

SetDISABLE_USERS sets DISABLE_USERS field to given value.

### HasDISABLE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) HasDISABLE_USERS() bool`

HasDISABLE_USERS returns a boolean if a field has been set.

### GetENABLE_COMMUNITIES

`func (o *PropagationStoreConfigurationSalesforceContacts) GetENABLE_COMMUNITIES() bool`

GetENABLE_COMMUNITIES returns the ENABLE_COMMUNITIES field if non-nil, zero value otherwise.

### GetENABLE_COMMUNITIESOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetENABLE_COMMUNITIESOk() (*bool, bool)`

GetENABLE_COMMUNITIESOk returns a tuple with the ENABLE_COMMUNITIES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENABLE_COMMUNITIES

`func (o *PropagationStoreConfigurationSalesforceContacts) SetENABLE_COMMUNITIES(v bool)`

SetENABLE_COMMUNITIES sets ENABLE_COMMUNITIES field to given value.

### HasENABLE_COMMUNITIES

`func (o *PropagationStoreConfigurationSalesforceContacts) HasENABLE_COMMUNITIES() bool`

HasENABLE_COMMUNITIES returns a boolean if a field has been set.

### GetFREEZE_USER_FLAG

`func (o *PropagationStoreConfigurationSalesforceContacts) GetFREEZE_USER_FLAG() bool`

GetFREEZE_USER_FLAG returns the FREEZE_USER_FLAG field if non-nil, zero value otherwise.

### GetFREEZE_USER_FLAGOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetFREEZE_USER_FLAGOk() (*bool, bool)`

GetFREEZE_USER_FLAGOk returns a tuple with the FREEZE_USER_FLAG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFREEZE_USER_FLAG

`func (o *PropagationStoreConfigurationSalesforceContacts) SetFREEZE_USER_FLAG(v bool)`

SetFREEZE_USER_FLAG sets FREEZE_USER_FLAG field to given value.

### HasFREEZE_USER_FLAG

`func (o *PropagationStoreConfigurationSalesforceContacts) HasFREEZE_USER_FLAG() bool`

HasFREEZE_USER_FLAG returns a boolean if a field has been set.

### GetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSalesforceContacts) GetOAUTH_ACCESS_TOKEN() string`

GetOAUTH_ACCESS_TOKEN returns the OAUTH_ACCESS_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_ACCESS_TOKENOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetOAUTH_ACCESS_TOKENOk() (*string, bool)`

GetOAUTH_ACCESS_TOKENOk returns a tuple with the OAUTH_ACCESS_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_ACCESS_TOKEN

`func (o *PropagationStoreConfigurationSalesforceContacts) SetOAUTH_ACCESS_TOKEN(v string)`

SetOAUTH_ACCESS_TOKEN sets OAUTH_ACCESS_TOKEN field to given value.


### GetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfigurationSalesforceContacts) GetOAUTH_REFRESH_TOKEN() string`

GetOAUTH_REFRESH_TOKEN returns the OAUTH_REFRESH_TOKEN field if non-nil, zero value otherwise.

### GetOAUTH_REFRESH_TOKENOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetOAUTH_REFRESH_TOKENOk() (*string, bool)`

GetOAUTH_REFRESH_TOKENOk returns a tuple with the OAUTH_REFRESH_TOKEN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_REFRESH_TOKEN

`func (o *PropagationStoreConfigurationSalesforceContacts) SetOAUTH_REFRESH_TOKEN(v string)`

SetOAUTH_REFRESH_TOKEN sets OAUTH_REFRESH_TOKEN field to given value.


### GetPERMISSION_SET_MANAGEMENT

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPERMISSION_SET_MANAGEMENT() string`

GetPERMISSION_SET_MANAGEMENT returns the PERMISSION_SET_MANAGEMENT field if non-nil, zero value otherwise.

### GetPERMISSION_SET_MANAGEMENTOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPERMISSION_SET_MANAGEMENTOk() (*string, bool)`

GetPERMISSION_SET_MANAGEMENTOk returns a tuple with the PERMISSION_SET_MANAGEMENT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPERMISSION_SET_MANAGEMENT

`func (o *PropagationStoreConfigurationSalesforceContacts) SetPERMISSION_SET_MANAGEMENT(v string)`

SetPERMISSION_SET_MANAGEMENT sets PERMISSION_SET_MANAGEMENT field to given value.

### HasPERMISSION_SET_MANAGEMENT

`func (o *PropagationStoreConfigurationSalesforceContacts) HasPERMISSION_SET_MANAGEMENT() bool`

HasPERMISSION_SET_MANAGEMENT returns a boolean if a field has been set.

### GetPROFILE_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPROFILE_ID() string`

GetPROFILE_ID returns the PROFILE_ID field if non-nil, zero value otherwise.

### GetPROFILE_IDOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPROFILE_IDOk() (*string, bool)`

GetPROFILE_IDOk returns a tuple with the PROFILE_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROFILE_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) SetPROFILE_ID(v string)`

SetPROFILE_ID sets PROFILE_ID field to given value.

### HasPROFILE_ID

`func (o *PropagationStoreConfigurationSalesforceContacts) HasPROFILE_ID() bool`

HasPROFILE_ID returns a boolean if a field has been set.

### GetPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPROVISION_DISABLED_USERS() bool`

GetPROVISION_DISABLED_USERS returns the PROVISION_DISABLED_USERS field if non-nil, zero value otherwise.

### GetPROVISION_DISABLED_USERSOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetPROVISION_DISABLED_USERSOk() (*bool, bool)`

GetPROVISION_DISABLED_USERSOk returns a tuple with the PROVISION_DISABLED_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) SetPROVISION_DISABLED_USERS(v bool)`

SetPROVISION_DISABLED_USERS sets PROVISION_DISABLED_USERS field to given value.

### HasPROVISION_DISABLED_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) HasPROVISION_DISABLED_USERS() bool`

HasPROVISION_DISABLED_USERS returns a boolean if a field has been set.

### GetRECORD_TYPE

`func (o *PropagationStoreConfigurationSalesforceContacts) GetRECORD_TYPE() EnumPropagationStoreTypeSalesforceContactsRecordType`

GetRECORD_TYPE returns the RECORD_TYPE field if non-nil, zero value otherwise.

### GetRECORD_TYPEOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetRECORD_TYPEOk() (*EnumPropagationStoreTypeSalesforceContactsRecordType, bool)`

GetRECORD_TYPEOk returns a tuple with the RECORD_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRECORD_TYPE

`func (o *PropagationStoreConfigurationSalesforceContacts) SetRECORD_TYPE(v EnumPropagationStoreTypeSalesforceContactsRecordType)`

SetRECORD_TYPE sets RECORD_TYPE field to given value.


### GetSALESFORCE_DOMAIN

`func (o *PropagationStoreConfigurationSalesforceContacts) GetSALESFORCE_DOMAIN() string`

GetSALESFORCE_DOMAIN returns the SALESFORCE_DOMAIN field if non-nil, zero value otherwise.

### GetSALESFORCE_DOMAINOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetSALESFORCE_DOMAINOk() (*string, bool)`

GetSALESFORCE_DOMAINOk returns a tuple with the SALESFORCE_DOMAIN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSALESFORCE_DOMAIN

`func (o *PropagationStoreConfigurationSalesforceContacts) SetSALESFORCE_DOMAIN(v string)`

SetSALESFORCE_DOMAIN sets SALESFORCE_DOMAIN field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationSalesforceContacts) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationSalesforceContacts) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


