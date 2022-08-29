# \ResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResource**](ResourcesApi.md#CreateResource) | **Post** /v1/environments/{environmentID}/resources | CREATE Resource
[**DeleteResource**](ResourcesApi.md#DeleteResource) | **Delete** /v1/environments/{environmentID}/resources/{resourceID} | DELETE Resource
[**ReadAllResources**](ResourcesApi.md#ReadAllResources) | **Get** /v1/environments/{environmentID}/resources | READ All Resources
[**ReadOneResource**](ResourcesApi.md#ReadOneResource) | **Get** /v1/environments/{environmentID}/resources/{resourceID} | READ One Resource
[**UpdateResource**](ResourcesApi.md#UpdateResource) | **Put** /v1/environments/{environmentID}/resources/{resourceID} | UPDATE Resource



## CreateResource

> Resource CreateResource(ctx, environmentID).Resource(resource).Execute()

CREATE Resource

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
    resource := *openapiclient.NewResource("Name_example") // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.CreateResource(context.Background(), environmentID).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.CreateResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.CreateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResource

> DeleteResource(ctx, environmentID, resourceID).Execute()

DELETE Resource

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
    resourceID := "resourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.DeleteResource(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.DeleteResource``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequest struct via the builder pattern


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


## ReadAllResources

> EntityArray ReadAllResources(ctx, environmentID).Execute()

READ All Resources

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
    resp, r, err := apiClient.ResourcesApi.ReadAllResources(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ReadAllResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllResources`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ReadAllResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllResourcesRequest struct via the builder pattern


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


## ReadOneResource

> Resource ReadOneResource(ctx, environmentID, resourceID).Execute()

READ One Resource

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
    resourceID := "resourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.ReadOneResource(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.ReadOneResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.ReadOneResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Resource**](Resource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResource

> Resource UpdateResource(ctx, environmentID, resourceID).Resource(resource).Execute()

UPDATE Resource

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
    resourceID := "resourceID_example" // string | 
    resource := *openapiclient.NewResource("Name_example") // Resource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourcesApi.UpdateResource(context.Background(), environmentID, resourceID).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourcesApi.UpdateResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResource`: Resource
    fmt.Fprintf(os.Stdout, "Response from `ResourcesApi.UpdateResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

