# \NotificationsNotificationsSettingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDNotificationsSettingsDelete**](NotificationsNotificationsSettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsDelete) | **Delete** /v1/environments/{environmentID}/notificationsSettings | DELETE Notifications Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsGet**](NotificationsNotificationsSettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsGet) | **Get** /v1/environments/{environmentID}/notificationsSettings | READ Notifications Settings
[**V1EnvironmentsEnvironmentIDNotificationsSettingsPut**](NotificationsNotificationsSettingsApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsPut) | **Put** /v1/environments/{environmentID}/notificationsSettings | UPDATE Notifications Settings



## V1EnvironmentsEnvironmentIDNotificationsSettingsDelete

> V1EnvironmentsEnvironmentIDNotificationsSettingsDelete(ctx, environmentID).Execute()

DELETE Notifications Settings

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
    resp, r, err := apiClient.NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsDelete(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsGet

> V1EnvironmentsEnvironmentIDNotificationsSettingsGet(ctx, environmentID).Execute()

READ Notifications Settings

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
    resp, r, err := apiClient.NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsPut

> V1EnvironmentsEnvironmentIDNotificationsSettingsPut(ctx, environmentID).Execute()

UPDATE Notifications Settings

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
    resp, r, err := apiClient.NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPut(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsNotificationsSettingsApi.V1EnvironmentsEnvironmentIDNotificationsSettingsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsPutRequest struct via the builder pattern


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

