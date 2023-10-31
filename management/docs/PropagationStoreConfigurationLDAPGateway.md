# PropagationStoreConfigurationLDAPGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ATTRIBUTE_METADATA** | Pointer to **string** | User-defined attribute metadata. | [optional] 
**CLIENT_ID** | **string** | Identifier of the client for authentication. | 
**CLIENT_SECRET** | **string** | Secret of the client for authentication. | 
**CREATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be created. | [optional] 
**DELETE_USERS** | Pointer to **bool** | Whether or not users are allowed to be deleted. | [optional] 
**ENVIRONMENT_ID** | **string** | Identifier, a UUID, of the environment the connector services. | 
**GATEWAY_BASE_URL** | **string** | Base URL of the gateway. | 
**GATEWAY_ID** | **string** | Identifier of the gateway to which the connector connects. | 
**LDAP_TYPE** | [**EnumPropagationStoreTypeLDAPGatewayLDAPType**](EnumPropagationStoreTypeLDAPGatewayLDAPType.md) |  | 
**OAUTH_URL** | **string** | OAuth token request endpoint. | 
**UPDATE_USERS** | Pointer to **bool** | Whether or not users are allowed to be updated. | [optional] 

## Methods

### NewPropagationStoreConfigurationLDAPGateway

`func NewPropagationStoreConfigurationLDAPGateway(cLIENTID string, cLIENTSECRET string, eNVIRONMENTID string, gATEWAYBASEURL string, gATEWAYID string, lDAPTYPE EnumPropagationStoreTypeLDAPGatewayLDAPType, oAUTHURL string, ) *PropagationStoreConfigurationLDAPGateway`

NewPropagationStoreConfigurationLDAPGateway instantiates a new PropagationStoreConfigurationLDAPGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropagationStoreConfigurationLDAPGatewayWithDefaults

`func NewPropagationStoreConfigurationLDAPGatewayWithDefaults() *PropagationStoreConfigurationLDAPGateway`

NewPropagationStoreConfigurationLDAPGatewayWithDefaults instantiates a new PropagationStoreConfigurationLDAPGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetATTRIBUTE_METADATA

`func (o *PropagationStoreConfigurationLDAPGateway) GetATTRIBUTE_METADATA() string`

GetATTRIBUTE_METADATA returns the ATTRIBUTE_METADATA field if non-nil, zero value otherwise.

### GetATTRIBUTE_METADATAOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetATTRIBUTE_METADATAOk() (*string, bool)`

GetATTRIBUTE_METADATAOk returns a tuple with the ATTRIBUTE_METADATA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetATTRIBUTE_METADATA

`func (o *PropagationStoreConfigurationLDAPGateway) SetATTRIBUTE_METADATA(v string)`

SetATTRIBUTE_METADATA sets ATTRIBUTE_METADATA field to given value.

### HasATTRIBUTE_METADATA

`func (o *PropagationStoreConfigurationLDAPGateway) HasATTRIBUTE_METADATA() bool`

HasATTRIBUTE_METADATA returns a boolean if a field has been set.

### GetCLIENT_ID

`func (o *PropagationStoreConfigurationLDAPGateway) GetCLIENT_ID() string`

GetCLIENT_ID returns the CLIENT_ID field if non-nil, zero value otherwise.

### GetCLIENT_IDOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetCLIENT_IDOk() (*string, bool)`

GetCLIENT_IDOk returns a tuple with the CLIENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_ID

`func (o *PropagationStoreConfigurationLDAPGateway) SetCLIENT_ID(v string)`

SetCLIENT_ID sets CLIENT_ID field to given value.


### GetCLIENT_SECRET

`func (o *PropagationStoreConfigurationLDAPGateway) GetCLIENT_SECRET() string`

GetCLIENT_SECRET returns the CLIENT_SECRET field if non-nil, zero value otherwise.

### GetCLIENT_SECRETOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetCLIENT_SECRETOk() (*string, bool)`

GetCLIENT_SECRETOk returns a tuple with the CLIENT_SECRET field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCLIENT_SECRET

`func (o *PropagationStoreConfigurationLDAPGateway) SetCLIENT_SECRET(v string)`

SetCLIENT_SECRET sets CLIENT_SECRET field to given value.


### GetCREATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) GetCREATE_USERS() bool`

GetCREATE_USERS returns the CREATE_USERS field if non-nil, zero value otherwise.

### GetCREATE_USERSOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetCREATE_USERSOk() (*bool, bool)`

GetCREATE_USERSOk returns a tuple with the CREATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCREATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) SetCREATE_USERS(v bool)`

SetCREATE_USERS sets CREATE_USERS field to given value.

### HasCREATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) HasCREATE_USERS() bool`

HasCREATE_USERS returns a boolean if a field has been set.

### GetDELETE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) GetDELETE_USERS() bool`

GetDELETE_USERS returns the DELETE_USERS field if non-nil, zero value otherwise.

### GetDELETE_USERSOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetDELETE_USERSOk() (*bool, bool)`

GetDELETE_USERSOk returns a tuple with the DELETE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDELETE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) SetDELETE_USERS(v bool)`

SetDELETE_USERS sets DELETE_USERS field to given value.

### HasDELETE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) HasDELETE_USERS() bool`

HasDELETE_USERS returns a boolean if a field has been set.

### GetENVIRONMENT_ID

`func (o *PropagationStoreConfigurationLDAPGateway) GetENVIRONMENT_ID() string`

GetENVIRONMENT_ID returns the ENVIRONMENT_ID field if non-nil, zero value otherwise.

### GetENVIRONMENT_IDOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetENVIRONMENT_IDOk() (*string, bool)`

GetENVIRONMENT_IDOk returns a tuple with the ENVIRONMENT_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetENVIRONMENT_ID

`func (o *PropagationStoreConfigurationLDAPGateway) SetENVIRONMENT_ID(v string)`

SetENVIRONMENT_ID sets ENVIRONMENT_ID field to given value.


### GetGATEWAY_BASE_URL

`func (o *PropagationStoreConfigurationLDAPGateway) GetGATEWAY_BASE_URL() string`

GetGATEWAY_BASE_URL returns the GATEWAY_BASE_URL field if non-nil, zero value otherwise.

### GetGATEWAY_BASE_URLOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetGATEWAY_BASE_URLOk() (*string, bool)`

GetGATEWAY_BASE_URLOk returns a tuple with the GATEWAY_BASE_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGATEWAY_BASE_URL

`func (o *PropagationStoreConfigurationLDAPGateway) SetGATEWAY_BASE_URL(v string)`

SetGATEWAY_BASE_URL sets GATEWAY_BASE_URL field to given value.


### GetGATEWAY_ID

`func (o *PropagationStoreConfigurationLDAPGateway) GetGATEWAY_ID() string`

GetGATEWAY_ID returns the GATEWAY_ID field if non-nil, zero value otherwise.

### GetGATEWAY_IDOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetGATEWAY_IDOk() (*string, bool)`

GetGATEWAY_IDOk returns a tuple with the GATEWAY_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGATEWAY_ID

`func (o *PropagationStoreConfigurationLDAPGateway) SetGATEWAY_ID(v string)`

SetGATEWAY_ID sets GATEWAY_ID field to given value.


### GetLDAP_TYPE

`func (o *PropagationStoreConfigurationLDAPGateway) GetLDAP_TYPE() EnumPropagationStoreTypeLDAPGatewayLDAPType`

GetLDAP_TYPE returns the LDAP_TYPE field if non-nil, zero value otherwise.

### GetLDAP_TYPEOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetLDAP_TYPEOk() (*EnumPropagationStoreTypeLDAPGatewayLDAPType, bool)`

GetLDAP_TYPEOk returns a tuple with the LDAP_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLDAP_TYPE

`func (o *PropagationStoreConfigurationLDAPGateway) SetLDAP_TYPE(v EnumPropagationStoreTypeLDAPGatewayLDAPType)`

SetLDAP_TYPE sets LDAP_TYPE field to given value.


### GetOAUTH_URL

`func (o *PropagationStoreConfigurationLDAPGateway) GetOAUTH_URL() string`

GetOAUTH_URL returns the OAUTH_URL field if non-nil, zero value otherwise.

### GetOAUTH_URLOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetOAUTH_URLOk() (*string, bool)`

GetOAUTH_URLOk returns a tuple with the OAUTH_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAUTH_URL

`func (o *PropagationStoreConfigurationLDAPGateway) SetOAUTH_URL(v string)`

SetOAUTH_URL sets OAUTH_URL field to given value.


### GetUPDATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) GetUPDATE_USERS() bool`

GetUPDATE_USERS returns the UPDATE_USERS field if non-nil, zero value otherwise.

### GetUPDATE_USERSOk

`func (o *PropagationStoreConfigurationLDAPGateway) GetUPDATE_USERSOk() (*bool, bool)`

GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUPDATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) SetUPDATE_USERS(v bool)`

SetUPDATE_USERS sets UPDATE_USERS field to given value.

### HasUPDATE_USERS

`func (o *PropagationStoreConfigurationLDAPGateway) HasUPDATE_USERS() bool`

HasUPDATE_USERS returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


