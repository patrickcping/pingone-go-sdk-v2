# \UsersEnableUsersApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadUserEnabled**](UsersEnableUsersApi.md#ReadUserEnabled) | **Get** /v1/environments/{environmentID}/users/{userID}/enabled | READ User Enabled
[**UpdateUserEnabled**](UsersEnableUsersApi.md#UpdateUserEnabled) | **Put** /v1/environments/{environmentID}/users/{userID}/enabled | UPDATE User Enabled



## ReadUserEnabled

> UserEnabled ReadUserEnabled(ctx, environmentID, userID).Execute()

READ User Enabled

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
    resp, r, err := apiClient.UsersEnableUsersApi.ReadUserEnabled(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersApi.ReadUserEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserEnabled`: UserEnabled
    fmt.Fprintf(os.Stdout, "Response from `UsersEnableUsersApi.ReadUserEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserEnabled**](UserEnabled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserEnabled

> UserEnabled UpdateUserEnabled(ctx, environmentID, userID).UserEnabled(userEnabled).Execute()

UPDATE User Enabled

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
    userEnabled := *openapiclient.NewUserEnabled() // UserEnabled |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersEnableUsersApi.UpdateUserEnabled(context.Background(), environmentID, userID).UserEnabled(userEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersApi.UpdateUserEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserEnabled`: UserEnabled
    fmt.Fprintf(os.Stdout, "Response from `UsersEnableUsersApi.UpdateUserEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userEnabled** | [**UserEnabled**](UserEnabled.md) |  | 

### Return type

[**UserEnabled**](UserEnabled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

