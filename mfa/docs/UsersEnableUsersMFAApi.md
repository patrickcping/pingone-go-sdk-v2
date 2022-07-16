# \UsersEnableUsersMFAApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet**](UsersEnableUsersMFAApi.md#V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet) | **Get** /v1/environments/{envID}/users/{userID}/mfaEnabled | READ User MFA Enabled
[**V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut**](UsersEnableUsersMFAApi.md#V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut) | **Put** /v1/environments/{envID}/users/{userID}/mfaEnabled | UPDATE User MFA Enabled



## V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet

> V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet(ctx, envID, userID).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersEnableUsersMFAApi.V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet(context.Background(), envID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersMFAApi.V1EnvironmentsEnvIDUsersUserIDMfaEnabledGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMfaEnabledGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut

> V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut(ctx, envID, userID).Body(body).Execute()

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
    envID := "envID_example" // string | 
    userID := "userID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersEnableUsersMFAApi.V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut(context.Background(), envID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersEnableUsersMFAApi.V1EnvironmentsEnvIDUsersUserIDMfaEnabledPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDMfaEnabledPutRequest struct via the builder pattern


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

