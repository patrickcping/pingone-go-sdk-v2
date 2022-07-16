# \DeviceAuthenticationPolicyApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut**](DeviceAuthenticationPolicyApi.md#V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut) | **Put** /v1/environments/{environmentID}/deviceAuthenticationPolicy/{deviceAuthPolicyID} | UPDATE Device Authentication Policy
[**V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet**](DeviceAuthenticationPolicyApi.md#V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet) | **Get** /v1/environments/{environmentID}/deviceAuthenticationPolicy | READ Device Authentication Policy



## V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut

> V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut(ctx, environmentID, deviceAuthPolicyID).Body(body).Execute()

UPDATE Device Authentication Policy

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
    deviceAuthPolicyID := "deviceAuthPolicyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut(context.Background(), environmentID, deviceAuthPolicyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**deviceAuthPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet

> V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet(ctx, environmentID).Execute()

READ Device Authentication Policy

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
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyGetRequest struct via the builder pattern


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

