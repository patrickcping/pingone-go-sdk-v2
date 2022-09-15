# ApplicationOIDCInternalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 

## Methods

### NewApplicationOIDCInternalAllOf

`func NewApplicationOIDCInternalAllOf(tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, ) *ApplicationOIDCInternalAllOf`

NewApplicationOIDCInternalAllOf instantiates a new ApplicationOIDCInternalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCInternalAllOfWithDefaults

`func NewApplicationOIDCInternalAllOfWithDefaults() *ApplicationOIDCInternalAllOf`

NewApplicationOIDCInternalAllOfWithDefaults instantiates a new ApplicationOIDCInternalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPkceEnforcement

`func (o *ApplicationOIDCInternalAllOf) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *ApplicationOIDCInternalAllOf) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *ApplicationOIDCInternalAllOf) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *ApplicationOIDCInternalAllOf) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationOIDCInternalAllOf) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationOIDCInternalAllOf) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationOIDCInternalAllOf) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


