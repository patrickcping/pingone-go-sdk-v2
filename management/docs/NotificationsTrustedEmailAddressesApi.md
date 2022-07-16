# \NotificationsTrustedEmailAddressesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet**](NotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails | READ All Trusted Email Addresses
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost**](NotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost) | **Post** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails | CREATE Trusted Email Address
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete**](NotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete) | **Delete** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | DELETE Trusted Email Address
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet**](NotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | READ One Trusted Email Address
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost**](NotificationsTrustedEmailAddressesApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost) | **Post** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | Resend Verification Code To Email



## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet(ctx, environmentID, emailDomainID).Execute()

READ All Trusted Email Addresses

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
    emailDomainID := "emailDomainID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost(ctx, environmentID, emailDomainID).Body(body).Execute()

CREATE Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost(context.Background(), environmentID, emailDomainID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete(ctx, environmentID, emailDomainID, trustedEmailId).Execute()

DELETE Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete(context.Background(), environmentID, emailDomainID, trustedEmailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet(ctx, environmentID, emailDomainID, trustedEmailId).Execute()

READ One Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet(context.Background(), environmentID, emailDomainID, trustedEmailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost(ctx, environmentID, emailDomainID, trustedEmailId).ContentType(contentType).Execute()

Resend Verification Code To Email

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 
    contentType := "application/vnd.pingidentity.trustedEmail.sendVerificationCode+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost(context.Background(), environmentID, emailDomainID, trustedEmailId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailAddressesApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 
**trustedEmailId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDTrustedEmailsTrustedEmailIdPostRequest struct via the builder pattern


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

