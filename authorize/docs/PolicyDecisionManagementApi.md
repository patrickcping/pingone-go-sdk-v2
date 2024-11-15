# \PolicyDecisionManagementApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDecisionEndpoint**](PolicyDecisionManagementApi.md#CreateDecisionEndpoint) | **Post** /environments/{environmentID}/decisionEndpoints | CREATE Decision Endpoint
[**DeleteDecisionEndpoint**](PolicyDecisionManagementApi.md#DeleteDecisionEndpoint) | **Delete** /environments/{environmentID}/decisionEndpoints/{decisionEndpointID} | DELETE Decision Endpoint
[**ReadAllDecisionEndpoints**](PolicyDecisionManagementApi.md#ReadAllDecisionEndpoints) | **Get** /environments/{environmentID}/decisionEndpoints | READ All Decision Endpoints
[**ReadOneDecisionEndpoint**](PolicyDecisionManagementApi.md#ReadOneDecisionEndpoint) | **Get** /environments/{environmentID}/decisionEndpoints/{decisionEndpointID} | READ One Decision Endpoint
[**UpdateDecisionEndpoint**](PolicyDecisionManagementApi.md#UpdateDecisionEndpoint) | **Put** /environments/{environmentID}/decisionEndpoints/{decisionEndpointID} | UPDATE Decision Endpoint



## CreateDecisionEndpoint

> DecisionEndpoint CreateDecisionEndpoint(ctx, environmentID).DecisionEndpoint(decisionEndpoint).Execute()

CREATE Decision Endpoint

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
    decisionEndpoint := *openapiclient.NewDecisionEndpoint("Description_example", "Name_example", false) // DecisionEndpoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDecisionManagementApi.CreateDecisionEndpoint(context.Background(), environmentID).DecisionEndpoint(decisionEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDecisionManagementApi.CreateDecisionEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDecisionEndpoint`: DecisionEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PolicyDecisionManagementApi.CreateDecisionEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDecisionEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decisionEndpoint** | [**DecisionEndpoint**](DecisionEndpoint.md) |  | 

### Return type

[**DecisionEndpoint**](DecisionEndpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDecisionEndpoint

> DeleteDecisionEndpoint(ctx, environmentID, decisionEndpointID).Execute()

DELETE Decision Endpoint

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
    decisionEndpointID := "decisionEndpointID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyDecisionManagementApi.DeleteDecisionEndpoint(context.Background(), environmentID, decisionEndpointID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDecisionManagementApi.DeleteDecisionEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**decisionEndpointID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDecisionEndpointRequest struct via the builder pattern


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


## ReadAllDecisionEndpoints

READ All Decision Endpoints

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllDecisionEndpoints(ctx, environmentID).Execute()

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
    pagedIterator := api.ReadAllDecisionEndpoints(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllDecisionEndpoints``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllDecisionEndpoints`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllDecisionEndpoints`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllDecisionEndpoints(ctx, environmentID).ExecuteInitialPage()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDecisionManagementApi.ReadAllDecisionEndpoints(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDecisionManagementApi.ReadAllDecisionEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllDecisionEndpoints`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `PolicyDecisionManagementApi.ReadAllDecisionEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDecisionEndpointsRequest struct via the builder pattern


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


## ReadOneDecisionEndpoint

> DecisionEndpoint ReadOneDecisionEndpoint(ctx, environmentID, decisionEndpointID).Execute()

READ One Decision Endpoint

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
    decisionEndpointID := "decisionEndpointID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDecisionManagementApi.ReadOneDecisionEndpoint(context.Background(), environmentID, decisionEndpointID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDecisionManagementApi.ReadOneDecisionEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDecisionEndpoint`: DecisionEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PolicyDecisionManagementApi.ReadOneDecisionEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**decisionEndpointID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDecisionEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DecisionEndpoint**](DecisionEndpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDecisionEndpoint

> DecisionEndpoint UpdateDecisionEndpoint(ctx, environmentID, decisionEndpointID).DecisionEndpoint(decisionEndpoint).Execute()

UPDATE Decision Endpoint

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
    decisionEndpointID := "decisionEndpointID_example" // string | 
    decisionEndpoint := *openapiclient.NewDecisionEndpoint("Description_example", "Name_example", false) // DecisionEndpoint |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyDecisionManagementApi.UpdateDecisionEndpoint(context.Background(), environmentID, decisionEndpointID).DecisionEndpoint(decisionEndpoint).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyDecisionManagementApi.UpdateDecisionEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDecisionEndpoint`: DecisionEndpoint
    fmt.Fprintf(os.Stdout, "Response from `PolicyDecisionManagementApi.UpdateDecisionEndpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**decisionEndpointID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDecisionEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **decisionEndpoint** | [**DecisionEndpoint**](DecisionEndpoint.md) |  | 

### Return type

[**DecisionEndpoint**](DecisionEndpoint.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

