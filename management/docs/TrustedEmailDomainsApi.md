# \TrustedEmailDomainsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrustedEmailDomain**](TrustedEmailDomainsApi.md#CreateTrustedEmailDomain) | **Post** /v1/environments/{environmentID}/emailDomains | CREATE Trusted Email Domain
[**DeleteTrustedEmailDomain**](TrustedEmailDomainsApi.md#DeleteTrustedEmailDomain) | **Delete** /v1/environments/{environmentID}/emailDomains/{emailDomainID} | DELETE Trusted Email Domain
[**ReadAllTrustedEmailDomains**](TrustedEmailDomainsApi.md#ReadAllTrustedEmailDomains) | **Get** /v1/environments/{environmentID}/emailDomains | READ All Trusted Email Domains
[**ReadOneTrustedEmailDomain**](TrustedEmailDomainsApi.md#ReadOneTrustedEmailDomain) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID} | READ One Trusted Email Domain
[**ReadTrustedEmailDomainDKIMStatus**](TrustedEmailDomainsApi.md#ReadTrustedEmailDomainDKIMStatus) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/dkim | READ Trusted Email Domain DKIM Status
[**ReadTrustedEmailDomainOwnershipStatus**](TrustedEmailDomainsApi.md#ReadTrustedEmailDomainOwnershipStatus) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/ownership | READ Trusted Email Domain Ownership Status
[**ReadTrustedEmailDomainSPFStatus**](TrustedEmailDomainsApi.md#ReadTrustedEmailDomainSPFStatus) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/spf | READ Trusted Email Domain SPF Status



## CreateTrustedEmailDomain

> EmailDomain CreateTrustedEmailDomain(ctx, environmentID).EmailDomain(emailDomain).Execute()

CREATE Trusted Email Domain

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
    emailDomain := *openapiclient.NewEmailDomain("DomainName_example") // EmailDomain |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailDomainsApi.CreateTrustedEmailDomain(context.Background(), environmentID).EmailDomain(emailDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.CreateTrustedEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrustedEmailDomain`: EmailDomain
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.CreateTrustedEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrustedEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **emailDomain** | [**EmailDomain**](EmailDomain.md) |  | 

### Return type

[**EmailDomain**](EmailDomain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTrustedEmailDomain

> DeleteTrustedEmailDomain(ctx, environmentID, emailDomainID).Execute()

DELETE Trusted Email Domain

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
    r, err := apiClient.TrustedEmailDomainsApi.DeleteTrustedEmailDomain(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.DeleteTrustedEmailDomain``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteTrustedEmailDomainRequest struct via the builder pattern


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


## ReadAllTrustedEmailDomains

> EntityArray ReadAllTrustedEmailDomains(ctx, environmentID).Execute()

READ All Trusted Email Domains

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrustedEmailDomainsApi.ReadAllTrustedEmailDomains(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.ReadAllTrustedEmailDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllTrustedEmailDomains`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.ReadAllTrustedEmailDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllTrustedEmailDomainsRequest struct via the builder pattern


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


## ReadOneTrustedEmailDomain

> EmailDomain ReadOneTrustedEmailDomain(ctx, environmentID, emailDomainID).Execute()

READ One Trusted Email Domain

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
    resp, r, err := apiClient.TrustedEmailDomainsApi.ReadOneTrustedEmailDomain(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.ReadOneTrustedEmailDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneTrustedEmailDomain`: EmailDomain
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.ReadOneTrustedEmailDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneTrustedEmailDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailDomain**](EmailDomain.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTrustedEmailDomainDKIMStatus

> EmailDomainDKIMStatus ReadTrustedEmailDomainDKIMStatus(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain DKIM Status

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
    resp, r, err := apiClient.TrustedEmailDomainsApi.ReadTrustedEmailDomainDKIMStatus(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.ReadTrustedEmailDomainDKIMStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTrustedEmailDomainDKIMStatus`: EmailDomainDKIMStatus
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.ReadTrustedEmailDomainDKIMStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTrustedEmailDomainDKIMStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailDomainDKIMStatus**](EmailDomainDKIMStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTrustedEmailDomainOwnershipStatus

> EmailDomainOwnershipStatus ReadTrustedEmailDomainOwnershipStatus(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain Ownership Status

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
    resp, r, err := apiClient.TrustedEmailDomainsApi.ReadTrustedEmailDomainOwnershipStatus(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.ReadTrustedEmailDomainOwnershipStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTrustedEmailDomainOwnershipStatus`: EmailDomainOwnershipStatus
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.ReadTrustedEmailDomainOwnershipStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTrustedEmailDomainOwnershipStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailDomainOwnershipStatus**](EmailDomainOwnershipStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTrustedEmailDomainSPFStatus

> EmailDomainSPFStatus ReadTrustedEmailDomainSPFStatus(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain SPF Status

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
    resp, r, err := apiClient.TrustedEmailDomainsApi.ReadTrustedEmailDomainSPFStatus(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrustedEmailDomainsApi.ReadTrustedEmailDomainSPFStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTrustedEmailDomainSPFStatus`: EmailDomainSPFStatus
    fmt.Fprintf(os.Stdout, "Response from `TrustedEmailDomainsApi.ReadTrustedEmailDomainSPFStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**emailDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadTrustedEmailDomainSPFStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailDomainSPFStatus**](EmailDomainSPFStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

