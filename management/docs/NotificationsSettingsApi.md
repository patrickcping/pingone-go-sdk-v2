# \NotificationsSettingsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNotificationsSettings**](NotificationsSettingsApi.md#DeleteNotificationsSettings) | **Delete** /environments/{environmentID}/notificationsSettings | DELETE Notifications Settings
[**ReadNotificationsSettings**](NotificationsSettingsApi.md#ReadNotificationsSettings) | **Get** /environments/{environmentID}/notificationsSettings | READ Notifications Settings
[**UpdateNotificationsSettings**](NotificationsSettingsApi.md#UpdateNotificationsSettings) | **Put** /environments/{environmentID}/notificationsSettings | UPDATE Notifications Settings



## DeleteNotificationsSettings

> NotificationsSettings DeleteNotificationsSettings(ctx, environmentID).Execute()

DELETE Notifications Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsSettingsApi.DeleteNotificationsSettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsApi.DeleteNotificationsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNotificationsSettings`: NotificationsSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsApi.DeleteNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsSettings**](NotificationsSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNotificationsSettings

> NotificationsSettings ReadNotificationsSettings(ctx, environmentID).Execute()

READ Notifications Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsSettingsApi.ReadNotificationsSettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsApi.ReadNotificationsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNotificationsSettings`: NotificationsSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsApi.ReadNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsSettings**](NotificationsSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsSettings

> NotificationsSettings UpdateNotificationsSettings(ctx, environmentID).NotificationsSettings(notificationsSettings).Execute()

UPDATE Notifications Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    notificationsSettings := *openapiclient.NewNotificationsSettings() // NotificationsSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsSettingsApi.UpdateNotificationsSettings(context.Background(), environmentID).NotificationsSettings(notificationsSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsApi.UpdateNotificationsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsSettings`: NotificationsSettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsApi.UpdateNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsSettings** | [**NotificationsSettings**](NotificationsSettings.md) |  | 

### Return type

[**NotificationsSettings**](NotificationsSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

