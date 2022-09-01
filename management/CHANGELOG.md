# v0.7.0 (Unreleased)

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