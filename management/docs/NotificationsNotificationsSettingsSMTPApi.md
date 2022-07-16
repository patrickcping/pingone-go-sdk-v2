# \NotificationsNotificationsSettingsSMTPApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet**](NotificationsNotificationsSettingsSMTPApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet) | **Get** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | READ Notifications Settings (SMTP)
[**V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut**](NotificationsNotificationsSettingsSMTPApi.md#V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut) | **Put** /v1/environments/{environmentID}/notificationsSettings/emailDeliverySettings | UPDATE Notifications Settings (SMTP)



## V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet

> V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet(ctx, environmentID).Execute()

READ Notifications Settings (SMTP)

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
    resp, r, err := apiClient.NotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut

> V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut(ctx, environmentID).Body(body).Execute()

UPDATE Notifications Settings (SMTP)

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
    resp, r, err := apiClient.NotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsNotificationsSettingsSMTPApi.V1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDNotificationsSettingsEmailDeliverySettingsPutRequest struct via the builder pattern


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

