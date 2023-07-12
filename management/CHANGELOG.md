# v0.24.0 (Unreleased)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* **Bug** Corrected data type for `CreatedAt`, `UpdatedAt` on the `User` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected data type for `LockedAt`, `UnlockAt` on the `UserAccount` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected read-only status for `LockedAt`, `SecondsUntilUnlock` and `UnlockAt` on the `UserAccount` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected required status for `CanAuthenticate` and `Status` on the `UserAccount` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected data type for `At` on the `UserLastSignOn` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected data type for `Type` on the `UserPasswordExternalGateway` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Bug** Corrected required status for `Href` on the `UserPhoto` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Enhancement** Added read-only `EmailVerified` to the `User` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Enhancement** Added optional `SuppressVerificationCode` to the `UserLifecycle` object. [#221](https://github.com/patrickcping/pingone-go-sdk-v2/pull/221)
* **Enhancement** Updated the `KeyRotationPolicy` model such that `ValidityPeriod` is now optional and has a default value, and `RotationPeriod` has a default value. [#220](https://github.com/patrickcping/pingone-go-sdk-v2/pull/220)

# v0.23.0 (2023-07-04)

* **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
* **Bug** Fixed Phone Notification Settings `POST` and `PUT` request payload data model. [#195](https://github.com/patrickcping/pingone-go-sdk-v2/pull/195)
* **Bug** Corrected `Requests` from object to array in the `NotificationsSettingsPhoneDeliverySettingsCustom` object. [#195](https://github.com/patrickcping/pingone-go-sdk-v2/pull/195)
* **Bug** Corrected `Name` as required property of `NotificationsSettingsPhoneDeliverySettingsCustom` object. [#195](https://github.com/patrickcping/pingone-go-sdk-v2/pull/195)
* **Bug** Corrected `DeleteNotificationsSettings` response code and payload. [#203](https://github.com/patrickcping/pingone-go-sdk-v2/pull/203)
* **Bug** Corrected `UpdateNotificationsSettings` request payload. [#203](https://github.com/patrickcping/pingone-go-sdk-v2/pull/203)
* **Enhancement** Added `Numbers` array to the `NotificationsSettingsPhoneDeliverySettingsCustom` object. [#195](https://github.com/patrickcping/pingone-go-sdk-v2/pull/195)
* **Enhancement** Added `PhoneDeliverySettings` array to the `EntityArray` object. [#195](https://github.com/patrickcping/pingone-go-sdk-v2/pull/195)
* **Enhancement** Added `Environment`, `DeliveryMode` and `Whitelist` attributes to the `NotificationsSettings` object. [#203](https://github.com/patrickcping/pingone-go-sdk-v2/pull/203)
* **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# v0.22.0 (2023-05-30)

* **Enhancement** Added `SloWindow` optional attribute to the Application SAML objects. [#179](https://github.com/patrickcping/pingone-go-sdk-v2/pull/179)
* **Enhancement** Added SLO optional attributes to the SAML External Identity Provider object. [#179](https://github.com/patrickcping/pingone-go-sdk-v2/pull/179)
* **Enhancement** Added `NewUserProvisioning` to the `SignOnPolicyActionLogin` model. [#181](https://github.com/patrickcping/pingone-go-sdk-v2/pull/181)

# v0.21.0 (2023-05-22)

* **Note** Deprecated `bundleId` and `packageName` at the root level of the `ApplicationOIDC` model. Customers should use `mobile.bundleId` and `mobile.packageName` going forward. [#172](https://github.com/patrickcping/pingone-go-sdk-v2/pull/172)
* **Feature** Support for Forms. [#176](https://github.com/patrickcping/pingone-go-sdk-v2/pull/176)
* **Feature** Support for Forms Recaptcha configuration. [#176](https://github.com/patrickcping/pingone-go-sdk-v2/pull/176)
* **Enhancement** Added `filterOptions.ipAddressExposed` and `filterOptions.userAgentExposed` to the `Subscription` (webhook) data model. [#173](https://github.com/patrickcping/pingone-go-sdk-v2/pull/173)

# v0.20.0 (2023-05-19)

* **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
* **Enhancement** Add support for enumerated values and regex validation to the schema attribute model. [#161](https://github.com/patrickcping/pingone-go-sdk-v2/pull/161)
* **Enhancement** Add support for setting `Email` quotas in notification policies. [#162](https://github.com/patrickcping/pingone-go-sdk-v2/pull/162)

# v0.19.1 (2023-04-28)

* **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)

# v0.19.0 (2023-04-24)

* **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* **Enhancement** Added support for the Google Play Integrity API, `GooglePlay` for the `ApplicationOIDC` object model.  *Breaking change note*: This is now a required attribute when configuring integrity detection on Android devices. [#153](https://github.com/patrickcping/pingone-go-sdk-v2/pull/153)

# v0.18.0 (2023-04-18)

* **Enhancement** Add `CustomCRL` to the `Certificate` data model [#136](https://github.com/patrickcping/pingone-go-sdk-v2/pull/136)
* **Enhancement** Add notifications policy country limit attributes for `NotificationsPolicy` model. [#142](https://github.com/patrickcping/pingone-go-sdk-v2/pull/142)
* **Enhancement** Expand the `ApplicationAttributeMapping` model for attribute scoping. [#143](https://github.com/patrickcping/pingone-go-sdk-v2/pull/143)

# v0.17.1 (2023-03-20)

* **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)

# v0.17.0 (2023-03-20)

* **Note** Added `FlowPolicy` and `FlowPolicyAssignment` data model descriptions. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
* **Bug fix** `priority` attribute in the `FlowPolicyAssignment` data model corrected to be a required field, with a minimum value of `1`. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
* **Bug fix** `flowPolicy` attribute in the `FlowPolicyAssignment` data model corrected to the `FlowPolicyAssignmentFlowPolicy` data model. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
* **Bug fix** `idpSigning.algorithm` attribute in the `ApplicationSAML` data model corrected to be writable, set as ENUM.

# v0.16.0 (2023-03-10)

* **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)
* **Breaking change** Agreement revision content type changed to an enum `EnumAgreementRevisionContentType`. [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
* **Breaking change** Required attributes now added to the `AgreementLanguageRevision` data model [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
* **Feature** Support for Key Rotation Policies [#123](https://github.com/patrickcping/pingone-go-sdk-v2/pull/123)
* **Bug fix** `reconsentPeriodDays` attribute in the `Agreement` data model corrected to be a floating point number, as in the API. [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
* **Bug fix** `requireReconsent` attribute corrected in `AgreementLanguageRevision` data model. [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
* **Bug fix** Fix the `PUT` request payload to be the `AgreementLanguage` data model [#118](https://github.com/patrickcping/pingone-go-sdk-v2/pull/118)
* **Bug fix** Corrected the RFC3339 date time string format for `Agreement.ConsentCountsUpdatedAt`, `EntityArrayEmbeddedLanguagesInner.CreatedAt`, `EntityArrayEmbeddedLanguagesInner.UpdatedAt`, `Image.CreatedAt`, `Language.CreatedAt`, `Language.UpdatedAt`, `LanguageLocalizationStatus.CreatedAt`, `LanguageLocalizationStatus.UpdatedAt`, `License.BeginsAt`, `License.ExpiresAt` and `License.TerminatesAt` [#119](https://github.com/patrickcping/pingone-go-sdk-v2/pull/119)
* **Bug fix** Corrected `Agreement.consentCountsUpdatedAt` => `Agreement.consentsAggregatedAt` [#120](https://github.com/patrickcping/pingone-go-sdk-v2/pull/120)
* **Bug fix** Corrected `Agreement.expiredUserConsents` => `Agreement.totalExpiredConsents` [#120](https://github.com/patrickcping/pingone-go-sdk-v2/pull/120)
* **Bug fix** Corrected `Agreement.totalUserConsents` => `Agreement.totalConsents` [#120](https://github.com/patrickcping/pingone-go-sdk-v2/pull/120)
* **Enhancement** Add support for `ApplicationPingOneAdminConsole` model for applications [#114](https://github.com/patrickcping/pingone-go-sdk-v2/pull/114)
* **Enhancement** Add support for Flow policies and application flow policy assignments [#115](https://github.com/patrickcping/pingone-go-sdk-v2/pull/115)
* **Enhancement** `effectiveAt` and `notValidAfter` attributes on the `AgreementLanguageRevision` changed to be date/time format. [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
* **Enhancement** Add `INTERNAL` value to the `EnumOrganizationType` enum. [#124](https://github.com/patrickcping/pingone-go-sdk-v2/pull/124)
* **Enhancement** Add `default` and `environment` attributes to the `Population` data model.

# v0.15.0 (2023-02-22)

* **Note** bump `golang.org/x/net` v0.6.0 => v0.7.0 [#103](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* **Note** bump `golang.org/x/oauth2` v0.4.0 => v0.5.0 [#102](https://github.com/patrickcping/pingone-go-sdk-v2/pull/102)
* **Note** Handle file close errors in certificate management, images and password policies API [#111](https://github.com/patrickcping/pingone-go-sdk-v2/pull/111)
* **Breaking change** `OrganizationsApi.ReadOneOrganizations` changed to `OrganizationsApi.ReadOneOrganization` [#102](https://github.com/patrickcping/pingone-go-sdk-v2/pull/102)
* **Breaking change** `EnumGatewayLDAPSecurity` changed to `EnumGatewayTypeLDAPSecurity` [#107](https://github.com/patrickcping/pingone-go-sdk-v2/pull/107)
* **Breaking change** `GatewayLDAP` changed to `GatewayTypeLDAP` [#107](https://github.com/patrickcping/pingone-go-sdk-v2/pull/107)
* **Breaking change** Make `Address` a required attribute on the `NotificationsSettingsEmailDeliverySettingsFrom` model [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
* **Enhancement** Add `limit` parameter to `OrganizationsApi.ReadAllOrganizations` [#102](https://github.com/patrickcping/pingone-go-sdk-v2/pull/102)
* **Enhancement** Add support for Huawei HMS Push service [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* **Enhancement** Add support for RADIUS gateways [#107](https://github.com/patrickcping/pingone-go-sdk-v2/pull/107)
* **Enhancement** Add `Protocol` and `Password` attributes to the `NotificationsSettingsEmailDeliverySettings` model [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
* **Bug fix** - Correct the `EmailDomainTrustedEmail` API model [#103](https://github.com/patrickcping/pingone-go-sdk-v2/pull/103)
* **Bug fix** - Correct the "UPDATE Notifications Policy" function name (fix the operation ID) [#108](https://github.com/patrickcping/pingone-go-sdk-v2/pull/108)
* **Bug fix** - Add the `DeleteEmailDeliverySettings` API call back in to the `NotificationsSettingsSMTPApi` API [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)

# v0.14.0 (2023-01-12)

* **Bug fix** - Correct the `TemplateContent` API model [#97](https://github.com/patrickcping/pingone-go-sdk-v2/pull/97)

# v0.13.0 (2023-01-09)

* **Breaking change** Moved `AssignActorRoles` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Breaking change** Moved `Tags` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Breaking change** Moved `SupportUnsignedRequestObject` from all `Application` models to just `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)
* **Feature** Support for Notifications Settings [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Feature** Support for Notifications Policies [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Feature** Support for Notifications Templates and Contents [#85](https://github.com/patrickcping/pingone-go-sdk-v2/pull/85)
* **Enhancement** Add support for the WS-Federation application type [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `HiddenFromAppPortal` property on the `Application` models [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `AllowWildcardInRedirectUris` property on the `ApplicationOIDC` model [#96](https://github.com/patrickcping/pingone-go-sdk-v2/pull/96)
* **Enhancement** Add support for `InitiateLoginUri` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `RefreshTokenRollingGracePeriodDuration` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `TargetLinkUri` property on the `ApplicationOIDC` model [#91](https://github.com/patrickcping/pingone-go-sdk-v2/pull/91)
* **Enhancement** Add support for `HomePageUrl` property on the `ApplicationSAML` model [#96](https://github.com/patrickcping/pingone-go-sdk-v2/pull/96)
* **Enhancement** Add boolean data type support to the Sign On Policy `Equals` Condition [#93](https://github.com/patrickcping/pingone-go-sdk-v2/pull/93)

# v0.12.0 (2022-11-06)

* **Breaking change** Removed the `EnumLicensePackage` enum model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)
* **Feature** Support for Branding Settings [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
* **Feature** Support for Branding Themes [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
* **Feature** Support for Resource Client Secrets [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `idToken` and `userInfo` attributes to the `ResourceAttribute` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `mappedClaims` attributes to the `ResourceScope` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Add `introspectEndpointAuthMethod` attributes to the `Resource` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
* **Enhancement** Added the `TERMINATED` value to the `EnumLicenseStatus` model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)

# v0.11.2 (2022-10-15)

* **Bug fix** - Correct the `Image` API model [#74](https://github.com/patrickcping/pingone-go-sdk-v2/pull/74)

# v0.11.1 (2022-10-10)

* **Enhancement** Add `uriPrefix` (mobile) and `excludedPlatforms` (mobile device integrity check) to the `ApplicationOIDC` data model [#71](https://github.com/patrickcping/pingone-go-sdk-v2/pull/71)

# v0.11.0 (2022-10-09)

* **Feature** Support for Licenses [#64](https://github.com/patrickcping/pingone-go-sdk-v2/pull/64)
* **Feature** Support for Languages [#63](https://github.com/patrickcping/pingone-go-sdk-v2/pull/63)

# v0.10.0 (2022-09-18)

* **Bug fix** - Correct the http endpoint headers object in the `Subscription` model [#52](https://github.com/patrickcping/pingone-go-sdk-v2/pull/52)
* **Bug fix** - Correct the `createGroupNesting` response object [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Breaking change** `CreateApplication201Response` changes to `ReadOneApplication200Response` model [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for External Link applications [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for the PingOne Portal application [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Support for the PingOne Self Service application [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Feature** Add the `readOneGroupNesting` API [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Feature** Support for Alerts [#54](https://github.com/patrickcping/pingone-go-sdk-v2/pull/54)
* **Enhancement** Add `type` attribute to the `GroupNesting` data model [#53](https://github.com/patrickcping/pingone-go-sdk-v2/pull/53)
* **Enhancement** Support for `kerberos` attribute model in the `Application` data model [#49](https://github.com/patrickcping/pingone-go-sdk-v2/pull/49)
* **Enhancement** - Add `PINGID_WINLOGIN_PASSWORDLESS_AUTHENTICATION` and `PINGID_AUTHENTICATION` to the `EnumSignOnPolicyType` model for workforce based sign on policy actions [#47](https://github.com/patrickcping/pingone-go-sdk-v2/pull/47)
* **Enhancement** - Added `SignOnPolicyActionPingIDWinLoginPasswordless` model to support the PingID Windows Passwordless sign on policy action [#47](https://github.com/patrickcping/pingone-go-sdk-v2/pull/47)
* **Enhancement** - Added `application` attribute to the `ApplicationAttributeMapping` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateApplication201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateApplicationRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateGateway201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateGatewayRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedApplicationsInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedAttributesInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedGatewaysInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `IdentityProvider` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `SignOnPolicyAction` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `SignOnPolicyActionCommonConditionOrOrInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateApplicationRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateDomain200Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)

# v0.9.0 (2022-09-11)

* **Bug fix** - Correct the `CreateDomain` API response object [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct the `UpdateDomain` API response object [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct the `CustomDomainCertificate` conflict by adding `CustomDomainCertificateRequest` model [#45](https://github.com/patrickcping/pingone-go-sdk-v2/pull/45)
* **Bug fix** - Correct `Gateway` model attributes [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Bug fix** - Correct `GatewayLDAP` model attributes [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Bug fix** - Correct the `EnumGatewayPasswordAuthority` values [#46](https://github.com/patrickcping/pingone-go-sdk-v2/pull/46)
* **Feature** Support for Subscriptions data model [#42](https://github.com/patrickcping/pingone-go-sdk-v2/pull/42)
* **Feature** Support for `EmailDomain` and `EmailDomainTrustedEmail` data models [#43](https://github.com/patrickcping/pingone-go-sdk-v2/pull/43)
* **Breaking change** `createdAt` and `updatedAt` attributes changed to date/time data type [#42](https://github.com/patrickcping/pingone-go-sdk-v2/pull/42)
* **Breaking change** `lastUsedAt` attribute on the Gateway Credential model changed to date/time data type [#43](https://github.com/patrickcping/pingone-go-sdk-v2/pull/43)
* **Breaking change** Remove Device Authentication Policy API (moved to MFA module) [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)

# v0.8.0 (2022-09-04)

* **Bug fix** - Fix `Gateway` response dereferencing for non-`LDAP` types [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Bug fix** - Fix `ExportCSR` headers, to determine the response type of the CSR [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `ImportCSRResponse` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `GetKey` API when exporting the public certificate [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Bug fix** - Correct the `CreateCertificateFromFile` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Enhancement** - Add `Links` to `Gateway` model [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Enhancement** - Add `X_PEM_FILE` to `EnumCSRExportHeader` (exporting CSR formats) [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
* **Feature** Support for Read All Gateway Credentials [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
* **Feature** Support for Read One Gateway Credential [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)

# v0.7.0 (2022-09-01)

* **Bug fix** - Correct `ApplicationSAML` model `IdpSigning` attribute [#38](https://github.com/patrickcping/pingone-go-sdk-v2/pull/38)
* **Bug fix** - Correct `ApplicationSAML` model to include read only `Algorithm` attribute [#38](https://github.com/patrickcping/pingone-go-sdk-v2/pull/38)
* **Note** Uplift API version to 2022-08-02 [#32](https://github.com/patrickcping/pingone-go-sdk-v2/pull/32)
* **Note** Documentation formatting change `Certificate` model [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyAlgorithm` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyStatus` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeySignagureAlgorithm` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Note** Documentation formatting change `EnumCertificateKeyUsageType` [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Breaking change** Added required fields to the SAML Identity Provider constructor [#31](https://github.com/patrickcping/pingone-go-sdk-v2/pull/31)
* **Breaking change** Changed required attributes on the `Certificate` model [#35](https://github.com/patrickcping/pingone-go-sdk-v2/pull/35)
* **Enhancement** Support for Kerberos Gateway [#32](https://github.com/patrickcping/pingone-go-sdk-v2/pull/32)
* **Enhancement** `EnumCertificateKeyUsageType` includes missing enum values [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33)
* **Enhancement** `EnumCertificateKeyStatus` includes missing enum values [#37](https://github.com/patrickcping/pingone-go-sdk-v2/pull/37)
* **Enhancement** `EnumCertificateKeySignagureAlgorithm` includes missing enum values [#33](https://github.com/patrickcping/pingone-go-sdk-v2/pull/33), [#36](https://github.com/patrickcping/pingone-go-sdk-v2/pull/36)
* **Enhancement** Better define the Certificate Key update model [#34](https://github.com/patrickcping/pingone-go-sdk-v2/pull/34)

# v0.6.0 (2022-08-29)

* **Bug fix** - Correct `PINGONE_API` from incorrect `PING_ONE_API` in `EnumResourceType` enum [#24](https://github.com/patrickcping/pingone-go-sdk-v2/pull/24)
* **Bug fix** - Correct `EnumIdentityProviderAttributeMappingUpdate` values [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Bug fix** - Fix `IdentityProvider` oneOf decode routine [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Feature** Support for Custom Domains [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Feature** Support for Keypairs [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Feature** Support for Certificates [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
* **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)
* **Breaking change** API name change for Agreement Languages [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Agreement Revisions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Agreements [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Attribute Mapping [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Role Assignment [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Secret [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Application Sign on Policy Assignment [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Applications [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Branding Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Branding Themes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Enable Users [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway Credentials [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway Instances [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateway role assignments [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Gateways [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Group Memberships [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Identity Provider Attributes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Identity Providers [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Language localization [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Languages [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Linked Accounts [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for MFA Pairing Keys [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Settings SMTP [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Notification Templates [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Phone Delivery Settings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Mappings [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Plans [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Revisions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Rules [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Store Metadata [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Propagation Stores [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resource Attributes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resource Scopes [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Resources [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sessions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sign On Policies [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Sign On Policy Actions [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Trusted Email Addresses [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Trusted Email Domains [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Accounts [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Agreement Consents [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User ID Verification [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Passwords [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Populations [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for User Role Assignments [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)
* **Breaking change** API name change for Users [#30](https://github.com/patrickcping/pingone-go-sdk-v2/pull/30)

# v0.5.0 (2022-08-17)

* **Bug fix** - `SignOnPolicyActionProgressiveProfiling` object includes fixed `attributes` attribute, from object type to array type
* **Bug fix** - Restructure `SignOnPolicyAction` model to properly separate policy action attributes from each other
* **Bug fix** - Restructure `SignOnPolicyActionMFA` model to correct `noDevicesMode` attribute (fixed incorrect `noDeviceMode`)
* **Bug fix** - Corrections to `SignOnPolicyActionCommonConditions` model
* **Bug fix** - Moved `confirmIdentityProviderAttributes` boolean to the registration sub model of `SignOnPolicyAction` 

# v0.4.0 (2022-08-10)

* **Enhancement** Added generic `_links` to `Application` model
* **Breaking change** `id` made required in `spVerification.certificates` of the `ApplicationSAML` model
* **Enhancement** Add `mobile.passcodeRefreshDuration` to `ApplicationOIDC` model
* **Breaking change** `accessControl.role.type` made an enum in the `Application` model
* **Enhancement** Added `PING_ONE_DAVINCI` product type

# v0.3.0 (2022-08-05)

* **Enhancement** - *API update 2022-07-18* - `SignOnPolicyAction` object includes new attribute `DeviceAuthenticationPolicy` to replace deprecated `Applications`, `Voice`, `Sms`, `SecurityKey`, `Email`, `BoundBiometrics` and `Authenticator` attributes [#11](https://github.com/patrickcping/pingone-go-sdk-v2/pull/11)
* **Feature** Agreements, Agreement Languages and Agreement Revisions [#14](https://github.com/patrickcping/pingone-go-sdk-v2/pull/14)
* **Feature** Support for Identity Provider models [#15](https://github.com/patrickcping/pingone-go-sdk-v2/pull/15)
* **Bug** fix for `PasswordPolicyMinCharacters` special character issue [#17](https://github.com/patrickcping/pingone-go-sdk-v2/pull/17)

# v0.2.0 (2022-07-17)

* **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
* **Feature** Full support for enum values

# v0.1.0 (2022-07-16)

Initial release - rebasing versions to reflect module stability