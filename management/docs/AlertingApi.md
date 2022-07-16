# \AlertingApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete**](AlertingApi.md#V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete) | **Delete** /v1/environments/{environmentID}/alertChannels/{alertChannelID} | DELETE Alert Channel
[**V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut**](AlertingApi.md#V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut) | **Put** /v1/environments/{environmentID}/alertChannels/{alertChannelID} | UPDATE Alert Channel
[**V1EnvironmentsEnvironmentIDAlertChannelsGet**](AlertingApi.md#V1EnvironmentsEnvironmentIDAlertChannelsGet) | **Get** /v1/environments/{environmentID}/alertChannels | READ All Alert Channels per Environment
[**V1EnvironmentsEnvironmentIDAlertChannelsPost**](AlertingApi.md#V1EnvironmentsEnvironmentIDAlertChannelsPost) | **Post** /v1/environments/{environmentID}/alertChannels | CREATE Alert Channel (Email)



## V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete

> V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete(ctx, environmentID, alertChannelID).Execute()

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
    resp, r, err := apiClient.AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete(context.Background(), environmentID, alertChannelID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut

> V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut(ctx, environmentID, alertChannelID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut(context.Background(), environmentID, alertChannelID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAlertChannelsAlertChannelIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAlertChannelsGet

> V1EnvironmentsEnvironmentIDAlertChannelsGet(ctx, environmentID).Execute()

READ All Alert Channels per Environment

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
    resp, r, err := apiClient.AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAlertChannelsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAlertChannelsPost

> V1EnvironmentsEnvironmentIDAlertChannelsPost(ctx, environmentID).Body(body).Execute()

CREATE Alert Channel (Email)

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
    resp, r, err := apiClient.AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingApi.V1EnvironmentsEnvironmentIDAlertChannelsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAlertChannelsPostRequest struct via the builder pattern


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

