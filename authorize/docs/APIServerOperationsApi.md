# \APIServerOperationsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIServerOperation**](APIServerOperationsApi.md#CreateAPIServerOperation) | **Post** /environments/{environmentID}/apiServers/{apiServerID}/operations | CREATE API Server Operation
[**DeleteAPIServerOperation**](APIServerOperationsApi.md#DeleteAPIServerOperation) | **Delete** /environments/{environmentID}/apiServers/{apiServerID}/operations/{apiServerOperationID} | DELETE API Server Operation
[**ReadAllAPIServerOperations**](APIServerOperationsApi.md#ReadAllAPIServerOperations) | **Get** /environments/{environmentID}/apiServers/{apiServerID}/operations | READ All API Server Operations
[**ReadOneAPIServerOperation**](APIServerOperationsApi.md#ReadOneAPIServerOperation) | **Get** /environments/{environmentID}/apiServers/{apiServerID}/operations/{apiServerOperationID} | READ One API Server Operation
[**UpdateAPIServerOperation**](APIServerOperationsApi.md#UpdateAPIServerOperation) | **Put** /environments/{environmentID}/apiServers/{apiServerID}/operations/{apiServerOperationID} | UPDATE API Server Operation



## CreateAPIServerOperation

> APIServerOperation CreateAPIServerOperation(ctx, environmentID, apiServerID).APIServerOperation(aPIServerOperation).Execute()

CREATE API Server Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
    apiServerID := "apiServerID_example" // string | 
    aPIServerOperation := *openapiclient.NewAPIServerOperation("Name_example", []openapiclient.APIServerOperationPathsInner{*openapiclient.NewAPIServerOperationPathsInner("Pattern_example", openapiclient.EnumAPIServerOperationPathPatternType("EXACT"))}) // APIServerOperation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerOperationsApi.CreateAPIServerOperation(context.Background(), environmentID, apiServerID).APIServerOperation(aPIServerOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerOperationsApi.CreateAPIServerOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAPIServerOperation`: APIServerOperation
    fmt.Fprintf(os.Stdout, "Response from `APIServerOperationsApi.CreateAPIServerOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIServerOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPIServerOperation** | [**APIServerOperation**](APIServerOperation.md) |  | 

### Return type

[**APIServerOperation**](APIServerOperation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIServerOperation

> DeleteAPIServerOperation(ctx, environmentID, apiServerID, apiServerOperationID).Execute()

DELETE API Server Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
    apiServerID := "apiServerID_example" // string | 
    apiServerOperationID := "apiServerOperationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.APIServerOperationsApi.DeleteAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerOperationsApi.DeleteAPIServerOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 
**apiServerOperationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIServerOperationRequest struct via the builder pattern


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


## ReadAllAPIServerOperations

READ All API Server Operations

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllAPIServerOperations(ctx, environmentID, apiServerID).Execute()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllAPIServerOperations(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllAPIServerOperations``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllAPIServerOperations`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllAPIServerOperations`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllAPIServerOperations(ctx, environmentID, apiServerID).ExecuteInitialPage()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
    apiServerID := "apiServerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerOperationsApi.ReadAllAPIServerOperations(context.Background(), environmentID, apiServerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerOperationsApi.ReadAllAPIServerOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAPIServerOperations`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `APIServerOperationsApi.ReadAllAPIServerOperations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllAPIServerOperationsRequest struct via the builder pattern


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


## ReadOneAPIServerOperation

> APIServerOperation ReadOneAPIServerOperation(ctx, environmentID, apiServerID, apiServerOperationID).Execute()

READ One API Server Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
    apiServerID := "apiServerID_example" // string | 
    apiServerOperationID := "apiServerOperationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerOperationsApi.ReadOneAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerOperationsApi.ReadOneAPIServerOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAPIServerOperation`: APIServerOperation
    fmt.Fprintf(os.Stdout, "Response from `APIServerOperationsApi.ReadOneAPIServerOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 
**apiServerOperationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAPIServerOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**APIServerOperation**](APIServerOperation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIServerOperation

> APIServerOperation UpdateAPIServerOperation(ctx, environmentID, apiServerID, apiServerOperationID).APIServerOperation(aPIServerOperation).Execute()

UPDATE API Server Operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
    apiServerID := "apiServerID_example" // string | 
    apiServerOperationID := "apiServerOperationID_example" // string | 
    aPIServerOperation := *openapiclient.NewAPIServerOperation("Name_example", []openapiclient.APIServerOperationPathsInner{*openapiclient.NewAPIServerOperationPathsInner("Pattern_example", openapiclient.EnumAPIServerOperationPathPatternType("EXACT"))}) // APIServerOperation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerOperationsApi.UpdateAPIServerOperation(context.Background(), environmentID, apiServerID, apiServerOperationID).APIServerOperation(aPIServerOperation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerOperationsApi.UpdateAPIServerOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAPIServerOperation`: APIServerOperation
    fmt.Fprintf(os.Stdout, "Response from `APIServerOperationsApi.UpdateAPIServerOperation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 
**apiServerOperationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIServerOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **aPIServerOperation** | [**APIServerOperation**](APIServerOperation.md) |  | 

### Return type

[**APIServerOperation**](APIServerOperation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

