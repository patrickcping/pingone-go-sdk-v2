# \DeviceAuthenticationPolicyApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut**](DeviceAuthenticationPolicyApi.md#V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut) | **Put** /v1/environments/{envID}/deviceAuthenticationPolicy/{deviceAuthPolicyID} | UPDATE Device Authentication Policy
[**V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet**](DeviceAuthenticationPolicyApi.md#V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet) | **Get** /v1/environments/{envID}/deviceAuthenticationPolicy | READ Device Authentication Policy



## V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut

> V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut(ctx, envID, deviceAuthPolicyID).Body(body).Execute()

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
    envID := "envID_example" // string | 
    deviceAuthPolicyID := "deviceAuthPolicyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut(context.Background(), envID, deviceAuthPolicyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.V1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**deviceAuthPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet

> V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet(ctx, envID).Execute()

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
    envID := "envID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet(context.Background(), envID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.V1EnvironmentsEnvIDDeviceAuthenticationPolicyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDDeviceAuthenticationPolicyGetRequest struct via the builder pattern


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

