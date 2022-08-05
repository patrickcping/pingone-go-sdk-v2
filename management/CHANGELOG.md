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