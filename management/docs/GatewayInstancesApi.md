# \GatewayInstancesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllGatewayInstances**](GatewayInstancesApi.md#ReadAllGatewayInstances) | **Get** /environments/{environmentID}/gateways/{gatewayID}/instances | READ All Gateway Instances
[**ReadOneGatewayInstance**](GatewayInstancesApi.md#ReadOneGatewayInstance) | **Get** /environments/{environmentID}/gateways/{gatewayID}/instances/{instanceID} | READ One Gateway Instance



## ReadAllGatewayInstances

READ All Gateway Instances

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllGatewayInstances(ctx, environmentID, gatewayID).Execute()

#### Example

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllGatewayInstances(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllGatewayInstances``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllGatewayInstances`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllGatewayInstances`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllGatewayInstances(ctx, environmentID, gatewayID).ExecuteInitialPage()

#### Example

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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayInstancesApi.ReadAllGatewayInstances(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayInstancesApi.ReadAllGatewayInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGatewayInstances`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `GatewayInstancesApi.ReadAllGatewayInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGatewayInstancesRequest struct via the builder pattern


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


## ReadOneGatewayInstance

> GatewayInstance ReadOneGatewayInstance(ctx, environmentID, gatewayID, instanceID).Execute()

READ One Gateway Instance

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
    gatewayID := "gatewayID_example" // string | 
    instanceID := "instanceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayInstancesApi.ReadOneGatewayInstance(context.Background(), environmentID, gatewayID, instanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayInstancesApi.ReadOneGatewayInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayInstance`: GatewayInstance
    fmt.Fprintf(os.Stdout, "Response from `GatewayInstancesApi.ReadOneGatewayInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**instanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GatewayInstance**](GatewayInstance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

