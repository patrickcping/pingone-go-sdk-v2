# \ResourceScopesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceScope**](ResourceScopesApi.md#CreateResourceScope) | **Post** /environments/{environmentID}/resources/{resourceID}/scopes | CREATE PingOne access control scope
[**DeleteResourceScope**](ResourceScopesApi.md#DeleteResourceScope) | **Delete** /environments/{environmentID}/resources/{resourceID}/scopes/{scopeID} | DELETE Scope
[**ReadAllResourceScopes**](ResourceScopesApi.md#ReadAllResourceScopes) | **Get** /environments/{environmentID}/resources/{resourceID}/scopes | READ All Scopes (Resource)
[**ReadOneResourceScope**](ResourceScopesApi.md#ReadOneResourceScope) | **Get** /environments/{environmentID}/resources/{resourceID}/scopes/{scopeID} | READ One Scope
[**UpdateResourceScope**](ResourceScopesApi.md#UpdateResourceScope) | **Put** /environments/{environmentID}/resources/{resourceID}/scopes/{scopeID} | UPDATE PingOne access control scope



## CreateResourceScope

> ResourceScope CreateResourceScope(ctx, environmentID, resourceID).ResourceScope(resourceScope).Execute()

CREATE PingOne access control scope

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
    resourceID := "resourceID_example" // string | 
    resourceScope := *openapiclient.NewResourceScope("Name_example") // ResourceScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceScopesApi.CreateResourceScope(context.Background(), environmentID, resourceID).ResourceScope(resourceScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceScopesApi.CreateResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ResourceScopesApi.CreateResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceScope** | [**ResourceScope**](ResourceScope.md) |  | 

### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceScope

> DeleteResourceScope(ctx, environmentID, resourceID, scopeID).Execute()

DELETE Scope

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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceScopesApi.DeleteResourceScope(context.Background(), environmentID, resourceID, scopeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceScopesApi.DeleteResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceScopeRequest struct via the builder pattern


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


## ReadAllResourceScopes

> EntityArrayPagedIterator ReadAllResourceScopes(ctx, environmentID, resourceID).Execute()

READ All Scopes (Resource)

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
    resourceID := "resourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceScopesApi.ReadAllResourceScopes(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceScopesApi.ReadAllResourceScopes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllResourceScopes`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `ResourceScopesApi.ReadAllResourceScopes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllResourceScopesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ReadOneResourceScope

> ResourceScope ReadOneResourceScope(ctx, environmentID, resourceID, scopeID).Execute()

READ One Scope

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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceScopesApi.ReadOneResourceScope(context.Background(), environmentID, resourceID, scopeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceScopesApi.ReadOneResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ResourceScopesApi.ReadOneResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceScope

> ResourceScope UpdateResourceScope(ctx, environmentID, resourceID, scopeID).ResourceScope(resourceScope).Execute()

UPDATE PingOne access control scope

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
    resourceID := "resourceID_example" // string | 
    scopeID := "scopeID_example" // string | 
    resourceScope := *openapiclient.NewResourceScope("Name_example") // ResourceScope |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceScopesApi.UpdateResourceScope(context.Background(), environmentID, resourceID, scopeID).ResourceScope(resourceScope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceScopesApi.UpdateResourceScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceScope`: ResourceScope
    fmt.Fprintf(os.Stdout, "Response from `ResourceScopesApi.UpdateResourceScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**scopeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceScope** | [**ResourceScope**](ResourceScope.md) |  | 

### Return type

[**ResourceScope**](ResourceScope.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

