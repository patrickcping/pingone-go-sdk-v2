# \IdentityProvidersApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProvider**](IdentityProvidersApi.md#CreateIdentityProvider) | **Post** /v1/environments/{environmentID}/identityProviders | CREATE Identity Provider
[**DeleteIdentityProvider**](IdentityProvidersApi.md#DeleteIdentityProvider) | **Delete** /v1/environments/{environmentID}/identityProviders/{providerID} | DELETE Identity Provider
[**ReadAllIdentityProviders**](IdentityProvidersApi.md#ReadAllIdentityProviders) | **Get** /v1/environments/{environmentID}/identityProviders | READ All Identity Providers
[**ReadOneIdentityProvider**](IdentityProvidersApi.md#ReadOneIdentityProvider) | **Get** /v1/environments/{environmentID}/identityProviders/{providerID} | READ One Identity Provider
[**UpdateIdentityProvider**](IdentityProvidersApi.md#UpdateIdentityProvider) | **Put** /v1/environments/{environmentID}/identityProviders/{providerID} | UPDATE Identity Provider



## CreateIdentityProvider

> IdentityProvider CreateIdentityProvider(ctx, environmentID).IdentityProvider(identityProvider).Execute()

CREATE Identity Provider

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
    identityProvider := openapiclient.IdentityProvider{IdentityProviderApple: openapiclient.NewIdentityProviderApple(false, "Name_example", openapiclient.EnumIdentityProviderExt("FACEBOOK"), "ClientId_example", "ClientSecretSigningKey_example", "KeyId_example", "TeamId_example")} // IdentityProvider |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProvidersApi.CreateIdentityProvider(context.Background(), environmentID).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.CreateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersApi.CreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) |  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, environmentID, providerID).Execute()

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
    resp, r, err := apiClient.IdentityProvidersApi.DeleteIdentityProvider(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.DeleteIdentityProvider``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


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


## ReadAllIdentityProviders

> EntityArray ReadAllIdentityProviders(ctx, environmentID).Execute()

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
    resp, r, err := apiClient.IdentityProvidersApi.ReadAllIdentityProviders(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.ReadAllIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllIdentityProviders`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersApi.ReadAllIdentityProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllIdentityProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneIdentityProvider

> IdentityProvider ReadOneIdentityProvider(ctx, environmentID, providerID).Execute()

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
    resp, r, err := apiClient.IdentityProvidersApi.ReadOneIdentityProvider(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.ReadOneIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersApi.ReadOneIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvider

> IdentityProvider UpdateIdentityProvider(ctx, environmentID, providerID).IdentityProvider(identityProvider).Execute()

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
    identityProvider := openapiclient.IdentityProvider{IdentityProviderApple: openapiclient.NewIdentityProviderApple(false, "Name_example", openapiclient.EnumIdentityProviderExt("FACEBOOK"), "ClientId_example", "ClientSecretSigningKey_example", "KeyId_example", "TeamId_example")} // IdentityProvider |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProvidersApi.UpdateIdentityProvider(context.Background(), environmentID, providerID).IdentityProvider(identityProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProvidersApi.UpdateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProvider`: IdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `IdentityProvidersApi.UpdateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProvider** | [**IdentityProvider**](IdentityProvider.md) |  | 

### Return type

[**IdentityProvider**](IdentityProvider.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

