# EntityArrayEmbeddedApplicationsInner

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

### NewEntityArrayEmbeddedApplicationsInner

`func NewEntityArrayEmbeddedApplicationsInner(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType, acsUrls []string, assertionDuration int32, spEntityId string, grantTypes []EnumApplicationOIDCGrantType, tokenEndpointAuthMethod EnumApplicationOIDCTokenAuthMethod, ) *EntityArrayEmbeddedApplicationsInner`

NewEntityArrayEmbeddedApplicationsInner instantiates a new EntityArrayEmbeddedApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedApplicationsInnerWithDefaults

`func NewEntityArrayEmbeddedApplicationsInnerWithDefaults() *EntityArrayEmbeddedApplicationsInner`

NewEntityArrayEmbeddedApplicationsInnerWithDefaults instantiates a new EntityArrayEmbeddedApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *EntityArrayEmbeddedApplicationsInner) GetAccessControl() ApplicationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetAccessControlOk() (*ApplicationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *EntityArrayEmbeddedApplicationsInner) SetAccessControl(v ApplicationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *EntityArrayEmbeddedApplicationsInner) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetAssignActorRoles

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssignActorRoles() bool`

GetAssignActorRoles returns the AssignActorRoles field if non-nil, zero value otherwise.

### GetAssignActorRolesOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssignActorRolesOk() (*bool, bool)`

GetAssignActorRolesOk returns a tuple with the AssignActorRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignActorRoles

`func (o *EntityArrayEmbeddedApplicationsInner) SetAssignActorRoles(v bool)`

SetAssignActorRoles sets AssignActorRoles field to given value.

### HasAssignActorRoles

`func (o *EntityArrayEmbeddedApplicationsInner) HasAssignActorRoles() bool`

HasAssignActorRoles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *EntityArrayEmbeddedApplicationsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedApplicationsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedApplicationsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EntityArrayEmbeddedApplicationsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntityArrayEmbeddedApplicationsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *EntityArrayEmbeddedApplicationsInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedApplicationsInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedApplicationsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIcon

`func (o *EntityArrayEmbeddedApplicationsInner) GetIcon() ApplicationIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetIconOk() (*ApplicationIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EntityArrayEmbeddedApplicationsInner) SetIcon(v ApplicationIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EntityArrayEmbeddedApplicationsInner) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *EntityArrayEmbeddedApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedApplicationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedApplicationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginPageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) GetLoginPageUrl() string`

GetLoginPageUrl returns the LoginPageUrl field if non-nil, zero value otherwise.

### GetLoginPageUrlOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetLoginPageUrlOk() (*string, bool)`

GetLoginPageUrlOk returns a tuple with the LoginPageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) SetLoginPageUrl(v string)`

SetLoginPageUrl sets LoginPageUrl field to given value.

### HasLoginPageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) HasLoginPageUrl() bool`

HasLoginPageUrl returns a boolean if a field has been set.

### GetName

`func (o *EntityArrayEmbeddedApplicationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedApplicationsInner) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *EntityArrayEmbeddedApplicationsInner) GetProtocol() EnumApplicationProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetProtocolOk() (*EnumApplicationProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *EntityArrayEmbeddedApplicationsInner) SetProtocol(v EnumApplicationProtocol)`

SetProtocol sets Protocol field to given value.


### GetTags

`func (o *EntityArrayEmbeddedApplicationsInner) GetTags() []EnumApplicationTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetTagsOk() (*[]EnumApplicationTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EntityArrayEmbeddedApplicationsInner) SetTags(v []EnumApplicationTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EntityArrayEmbeddedApplicationsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *EntityArrayEmbeddedApplicationsInner) GetType() EnumApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetTypeOk() (*EnumApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityArrayEmbeddedApplicationsInner) SetType(v EnumApplicationType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityArrayEmbeddedApplicationsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSupportUnsignedRequestObject

`func (o *EntityArrayEmbeddedApplicationsInner) GetSupportUnsignedRequestObject() bool`

GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field if non-nil, zero value otherwise.

### GetSupportUnsignedRequestObjectOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSupportUnsignedRequestObjectOk() (*bool, bool)`

GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUnsignedRequestObject

`func (o *EntityArrayEmbeddedApplicationsInner) SetSupportUnsignedRequestObject(v bool)`

SetSupportUnsignedRequestObject sets SupportUnsignedRequestObject field to given value.

### HasSupportUnsignedRequestObject

`func (o *EntityArrayEmbeddedApplicationsInner) HasSupportUnsignedRequestObject() bool`

HasSupportUnsignedRequestObject returns a boolean if a field has been set.

### GetAcsUrls

`func (o *EntityArrayEmbeddedApplicationsInner) GetAcsUrls() []string`

GetAcsUrls returns the AcsUrls field if non-nil, zero value otherwise.

### GetAcsUrlsOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetAcsUrlsOk() (*[]string, bool)`

GetAcsUrlsOk returns a tuple with the AcsUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrls

`func (o *EntityArrayEmbeddedApplicationsInner) SetAcsUrls(v []string)`

SetAcsUrls sets AcsUrls field to given value.


### GetAssertionDuration

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssertionDuration() int32`

GetAssertionDuration returns the AssertionDuration field if non-nil, zero value otherwise.

### GetAssertionDurationOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssertionDurationOk() (*int32, bool)`

GetAssertionDurationOk returns a tuple with the AssertionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionDuration

`func (o *EntityArrayEmbeddedApplicationsInner) SetAssertionDuration(v int32)`

SetAssertionDuration sets AssertionDuration field to given value.


### GetAssertionSigned

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssertionSigned() bool`

GetAssertionSigned returns the AssertionSigned field if non-nil, zero value otherwise.

### GetAssertionSignedOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetAssertionSignedOk() (*bool, bool)`

GetAssertionSignedOk returns a tuple with the AssertionSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionSigned

`func (o *EntityArrayEmbeddedApplicationsInner) SetAssertionSigned(v bool)`

SetAssertionSigned sets AssertionSigned field to given value.

### HasAssertionSigned

`func (o *EntityArrayEmbeddedApplicationsInner) HasAssertionSigned() bool`

HasAssertionSigned returns a boolean if a field has been set.

### GetIdpSigningtype

`func (o *EntityArrayEmbeddedApplicationsInner) GetIdpSigningtype() ApplicationSAMLAllOfIdpSigningtype`

GetIdpSigningtype returns the IdpSigningtype field if non-nil, zero value otherwise.

### GetIdpSigningtypeOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetIdpSigningtypeOk() (*ApplicationSAMLAllOfIdpSigningtype, bool)`

GetIdpSigningtypeOk returns a tuple with the IdpSigningtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSigningtype

`func (o *EntityArrayEmbeddedApplicationsInner) SetIdpSigningtype(v ApplicationSAMLAllOfIdpSigningtype)`

SetIdpSigningtype sets IdpSigningtype field to given value.

### HasIdpSigningtype

`func (o *EntityArrayEmbeddedApplicationsInner) HasIdpSigningtype() bool`

HasIdpSigningtype returns a boolean if a field has been set.

### GetNameIdFormat

`func (o *EntityArrayEmbeddedApplicationsInner) GetNameIdFormat() string`

GetNameIdFormat returns the NameIdFormat field if non-nil, zero value otherwise.

### GetNameIdFormatOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetNameIdFormatOk() (*string, bool)`

GetNameIdFormatOk returns a tuple with the NameIdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIdFormat

`func (o *EntityArrayEmbeddedApplicationsInner) SetNameIdFormat(v string)`

SetNameIdFormat sets NameIdFormat field to given value.

### HasNameIdFormat

`func (o *EntityArrayEmbeddedApplicationsInner) HasNameIdFormat() bool`

HasNameIdFormat returns a boolean if a field has been set.

### GetResponseSigned

`func (o *EntityArrayEmbeddedApplicationsInner) GetResponseSigned() bool`

GetResponseSigned returns the ResponseSigned field if non-nil, zero value otherwise.

### GetResponseSignedOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetResponseSignedOk() (*bool, bool)`

GetResponseSignedOk returns a tuple with the ResponseSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSigned

`func (o *EntityArrayEmbeddedApplicationsInner) SetResponseSigned(v bool)`

SetResponseSigned sets ResponseSigned field to given value.

### HasResponseSigned

`func (o *EntityArrayEmbeddedApplicationsInner) HasResponseSigned() bool`

HasResponseSigned returns a boolean if a field has been set.

### GetSloBinding

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloBinding() EnumApplicationSAMLSloBinding`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloBindingOk() (*EnumApplicationSAMLSloBinding, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *EntityArrayEmbeddedApplicationsInner) SetSloBinding(v EnumApplicationSAMLSloBinding)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *EntityArrayEmbeddedApplicationsInner) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### GetSloEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetSloResponseEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloResponseEndpoint() string`

GetSloResponseEndpoint returns the SloResponseEndpoint field if non-nil, zero value otherwise.

### GetSloResponseEndpointOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSloResponseEndpointOk() (*string, bool)`

GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloResponseEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) SetSloResponseEndpoint(v string)`

SetSloResponseEndpoint sets SloResponseEndpoint field to given value.

### HasSloResponseEndpoint

`func (o *EntityArrayEmbeddedApplicationsInner) HasSloResponseEndpoint() bool`

HasSloResponseEndpoint returns a boolean if a field has been set.

### GetSpEntityId

`func (o *EntityArrayEmbeddedApplicationsInner) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *EntityArrayEmbeddedApplicationsInner) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.


### GetSpVerification

`func (o *EntityArrayEmbeddedApplicationsInner) GetSpVerification() ApplicationSAMLAllOfSpVerification`

GetSpVerification returns the SpVerification field if non-nil, zero value otherwise.

### GetSpVerificationOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool)`

GetSpVerificationOk returns a tuple with the SpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpVerification

`func (o *EntityArrayEmbeddedApplicationsInner) SetSpVerification(v ApplicationSAMLAllOfSpVerification)`

SetSpVerification sets SpVerification field to given value.

### HasSpVerification

`func (o *EntityArrayEmbeddedApplicationsInner) HasSpVerification() bool`

HasSpVerification returns a boolean if a field has been set.

### GetMobile

`func (o *EntityArrayEmbeddedApplicationsInner) GetMobile() ApplicationOIDCAllOfMobile`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetMobileOk() (*ApplicationOIDCAllOfMobile, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *EntityArrayEmbeddedApplicationsInner) SetMobile(v ApplicationOIDCAllOfMobile)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *EntityArrayEmbeddedApplicationsInner) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetBundleId

`func (o *EntityArrayEmbeddedApplicationsInner) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *EntityArrayEmbeddedApplicationsInner) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.

### HasBundleId

`func (o *EntityArrayEmbeddedApplicationsInner) HasBundleId() bool`

HasBundleId returns a boolean if a field has been set.

### GetPackageName

`func (o *EntityArrayEmbeddedApplicationsInner) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *EntityArrayEmbeddedApplicationsInner) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *EntityArrayEmbeddedApplicationsInner) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetGrantTypes

`func (o *EntityArrayEmbeddedApplicationsInner) GetGrantTypes() []EnumApplicationOIDCGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetGrantTypesOk() (*[]EnumApplicationOIDCGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *EntityArrayEmbeddedApplicationsInner) SetGrantTypes(v []EnumApplicationOIDCGrantType)`

SetGrantTypes sets GrantTypes field to given value.


### GetHomePageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *EntityArrayEmbeddedApplicationsInner) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### GetPkceEnforcement

`func (o *EntityArrayEmbeddedApplicationsInner) GetPkceEnforcement() EnumApplicationOIDCPKCEOption`

GetPkceEnforcement returns the PkceEnforcement field if non-nil, zero value otherwise.

### GetPkceEnforcementOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetPkceEnforcementOk() (*EnumApplicationOIDCPKCEOption, bool)`

GetPkceEnforcementOk returns a tuple with the PkceEnforcement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceEnforcement

`func (o *EntityArrayEmbeddedApplicationsInner) SetPkceEnforcement(v EnumApplicationOIDCPKCEOption)`

SetPkceEnforcement sets PkceEnforcement field to given value.

### HasPkceEnforcement

`func (o *EntityArrayEmbeddedApplicationsInner) HasPkceEnforcement() bool`

HasPkceEnforcement returns a boolean if a field has been set.

### GetPostLogoutRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) GetPostLogoutRedirectUris() []string`

GetPostLogoutRedirectUris returns the PostLogoutRedirectUris field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUrisOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetPostLogoutRedirectUrisOk() (*[]string, bool)`

GetPostLogoutRedirectUrisOk returns a tuple with the PostLogoutRedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) SetPostLogoutRedirectUris(v []string)`

SetPostLogoutRedirectUris sets PostLogoutRedirectUris field to given value.

### HasPostLogoutRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) HasPostLogoutRedirectUris() bool`

HasPostLogoutRedirectUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *EntityArrayEmbeddedApplicationsInner) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetRefreshTokenDuration

`func (o *EntityArrayEmbeddedApplicationsInner) GetRefreshTokenDuration() int32`

GetRefreshTokenDuration returns the RefreshTokenDuration field if non-nil, zero value otherwise.

### GetRefreshTokenDurationOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetRefreshTokenDurationOk() (*int32, bool)`

GetRefreshTokenDurationOk returns a tuple with the RefreshTokenDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenDuration

`func (o *EntityArrayEmbeddedApplicationsInner) SetRefreshTokenDuration(v int32)`

SetRefreshTokenDuration sets RefreshTokenDuration field to given value.

### HasRefreshTokenDuration

`func (o *EntityArrayEmbeddedApplicationsInner) HasRefreshTokenDuration() bool`

HasRefreshTokenDuration returns a boolean if a field has been set.

### GetRefreshTokenRollingDuration

`func (o *EntityArrayEmbeddedApplicationsInner) GetRefreshTokenRollingDuration() int32`

GetRefreshTokenRollingDuration returns the RefreshTokenRollingDuration field if non-nil, zero value otherwise.

### GetRefreshTokenRollingDurationOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetRefreshTokenRollingDurationOk() (*int32, bool)`

GetRefreshTokenRollingDurationOk returns a tuple with the RefreshTokenRollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenRollingDuration

`func (o *EntityArrayEmbeddedApplicationsInner) SetRefreshTokenRollingDuration(v int32)`

SetRefreshTokenRollingDuration sets RefreshTokenRollingDuration field to given value.

### HasRefreshTokenRollingDuration

`func (o *EntityArrayEmbeddedApplicationsInner) HasRefreshTokenRollingDuration() bool`

HasRefreshTokenRollingDuration returns a boolean if a field has been set.

### GetResponseTypes

`func (o *EntityArrayEmbeddedApplicationsInner) GetResponseTypes() []EnumApplicationOIDCResponseType`

GetResponseTypes returns the ResponseTypes field if non-nil, zero value otherwise.

### GetResponseTypesOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetResponseTypesOk() (*[]EnumApplicationOIDCResponseType, bool)`

GetResponseTypesOk returns a tuple with the ResponseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTypes

`func (o *EntityArrayEmbeddedApplicationsInner) SetResponseTypes(v []EnumApplicationOIDCResponseType)`

SetResponseTypes sets ResponseTypes field to given value.

### HasResponseTypes

`func (o *EntityArrayEmbeddedApplicationsInner) HasResponseTypes() bool`

HasResponseTypes returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *EntityArrayEmbeddedApplicationsInner) GetTokenEndpointAuthMethod() EnumApplicationOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *EntityArrayEmbeddedApplicationsInner) GetTokenEndpointAuthMethodOk() (*EnumApplicationOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *EntityArrayEmbeddedApplicationsInner) SetTokenEndpointAuthMethod(v EnumApplicationOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


