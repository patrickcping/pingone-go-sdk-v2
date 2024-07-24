# \AuthorizeEditorProcessorsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProcessor**](AuthorizeEditorProcessorsApi.md#CreateProcessor) | **Post** /environments/{environmentID}/authorizationProcessors | Create a Processor
[**DeleteProcessor**](AuthorizeEditorProcessorsApi.md#DeleteProcessor) | **Delete** /environments/{environmentID}/authorizationProcessors/{processorID} | Delete a Processor
[**GetProcessor**](AuthorizeEditorProcessorsApi.md#GetProcessor) | **Get** /environments/{environmentID}/authorizationProcessors/{processorID} | Get a Processor by ID
[**ListProcessors**](AuthorizeEditorProcessorsApi.md#ListProcessors) | **Get** /environments/{environmentID}/authorizationProcessors | List Processors
[**UpdateProcessor**](AuthorizeEditorProcessorsApi.md#UpdateProcessor) | **Put** /environments/{environmentID}/authorizationProcessors/{processorID} | Update a Processor



## CreateProcessor

> AuthorizeEditorDataDefinitionsProcessorDefinitionDTO CreateProcessor(ctx, environmentID).AuthorizeEditorDataDefinitionsProcessorDefinitionDTO(authorizeEditorDataDefinitionsProcessorDefinitionDTO).Execute()

Create a Processor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the processor.
    authorizeEditorDataDefinitionsProcessorDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataProcessorDTO("Name_example", "Type_example")) // AuthorizeEditorDataDefinitionsProcessorDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorProcessorsApi.CreateProcessor(context.Background(), environmentID).AuthorizeEditorDataDefinitionsProcessorDefinitionDTO(authorizeEditorDataDefinitionsProcessorDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorProcessorsApi.CreateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProcessor`: AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorProcessorsApi.CreateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the processor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataDefinitionsProcessorDefinitionDTO** | [**AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProcessor

> DeleteProcessor(ctx, environmentID, processorID).Execute()

Delete a Processor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the processor.
    processorID := "processorID_example" // string | The ID of the processor to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorProcessorsApi.DeleteProcessor(context.Background(), environmentID, processorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorProcessorsApi.DeleteProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the processor. | 
**processorID** | **string** | The ID of the processor to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProcessorRequest struct via the builder pattern


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


## GetProcessor

> AuthorizeEditorDataDefinitionsProcessorDefinitionDTO GetProcessor(ctx, environmentID, processorID).Execute()

Get a Processor by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the processor.
    processorID := "processorID_example" // string | The ID of the processor to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorProcessorsApi.GetProcessor(context.Background(), environmentID, processorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorProcessorsApi.GetProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProcessor`: AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorProcessorsApi.GetProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the processor. | 
**processorID** | **string** | The ID of the processor to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProcessors

> ListProcessors200Response ListProcessors(ctx, environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()

List Processors



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list processors.
    filter := "filter_example" // string | The parent filter that may be used to list the root processors or children of a processor. The following filter options can be applied to processors:     * `parent not pr` (root processors)     * `parent eq id` (processors with parent id equal to id) (optional)
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorProcessorsApi.ListProcessors(context.Background(), environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorProcessorsApi.ListProcessors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProcessors`: ListProcessors200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorProcessorsApi.ListProcessors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list processors. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListProcessorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The parent filter that may be used to list the root processors or children of a processor. The following filter options can be applied to processors:     * &#x60;parent not pr&#x60; (root processors)     * &#x60;parent eq id&#x60; (processors with parent id equal to id) | 
 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListProcessors200Response**](ListProcessors200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProcessor

> AuthorizeEditorDataDefinitionsProcessorDefinitionDTO UpdateProcessor(ctx, environmentID, processorID).AuthorizeEditorDataDefinitionsProcessorDefinitionDTO(authorizeEditorDataDefinitionsProcessorDefinitionDTO).Execute()

Update a Processor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the processor.
    processorID := "processorID_example" // string | The ID of the processor to be updated
    authorizeEditorDataDefinitionsProcessorDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataProcessorDTO("Name_example", "Type_example")) // AuthorizeEditorDataDefinitionsProcessorDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorProcessorsApi.UpdateProcessor(context.Background(), environmentID, processorID).AuthorizeEditorDataDefinitionsProcessorDefinitionDTO(authorizeEditorDataDefinitionsProcessorDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorProcessorsApi.UpdateProcessor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProcessor`: AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorProcessorsApi.UpdateProcessor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the processor. | 
**processorID** | **string** | The ID of the processor to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProcessorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataDefinitionsProcessorDefinitionDTO** | [**AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsProcessorDefinitionDTO**](AuthorizeEditorDataDefinitionsProcessorDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
