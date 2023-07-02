# v0.8.0 (Unreleased)

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