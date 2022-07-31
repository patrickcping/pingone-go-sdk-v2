# \UsersMFADevicesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/devices | DELETE Device Order
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | DELETE MFA User Device
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | READ One MFA User Device
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut) | **Put** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID}/logs | SEND MFA Device Logs
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut) | **Put** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID}/nickname | UPDATE Device Nickname
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost) | **Post** /v1/environments/{environmentID}/users/{userID}/devices/{deviceID} | ACTIVATE MFA User Device
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet) | **Get** /v1/environments/{environmentID}/users/{userID}/devices | READ All MFA User Devices
[**V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost**](UsersMFADevicesApi.md#V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost) | **Post** /v1/environments/{environmentID}/users/{userID}/devices | SET Device Order



## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete(ctx, environmentID, userID).ContentType(contentType).Execute()

DELETE Device Order

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
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.order.remove+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete(context.Background(), environmentID, userID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete(ctx, environmentID, userID, deviceID).Execute()

DELETE MFA User Device

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
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete(context.Background(), environmentID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet(ctx, environmentID, userID, deviceID).Execute()

READ One MFA User Device

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
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet(context.Background(), environmentID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut(ctx, environmentID, userID, deviceID).Body(body).Execute()

SEND MFA Device Logs

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
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut(context.Background(), environmentID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut(ctx, environmentID, userID, deviceID).Body(body).Execute()

UPDATE Device Nickname

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
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut(context.Background(), environmentID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost(ctx, environmentID, userID, deviceID).ContentType(contentType).Body(body).Execute()

ACTIVATE MFA User Device

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
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    contentType := "application/vnd.pingidentity.device.activate+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost(context.Background(), environmentID, userID, deviceID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet(ctx, environmentID, userID).Execute()

READ All MFA User Devices

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost

> V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost(ctx, environmentID, userID).ContentType(contentType).Body(body).Execute()

SET Device Order

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
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.reorder+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost(context.Background(), environmentID, userID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvironmentIDUsersUserIDDevicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDDevicesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
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

