# \CapabilitiesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDCapabilitiesGet**](CapabilitiesApi.md#EnvironmentsEnvironmentIDCapabilitiesGet) | **Get** /environments/{environmentID}/capabilities | READ Environment Capabilities
[**OrganizationsOrganizationIDCapabilitiesGet**](CapabilitiesApi.md#OrganizationsOrganizationIDCapabilitiesGet) | **Get** /organizations/{organizationID}/capabilities | READ Organization Capabilities



## EnvironmentsEnvironmentIDCapabilitiesGet

> EnvironmentsEnvironmentIDCapabilitiesGet(ctx, environmentID).Execute()

READ Environment Capabilities

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
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CapabilitiesApi.EnvironmentsEnvironmentIDCapabilitiesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesApi.EnvironmentsEnvironmentIDCapabilitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDCapabilitiesGetRequest struct via the builder pattern


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


## OrganizationsOrganizationIDCapabilitiesGet

> OrganizationsOrganizationIDCapabilitiesGet(ctx, organizationID).Execute()

READ Organization Capabilities

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CapabilitiesApi.OrganizationsOrganizationIDCapabilitiesGet(context.Background(), organizationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesApi.OrganizationsOrganizationIDCapabilitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationsOrganizationIDCapabilitiesGetRequest struct via the builder pattern


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

