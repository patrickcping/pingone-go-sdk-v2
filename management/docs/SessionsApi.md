# \SessionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet**](SessionsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet) | **Get** /v1/environments/{environmentID}/users/{userID}/sessions | READ All Sessions
[**V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete**](SessionsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/sessions/{sessionID} | DELETE Session
[**V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet**](SessionsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/sessions/{sessionID} | READ One Session



## V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet

> V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet(ctx, environmentID, userID).Execute()

READ All Sessions

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDSessionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete

> V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete(ctx, environmentID, userID, sessionID).Execute()

DELETE Session

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
    userID := "userID_example" // string | 
    sessionID := "sessionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete(context.Background(), environmentID, userID, sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet(ctx, environmentID, userID, sessionID).Execute()

READ One Session

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
    userID := "userID_example" // string | 
    sessionID := "sessionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet(context.Background(), environmentID, userID, sessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionsApi.V1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**sessionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDSessionsSessionIDGetRequest struct via the builder pattern


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

