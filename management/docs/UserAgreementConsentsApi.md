# \UserAgreementConsentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet**](UserAgreementConsentsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet) | **Get** /v1/environments/{environmentID}/users/{userID}/agreementConsents/{agreementID} | READ One User Agreement Consent
[**V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost**](UserAgreementConsentsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost) | **Post** /v1/environments/{environmentID}/users/{userID}/agreementConsents/{agreementID} | Revoke Agreement
[**V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet**](UserAgreementConsentsApi.md#V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet) | **Get** /v1/environments/{environmentID}/users/{userID}/agreementConsents | READ All User Agreement Consents



## V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet

> V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet(ctx, environmentID, userID, agreementID).Execute()

READ One User Agreement Consent

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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet(context.Background(), environmentID, userID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet``: %v\n", err)
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
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost

> V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost(ctx, environmentID, userID, agreementID).ContentType(contentType).Execute()

Revoke Agreement

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
    agreementID := "agreementID_example" // string | 
    contentType := "application/vnd.pingidentity.consent.revoke+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost(context.Background(), environmentID, userID, agreementID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost``: %v\n", err)
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
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet

> V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet(ctx, environmentID, userID).Execute()

READ All User Agreement Consents

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
    resp, r, err := apiClient.UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.V1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest struct via the builder pattern


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

