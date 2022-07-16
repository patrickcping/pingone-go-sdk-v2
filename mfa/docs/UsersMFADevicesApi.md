# \UsersMFADevicesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDDevicesDelete**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDelete) | **Delete** /v1/environments/{envID}/users/{userID}/devices | DELETE Device Order
[**V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/devices/{deviceID} | DELETE MFA User Device
[**V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet) | **Get** /v1/environments/{envID}/users/{userID}/devices/{deviceID} | READ One MFA User Device
[**V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut) | **Put** /v1/environments/{envID}/users/{userID}/devices/{deviceID}/logs | SEND MFA Device Logs
[**V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut) | **Put** /v1/environments/{envID}/users/{userID}/devices/{deviceID}/nickname | UPDATE Device Nickname
[**V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost) | **Post** /v1/environments/{envID}/users/{userID}/devices/{deviceID} | ACTIVATE MFA User Device
[**V1EnvironmentsEnvIDUsersUserIDDevicesGet**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesGet) | **Get** /v1/environments/{envID}/users/{userID}/devices | READ All MFA User Devices
[**V1EnvironmentsEnvIDUsersUserIDDevicesPost**](UsersMFADevicesApi.md#V1EnvironmentsEnvIDUsersUserIDDevicesPost) | **Post** /v1/environments/{envID}/users/{userID}/devices | SET Device Order



## V1EnvironmentsEnvIDUsersUserIDDevicesDelete

> V1EnvironmentsEnvIDUsersUserIDDevicesDelete(ctx, envID, userID).ContentType(contentType).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.order.remove+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDelete(context.Background(), envID, userID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete

> V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete(ctx, envID, userID, deviceID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete(context.Background(), envID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet

> V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet(ctx, envID, userID, deviceID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet(context.Background(), envID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut

> V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut(ctx, envID, userID, deviceID).Body(body).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut(context.Background(), envID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDLogsPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut

> V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut(ctx, envID, userID, deviceID).Body(body).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut(context.Background(), envID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDNicknamePutRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost

> V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost(ctx, envID, userID, deviceID).ContentType(contentType).Body(body).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    contentType := "application/vnd.pingidentity.device.activate+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost(context.Background(), envID, userID, deviceID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**deviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesDeviceIDPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesGet

> V1EnvironmentsEnvIDUsersUserIDDevicesGet(ctx, envID, userID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDDevicesPost

> V1EnvironmentsEnvIDUsersUserIDDevicesPost(ctx, envID, userID).ContentType(contentType).Body(body).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.reorder+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesPost(context.Background(), envID, userID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFADevicesApi.V1EnvironmentsEnvIDUsersUserIDDevicesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDDevicesPostRequest struct via the builder pattern


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

