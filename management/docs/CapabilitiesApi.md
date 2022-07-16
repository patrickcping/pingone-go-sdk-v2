# \CapabilitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDCapabilitiesGet**](CapabilitiesApi.md#V1EnvironmentsEnvironmentIDCapabilitiesGet) | **Get** /v1/environments/{environmentID}/capabilities | READ Environment Capabilities
[**V1OrganizationsOrganizationIDCapabilitiesGet**](CapabilitiesApi.md#V1OrganizationsOrganizationIDCapabilitiesGet) | **Get** /v1/organizations/{organizationID}/capabilities | READ Organization Capabilities



## V1EnvironmentsEnvironmentIDCapabilitiesGet

> V1EnvironmentsEnvironmentIDCapabilitiesGet(ctx, environmentID).Execute()

READ Environment Capabilities

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
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CapabilitiesApi.V1EnvironmentsEnvironmentIDCapabilitiesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesApi.V1EnvironmentsEnvironmentIDCapabilitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCapabilitiesGetRequest struct via the builder pattern


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


## V1OrganizationsOrganizationIDCapabilitiesGet

> V1OrganizationsOrganizationIDCapabilitiesGet(ctx, organizationID).Execute()

READ Organization Capabilities

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
    resp, r, err := apiClient.CapabilitiesApi.V1OrganizationsOrganizationIDCapabilitiesGet(context.Background(), organizationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesApi.V1OrganizationsOrganizationIDCapabilitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDCapabilitiesGetRequest struct via the builder pattern


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

