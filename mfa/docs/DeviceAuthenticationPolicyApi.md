# \DeviceAuthenticationPolicyApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceAuthenticationPolicies**](DeviceAuthenticationPolicyApi.md#CreateDeviceAuthenticationPolicies) | **Post** /v1/environments/{environmentID}/deviceAuthenticationPolicies | CREATE Device Authentication Policy
[**DeleteDeviceAuthenticationPolicy**](DeviceAuthenticationPolicyApi.md#DeleteDeviceAuthenticationPolicy) | **Delete** /v1/environments/{environmentID}/deviceAuthenticationPolicy/{deviceAuthPolicyID} | DELETE Device Authentication Policy
[**ReadDeviceAuthenticationPolicies**](DeviceAuthenticationPolicyApi.md#ReadDeviceAuthenticationPolicies) | **Get** /v1/environments/{environmentID}/deviceAuthenticationPolicies | READ Device Authentication Policies
[**ReadOneDeviceAuthenticationPolicy**](DeviceAuthenticationPolicyApi.md#ReadOneDeviceAuthenticationPolicy) | **Get** /v1/environments/{environmentID}/deviceAuthenticationPolicy/{deviceAuthPolicyID} | READ One Device Authentication Policy
[**V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut**](DeviceAuthenticationPolicyApi.md#V1EnvironmentsEnvironmentIDDeviceAuthenticationPolicyDeviceAuthPolicyIDPut) | **Put** /v1/environments/{environmentID}/deviceAuthenticationPolicy/{deviceAuthPolicyID} | UPDATE Device Authentication Policy



## CreateDeviceAuthenticationPolicies

> DeviceAuthenticationPolicy CreateDeviceAuthenticationPolicies(ctx, environmentID).DeviceAuthenticationPolicy(deviceAuthenticationPolicy).Execute()

CREATE Device Authentication Policy

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
    deviceAuthenticationPolicy := *openapiclient.NewDeviceAuthenticationPolicy() // DeviceAuthenticationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies(context.Background(), environmentID).DeviceAuthenticationPolicy(deviceAuthenticationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceAuthenticationPolicies`: DeviceAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceAuthenticationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAuthenticationPolicy** | [**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md) |  | 

### Return type

[**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceAuthenticationPolicy

> DeleteDeviceAuthenticationPolicy(ctx, environmentID, deviceAuthPolicyID).Execute()

DELETE Device Authentication Policy

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.DeleteDeviceAuthenticationPolicy(context.Background(), environmentID, deviceAuthPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.DeleteDeviceAuthenticationPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDeviceAuthenticationPolicyRequest struct via the builder pattern


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


## ReadDeviceAuthenticationPolicies

> EntityArray ReadDeviceAuthenticationPolicies(ctx, environmentID).Execute()

READ Device Authentication Policies

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
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDeviceAuthenticationPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDeviceAuthenticationPoliciesRequest struct via the builder pattern


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


## ReadOneDeviceAuthenticationPolicy

> DeviceAuthenticationPolicy ReadOneDeviceAuthenticationPolicy(ctx, environmentID, deviceAuthPolicyID).Execute()

READ One Device Authentication Policy

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy(context.Background(), environmentID, deviceAuthPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDeviceAuthenticationPolicy`: DeviceAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**deviceAuthPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDeviceAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

