# \GatewayManagementGatewayInstancesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllGatewayInstances**](GatewayManagementGatewayInstancesApi.md#ReadAllGatewayInstances) | **Get** /v1/environments/{environmentID}/gateways/{gatewayID}/instances | READ All Gateway Instances
[**ReadOneGatewayInstance**](GatewayManagementGatewayInstancesApi.md#ReadOneGatewayInstance) | **Get** /v1/environments/{environmentID}/gateways/{gatewayID}/instances/{instanceID} | READ One Gateway Instance



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
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayManagementGatewayInstancesApi.ReadAllGatewayInstances(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayManagementGatewayInstancesApi.ReadAllGatewayInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGatewayInstances`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `GatewayManagementGatewayInstancesApi.ReadAllGatewayInstances`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    gatewayID := "gatewayID_example" // string | 
    instanceID := "instanceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayManagementGatewayInstancesApi.ReadOneGatewayInstance(context.Background(), environmentID, gatewayID, instanceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayManagementGatewayInstancesApi.ReadOneGatewayInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayInstance`: GatewayInstance
    fmt.Fprintf(os.Stdout, "Response from `GatewayManagementGatewayInstancesApi.ReadOneGatewayInstance`: %v\n", resp)
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

