# \ApplicationAttributeMappingApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationAttributeMapping**](ApplicationAttributeMappingApi.md#CreateApplicationAttributeMapping) | **Post** /environments/{environmentID}/applications/{applicationID}/attributes | CREATE Application Attribute Mapping
[**DeleteApplicationAttributeMapping**](ApplicationAttributeMappingApi.md#DeleteApplicationAttributeMapping) | **Delete** /environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID} | DELETE Application Attribute Mapping
[**ReadAllApplicationAttributeMappings**](ApplicationAttributeMappingApi.md#ReadAllApplicationAttributeMappings) | **Get** /environments/{environmentID}/applications/{applicationID}/attributes | READ All Application Attribute Mappings
[**ReadOneApplicationAttributeMapping**](ApplicationAttributeMappingApi.md#ReadOneApplicationAttributeMapping) | **Get** /environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID} | READ One Application Attribute Mapping
[**UpdateApplicationAttributeMapping**](ApplicationAttributeMappingApi.md#UpdateApplicationAttributeMapping) | **Put** /environments/{environmentID}/applications/{applicationID}/attributes/{attrMappingID} | UPDATE Application Attribute Mapping



## CreateApplicationAttributeMapping

> ApplicationAttributeMapping CreateApplicationAttributeMapping(ctx, environmentID, applicationID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()

CREATE Application Attribute Mapping

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
    applicationID := "applicationID_example" // string | 
    applicationAttributeMapping := *openapiclient.NewApplicationAttributeMapping("Name_example", false, "Value_example") // ApplicationAttributeMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAttributeMappingApi.CreateApplicationAttributeMapping(context.Background(), environmentID, applicationID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAttributeMappingApi.CreateApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAttributeMappingApi.CreateApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationAttributeMapping** | [**ApplicationAttributeMapping**](ApplicationAttributeMapping.md) |  | 

### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationAttributeMapping

> DeleteApplicationAttributeMapping(ctx, environmentID, applicationID, attrMappingID).Execute()

DELETE Application Attribute Mapping

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
    applicationID := "applicationID_example" // string | 
    attrMappingID := "attrMappingID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationAttributeMappingApi.DeleteApplicationAttributeMapping(context.Background(), environmentID, applicationID, attrMappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAttributeMappingApi.DeleteApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationAttributeMappingRequest struct via the builder pattern


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


## ReadAllApplicationAttributeMappings

> EntityArray ReadAllApplicationAttributeMappings(ctx, environmentID, applicationID).Execute()

READ All Application Attribute Mappings

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
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationAttributeMappings`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAttributeMappingApi.ReadAllApplicationAttributeMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationAttributeMappingsRequest struct via the builder pattern


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


## ReadOneApplicationAttributeMapping

> ApplicationAttributeMapping ReadOneApplicationAttributeMapping(ctx, environmentID, applicationID, attrMappingID).Execute()

READ One Application Attribute Mapping

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
    applicationID := "applicationID_example" // string | 
    attrMappingID := "attrMappingID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping(context.Background(), environmentID, applicationID, attrMappingID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAttributeMappingApi.ReadOneApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationAttributeMapping

> ApplicationAttributeMapping UpdateApplicationAttributeMapping(ctx, environmentID, applicationID, attrMappingID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()

UPDATE Application Attribute Mapping

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
    applicationID := "applicationID_example" // string | 
    attrMappingID := "attrMappingID_example" // string | 
    applicationAttributeMapping := *openapiclient.NewApplicationAttributeMapping("Name_example", false, "Value_example") // ApplicationAttributeMapping |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationAttributeMappingApi.UpdateApplicationAttributeMapping(context.Background(), environmentID, applicationID, attrMappingID).ApplicationAttributeMapping(applicationAttributeMapping).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAttributeMappingApi.UpdateApplicationAttributeMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationAttributeMapping`: ApplicationAttributeMapping
    fmt.Fprintf(os.Stdout, "Response from `ApplicationAttributeMappingApi.UpdateApplicationAttributeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**attrMappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationAttributeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationAttributeMapping** | [**ApplicationAttributeMapping**](ApplicationAttributeMapping.md) |  | 

### Return type

[**ApplicationAttributeMapping**](ApplicationAttributeMapping.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

