# v0.19.1 (Unreleased)

* **Note** Add backoff retry on transient 404 errors expected after environment creation.

# v0.19.0 (2025-02-10)

* **Note** Enhance backoff retry logic for transient errors according to [best practices](https://apidocs.pingidentity.com/pingone/platform/v1/api/#retries-best-practice-for-managing-transient-api-errors). [#418](https://github.com/patrickcping/pingone-go-sdk-v2/pull/418) [#421](https://github.com/patrickcping/pingone-go-sdk-v2/pull/421)
* **Enhancement** Added API functions for `X-Ping-External-Transaction-ID` and `X-Ping-External-Session-ID` for transaction and session telemetry. [#419](https://github.com/patrickcping/pingone-go-sdk-v2/pull/419)

# v0.18.0 (2024-12-18)

* **Bug** Add missing protect predictor composite conditions `IP_RANGE` and `IP_COMPARISON`. [#401](https://github.com/patrickcping/pingone-go-sdk-v2/pull/401)

# v0.17.0 (2024-11-15)

* **Breaking change** `(Api[a-zA-Z]Request).Execute()` and `(*Api[a-zA-Z]Request).[a-zA-Z]Execute()` API functions now returns the `EntityArrayPagedIterator` data type to for code clients to implement paging of results. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
* **Note** Upgraded go version to 1.23 to align with the go [release policy](https://go.dev/doc/devel/release#policy). [#395](https://github.com/patrickcping/pingone-go-sdk-v2/pull/395)
* **Enhancement** Added `(Api[a-zA-Z]Request).ExecuteInitialPage()` and `(*Api[a-zA-Z]Request).[a-zA-Z]ExecuteInitialPage()` API functions to just return the initial page of a paged response. [#392](https://github.com/patrickcping/pingone-go-sdk-v2/pull/392)
* **Enhancement** Added the `ShouldValidatePayloadSignature` field to the `RiskPredictorDevice` model. [#380](https://github.com/patrickcping/pingone-go-sdk-v2/pull/380)
* **Enhancement** Added the `IncludeRepeatedEventsWithoutSdk` field to the `RiskPredictorBotDetection` model. [#382](https://github.com/patrickcping/pingone-go-sdk-v2/pull/382)
* **Enhancement** Support the `TRAFFIC_ANOMALY` risk predictor type (new `RiskPredictorTrafficAnomaly` model). [#383](https://github.com/patrickcping/pingone-go-sdk-v2/pull/383)

# v0.16.0 (2024-07-04)

* **Breaking change** `Links` objects changed from `LinksHATEOAS` object to `map[string]LinksHATEOASValue` type. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)
* **Enhancement** Added the `HALApi` service to be able to follow HAL links in responses. [#366](https://github.com/patrickcping/pingone-go-sdk-v2/pull/366)

# v0.15.1 (2024-06-18)

* **Bug** Corrected `DomainWhiteList` field in the `RiskPredictorAdversaryInTheMiddle` model. [#361](https://github.com/patrickcping/pingone-go-sdk-v2/pull/361)

# v0.15.0 (2024-06-07)

* **Enhancement** Add the `com.au` top level domain to the connection configuration. [#351](https://github.com/patrickcping/pingone-go-sdk-v2/pull/351)

# v0.14.1 (2024-03-25)

* **Bug** Fixed error "Error when calling `ReadAllRiskPredictors`: Data failed to match schemas in oneOf(RiskPredictor)" when the `EMAIL_REPUTATION` predictor type is in the API's response payload. [#332](https://github.com/patrickcping/pingone-go-sdk-v2/pull/332)

# v0.14.0 (2024-03-15)

* **Enhancement** Added the `EMAIL_REPUTATION` predictor type. [#330](https://github.com/patrickcping/pingone-go-sdk-v2/pull/330)

# v0.13.0 (2024-02-23)

* **Enhancement** Added the `RiskPredictorAdversaryInTheMiddle` property to the `RiskPredictor` object model and extended the `EnumPredictorType` enum with `ADVERSARY_IN_THE_MIDDLE` to support the adversary in the middle predictor. [#321](https://github.com/patrickcping/pingone-go-sdk-v2/pull/321)

# v0.12.2 (2024-01-12)

* **Note** Upgrade GO to `v1.21` [#306](https://github.com/patrickcping/pingone-go-sdk-v2/pull/306)
* **Note** Updated the default UserAgent string format. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)
* **Enhancement** Add ability to append custom user-agent information to HTTP requests. [#305](https://github.com/patrickcping/pingone-go-sdk-v2/pull/305)

# v0.12.1 (2023-12-27)

* **Note** Remove redundant data models and documentation. [#300](https://github.com/patrickcping/pingone-go-sdk-v2/pull/300)

# v0.12.0 (2023-11-10)

* **Enhancement** Add API error handling for `409` errors. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)
* **Enhancement** Added the `ReferencedValues` property to the `P1ErrorDetailsInnerInnerError` object model. [#279](https://github.com/patrickcping/pingone-go-sdk-v2/pull/279)

# v0.11.0 (2023-10-16)

* **Note** Unskip read only attributes when converting API models to maps. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* **Note** Updated the default UserAgent string format. [#255](https://github.com/patrickcping/pingone-go-sdk-v2/pull/255)
* **Enhancement** Enhance the `EnumFlowType` with new supported values. [#256](https://github.com/patrickcping/pingone-go-sdk-v2/pull/256)
* **Enhancement** Added the `Compositions` array to the `RiskPredictorComposite` model, to replace the `Composition` attribute that has now been deprecated. [#256](https://github.com/patrickcping/pingone-go-sdk-v2/pull/256)

# v0.10.0 (2023-09-15)

* **Enhancement** Add support for a Bot detection specific data model. [#252](https://github.com/patrickcping/pingone-go-sdk-v2/pull/252)

# v0.9.0 (2023-08-08)

* **Enhancement** Implement HATEOAS links for API response objects. [#227](https://github.com/patrickcping/pingone-go-sdk-v2/pull/227)
* **Enhancement** Add support for Suspicious device and Bot detection predictor types. [#229](https://github.com/patrickcping/pingone-go-sdk-v2/pull/229), [#231](https://github.com/patrickcping/pingone-go-sdk-v2/pull/231)

# v0.8.1 (2023-07-12)

* **Note** Code optimisation for API response processing. [#216](https://github.com/patrickcping/pingone-go-sdk-v2/pull/216)
* **Note** Allow user-defined values for the `UserAgent` configuration parameter. [#217](https://github.com/patrickcping/pingone-go-sdk-v2/pull/217)
* **Note** Add parameter to explicitly define http proxy support without the standard environment variables. [#223](https://github.com/patrickcping/pingone-go-sdk-v2/pull/223)

# v0.8.0 (2023-07-04)

* **Note** Now suppresses errors when attempting to unmarshal an ENUM value from JSON that isn't yet supported in the SDK.  The value is now returned as `UNKNOWN`. [#208](https://github.com/patrickcping/pingone-go-sdk-v2/pull/208)
* **Enhancement** Add parameters `protocol`, `baseDomain` and `baseHostname` to server configuration. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)
* **Enhancement** Add ability to set default server index and a server variable default value on the client configuration as an alternative option to setting them in the context. [#205](https://github.com/patrickcping/pingone-go-sdk-v2/pull/205)

# v0.7.1 (2023-06-05)

* **Bug** Fix the risk policy result for the `RiskPolicySet.DefaultResult` and the `RiskPolicyResult.Level` attribute. [#183](https://github.com/patrickcping/pingone-go-sdk-v2/pull/183)

# v0.7.0 (2023-05-30)

* **Enhancement** Add `Triggers` to the `RiskPolicySet` model to support "staging" risk policies. [#180](https://github.com/patrickcping/pingone-go-sdk-v2/pull/180)
* **Bug** Fix the `RiskPolicyResult` object, where the `Level` should only be `LOW`. [#178](https://github.com/patrickcping/pingone-go-sdk-v2/pull/178)

# v0.6.0 (2023-05-19)

* **Breaking change** `RiskPolicySetRiskPoliciesInner` model changed to `RiskPolicy`. [#170](https://github.com/patrickcping/pingone-go-sdk-v2/pull/170)
* **Note** Change default useragent for HTTP requests. [#160](https://github.com/patrickcping/pingone-go-sdk-v2/pull/160)
* **Bug** Fix the Risk Predictor data model. [#164](https://github.com/patrickcping/pingone-go-sdk-v2/pull/164)
* **Enhancement** Enhance the Risk Predictor data model. [#164](https://github.com/patrickcping/pingone-go-sdk-v2/pull/164)
* **Enhancement** Support score based risk policies. [#170](https://github.com/patrickcping/pingone-go-sdk-v2/pull/170)

# v0.5.1 (2023-04-28)

* **Bug** Fix for retry for conditions based on PingOne error response object. [#157](https://github.com/patrickcping/pingone-go-sdk-v2/pull/157)

# v0.5.0 (2023-04-24)

* **Breaking change** `RiskPredictor` model changed to `RiskPredictorCommon`, replaced with the predictor type model. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/151)
* **Note** Add retry logic for retryable HTTP status codes. [#147](https://github.com/patrickcping/pingone-go-sdk-v2/pull/147)
* **Bug fix** Fix for Risk Predictor API request/responses not including the individual predictor type data models. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/151)

# v0.4.0 (2023-04-18)

* **Feature** Full support for risk predictors. [#135](https://github.com/patrickcping/pingone-go-sdk-v2/pull/135)

# v0.3.4 (2023-03-20)

* **Note** Adjust client request/response debug and fix linting issues. [#133](https://github.com/patrickcping/pingone-go-sdk-v2/pull/133)

# v0.3.3 (2023-03-10)

* **Note** bump codegen v6.2.1 => v6.4.0 [#122](https://github.com/patrickcping/pingone-go-sdk-v2/pull/122)

# v0.3.2 (2023-02-22)

* **Note** bump `golang.org/x/net` v0.2.0 => v0.7.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)
* **Note** bump `golang.org/x/oauth2` v0.2.0 => v0.5.0 [#109](https://github.com/patrickcping/pingone-go-sdk-v2/pull/109)

# v0.3.1 (2023-01-09)

* **Note** bump codegen v6.2.0 => v6.2.1 [#84](https://github.com/patrickcping/pingone-go-sdk-v2/pull/84)

# v0.3.0 (2022-08-29)

* **Enhancement** Add 404 response processing for all requests [#28](https://github.com/patrickcping/pingone-go-sdk-v2/pull/28)
* **Enhancement** Add missing 400 response processing [#29](https://github.com/patrickcping/pingone-go-sdk-v2/pull/29)

# v0.2.0 (2022-07-17)

* **Security** Bump `golang.org/x/net` for `CVE-2021-33194` : https://deps.dev/advisory/osv/GO-2021-0238
* **Feature** Full support for enum values

# v0.1.0 (2022-07-16)

Initial release - rebasing versions to reflect module stability