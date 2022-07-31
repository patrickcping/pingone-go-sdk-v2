# \UsersEnableUsersMFAApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet**](UsersEnableUsersMFAApi.md#V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet) | **Get** /v1/environments/{environmentID}/users/{userID}/mfaEnabled | READ User MFA Enabled
[**V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut**](UsersEnableUsersMFAApi.md#V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut) | **Put** /v1/environments/{environmentID}/users/{userID}/mfaEnabled | UPDATE User MFA Enabled



## V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet

> V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet(ctx, environmentID, userID).Execute()

READ User MFA Enabled

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
    resp, r, err := apiClient.UsersEnableUsersMFAApi.V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersMFAApi.V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut

> V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut(ctx, environmentID, userID).Body(body).Execute()

UPDATE User MFA Enabled

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersEnableUsersMFAApi.V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut(context.Background(), environmentID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersMFAApi.V1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDMfaEnabledPutRequest struct via the builder pattern


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

