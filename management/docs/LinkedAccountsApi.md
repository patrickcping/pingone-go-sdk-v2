# \LinkedAccountsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet**](LinkedAccountsApi.md#EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet) | **Get** /environments/{environmentID}/users/{userID}/linkedAccounts | READ Linked Accounts
[**EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete**](LinkedAccountsApi.md#EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete) | **Delete** /environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID} | DELETE Linked Account
[**EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet**](LinkedAccountsApi.md#EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet) | **Get** /environments/{environmentID}/users/{userID}/linkedAccounts/{linkedAccountID} | READ One Linked Account



## EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet

> EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet(ctx, environmentID, userID).Execute()

READ Linked Accounts

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete

> EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(ctx, environmentID, userID, linkedAccountID).Execute()

DELETE Linked Account

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
    userID := "userID_example" // string | 
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete(context.Background(), environmentID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDelete``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet

> EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet(ctx, environmentID, userID, linkedAccountID).Execute()

READ One Linked Account

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
    userID := "userID_example" // string | 
    linkedAccountID := "linkedAccountID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet(context.Background(), environmentID, userID, linkedAccountID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkedAccountsApi.EnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGet``: %v\n", err)
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
**linkedAccountID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDLinkedAccountsLinkedAccountIDGetRequest struct via the builder pattern


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

