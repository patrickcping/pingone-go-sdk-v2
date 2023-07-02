# \EnvironmentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentActiveLicense**](EnvironmentsApi.md#CreateEnvironmentActiveLicense) | **Post** /environments | CREATE Environment (Active License)
[**DeleteEnvironment**](EnvironmentsApi.md#DeleteEnvironment) | **Delete** /environments/{environmentID} | DELETE Environment
[**ReadAllEnvironments**](EnvironmentsApi.md#ReadAllEnvironments) | **Get** /environments | READ All Environments
[**ReadOneEnvironment**](EnvironmentsApi.md#ReadOneEnvironment) | **Get** /environments/{environmentID} | READ One Environment
[**UpdateEnvironment**](EnvironmentsApi.md#UpdateEnvironment) | **Put** /environments/{environmentID} | UPDATE Environment
[**UpdateEnvironmentType**](EnvironmentsApi.md#UpdateEnvironmentType) | **Put** /environments/{environmentID}/type | UPDATE Environment Type



## CreateEnvironmentActiveLicense

> Environment CreateEnvironmentActiveLicense(ctx).Environment(environment).Execute()

CREATE Environment (Active License)

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
    environment := *openapiclient.NewEnvironment(*openapiclient.NewEnvironmentLicense("Id_example"), "Name_example", openapiclient.EnumRegionCode("AP"), openapiclient.EnumEnvironmentType("PRODUCTION")) // Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.CreateEnvironmentActiveLicense(context.Background()).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironmentActiveLicense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEnvironmentActiveLicense`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironmentActiveLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentActiveLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environment** | [**Environment**](Environment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironment

> DeleteEnvironment(ctx, environmentID).Execute()

DELETE Environment

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
    r, err := apiClient.EnvironmentsApi.DeleteEnvironment(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEnvironmentRequest struct via the builder pattern


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


## ReadAllEnvironments

> EntityArray ReadAllEnvironments(ctx).Limit(limit).Filter(filter).Execute()

READ All Environments

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
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    filter := "name sw "S" and license.id eq "34f0efac-21d9-4a17-8a35-196bb3362983"" // string | Adding a SCIM filter for an environment to display only those resources associated with the specified environment. 'sw', 'eq' and 'and' are supported (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.ReadAllEnvironments(context.Background()).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ReadAllEnvironments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllEnvironments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ReadAllEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadAllEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** | Adding a SCIM filter for an environment to display only those resources associated with the specified environment. &#39;sw&#39;, &#39;eq&#39; and &#39;and&#39; are supported | 

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


## ReadOneEnvironment

> Environment ReadOneEnvironment(ctx, environmentID).Execute()

READ One Environment

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
    resp, r, err := apiClient.EnvironmentsApi.ReadOneEnvironment(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ReadOneEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ReadOneEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironment

> Environment UpdateEnvironment(ctx, environmentID).Environment(environment).Execute()

UPDATE Environment

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
    environment := *openapiclient.NewEnvironment(*openapiclient.NewEnvironmentLicense("Id_example"), "Name_example", openapiclient.EnumRegionCode("AP"), openapiclient.EnumEnvironmentType("PRODUCTION")) // Environment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateEnvironment(context.Background(), environmentID).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironment`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environment** | [**Environment**](Environment.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnvironmentType

> Environment UpdateEnvironmentType(ctx, environmentID).UpdateEnvironmentTypeRequest(updateEnvironmentTypeRequest).Execute()

UPDATE Environment Type

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
    updateEnvironmentTypeRequest := *openapiclient.NewUpdateEnvironmentTypeRequest() // UpdateEnvironmentTypeRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.UpdateEnvironmentType(context.Background(), environmentID).UpdateEnvironmentTypeRequest(updateEnvironmentTypeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.UpdateEnvironmentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEnvironmentType`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.UpdateEnvironmentType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnvironmentTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEnvironmentTypeRequest** | [**UpdateEnvironmentTypeRequest**](UpdateEnvironmentTypeRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

