# \UserMFADevicesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDelete**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDelete) | **Delete** /environments/{environmentID}/users/{userID}/devices | DELETE Device Order
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete) | **Delete** /environments/{environmentID}/users/{userID}/devices/{deviceID} | DELETE MFA User Device
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet) | **Get** /environments/{environmentID}/users/{userID}/devices/{deviceID} | READ One MFA User Device
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut) | **Put** /environments/{environmentID}/users/{userID}/devices/{deviceID}/logs | SEND MFA Device Logs
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut) | **Put** /environments/{environmentID}/users/{userID}/devices/{deviceID}/nickname | UPDATE Device Nickname
[**EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost) | **Post** /environments/{environmentID}/users/{userID}/devices/{deviceID} | ACTIVATE MFA User Device
[**EnvironmentsEnvironmentIDUsersUserIDDevicesGet**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesGet) | **Get** /environments/{environmentID}/users/{userID}/devices | READ All MFA User Devices
[**EnvironmentsEnvironmentIDUsersUserIDDevicesPost**](UserMFADevicesApi.md#EnvironmentsEnvironmentIDUsersUserIDDevicesPost) | **Post** /environments/{environmentID}/users/{userID}/devices | SET Device Order



## EnvironmentsEnvironmentIDUsersUserIDDevicesDelete

> EnvironmentsEnvironmentIDUsersUserIDDevicesDelete(ctx, environmentID, userID).ContentType(contentType).Execute()

DELETE Device Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.order.remove+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDelete(context.Background(), environmentID, userID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete

> EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete(ctx, environmentID, userID, deviceID).Execute()

DELETE MFA User Device

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete(context.Background(), environmentID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet

> EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet(ctx, environmentID, userID, deviceID).Execute()

READ One MFA User Device

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet(context.Background(), environmentID, userID, deviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut

> EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut(ctx, environmentID, userID, deviceID).Body(body).Execute()

SEND MFA Device Logs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut(context.Background(), environmentID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDLogsPutRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut

> EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut(ctx, environmentID, userID, deviceID).Body(body).Execute()

UPDATE Device Nickname

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut(context.Background(), environmentID, userID, deviceID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDNicknamePutRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost

> EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost(ctx, environmentID, userID, deviceID).ContentType(contentType).Body(body).Execute()

ACTIVATE MFA User Device

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    deviceID := "deviceID_example" // string | 
    contentType := "application/vnd.pingidentity.device.activate+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost(context.Background(), environmentID, userID, deviceID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesDeviceIDPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesGet

> EnvironmentsEnvironmentIDUsersUserIDDevicesGet(ctx, environmentID, userID).Execute()

READ All MFA User Devices

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDDevicesPost

> EnvironmentsEnvironmentIDUsersUserIDDevicesPost(ctx, environmentID, userID).ContentType(contentType).Body(body).Execute()

SET Device Order

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    contentType := "application/vnd.pingidentity.devices.reorder+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesPost(context.Background(), environmentID, userID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFADevicesApi.EnvironmentsEnvironmentIDUsersUserIDDevicesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDDevicesPostRequest struct via the builder pattern


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

