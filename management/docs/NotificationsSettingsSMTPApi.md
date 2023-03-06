# \NotificationsSettingsSMTPApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmailDeliverySettings**](NotificationsSettingsSMTPApi.md#DeleteEmailDeliverySettings) | **Delete** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | DELETE Email Delivery Settings
[**ReadEmailNotificationsSettings**](NotificationsSettingsSMTPApi.md#ReadEmailNotificationsSettings) | **Get** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | READ Email Notifications Settings
[**UpdateEmailNotificationsSettings**](NotificationsSettingsSMTPApi.md#UpdateEmailNotificationsSettings) | **Put** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | UPDATE Email Notifications Settings



## DeleteEmailDeliverySettings

> DeleteEmailDeliverySettings(ctx, environmentID).Execute()

DELETE Email Delivery Settings

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
    r, err := apiClient.NotificationsSettingsSMTPApi.DeleteEmailDeliverySettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsSMTPApi.DeleteEmailDeliverySettings``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEmailDeliverySettingsRequest struct via the builder pattern


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


## ReadEmailNotificationsSettings

> NotificationsSettingsEmailDeliverySettings ReadEmailNotificationsSettings(ctx, environmentID).Execute()

READ Email Notifications Settings

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
    resp, r, err := apiClient.NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadEmailNotificationsSettings`: NotificationsSettingsEmailDeliverySettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsSMTPApi.ReadEmailNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadEmailNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationsSettingsEmailDeliverySettings**](NotificationsSettingsEmailDeliverySettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEmailNotificationsSettings

> NotificationsSettingsEmailDeliverySettings UpdateEmailNotificationsSettings(ctx, environmentID).NotificationsSettingsEmailDeliverySettings(notificationsSettingsEmailDeliverySettings).Execute()

UPDATE Email Notifications Settings

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
    notificationsSettingsEmailDeliverySettings := *openapiclient.NewNotificationsSettingsEmailDeliverySettings() // NotificationsSettingsEmailDeliverySettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsSettingsSMTPApi.UpdateEmailNotificationsSettings(context.Background(), environmentID).NotificationsSettingsEmailDeliverySettings(notificationsSettingsEmailDeliverySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsSMTPApi.UpdateEmailNotificationsSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEmailNotificationsSettings`: NotificationsSettingsEmailDeliverySettings
    fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsSMTPApi.UpdateEmailNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEmailNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsSettingsEmailDeliverySettings** | [**NotificationsSettingsEmailDeliverySettings**](NotificationsSettingsEmailDeliverySettings.md) |  | 

### Return type

[**NotificationsSettingsEmailDeliverySettings**](NotificationsSettingsEmailDeliverySettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

