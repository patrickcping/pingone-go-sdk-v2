# \AlertingApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertChannel**](AlertingApi.md#CreateAlertChannel) | **Post** /v1/environments/{environmentID}/alertChannels | CREATE Alert Channel
[**DeleteAlertChannel**](AlertingApi.md#DeleteAlertChannel) | **Delete** /v1/environments/{environmentID}/alertChannels/{alertChannelID} | DELETE Alert Channel
[**ReadAllAlertChannels**](AlertingApi.md#ReadAllAlertChannels) | **Get** /v1/environments/{environmentID}/alertChannels | READ All Alert Channels
[**ReadOneAlertChannel**](AlertingApi.md#ReadOneAlertChannel) | **Get** /v1/environments/{environmentID}/alertChannels/{alertChannelID} | READ One Alert Channel
[**UpdateAlertChannel**](AlertingApi.md#UpdateAlertChannel) | **Put** /v1/environments/{environmentID}/alertChannels/{alertChannelID} | UPDATE Alert Channel



## CreateAlertChannel

> AlertChannel CreateAlertChannel(ctx, environmentID).AlertChannel(alertChannel).Execute()

CREATE Alert Channel

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
    alertChannel := *openapiclient.NewAlertChannel(openapiclient.EnumAlertChannelType("EMAIL"), []string{"Addresses_example"}) // AlertChannel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingApi.CreateAlertChannel(context.Background(), environmentID).AlertChannel(alertChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.CreateAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertingApi.CreateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertChannel** | [**AlertChannel**](AlertChannel.md) |  | 

### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlertChannel

> DeleteAlertChannel(ctx, environmentID, alertChannelID).Execute()

DELETE Alert Channel

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
    alertChannelID := "alertChannelID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingApi.DeleteAlertChannel(context.Background(), environmentID, alertChannelID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.DeleteAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**alertChannelID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertChannelRequest struct via the builder pattern


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


## ReadAllAlertChannels

> EntityArray ReadAllAlertChannels(ctx, environmentID).Execute()

READ All Alert Channels

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
    resp, r, err := apiClient.AlertingApi.ReadAllAlertChannels(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.ReadAllAlertChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAlertChannels`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AlertingApi.ReadAllAlertChannels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllAlertChannelsRequest struct via the builder pattern


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


## ReadOneAlertChannel

> AlertChannel ReadOneAlertChannel(ctx, environmentID, alertChannelID).Execute()

READ One Alert Channel

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
    alertChannelID := "alertChannelID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingApi.ReadOneAlertChannel(context.Background(), environmentID, alertChannelID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.ReadOneAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertingApi.ReadOneAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**alertChannelID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAlertChannel

> AlertChannel UpdateAlertChannel(ctx, environmentID, alertChannelID).AlertChannel(alertChannel).Execute()

UPDATE Alert Channel

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
    alertChannelID := "alertChannelID_example" // string | 
    alertChannel := *openapiclient.NewAlertChannel(openapiclient.EnumAlertChannelType("EMAIL"), []string{"Addresses_example"}) // AlertChannel |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingApi.UpdateAlertChannel(context.Background(), environmentID, alertChannelID).AlertChannel(alertChannel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.UpdateAlertChannel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAlertChannel`: AlertChannel
    fmt.Fprintf(os.Stdout, "Response from `AlertingApi.UpdateAlertChannel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**alertChannelID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAlertChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alertChannel** | [**AlertChannel**](AlertChannel.md) |  | 

### Return type

[**AlertChannel**](AlertChannel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

