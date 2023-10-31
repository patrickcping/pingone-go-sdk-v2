# v0.5.0 (Unreleased)

* **Enhancement** Added the `Default` property to the `CredentialTypeMetaDataFieldsInner` object model. [#275](https://github.com/patrickcping/pingone-go-sdk-v2/pull/275)
* **Enhancement** Added `Multiple` object attribute to the `CredentialType` object model. [#275](https://github.com/patrickcping/pingone-go-sdk-v2/pull/275)

# v0.4.1 (2023-10-16)

* **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)

# v0.4.0 (2023-08-22)

* **Enhancement** Add parameters `OnDelete` to the `CredentialType` model. [#240](https://github.com/patrickcping/pingone-go-sdk-v2/pull/240)
* **Enhancement** Add parameters `FileSupport` to the `CredentialTypeMetaDataFieldsInner` model. [#240](https://github.com/patrickcping/pingone-go-sdk-v2/pull/240)

# v0.3.1 (2023-08-15)

* **Bug** Corrected situations where `EntityArrayEmbeddedItemsInner` object dereferencing led to panic error. [#237](https://github.com/patrickcping/pingone-go-sdk-v2/pull/237)

# v0.3.0 (2023-08-08)

* **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)

# v0.2.1 (2023-07-12)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)

# v0.2.0 (2023-07-04)

* **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
* **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# v0.1.0 (2023-05-19)

* **Initial release** [#166](https://github.com/patrickcping/pingone-go-sdk-v2/pull/166)
