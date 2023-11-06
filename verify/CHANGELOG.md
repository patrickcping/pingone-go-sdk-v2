# v0.4.0 (2023-11-06)

* **Enhancement** Added the `InspectionType` property to the `GovernmentIdConfiguration` object model. [#278](https://github.com/patrickcping/pingone-go-sdk-v2/pull/278)

# v0.3.1 (2023-10-16)

* **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)

# v0.3.0 (2023-08-08)

* **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
* **Enhancement** New voice configuration object option on the verify policy API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)
* **Enhancement** New voice phrase API, which is a prerequisite to support the management of the voice configuration object via the verify policy API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)
* **Enhancement** New voice phrase contents API. [#228](https://github.com/patrickcping/pingone-go-sdk-v2/pull/228)

# v0.2.1 (2023-07-12)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)

# v0.2.0 (2023-07-04)

* **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
* **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# v0.1.0 (2023-06-19)

* **Initial release** [#186](https://github.com/patrickcping/pingone-go-sdk-v2/pull/186), [#194](https://github.com/patrickcping/pingone-go-sdk-v2/pull/194)
