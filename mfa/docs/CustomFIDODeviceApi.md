# \CustomFIDODeviceApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFidoDevice**](CustomFIDODeviceApi.md#CreateFidoDevice) | **Post** /v1/environments/{environmentID}/fidoDevicesMetadata | CREATE Custom FIDO Device
[**DeleteFidoDevice**](CustomFIDODeviceApi.md#DeleteFidoDevice) | **Delete** /v1/environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | DELETE Custom FIDO Device
[**ReadFidoDevices**](CustomFIDODeviceApi.md#ReadFidoDevices) | **Get** /v1/environments/{environmentID}/fidoDevicesMetadata | READ All Custom FIDO Devices
[**ReadOneFidoDevice**](CustomFIDODeviceApi.md#ReadOneFidoDevice) | **Get** /v1/environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | READ One Custom FIDO Device
[**UpdateFIDODevice**](CustomFIDODeviceApi.md#UpdateFIDODevice) | **Put** /v1/environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | UPDATE Custom FIDO Device



## CreateFidoDevice

> map[string]interface{} CreateFidoDevice(ctx, environmentID).Body(body).Execute()

CREATE Custom FIDO Device

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFIDODeviceApi.CreateFidoDevice(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFIDODeviceApi.CreateFidoDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFidoDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomFIDODeviceApi.CreateFidoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFidoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFidoDevice

> DeleteFidoDevice(ctx, environmentID, fidoDeviceID).Execute()

DELETE Custom FIDO Device

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
    fidoDeviceID := "fidoDeviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFIDODeviceApi.DeleteFidoDevice(context.Background(), environmentID, fidoDeviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFIDODeviceApi.DeleteFidoDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoDeviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFidoDeviceRequest struct via the builder pattern


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


## ReadFidoDevices

> EntityArray ReadFidoDevices(ctx, environmentID).Execute()

READ All Custom FIDO Devices

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
    resp, r, err := apiClient.CustomFIDODeviceApi.ReadFidoDevices(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFIDODeviceApi.ReadFidoDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFidoDevices`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CustomFIDODeviceApi.ReadFidoDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFidoDevicesRequest struct via the builder pattern


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


## ReadOneFidoDevice

> map[string]interface{} ReadOneFidoDevice(ctx, environmentID, fidoDeviceID).Execute()

READ One Custom FIDO Device

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
    fidoDeviceID := "fidoDeviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFIDODeviceApi.ReadOneFidoDevice(context.Background(), environmentID, fidoDeviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFIDODeviceApi.ReadOneFidoDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneFidoDevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomFIDODeviceApi.ReadOneFidoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoDeviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFidoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFIDODevice

> map[string]interface{} UpdateFIDODevice(ctx, environmentID, fidoDeviceID).Body(body).Execute()

UPDATE Custom FIDO Device

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
    fidoDeviceID := "fidoDeviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomFIDODeviceApi.UpdateFIDODevice(context.Background(), environmentID, fidoDeviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomFIDODeviceApi.UpdateFIDODevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFIDODevice`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomFIDODeviceApi.UpdateFIDODevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoDeviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFIDODeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

