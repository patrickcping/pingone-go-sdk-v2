# \UserIDVerificationApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet**](UserIDVerificationApi.md#EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet) | **Get** /environments/{environmentID}/users/{userID}/verifyTransactions | READ All ID Verification Transaction Records for a User
[**EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost**](UserIDVerificationApi.md#EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost) | **Post** /environments/{environmentID}/users/{userID}/verifyTransactions | CREATE ID Verification Transaction Record for a User
[**EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete**](UserIDVerificationApi.md#EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete) | **Delete** /environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | DELETE ID Verification Transaction Record for a User
[**EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet**](UserIDVerificationApi.md#EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet) | **Get** /environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | READ ID Verification Transaction Record for a User
[**EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut**](UserIDVerificationApi.md#EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut) | **Put** /environments/{environmentID}/users/{userID}/verifyTransactions/{transactionID} | UPDATE ID Verification Transaction Record for a User



## EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet

> EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet(ctx, environmentID, userID).Execute()

READ All ID Verification Transaction Records for a User

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
    r, err := apiClient.UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost

> EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost(ctx, environmentID, userID).Execute()

CREATE ID Verification Transaction Record for a User

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
    r, err := apiClient.UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete

> EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete(ctx, environmentID, userID, transactionID).Execute()

DELETE ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete(context.Background(), environmentID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet

> EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet(ctx, environmentID, userID, transactionID).Execute()

READ ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet(context.Background(), environmentID, userID, transactionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut

> EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut(ctx, environmentID, userID, transactionID).Body(body).Execute()

UPDATE ID Verification Transaction Record for a User

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
    transactionID := "transactionID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut(context.Background(), environmentID, userID, transactionID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserIDVerificationApi.EnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDVerifyTransactionsTransactionIDPutRequest struct via the builder pattern


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

