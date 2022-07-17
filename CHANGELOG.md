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