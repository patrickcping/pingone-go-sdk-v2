# \NotificationsTrustedEmailDomainsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete) | **Delete** /v1/environments/{environmentID}/emailDomains/{emailDomainID} | DELETE Trusted Email Domain
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/dkim | READ Trusted Email Domain DKIM Status
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID} | READ One Trusted Email Domain
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/ownership | READ Trusted Email Domain Ownership Status
[**V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet) | **Get** /v1/environments/{environmentID}/emailDomains/{emailDomainID}/spf | READ Trusted Email Domain SPF Status
[**V1EnvironmentsEnvironmentIDEmailDomainsGet**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsGet) | **Get** /v1/environments/{environmentID}/emailDomains | READ All Trusted Email Domains
[**V1EnvironmentsEnvironmentIDEmailDomainsPost**](NotificationsTrustedEmailDomainsApi.md#V1EnvironmentsEnvironmentIDEmailDomainsPost) | **Post** /v1/environments/{environmentID}/emailDomains | CREATE Trusted Email Domain



## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete(ctx, environmentID, emailDomainID).Execute()

DELETE Trusted Email Domain

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
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain DKIM Status

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
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDDkimGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet(ctx, environmentID, emailDomainID).Execute()

READ One Trusted Email Domain

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
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain Ownership Status

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
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDOwnershipGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet

> V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet(ctx, environmentID, emailDomainID).Execute()

READ Trusted Email Domain SPF Status

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
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet(context.Background(), environmentID, emailDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsEmailDomainIDSpfGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsGet

> V1EnvironmentsEnvironmentIDEmailDomainsGet(ctx, environmentID).Execute()

READ All Trusted Email Domains

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDEmailDomainsPost

> V1EnvironmentsEnvironmentIDEmailDomainsPost(ctx, environmentID).Body(body).Execute()

CREATE Trusted Email Domain

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTrustedEmailDomainsApi.V1EnvironmentsEnvironmentIDEmailDomainsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEmailDomainsPostRequest struct via the builder pattern


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

