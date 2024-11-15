# \LicensesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllLicenses**](LicensesApi.md#ReadAllLicenses) | **Get** /organizations/{organizationID}/licenses | READ All Licenses
[**ReadOneLicense**](LicensesApi.md#ReadOneLicense) | **Get** /organizations/{organizationID}/licenses/{licenseID} | READ One License
[**ReadOneLicenseName**](LicensesApi.md#ReadOneLicenseName) | **Get** /organizations/{organizationID}/licenses/{licenseID}/name | READ One License Name
[**UpdateOneLicenseName**](LicensesApi.md#UpdateOneLicenseName) | **Put** /organizations/{organizationID}/licenses/{licenseID}/name | Update One License Name



## ReadAllLicenses

> EntityArrayPagedIterator ReadAllLicenses(ctx, organizationID).Filter(filter).Order(order).Execute()

READ All Licenses

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    organizationID := "organizationID_example" // string | 
    filter := "filter=beginsAt lt "{{now}}"" // string | For organizations with more than one license, you can use a filter to return the list of licenses in descending order. (optional)
    order := "order=-beginsAt" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.ReadAllLicenses(context.Background(), organizationID).Filter(filter).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.ReadAllLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllLicenses`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.ReadAllLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | For organizations with more than one license, you can use a filter to return the list of licenses in descending order. | 
 **order** | **string** |  | 

### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneLicense

> License ReadOneLicense(ctx, organizationID, licenseID).Execute()

READ One License

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.ReadOneLicense(context.Background(), organizationID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.ReadOneLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneLicense`: License
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.ReadOneLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**License**](License.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneLicenseName

> LicenseName ReadOneLicenseName(ctx, organizationID, licenseID).Execute()

READ One License Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.ReadOneLicenseName(context.Background(), organizationID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.ReadOneLicenseName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneLicenseName`: LicenseName
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.ReadOneLicenseName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneLicenseNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LicenseName**](LicenseName.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOneLicenseName

> LicenseName UpdateOneLicenseName(ctx, organizationID, licenseID).LicenseName(licenseName).Execute()

Update One License Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 
    licenseName := *openapiclient.NewLicenseName("Name_example") // LicenseName |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.UpdateOneLicenseName(context.Background(), organizationID, licenseID).LicenseName(licenseName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.UpdateOneLicenseName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOneLicenseName`: LicenseName
    fmt.Fprintf(os.Stdout, "Response from `LicensesApi.UpdateOneLicenseName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOneLicenseNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **licenseName** | [**LicenseName**](LicenseName.md) |  | 

### Return type

[**LicenseName**](LicenseName.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

