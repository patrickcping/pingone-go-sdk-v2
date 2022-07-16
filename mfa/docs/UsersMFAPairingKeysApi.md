# \UsersMFAPairingKeysApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete**](UsersMFAPairingKeysApi.md#V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete) | **Delete** /v1/environments/{envID}/users/{userID}/pairingKeys/{pairingKeyID} | DELETE MFA Pairing Key
[**V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet**](UsersMFAPairingKeysApi.md#V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet) | **Get** /v1/environments/{envID}/users/{userID}/pairingKeys/{pairingKeyID} | READ One MFA Pairing Key
[**V1EnvironmentsEnvIDUsersUserIDPairingKeysPost**](UsersMFAPairingKeysApi.md#V1EnvironmentsEnvIDUsersUserIDPairingKeysPost) | **Post** /v1/environments/{envID}/users/{userID}/pairingKeys | CREATE MFA Pairing Key



## V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete

> V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete(ctx, envID, userID, pairingKeyID).Execute()

DELETE MFA Pairing Key



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
    pairingKeyID := "pairingKeyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete(context.Background(), envID, userID, pairingKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDelete``: %v\n", err)
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
**pairingKeyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet

> V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet(ctx, envID, userID, pairingKeyID).Execute()

READ One MFA Pairing Key



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
    pairingKeyID := "pairingKeyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet(context.Background(), envID, userID, pairingKeyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGet``: %v\n", err)
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
**pairingKeyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDPairingKeysPairingKeyIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvIDUsersUserIDPairingKeysPost

> V1EnvironmentsEnvIDUsersUserIDPairingKeysPost(ctx, envID, userID).Body(body).Execute()

CREATE MFA Pairing Key



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
    resp, r, err := apiClient.UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPost(context.Background(), envID, userID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersMFAPairingKeysApi.V1EnvironmentsEnvIDUsersUserIDPairingKeysPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvIDUsersUserIDPairingKeysPostRequest struct via the builder pattern


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

