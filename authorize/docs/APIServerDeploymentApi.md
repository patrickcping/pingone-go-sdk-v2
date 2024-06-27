# \APIServerDeploymentApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeployAPIServer**](APIServerDeploymentApi.md#DeployAPIServer) | **Post** /environments/{environmentID}/apiServers/{apiServerID}/deployment | Deploy API Server
[**ReadDeploymentStatus**](APIServerDeploymentApi.md#ReadDeploymentStatus) | **Get** /environments/{environmentID}/apiServers/{apiServerID}/deployment | READ API Server Deployment Status



## DeployAPIServer

> APIServerDeployment DeployAPIServer(ctx, environmentID, apiServerID).ContentType(contentType).Execute()

Deploy API Server

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
    contentType := "contentType_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerDeploymentApi.DeployAPIServer(context.Background(), environmentID, apiServerID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerDeploymentApi.DeployAPIServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeployAPIServer`: APIServerDeployment
    fmt.Fprintf(os.Stdout, "Response from `APIServerDeploymentApi.DeployAPIServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeployAPIServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

### Return type

[**APIServerDeployment**](APIServerDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDeploymentStatus

> APIServerDeployment ReadDeploymentStatus(ctx, environmentID, apiServerID).Execute()

READ API Server Deployment Status

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APIServerDeploymentApi.ReadDeploymentStatus(context.Background(), environmentID, apiServerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIServerDeploymentApi.ReadDeploymentStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDeploymentStatus`: APIServerDeployment
    fmt.Fprintf(os.Stdout, "Response from `APIServerDeploymentApi.ReadDeploymentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDeploymentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**APIServerDeployment**](APIServerDeployment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

