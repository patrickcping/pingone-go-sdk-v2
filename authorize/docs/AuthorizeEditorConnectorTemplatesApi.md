# \AuthorizeEditorConnectorTemplatesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectorTemplate**](AuthorizeEditorConnectorTemplatesApi.md#GetConnectorTemplate) | **Get** /environments/{environmentID}/authorizationConnectorTemplates/{code} | Get a Connector Template by code
[**ListConnectorTemplates**](AuthorizeEditorConnectorTemplatesApi.md#ListConnectorTemplates) | **Get** /environments/{environmentID}/authorizationConnectorTemplates | List Connector Templates



## GetConnectorTemplate

> AuthorizeEditorDataConnectorsConnectorTemplateDTO GetConnectorTemplate(ctx, environmentID, code).Execute()

Get a Connector Template by code



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the connector template.
    code := "code_example" // string | The code of the connector template to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConnectorTemplatesApi.GetConnectorTemplate(context.Background(), environmentID, code).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConnectorTemplatesApi.GetConnectorTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorTemplate`: AuthorizeEditorDataConnectorsConnectorTemplateDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConnectorTemplatesApi.GetConnectorTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the connector template. | 
**code** | **string** | The code of the connector template to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataConnectorsConnectorTemplateDTO**](AuthorizeEditorDataConnectorsConnectorTemplateDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectorTemplates

> ListConnectorTemplates200Response ListConnectorTemplates(ctx, environmentID).Limit(limit).Cursor(cursor).Execute()

List Connector Templates



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list connector templates.
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConnectorTemplatesApi.ListConnectorTemplates(context.Background(), environmentID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConnectorTemplatesApi.ListConnectorTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectorTemplates`: ListConnectorTemplates200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConnectorTemplatesApi.ListConnectorTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list connector templates. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectorTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListConnectorTemplates200Response**](ListConnectorTemplates200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
