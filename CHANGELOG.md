# Release (Unreleased)

* `github.com/patrickcping/pingone-go-sdk-v2` : v0.3.4
    * **Note** bump `github.com/golangci/golangci-lint` v1.47.2 => v1.49.0
    * **Note** bump `github.com/securego/gosec/v2` v2.12.0 => v2.13.1
* `github.com/patrickcping/pingone-go-sdk-v2/management` : [v0.6.0](./management/CHANGELOG.md)
    * **Bug fix** - Correct `PINGONE_API` from incorrect `PING_ONE_API` in `EnumResourceType` enum
    * **Feature** Support for Custom Domains [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
    * **Feature** Support for Keypairs [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
    * **Feature** Support for Certificates [#26](https://github.com/patrickcping/pingone-go-sdk-v2/pull/26)
    * **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* `github.com/patrickcping/pingone-go-sdk-v2/mfa` : [v0.4.0](./management/CHANGELOG.md)
    * **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* `github.com/patrickcping/pingone-go-sdk-v2/risk` : [v0.3.0](./management/CHANGELOG.md)
    * **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)

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