# \NotificationsPhoneDeliverySettingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete**](NotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete) | **Delete** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | DELETE Phone Delivery Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet**](NotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet) | **Get** /v1/environments/{environmentID}/notificationsSettings/phoneDeliverySettings | READ All Phone Delivery Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet**](NotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet) | **Get** /v1/environments/{environmentID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId} | READ One Phone Delivery Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut**](NotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut) | **Put** /v1/environments/{environmentID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsId} | UPDATE Phone Delivery Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost**](NotificationsPhoneDeliverySettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost) | **Post** /v1/environments/{environmentID}/notificationsSettings/phoneDeliverySettings | CREATE Phone Delivery Settings (Syniverse)



## V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete

> V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete(ctx, environmentID).Execute()

DELETE Phone Delivery Settings

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
    resp, r, err := apiClient.NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet

> V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet(ctx, environmentID).Execute()

READ All Phone Delivery Settings

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
    resp, r, err := apiClient.NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet

> V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet(ctx, environmentID, phoneDeliverySettingsId).Execute()

READ One Phone Delivery Settings

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
    phoneDeliverySettingsId := "phoneDeliverySettingsId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet(context.Background(), environmentID, phoneDeliverySettingsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**phoneDeliverySettingsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut

> V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut(ctx, environmentID, phoneDeliverySettingsId).Body(body).Execute()

UPDATE Phone Delivery Settings

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
    phoneDeliverySettingsId := "phoneDeliverySettingsId_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut(context.Background(), environmentID, phoneDeliverySettingsId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**phoneDeliverySettingsId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPhoneDeliverySettingsIdPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost

> V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost(ctx, environmentID).Body(body).Execute()

CREATE Phone Delivery Settings (Syniverse)

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
    resp, r, err := apiClient.NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPhoneDeliverySettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsPhoneDeliverySettingsPostRequest struct via the builder pattern


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

