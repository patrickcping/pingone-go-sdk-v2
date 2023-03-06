# \UserMFAPairingKeysApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete**](UserMFAPairingKeysApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete) | **Delete** /v1/environments/{environmentID}/users/{userID}/pairingKeys/{pairingKeyID} | DELETE MFA Pairing Key
[**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet**](UserMFAPairingKeysApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/pairingKeys/{pairingKeyID} | READ One MFA Pairing Key
[**V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost**](UserMFAPairingKeysApi.md#V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost) | **Post** /v1/environments/{environmentID}/users/{userID}/pairingKeys | CREATE MFA Pairing Key



## V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete

> V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete(ctx, environmentID, userID, pairingKeyID).Execute()

DELETE MFA Pairing Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    pairingKeyID := "pairingKeyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete(context.Background(), environmentID, userID, pairingKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDelete``: %v\n", err)
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
**pairingKeyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet(ctx, environmentID, userID, pairingKeyID).Execute()

READ One MFA Pairing Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    pairingKeyID := "pairingKeyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet(context.Background(), environmentID, userID, pairingKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGet``: %v\n", err)
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
**pairingKeyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPairingKeyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost

> V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost(ctx, environmentID, userID).Body(body).Execute()

CREATE MFA Pairing Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost(context.Background(), environmentID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMFAPairingKeysApi.V1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDPairingKeysPostRequest struct via the builder pattern


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

