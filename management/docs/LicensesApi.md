# \LicensesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1OrganizationsOrganizationIDLicensesGet**](LicensesApi.md#V1OrganizationsOrganizationIDLicensesGet) | **Get** /v1/organizations/{organizationID}/licenses | READ All Licenses
[**V1OrganizationsOrganizationIDLicensesLicenseIDGet**](LicensesApi.md#V1OrganizationsOrganizationIDLicensesLicenseIDGet) | **Get** /v1/organizations/{organizationID}/licenses/{licenseID} | READ One License
[**V1OrganizationsOrganizationIDLicensesLicenseIDNameGet**](LicensesApi.md#V1OrganizationsOrganizationIDLicensesLicenseIDNameGet) | **Get** /v1/organizations/{organizationID}/licenses/{licenseID}/name | READ One License Name
[**V1OrganizationsOrganizationIDLicensesLicenseIDNamePut**](LicensesApi.md#V1OrganizationsOrganizationIDLicensesLicenseIDNamePut) | **Put** /v1/organizations/{organizationID}/licenses/{licenseID}/name | Update One License Name



## V1OrganizationsOrganizationIDLicensesGet

> V1OrganizationsOrganizationIDLicensesGet(ctx, organizationID).Execute()

READ All Licenses

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationID := "organizationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.V1OrganizationsOrganizationIDLicensesGet(context.Background(), organizationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.V1OrganizationsOrganizationIDLicensesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDLicensesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsOrganizationIDLicensesLicenseIDGet

> V1OrganizationsOrganizationIDLicensesLicenseIDGet(ctx, organizationID, licenseID).Execute()

READ One License

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDGet(context.Background(), organizationID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDLicensesLicenseIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsOrganizationIDLicensesLicenseIDNameGet

> V1OrganizationsOrganizationIDLicensesLicenseIDNameGet(ctx, organizationID, licenseID).Execute()

READ One License Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDNameGet(context.Background(), organizationID, licenseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDLicensesLicenseIDNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1OrganizationsOrganizationIDLicensesLicenseIDNamePut

> V1OrganizationsOrganizationIDLicensesLicenseIDNamePut(ctx, organizationID, licenseID).Body(body).Execute()

Update One License Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDNamePut(context.Background(), organizationID, licenseID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensesApi.V1OrganizationsOrganizationIDLicensesLicenseIDNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDLicensesLicenseIDNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

