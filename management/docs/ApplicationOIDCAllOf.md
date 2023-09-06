# ApplicationOIDCAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalRefreshTokenReplayProtectionEnabled** | Pointer to **bool** | When set to &#x60;true&#x60; (the default), if you attempt to reuse the refresh token, the authorization server immediately revokes the reused refresh token, as well as all descendant tokens. Setting this to null equates to a &#x60;false&#x60; setting. | [optional] [default to true]
**AllowWildcardInRedirectUris** | Pointer to **bool** | A boolean to specify whether wildcards are allowed in redirect URIs. For more information, see [Wildcards in Redirect URIs](https://docs.pingidentity.com/csh?context&#x3D;p1_c_wildcard_redirect_uri). | [optional] 
**AssignActorRoles** | Pointer to **bool** | A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request. | [optional] 
**Mobile** | Pointer to [**ApplicationOIDCAllOfMobile**](ApplicationOIDCAllOfMobile.md) |  | [optional] 
**BundleId** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.bundleId&#x60; instead.  A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  | [optional] 
**PackageName** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.packageName&#x60; instead.  A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  | [optional] 
**Kerberos** | Pointer to [**ApplicationOIDCAllOfKerberos**](ApplicationOIDCAllOfKerberos.md) |  | [optional] 
**GrantTypes** | [**[]EnumApplicationOIDCGrantType**](EnumApplicationOIDCGrantType.md) | A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS. | 
**HomePageUrl** | Pointer to **string** | A string that specifies the custom home page URL for the application. | [optional] 
**InitiateLoginUri** | Pointer to **string** | A string that specifies the URI to use for third-parties to begin the sign-on process for the application. If specified, PingOne redirects users to this URI to initiate SSO to PingOne. The application is responsible for implementing the relevant OIDC flow when the initiate login URI is requested. This property is required if you want the application to appear in the PingOne Application Portal. See the OIDC specification section of [Initiating Login from a Third Party](https://openid.net/specs/openid-connect-core-1_0.html#ThirdPartyInitiatedLogin) for more information. | [optional] 
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | A string that specifies the URLs that the browser can be redirected to after logout. | [optional] 
**RedirectUris** | Pointer to **[]string** | A string that specifies the callback URI for the authentication response. | [optional] 
**RefreshTokenDuration** | Pointer to **int32** | An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the &#x60;refreshTokenRollingDuration&#x60; property is specified for the application, then this property must be less than or equal to the value of &#x60;refreshTokenRollingDuration&#x60;. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] [default to 2592000]
**RefreshTokenRollingDuration** | Pointer to **int32** | An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] 
**RefreshTokenRollingGracePeriodDuration** | Pointer to **int32** | The number of seconds that a refresh token may be reused after having been exchanged for a new set of tokens. This is useful in the case of network errors on the client. Valid values are between 0 and 86400 seconds. Null is treated the same as 0. | [optional] 
**ResponseTypes** | Pointer to [**[]EnumApplicationOIDCResponseType**](EnumApplicationOIDCResponseType.md) | A string that specifies the code or token type returned by an authorization request. Options are TOKEN, ID_TOKEN, and CODE. Note that CODE cannot be used in an authorization request with TOKEN or ID_TOKEN because PingOne does not currently support OIDC hybrid flows. | [optional] 
**RequireSignedRequestObject** | Pointer to **bool** | Indicates that the Java Web Token (JWT) for the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter is required to be signed. If &#x60;false&#x60; or null (default), a signed request object is not required. Both &#x60;supportUnsignedRequestObject&#x60; and this property cannot be set to &#x60;true&#x60;. | [optional] 
**SupportUnsignedRequestObject** | Pointer to **bool** | A boolean that specifies whether the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed. | [optional] 
**Tags** | Pointer to [**[]EnumApplicationTags**](EnumApplicationTags.md) | An array that specifies the list of labels associated with the application. Options are &#x60;PING_FED_CONNECTION_INTEGRATION&#x60;.  Only applicable for creating worker applications. | [optional] 
**TargetLinkUri** | Pointer to **string** | The URI for the application. If specified, PingOne will redirect application users to this URI after a user is authenticated. In the PingOne admin console, this becomes the value of the &#x60;target_link_uri&#x60; parameter used for the Initiate Single Sign-On URL field. | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 
**ParRequirement** | Pointer to [**EnumApplicationOIDCPARRequirement**](EnumApplicationOIDCPARRequirement.md) |  | [optional] [default to ENUMAPPLICATIONOIDCPARREQUIREMENT_OPTIONAL]
**ParTimeout** | Pointer to **int32** | PAR timeout in seconds. Must be between &#x60;1&#x60; and &#x60;600&#x60;. The default value is &#x60;60&#x60;. | [optional] [default to 60]

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

### GetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDCAllOf) GetAdditionalRefreshTokenReplayProtectionEnabled() bool`

GetAdditionalRefreshTokenReplayProtectionEnabled returns the AdditionalRefreshTokenReplayProtectionEnabled field if non-nil, zero value otherwise.

### GetAdditionalRefreshTokenReplayProtectionEnabledOk

`func (o *ApplicationOIDCAllOf) GetAdditionalRefreshTokenReplayProtectionEnabledOk() (*bool, bool)`

GetAdditionalRefreshTokenReplayProtectionEnabledOk returns a tuple with the AdditionalRefreshTokenReplayProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDCAllOf) SetAdditionalRefreshTokenReplayProtectionEnabled(v bool)`

SetAdditionalRefreshTokenReplayProtectionEnabled sets AdditionalRefreshTokenReplayProtectionEnabled field to given value.

### HasAdditionalRefreshTokenReplayProtectionEnabled

`func (o *ApplicationOIDCAllOf) HasAdditionalRefreshTokenReplayProtectionEnabled() bool`

HasAdditionalRefreshTokenReplayProtectionEnabled returns a boolean if a field has been set.

### GetAllowWildcardInRedirectUris

`func (o *ApplicationOIDCAllOf) GetAllowWildcardInRedirectUris() bool`

GetAllowWildcardInRedirectUris returns the AllowWildcardInRedirectUris field if non-nil, zero value otherwise.

### GetAllowWildcardInRedirectUrisOk

`func (o *ApplicationOIDCAllOf) GetAllowWildcardInRedirectUrisOk() (*bool, bool)`

GetAllowWildcardInRedirectUrisOk returns a tuple with the AllowWildcardInRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardInRedirectUris

`func (o *ApplicationOIDCAllOf) SetAllowWildcardInRedirectUris(v bool)`

SetAllowWildcardInRedirectUris sets AllowWildcardInRedirectUris field to given value.

### HasAllowWildcardInRedirectUris

`func (o *ApplicationOIDCAllOf) HasAllowWildcardInRedirectUris() bool`

HasAllowWildcardInRedirectUris returns a boolean if a field has been set.

### GetAssignActorRoles

`func (o *ApplicationOIDCAllOf) GetAssignActorRoles() bool`

GetAssignActorRoles returns the AssignActorRoles field if non-nil, zero value otherwise.

### GetAssignActorRolesOk

`func (o *ApplicationOIDCAllOf) GetAssignActorRolesOk() (*bool, bool)`

GetAssignActorRolesOk returns a tuple with the AssignActorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignActorRoles

`func (o *ApplicationOIDCAllOf) SetAssignActorRoles(v bool)`

SetAssignActorRoles sets AssignActorRoles field to given value.

### HasAssignActorRoles

`func (o *ApplicationOIDCAllOf) HasAssignActorRoles() bool`

HasAssignActorRoles returns a boolean if a field has been set.

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

### GetKerberos

`func (o *ApplicationOIDCAllOf) GetKerberos() ApplicationOIDCAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *ApplicationOIDCAllOf) GetKerberosOk() (*ApplicationOIDCAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *ApplicationOIDCAllOf) SetKerberos(v ApplicationOIDCAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *ApplicationOIDCAllOf) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

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

### GetInitiateLoginUri

`func (o *ApplicationOIDCAllOf) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *ApplicationOIDCAllOf) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *ApplicationOIDCAllOf) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *ApplicationOIDCAllOf) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

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

### GetRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingGracePeriodDuration() int32`

GetRefreshTokenRollingGracePeriodDuration returns the RefreshTokenRollingGracePeriodDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodDurationOk

`func (o *ApplicationOIDCAllOf) GetRefreshTokenRollingGracePeriodDurationOk() (*int32, bool)`

GetRefreshTokenRollingGracePeriodDurationOk returns a tuple with the RefreshTokenRollingGracePeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDCAllOf) SetRefreshTokenRollingGracePeriodDuration(v int32)`

SetRefreshTokenRollingGracePeriodDuration sets RefreshTokenRollingGracePeriodDuration field to given value.

### HasRefreshTokenRollingGracePeriodDuration

`func (o *ApplicationOIDCAllOf) HasRefreshTokenRollingGracePeriodDuration() bool`

HasRefreshTokenRollingGracePeriodDuration returns a boolean if a field has been set.

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

### GetRequireSignedRequestObject

`func (o *ApplicationOIDCAllOf) GetRequireSignedRequestObject() bool`

GetRequireSignedRequestObject returns the RequireSignedRequestObject field if non-nil, zero value otherwise.

### GetRequireSignedRequestObjectOk

`func (o *ApplicationOIDCAllOf) GetRequireSignedRequestObjectOk() (*bool, bool)`

GetRequireSignedRequestObjectOk returns a tuple with the RequireSignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequestObject

`func (o *ApplicationOIDCAllOf) SetRequireSignedRequestObject(v bool)`

SetRequireSignedRequestObject sets RequireSignedRequestObject field to given value.

### HasRequireSignedRequestObject

`func (o *ApplicationOIDCAllOf) HasRequireSignedRequestObject() bool`

HasRequireSignedRequestObject returns a boolean if a field has been set.

### GetSupportUnsignedRequestObject

`func (o *ApplicationOIDCAllOf) GetSupportUnsignedRequestObject() bool`

GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field if non-nil, zero value otherwise.

### GetSupportUnsignedRequestObjectOk

`func (o *ApplicationOIDCAllOf) GetSupportUnsignedRequestObjectOk() (*bool, bool)`

GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUnsignedRequestObject

`func (o *ApplicationOIDCAllOf) SetSupportUnsignedRequestObject(v bool)`

SetSupportUnsignedRequestObject sets SupportUnsignedRequestObject field to given value.

### HasSupportUnsignedRequestObject

`func (o *ApplicationOIDCAllOf) HasSupportUnsignedRequestObject() bool`

HasSupportUnsignedRequestObject returns a boolean if a field has been set.

### GetTags

`func (o *ApplicationOIDCAllOf) GetTags() []EnumApplicationTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApplicationOIDCAllOf) GetTagsOk() (*[]EnumApplicationTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApplicationOIDCAllOf) SetTags(v []EnumApplicationTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApplicationOIDCAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetLinkUri

`func (o *ApplicationOIDCAllOf) GetTargetLinkUri() string`

GetTargetLinkUri returns the TargetLinkUri field if non-nil, zero value otherwise.

### GetTargetLinkUriOk

`func (o *ApplicationOIDCAllOf) GetTargetLinkUriOk() (*string, bool)`

GetTargetLinkUriOk returns a tuple with the TargetLinkUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLinkUri

`func (o *ApplicationOIDCAllOf) SetTargetLinkUri(v string)`

SetTargetLinkUri sets TargetLinkUri field to given value.

### HasTargetLinkUri

`func (o *ApplicationOIDCAllOf) HasTargetLinkUri() bool`

HasTargetLinkUri returns a boolean if a field has been set.

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


### GetParRequirement

`func (o *ApplicationOIDCAllOf) GetParRequirement() EnumApplicationOIDCPARRequirement`

GetParRequirement returns the ParRequirement field if non-nil, zero value otherwise.

### GetParRequirementOk

`func (o *ApplicationOIDCAllOf) GetParRequirementOk() (*EnumApplicationOIDCPARRequirement, bool)`

GetParRequirementOk returns a tuple with the ParRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequirement

`func (o *ApplicationOIDCAllOf) SetParRequirement(v EnumApplicationOIDCPARRequirement)`

SetParRequirement sets ParRequirement field to given value.

### HasParRequirement

`func (o *ApplicationOIDCAllOf) HasParRequirement() bool`

HasParRequirement returns a boolean if a field has been set.

### GetParTimeout

`func (o *ApplicationOIDCAllOf) GetParTimeout() int32`

GetParTimeout returns the ParTimeout field if non-nil, zero value otherwise.

### GetParTimeoutOk

`func (o *ApplicationOIDCAllOf) GetParTimeoutOk() (*int32, bool)`

GetParTimeoutOk returns a tuple with the ParTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParTimeout

`func (o *ApplicationOIDCAllOf) SetParTimeout(v int32)`

SetParTimeout sets ParTimeout field to given value.

### HasParTimeout

`func (o *ApplicationOIDCAllOf) HasParTimeout() bool`

HasParTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


