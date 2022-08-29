# \UserAccountsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDPost**](UserAccountsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPost) | **Post** /v1/environments/{environmentID}/users/{userID} | User Account Unlock



## V1EnvironmentsEnvironmentIDUsersUserIDPost

> EntityArray V1EnvironmentsEnvironmentIDUsersUserIDPost(ctx, environmentID, userID).ContentType(contentType).Execute()

User Account Unlock

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
    contentType := "application/vnd.pingidentity.account.unlock+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDPost(context.Background(), environmentID, userID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `V1EnvironmentsEnvironmentIDUsersUserIDPost`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `UserAccountsApi.V1EnvironmentsEnvironmentIDUsersUserIDPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

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

