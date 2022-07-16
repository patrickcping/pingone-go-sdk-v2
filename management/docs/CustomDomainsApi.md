# \CustomDomainsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete**](CustomDomainsApi.md#V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete) | **Delete** /v1/environments/{environmentID}/customDomains/{customDomainID} | DELETE Domain
[**V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet**](CustomDomainsApi.md#V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet) | **Get** /v1/environments/{environmentID}/customDomains/{customDomainID} | READ One Domain
[**V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost**](CustomDomainsApi.md#V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost) | **Post** /v1/environments/{environmentID}/customDomains/{customDomainID} | Import Certificate
[**V1EnvironmentsEnvironmentIDCustomDomainsGet**](CustomDomainsApi.md#V1EnvironmentsEnvironmentIDCustomDomainsGet) | **Get** /v1/environments/{environmentID}/customDomains | READ All Domains
[**V1EnvironmentsEnvironmentIDCustomDomainsPost**](CustomDomainsApi.md#V1EnvironmentsEnvironmentIDCustomDomainsPost) | **Post** /v1/environments/{environmentID}/customDomains | CREATE Domain



## V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete

> V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete(ctx, environmentID, customDomainID).Execute()

DELETE Domain

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
    customDomainID := "customDomainID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete(context.Background(), environmentID, customDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet

> V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet(ctx, environmentID, customDomainID).Execute()

READ One Domain

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
    customDomainID := "customDomainID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet(context.Background(), environmentID, customDomainID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost

> V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost(ctx, environmentID, customDomainID).ContentType(contentType).Body(body).Execute()

Import Certificate

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
    customDomainID := "customDomainID_example" // string | 
    contentType := "application/vnd.pingidentity.certificate.import+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost(context.Background(), environmentID, customDomainID).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customDomainID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCustomDomainsCustomDomainIDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDCustomDomainsGet

> V1EnvironmentsEnvironmentIDCustomDomainsGet(ctx, environmentID).Execute()

READ All Domains

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
    resp, r, err := apiClient.CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCustomDomainsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCustomDomainsPost

> V1EnvironmentsEnvironmentIDCustomDomainsPost(ctx, environmentID).Body(body).Execute()

CREATE Domain

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
    resp, r, err := apiClient.CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomDomainsApi.V1EnvironmentsEnvironmentIDCustomDomainsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCustomDomainsPostRequest struct via the builder pattern


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

