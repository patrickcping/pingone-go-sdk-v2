# \UserAgreementConsentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet**](UserAgreementConsentsApi.md#EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet) | **Get** /environments/{environmentID}/users/{userID}/agreementConsents/{agreementID} | READ One User Agreement Consent
[**EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost**](UserAgreementConsentsApi.md#EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost) | **Post** /environments/{environmentID}/users/{userID}/agreementConsents/{agreementID} | Revoke Agreement
[**EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet**](UserAgreementConsentsApi.md#EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet) | **Get** /environments/{environmentID}/users/{userID}/agreementConsents | READ All User Agreement Consents



## EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet

> EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet(ctx, environmentID, userID, agreementID).Execute()

READ One User Agreement Consent

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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet(context.Background(), environmentID, userID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost

> EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost(ctx, environmentID, userID, agreementID).ContentType(contentType).Execute()

Revoke Agreement

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
    agreementID := "agreementID_example" // string | 
    contentType := "application/vnd.pingidentity.consent.revoke+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost(context.Background(), environmentID, userID, agreementID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsAgreementIDPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet

> EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet(ctx, environmentID, userID).Execute()

READ All User Agreement Consents

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
    r, err := apiClient.UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAgreementConsentsApi.EnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDUsersUserIDAgreementConsentsGetRequest struct via the builder pattern


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

