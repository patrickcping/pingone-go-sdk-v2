# \SubscriptionsWebhooksApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDSubscriptionsGet**](SubscriptionsWebhooksApi.md#V1EnvironmentsEnvironmentIDSubscriptionsGet) | **Get** /v1/environments/{environmentID}/subscriptions | READ All Subscriptions
[**V1EnvironmentsEnvironmentIDSubscriptionsPost**](SubscriptionsWebhooksApi.md#V1EnvironmentsEnvironmentIDSubscriptionsPost) | **Post** /v1/environments/{environmentID}/subscriptions | CREATE Subscriptions
[**V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete**](SubscriptionsWebhooksApi.md#V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete) | **Delete** /v1/environments/{environmentID}/subscriptions/{subscriptionID} | DELETE Subscription
[**V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet**](SubscriptionsWebhooksApi.md#V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet) | **Get** /v1/environments/{environmentID}/subscriptions/{subscriptionID} | READ One Subscription
[**V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut**](SubscriptionsWebhooksApi.md#V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut) | **Put** /v1/environments/{environmentID}/subscriptions/{subscriptionID} | UPDATE Subscription



## V1EnvironmentsEnvironmentIDSubscriptionsGet

> V1EnvironmentsEnvironmentIDSubscriptionsGet(ctx, environmentID).Execute()

READ All Subscriptions

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
    resp, r, err := apiClient.SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSubscriptionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSubscriptionsPost

> V1EnvironmentsEnvironmentIDSubscriptionsPost(ctx, environmentID).Body(body).Execute()

CREATE Subscriptions

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
    resp, r, err := apiClient.SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSubscriptionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete

> V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete(ctx, environmentID, subscriptionID).Execute()

DELETE Subscription

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
    subscriptionID := "subscriptionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete(context.Background(), environmentID, subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet

> V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet(ctx, environmentID, subscriptionID).Execute()

READ One Subscription

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
    subscriptionID := "subscriptionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet(context.Background(), environmentID, subscriptionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut

> V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut(ctx, environmentID, subscriptionID).Body(body).Execute()

UPDATE Subscription

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
    subscriptionID := "subscriptionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut(context.Background(), environmentID, subscriptionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsWebhooksApi.V1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**subscriptionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSubscriptionsSubscriptionIDPutRequest struct via the builder pattern


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

