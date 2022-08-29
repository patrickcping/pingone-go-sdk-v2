# \BrandingThemesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDThemesGet**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesGet) | **Get** /v1/environments/{environmentID}/themes | READ Branding Themes
[**V1EnvironmentsEnvironmentIDThemesPost**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesPost) | **Post** /v1/environments/{environmentID}/themes | CREATE Branding Theme
[**V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet) | **Get** /v1/environments/{environmentID}/themes/{themeID}/default | READ Branding Theme Default
[**V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut) | **Put** /v1/environments/{environmentID}/themes/{themeID}/default | UPDATE Branding Theme Default
[**V1EnvironmentsEnvironmentIDThemesThemeIDDelete**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesThemeIDDelete) | **Delete** /v1/environments/{environmentID}/themes/{themeID} | DELETE Branding Theme
[**V1EnvironmentsEnvironmentIDThemesThemeIDGet**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesThemeIDGet) | **Get** /v1/environments/{environmentID}/themes/{themeID} | READ One Branding Theme
[**V1EnvironmentsEnvironmentIDThemesThemeIDPut**](BrandingThemesApi.md#V1EnvironmentsEnvironmentIDThemesThemeIDPut) | **Put** /v1/environments/{environmentID}/themes/{themeID} | UPDATE Branding Theme



## V1EnvironmentsEnvironmentIDThemesGet

> V1EnvironmentsEnvironmentIDThemesGet(ctx, environmentID).Authorization(authorization).Execute()

READ Branding Themes

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
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesGet(context.Background(), environmentID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDThemesPost

> V1EnvironmentsEnvironmentIDThemesPost(ctx, environmentID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

CREATE Branding Theme

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
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesPost(context.Background(), environmentID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet

> V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet(ctx, environmentID, themeID).Authorization(authorization).Execute()

READ Branding Theme Default

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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet(context.Background(), environmentID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDefaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesThemeIDDefaultGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut

> V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut(ctx, environmentID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

UPDATE Branding Theme Default

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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut(context.Background(), environmentID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDefaultPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesThemeIDDefaultPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDThemesThemeIDDelete

> V1EnvironmentsEnvironmentIDThemesThemeIDDelete(ctx, environmentID, themeID).Authorization(authorization).Execute()

DELETE Branding Theme

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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDelete(context.Background(), environmentID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesThemeIDDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDThemesThemeIDGet

> V1EnvironmentsEnvironmentIDThemesThemeIDGet(ctx, environmentID, themeID).Authorization(authorization).Execute()

READ One Branding Theme

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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDGet(context.Background(), environmentID, themeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesThemeIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDThemesThemeIDPut

> V1EnvironmentsEnvironmentIDThemesThemeIDPut(ctx, environmentID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

UPDATE Branding Theme

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
    themeID := "themeID_example" // string | 
    authorization := "{{jwtToken}}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDPut(context.Background(), environmentID, themeID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.V1EnvironmentsEnvironmentIDThemesThemeIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDThemesThemeIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
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

