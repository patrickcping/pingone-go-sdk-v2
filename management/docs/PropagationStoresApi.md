# \PropagationStoresApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePropagationStore**](PropagationStoresApi.md#CreatePropagationStore) | **Post** /environments/{environmentID}/propagation/stores | CREATE Propagation Store
[**DeletePropagationStore**](PropagationStoresApi.md#DeletePropagationStore) | **Delete** /environments/{environmentID}/propagation/stores/{storeID} | DELETE Store
[**ReadAllStores**](PropagationStoresApi.md#ReadAllStores) | **Get** /environments/{environmentID}/propagation/stores | READ All Stores
[**ReadOnePropagationStore**](PropagationStoresApi.md#ReadOnePropagationStore) | **Get** /environments/{environmentID}/propagation/stores/{storeID} | READ One Store
[**TestConnectionConfiguration**](PropagationStoresApi.md#TestConnectionConfiguration) | **Post** /environments/{environmentID}/propagation/stores/connection/status | TEST Connection Configuration
[**UpdatePropagationStore**](PropagationStoresApi.md#UpdatePropagationStore) | **Put** /environments/{environmentID}/propagation/stores/{storeID} | UPDATE Store



## CreatePropagationStore

> PropagationStore CreatePropagationStore(ctx, environmentID).PropagationStore(propagationStore).Execute()

CREATE Propagation Store

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
    propagationStore := *openapiclient.NewPropagationStore(map[string]interface{}{"key": interface{}(123)}, "Name_example", openapiclient.EnumPropagationStoreType("Aquera")) // PropagationStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.CreatePropagationStore(context.Background(), environmentID).PropagationStore(propagationStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.CreatePropagationStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePropagationStore`: PropagationStore
    fmt.Fprintf(os.Stdout, "Response from `PropagationStoresApi.CreatePropagationStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePropagationStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propagationStore** | [**PropagationStore**](PropagationStore.md) |  | 

### Return type

[**PropagationStore**](PropagationStore.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePropagationStore

> DeletePropagationStore(ctx, environmentID, storeID).Execute()

DELETE Store

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
    storeID := "storeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoresApi.DeletePropagationStore(context.Background(), environmentID, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.DeletePropagationStore``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePropagationStoreRequest struct via the builder pattern


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


## ReadAllStores

> EntityArray ReadAllStores(ctx, environmentID).Execute()

READ All Stores

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
    resp, r, err := apiClient.PropagationStoresApi.ReadAllStores(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.ReadAllStores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllStores`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PropagationStoresApi.ReadAllStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllStoresRequest struct via the builder pattern


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


## ReadOnePropagationStore

> PropagationStore ReadOnePropagationStore(ctx, environmentID, storeID).Execute()

READ One Store

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
    storeID := "storeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.ReadOnePropagationStore(context.Background(), environmentID, storeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.ReadOnePropagationStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePropagationStore`: PropagationStore
    fmt.Fprintf(os.Stdout, "Response from `PropagationStoresApi.ReadOnePropagationStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePropagationStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PropagationStore**](PropagationStore.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnectionConfiguration

> TestConnectionConfiguration(ctx, environmentID).ContentType(contentType).PropagationStore(propagationStore).Execute()

TEST Connection Configuration

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
    contentType := openapiclient.EnumPropagationStoreConnectionStatusContentType("application/vnd.pingidentity.connection.check+json") // EnumPropagationStoreConnectionStatusContentType |  (optional)
    propagationStore := *openapiclient.NewPropagationStore(map[string]interface{}{"key": interface{}(123)}, "Name_example", openapiclient.EnumPropagationStoreType("Aquera")) // PropagationStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoresApi.TestConnectionConfiguration(context.Background(), environmentID).ContentType(contentType).PropagationStore(propagationStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.TestConnectionConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiTestConnectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | [**EnumPropagationStoreConnectionStatusContentType**](EnumPropagationStoreConnectionStatusContentType.md) |  | 
 **propagationStore** | [**PropagationStore**](PropagationStore.md) |  | 

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


## UpdatePropagationStore

> PropagationStore UpdatePropagationStore(ctx, environmentID, storeID).PropagationStore(propagationStore).Execute()

UPDATE Store

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
    storeID := "storeID_example" // string | 
    propagationStore := *openapiclient.NewPropagationStore(map[string]interface{}{"key": interface{}(123)}, "Name_example", openapiclient.EnumPropagationStoreType("Aquera")) // PropagationStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationStoresApi.UpdatePropagationStore(context.Background(), environmentID, storeID).PropagationStore(propagationStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoresApi.UpdatePropagationStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePropagationStore`: PropagationStore
    fmt.Fprintf(os.Stdout, "Response from `PropagationStoresApi.UpdatePropagationStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePropagationStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propagationStore** | [**PropagationStore**](PropagationStore.md) |  | 

### Return type

[**PropagationStore**](PropagationStore.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

