# \PropagationStoresApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost) | **Post** /v1/environments/{environmentID}/propagation/stores/connection/status | TEST Connection Configuration
[**V1EnvironmentsEnvironmentIDPropagationStoresGet**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresGet) | **Get** /v1/environments/{environmentID}/propagation/stores | READ All Stores
[**V1EnvironmentsEnvironmentIDPropagationStoresPost**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresPost) | **Post** /v1/environments/{environmentID}/propagation/stores | CREATE Store (Aquera)
[**V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete) | **Delete** /v1/environments/{environmentID}/propagation/stores/{storeID} | DELETE Store
[**V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet) | **Get** /v1/environments/{environmentID}/propagation/stores/{storeID} | READ One Store
[**V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut**](PropagationStoresApi.md#V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut) | **Put** /v1/environments/{environmentID}/propagation/stores/{storeID} | UPDATE Store



## V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost

> V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost(ctx, environmentID).ContentType(contentType).Body(body).Execute()

TEST Connection Configuration

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
    contentType := "application/vnd.pingidentity.connection.check+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost(context.Background(), environmentID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresConnectionStatusPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationStoresGet

> V1EnvironmentsEnvironmentIDPropagationStoresGet(ctx, environmentID).Accept(accept).Execute()

READ All Stores

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
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresGet(context.Background(), environmentID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDPropagationStoresPost

> V1EnvironmentsEnvironmentIDPropagationStoresPost(ctx, environmentID).Body(body).Execute()

CREATE Store (Aquera)

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete

> V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete(ctx, environmentID, storeID).Accept(accept).Execute()

DELETE Store

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
    storeID := "storeID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete(context.Background(), environmentID, storeID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresStoreIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet

> V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet(ctx, environmentID, storeID).Accept(accept).Execute()

READ One Store

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
    storeID := "storeID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet(context.Background(), environmentID, storeID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresStoreIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut

> V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut(ctx, environmentID, storeID).Body(body).Execute()

UPDATE Store

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
    storeID := "storeID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut(context.Background(), environmentID, storeID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.V1EnvironmentsEnvironmentIDPropagationStoresStoreIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoresStoreIDPutRequest struct via the builder pattern


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

