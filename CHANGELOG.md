# Release (2025-07-16)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.14.0
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.59.0 => v0.60.0 [#469](https://github.com/patrickcping/pingone-go-sdk-v2/pull/469)
  * **Enhancement** Add support for Singapore in region selection. [#468](https://github.com/patrickcping/pingone-go-sdk-v2/pull/468)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.60.0](./management/CHANGELOG.md)  
  * **Enhancement** Added the `AU` and `SG` region codes to the `EnumPropagationStoreTypePingOneRegion` model. [#468](https://github.com/patrickcping/pingone-go-sdk-v2/pull/468)
  * **Enhancement** Added the `SG` region code to the `EnumRegionCode` model. [#468](https://github.com/patrickcping/pingone-go-sdk-v2/pull/468)
  * **Enhancement** Added the `AU` and `SG` region codes to the `EnumRegionCodeLicense` model. [#468](https://github.com/patrickcping/pingone-go-sdk-v2/pull/468)

# Release (2025-06-27)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.13.0
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.58.0 => v0.59.0 [#466](https://github.com/patrickcping/pingone-go-sdk-v2/pull/466)
  * **Enhancement** Add support for PingID v2 environments in product selection. [#465](https://github.com/patrickcping/pingone-go-sdk-v2/pull/465)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.59.0](./management/CHANGELOG.md)  
  * **Enhancement** Added the `PING_ONE_ID` product type to support PingID (v2). [#465](https://github.com/patrickcping/pingone-go-sdk-v2/pull/465)

# Release (2025-06-24)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.19
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.57.1 => v0.58.0 [#464](https://github.com/patrickcping/pingone-go-sdk-v2/pull/464)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.58.0](./management/CHANGELOG.md)
  * **Enhancement** Add support for `VirtualServerIdSettings` in SAML application objects. [#462](https://github.com/patrickcping/pingone-go-sdk-v2/pull/462)

# Release (2025-06-11)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.18
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.8.0 => v0.8.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.11.0 => v0.11.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.57.0 => v0.57.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.23.0 => v0.23.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.19.0 => v0.19.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.9.0 => v0.9.1 [#461](https://github.com/patrickcping/pingone-go-sdk-v2/pull/461)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.8.1](./authorize/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.11.1](./credentials/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.57.1](./management/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.23.1](./mfa/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.19.1](./risk/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.9.1](./verify/CHANGELOG.md)
  * **Note** Add backoff retry on transient 404 errors expected after environment creation. [#460](https://github.com/patrickcping/pingone-go-sdk-v2/pull/460)

# Release (2025-05-29)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.17
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.56.0 => v0.57.0 [#459](https://github.com/patrickcping/pingone-go-sdk-v2/pull/459)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.57.0](./management/CHANGELOG.md)
  * **Feature** Add support for the Locale Translations API. [#457](https://github.com/patrickcping/pingone-go-sdk-v2/pull/457)
  * **Enhancement** Add support for assigning application scope to custom roles. [#458](https://github.com/patrickcping/pingone-go-sdk-v2/pull/458)

# Release (2025-05-20)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.16
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.55.0 => v0.56.0 [#456](https://github.com/patrickcping/pingone-go-sdk-v2/pull/456)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.56.0](./management/CHANGELOG.md)
  * **Enhancement** Support `SubjectNameIdentifierFormat` on WS-FED application data models. [#455](https://github.com/patrickcping/pingone-go-sdk-v2/pull/455)
  * **Bug** Corrected `WS_FED` enum value typo on the `EnumApplicationProtocol` model. [#455](https://github.com/patrickcping/pingone-go-sdk-v2/pull/455)

# Release (2025-05-13)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.15
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.54.0 => v0.55.0 [#454](https://github.com/patrickcping/pingone-go-sdk-v2/pull/454)
  * **Note** bump `golang.org/x/oauth2` v0.29.0 => v0.30.0 [#454](https://github.com/patrickcping/pingone-go-sdk-v2/pull/454)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.55.0](./management/CHANGELOG.md)
  * **Feature** Add support for the Environment Status API. [#453](https://github.com/patrickcping/pingone-go-sdk-v2/pull/453)
  * **Enhancement** Support `HardDeletedAllowedAt`, `SoftDeletedAt`, `Status` on the `Environment` data model. [#453](https://github.com/patrickcping/pingone-go-sdk-v2/pull/453)
  * **Enhancement** Support `AlternativeIdentifiers` on the `Population` data model. [#449](https://github.com/patrickcping/pingone-go-sdk-v2/pull/449)
  * **Enhancement** Support `PreferredLanguage` on the `Population` data model. [#449](https://github.com/patrickcping/pingone-go-sdk-v2/pull/449)
  * **Enhancement** Support `Theme` on the `Population` data model. [#449](https://github.com/patrickcping/pingone-go-sdk-v2/pull/449)
  * **Bug** Corrected `EnvironmentRegion` model unmarshalling when using custom region string values. [#452](https://github.com/patrickcping/pingone-go-sdk-v2/pull/452)

# Release (2025-04-28)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.14
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.53.0 => v0.54.0 [#447](https://github.com/patrickcping/pingone-go-sdk-v2/pull/447)
  * **Note** bump `golang.org/x/oauth2` v0.28.0 => v0.29.0 [#441](https://github.com/patrickcping/pingone-go-sdk-v2/pull/441)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.54.0](./management/CHANGELOG.md)
  * **Enhancement** Support the new `LINKEDIN_OIDC` Identity provider type. [#445](https://github.com/patrickcping/pingone-go-sdk-v2/pull/445)

# Release (2025-03-11)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.13
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.52.0 => v0.53.0 [#436](https://github.com/patrickcping/pingone-go-sdk-v2/pull/436)
  * **Note** bump `golang.org/x/oauth2` v0.27.0 => v0.28.0 [#436](https://github.com/patrickcping/pingone-go-sdk-v2/pull/436)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.53.0](./management/CHANGELOG.md)
  * **Enhancement** Added `IdpSignoff` to the `ApplicationOIDC` model. [#433](https://github.com/patrickcping/pingone-go-sdk-v2/pull/433)

# Release (2025-03-04)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.12
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.51.0 => v0.52.0 [#432](https://github.com/patrickcping/pingone-go-sdk-v2/pull/432)
  * **Note** bump `golang.org/x/oauth2` v0.26.0 => v0.27.0 [#432](https://github.com/patrickcping/pingone-go-sdk-v2/pull/432)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.52.0](./management/CHANGELOG.md)
  * **Breaking change** The `AdministratorSecurity` model has been aligned to the service. [#429](https://github.com/patrickcping/pingone-go-sdk-v2/pull/429)
 
# Release (2025-02-24)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.11
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.50.0 => v0.51.0 [#428](https://github.com/patrickcping/pingone-go-sdk-v2/pull/428)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.51.0](./management/CHANGELOG.md)
  * **Breaking change** The Microsoft Identity provider now uses the `IdentityProviderMicrosoft` model when using the `IdentityProvider` model. [#427](https://github.com/patrickcping/pingone-go-sdk-v2/pull/427)

# Release (2025-02-21)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.10
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.49.0 => v0.50.0 [#426](https://github.com/patrickcping/pingone-go-sdk-v2/pull/426)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.50.0](./management/CHANGELOG.md)
  * **Breaking change** Added support for the `TenantId` field on Microsoft Identity providers. The Microsoft Identity provider now uses the `IdentityProviderMicrosoft` model. [#425](https://github.com/patrickcping/pingone-go-sdk-v2/pull/425)
 
# Release (2025-02-10)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.9
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.7.0 => v0.8.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.10.0 => v0.11.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.48.0 => v0.49.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.22.0 => v0.23.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.18.0 => v0.19.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.8.0 => v0.9.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
  * **Note** bump `golang.org/x/oauth2` v0.25.0 => v0.26.0 [#422](https://github.com/patrickcping/pingone-go-sdk-v2/pull/422)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.8.0](./authorize/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.11.0](./credentials/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.49.0](./management/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)
  * **Bug** Corrected `data matches more than one schema in oneOf(NotificationsSettingsEmailDeliverySettings)` error for email notification settings on new environments (again). [#417](https://github.com/patrickcping/pingone-go-sdk-v2/pull/417)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.23.0](./mfa/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.19.0](./risk/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.9.0](./verify/CHANGELOG.md)
  * **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
  * **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)

# Release (2025-02-05)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.8
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.47.0 => v0.48.0 [#416](https://github.com/patrickcping/pingone-go-sdk-v2/pull/416)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.48.0](./management/CHANGELOG.md)
  * **Enhancement** Added the `filter` query string parameter function to the `ReadAllCustomAdminRoles(..)` API request model. [#414](https://github.com/patrickcping/pingone-go-sdk-v2/pull/414)
  * **Bug** Corrected `data matches more than one schema in oneOf(NotificationsSettingsEmailDeliverySettings)` error for email notification settings on new environments. [#415](https://github.com/patrickcping/pingone-go-sdk-v2/pull/415)

# Release (2025-01-23)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.7
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.46.0 => v0.47.0 [#413](https://github.com/patrickcping/pingone-go-sdk-v2/pull/413)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.47.0](./management/CHANGELOG.md)
  * **Enhancement** Added `EnvironmentDnsRecord` to the `EmailDomainOwnershipStatus` model. [#412](https://github.com/patrickcping/pingone-go-sdk-v2/pull/412)

# Release (2025-01-21)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.6
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.45.0 => v0.46.0 [#411](https://github.com/patrickcping/pingone-go-sdk-v2/pull/411)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.21.0 => v0.22.0 [#411](https://github.com/patrickcping/pingone-go-sdk-v2/pull/411)
  * **Note** bump `golang.org/x/oauth2` v0.24.0 => v0.25.0 [#411](https://github.com/patrickcping/pingone-go-sdk-v2/pull/411)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.46.0](./management/CHANGELOG.md)
  * **Enhancement** Added roles `Advanced Identity Cloud Super Admin`, `Advanced Identity Cloud Tenant Admin` and `Custom Roles Admin` to the `EnumRoleName` model. [#407](https://github.com/patrickcping/pingone-go-sdk-v2/pull/407)
  * **Enhancement** Added `BlastRadiusMitigation` to the `GatewayTypeRADIUSAllOfRadiusClients` model. [#410](https://github.com/patrickcping/pingone-go-sdk-v2/pull/410)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.22.0](./mfa/CHANGELOG.md)
  * **Enhancement** Added `UserPresenceTimeout` to the `FIDO2Policy` model. [#408](https://github.com/patrickcping/pingone-go-sdk-v2/pull/408)
  * **Enhancement** Added `PublicKeyCredentialHints` to the `FIDO2Policy` model. [#409](https://github.com/patrickcping/pingone-go-sdk-v2/pull/409)

# Release (2024-12-18)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.5
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.44.0 => v0.45.0 [#403](https://github.com/patrickcping/pingone-go-sdk-v2/pull/403)
  * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.17.0 => v0.18.0 [#403](https://github.com/patrickcping/pingone-go-sdk-v2/pull/403)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.45.0](./management/CHANGELOG.md)
  * **Enhancement** Added `SUSPICIOUS_TRAFFIC` enum to `EnumAlertChannelAlertType` model. [#399](https://github.com/patrickcping/pingone-go-sdk-v2/pull/399)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.18.0](./risk/CHANGELOG.md)
  * **Bug** Add missing protect predictor composite conditions `IP_RANGE` and `IP_COMPARISON`. [#401](https://github.com/patrickcping/pingone-go-sdk-v2/pull/401)

# Release (2024-11-15)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.4
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.6.0 => v0.7.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.9.0 => v0.10.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.43.0 => v0.44.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.20.0 => v0.21.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.16.0 => v0.17.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.7.0 => v0.8.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
    * **Note** bump `golang.org/x/oauth2` v0.21.0 => v0.24.0 [#397](https://github.com/patrickcping/pingone-go-sdk-v2/pull/397)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.3.3](./agreementmanagement/CHANGELOG.md)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.7.0](./authorize/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.10.0](./credentials/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.44.0](./management/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Breaking change** Model `NotificationsSettingsEmailDeliverySettings` is now a compound model supporting both SMTP and custom notification models. [#386](https://github.com/patrickcping/pingone-go-sdk-v2/pull/386)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Feature** Added support for the Administrator Security API. [#381](https://github.com/patrickcping/pingone-go-sdk-v2/pull/381)
    * **Feature** Added support for custom email notification providers. [#386](https://github.com/patrickcping/pingone-go-sdk-v2/pull/386)
    * **Feature** Added support for the Custom Admin Roles API. [#388](https://github.com/patrickcping/pingone-go-sdk-v2/pull/388)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Enhancement** Added the `SessionNotOnOrAfterDuration` field to the `ApplicationSAML` data model. [#387](https://github.com/patrickcping/pingone-go-sdk-v2/pull/387)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.21.0](./mfa/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Enhancement** Added the `UriParameters` field to the `DeviceAuthenticationPolicyTotp` data model. [#384](https://github.com/patrickcping/pingone-go-sdk-v2/pull/384)
    * **Enhancement** Added the `OtpLength` field to the `DeviceAuthenticationPolicyOfflineDeviceOtp` data model. [#385](https://github.com/patrickcping/pingone-go-sdk-v2/pull/385)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.17.0](./risk/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Enhancement** Added the `ShouldValidatePayloadSignature` field to the `RiskPredictorDevice` model. [#380](https://github.com/patrickcping/pingone-go-sdk-v2/pull/380)
    * **Enhancement** Added the `IncludeRepeatedEventsWithoutSdk` field to the `RiskPredictorBotDetection` model. [#382](https://github.com/patrickcping/pingone-go-sdk-v2/pull/382)
    * **Enhancement** Support the `TRAFFIC_ANOMALY` risk predictor type (new `RiskPredictorTrafficAnomaly` model). [#383](https://github.com/patrickcping/pingone-go-sdk-v2/pull/383)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.8.0](./verify/CHANGELOG.md)
    * **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
    * **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
    * **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)

# Release (2024-07-22)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.3
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.42.0 => v0.43.0 [#371](https://github.com/patrickcping/pingone-go-sdk-v2/pull/371)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.6.0 => v0.7.0 [#371](https://github.com/patrickcping/pingone-go-sdk-v2/pull/371)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.43.0](./management/CHANGELOG.md)
    * **Breaking change** API name change for Integration Catalog [#370](https://github.com/patrickcping/pingone-go-sdk-v2/pull/370)
    * **Enhancement** Added "Integration", "Integration Version",  "Integration Version Attribute" data models. [#370](https://github.com/patrickcping/pingone-go-sdk-v2/pull/370)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.7.0](./verify/CHANGELOG.md)
    * **Enhancement** Added `Retry` field to the `LivenessConfiguration` model. [#369](https://github.com/patrickcping/pingone-go-sdk-v2/pull/369)
    * **Enhancement** Added `FailExpiredId`, `Provider` and `Retry` fields to the `GovernmentIdConfiguration` object. [#369](https://github.com/patrickcping/pingone-go-sdk-v2/pull/369)

# Release (2024-07-04)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.5.0 => v0.6.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.8.0 => v0.9.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.41.0 => v0.42.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.19.0 => v0.20.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.15.1 => v0.16.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.5.0 => v0.6.0 [#368](https://github.com/patrickcping/pingone-go-sdk-v2/pull/368)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.6.0](./authorize/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.9.0](./credentials/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.42.0](./management/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added trials ENUM values to the `EnumSolutionType` model. [#364](https://github.com/patrickcping/pingone-go-sdk-v2/pull/364)
    * **Enhancement** Added `AlertName` field to the `AlertChannel` model. [#363](https://github.com/patrickcping/pingone-go-sdk-v2/pull/363)
    * **Enhancement** Added support for granting the `Application Owner` role. [#365](https://github.com/patrickcping/pingone-go-sdk-v2/pull/365)
    * **Bug** Remove unsupported `AlertingApi.ReadOneAlertChannel`. [#363](https://github.com/patrickcping/pingone-go-sdk-v2/pull/363)
    * **Bug** Corrected `EnumAlertChannelAlertType` enum values. [#363](https://github.com/patrickcping/pingone-go-sdk-v2/pull/363)
    * **Bug** Corrected `EnumGatewayVendor` enum values. [#367](https://github.com/patrickcping/pingone-go-sdk-v2/pull/367)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.20.0](./mfa/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.16.0](./risk/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.6.0](./verify/CHANGELOG.md)
    * **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
    * **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)

# Release (2024-06-18)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.40.0 => v0.41.0 [#362](https://github.com/patrickcping/pingone-go-sdk-v2/pull/362)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.15.0 => v0.15.1 [#362](https://github.com/patrickcping/pingone-go-sdk-v2/pull/362)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.41.0](./management/CHANGELOG.md)
    * **Enhancement** Added `SpEncryption` to the `ApplicationSAML` and associated models to control encryption of SAML application assertions. [#348](https://github.com/patrickcping/pingone-go-sdk-v2/pull/348)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.15.1](./risk/CHANGELOG.md)
    * **Bug** Corrected `DomainWhiteList` field in the `RiskPredictorAdversaryInTheMiddle` model. [#361](https://github.com/patrickcping/pingone-go-sdk-v2/pull/361)

# Release (2024-06-07)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.12.0
    * **Note** Removal of retracted module `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement`. [#352](https://github.com/patrickcping/pingone-go-sdk-v2/pull/352)
    * **Note** Deprecated the region name.  Region codes should be used instead. [#358](https://github.com/patrickcping/pingone-go-sdk-v2/pull/358)
    * **Note** bump `golang.org/x/oauth2` v0.19.0 => v0.21.0 [#341](https://github.com/patrickcping/pingone-go-sdk-v2/pull/341)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.4.1 => v0.5.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.7.0 => v0.8.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.39.0 => v0.40.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.18.3 => v0.19.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.14.1 => v0.15.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.4.1 => v0.5.0 [#360](https://github.com/patrickcping/pingone-go-sdk-v2/pull/360)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.3.2](./agreementmanagement/CHANGELOG.md)
    * **MODULE RETRACTION** The API endpoint in this module suffers major loss of function. The module has been retracted but is retained in the source repo for the purpose of retration. [#352](https://github.com/patrickcping/pingone-go-sdk-v2/pull/352)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.5.0](./authorize/CHANGELOG.md)
    * **Note** Removed unused `Operations` field from the `APIServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)
    * **Feature** Add support for Application Resource Permissions API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
    * **Feature** Add support for Application Resources API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
    * **Feature** Add support for Application Roles API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
    * **Feature** Add support for Application Role Assignments API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
    * **Feature** Add support for Application Role Permissions API. [#344](https://github.com/patrickcping/pingone-go-sdk-v2/pull/344)
    * **Feature** Add support for API Server Operations. [#354](https://github.com/patrickcping/pingone-go-sdk-v2/pull/354)
    * **Feature** Add support for API Server Deployment.
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
    * **Enhancement** Added `AccessControl` and `Directory` fields to the `APIServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)
    * **Enhancement** Added `Type` field to the `APIServerAuthorizationServer` model. [#353](https://github.com/patrickcping/pingone-go-sdk-v2/pull/353)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.8.0](./credentials/CHANGELOG.md)
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.40.0](./management/CHANGELOG.md)
    * **Breaking Change** Removed the `FormSocialLoginButtonStyles` data model.  Use the `FormStyles` data model going forward. [#350](https://github.com/patrickcping/pingone-go-sdk-v2/pull/350)
    * **Note** Removed unnecessary `Width` and `IconSrc` fields from `FormSocialLoginButton` and associated data models. [#350](https://github.com/patrickcping/pingone-go-sdk-v2/pull/350)
    * **Feature** Add support for Application Resource API. [#346](https://github.com/patrickcping/pingone-go-sdk-v2/pull/346)
    * **Feature** Add support for Application Resource Permissions API. [#346](https://github.com/patrickcping/pingone-go-sdk-v2/pull/346)
    * **Feature** Add support for User Application Role Assignment API. [#346](https://github.com/patrickcping/pingone-go-sdk-v2/pull/346)
    * **Enhancement** Added `ApplicationPermissionsSettings` to the `Resource` model. [#346](https://github.com/patrickcping/pingone-go-sdk-v2/pull/346)
    * **Enhancement** Added `Key` field to `FormSocialLoginButton` and associated data models. [#350](https://github.com/patrickcping/pingone-go-sdk-v2/pull/350)
    * **Enhancement** Added the `DeletePreviousResourceSecret` function to control resource secret rotation. [#347](https://github.com/patrickcping/pingone-go-sdk-v2/pull/347)
    * **Enhancement** Added `Previous` to the `ResourceSecret` model to control resource secret rotation. [#347](https://github.com/patrickcping/pingone-go-sdk-v2/pull/347)
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
    * **Enhancement** Add the `AU` region code to the `EnumRegionCode` model. [#358](https://github.com/patrickcping/pingone-go-sdk-v2/pull/358)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.19.0](./mfa/CHANGELOG.md)
    * **Breaking change** Remove optional `DeviceAuthenticationPolicyMobileOtpWindow` from `DeviceAuthenticationPolicyMobileOtp` constructor. [#343](https://github.com/patrickcping/pingone-go-sdk-v2/pull/343)
    * **Enhancement** Added `PromptForNicknameOnPairing` boolean field to the `DeviceAuthenticationPolicyFido2`, `DeviceAuthenticationPolicyMobile`, `DeviceAuthenticationPolicyOfflineDevice` and `DeviceAuthenticationPolicyTotp` models. [#349](https://github.com/patrickcping/pingone-go-sdk-v2/pull/349)
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.15.0](./risk/CHANGELOG.md)
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.5.0](./verify/CHANGELOG.md)
    * **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)

# Release (2024-05-01)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.9
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.6.2 => v0.7.0 [#340](https://github.com/patrickcping/pingone-go-sdk-v2/pull/340)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.38.0 => v0.39.0 [#340](https://github.com/patrickcping/pingone-go-sdk-v2/pull/340)
    * **Note** bump `golang.org/x/net` v0.22.0 => v0.23.0 [#340](https://github.com/patrickcping/pingone-go-sdk-v2/pull/340)
    * **Note** bump `golang.org/x/oauth2` v0.18.0 => v0.19.0 [#340](https://github.com/patrickcping/pingone-go-sdk-v2/pull/340)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.7.0](./credentials/CHANGELOG.md)
    * **Enhancement** Add `Management` to the `CredentialType` data model. [#339](https://github.com/patrickcping/pingone-go-sdk-v2/pull/339)
    * **Enhancement** Add `Required` to the `CredentialTypeMetaDataFieldsInner` data model. [#339](https://github.com/patrickcping/pingone-go-sdk-v2/pull/339)
    * **Bug** Fixed `DeletedAt` property (made read-only) in the `CredentialType` data model and update the data type. [#339](https://github.com/patrickcping/pingone-go-sdk-v2/pull/339)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.39.0](./management/CHANGELOG.md)
    * **Enhancement** Added the `NetworkPolicyServer` property to the `GatewayTypeRADIUS` data model. [#336](https://github.com/patrickcping/pingone-go-sdk-v2/pull/336)

# Release (2024-03-25)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.8
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.14.0 => v0.14.1 [#333](https://github.com/patrickcping/pingone-go-sdk-v2/pull/333)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.14.1](./risk/CHANGELOG.md)
    * **Bug** Fixed error "Error when calling `ReadAllRiskPredictors`: Data failed to match schemas in oneOf(RiskPredictor)" when the `EMAIL_REPUTATION` predictor type is in the API's response payload. [#332](https://github.com/patrickcping/pingone-go-sdk-v2/pull/332)

# Release (2024-03-15)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.7
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.37.0 => v0.38.0 [#331](https://github.com/patrickcping/pingone-go-sdk-v2/pull/331)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.13.0 => v0.14.0 [#331](https://github.com/patrickcping/pingone-go-sdk-v2/pull/331)
    * **Note** bump `golang.org/x/oauth2` v0.17.0 => v0.18.0 [#331](https://github.com/patrickcping/pingone-go-sdk-v2/pull/331)
    * **Note** bump `google.golang.org/protobuf` v1.31.0 => v1.33.0 [#331](https://github.com/patrickcping/pingone-go-sdk-v2/pull/331)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.38.0](./management/CHANGELOG.md)
    * **Enhancement** Added the `DefaultTargetUrl` property to the `ApplicationSAML` data model. [#325](https://github.com/patrickcping/pingone-go-sdk-v2/pull/325)
    * **Enhancement** Added the `UpdateUserOnSuccessfulAuthentication` property to the `GatewayTypeLDAPAllOfUserTypes` data model. [#328](https://github.com/patrickcping/pingone-go-sdk-v2/pull/328)
    * **Enhancement** Added the `PkceMethod` ENUM property to the `IdentityProviderOIDC` data model. [#329](https://github.com/patrickcping/pingone-go-sdk-v2/pull/329)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.14.0](./risk/CHANGELOG.md)
    * **Enhancement** Added the `EMAIL_REPUTATION` predictor type. [#330](https://github.com/patrickcping/pingone-go-sdk-v2/pull/330)

# Release (2024-02-23)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.6
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.36.0 => v0.37.0 [#322](https://github.com/patrickcping/pingone-go-sdk-v2/pull/322)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.12.2 => v0.13.0 [#322](https://github.com/patrickcping/pingone-go-sdk-v2/pull/322)
    * **Note** bump `golang.org/x/oauth2` v0.16.0 => v0.17.0 [#318](https://github.com/patrickcping/pingone-go-sdk-v2/pull/318)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.37.0](./management/CHANGELOG.md)
    * **Enhancement** Add new API operations `ReadOnePopulationDefaultIdp` and `UpdatePopulationDefaultIdp` to support setting default identity providers to populations. [#319](https://github.com/patrickcping/pingone-go-sdk-v2/pull/319)
    * **Enhancement** Added the `Signing` property to the `ApplicationOIDC` data model, to support assigning custom defined KRPs to a supported application. [#320](https://github.com/patrickcping/pingone-go-sdk-v2/pull/320)
    * **Enhancement** Added the `DevicePathId`, `DeviceCustomVerificationUri`, `DeviceTimeout`, `DevicePollingInterval` properties to the `ApplicationOIDC` data model and extended the `EnumApplicationOIDCGrantType` ENUM, to support the `DEVICE_CODE` application grant type. [#320](https://github.com/patrickcping/pingone-go-sdk-v2/pull/320)
    * **Enhancement** Added the `Jwks`, `JwksUrl` properties to the `ApplicationOIDC` data model and extended the `EnumApplicationOIDCTokenAuthMethod` ENUM, to support the `PRIVATE_KEY_JWT` and `CLIENT_SECRET_JWT` token endpoint auth methods. [#320](https://github.com/patrickcping/pingone-go-sdk-v2/pull/320)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.13.0](./risk/CHANGELOG.md)
    * **Enhancement** Added the `RiskPredictorAdversaryInTheMiddle` property to the `RiskPredictor` object model and extended the `EnumPredictorType` enum with `ADVERSARY_IN_THE_MIDDLE` to support the adversary in the middle predictor. [#321](https://github.com/patrickcping/pingone-go-sdk-v2/pull/321)

# Release (2024-01-30)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.5
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.35.0 => v0.36.0 [#316](https://github.com/patrickcping/pingone-go-sdk-v2/pull/316)
    * **Note** bump `golang.org/x/oauth2` v0.15.0 => v0.16.0 [#316](https://github.com/patrickcping/pingone-go-sdk-v2/pull/316)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.36.0](./management/CHANGELOG.md)
    * **Breaking Change** Notification template names, used on the `NotificationsTemplatesApiService` service now uses an enum which defines the templates supported in the API path. [#314](https://github.com/patrickcping/pingone-go-sdk-v2/pull/314)
    * **Enhancement** Add ability to designate an application's client secret as a previous secret with an expiry date up to 30 days, for rotation purposes. [#311](https://github.com/patrickcping/pingone-go-sdk-v2/pull/311)
    * **Enhancement** Add ability to set an icon on the `Environment` object model. [#313](https://github.com/patrickcping/pingone-go-sdk-v2/pull/313)

# Release (2024-01-12)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.4
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.3.0 => v0.3.1 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.4.0 => v0.4.1 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.6.1 => v0.6.2 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.34.0 => v0.35.0 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.18.2 => v0.18.3 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.12.1 => v0.12.2 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.4.0 => v0.4.1 [#308](https://github.com/patrickcping/pingone-go-sdk-v2/pull/308)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.3.1](./agreementmanagement/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.4.1](./authorize/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.6.2](./credentials/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.35.0](./management/CHANGELOG.md)
    * **Breaking Change** Add ability to use a freetext region code on environment creation for non-production environments.  The `Region` parameter on the `Environment` object model now takes a `EnvironmentRegion` object, where one of `EnumRegionCode` or `string`. [#304](https://github.com/patrickcping/pingone-go-sdk-v2/pull/304)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.18.3](./mfa/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.12.2](./risk/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.4.1](./verify/CHANGELOG.md)
    * **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
    * **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
    * **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)

# Release (2023-12-27)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.3
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.6.0 => v0.6.1 [#301](https://github.com/patrickcping/pingone-go-sdk-v2/pull/301)
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.33.0 => v0.34.0 [#301](https://github.com/patrickcping/pingone-go-sdk-v2/pull/301)
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.18.1 => v0.18.2 [#301](https://github.com/patrickcping/pingone-go-sdk-v2/pull/301)
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.12.0 => v0.12.1 [#301](https://github.com/patrickcping/pingone-go-sdk-v2/pull/301)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.6.1](./credentials/CHANGELOG.md)
    * **Note** Remove redundant data models and documentation. [#300](https://github.com/patrickcping/pingone-go-sdk-v2/pull/300)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.34.0](./management/CHANGELOG.md)
    * **Note** Adjust CORS origins documentation. [#291](https://github.com/patrickcping/pingone-go-sdk-v2/pull/291)
    * **Note** Remove redundant data models and documentation. [#300](https://github.com/patrickcping/pingone-go-sdk-v2/pull/300)
    * **Feature** Add support for Identity Propagation Plans API. [#299](https://github.com/patrickcping/pingone-go-sdk-v2/pull/299)
    * **Enhancement** Change `Type` property in the `ApplicationAccessControlGroup` object model to be an ENUM. [#295](https://github.com/patrickcping/pingone-go-sdk-v2/pull/295)
    * **Enhancement** Added the `Include` query string parameter to the `ApiReadFormRequest` API. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Enhancement** Added ability to configure `FormFieldTextblob` form controls. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Enhancement** Added the `ShowPasswordRequirements` property to password based form controls. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Enhancement** Created `FormStyles` object model to make usage simpler. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Enhancement** Added `FormStylesPadding` object model to support custom style override of form controls. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed the `OtherOptionLabel` property for form field models. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed the `OtherOptionInputLabel` property for form field models. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed the `Alignment` property for form field style models. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed required propertys for form field models. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed the `Options` form field property object. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Fixed the `FormFieldCombobox` form object. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Removal of unnecessary `Key` property from the `FormRecaptchaV2` form object. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
    * **Bug** Added required `Key` property to the `FormQrCode` form object. [#297](https://github.com/patrickcping/pingone-go-sdk-v2/pull/297)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.18.2](./mfa/CHANGELOG.md)
    * **Note** Remove redundant data models and documentation. [#300](https://github.com/patrickcping/pingone-go-sdk-v2/pull/300)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.12.1](./risk/CHANGELOG.md)
    * **Note** Remove redundant data models and documentation. [#300](https://github.com/patrickcping/pingone-go-sdk-v2/pull/300)

# Release (2023-11-29) (2)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.2
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.18.0 => v0.18.1 [#290](https://github.com/patrickcping/pingone-go-sdk-v2/pull/290)
     * **Note** bump `golang.org/x/oauth2` v0.14.0 => v0.15.0 [#290](https://github.com/patrickcping/pingone-go-sdk-v2/pull/290)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.18.1](./mfa/CHANGELOG.md)
    * **Bug** Fixed invalid value for `EnumMFADevicePolicySelection` model.  `ALWAYS_PROMPT_TO_SELECT` corrected to `ALWAYS_DISPLAY_DEVICES`. [#289](https://github.com/patrickcping/pingone-go-sdk-v2/pull/289)

# Release (2023-11-29)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.1
     * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.32.0 => v0.33.0 [#287](https://github.com/patrickcping/pingone-go-sdk-v2/pull/287)
     * **Note** bump `golang.org/x/oauth2` v0.13.0 => v0.14.0 [#287](https://github.com/patrickcping/pingone-go-sdk-v2/pull/287)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.33.0](./management/CHANGELOG.md)
    * **Feature** Add support for User Account API (allowing lock and unlock). [#282](https://github.com/patrickcping/pingone-go-sdk-v2/pull/282)
    * **Enhancement** Expand the `GroupMembership` data model. [#284](https://github.com/patrickcping/pingone-go-sdk-v2/pull/284)
    * **Enhancement** Simplified the `GroupMembershipApi` request and response payload models. [#284](https://github.com/patrickcping/pingone-go-sdk-v2/pull/284)
    * **Enhancement** Corrected and expanded the `SchemaAttributePatch` request and response payload. [#285](https://github.com/patrickcping/pingone-go-sdk-v2/pull/285)
    * **Enhancement** Added `CorsSettings` object attribute to the `ApplicationOIDC`, `ApplicationSAML` and `ApplicationWSFED` object models. [#286](https://github.com/patrickcping/pingone-go-sdk-v2/pull/286)

# Release (2023-11-10)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.11.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.2.2 => v0.3.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.3.1 => v0.4.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.5.0 => v0.6.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.31.0 => v0.32.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.17.0 => v0.18.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.11.0 => v0.12.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.3.1 => v0.4.0 [#280](https://github.com/patrickcping/pingone-go-sdk-v2/pull/280)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.3.0](./agreementmanagement/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.4.0](./authorize/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.6.0](./credentials/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.32.0](./management/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.18.0](./mfa/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.12.0](./risk/CHANGELOG.md)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.4.0](./verify/CHANGELOG.md)
    * **Enhancement** Added the `InspectionType` property to the `GovernmentIdConfiguration` object model. [#278](https://github.com/patrickcping/pingone-go-sdk-v2/pull/278)
    * **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
    * **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)

# Release (2023-11-01)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.9
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.4.1 => v0.5.0 [#277](https://github.com/patrickcping/pingone-go-sdk-v2/pull/277)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.30.0 => v0.31.0 [#277](https://github.com/patrickcping/pingone-go-sdk-v2/pull/277)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.5.0](./credentials/CHANGELOG.md)
    * **Enhancement** Added the `Default` property to the `CredentialTypeMetaDataFieldsInner` object model. [#275](https://github.com/patrickcping/pingone-go-sdk-v2/pull/275)
    * **Enhancement** Added `Multiple` object attribute to the `CredentialType` object model. [#275](https://github.com/patrickcping/pingone-go-sdk-v2/pull/275)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.31.0](./management/CHANGELOG.md)
    * **Enhancement** Added `DisplayName`, `SourceId` and `SourceType` object attributes to the `Group` object model. [#264](https://github.com/patrickcping/pingone-go-sdk-v2/pull/264)
    * **Enhancement** Added `TlsClientAuthKeyPair` object attributes to the `Subscription` object model. [#265](https://github.com/patrickcping/pingone-go-sdk-v2/pull/265)
    * **Enhancement** Added `OUTBOUND_MTLS` to the `EnumCertificateKeyUsageType` enum. [#265](https://github.com/patrickcping/pingone-go-sdk-v2/pull/265)
    * **Enhancement** Added ability to set a PKCS12 keystore password when building a `CreateKeyRequest`. [#266](https://github.com/patrickcping/pingone-go-sdk-v2/pull/266)
    * **Enhancement** Better define the `Role` and `RolePermissionsInner` data models. [#270](https://github.com/patrickcping/pingone-go-sdk-v2/pull/270)
    * **Enhancement** Support group role assignments. [#271](https://github.com/patrickcping/pingone-go-sdk-v2/pull/271)
    * **Enhancement** Added `Tags` object attribute to the `BillOfMaterialsProductsInner` model, to facilitate creation of DaVinci enabled environments without example configuration. [#272](https://github.com/patrickcping/pingone-go-sdk-v2/pull/272)
    * **Enhancement** Added support for the Propagation Stores API. [#276](https://github.com/patrickcping/pingone-go-sdk-v2/pull/276)
    * **Bug** Corrected situations where `EntityArrayEmbeddedGatewaysInner` unmarshal did not return objects of `type` `ENUMGATEWAYTYPE_RADIUS` properly. [#273](https://github.com/patrickcping/pingone-go-sdk-v2/pull/273)

# Release (2023-10-16)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.8
    * **Note** bump `github.com/securego/gosec/v2` v2.17.0 => v2.18.1 [#261](https://github.com/patrickcping/pingone-go-sdk-v2/pull/261)
    * **Note** bump `golang.org/x/oauth2` v0.12.0 => v0.13.0 [#259](https://github.com/patrickcping/pingone-go-sdk-v2/pull/259)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.2.1 => v0.2.2 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.3.0 => v0.3.1 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.4.0 => v0.4.1 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.29.0 => v0.30.0 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.16.0 => v0.17.0 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.10.0 => v0.11.0 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.3.0 => v0.3.1 [#263](https://github.com/patrickcping/pingone-go-sdk-v2/pull/263)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.2.2](./agreementmanagement/CHANGELOG.md)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.3.1](./authorize/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.4.1](./credentials/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.30.0](./management/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Enhancement** Added `Resource` object attribute to the `ResourceScope` object model. [#262](https://github.com/patrickcping/pingone-go-sdk-v2/pull/262)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.17.0](./mfa/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Enhancement** Add `Users` object to the `MFASettings` model to support setting default MFA settings for registering users. [#256](https://github.com/patrickcping/pingone-go-sdk-v2/pull/256)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.11.0](./risk/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Enhancement** Enhance the `EnumFlowType` with new supported values. [#256](https://github.com/patrickcping/pingone-go-sdk-v2/pull/256)
    * **Enhancement** Added the `Compositions` array to the `RiskPredictorComposite` model, to replace the `Composition` attribute that has now been deprecated. [#256](https://github.com/patrickcping/pingone-go-sdk-v2/pull/256)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.3.1](./verify/CHANGELOG.md)
    * **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
    * **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)

# Release (2023-09-15)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.7
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.28.0 => v0.29.0 [#254](https://github.com/patrickcping/pingone-go-sdk-v2/pull/254)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.9.0 => v0.10.0 [#254](https://github.com/patrickcping/pingone-go-sdk-v2/pull/254)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.29.0](./management/CHANGELOG.md)
    * **Enhancement** Added the `application/vnd.pingidentity.user.import+json` content-type header to the UsersAPI, to be able to import passwords. [#253](https://github.com/patrickcping/pingone-go-sdk-v2/pull/253)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.10.0](./management/CHANGELOG.md)
    * **Enhancement** Add support for a Bot detection specific data model. [#252](https://github.com/patrickcping/pingone-go-sdk-v2/pull/252)

# Release (2023-09-11)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.6
    * **Note** bump `golang.org/x/oauth2` v0.11.0 => v0.12.0 [#251](https://github.com/patrickcping/pingone-go-sdk-v2/pull/251)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.27.0 => v0.28.0 [#251](https://github.com/patrickcping/pingone-go-sdk-v2/pull/251)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.28.0](./management/CHANGELOG.md)
    * **Breaking Change** Removed deprecated enum values `SHA224withRSA` and `SHA224withECDSA` from the `EnumCertificateKeySignagureAlgorithm` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added enum values to `EnumApplicationWSFEDIDPSigningAlgorithm` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to model attributes that use the `EnumApplicationSAMLSloBinding` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to model attributes that use the `EnumGatewayTypeLDAPSecurity` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to model attributes that use the `EnumIdentityProvider` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to model attributes that use the `EnumSchemaAttributeType` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to model attributes that use the `EnumTemplateContentPushCategory` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to the `AssertionSigned` attributes on the `ApplicationSAML` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added default value to the `ResponseSigned` attributes on the `ApplicationSAML` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added support for the `AuthnRequestSigned` attribute on the `ApplicationSAMLAllOfSpVerification` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added support for the `Algorithm` attribute on the `IdentityProviderSAMLAllOfSpSigning` model. [#247](https://github.com/patrickcping/pingone-go-sdk-v2/pull/247)
    * **Enhancement** Added support for new attributes `AdditionalRefreshTokenReplayProtectionEnabled`, `RequireSignedRequestObject`, `ParRequirement`, `ParTimeout` to the `ApplicationOIDC` and `ApplicationOIDCAllOf` data models. [#248](https://github.com/patrickcping/pingone-go-sdk-v2/pull/248)

# Release (2023-09-05)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.5
    * **Note** bump `github.com/golangci/golangci-lint` v1.54.1 => v1.54.2 [#246](https://github.com/patrickcping/pingone-go-sdk-v2/pull/246)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.26.0 => v0.27.0 [#246](https://github.com/patrickcping/pingone-go-sdk-v2/pull/246)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.27.0](./management/CHANGELOG.md)
    * **Enhancement** Added `EnableRequestedAuthnContext` to the `ApplicationSAML` and `ApplicationSAMLAllOf` data models. [#245](https://github.com/patrickcping/pingone-go-sdk-v2/pull/245)

# Release (2023-08-22)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.4
    * **Note** bump `github.com/securego/gosec/v2` v2.16.0 => v2.17.0 [#242](https://github.com/patrickcping/pingone-go-sdk-v2/pull/242)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.3.1 => v0.4.0 [#242](https://github.com/patrickcping/pingone-go-sdk-v2/pull/242)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.4.0](./credentials/CHANGELOG.md)
    * **Enhancement** Add parameters `OnDelete` to the `CredentialType` model. [#240](https://github.com/patrickcping/pingone-go-sdk-v2/pull/240)
    * **Enhancement** Add parameters `FileSupport` to the `CredentialTypeMetaDataFieldsInner` model. [#240](https://github.com/patrickcping/pingone-go-sdk-v2/pull/240)

# Release (2023-08-15)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.3
    * **Note** bump `github.com/golangci/golangci-lint` v1.53.3 => v1.54.1 [#238](https://github.com/patrickcping/pingone-go-sdk-v2/pull/238)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.3.0 => v0.3.1 [#238](https://github.com/patrickcping/pingone-go-sdk-v2/pull/238)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.25.0 => v0.26.0 [#238](https://github.com/patrickcping/pingone-go-sdk-v2/pull/238)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.26.0](./management/CHANGELOG.md)
    * **Bug** Fixed inclusion of `FormManagementApi` and `RecaptchaConfigurationApi` API to the client. [#235](https://github.com/patrickcping/pingone-go-sdk-v2/pull/235)
    * **Enhancement** Clarified the `headers` parameter in the `SubscriptionHttpEndpoint` model of the subscriptions API endpoint to be a map of strings rather than a map of any data type. [#234](https://github.com/patrickcping/pingone-go-sdk-v2/pull/234)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.3.1](./credentials/CHANGELOG.md)
    * **Bug** Corrected situations where `EntityArrayEmbeddedItemsInner` object dereferencing led to panic error. [#237](https://github.com/patrickcping/pingone-go-sdk-v2/pull/237)

# Release (2023-08-08)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.2.1 => v0.3.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.2.1 => v0.3.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.24.0 => v0.25.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.15.0 => v0.16.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.8.1 => v0.9.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.2.1 => v0.3.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
    * **Note** bump `golang.org/x/oauth2` v0.10.0 => v0.11.0 [#233](https://github.com/patrickcping/pingone-go-sdk-v2/pull/233)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.3.0](./authorize/CHANGELOG.md)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.3.0](./credentials/CHANGELOG.md)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.25.0](./management/CHANGELOG.md)
    * **Enhancement** Implement basic cursor for paging results (`ApiReadAllEnvironmentsRequest`, `ApiReadAllGroupMembershipsForUserRequest`, `ApiReadAllGroupsRequest`, `ApiReadAllPopulationsRequest`, `ApiReadAllUsersRequest`). [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
    * **Enhancement** Add new DaVinci admin roles to `EnumRoleName` model. [#230](https://github.com/patrickcping/pingone-go-sdk-v2/pull/230)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.16.0](./mfa/CHANGELOG.md)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.9.0](./risk/CHANGELOG.md)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
    * **Enhancement** Add support for Suspicious device and Bot detection predictor types. [#229](https://github.com/patrickcping/pingone-go-sdk-v2/pull/229), [#231](https://github.com/patrickcping/pingone-go-sdk-v2/pull/231)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.3.0](./verify/CHANGELOG.md)
    * **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
    * **Enhancement** New voice configuration object option on the verify policy API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)
    * **Enhancement** New voice phrase API, which is a prerequisite to support the management of the voice configuration object via the verify policy API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)
    * **Enhancement** New voice phrase contents API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)

# Release (2023-07-18)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.1
    * **Bug** Fixed service hostname override client parameter validation that reported an invalid https address. [#225](https://github.com/patrickcping/pingone-go-sdk-v2/pull/225)

# Release (2023-07-12)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.10.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.2.0 => v0.2.1 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.2.0 => v0.2.1 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.2.0 => v0.2.1 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.23.0 => v0.24.0 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.14.0 => v0.15.0 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.8.0 => v0.8.1 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.2.0 => v0.2.1 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Note** bump `golang.org/x/oauth2` v0.9.0 => v0.10.0 [#224](https://github.com/patrickcping/pingone-go-sdk-v2/pull/224)
    * **Enhancement** Added ability to override the useragent at the point of client init. [#218](https://github.com/patrickcping/pingone-go-sdk-v2/pull/218)
    * **Enhancement** Add environment parameter defaults to client init. [#222](https://github.com/patrickcping/pingone-go-sdk-v2/pull/222)
    * **Enhancement** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.2.1](./agreementmanagement/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.2.1](./authorize/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.2.1](./credentials/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.24.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.15.0](./mfa/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
    * **Enhancement** Add optional ENUM attribute `NewDeviceNotification` to the `DeviceAuthenticationPolicy` model. [#215](https://github.com/patrickcping/pingone-go-sdk-v2/pull/215)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.8.1](./risk/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.2.1](./verify/CHANGELOG.md)
    * **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
    * **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
    * **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)

# Release (2023-07-04)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.9.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.13.0 => v0.14.0 [#214](https://github.com/patrickcping/pingone-go-sdk-v2/pull/214)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.14.0](./mfa/CHANGELOG.md)
    * **Breaking change** Corrected the `CreateDeviceAuthenticationPolicies` return data model object to return a new object `DeviceAuthenticationPolicyPostResponse` that reflects an `EntityArray` object returned on migration. [#213](https://github.com/patrickcping/pingone-go-sdk-v2/pull/213)
    * **Bug** Fixed the inability to select the content type header of the `CreateDeviceAuthenticationPolicies` API. [#213](https://github.com/patrickcping/pingone-go-sdk-v2/pull/213)

# Release (2023-07-04)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.9.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.1.4 => v0.2.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.7 => v0.2.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/credentials` v0.1.0 => v0.2.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.22.0 => v0.23.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.12.0 => v0.13.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.7.1 => v0.8.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/verify` v0.1.0 => v0.2.0 [#212](https://github.com/patrickcping/pingone-go-sdk-v2/pull/212)
    * **Enhancement** Added ability to override the default server and server variables at the point of client init. [#206](https://github.com/patrickcping/pingone-go-sdk-v2/pull/206)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.2.0](./agreementmanagement/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.2.0](./authorize/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.2.0](./credentials/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.23.0](./management/CHANGELOG.md)
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
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.13.0](./mfa/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Breaking change** Changed the request and response payload of the `DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies` API to account for the ability to migrate device authentication policies from legacy FIDO to upgraded FIDO2 policies. [#211](https://github.com/patrickcping/pingone-go-sdk-v2/pull/211)
    * **Breaking change** Migrated `CreateMFAPushCredentialRequest` and `UpdateMFAPushCredentialRequest` to `MFAPushCredentialRequest`. [#201](https://github.com/patrickcping/pingone-go-sdk-v2/pull/201)
    * **Breaking change** Migrated `EntityArrayEmbeddedPushCredentialsInner` to `MFAPushCredentialResponse`. [#201](https://github.com/patrickcping/pingone-go-sdk-v2/pull/201)
    * **Note** Deprecated old FIDO policy API. [#202](https://github.com/patrickcping/pingone-go-sdk-v2/pull/202)
    * **Feature** Support for upgraded FIDO2 policies API. [#202](https://github.com/patrickcping/pingone-go-sdk-v2/pull/202), [#209](https://github.com/patrickcping/pingone-go-sdk-v2/pull/209), [#210](https://github.com/patrickcping/pingone-go-sdk-v2/pull/210)
    * **Enhancement** Add support for migrating device authentication policies from legacy FIDO to upgraded FIDO2 policies. [#211](https://github.com/patrickcping/pingone-go-sdk-v2/pull/211)
    * **Enhancement** Add support for FIDO2 policies API in MFA Device Policies. [#202](https://github.com/patrickcping/pingone-go-sdk-v2/pull/202)
    * **Enhancement** Add optional attribute `PairingDisabled` to each MFA Device Policy device type. [#202](https://github.com/patrickcping/pingone-go-sdk-v2/pull/202), [#204](https://github.com/patrickcping/pingone-go-sdk-v2/pull/204)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.8.0](./risk/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.2.0](./verify/CHANGELOG.md)
    * **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
    * **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
    * **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# Release (2023-06-19)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.8.0
    * **Note** bump `github.com/golangci/golangci-lint` from v1.52.2 to v1.53.3 [#184](https://github.com/patrickcping/pingone-go-sdk-v2/pull/184) [#199](https://github.com/patrickcping/pingone-go-sdk-v2/pull/199)
    * **Note** bump `golang.org/x/oauth2` v0.8.0 => v0.9.0 [#199](https://github.com/patrickcping/pingone-go-sdk-v2/pull/199)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.11.0 => v0.12.0 [#199](https://github.com/patrickcping/pingone-go-sdk-v2/pull/199)
    * **Enhancement** Common client support for the new PingOne Verify module. [#191](https://github.com/patrickcping/pingone-go-sdk-v2/pull/191)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.12.0](./mfa/CHANGELOG.md)
    * **Note** Deprecated FCM key authentication for Google Play based mobile devices. [#196](https://github.com/patrickcping/pingone-go-sdk-v2/pull/196)
    * **Enhancement** Add support for Firebase Cloud Messaging for sending push messages for Google Play based mobile devices. [#196](https://github.com/patrickcping/pingone-go-sdk-v2/pull/196)
    * **Enhancement** Fix `TimeUnit` enum in the `DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime` model. [#190](https://github.com/patrickcping/pingone-go-sdk-v2/pull/190)
* `github.com/patrickcping/pingone-go-sdk-v2/verify` : [v0.1.0](./verify/CHANGELOG.md)
    * **Initial release** [#186](https://github.com/patrickcping/pingone-go-sdk-v2/pull/186), [#194](https://github.com/patrickcping/pingone-go-sdk-v2/pull/194)

# Release (2023-06-05)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.7.3
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.7.0 => v0.7.1 [#185](https://github.com/patrickcping/pingone-go-sdk-v2/pull/185)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.7.1](./risk/CHANGELOG.md)
    * **Bug** Fix the risk policy result for the `RiskPolicySet.DefaultResult` and the `RiskPolicyResult.Level` attribute. [#183](https://github.com/patrickcping/pingone-go-sdk-v2/pull/183)

# Release (2023-05-30)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.7.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.21.0 => v0.22.0 [#182](https://github.com/patrickcping/pingone-go-sdk-v2/pull/182)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.6.0 => v0.7.0 [#182](https://github.com/patrickcping/pingone-go-sdk-v2/pull/182)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.22.0](./management/CHANGELOG.md)
    * **Enhancement** Added `SloWindow` optional attribute to the Application SAML objects. [#179](https://github.com/patrickcping/pingone-go-sdk-v2/pull/179)
    * **Enhancement** Added SLO optional attributes to the SAML External Identity Provider object. [#179](https://github.com/patrickcping/pingone-go-sdk-v2/pull/179)
    * **Enhancement** Added `NewUserProvisioning` to the `SignOnPolicyActionLogin` model. [#181](https://github.com/patrickcping/pingone-go-sdk-v2/pull/181)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.7.0](./risk/CHANGELOG.md)
    * **Enhancement** Add `Triggers` to the `RiskPolicySet` model to support "staging" risk policies. [#180](https://github.com/patrickcping/pingone-go-sdk-v2/pull/180)
    * **Bug** Fix the `RiskPolicyResult` object, where the `Level` should only be `LOW`. [#178](https://github.com/patrickcping/pingone-go-sdk-v2/pull/178)

# Release (2023-05-22)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.7.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.20.0 => v0.21.0 [#177](https://github.com/patrickcping/pingone-go-sdk-v2/pull/177)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.10.0 => v0.11.0 [#177](https://github.com/patrickcping/pingone-go-sdk-v2/pull/177)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.21.0](./management/CHANGELOG.md)
    * **Note** Deprecated `bundleId` and `packageName` at the root level of the `ApplicationOIDC` model. Customers should use `mobile.bundleId` and `mobile.packageName` going forward. [#172](https://github.com/patrickcping/pingone-go-sdk-v2/pull/172)
    * **Feature** Support for Forms. [#176](https://github.com/patrickcping/pingone-go-sdk-v2/pull/176)
    * **Feature** Support for Forms Recaptcha configuration. [#176](https://github.com/patrickcping/pingone-go-sdk-v2/pull/176)
    * **Enhancement** Added `filterOptions.ipAddressExposed` and `filterOptions.userAgentExposed` to the `Subscription` (webhook) data model. [#173](https://github.com/patrickcping/pingone-go-sdk-v2/pull/173)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.11.0](./mfa/CHANGELOG.md)
    * **Enhancement** Support for `PhoneExtensions` in the `MFASettings` model. [#175](https://github.com/patrickcping/pingone-go-sdk-v2/pull/175)

# Release (2023-05-19)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.7.0
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.1.3 => v0.1.4 [#171](https://github.com/patrickcping/pingone-go-sdk-v2/pull/171)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.6 => v0.1.7 [#171](https://github.com/patrickcping/pingone-go-sdk-v2/pull/171)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.19.1 => v0.20.0 [#171](https://github.com/patrickcping/pingone-go-sdk-v2/pull/171)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.9.3 => v0.10.0 [#171](https://github.com/patrickcping/pingone-go-sdk-v2/pull/171)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.5.1 => v0.6.0 [#171](https://github.com/patrickcping/pingone-go-sdk-v2/pull/171)
    * **Note** bump `golang.org/x/oauth2` from v0.7.0 to v0.8.0 [#165](https://github.com/patrickcping/pingone-go-sdk-v2/pull/165)
    * **Enhancement** Common client support for the new PingOne Credentials module. [#169](https://github.com/patrickcping/pingone-go-sdk-v2/pull/169)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.1.4](./agreementmanagement/CHANGELOG.md)
    * **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.7](./authorize/CHANGELOG.md)
    * **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
* `github.com/patrickcping/pingone-go-sdk-v2/credentials` : [v0.1.0](./credentials/CHANGELOG.md)
    * **Initial release** [#166](https://github.com/patrickcping/pingone-go-sdk-v2/pull/166)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.20.0](./management/CHANGELOG.md)
    * **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
    * **Enhancement** Add support for enumerated values and regex validation to the schema attribute model. [#161](https://github.com/patrickcping/pingone-go-sdk-v2/pull/161)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.10.0](./mfa/CHANGELOG.md)
    * **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
    * **Enhancement** Support for `PairingKeyLifetime` and `PushLimit` in the `DeviceAuthenticationPolicyMobileApplicationsInner` model. [#159](https://github.com/patrickcping/pingone-go-sdk-v2/pull/159)
    * **Enhancement** Add support for setting `Email` quotas in notification policies. [#162](https://github.com/patrickcping/pingone-go-sdk-v2/pull/162)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.6.0](./risk/CHANGELOG.md)
    * **Breaking change** `RiskPolicySetRiskPoliciesInner` model changed to `RiskPolicy`. [#170](https://github.com/patrickcping/pingone-go-sdk-v2/pull/170)
    * **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
    * **Bug** Fix the Risk Predictor data model. [#164](https://github.com/patrickcping/pingone-go-sdk-v2/pull/164)
    * **Enhancement** Enhance the Risk Predictor data model. [#164](https://github.com/patrickcping/pingone-go-sdk-v2/pull/164)
    * **Enhancement** Support score based risk policies. [#170](https://github.com/patrickcping/pingone-go-sdk-v2/pull/170)

# Release (2023-04-28)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.4
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.1.2 => v0.1.3 [#158](https://github.com/patrickcping/pingone-go-sdk-v2/pull/158)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.5 => v0.1.6 [#158](https://github.com/patrickcping/pingone-go-sdk-v2/pull/158)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.19.0 => v0.19.1 [#158](https://github.com/patrickcping/pingone-go-sdk-v2/pull/158)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.9.2 => v0.9.3 [#158](https://github.com/patrickcping/pingone-go-sdk-v2/pull/158)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.5.0 => v0.5.1 [#158](https://github.com/patrickcping/pingone-go-sdk-v2/pull/158)
    * **Note** Remove unnecessary log output. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.1.3](./agreementmanagement/CHANGELOG.md)
    * **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.6](./authorize/CHANGELOG.md)
    * **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.19.1](./management/CHANGELOG.md)
    * **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.9.3](./mfa/CHANGELOG.md)
    * **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.5.1](./risk/CHANGELOG.md)
    * **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)

# Release (2023-04-24)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.3
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.1.1 => v0.1.2 [#155](https://github.com/patrickcping/pingone-go-sdk-v2/pull/155)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.4 => v0.1.5 [#155](https://github.com/patrickcping/pingone-go-sdk-v2/pull/155)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.18.0 => v0.19.0 [#155](https://github.com/patrickcping/pingone-go-sdk-v2/pull/155)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.9.1 => v0.9.2 [#155](https://github.com/patrickcping/pingone-go-sdk-v2/pull/155)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.4.0 => v0.5.0 [#155](https://github.com/patrickcping/pingone-go-sdk-v2/pull/155)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.1.2](./agreementmanagement/CHANGELOG.md)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.5](./authorize/CHANGELOG.md)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.19.0](./management/CHANGELOG.md)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
    * **Enhancement** Added support for the Google Play Integrity API, `GooglePlay` for the `ApplicationOIDC` object model.  *Breaking change note*: This is now a required attribute when configuring integrity detection on Android devices. [#153](https://github.com/patrickcping/pingone-go-sdk-v2/pull/153)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.9.2](./mfa/CHANGELOG.md)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.5.0](./risk/CHANGELOG.md)
    * **Breaking change** `RiskPredictor` model changed to `RiskPredictorCommon`, replaced with the predictor type model. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/151)
    * **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
    * **Bug fix** Fix for Risk Predictor API request/responses not including the individual predictor type data models. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/151)

# Release (2023-04-18)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.2
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.17.1 => v0.18.0 [#146](https://github.com/patrickcping/pingone-go-sdk-v2/pull/146)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.3.4 => v0.4.0 [#146](https://github.com/patrickcping/pingone-go-sdk-v2/pull/146)
    * **Note** bump `github.com/golangci/golangci-lint` from v1.52.0 to v1.52.2 [#138](https://github.com/patrickcping/pingone-go-sdk-v2/pull/138)
    * **Note** bump `golang.org/x/oauth2` from v0.6.0 to v0.7.0 [#140](https://github.com/patrickcping/pingone-go-sdk-v2/pull/140)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.18.0](./management/CHANGELOG.md)
    * **Enhancement** Add `CustomCRL` to the `Certificate` data model [#136](https://github.com/patrickcping/pingone-go-sdk-v2/pull/136)
    * **Enhancement** Add notifications policy country limit attributes for `NotificationsPolicy` model. [#142](https://github.com/patrickcping/pingone-go-sdk-v2/pull/142)
    * **Enhancement** Expand the `ApplicationAttributeMapping` model for attribute scoping. [#143](https://github.com/patrickcping/pingone-go-sdk-v2/pull/143)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.4.0](./risk/CHANGELOG.md)
    * **Feature** Full support for risk predictors. [#135](https://github.com/patrickcping/pingone-go-sdk-v2/pull/135)

# Release (2023-03-20)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.6.1
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` v0.1.0 => v0.1.1 [#134](https://github.com/patrickcping/pingone-go-sdk-v2/pull/134)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/authorize` v0.1.3 => v0.1.4 [#134](https://github.com/patrickcping/pingone-go-sdk-v2/pull/134)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.17.0 => v0.17.1 [#134](https://github.com/patrickcping/pingone-go-sdk-v2/pull/134)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/management` v0.16.0 => v0.17.0 [#132](https://github.com/patrickcping/pingone-go-sdk-v2/pull/132)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/mfa` v0.9.0 => v0.9.1 [#134](https://github.com/patrickcping/pingone-go-sdk-v2/pull/134)
    * **Note** bump `github.com/patrickcping/pingone-go-sdk-v2/risk` v0.3.3 => v0.3.4 [#134](https://github.com/patrickcping/pingone-go-sdk-v2/pull/134)
    * **Note** bump `github.com/golangci/golangci-lint` from v1.51.2 to v1.52.0 [#129](https://github.com/patrickcping/pingone-go-sdk-v2/pull/129)
* `github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement` : [v0.1.1](./agreementmanagement/CHANGELOG.md)
    * **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)
* `github.com/patrickcping/pingone-go-sdk-v2/authorize` : [v0.1.4](./authorize/CHANGELOG.md)
    * **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.17.1](./management/CHANGELOG.md)
    * **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.17.0](./management/CHANGELOG.md)
    * **Note** Added `FlowPolicy` and `FlowPolicyAssignment` data model descriptions. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `priority` attribute in the `FlowPolicyAssignment` data model corrected to be a required field, with a minimum value of `1`. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `flowPolicy` attribute in the `FlowPolicyAssignment` data model corrected to the `FlowPolicyAssignmentFlowPolicy` data model. [#127](https://github.com/patrickcping/pingone-go-sdk-v2/pull/127)
    * **Bug fix** `idpSigning.algorithm` attribute in the `ApplicationSAML` data model corrected to be writable, set as ENUM.
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.9.1](./mfa/CHANGELOG.md)
    * **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.4](./risk/CHANGELOG.md)
    * **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)

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
