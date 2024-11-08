# \NotificationsTemplatesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContent**](NotificationsTemplatesApi.md#CreateContent) | **Post** /environments/{environmentID}/templates/{templateName}/contents | CREATE Content
[**DeleteBulkVariantContents**](NotificationsTemplatesApi.md#DeleteBulkVariantContents) | **Delete** /environments/{environmentID}/templates/{templateName}/contents | DELETE Bulk Variant Contents
[**DeleteContent**](NotificationsTemplatesApi.md#DeleteContent) | **Delete** /environments/{environmentID}/templates/{templateName}/contents/{contentID} | DELETE Content
[**PatchBulkVariantContents**](NotificationsTemplatesApi.md#PatchBulkVariantContents) | **Patch** /environments/{environmentID}/templates/{templateName}/contents | PATCH Bulk Variant Contents
[**ReadAllTemplateContents**](NotificationsTemplatesApi.md#ReadAllTemplateContents) | **Get** /environments/{environmentID}/templates/{templateName}/contents | READ All Contents
[**ReadAllTemplates**](NotificationsTemplatesApi.md#ReadAllTemplates) | **Get** /environments/{environmentID}/templates | READ All Templates
[**ReadOneContent**](NotificationsTemplatesApi.md#ReadOneContent) | **Get** /environments/{environmentID}/templates/{templateName}/contents/{contentID} | READ One Content
[**ReadOneTemplate**](NotificationsTemplatesApi.md#ReadOneTemplate) | **Get** /environments/{environmentID}/templates/{templateName} | READ One Template
[**UpdateContent**](NotificationsTemplatesApi.md#UpdateContent) | **Put** /environments/{environmentID}/templates/{templateName}/contents/{contentID} | UPDATE Content



## CreateContent

> TemplateContent CreateContent(ctx, environmentID, templateName).TemplateContent(templateContent).Execute()

CREATE Content

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    templateContent := openapiclient.TemplateContent{TemplateContentEmail: openapiclient.NewTemplateContentEmail("Locale_example", openapiclient.EnumTemplateContentDeliveryMethod("Email"), "Body_example")} // TemplateContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.CreateContent(context.Background(), environmentID, templateName).TemplateContent(templateContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.CreateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContent`: TemplateContent
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.CreateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **templateContent** | [**TemplateContent**](TemplateContent.md) |  | 

### Return type

[**TemplateContent**](TemplateContent.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBulkVariantContents

> DeleteBulkVariantContents(ctx, environmentID, templateName).Filter(filter).Execute()

DELETE Bulk Variant Contents

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    filter := "variant eq {{variantName}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsTemplatesApi.DeleteBulkVariantContents(context.Background(), environmentID, templateName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.DeleteBulkVariantContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkVariantContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** |  | 

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


## DeleteContent

> DeleteContent(ctx, environmentID, templateName, contentID).Execute()

DELETE Content

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    contentID := "contentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsTemplatesApi.DeleteContent(context.Background(), environmentID, templateName, contentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.DeleteContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContentRequest struct via the builder pattern


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


## PatchBulkVariantContents

> EntityArray PatchBulkVariantContents(ctx, environmentID, templateName).Filter(filter).Body(body).Execute()

PATCH Bulk Variant Contents

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    filter := "variant eq {{variantName}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.PatchBulkVariantContents(context.Background(), environmentID, templateName).Filter(filter).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.PatchBulkVariantContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchBulkVariantContents`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.PatchBulkVariantContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchBulkVariantContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllTemplateContents

> EntityArray ReadAllTemplateContents(ctx, environmentID, templateName).Limit(limit).Execute()

READ All Contents

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.ReadAllTemplateContents(context.Background(), environmentID, templateName).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.ReadAllTemplateContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllTemplateContents`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.ReadAllTemplateContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllTemplateContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 

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


## ReadAllTemplates

> EntityArray ReadAllTemplates(ctx, environmentID).Limit(limit).Filter(filter).Order(order).Execute()

READ All Templates

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
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    filter := "(createdAt lt "2018-08-30") and (updatedAt gt "2018-07-30")" // string |  (optional)
    order := "-createdAt" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.ReadAllTemplates(context.Background(), environmentID).Limit(limit).Filter(filter).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.ReadAllTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllTemplates`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.ReadAllTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** |  | 
 **order** | **string** |  | 

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


## ReadOneContent

> TemplateContent ReadOneContent(ctx, environmentID, templateName, contentID).Execute()

READ One Content

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    contentID := "contentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.ReadOneContent(context.Background(), environmentID, templateName, contentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.ReadOneContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneContent`: TemplateContent
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.ReadOneContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TemplateContent**](TemplateContent.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneTemplate

> Template ReadOneTemplate(ctx, environmentID, templateName).Execute()

READ One Template

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.ReadOneTemplate(context.Background(), environmentID, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.ReadOneTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneTemplate`: Template
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.ReadOneTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Template**](Template.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContent

> TemplateContent UpdateContent(ctx, environmentID, templateName, contentID).TemplateContent(templateContent).Execute()

UPDATE Content

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
    templateName := openapiclient.EnumTemplateName("credential_issued") // EnumTemplateName | 
    contentID := "contentID_example" // string | 
    templateContent := openapiclient.TemplateContent{TemplateContentEmail: openapiclient.NewTemplateContentEmail("Locale_example", openapiclient.EnumTemplateContentDeliveryMethod("Email"), "Body_example")} // TemplateContent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.UpdateContent(context.Background(), environmentID, templateName, contentID).TemplateContent(templateContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.UpdateContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContent`: TemplateContent
    fmt.Fprintf(os.Stdout, "Response from `NotificationsTemplatesApi.UpdateContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | [**EnumTemplateName**](.md) |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **templateContent** | [**TemplateContent**](TemplateContent.md) |  | 

### Return type

[**TemplateContent**](TemplateContent.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

