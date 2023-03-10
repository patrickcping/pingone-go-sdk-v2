# \NotificationsPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationsPolicy**](NotificationsPoliciesApi.md#CreateNotificationsPolicy) | **Post** /v1/environments/{environmentID}/notificationsPolicies | CREATE Notifications Policy
[**DeleteNotificationsPolicy**](NotificationsPoliciesApi.md#DeleteNotificationsPolicy) | **Delete** /v1/environments/{environmentID}/notificationsPolicies/{notificationsPolicyID} | DELETE Notifications Policy
[**ReadAllNotificationsPolicies**](NotificationsPoliciesApi.md#ReadAllNotificationsPolicies) | **Get** /v1/environments/{environmentID}/notificationsPolicies | READ All Notifications Policies
[**ReadOneNotificationsPolicy**](NotificationsPoliciesApi.md#ReadOneNotificationsPolicy) | **Get** /v1/environments/{environmentID}/notificationsPolicies/{notificationsPolicyID} | READ One Notifications Policy
[**UpdateNotificationsPolicy**](NotificationsPoliciesApi.md#UpdateNotificationsPolicy) | **Put** /v1/environments/{environmentID}/notificationsPolicies/{notificationsPolicyID} | UPDATE Notifications Policy



## CreateNotificationsPolicy

> NotificationsPolicy CreateNotificationsPolicy(ctx, environmentID).NotificationsPolicy(notificationsPolicy).Execute()

CREATE Notifications Policy

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
    notificationsPolicy := *openapiclient.NewNotificationsPolicy("Name_example", []openapiclient.NotificationsPolicyQuotasInner{*openapiclient.NewNotificationsPolicyQuotasInner(openapiclient.EnumNotificationsPolicyQuotaItemType("USER"), []openapiclient.EnumNotificationsPolicyQuotaDeliveryMethods{openapiclient.EnumNotificationsPolicyQuotaDeliveryMethods("SMS")})}) // NotificationsPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsPoliciesApi.CreateNotificationsPolicy(context.Background(), environmentID).NotificationsPolicy(notificationsPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPoliciesApi.CreateNotificationsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNotificationsPolicy`: NotificationsPolicy
    fmt.Fprintf(os.Stdout, "Response from `NotificationsPoliciesApi.CreateNotificationsPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsPolicy** | [**NotificationsPolicy**](NotificationsPolicy.md) |  | 

### Return type

[**NotificationsPolicy**](NotificationsPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationsPolicy

> DeleteNotificationsPolicy(ctx, environmentID, notificationsPolicyID).Execute()

DELETE Notifications Policy

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
    notificationsPolicyID := "notificationsPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsPoliciesApi.DeleteNotificationsPolicy(context.Background(), environmentID, notificationsPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPoliciesApi.DeleteNotificationsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**notificationsPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationsPolicyRequest struct via the builder pattern


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


## ReadAllNotificationsPolicies

> EntityArray ReadAllNotificationsPolicies(ctx, environmentID).Execute()

READ All Notifications Policies

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
    resp, r, err := apiClient.NotificationsPoliciesApi.ReadAllNotificationsPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPoliciesApi.ReadAllNotificationsPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllNotificationsPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `NotificationsPoliciesApi.ReadAllNotificationsPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllNotificationsPoliciesRequest struct via the builder pattern


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


## ReadOneNotificationsPolicy

> NotificationsPolicy ReadOneNotificationsPolicy(ctx, environmentID, notificationsPolicyID).Execute()

READ One Notifications Policy

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
    notificationsPolicyID := "notificationsPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsPoliciesApi.ReadOneNotificationsPolicy(context.Background(), environmentID, notificationsPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPoliciesApi.ReadOneNotificationsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneNotificationsPolicy`: NotificationsPolicy
    fmt.Fprintf(os.Stdout, "Response from `NotificationsPoliciesApi.ReadOneNotificationsPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**notificationsPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneNotificationsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationsPolicy**](NotificationsPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsPolicy

> NotificationsPolicy UpdateNotificationsPolicy(ctx, environmentID, notificationsPolicyID).NotificationsPolicy(notificationsPolicy).Execute()

UPDATE Notifications Policy

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
    notificationsPolicyID := "notificationsPolicyID_example" // string | 
    notificationsPolicy := *openapiclient.NewNotificationsPolicy("Name_example", []openapiclient.NotificationsPolicyQuotasInner{*openapiclient.NewNotificationsPolicyQuotasInner(openapiclient.EnumNotificationsPolicyQuotaItemType("USER"), []openapiclient.EnumNotificationsPolicyQuotaDeliveryMethods{openapiclient.EnumNotificationsPolicyQuotaDeliveryMethods("SMS")})}) // NotificationsPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsPoliciesApi.UpdateNotificationsPolicy(context.Background(), environmentID, notificationsPolicyID).NotificationsPolicy(notificationsPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsPoliciesApi.UpdateNotificationsPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationsPolicy`: NotificationsPolicy
    fmt.Fprintf(os.Stdout, "Response from `NotificationsPoliciesApi.UpdateNotificationsPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**notificationsPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notificationsPolicy** | [**NotificationsPolicy**](NotificationsPolicy.md) |  | 

### Return type

[**NotificationsPolicy**](NotificationsPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

