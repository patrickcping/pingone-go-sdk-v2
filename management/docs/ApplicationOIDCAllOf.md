# ApplicationOIDCAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mobile** | Pointer to [**ApplicationOIDCAllOfMobile**](ApplicationOIDCAllOfMobile.md) |  | [optional] 
**BundleId** | Pointer to **string** | A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable. | [optional] 
**PackageName** | Pointer to **string** | A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable. | [optional] 
**GrantTypes** | [**[]EnumApplicationOIDCGrantType**](EnumApplicationOIDCGrantType.md) | A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS. | 
**HomePageUrl** | Pointer to **string** | A string that specifies the custom home page URL for the application. | [optional] 
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | A string that specifies the URLs that the browser can be redirected to after logout. | [optional] 
**RedirectUris** | Pointer to **[]string** | A string that specifies the callback URI for the authentication response. | [optional] 
**RefreshTokenDuration** | Pointer to **int32** | An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the refreshTokenRollingDuration property is specified for the application, then this property must be less than or equal to the value of refreshTokenRollingDuration. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] 
**RefreshTokenRollingDuration** | Pointer to **int32** | An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] 
**ResponseTypes** | Pointer to [**[]EnumApplicationOIDCResponseType**](EnumApplicationOIDCResponseType.md) | A string that specifies the code or token type returned by an authorization request. Options are TOKEN, ID_TOKEN, and CODE. Note that CODE cannot be used in an authorization request with TOKEN or ID_TOKEN because PingOne does not currently support OIDC hybrid flows. | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 

## Methods

### NewApplicationOIDCAllOf

`func NewApplicationOIDCAllOf(grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, ) *ApplicationOIDCAllOf`

NewApplicationOIDCAllOf instantiates a new ApplicationOIDCAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOIDCAllOfWithDefaults

`func NewApplicationOIDCAllOfWithDefaults() *ApplicationOIDCAllOf`

NewApplicationOIDCAllOfWithDefaults instantiates a new ApplicationOIDCAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobile

`func (o *ApplicationOIDCAllOf) GetMobile() ApplicationOIDCAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *ApplicationOIDCAllOf) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *ApplicationOIDCAllOf) SetMobile(v ApplicationOIDCAllOfMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *ApplicationOIDCAllOf) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetBundleId

`func (o *ApplicationOIDCAllOf) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *ApplicationOIDCAllOf) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *ApplicationOIDCAllOf) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *ApplicationOIDCAllOf) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *ApplicationOIDCAllOf) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ApplicationOIDCAllOf) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ApplicationOIDCAllOf) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ApplicationOIDCAllOf) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetGrantTypes

`func (o *ApplicationOIDCAllOf) GetGrantTypes() []EnumApplicationOIDCGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *ApplicationOIDCAllOf) GetGrantTypesOk() (*[]EnumApplicationOIDCGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *ApplicationOIDCAllOf) SetGrantTypes(v []EnumApplicationOIDCGrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetHomePageUrl

`func (o *ApplicationOIDCAllOf) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ApplicationOIDCAllOf) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ApplicationOIDCAllOf) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *ApplicationOIDCAllOf) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### GetPkceEnforcement

`func (o *ApplicationOIDCAllOf) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *ApplicationOIDCAllOf) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *ApplicationOIDCAllOf) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *ApplicationOIDCAllOf) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *ApplicationOIDCAllOf) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *ApplicationOIDCAllOf) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *ApplicationOIDCAllOf) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *ApplicationOIDCAllOf) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *ApplicationOIDCAllOf) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ApplicationOIDCAllOf) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ApplicationOIDCAllOf) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *ApplicationOIDCAllOf) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *ApplicationOIDCAllOf) GetRefreshTokenDuration() int32`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *ApplicationOIDCAllOf) GetRefreshTokenDurationOk() (*int32, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *ApplicationOIDCAllOf) SetRefreshTokenDuration(v int32)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *ApplicationOIDCAllOf) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingDuration

`func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingDuration() int32`

GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingDurationOk

`func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingDurationOk() (*int32, bool)`

GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingDuration

`func (o *ApplicationOIDCAllOf) SetRefreshTokenRollingDuration(v int32)`

SetRefreshTokenRollingDuration sets RefreshTokenRollingDuration field to given value.

### HasRefreshTokenRollingDuration

`func (o *ApplicationOIDCAllOf) HasRefreshTokenRollingDuration() bool`

HasRefreshTokenRollingDuration returns a boolean if a field has been set.

### GetResponseTypes

`func (o *ApplicationOIDCAllOf) GetResponseTypes() []EnumApplicationOIDCResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *ApplicationOIDCAllOf) GetResponseTypesOk() (*[]EnumApplicationOIDCResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *ApplicationOIDCAllOf) SetResponseTypes(v []EnumApplicationOIDCResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *ApplicationOIDCAllOf) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationOIDCAllOf) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationOIDCAllOf) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationOIDCAllOf) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


