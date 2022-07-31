# Release (Unreleased)

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
    * **Bug fix** - Corrections to `SignOnPolicyActionCommonCondition` model

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