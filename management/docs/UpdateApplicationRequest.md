# UpdateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**AccessControl** | Pointer to [**ApplicationAccessControl**](ApplicationAccessControl.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the application. | [optional] 
**Enabled** | **bool** | A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**HiddenFromAppPortal** | Pointer to **bool** | A boolean to specify whether the application is hidden in the application portal despite the configured group access policy. | [optional] 
**Icon** | Pointer to [**ApplicationIcon**](ApplicationIcon.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**LoginPageUrl** | Pointer to **string** | A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains. | [optional] 
**Name** | **string** | A string that specifies the name of the application. This is a required property. | 
**Protocol** | [**EnumApplicationProtocol**](EnumApplicationProtocol.md) |  | 
**Type** | [**EnumApplicationType**](EnumApplicationType.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**HomePageUrl** | **string** | A string that specifies the custom home page URL for the application. | 
**AcsUrls** | **[]string** | A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property. | 
**AssertionDuration** | **int32** | An integer that specifies the assertion validity duration in seconds. This is a required property. | 
**AssertionSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion itself should be signed. The default value is &#x60;true&#x60;. | [optional] [default to true]
**CorsSettings** | Pointer to [**ApplicationCorsSettings**](ApplicationCorsSettings.md) |  | [optional] 
**DefaultTargetUrl** | Pointer to **string** | This is used as the RelayState parameter by the IdP to deep link into the application after authentication. This value can be overridden by the applicationUrl query parameter for GET Identity Provider Initiated SSO. Although both of these parameters are generally URLs, because they are used as deep links, this is not enforced. If neither defaultTargetUrl nor applicationUrl is specified during a SAML authentication flow, no RelayState value is supplied to the application. The defaultTargetUrl (or the applicationUrl) value is passed to the SAML application&#39;s ACS URL as a separate RelayState key value (not within the SAMLResponse key value). | [optional] 
**EnableRequestedAuthnContext** | Pointer to **bool** | Indicates whether &#x60;requestedAuthnContext&#x60; is taken into account in policy decision-making during authentication. | [optional] 
**IdpSigning** | [**ApplicationWSFEDAllOfIdpSigning**](ApplicationWSFEDAllOfIdpSigning.md) |  | 
**NameIdFormat** | Pointer to **string** | A string that specifies the format of the Subject NameID attibute in the SAML assertion | [optional] 
**ResponseSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion response itself should be signed. The default value is &#x60;false&#x60;. | [optional] [default to false]
**SessionNotOnOrAfterDuration** | Pointer to **int32** | Update this value if the SAML application requires a different &#x60;SessionNotOnOrAfter&#x60; attribute value within the &#x60;AuthnStatement&#x60; element than the &#x60;NotOnOrAfter&#x60; value set by the &#x60;assertionDuration&#x60; property. | [optional] 
**SloBinding** | Pointer to [**EnumApplicationSAMLSloBinding**](EnumApplicationSAMLSloBinding.md) |  | [optional] [default to ENUMAPPLICATIONSAMLSLOBINDING_POST]
**SloEndpoint** | Pointer to **string** | The single logout endpoint URL. | [optional] 
**SloResponseEndpoint** | Pointer to **string** | A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response. | [optional] 
**SloWindow** | Pointer to **int32** | Defines how long PingOne can exchange logout messages with the application, specifically a &#x60;LogoutRequest&#x60; from the application, since the initial request. PingOne can also send a &#x60;LogoutRequest&#x60; to the application when a single logout is initiated by the user from other session participants, such as an application or identity provider. This setting is per application. The SLO logout is separate from the user session logout that revokes all tokens. | [optional] 
**SpEncryption** | Pointer to [**ApplicationSAMLAllOfSpEncryption**](ApplicationSAMLAllOfSpEncryption.md) |  | [optional] 
**SpEntityId** | **string** | A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment. | 
**SpVerification** | Pointer to [**ApplicationSAMLAllOfSpVerification**](ApplicationSAMLAllOfSpVerification.md) |  | [optional] 
**Template** | Pointer to [**ApplicationTemplate**](ApplicationTemplate.md) |  | [optional] 
**VirtualServerIdSettings** | Pointer to [**ApplicationSAMLAllOfVirtualServerIdSettings**](ApplicationSAMLAllOfVirtualServerIdSettings.md) |  | [optional] 
**AdditionalRefreshTokenReplayProtectionEnabled** | Pointer to **bool** | When set to &#x60;true&#x60; (the default), if you attempt to reuse the refresh token, the authorization server immediately revokes the reused refresh token, as well as all descendant tokens. Setting this to null equates to a &#x60;false&#x60; setting. | [optional] [default to true]
**AllowWildcardInRedirectUris** | Pointer to **bool** | A boolean to specify whether wildcards are allowed in redirect URIs. For more information, see [Wildcards in Redirect URIs](https://docs.pingidentity.com/csh?context&#x3D;p1_c_wildcard_redirect_uri). | [optional] 
**AssignActorRoles** | Pointer to **bool** | A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request. | [optional] 
**DevicePathId** | Pointer to **string** | A string that specifies a unique identifier within an environment for a device authorization grant flow to provide a short identifier to the application. This property is ignored when the &#x60;deviceCustomVerificationUri&#x60; property is configured. The string can contain any letters, numbers, and some special characters (regex &#x60;a-zA-Z0-9_-&#x60;). It can have a length of no more than 50 characters (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;50&#x60;). | [optional] 
**DeviceCustomVerificationUri** | Pointer to **string** | A string that specifies an optional custom verification URI that is returned for the &#x60;/device_authorization&#x60; endpoint. | [optional] 
**DeviceTimeout** | Pointer to **int32** | An integer that specifies the length of time (in seconds) that the &#x60;userCode&#x60; and &#x60;deviceCode&#x60; returned by the &#x60;/device_authorization&#x60; endpoint are valid. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;600&#x60; seconds. It can have a value of no more than &#x60;3600&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;3600&#x60;). | [optional] [default to 600]
**DevicePollingInterval** | Pointer to **int32** | An integer that specifies the frequency (in seconds) for the client to poll the &#x60;/as/token&#x60; endpoint. This property is required only for applications in which the &#x60;grantTypes&#x60; property is set to &#x60;device_code&#x60;. The default value is &#x60;5&#x60; seconds. It can have a value of no more than &#x60;60&#x60; seconds (&#x60;min&#x60;/&#x60;max&#x60;&#x3D;&#x60;1&#x60;/&#x60;60&#x60;). | [optional] [default to 5]
**IdpSignoff** | Pointer to **bool** | Set this to true to allow an application to request to terminate a user session using only the ID token. The application is not required to have access to the session token cookie. | [optional] 
**Jwks** | Pointer to **string** | A JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwksUrl&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwksUrl&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**JwksUrl** | Pointer to **string** | A URL (supports &#x60;https://&#x60; only) that provides access to a JWKS string that validates the signature of signed JWTs for applications that use the &#x60;PRIVATE_KEY_JWT&#x60; option for the &#x60;tokenEndpointAuthMethod&#x60;. This property is required when &#x60;tokenEndpointAuthMethod&#x60; is &#x60;PRIVATE_KEY_JWT&#x60; and the &#x60;jwks&#x60; property is empty. For more information, see [Create a private_key_jwt JWKS string](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-private_key_jwt-jwks-string). This property is also required if the optional &#x60;request&#x60; property JWT on the authorize endpoint is signed using the RS256 (or RS384, RS512) signing algorithm and the &#x60;jwks&#x60; property is empty. For more infornmation about signing the request property JWT, see [Create a request property JWT](https://apidocs.pingidentity.com/pingone/platform/v1/api/#create-a-request-property-jwt). | [optional] 
**Mobile** | Pointer to [**ApplicationOIDCAllOfMobile**](ApplicationOIDCAllOfMobile.md) |  | [optional] 
**BundleId** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.bundleId&#x60; instead.  A string that specifies the bundle associated with the application, for push notifications in native apps. The value of the bundleId property is unique per environment, and once defined, is immutable.  | [optional] 
**PackageName** | Pointer to **string** | **Deprecation Notice** This field is deprecated and will be removed in a future release. Use &#x60;mobile.packageName&#x60; instead.  A string that specifies the package name associated with the application, for push notifications in native apps. The value of the mobile.packageName property is unique per environment, and once defined, is immutable.  | [optional] 
**Kerberos** | Pointer to [**ApplicationWSFEDAllOfKerberos**](ApplicationWSFEDAllOfKerberos.md) |  | [optional] 
**GrantTypes** | [**[]EnumApplicationOIDCGrantType**](EnumApplicationOIDCGrantType.md) | A string that specifies the grant type for the authorization request. This is a required property. Options are AUTHORIZATION_CODE, IMPLICIT, REFRESH_TOKEN, CLIENT_CREDENTIALS. | 
**InitiateLoginUri** | Pointer to **string** | A string that specifies the URI to use for third-parties to begin the sign-on process for the application. If specified, PingOne redirects users to this URI to initiate SSO to PingOne. The application is responsible for implementing the relevant OIDC flow when the initiate login URI is requested. This property is required if you want the application to appear in the PingOne Application Portal. See the OIDC specification section of [Initiating Login from a Third Party](https://openid.net/specs/openid-connect-core-1_0.html#ThirdPartyInitiatedLogin) for more information. | [optional] 
**PkceEnforcement** | Pointer to [**EnumApplicationOIDCPKCEOption**](EnumApplicationOIDCPKCEOption.md) |  | [optional] 
**PostLogoutRedirectUris** | Pointer to **[]string** | A string that specifies the URLs that the browser can be redirected to after logout. | [optional] 
**RedirectUris** | Pointer to **[]string** | A string that specifies the callback URI for the authentication response. | [optional] 
**RefreshTokenDuration** | Pointer to **int32** | An integer that specifies the lifetime in seconds of the refresh token. If a value is not provided, the default value is 2592000, or 30 days. Valid values are between 60 and 2147483647. If the &#x60;refreshTokenRollingDuration&#x60; property is specified for the application, then this property must be less than or equal to the value of &#x60;refreshTokenRollingDuration&#x60;. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] [default to 2592000]
**RefreshTokenRollingDuration** | Pointer to **int32** | An integer that specifies the number of seconds a refresh token can be exchanged before re-authentication is required. If a value is not provided, the refresh token is valid forever. Valid values are between 60 and 2147483647. After this property is set, the value cannot be nullified. This value is used to generate the value for the exp claim when minting a new refresh token. | [optional] 
**RefreshTokenRollingGracePeriodDuration** | Pointer to **int32** | The number of seconds that a refresh token may be reused after having been exchanged for a new set of tokens. This is useful in the case of network errors on the client. Valid values are between 0 and 86400 seconds. Null is treated the same as 0. | [optional] 
**ResponseTypes** | Pointer to [**[]EnumApplicationOIDCResponseType**](EnumApplicationOIDCResponseType.md) | The code or token type returned by an authorization request. Options are &#x60;TOKEN&#x60;, &#x60;ID_TOKEN&#x60;, and &#x60;CODE&#x60;. For hybrid flows that specify &#x60;CODE&#x60; with &#x60;TOKEN&#x60; or &#x60;ID_TOKEN&#x60;, see [Hybrid grant type](https://apidocs.pingidentity.com/pingone/main/v1/api/#hybrid-grant-type). | [optional] 
**RequireSignedRequestObject** | Pointer to **bool** | Indicates that the Java Web Token (JWT) for the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter is required to be signed. If &#x60;false&#x60; or null (default), a signed request object is not required. Both &#x60;supportUnsignedRequestObject&#x60; and this property cannot be set to &#x60;true&#x60;. | [optional] 
**SupportUnsignedRequestObject** | Pointer to **bool** | A boolean that specifies whether the [request query](https://openid.net/specs/openid-connect-core-1_0.html#RequestObject) parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed. | [optional] 
**Tags** | Pointer to [**[]EnumApplicationTags**](EnumApplicationTags.md) | An array that specifies the list of labels associated with the application. Options are &#x60;PING_FED_CONNECTION_INTEGRATION&#x60;.  Only applicable for creating worker applications. | [optional] 
**TargetLinkUri** | Pointer to **string** | The URI for the application. If specified, PingOne will redirect application users to this URI after a user is authenticated. In the PingOne admin console, this becomes the value of the &#x60;target_link_uri&#x60; parameter used for the Initiate Single Sign-On URL field. | [optional] 
**TokenEndpointAuthMethod** | [**EnumApplicationOIDCTokenAuthMethod**](EnumApplicationOIDCTokenAuthMethod.md) |  | 
**ParRequirement** | Pointer to [**EnumApplicationOIDCPARRequirement**](EnumApplicationOIDCPARRequirement.md) |  | [optional] [default to ENUMAPPLICATIONOIDCPARREQUIREMENT_OPTIONAL]
**ParTimeout** | Pointer to **int32** | PAR timeout in seconds. Must be between &#x60;1&#x60; and &#x60;600&#x60;. The default value is &#x60;60&#x60;. | [optional] [default to 60]
**Signing** | Pointer to [**ApplicationOIDCAllOfSigning**](ApplicationOIDCAllOfSigning.md) |  | [optional] 
**AudienceRestriction** | Pointer to **string** | The service provider ID. Defaults to &#x60;urn:federation:MicrosoftOnline&#x60;. | [optional] [default to "urn:federation:MicrosoftOnline"]
**DomainName** | **string** | The federated domain name (for example, the Azure custom domain). | 
**ReplyUrl** | **string** | The URL that the replying party (such as, Office365) uses to accept submissions of RequestSecurityTokenResponse messages that are a result of SSO requests. | 
**SubjectNameIdentifierFormat** | Pointer to [**EnumApplicationWSFEDSubjectNameIdentifierFormat**](EnumApplicationWSFEDSubjectNameIdentifierFormat.md) |  | [optional] 
**ApplyDefaultTheme** | **bool** | If &#x60;true&#x60;, applies the default theme to the self service application. | 
**EnableDefaultThemeFooter** | Pointer to **bool** | If &#x60;true&#x60;, shows the default theme footer on the self service application. Applies only if &#x60;applyDefaultTheme&#x60; is also &#x60;true&#x60;. | [optional] 

## Methods

### NewUpdateApplicationRequest

`func NewUpdateApplicationRequest(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, homePageUrl string, acsUrls []string, assertionDuration int32, idpSigning ApplicationWSFEDAllOfIdpSigning, spEntityId string, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, domainName string, replyUrl string, applyDefaultTheme bool, ) *UpdateApplicationRequest`

NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestWithDefaults

`func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest`

NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UpdateApplicationRequest) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateApplicationRequest) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateApplicationRequest) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateApplicationRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccessControl

`func (o *UpdateApplicationRequest) GetAccessControl() ApplicationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *UpdateApplicationRequest) GetAccessControlOk() (*ApplicationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *UpdateApplicationRequest) SetAccessControl(v ApplicationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *UpdateApplicationRequest) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateApplicationRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateApplicationRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateApplicationRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateApplicationRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateApplicationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateApplicationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateApplicationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *UpdateApplicationRequest) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateApplicationRequest) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateApplicationRequest) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateApplicationRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetHiddenFromAppPortal

`func (o *UpdateApplicationRequest) GetHiddenFromAppPortal() bool`

GetHiddenFromAppPortal returns the HiddenFromAppPortal field if non-nil, zero value otherwise.

### GetHiddenFromAppPortalOk

`func (o *UpdateApplicationRequest) GetHiddenFromAppPortalOk() (*bool, bool)`

GetHiddenFromAppPortalOk returns a tuple with the HiddenFromAppPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromAppPortal

`func (o *UpdateApplicationRequest) SetHiddenFromAppPortal(v bool)`

SetHiddenFromAppPortal sets HiddenFromAppPortal field to given value.

### HasHiddenFromAppPortal

`func (o *UpdateApplicationRequest) HasHiddenFromAppPortal() bool`

HasHiddenFromAppPortal returns a boolean if a field has been set.

### GetIcon

`func (o *UpdateApplicationRequest) GetIcon() ApplicationIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *UpdateApplicationRequest) GetIconOk() (*ApplicationIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *UpdateApplicationRequest) SetIcon(v ApplicationIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *UpdateApplicationRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *UpdateApplicationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateApplicationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateApplicationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateApplicationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginPageUrl

`func (o *UpdateApplicationRequest) GetLoginPageUrl() string`

GetLoginPageUrl returns the LoginPageUrl field if non-nil, zero value otherwise.

### GetLoginPageUrlOk

`func (o *UpdateApplicationRequest) GetLoginPageUrlOk() (*string, bool)`

GetLoginPageUrlOk returns a tuple with the LoginPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageUrl

`func (o *UpdateApplicationRequest) SetLoginPageUrl(v string)`

SetLoginPageUrl sets LoginPageUrl field to given value.

### HasLoginPageUrl

`func (o *UpdateApplicationRequest) HasLoginPageUrl() bool`

HasLoginPageUrl returns a boolean if a field has been set.

### GetName

`func (o *UpdateApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *UpdateApplicationRequest) GetProtocol() EnumApplicationProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *UpdateApplicationRequest) GetProtocolOk() (*EnumApplicationProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *UpdateApplicationRequest) SetProtocol(v EnumApplicationProtocol)`

SetProtocol sets Protocol field to given value.


### GetType

`func (o *UpdateApplicationRequest) GetType() EnumApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateApplicationRequest) GetTypeOk() (*EnumApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateApplicationRequest) SetType(v EnumApplicationType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *UpdateApplicationRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateApplicationRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateApplicationRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateApplicationRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHomePageUrl

`func (o *UpdateApplicationRequest) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *UpdateApplicationRequest) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *UpdateApplicationRequest) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.


### GetAcsUrls

`func (o *UpdateApplicationRequest) GetAcsUrls() []string`

GetAcsUrls returns the AcsUrls field if non-nil, zero value otherwise.

### GetAcsUrlsOk

`func (o *UpdateApplicationRequest) GetAcsUrlsOk() (*[]string, bool)`

GetAcsUrlsOk returns a tuple with the AcsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrls

`func (o *UpdateApplicationRequest) SetAcsUrls(v []string)`

SetAcsUrls sets AcsUrls field to given value.


### GetAssertionDuration

`func (o *UpdateApplicationRequest) GetAssertionDuration() int32`

GetAssertionDuration returns the AssertionDuration field if non-nil, zero value otherwise.

### GetAssertionDurationOk

`func (o *UpdateApplicationRequest) GetAssertionDurationOk() (*int32, bool)`

GetAssertionDurationOk returns a tuple with the AssertionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionDuration

`func (o *UpdateApplicationRequest) SetAssertionDuration(v int32)`

SetAssertionDuration sets AssertionDuration field to given value.


### GetAssertionSigned

`func (o *UpdateApplicationRequest) GetAssertionSigned() bool`

GetAssertionSigned returns the AssertionSigned field if non-nil, zero value otherwise.

### GetAssertionSignedOk

`func (o *UpdateApplicationRequest) GetAssertionSignedOk() (*bool, bool)`

GetAssertionSignedOk returns a tuple with the AssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSigned

`func (o *UpdateApplicationRequest) SetAssertionSigned(v bool)`

SetAssertionSigned sets AssertionSigned field to given value.

### HasAssertionSigned

`func (o *UpdateApplicationRequest) HasAssertionSigned() bool`

HasAssertionSigned returns a boolean if a field has been set.

### GetCorsSettings

`func (o *UpdateApplicationRequest) GetCorsSettings() ApplicationCorsSettings`

GetCorsSettings returns the CorsSettings field if non-nil, zero value otherwise.

### GetCorsSettingsOk

`func (o *UpdateApplicationRequest) GetCorsSettingsOk() (*ApplicationCorsSettings, bool)`

GetCorsSettingsOk returns a tuple with the CorsSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsSettings

`func (o *UpdateApplicationRequest) SetCorsSettings(v ApplicationCorsSettings)`

SetCorsSettings sets CorsSettings field to given value.

### HasCorsSettings

`func (o *UpdateApplicationRequest) HasCorsSettings() bool`

HasCorsSettings returns a boolean if a field has been set.

### GetDefaultTargetUrl

`func (o *UpdateApplicationRequest) GetDefaultTargetUrl() string`

GetDefaultTargetUrl returns the DefaultTargetUrl field if non-nil, zero value otherwise.

### GetDefaultTargetUrlOk

`func (o *UpdateApplicationRequest) GetDefaultTargetUrlOk() (*string, bool)`

GetDefaultTargetUrlOk returns a tuple with the DefaultTargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetUrl

`func (o *UpdateApplicationRequest) SetDefaultTargetUrl(v string)`

SetDefaultTargetUrl sets DefaultTargetUrl field to given value.

### HasDefaultTargetUrl

`func (o *UpdateApplicationRequest) HasDefaultTargetUrl() bool`

HasDefaultTargetUrl returns a boolean if a field has been set.

### GetEnableRequestedAuthnContext

`func (o *UpdateApplicationRequest) GetEnableRequestedAuthnContext() bool`

GetEnableRequestedAuthnContext returns the EnableRequestedAuthnContext field if non-nil, zero value otherwise.

### GetEnableRequestedAuthnContextOk

`func (o *UpdateApplicationRequest) GetEnableRequestedAuthnContextOk() (*bool, bool)`

GetEnableRequestedAuthnContextOk returns a tuple with the EnableRequestedAuthnContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRequestedAuthnContext

`func (o *UpdateApplicationRequest) SetEnableRequestedAuthnContext(v bool)`

SetEnableRequestedAuthnContext sets EnableRequestedAuthnContext field to given value.

### HasEnableRequestedAuthnContext

`func (o *UpdateApplicationRequest) HasEnableRequestedAuthnContext() bool`

HasEnableRequestedAuthnContext returns a boolean if a field has been set.

### GetIdpSigning

`func (o *UpdateApplicationRequest) GetIdpSigning() ApplicationWSFEDAllOfIdpSigning`

GetIdpSigning returns the IdpSigning field if non-nil, zero value otherwise.

### GetIdpSigningOk

`func (o *UpdateApplicationRequest) GetIdpSigningOk() (*ApplicationWSFEDAllOfIdpSigning, bool)`

GetIdpSigningOk returns a tuple with the IdpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSigning

`func (o *UpdateApplicationRequest) SetIdpSigning(v ApplicationWSFEDAllOfIdpSigning)`

SetIdpSigning sets IdpSigning field to given value.


### GetNameIdFormat

`func (o *UpdateApplicationRequest) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *UpdateApplicationRequest) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *UpdateApplicationRequest) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *UpdateApplicationRequest) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetResponseSigned

`func (o *UpdateApplicationRequest) GetResponseSigned() bool`

GetResponseSigned returns the ResponseSigned field if non-nil, zero value otherwise.

### GetResponseSignedOk

`func (o *UpdateApplicationRequest) GetResponseSignedOk() (*bool, bool)`

GetResponseSignedOk returns a tuple with the ResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSigned

`func (o *UpdateApplicationRequest) SetResponseSigned(v bool)`

SetResponseSigned sets ResponseSigned field to given value.

### HasResponseSigned

`func (o *UpdateApplicationRequest) HasResponseSigned() bool`

HasResponseSigned returns a boolean if a field has been set.

### GetSessionNotOnOrAfterDuration

`func (o *UpdateApplicationRequest) GetSessionNotOnOrAfterDuration() int32`

GetSessionNotOnOrAfterDuration returns the SessionNotOnOrAfterDuration field if non-nil, zero value otherwise.

### GetSessionNotOnOrAfterDurationOk

`func (o *UpdateApplicationRequest) GetSessionNotOnOrAfterDurationOk() (*int32, bool)`

GetSessionNotOnOrAfterDurationOk returns a tuple with the SessionNotOnOrAfterDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionNotOnOrAfterDuration

`func (o *UpdateApplicationRequest) SetSessionNotOnOrAfterDuration(v int32)`

SetSessionNotOnOrAfterDuration sets SessionNotOnOrAfterDuration field to given value.

### HasSessionNotOnOrAfterDuration

`func (o *UpdateApplicationRequest) HasSessionNotOnOrAfterDuration() bool`

HasSessionNotOnOrAfterDuration returns a boolean if a field has been set.

### GetSloBinding

`func (o *UpdateApplicationRequest) GetSloBinding() EnumApplicationSAMLSloBinding`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *UpdateApplicationRequest) GetSloBindingOk() (*EnumApplicationSAMLSloBinding, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *UpdateApplicationRequest) SetSloBinding(v EnumApplicationSAMLSloBinding)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *UpdateApplicationRequest) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### GetSloEndpoint

`func (o *UpdateApplicationRequest) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *UpdateApplicationRequest) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *UpdateApplicationRequest) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *UpdateApplicationRequest) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetSloResponseEndpoint

`func (o *UpdateApplicationRequest) GetSloResponseEndpoint() string`

GetSloResponseEndpoint returns the SloResponseEndpoint field if non-nil, zero value otherwise.

### GetSloResponseEndpointOk

`func (o *UpdateApplicationRequest) GetSloResponseEndpointOk() (*string, bool)`

GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloResponseEndpoint

`func (o *UpdateApplicationRequest) SetSloResponseEndpoint(v string)`

SetSloResponseEndpoint sets SloResponseEndpoint field to given value.

### HasSloResponseEndpoint

`func (o *UpdateApplicationRequest) HasSloResponseEndpoint() bool`

HasSloResponseEndpoint returns a boolean if a field has been set.

### GetSloWindow

`func (o *UpdateApplicationRequest) GetSloWindow() int32`

GetSloWindow returns the SloWindow field if non-nil, zero value otherwise.

### GetSloWindowOk

`func (o *UpdateApplicationRequest) GetSloWindowOk() (*int32, bool)`

GetSloWindowOk returns a tuple with the SloWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloWindow

`func (o *UpdateApplicationRequest) SetSloWindow(v int32)`

SetSloWindow sets SloWindow field to given value.

### HasSloWindow

`func (o *UpdateApplicationRequest) HasSloWindow() bool`

HasSloWindow returns a boolean if a field has been set.

### GetSpEncryption

`func (o *UpdateApplicationRequest) GetSpEncryption() ApplicationSAMLAllOfSpEncryption`

GetSpEncryption returns the SpEncryption field if non-nil, zero value otherwise.

### GetSpEncryptionOk

`func (o *UpdateApplicationRequest) GetSpEncryptionOk() (*ApplicationSAMLAllOfSpEncryption, bool)`

GetSpEncryptionOk returns a tuple with the SpEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEncryption

`func (o *UpdateApplicationRequest) SetSpEncryption(v ApplicationSAMLAllOfSpEncryption)`

SetSpEncryption sets SpEncryption field to given value.

### HasSpEncryption

`func (o *UpdateApplicationRequest) HasSpEncryption() bool`

HasSpEncryption returns a boolean if a field has been set.

### GetSpEntityId

`func (o *UpdateApplicationRequest) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *UpdateApplicationRequest) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *UpdateApplicationRequest) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.


### GetSpVerification

`func (o *UpdateApplicationRequest) GetSpVerification() ApplicationSAMLAllOfSpVerification`

GetSpVerification returns the SpVerification field if non-nil, zero value otherwise.

### GetSpVerificationOk

`func (o *UpdateApplicationRequest) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool)`

GetSpVerificationOk returns a tuple with the SpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVerification

`func (o *UpdateApplicationRequest) SetSpVerification(v ApplicationSAMLAllOfSpVerification)`

SetSpVerification sets SpVerification field to given value.

### HasSpVerification

`func (o *UpdateApplicationRequest) HasSpVerification() bool`

HasSpVerification returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateApplicationRequest) GetTemplate() ApplicationTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateApplicationRequest) GetTemplateOk() (*ApplicationTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateApplicationRequest) SetTemplate(v ApplicationTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateApplicationRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetVirtualServerIdSettings

`func (o *UpdateApplicationRequest) GetVirtualServerIdSettings() ApplicationSAMLAllOfVirtualServerIdSettings`

GetVirtualServerIdSettings returns the VirtualServerIdSettings field if non-nil, zero value otherwise.

### GetVirtualServerIdSettingsOk

`func (o *UpdateApplicationRequest) GetVirtualServerIdSettingsOk() (*ApplicationSAMLAllOfVirtualServerIdSettings, bool)`

GetVirtualServerIdSettingsOk returns a tuple with the VirtualServerIdSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServerIdSettings

`func (o *UpdateApplicationRequest) SetVirtualServerIdSettings(v ApplicationSAMLAllOfVirtualServerIdSettings)`

SetVirtualServerIdSettings sets VirtualServerIdSettings field to given value.

### HasVirtualServerIdSettings

`func (o *UpdateApplicationRequest) HasVirtualServerIdSettings() bool`

HasVirtualServerIdSettings returns a boolean if a field has been set.

### GetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *UpdateApplicationRequest) GetAdditionalRefreshTokenReplayProtectionEnabled() bool`

GetAdditionalRefreshTokenReplayProtectionEnabled returns the AdditionalRefreshTokenReplayProtectionEnabled field if non-nil, zero value otherwise.

### GetAdditionalRefreshTokenReplayProtectionEnabledOk

`func (o *UpdateApplicationRequest) GetAdditionalRefreshTokenReplayProtectionEnabledOk() (*bool, bool)`

GetAdditionalRefreshTokenReplayProtectionEnabledOk returns a tuple with the AdditionalRefreshTokenReplayProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRefreshTokenReplayProtectionEnabled

`func (o *UpdateApplicationRequest) SetAdditionalRefreshTokenReplayProtectionEnabled(v bool)`

SetAdditionalRefreshTokenReplayProtectionEnabled sets AdditionalRefreshTokenReplayProtectionEnabled field to given value.

### HasAdditionalRefreshTokenReplayProtectionEnabled

`func (o *UpdateApplicationRequest) HasAdditionalRefreshTokenReplayProtectionEnabled() bool`

HasAdditionalRefreshTokenReplayProtectionEnabled returns a boolean if a field has been set.

### GetAllowWildcardInRedirectUris

`func (o *UpdateApplicationRequest) GetAllowWildcardInRedirectUris() bool`

GetAllowWildcardInRedirectUris returns the AllowWildcardInRedirectUris field if non-nil, zero value otherwise.

### GetAllowWildcardInRedirectUrisOk

`func (o *UpdateApplicationRequest) GetAllowWildcardInRedirectUrisOk() (*bool, bool)`

GetAllowWildcardInRedirectUrisOk returns a tuple with the AllowWildcardInRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWildcardInRedirectUris

`func (o *UpdateApplicationRequest) SetAllowWildcardInRedirectUris(v bool)`

SetAllowWildcardInRedirectUris sets AllowWildcardInRedirectUris field to given value.

### HasAllowWildcardInRedirectUris

`func (o *UpdateApplicationRequest) HasAllowWildcardInRedirectUris() bool`

HasAllowWildcardInRedirectUris returns a boolean if a field has been set.

### GetAssignActorRoles

`func (o *UpdateApplicationRequest) GetAssignActorRoles() bool`

GetAssignActorRoles returns the AssignActorRoles field if non-nil, zero value otherwise.

### GetAssignActorRolesOk

`func (o *UpdateApplicationRequest) GetAssignActorRolesOk() (*bool, bool)`

GetAssignActorRolesOk returns a tuple with the AssignActorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignActorRoles

`func (o *UpdateApplicationRequest) SetAssignActorRoles(v bool)`

SetAssignActorRoles sets AssignActorRoles field to given value.

### HasAssignActorRoles

`func (o *UpdateApplicationRequest) HasAssignActorRoles() bool`

HasAssignActorRoles returns a boolean if a field has been set.

### GetDevicePathId

`func (o *UpdateApplicationRequest) GetDevicePathId() string`

GetDevicePathId returns the DevicePathId field if non-nil, zero value otherwise.

### GetDevicePathIdOk

`func (o *UpdateApplicationRequest) GetDevicePathIdOk() (*string, bool)`

GetDevicePathIdOk returns a tuple with the DevicePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePathId

`func (o *UpdateApplicationRequest) SetDevicePathId(v string)`

SetDevicePathId sets DevicePathId field to given value.

### HasDevicePathId

`func (o *UpdateApplicationRequest) HasDevicePathId() bool`

HasDevicePathId returns a boolean if a field has been set.

### GetDeviceCustomVerificationUri

`func (o *UpdateApplicationRequest) GetDeviceCustomVerificationUri() string`

GetDeviceCustomVerificationUri returns the DeviceCustomVerificationUri field if non-nil, zero value otherwise.

### GetDeviceCustomVerificationUriOk

`func (o *UpdateApplicationRequest) GetDeviceCustomVerificationUriOk() (*string, bool)`

GetDeviceCustomVerificationUriOk returns a tuple with the DeviceCustomVerificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCustomVerificationUri

`func (o *UpdateApplicationRequest) SetDeviceCustomVerificationUri(v string)`

SetDeviceCustomVerificationUri sets DeviceCustomVerificationUri field to given value.

### HasDeviceCustomVerificationUri

`func (o *UpdateApplicationRequest) HasDeviceCustomVerificationUri() bool`

HasDeviceCustomVerificationUri returns a boolean if a field has been set.

### GetDeviceTimeout

`func (o *UpdateApplicationRequest) GetDeviceTimeout() int32`

GetDeviceTimeout returns the DeviceTimeout field if non-nil, zero value otherwise.

### GetDeviceTimeoutOk

`func (o *UpdateApplicationRequest) GetDeviceTimeoutOk() (*int32, bool)`

GetDeviceTimeoutOk returns a tuple with the DeviceTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTimeout

`func (o *UpdateApplicationRequest) SetDeviceTimeout(v int32)`

SetDeviceTimeout sets DeviceTimeout field to given value.

### HasDeviceTimeout

`func (o *UpdateApplicationRequest) HasDeviceTimeout() bool`

HasDeviceTimeout returns a boolean if a field has been set.

### GetDevicePollingInterval

`func (o *UpdateApplicationRequest) GetDevicePollingInterval() int32`

GetDevicePollingInterval returns the DevicePollingInterval field if non-nil, zero value otherwise.

### GetDevicePollingIntervalOk

`func (o *UpdateApplicationRequest) GetDevicePollingIntervalOk() (*int32, bool)`

GetDevicePollingIntervalOk returns a tuple with the DevicePollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePollingInterval

`func (o *UpdateApplicationRequest) SetDevicePollingInterval(v int32)`

SetDevicePollingInterval sets DevicePollingInterval field to given value.

### HasDevicePollingInterval

`func (o *UpdateApplicationRequest) HasDevicePollingInterval() bool`

HasDevicePollingInterval returns a boolean if a field has been set.

### GetIdpSignoff

`func (o *UpdateApplicationRequest) GetIdpSignoff() bool`

GetIdpSignoff returns the IdpSignoff field if non-nil, zero value otherwise.

### GetIdpSignoffOk

`func (o *UpdateApplicationRequest) GetIdpSignoffOk() (*bool, bool)`

GetIdpSignoffOk returns a tuple with the IdpSignoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSignoff

`func (o *UpdateApplicationRequest) SetIdpSignoff(v bool)`

SetIdpSignoff sets IdpSignoff field to given value.

### HasIdpSignoff

`func (o *UpdateApplicationRequest) HasIdpSignoff() bool`

HasIdpSignoff returns a boolean if a field has been set.

### GetJwks

`func (o *UpdateApplicationRequest) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *UpdateApplicationRequest) GetJwksOk() (*string, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *UpdateApplicationRequest) SetJwks(v string)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *UpdateApplicationRequest) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetJwksUrl

`func (o *UpdateApplicationRequest) GetJwksUrl() string`

GetJwksUrl returns the JwksUrl field if non-nil, zero value otherwise.

### GetJwksUrlOk

`func (o *UpdateApplicationRequest) GetJwksUrlOk() (*string, bool)`

GetJwksUrlOk returns a tuple with the JwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUrl

`func (o *UpdateApplicationRequest) SetJwksUrl(v string)`

SetJwksUrl sets JwksUrl field to given value.

### HasJwksUrl

`func (o *UpdateApplicationRequest) HasJwksUrl() bool`

HasJwksUrl returns a boolean if a field has been set.

### GetMobile

`func (o *UpdateApplicationRequest) GetMobile() ApplicationOIDCAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *UpdateApplicationRequest) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *UpdateApplicationRequest) SetMobile(v ApplicationOIDCAllOfMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *UpdateApplicationRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetBundleId

`func (o *UpdateApplicationRequest) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *UpdateApplicationRequest) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *UpdateApplicationRequest) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *UpdateApplicationRequest) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *UpdateApplicationRequest) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *UpdateApplicationRequest) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *UpdateApplicationRequest) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *UpdateApplicationRequest) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetKerberos

`func (o *UpdateApplicationRequest) GetKerberos() ApplicationWSFEDAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *UpdateApplicationRequest) GetKerberosOk() (*ApplicationWSFEDAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *UpdateApplicationRequest) SetKerberos(v ApplicationWSFEDAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *UpdateApplicationRequest) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetGrantTypes

`func (o *UpdateApplicationRequest) GetGrantTypes() []EnumApplicationOIDCGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *UpdateApplicationRequest) GetGrantTypesOk() (*[]EnumApplicationOIDCGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *UpdateApplicationRequest) SetGrantTypes(v []EnumApplicationOIDCGrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetInitiateLoginUri

`func (o *UpdateApplicationRequest) GetInitiateLoginUri() string`

GetInitiateLoginUri returns the InitiateLoginUri field if non-nil, zero value otherwise.

### GetInitiateLoginUriOk

`func (o *UpdateApplicationRequest) GetInitiateLoginUriOk() (*string, bool)`

GetInitiateLoginUriOk returns a tuple with the InitiateLoginUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiateLoginUri

`func (o *UpdateApplicationRequest) SetInitiateLoginUri(v string)`

SetInitiateLoginUri sets InitiateLoginUri field to given value.

### HasInitiateLoginUri

`func (o *UpdateApplicationRequest) HasInitiateLoginUri() bool`

HasInitiateLoginUri returns a boolean if a field has been set.

### GetPkceEnforcement

`func (o *UpdateApplicationRequest) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *UpdateApplicationRequest) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *UpdateApplicationRequest) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *UpdateApplicationRequest) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *UpdateApplicationRequest) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *UpdateApplicationRequest) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *UpdateApplicationRequest) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *UpdateApplicationRequest) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *UpdateApplicationRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *UpdateApplicationRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *UpdateApplicationRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *UpdateApplicationRequest) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *UpdateApplicationRequest) GetRefreshTokenDuration() int32`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *UpdateApplicationRequest) GetRefreshTokenDurationOk() (*int32, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *UpdateApplicationRequest) SetRefreshTokenDuration(v int32)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *UpdateApplicationRequest) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingDuration

`func (o *UpdateApplicationRequest) GetRefreshTokenRollingDuration() int32`

GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingDurationOk

`func (o *UpdateApplicationRequest) GetRefreshTokenRollingDurationOk() (*int32, bool)`

GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingDuration

`func (o *UpdateApplicationRequest) SetRefreshTokenRollingDuration(v int32)`

SetRefreshTokenRollingDuration sets RefreshTokenRollingDuration field to given value.

### HasRefreshTokenRollingDuration

`func (o *UpdateApplicationRequest) HasRefreshTokenRollingDuration() bool`

HasRefreshTokenRollingDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingGracePeriodDuration

`func (o *UpdateApplicationRequest) GetRefreshTokenRollingGracePeriodDuration() int32`

GetRefreshTokenRollingGracePeriodDuration returns the RefreshTokenRollingGracePeriodDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingGracePeriodDurationOk

`func (o *UpdateApplicationRequest) GetRefreshTokenRollingGracePeriodDurationOk() (*int32, bool)`

GetRefreshTokenRollingGracePeriodDurationOk returns a tuple with the RefreshTokenRollingGracePeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingGracePeriodDuration

`func (o *UpdateApplicationRequest) SetRefreshTokenRollingGracePeriodDuration(v int32)`

SetRefreshTokenRollingGracePeriodDuration sets RefreshTokenRollingGracePeriodDuration field to given value.

### HasRefreshTokenRollingGracePeriodDuration

`func (o *UpdateApplicationRequest) HasRefreshTokenRollingGracePeriodDuration() bool`

HasRefreshTokenRollingGracePeriodDuration returns a boolean if a field has been set.

### GetResponseTypes

`func (o *UpdateApplicationRequest) GetResponseTypes() []EnumApplicationOIDCResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *UpdateApplicationRequest) GetResponseTypesOk() (*[]EnumApplicationOIDCResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *UpdateApplicationRequest) SetResponseTypes(v []EnumApplicationOIDCResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *UpdateApplicationRequest) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetRequireSignedRequestObject

`func (o *UpdateApplicationRequest) GetRequireSignedRequestObject() bool`

GetRequireSignedRequestObject returns the RequireSignedRequestObject field if non-nil, zero value otherwise.

### GetRequireSignedRequestObjectOk

`func (o *UpdateApplicationRequest) GetRequireSignedRequestObjectOk() (*bool, bool)`

GetRequireSignedRequestObjectOk returns a tuple with the RequireSignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignedRequestObject

`func (o *UpdateApplicationRequest) SetRequireSignedRequestObject(v bool)`

SetRequireSignedRequestObject sets RequireSignedRequestObject field to given value.

### HasRequireSignedRequestObject

`func (o *UpdateApplicationRequest) HasRequireSignedRequestObject() bool`

HasRequireSignedRequestObject returns a boolean if a field has been set.

### GetSupportUnsignedRequestObject

`func (o *UpdateApplicationRequest) GetSupportUnsignedRequestObject() bool`

GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field if non-nil, zero value otherwise.

### GetSupportUnsignedRequestObjectOk

`func (o *UpdateApplicationRequest) GetSupportUnsignedRequestObjectOk() (*bool, bool)`

GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUnsignedRequestObject

`func (o *UpdateApplicationRequest) SetSupportUnsignedRequestObject(v bool)`

SetSupportUnsignedRequestObject sets SupportUnsignedRequestObject field to given value.

### HasSupportUnsignedRequestObject

`func (o *UpdateApplicationRequest) HasSupportUnsignedRequestObject() bool`

HasSupportUnsignedRequestObject returns a boolean if a field has been set.

### GetTags

`func (o *UpdateApplicationRequest) GetTags() []EnumApplicationTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateApplicationRequest) GetTagsOk() (*[]EnumApplicationTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateApplicationRequest) SetTags(v []EnumApplicationTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateApplicationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetLinkUri

`func (o *UpdateApplicationRequest) GetTargetLinkUri() string`

GetTargetLinkUri returns the TargetLinkUri field if non-nil, zero value otherwise.

### GetTargetLinkUriOk

`func (o *UpdateApplicationRequest) GetTargetLinkUriOk() (*string, bool)`

GetTargetLinkUriOk returns a tuple with the TargetLinkUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLinkUri

`func (o *UpdateApplicationRequest) SetTargetLinkUri(v string)`

SetTargetLinkUri sets TargetLinkUri field to given value.

### HasTargetLinkUri

`func (o *UpdateApplicationRequest) HasTargetLinkUri() bool`

HasTargetLinkUri returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *UpdateApplicationRequest) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *UpdateApplicationRequest) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *UpdateApplicationRequest) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetParRequirement

`func (o *UpdateApplicationRequest) GetParRequirement() EnumApplicationOIDCPARRequirement`

GetParRequirement returns the ParRequirement field if non-nil, zero value otherwise.

### GetParRequirementOk

`func (o *UpdateApplicationRequest) GetParRequirementOk() (*EnumApplicationOIDCPARRequirement, bool)`

GetParRequirementOk returns a tuple with the ParRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParRequirement

`func (o *UpdateApplicationRequest) SetParRequirement(v EnumApplicationOIDCPARRequirement)`

SetParRequirement sets ParRequirement field to given value.

### HasParRequirement

`func (o *UpdateApplicationRequest) HasParRequirement() bool`

HasParRequirement returns a boolean if a field has been set.

### GetParTimeout

`func (o *UpdateApplicationRequest) GetParTimeout() int32`

GetParTimeout returns the ParTimeout field if non-nil, zero value otherwise.

### GetParTimeoutOk

`func (o *UpdateApplicationRequest) GetParTimeoutOk() (*int32, bool)`

GetParTimeoutOk returns a tuple with the ParTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParTimeout

`func (o *UpdateApplicationRequest) SetParTimeout(v int32)`

SetParTimeout sets ParTimeout field to given value.

### HasParTimeout

`func (o *UpdateApplicationRequest) HasParTimeout() bool`

HasParTimeout returns a boolean if a field has been set.

### GetSigning

`func (o *UpdateApplicationRequest) GetSigning() ApplicationOIDCAllOfSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *UpdateApplicationRequest) GetSigningOk() (*ApplicationOIDCAllOfSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *UpdateApplicationRequest) SetSigning(v ApplicationOIDCAllOfSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *UpdateApplicationRequest) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetAudienceRestriction

`func (o *UpdateApplicationRequest) GetAudienceRestriction() string`

GetAudienceRestriction returns the AudienceRestriction field if non-nil, zero value otherwise.

### GetAudienceRestrictionOk

`func (o *UpdateApplicationRequest) GetAudienceRestrictionOk() (*string, bool)`

GetAudienceRestrictionOk returns a tuple with the AudienceRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceRestriction

`func (o *UpdateApplicationRequest) SetAudienceRestriction(v string)`

SetAudienceRestriction sets AudienceRestriction field to given value.

### HasAudienceRestriction

`func (o *UpdateApplicationRequest) HasAudienceRestriction() bool`

HasAudienceRestriction returns a boolean if a field has been set.

### GetDomainName

`func (o *UpdateApplicationRequest) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateApplicationRequest) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateApplicationRequest) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetReplyUrl

`func (o *UpdateApplicationRequest) GetReplyUrl() string`

GetReplyUrl returns the ReplyUrl field if non-nil, zero value otherwise.

### GetReplyUrlOk

`func (o *UpdateApplicationRequest) GetReplyUrlOk() (*string, bool)`

GetReplyUrlOk returns a tuple with the ReplyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyUrl

`func (o *UpdateApplicationRequest) SetReplyUrl(v string)`

SetReplyUrl sets ReplyUrl field to given value.


### GetSubjectNameIdentifierFormat

`func (o *UpdateApplicationRequest) GetSubjectNameIdentifierFormat() EnumApplicationWSFEDSubjectNameIdentifierFormat`

GetSubjectNameIdentifierFormat returns the SubjectNameIdentifierFormat field if non-nil, zero value otherwise.

### GetSubjectNameIdentifierFormatOk

`func (o *UpdateApplicationRequest) GetSubjectNameIdentifierFormatOk() (*EnumApplicationWSFEDSubjectNameIdentifierFormat, bool)`

GetSubjectNameIdentifierFormatOk returns a tuple with the SubjectNameIdentifierFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectNameIdentifierFormat

`func (o *UpdateApplicationRequest) SetSubjectNameIdentifierFormat(v EnumApplicationWSFEDSubjectNameIdentifierFormat)`

SetSubjectNameIdentifierFormat sets SubjectNameIdentifierFormat field to given value.

### HasSubjectNameIdentifierFormat

`func (o *UpdateApplicationRequest) HasSubjectNameIdentifierFormat() bool`

HasSubjectNameIdentifierFormat returns a boolean if a field has been set.

### GetApplyDefaultTheme

`func (o *UpdateApplicationRequest) GetApplyDefaultTheme() bool`

GetApplyDefaultTheme returns the ApplyDefaultTheme field if non-nil, zero value otherwise.

### GetApplyDefaultThemeOk

`func (o *UpdateApplicationRequest) GetApplyDefaultThemeOk() (*bool, bool)`

GetApplyDefaultThemeOk returns a tuple with the ApplyDefaultTheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyDefaultTheme

`func (o *UpdateApplicationRequest) SetApplyDefaultTheme(v bool)`

SetApplyDefaultTheme sets ApplyDefaultTheme field to given value.


### GetEnableDefaultThemeFooter

`func (o *UpdateApplicationRequest) GetEnableDefaultThemeFooter() bool`

GetEnableDefaultThemeFooter returns the EnableDefaultThemeFooter field if non-nil, zero value otherwise.

### GetEnableDefaultThemeFooterOk

`func (o *UpdateApplicationRequest) GetEnableDefaultThemeFooterOk() (*bool, bool)`

GetEnableDefaultThemeFooterOk returns a tuple with the EnableDefaultThemeFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDefaultThemeFooter

`func (o *UpdateApplicationRequest) SetEnableDefaultThemeFooter(v bool)`

SetEnableDefaultThemeFooter sets EnableDefaultThemeFooter field to given value.

### HasEnableDefaultThemeFooter

`func (o *UpdateApplicationRequest) HasEnableDefaultThemeFooter() bool`

HasEnableDefaultThemeFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


