# UpdateApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | Pointer to [**ApplicationAccessControl**](ApplicationAccessControl.md) |  | [optional] 
**AssignActorRoles** | Pointer to **bool** | A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request. | [optional] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the application. | [optional] 
**Enabled** | **bool** | A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Icon** | Pointer to [**ApplicationIcon**](ApplicationIcon.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the application ID. | [optional] [readonly] 
**LoginPageUrl** | Pointer to **string** | A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains. | [optional] 
**Name** | **string** | A string that specifies the name of the application. This is a required property. | 
**Protocol** | [**EnumApplicationProtocol**](EnumApplicationProtocol.md) |  | 
**Tags** | Pointer to [**[]EnumApplicationTags**](EnumApplicationTags.md) | An array that specifies the list of labels associated with the application. Options are PING_FED_CONNECTION_INTEGRATION. | [optional] 
**Type** | [**EnumApplicationType**](EnumApplicationType.md) |  | 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**SupportUnsignedRequestObject** | Pointer to **bool** | A boolean that specifies whether the request query parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed. | [optional] 
**AcsUrls** | **[]string** | A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property. | 
**AssertionDuration** | **int32** | An integer that specifies the assertion validity duration in seconds. This is a required property. | 
**AssertionSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion itself should be signed. The default value is true. | [optional] 
**IdpSigningtype** | Pointer to [**ApplicationSAMLAllOfIdpSigningtype**](ApplicationSAMLAllOfIdpSigningtype.md) |  | [optional] 
**NameIdFormat** | Pointer to **string** | A string that specifies the format of the Subject NameID attibute in the SAML assertion | [optional] 
**ResponseSigned** | Pointer to **bool** | A boolean that specifies whether the SAML assertion response itself should be signed. The default value is False. | [optional] 
**SloBinding** | Pointer to [**EnumApplicationSAMLSloBinding**](EnumApplicationSAMLSloBinding.md) |  | [optional] 
**SloEndpoint** | Pointer to **string** | A string that specifies the logout endpoint URL. This is an optional property. However, if a sloEndpoint logout endpoint URL is not defined, logout actions result in an error. | [optional] 
**SloResponseEndpoint** | Pointer to **string** | A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response. | [optional] 
**SpEntityId** | **string** | A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment. | 
**SpVerification** | Pointer to [**ApplicationSAMLAllOfSpVerification**](ApplicationSAMLAllOfSpVerification.md) |  | [optional] 
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

### NewUpdateApplicationRequest

`func NewUpdateApplicationRequest(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, acsUrls []string, assertionDuration int32, spEntityId string, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, ) *UpdateApplicationRequest`

NewUpdateApplicationRequest instantiates a new UpdateApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationRequestWithDefaults

`func NewUpdateApplicationRequestWithDefaults() *UpdateApplicationRequest`

NewUpdateApplicationRequestWithDefaults instantiates a new UpdateApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetCreatedAt

`func (o *UpdateApplicationRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateApplicationRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateApplicationRequest) SetCreatedAt(v string)`

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

`func (o *UpdateApplicationRequest) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateApplicationRequest) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateApplicationRequest) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateApplicationRequest) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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

### GetIdpSigningtype

`func (o *UpdateApplicationRequest) GetIdpSigningtype() ApplicationSAMLAllOfIdpSigningtype`

GetIdpSigningtype returns the IdpSigningtype field if non-nil, zero value otherwise.

### GetIdpSigningtypeOk

`func (o *UpdateApplicationRequest) GetIdpSigningtypeOk() (*ApplicationSAMLAllOfIdpSigningtype, bool)`

GetIdpSigningtypeOk returns a tuple with the IdpSigningtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSigningtype

`func (o *UpdateApplicationRequest) SetIdpSigningtype(v ApplicationSAMLAllOfIdpSigningtype)`

SetIdpSigningtype sets IdpSigningtype field to given value.

### HasIdpSigningtype

`func (o *UpdateApplicationRequest) HasIdpSigningtype() bool`

HasIdpSigningtype returns a boolean if a field has been set.

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

### HasHomePageUrl

`func (o *UpdateApplicationRequest) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


