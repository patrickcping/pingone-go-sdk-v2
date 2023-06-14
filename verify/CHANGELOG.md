# v0.1.1 (2023-06-13)

* **Breaking change** To align with underlying API, removed `HOURS` enum option in `EnumLongTimeUnit`, and replaced `EnumLongTimeUnit` and `EnumShortTimeUnit` with the simplified `EnumTimeUnit` model which contains `MINUTES` and `SECONDS` only. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/193)
* **Breaking change** `EmailPhoneConfiguration` model changed to the more intuitive naming of `OTPDeviceConfiguration`. [#151](https://github.com/patrickcping/pingone-go-sdk-v2/pull/193)
* **Note** Aforementioned are listed as breaking changes for proper documentation purposes, however the corrections were made prior to the deployment of the initial service using P1Verify Verify Policy capabilities.

# v0.1.0 (2023-05-26)

* **Initial release** [#186](https://github.com/patrickcping/pingone-go-sdk-v2/pull/186)
