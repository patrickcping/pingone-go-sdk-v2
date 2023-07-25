# \GatewayInstancesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllGatewayInstances**](GatewayInstancesApi.md#ReadAllGatewayInstances) | **Get** /environments/{environmentID}/gateways/{gatewayID}/instances | READ All Gateway Instances
[**ReadOneGatewayInstance**](GatewayInstancesApi.md#ReadOneGatewayInstance) | **Get** /environments/{environmentID}/gateways/{gatewayID}/instances/{instanceID} | READ One Gateway Instance



## ReadAllGatewayInstances

> EntityArray ReadAllGatewayInstances(ctx, environmentID, gatewayID).Execute()

READ All Gateway Instances

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayInstancesApi.ReadAllGatewayInstances(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayInstancesApi.ReadAllGatewayInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGatewayInstances`: EntityArray
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

[**EntityArray**](EntityArray.md)

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

