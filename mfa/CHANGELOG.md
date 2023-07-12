# v0.15.0 (2023-07-12)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)
* **Enhancement** Add optional ENUM attribute `NewDeviceNotification` to the `DeviceAuthenticationPolicy` model. [#215](https://github.com/patrickcping/pingone-go-sdk-v2/pull/215)

# v0.14.0 (2023-07-04)

* **Breaking change** Corrected the `CreateDeviceAuthenticationPolicies` return data model object to return a new object `DeviceAuthenticationPolicyPostResponse` that reflects an `EntityArray` object returned on migration. [#213](https://github.com/patrickcping/pingone-go-sdk-v2/pull/213)
* **Bug** Fixed the inability to select the content type header of the `CreateDeviceAuthenticationPolicies` API. [#213](https://github.com/patrickcping/pingone-go-sdk-v2/pull/213)

# v0.13.0 (2023-07-04)

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

# v0.12.0 (2023-06-19)

* **Note** Deprecated FCM key authentication for Google Play based mobile devices. [#196](https://github.com/patrickcping/pingone-go-sdk-v2/pull/196)
* **Enhancement** Add support for Firebase Cloud Messaging for sending push messages for Google Play based mobile devices. [#196](https://github.com/patrickcping/pingone-go-sdk-v2/pull/196)
* **Enhancement** Fix `TimeUnit` enum in the `DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime` model. [#190](https://github.com/patrickcping/pingone-go-sdk-v2/pull/190)

# v0.11.0 (2023-05-23)

* **Enhancement** Support for `PhoneExtensions` in the `MFASettings` model. [#175](https://github.com/patrickcping/pingone-go-sdk-v2/pull/175)

# v0.10.0 (2023-05-19)

* **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
* **Enhancement** Support for `PairingKeyLifetime` and `PushLimit` in the `DeviceAuthenticationPolicyMobileApplicationsInner` model. [#159](https://github.com/patrickcping/pingone-go-sdk-v2/pull/159)

# v0.9.3 (2023-04-28)

* **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)

# v0.9.2 (2023-04-24)

* **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)

# v0.9.1 (2023-03-20)

* **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)

# v0.9.0 (2023-03-10)

* **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)
* **Feature** Support for "User MFA Enabled" API and data model [#113](https://github.com/patrickcping/pingone-go-sdk-v2/pull/113)

# v0.8.0 (2023-02-22)

* **Note** bump `golang.org/x/net` v0.5.0 => v0.7.0 [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* **Note** bump `golang.org/x/oauth2` v0.4.0 => v0.5.0 [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* **Breaking change** `Key` property removed from the `MFAPushCredential` model object and assigned to `MFAPushCredentialFCM` and `MFAPushCredentialAPNS` individually [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)
* **Enhancement** Add support for Huawei HMS Push service [#105](https://github.com/patrickcping/pingone-go-sdk-v2/pull/105)

# v0.7.2 (2023-01-12)

* **Breaking change** Device selection `Authentication` model no longer required for the `MFASettings` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
* **Note** Deprecated device selection from the `MFASettings` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
* **Enhancement** - Added device selection parameters to the `DeviceAuthenticationPolicy` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)
* **Enhancement** - Added push timeout to the `DeviceAuthenticationPolicy` model [#98](https://github.com/patrickcping/pingone-go-sdk-v2/pull/98)

# v0.7.1 (2023-01-09)

* **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)

# v0.7.0 (2022-11-06)

* **Feature** Added FIDO Policy API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)
* **Feature** Added FIDO Custom Device Metadata API and model [#75](https://github.com/patrickcping/pingone-go-sdk-v2/pull/75)

# v0.6.1 (2022-10-10)

* **Bug fix** `lockout` made optional in the `MFASettings` model [#70](https://github.com/patrickcping/pingone-go-sdk-v2/pull/70)

# v0.6.0 (2022-10-09)

* **Bug fix** Corrected the Device policy API [#65](https://github.com/patrickcping/pingone-go-sdk-v2/pull/65)
* **Bug fix** Corrected the Application push credentials API model [#67](https://github.com/patrickcping/pingone-go-sdk-v2/pull/67)

# v0.5.1 (2022-09-18)

* **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredential201Response` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `CreateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `EntityArrayEmbeddedPushCredentialsInner` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Changed model dereferencing strategy for the `UpdateMFAPushCredentialRequest` model [#48](https://github.com/patrickcping/pingone-go-sdk-v2/pull/48)
* **Enhancement** - Add `environment` attribute block to `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)
* **Enhancement** - Add required attributes to the `MFASettings` model [#50](https://github.com/patrickcping/pingone-go-sdk-v2/pull/50)

# v0.5.0 (2022-09-11)

* **Feature** Added MFA Settings model [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)
* **Breaking change** `updatedAt` attributes changed to date/time data type [#44](https://github.com/patrickcping/pingone-go-sdk-v2/pull/44)

# v0.4.0 (2022-08-29)

* **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)

# v0.3.0 (2022-08-05)

* **Feature** Mfa device authentication policies [#13](https://github.com/patrickcping/pingone-go-sdk-v2/pull/13)

# v0.2.0 (2022-07-17)

* **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
* **Feature** Full support for enum values

# v0.1.0 (2022-07-16)

Initial release - rebasing versions to reflect module stability