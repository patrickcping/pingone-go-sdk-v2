# \IntegrationCatalogApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDIntegrationsGet**](IntegrationCatalogApi.md#V1EnvironmentsEnvironmentIDIntegrationsGet) | **Get** /v1/environments/{environmentID}/integrations | READ Integration Metadata
[**V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet**](IntegrationCatalogApi.md#V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet) | **Get** /v1/environments/{environmentID}/integrations/{integrationID} | READ One Integration Metadata
[**V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet**](IntegrationCatalogApi.md#V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet) | **Get** /v1/environments/{environmentID}/integrations/{integrationID}/versions | READ Integration Version Metadata
[**V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet**](IntegrationCatalogApi.md#V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet) | **Get** /v1/environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID}/asset | READ Integration Version Asset Download
[**V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet**](IntegrationCatalogApi.md#V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet) | **Get** /v1/environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID} | READ One Integration Version Metadata



## V1EnvironmentsEnvironmentIDIntegrationsGet

> V1EnvironmentsEnvironmentIDIntegrationsGet(ctx, environmentID).Execute()

READ Integration Metadata

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
    resp, r, err := apiClient.IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIntegrationsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet

> V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet(ctx, environmentID, integrationID).Execute()

READ One Integration Metadata

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIntegrationsIntegrationIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet

> V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet(ctx, environmentID, integrationID).Execute()

READ Integration Version Metadata

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet

> V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(ctx, environmentID, integrationID, integrationVersionID).Execute()

READ Integration Version Asset Download

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet

> V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(ctx, environmentID, integrationID, integrationVersionID).Execute()

READ One Integration Version Metadata

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.V1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGetRequest struct via the builder pattern


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

