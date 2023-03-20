# Release (2023-03-20)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.16.0 => v0.17.0 [#132](https://github.com/patrickcping/pingone-go-sdk-v2/pull/132)
    * **Note** bump `github.com/golangci/golangci-lint` from v1.51.2 to v1.52.0 [#129](https://github.com/patrickcping/pingone-go-sdk-v2/pull/129)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.17.0](./management/CHANGELOG.md)
    * **Note** Added `FlowPolicy` and `FlowPolicyAssignment` data model descriptions. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `priority` attribute in the `FlowPolicyAssignment` data model corrected to be a required field, with a minimum value of `1`. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `flowPolicy` attribute in the `FlowPolicyAssignment` data model corrected to the `FlowPolicyAssignmentFlowPolicy` data model. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `idpSigning.algorithm` attribute in the `ApplicationSAML` data model corrected to be writable, set as ENUM.

# Release (2023-03-10)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.0
    * **Note** bump `golang.org/x/oauth2` v0.5.0 => v0.6.0 [#121](https://github.com/patrickcping/pingone-go-sdk-v2/pull/121)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.2 => v0.1.3 [#126](https://github.com/patrickcping/pingone-go-sdk-v2/pull/126)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/management` v0.15.0 => v0.16.0 [#126](https://github.com/patrickcping/pingone-go-sdk-v2/pull/126)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.8.0 => v0.9.0 [#126](https://github.com/patrickcping/pingone-go-sdk-v2/pull/126)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.3.2 => v0.3.3 [#126](https://github.com/patrickcping/pingone-go-sdk-v2/pull/126)
    * **Feature** Support for agreement management endpoint API client [#117](https://github.com/patrickcping/pingone-go-sdk-v2/pull/117)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.1.0](./agreementmanagement/CHANGELOG.md)
    * **Initial release** [#117](https://github.com/patrickcping/pingone-go-sdk-v2/pull/117)
    * **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.3](./authorize/CHANGELOG.md)
    * **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.16.0](./management/CHANGELOG.md)
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
    * **Enhancement** `effectiveAt` and `notValidAfter` attributes on the `AgreementLanguageRevision` changed to be date/time format. [#116](https://github.com/patrickcping/pingone-go-sdk-v2/pull/116)
    * **Enhancement** Add `INTERNAL` value to the `EnumOrganizationType` enum. [#124](https://github.com/patrickcping/pingone-go-sdk-v2/pull/124)
    * **Enhancement** Add support for Flow policies and application flow policy assignments [#115](https://github.com/patrickcping/pingone-go-sdk-v2/pull/115)
    * **Enhancement** Add `default` and `environment` attributes to the `Population` data model.
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.9.0](./mfa/CHANGELOG.md)
    * **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)
    * **Feature** Support for "User MFA Enabled" API and data model [#113](https://github.com/patrickcping/pingone-go-sdk-v2/pull/113)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.3](./risk/CHANGELOG.md)
    * **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)

# Release (2023-02-22)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.5.3
    * **Note** bump `golang.org/x/net` v0.5.0 => v0.7.0 [#110](https://github.com/patrickcping/pingone-go-sdk-v2/pull/110)
    * **Note** bump `github.com/golangci/golangci-lint` v1.50.1 => v1.51.2 [#110](https://github.com/patrickcping/pingone-go-sdk-v2/pull/110)
    * **Note** bump `golang.org/x/oauth2` v0.4.0 => v0.5.0 [#110](https://github.com/patrickcping/pingone-go-sdk-v2/pull/110)
    * **Note** bump `github.com/securego/gosec/v2` v2.14.0 => v2.15.0 [#110](https://github.com/patrickcping/pingone-go-sdk-v2/pull/110)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.1 => v0.1.2 [#112](https://github.com/patrickcping/pingone-go-sdk-v2/pull/112)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/management` v0.14.0 => v0.15.0 [#112](https://github.com/patrickcping/pingone-go-sdk-v2/pull/112)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.7.2 => v0.8.0 [#112](https://github.com/patrickcping/pingone-go-sdk-v2/pull/112)
    * **Note** `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.3.1 => v0.3.2 [#112](https://github.com/patrickcping/pingone-go-sdk-v2/pull/112)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.2](./authorize/CHANGELOG.md)
    * **Note** bump `golang.org/x/net` v0.2.0 => v0.7.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
    * **Note** bump `golang.org/x/oauth2` v0.2.0 => v0.5.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.15.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.8.0](./mfa/CHANGELOG.md)
    * **Note** bump `golang.org/x/net` v0.5.0 => v0.7.0 [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
    * **Note** bump `golang.org/x/oauth2` v0.4.0 => v0.5.0 [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
    * **Breaking change** `Key` property removed from the `MFAPushCredential` model object and assigned to `MFAPushCredentialFCM` and `MFAPushCredentialAPNS` individually [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
    * **Enhancement** Add support for Huawei HMS Push service [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.2](./risk/CHANGELOG.md)
    * **Note** bump `golang.org/x/net` v0.2.0 => v0.7.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
    * **Note** bump `golang.org/x/oauth2` v0.2.0 => v0.5.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)

# Release (2023-01-12)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.5.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.13.0 => v0.14.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.7.1 => v0.7.2
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.14.0](./management/CHANGELOG.md)
    * **Bug fix** - Correct the `TemplateContent` API model [#97](https://github.com/patrickcping/pingone-go-sdk-v2/pull/97)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.7.2](./mfa/CHANGELOG.md)
    * **Breaking change** Device selection `Authentication` model no longer required for the `MFASettings` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
    * **Note** Deprecated device selection from the `MFASettings` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
    * **Enhancement** - Added device selection parameters to the `DeviceAuthenticationPolicy` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
    * **Enhancement** - Added push timeout to the `DeviceAuthenticationPolicy` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)

# Release (2023-01-09)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.5.1
    * **Breaking change** PingOne Fraud is now a non-selectable capability [#92](https://github.com/patrickcping/pingone-go-sdk-v2/pull/92)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.0 => v0.1.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.12.0 => v0.13.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.7.0 => v0.7.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.3.0 => v0.3.1
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.1](./authorize/CHANGELOG.md)
    * **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.13.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.7.1](./mfa/CHANGELOG.md)
    * **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.1](./risk/CHANGELOG.md)
    * **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)

# Release (2022-11-06-2)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.5.0
    * **Enhancement** Add support for initialising the client with an access token [#82](https://github.com/patrickcping/pingone-go-sdk-v2/pull/82)

# Release (2022-11-06)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.4.3
    * **Note** bump `github.com/golangci/golangci-lint` v1.49.0 => v1.50.1
    * **Note** bump `github.com/securego/gosec/v2` v2.13.1 => v2.14.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.11.2 => v0.12.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.6.1 => v0.7.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.12.0](./management/CHANGELOG.md)
    * **Breaking change** Removed the `EnumLicensePackage` enum model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)
    * **Feature** Support for Branding Settings [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
    * **Feature** Support for Branding Themes [#76](https://github.com/patrickcping/pingone-go-sdk-v2/pull/76)
    * **Feature** Support for Resource Client Secrets [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
    * **Enhancement** Add `idToken` and `userInfo` attributes to the `ResourceAttribute` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
    * **Enhancement** Add `mappedClaims` attributes to the `ResourceScope` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
    * **Enhancement** Add `introspectEndpointAuthMethod` attributes to the `Resource` data model [#78](https://github.com/patrickcping/pingone-go-sdk-v2/pull/78)
    * **Enhancement** Added the `TERMINATED` value to the `EnumLicenseStatus` model [#81](https://github.com/patrickcping/pingone-go-sdk-v2/pull/81)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.7.0](./mfa/CHANGELOG.md)
    * **Feature** Added FIDO Policy API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)
    * **Feature** Added FIDO Custom Device Metadata API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)

# Release (2022-10-15)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.4.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.11.1 => v0.11.2
    * **Bug fix** Fix error conversion where the model doesn't exist [#72](https://github.com/patrickcping/pingone-go-sdk-v2/pull/72)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.11.2](./management/CHANGELOG.md)
    * **Bug fix** - Correct the `Image` API model [#74](https://github.com/patrickcping/pingone-go-sdk-v2/pull/74)

# Release (2022-10-10)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.4.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.11.0 => v0.11.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.6.0 => v0.6.1
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.11.1](./management/CHANGELOG.md)
    * **Enhancement** Add `uriPrefix` (mobile) and `excludedPlatforms` (mobile device integrity check) to the `ApplicationOIDC` data model [#71](https://github.com/patrickcping/pingone-go-sdk-v2/pull/71)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.6.1](./mfa/CHANGELOG.md)
    * **Bug fix** `lockout` made optional in the `MFASettings` model [#70](https://github.com/patrickcping/pingone-go-sdk-v2/pull/70)

# Release (2022-10-09)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.4.0
    * **Note** Add `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.10.0 => v0.11.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.5.1 => v0.6.0
    * **Feature** Added error models and marshal helpers to `model` package [#66](https://github.com/patrickcping/pingone-go-sdk-v2/pull/66)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.0](./authorize/CHANGELOG.md)
    * **Feature** Initial release [#55](https://github.com/patrickcping/pingone-go-sdk-v2/pull/55)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.11.0](./management/CHANGELOG.md)
    * **Feature** Support for Licenses [#64](https://github.com/patrickcping/pingone-go-sdk-v2/pull/64)
    * **Feature** Support for Languages [#63](https://github.com/patrickcping/pingone-go-sdk-v2/pull/63)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.6.0](./mfa/CHANGELOG.md)
    * **Bug fix** Corrected the Device policy API [#65](https://github.com/patrickcping/pingone-go-sdk-v2/pull/65)
    * **Bug fix** Corrected the Application push credentials API model [#67](https://github.com/patrickcping/pingone-go-sdk-v2/pull/67)

# Release (2022-09-18)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.8
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.9.0 => v0.10.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.5.0 => v0.5.1
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.10.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.5.1](./management/CHANGELOG.md)
    * **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredential201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
    * **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
    * **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedPushCredentialsInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
    * **Enhancement** - Changed model dereferencing strategy for the `UpdateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
    * **Enhancement** - Add `environment` attribute block to `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)
    * **Enhancement** - Add required attributes to the `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)

# Release (2022-09-11)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.7
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.8.0 => v0.9.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.4.0 => v0.5.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.9.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.5.0](./management/CHANGELOG.md)
    * **Feature** Added MFA Settings model [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)
    * **Breaking change** `updatedAt` attributes changed to date/time data type [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)

# Release (2022-09-04)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.6
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.7.0 => v0.8.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.8.0](./management/CHANGELOG.md)
    * **Bug fix** - Fix `Gateway` response dereferencing for non-`LDAP` types [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
    * **Bug fix** - Fix `ExportCSR` headers, to determine the response type of the CSR [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
    * **Bug fix** - Correct the `ImportCSRResponse` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
    * **Bug fix** - Correct the `GetKey` API when exporting the public certificate [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
    * **Bug fix** - Correct the `CreateCertificateFromFile` API [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
    * **Enhancement** - Add `Links` to `Gateway` model [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
    * **Enhancement** - Add `X_PEM_FILE` to `EnumCSRExportHeader` (exporting CSR formats) [#41](https://github.com/patrickcping/pingone-go-sdk-v2/pull/41)
    * **Feature** Support for Read All Gateway Credentials [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)
    * **Feature** Support for Read One Gateway Credential [#40](https://github.com/patrickcping/pingone-go-sdk-v2/pull/40)

# Release (2022-09-01)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.5
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.6.0 => v0.7.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.7.0](./management/CHANGELOG.md)
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

# Release (2022-08-29)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.4
    * **Note** bump `github.com/golangci/golangci-lint` v1.47.2 => v1.49.0
    * **Note** bump `github.com/securego/gosec/v2` v2.12.0 => v2.13.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.5.0 => v0.6.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.3.0 => v0.4.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.2.0 => v0.3.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.6.0](./management/CHANGELOG.md)
    * **Bug fix** - Correct `PINGONE_API` from incorrect `PING_ONE_API` in `EnumResourceType` enum
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.4.0](./management/CHANGELOG.md)
    * **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
    * **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.0](./management/CHANGELOG.md)
    * **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
    * **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)

# Release (2022-08-17)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.3
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.4.0 => v0.5.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.5.0](./management/CHANGELOG.md)
    * **Bug fix** - `SignOnPolicyActionProgressiveProfiling` object includes fixed `attributes` attribute, from object type to array type
    * **Bug fix** - Restructure `SignOnPolicyAction` model to properly separate policy action attributes from each other
    * **Bug fix** - Restructure `SignOnPolicyActionMFA` model to correct `noDevicesMode` attribute (fixed incorrect `noDeviceMode`)
    * **Bug fix** - Corrections to `SignOnPolicyActionCommonConditions` model
    * **Bug fix** - Moved `confirmIdentityProviderAttributes` boolean to the registration sub model of `SignOnPolicyAction` 

# Release (2022-08-10)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.3.0 => v0.4.0
    * **Enhancement** Retarget the `DaVinci` product code to `PING_ONE_DAVINCI` instead of the deprecated `PING_ONE_ORCHESTRATE`
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.4.0](./management/CHANGELOG.md)
    * **Enhancement** Added generic `_links` to `Application` model
    * **Breaking change** `id` made required in `spVerification.certificates` of the `ApplicationSAML` model
    * **Enhancement** Add `mobile.passcodeRefreshDuration` to `ApplicationOIDC` model
    * **Breaking change** `accessControl.role.type` made an enum in the `Application` model
    * **Enhancement** Added `PING_ONE_DAVINCI` product type

# Release (2022-08-05)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.2.0 => v0.3.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.2.0 => v0.3.0
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.3.0](./management/CHANGELOG.md)
    * **Enhancement** - *API update 2022-07-18* - `SignOnPolicyAction` object includes new attribute `DeviceAuthenticationPolicy` to replace deprecated `Applications`, `Voice`, `Sms`, `SecurityKey`, `Email`, `BoundBiometrics` and `Authenticator` attributes [#11](https://github.com/patrickcping/pingone-go-sdk-v2/pull/11)
    * **Feature** Agreements, Agreement Languages and Agreement Revisions [#14](https://github.com/patrickcping/pingone-go-sdk-v2/pull/14)
    * **Feature** Support for Identity Provider models [#15](https://github.com/patrickcping/pingone-go-sdk-v2/pull/15)
    * **Bug** fix for `PasswordPolicyMinCharacters` special character issue [#17](https://github.com/patrickcping/pingone-go-sdk-v2/pull/17)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.3.0](./mfa/CHANGELOG.md)
    * **Feature** Mfa device authentication policies [#13](https://github.com/patrickcping/pingone-go-sdk-v2/pull/13)

# Release (2022-07-17)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.0
    * **Breaking change** Moved `RegionsAvailableList` (renamed from `AvailableRegionsList`), `FindRegionByAPICode` and `FindRegionByName` to `model` package
    * **Feature** Added `FindProductByName`, `FindProductByAPICode`  and `ProductsSelectableList` to `model` package
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.2.0](./management/CHANGELOG.md)
    * **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
    * **Feature** Full support for enum values
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.2.0](./mfa/CHANGELOG.md)
    * **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
    * **Feature** Full support for enum values
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.2.0](./risk/CHANGELOG.md)
    * **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
    * **Feature** Full support for enum values

# Release (2022-07-16-2)

* **Feature** `github.com/patrickcping/pingone-go-sdk-v2` : v0.2.0
    * New `pingone` package with aggregated client initialisation
    * New `environment`, `schemaAttribute` and `role` helper types

# Release (2022-07-16)

Initial release - rebasing versions to reflect module stability