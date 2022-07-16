# \GatewayManagementGatewayCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayCredential**](GatewayManagementGatewayCredentialsApi.md#CreateGatewayCredential) | **Post** /v1/environments/{environmentID}/gateways/{gatewayID}/credentials | CREATE Gateway Credentials
[**DeleteGatewayCredential**](GatewayManagementGatewayCredentialsApi.md#DeleteGatewayCredential) | **Delete** /v1/environments/{environmentID}/gateways/{gatewayID}/credentials/{credentialID} | DELETE Gateway Credentials



## CreateGatewayCredential

> GatewayCredential CreateGatewayCredential(ctx, environmentID, gatewayID).Execute()

CREATE Gateway Credentials

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
    resp, r, err := apiClient.GatewayManagementGatewayCredentialsApi.CreateGatewayCredential(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayManagementGatewayCredentialsApi.CreateGatewayCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayCredential`: GatewayCredential
    fmt.Fprintf(os.Stdout, "Response from `GatewayManagementGatewayCredentialsApi.CreateGatewayCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayCredential**](GatewayCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGatewayCredential

> DeleteGatewayCredential(ctx, environmentID, gatewayID, credentialID).Execute()

DELETE Gateway Credentials

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
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayManagementGatewayCredentialsApi.DeleteGatewayCredential(context.Background(), environmentID, gatewayID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayManagementGatewayCredentialsApi.DeleteGatewayCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayCredentialRequest struct via the builder pattern


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

