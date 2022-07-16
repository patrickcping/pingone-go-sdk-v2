# \UsersUserIDVerificationApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet**](UsersUserIDVerificationApi.md#V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet) | **Get** /v1/environments/{environmentID}/users/{userID}/verifyTransactions | READ All ID Verification Transaction Records for a User
[**V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost**](UsersUserIDVerificationApi.md#V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost) | **Post** /v1/environments/{environmentID}/users/{userID}/verifyTransactions | CREATE ID Verification Transaction Record for a User
[**V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete**](UsersUserIDVerificationApi.md#V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | DELETE ID Verification Transaction Record for a User
[**V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet**](UsersUserIDVerificationApi.md#V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | READ ID Verification Transaction Record for a User
[**V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut**](UsersUserIDVerificationApi.md#V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut) | **Put** /v1/environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | UPDATE ID Verification Transaction Record for a User



## V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet

> V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet(ctx, environmentID, userID).Execute()

READ All ID Verification Transaction Records for a User

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
    resp, r, err := apiClient.UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost

> V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost(ctx, environmentID, userID).Execute()

CREATE ID Verification Transaction Record for a User

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
    resp, r, err := apiClient.UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete

> V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete(ctx, environmentID, userID, transactionID).Execute()

DELETE ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete(context.Background(), environmentID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet(ctx, environmentID, userID, transactionID).Execute()

READ ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet(context.Background(), environmentID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut

> V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut(ctx, environmentID, userID, transactionID).Body(body).Execute()

UPDATE ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut(context.Background(), environmentID, userID, transactionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserIDVerificationApi.V1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut``: %v\n", err)
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
**transactionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPutRequest struct via the builder pattern


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

