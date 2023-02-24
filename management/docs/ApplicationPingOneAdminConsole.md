# ApplicationPingOneAdminConsole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**TokenEndpointAuthMethod** | Pointer to [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | [optional] 

## Methods

### NewApplicationPingOneAdminConsole

`func NewApplicationPingOneAdminConsole() *ApplicationPingOneAdminConsole`

NewApplicationPingOneAdminConsole instantiates a new ApplicationPingOneAdminConsole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPingOneAdminConsoleWithDefaults

`func NewApplicationPingOneAdminConsoleWithDefaults() *ApplicationPingOneAdminConsole`

NewApplicationPingOneAdminConsoleWithDefaults instantiates a new ApplicationPingOneAdminConsole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkceEnforcement

`func (o *ApplicationPingOneAdminConsole) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *ApplicationPingOneAdminConsole) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *ApplicationPingOneAdminConsole) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *ApplicationPingOneAdminConsole) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationPingOneAdminConsole) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationPingOneAdminConsole) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationPingOneAdminConsole) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *ApplicationPingOneAdminConsole) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


