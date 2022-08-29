# \NotificationsTemplatesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDTemplatesGet**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesGet) | **Get** /v1/environments/{environmentID}/templates | READ All Templates
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete) | **Delete** /v1/environments/{environmentID}/templates/{templateName}/contents/{contentID} | DELETE Content
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet) | **Get** /v1/environments/{environmentID}/templates/{templateName}/contents/{contentID} | READ One Content
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut) | **Put** /v1/environments/{environmentID}/templates/{templateName}/contents/{contentID} | UPDATE Push Content
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete) | **Delete** /v1/environments/{environmentID}/templates/{templateName}/contents | DELETE Bulk Variant Contents
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet) | **Get** /v1/environments/{environmentID}/templates/{templateName}/contents | READ All Contents
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch) | **Patch** /v1/environments/{environmentID}/templates/{templateName}/contents | PATCH Bulk Variant Contents
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost) | **Post** /v1/environments/{environmentID}/templates/{templateName}/contents | CREATE Push Content
[**V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet**](NotificationsTemplatesApi.md#V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet) | **Get** /v1/environments/{environmentID}/templates/{templateName} | READ One Template



## V1EnvironmentsEnvironmentIDTemplatesGet

> V1EnvironmentsEnvironmentIDTemplatesGet(ctx, environmentID).Execute()

READ All Templates

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
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete(ctx, environmentID, templateName, contentID).Execute()

DELETE Content

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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete(context.Background(), environmentID, templateName, contentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet(ctx, environmentID, templateName, contentID).Execute()

READ One Content

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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet(context.Background(), environmentID, templateName, contentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut(ctx, environmentID, templateName, contentID).Body(body).Execute()

UPDATE Push Content

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
    templateName := "templateName_example" // string | 
    contentID := "contentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut(context.Background(), environmentID, templateName, contentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 
**contentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsContentIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete(ctx, environmentID, templateName).Filter(filter).Execute()

DELETE Bulk Variant Contents

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
    templateName := "templateName_example" // string | 
    filter := "variant eq {{variantName}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete(context.Background(), environmentID, templateName).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet(ctx, environmentID, templateName).Execute()

READ All Contents

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
    templateName := "templateName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet(context.Background(), environmentID, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch(ctx, environmentID, templateName).Filter(filter).Body(body).Execute()

PATCH Bulk Variant Contents

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
    templateName := "templateName_example" // string | 
    filter := "variant eq {{variantName}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch(context.Background(), environmentID, templateName).Filter(filter).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost(ctx, environmentID, templateName).Body(body).Execute()

CREATE Push Content

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
    templateName := "templateName_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost(context.Background(), environmentID, templateName).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameContentsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet

> V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet(ctx, environmentID, templateName).Execute()

READ One Template

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
    templateName := "templateName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet(context.Background(), environmentID, templateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsTemplatesApi.V1EnvironmentsEnvironmentIDTemplatesTemplateNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**templateName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTemplatesTemplateNameGetRequest struct via the builder pattern


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

