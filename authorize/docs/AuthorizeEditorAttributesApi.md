# \AuthorizeEditorAttributesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttribute**](AuthorizeEditorAttributesApi.md#CreateAttribute) | **Post** /environments/{environmentID}/authorizationAttributes | Create an Attribute
[**DeleteAttribute**](AuthorizeEditorAttributesApi.md#DeleteAttribute) | **Delete** /environments/{environmentID}/authorizationAttributes/{attributeID} | Delete an Attribute
[**GetAttribute**](AuthorizeEditorAttributesApi.md#GetAttribute) | **Get** /environments/{environmentID}/authorizationAttributes/{attributeID} | Get an Attribute by ID
[**ListAttributes**](AuthorizeEditorAttributesApi.md#ListAttributes) | **Get** /environments/{environmentID}/authorizationAttributes | List Attributes
[**TestAttribute**](AuthorizeEditorAttributesApi.md#TestAttribute) | **Post** /environments/{environmentID}/authorizationAttributes/{attributeID} | Test an Attribute by ID
[**UpdateAttribute**](AuthorizeEditorAttributesApi.md#UpdateAttribute) | **Put** /environments/{environmentID}/authorizationAttributes/{attributeID} | Update an Attribute



## CreateAttribute

> AuthorizeEditorDataDefinitionsAttributeDefinitionDTO CreateAttribute(ctx, environmentID).AuthorizeEditorDataDefinitionsAttributeDefinitionDTO(authorizeEditorDataDefinitionsAttributeDefinitionDTO).Execute()

Create an Attribute



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the attribute.
    authorizeEditorDataDefinitionsAttributeDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataValueTypeDTO("Type_example")) // AuthorizeEditorDataDefinitionsAttributeDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorAttributesApi.CreateAttribute(context.Background(), environmentID).AuthorizeEditorDataDefinitionsAttributeDefinitionDTO(authorizeEditorDataDefinitionsAttributeDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.CreateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttribute`: AuthorizeEditorDataDefinitionsAttributeDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorAttributesApi.CreateAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataDefinitionsAttributeDefinitionDTO** | [**AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttribute

> DeleteAttribute(ctx, environmentID, attributeID).Execute()

Delete an Attribute



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the attribute.
    attributeID := "attributeID_example" // string | The ID of the attribute to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorAttributesApi.DeleteAttribute(context.Background(), environmentID, attributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.DeleteAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the attribute. | 
**attributeID** | **string** | The ID of the attribute to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeRequest struct via the builder pattern


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


## GetAttribute

> AuthorizeEditorDataDefinitionsAttributeDefinitionDTO GetAttribute(ctx, environmentID, attributeID).Execute()

Get an Attribute by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the attribute.
    attributeID := "attributeID_example" // string | The ID of the attribute to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorAttributesApi.GetAttribute(context.Background(), environmentID, attributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.GetAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttribute`: AuthorizeEditorDataDefinitionsAttributeDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorAttributesApi.GetAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the attribute. | 
**attributeID** | **string** | The ID of the attribute to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAttributes

> ListAttributes200Response ListAttributes(ctx, environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()

List Attributes



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list attributes.
    filter := "filter_example" // string | The parent filter that may be used to list the root attributes or children of an attribute. The following filter options can be applied to attributes:     * `parent not pr` (root attributes)     * `parent eq id` (attributes with parent id equal to id) (optional)
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorAttributesApi.ListAttributes(context.Background(), environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.ListAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAttributes`: ListAttributes200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorAttributesApi.ListAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list attributes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The parent filter that may be used to list the root attributes or children of an attribute. The following filter options can be applied to attributes:     * &#x60;parent not pr&#x60; (root attributes)     * &#x60;parent eq id&#x60; (attributes with parent id equal to id) | 
 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListAttributes200Response**](ListAttributes200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestAttribute

> AuthorizeEditorDataEntityTestResponseDTO TestAttribute(ctx, environmentID, attributeID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()

Test an Attribute by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    attributeID := "attributeID_example" // string | 
    authorizeEditorDataEntityTestRequestDTO := *openapiclient.NewAuthorizeEditorDataEntityTestRequestDTO([]openapiclient.AuthorizeEditorDataEntityTestingParameterDTO{*openapiclient.NewAuthorizeEditorDataEntityTestingParameterDTO("Key_example", "Value_example")}) // AuthorizeEditorDataEntityTestRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorAttributesApi.TestAttribute(context.Background(), environmentID, attributeID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.TestAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestAttribute`: AuthorizeEditorDataEntityTestResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorAttributesApi.TestAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataEntityTestRequestDTO** | [**AuthorizeEditorDataEntityTestRequestDTO**](AuthorizeEditorDataEntityTestRequestDTO.md) |  | 

### Return type

[**AuthorizeEditorDataEntityTestResponseDTO**](AuthorizeEditorDataEntityTestResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> AuthorizeEditorDataDefinitionsAttributeDefinitionDTO UpdateAttribute(ctx, environmentID, attributeID).AuthorizeEditorDataDefinitionsAttributeDefinitionDTO(authorizeEditorDataDefinitionsAttributeDefinitionDTO).Execute()

Update an Attribute



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the attribute.
    attributeID := "attributeID_example" // string | The ID of the attribute to be updated
    authorizeEditorDataDefinitionsAttributeDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsAttributeDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataValueTypeDTO("Type_example")) // AuthorizeEditorDataDefinitionsAttributeDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorAttributesApi.UpdateAttribute(context.Background(), environmentID, attributeID).AuthorizeEditorDataDefinitionsAttributeDefinitionDTO(authorizeEditorDataDefinitionsAttributeDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorAttributesApi.UpdateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttribute`: AuthorizeEditorDataDefinitionsAttributeDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorAttributesApi.UpdateAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the attribute. | 
**attributeID** | **string** | The ID of the attribute to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataDefinitionsAttributeDefinitionDTO** | [**AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsAttributeDefinitionDTO**](AuthorizeEditorDataDefinitionsAttributeDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
