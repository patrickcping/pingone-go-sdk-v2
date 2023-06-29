# \TrustedEmailAddressesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustedEmailAddress**](TrustedEmailAddressesApi.md#CreateTrustedEmailAddress) | **Post** /environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails | CREATE Trusted Email Address
[**DeleteTrustedEmailAddress**](TrustedEmailAddressesApi.md#DeleteTrustedEmailAddress) | **Delete** /environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | DELETE Trusted Email Address
[**ReadAllTrustedEmailAddresses**](TrustedEmailAddressesApi.md#ReadAllTrustedEmailAddresses) | **Get** /environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails | READ All Trusted Email Addresses
[**ReadOneTrustedEmailAddress**](TrustedEmailAddressesApi.md#ReadOneTrustedEmailAddress) | **Get** /environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | READ One Trusted Email Address
[**ResendVerificationCodeToEmail**](TrustedEmailAddressesApi.md#ResendVerificationCodeToEmail) | **Post** /environments/{environmentID}/emailDomains/{emailDomainID}/trustedEmails/{trustedEmailId} | Resend Verification Code To Email



## CreateTrustedEmailAddress

> EmailDomainTrustedEmail CreateTrustedEmailAddress(ctx, environmentID, emailDomainID).EmailDomainTrustedEmail(emailDomainTrustedEmail).Execute()

CREATE Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    emailDomainTrustedEmail := *openapiclient.NewEmailDomainTrustedEmail("EmailAddress_example") // EmailDomainTrustedEmail |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailAddressesApi.CreateTrustedEmailAddress(context.Background(), environmentID, emailDomainID).EmailDomainTrustedEmail(emailDomainTrustedEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailAddressesApi.CreateTrustedEmailAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrustedEmailAddress`: EmailDomainTrustedEmail
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailAddressesApi.CreateTrustedEmailAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrustedEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emailDomainTrustedEmail** | [**EmailDomainTrustedEmail**](EmailDomainTrustedEmail.md) |  | 

### Return type

[**EmailDomainTrustedEmail**](EmailDomainTrustedEmail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustedEmailAddress

> DeleteTrustedEmailAddress(ctx, environmentID, emailDomainID, trustedEmailId).Execute()

DELETE Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TrustedEmailAddressesApi.DeleteTrustedEmailAddress(context.Background(), environmentID, emailDomainID, trustedEmailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailAddressesApi.DeleteTrustedEmailAddress``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTrustedEmailAddressRequest struct via the builder pattern


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


## ReadAllTrustedEmailAddresses

> EntityArray ReadAllTrustedEmailAddresses(ctx, environmentID, emailDomainID).Execute()

READ All Trusted Email Addresses

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
    emailDomainID := "emailDomainID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllTrustedEmailAddresses`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailAddressesApi.ReadAllTrustedEmailAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllTrustedEmailAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ReadOneTrustedEmailAddress

> EmailDomainTrustedEmail ReadOneTrustedEmailAddress(ctx, environmentID, emailDomainID, trustedEmailId).Execute()

READ One Trusted Email Address

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailAddressesApi.ReadOneTrustedEmailAddress(context.Background(), environmentID, emailDomainID, trustedEmailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailAddressesApi.ReadOneTrustedEmailAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneTrustedEmailAddress`: EmailDomainTrustedEmail
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailAddressesApi.ReadOneTrustedEmailAddress`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadOneTrustedEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EmailDomainTrustedEmail**](EmailDomainTrustedEmail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendVerificationCodeToEmail

> EmailDomainTrustedEmail ResendVerificationCodeToEmail(ctx, environmentID, emailDomainID, trustedEmailId).ContentType(contentType).Execute()

Resend Verification Code To Email

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
    emailDomainID := "emailDomainID_example" // string | 
    trustedEmailId := "trustedEmailId_example" // string | 
    contentType := "application/vnd.pingidentity.trustedEmail.sendVerificationCode+json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailAddressesApi.ResendVerificationCodeToEmail(context.Background(), environmentID, emailDomainID, trustedEmailId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailAddressesApi.ResendVerificationCodeToEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResendVerificationCodeToEmail`: EmailDomainTrustedEmail
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailAddressesApi.ResendVerificationCodeToEmail`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiResendVerificationCodeToEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 

### Return type

[**EmailDomainTrustedEmail**](EmailDomainTrustedEmail.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

