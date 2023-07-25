# \IntegrationCatalogApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDIntegrationsGet**](IntegrationCatalogApi.md#EnvironmentsEnvironmentIDIntegrationsGet) | **Get** /environments/{environmentID}/integrations | READ Integration Metadata
[**EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet**](IntegrationCatalogApi.md#EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet) | **Get** /environments/{environmentID}/integrations/{integrationID} | READ One Integration Metadata
[**EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet**](IntegrationCatalogApi.md#EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions | READ Integration Version Metadata
[**EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet**](IntegrationCatalogApi.md#EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID}/asset | READ Integration Version Asset Download
[**EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet**](IntegrationCatalogApi.md#EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID} | READ One Integration Version Metadata



## EnvironmentsEnvironmentIDIntegrationsGet

> EnvironmentsEnvironmentIDIntegrationsGet(ctx, environmentID).Execute()

READ Integration Metadata

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
    r, err := apiClient.IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDIntegrationsGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet

> EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet(ctx, environmentID, integrationID).Execute()

READ One Integration Metadata

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDIntegrationsIntegrationIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet

> EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet(ctx, environmentID, integrationID).Execute()

READ Integration Version Metadata

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet

> EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(ctx, environmentID, integrationID, integrationVersionID).Execute()

READ Integration Version Asset Download

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDAssetGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet

> EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(ctx, environmentID, integrationID, integrationVersionID).Execute()

READ One Integration Version Metadata

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.EnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDIntegrationsIntegrationIDVersionsIntegrationVersionIDGetRequest struct via the builder pattern


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

