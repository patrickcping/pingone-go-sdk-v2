# \IdentityProviderManagementIdentityProvidersApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDIdentityProvidersGet**](IdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersGet) | **Get** /v1/environments/{environmentID}/identityProviders | READ All Identity Providers
[**V1EnvironmentsEnvironmentIDIdentityProvidersPost**](IdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersPost) | **Post** /v1/environments/{environmentID}/identityProviders | Discover OpenID Provider Metadata
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete**](IdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete) | **Delete** /v1/environments/{environmentID}/identityProviders/{providerID} | DELETE Identity Provider
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet**](IdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet) | **Get** /v1/environments/{environmentID}/identityProviders/{providerID} | READ One Identity Provider
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut**](IdentityProviderManagementIdentityProvidersApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut) | **Put** /v1/environments/{environmentID}/identityProviders/{providerID} | UPDATE Identity Provider



## V1EnvironmentsEnvironmentIDIdentityProvidersGet

> V1EnvironmentsEnvironmentIDIdentityProvidersGet(ctx, environmentID).Execute()

READ All Identity Providers

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
    resp, r, err := apiClient.IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersPost

> V1EnvironmentsEnvironmentIDIdentityProvidersPost(ctx, environmentID).ContentType(contentType).Body(body).Execute()

Discover OpenID Provider Metadata

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
    contentType := "application/vnd.pingidentity.openid-configuration.discover+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersPost(context.Background(), environmentID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete(ctx, environmentID, providerID).Execute()

DELETE Identity Provider

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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet(ctx, environmentID, providerID).Execute()

READ One Identity Provider

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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut(ctx, environmentID, providerID).Body(body).Execute()

UPDATE Identity Provider

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
    providerID := "providerID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut(context.Background(), environmentID, providerID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProvidersApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDPutRequest struct via the builder pattern


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

