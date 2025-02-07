# \PhoneDeliverySettingsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneDeliverySettings**](PhoneDeliverySettingsApi.md#CreatePhoneDeliverySettings) | **Post** /environments/{environmentID}/notificationsSettings/phoneDeliverySettings | CREATE Phone Delivery Settings
[**DeletePhoneDeliverySettings**](PhoneDeliverySettingsApi.md#DeletePhoneDeliverySettings) | **Delete** /environments/{environmentID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsID} | DELETE Phone Delivery Settings
[**ReadAllPhoneDeliverySettings**](PhoneDeliverySettingsApi.md#ReadAllPhoneDeliverySettings) | **Get** /environments/{environmentID}/notificationsSettings/phoneDeliverySettings | READ All Phone Delivery Settings
[**ReadOnePhoneDeliverySettings**](PhoneDeliverySettingsApi.md#ReadOnePhoneDeliverySettings) | **Get** /environments/{environmentID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsID} | READ One Phone Delivery Settings
[**UpdatePhoneDeliverySettings**](PhoneDeliverySettingsApi.md#UpdatePhoneDeliverySettings) | **Put** /environments/{environmentID}/notificationsSettings/phoneDeliverySettings/{phoneDeliverySettingsID} | UPDATE Phone Delivery Settings



## CreatePhoneDeliverySettings

> NotificationsSettingsPhoneDeliverySettings CreatePhoneDeliverySettings(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).NotificationsSettingsPhoneDeliverySettings(notificationsSettingsPhoneDeliverySettings).Execute()

CREATE Phone Delivery Settings

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
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    notificationsSettingsPhoneDeliverySettings := openapiclient.NotificationsSettingsPhoneDeliverySettings{NotificationsSettingsPhoneDeliverySettingsCustom: openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustom(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsProvider("CUSTOM_TWILIO"), "Name_example", []openapiclient.NotificationsSettingsPhoneDeliverySettingsCustomRequest{*openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustomRequest(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod("SMS"), "Url_example", openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod("GET"), openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat("FULL"))}, *openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod("BASIC")))} // NotificationsSettingsPhoneDeliverySettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhoneDeliverySettingsApi.CreatePhoneDeliverySettings(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).NotificationsSettingsPhoneDeliverySettings(notificationsSettingsPhoneDeliverySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhoneDeliverySettingsApi.CreatePhoneDeliverySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhoneDeliverySettings`: NotificationsSettingsPhoneDeliverySettings
    fmt.Fprintf(os.Stdout, "Response from `PhoneDeliverySettingsApi.CreatePhoneDeliverySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhoneDeliverySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **notificationsSettingsPhoneDeliverySettings** | [**NotificationsSettingsPhoneDeliverySettings**](NotificationsSettingsPhoneDeliverySettings.md) |  | 

### Return type

[**NotificationsSettingsPhoneDeliverySettings**](NotificationsSettingsPhoneDeliverySettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneDeliverySettings

> DeletePhoneDeliverySettings(ctx, environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

DELETE Phone Delivery Settings

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
    phoneDeliverySettingsID := "phoneDeliverySettingsID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PhoneDeliverySettingsApi.DeletePhoneDeliverySettings(context.Background(), environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhoneDeliverySettingsApi.DeletePhoneDeliverySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**phoneDeliverySettingsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhoneDeliverySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

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


## ReadAllPhoneDeliverySettings

> EntityArrayPagedIterator ReadAllPhoneDeliverySettings(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

READ All Phone Delivery Settings

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
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhoneDeliverySettingsApi.ReadAllPhoneDeliverySettings(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhoneDeliverySettingsApi.ReadAllPhoneDeliverySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllPhoneDeliverySettings`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `PhoneDeliverySettingsApi.ReadAllPhoneDeliverySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllPhoneDeliverySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOnePhoneDeliverySettings

> NotificationsSettingsPhoneDeliverySettings ReadOnePhoneDeliverySettings(ctx, environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

READ One Phone Delivery Settings

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
    phoneDeliverySettingsID := "phoneDeliverySettingsID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhoneDeliverySettingsApi.ReadOnePhoneDeliverySettings(context.Background(), environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhoneDeliverySettingsApi.ReadOnePhoneDeliverySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePhoneDeliverySettings`: NotificationsSettingsPhoneDeliverySettings
    fmt.Fprintf(os.Stdout, "Response from `PhoneDeliverySettingsApi.ReadOnePhoneDeliverySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**phoneDeliverySettingsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePhoneDeliverySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**NotificationsSettingsPhoneDeliverySettings**](NotificationsSettingsPhoneDeliverySettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePhoneDeliverySettings

> NotificationsSettingsPhoneDeliverySettings UpdatePhoneDeliverySettings(ctx, environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).NotificationsSettingsPhoneDeliverySettings(notificationsSettingsPhoneDeliverySettings).Execute()

UPDATE Phone Delivery Settings

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
    phoneDeliverySettingsID := "phoneDeliverySettingsID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    notificationsSettingsPhoneDeliverySettings := openapiclient.NotificationsSettingsPhoneDeliverySettings{NotificationsSettingsPhoneDeliverySettingsCustom: openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustom(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsProvider("CUSTOM_TWILIO"), "Name_example", []openapiclient.NotificationsSettingsPhoneDeliverySettingsCustomRequest{*openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustomRequest(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomDeliveryMethod("SMS"), "Url_example", openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomRequestMethod("GET"), openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomNumberFormat("FULL"))}, *openapiclient.NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication(openapiclient.EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod("BASIC")))} // NotificationsSettingsPhoneDeliverySettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PhoneDeliverySettingsApi.UpdatePhoneDeliverySettings(context.Background(), environmentID, phoneDeliverySettingsID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).NotificationsSettingsPhoneDeliverySettings(notificationsSettingsPhoneDeliverySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhoneDeliverySettingsApi.UpdatePhoneDeliverySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePhoneDeliverySettings`: NotificationsSettingsPhoneDeliverySettings
    fmt.Fprintf(os.Stdout, "Response from `PhoneDeliverySettingsApi.UpdatePhoneDeliverySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**phoneDeliverySettingsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePhoneDeliverySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **notificationsSettingsPhoneDeliverySettings** | [**NotificationsSettingsPhoneDeliverySettings**](NotificationsSettingsPhoneDeliverySettings.md) |  | 

### Return type

[**NotificationsSettingsPhoneDeliverySettings**](NotificationsSettingsPhoneDeliverySettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

