# \BrandingThemesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandingTheme**](BrandingThemesApi.md#CreateBrandingTheme) | **Post** /v1/environments/{environmentID}/themes | CREATE Branding Theme
[**DeleteBrandingTheme**](BrandingThemesApi.md#DeleteBrandingTheme) | **Delete** /v1/environments/{environmentID}/themes/{themeID} | DELETE Branding Theme
[**ReadBrandingThemeDefault**](BrandingThemesApi.md#ReadBrandingThemeDefault) | **Get** /v1/environments/{environmentID}/themes/{themeID}/default | READ Branding Theme Default
[**ReadBrandingThemes**](BrandingThemesApi.md#ReadBrandingThemes) | **Get** /v1/environments/{environmentID}/themes | READ Branding Themes
[**ReadOneBrandingTheme**](BrandingThemesApi.md#ReadOneBrandingTheme) | **Get** /v1/environments/{environmentID}/themes/{themeID} | READ One Branding Theme
[**UpdateBrandingTheme**](BrandingThemesApi.md#UpdateBrandingTheme) | **Put** /v1/environments/{environmentID}/themes/{themeID} | UPDATE Branding Theme
[**UpdateBrandingThemeDefault**](BrandingThemesApi.md#UpdateBrandingThemeDefault) | **Put** /v1/environments/{environmentID}/themes/{themeID}/default | UPDATE Branding Theme Default



## CreateBrandingTheme

> BrandingTheme CreateBrandingTheme(ctx, environmentID).BrandingTheme(brandingTheme).Execute()

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
    brandingTheme := *openapiclient.NewBrandingTheme(*openapiclient.NewBrandingThemeConfiguration("BackgroundColor_example", openapiclient.EnumBrandingThemeBackgroundType("NONE"), openapiclient.EnumBrandingLogoType("IMAGE")), false, openapiclient.EnumBrandingThemeTemplate("default")) // BrandingTheme |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.CreateBrandingTheme(context.Background(), environmentID).BrandingTheme(brandingTheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.CreateBrandingTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingTheme`: BrandingTheme
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.CreateBrandingTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brandingTheme** | [**BrandingTheme**](BrandingTheme.md) |  | 

### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrandingTheme

> DeleteBrandingTheme(ctx, environmentID, themeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.DeleteBrandingTheme(context.Background(), environmentID, themeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.DeleteBrandingTheme``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteBrandingThemeRequest struct via the builder pattern


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


## ReadBrandingThemeDefault

> BrandingThemeDefault ReadBrandingThemeDefault(ctx, environmentID, themeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.ReadBrandingThemeDefault(context.Background(), environmentID, themeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.ReadBrandingThemeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBrandingThemeDefault`: BrandingThemeDefault
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.ReadBrandingThemeDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBrandingThemeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BrandingThemeDefault**](BrandingThemeDefault.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadBrandingThemes

> EntityArray ReadBrandingThemes(ctx, environmentID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.ReadBrandingThemes(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.ReadBrandingThemes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBrandingThemes`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.ReadBrandingThemes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBrandingThemesRequest struct via the builder pattern


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


## ReadOneBrandingTheme

> BrandingTheme ReadOneBrandingTheme(ctx, environmentID, themeID).Execute()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.ReadOneBrandingTheme(context.Background(), environmentID, themeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.ReadOneBrandingTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneBrandingTheme`: BrandingTheme
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.ReadOneBrandingTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneBrandingThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrandingTheme

> BrandingTheme UpdateBrandingTheme(ctx, environmentID, themeID).BrandingTheme(brandingTheme).Execute()

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
    brandingTheme := *openapiclient.NewBrandingTheme(*openapiclient.NewBrandingThemeConfiguration("BackgroundColor_example", openapiclient.EnumBrandingThemeBackgroundType("NONE"), openapiclient.EnumBrandingLogoType("IMAGE")), false, openapiclient.EnumBrandingThemeTemplate("default")) // BrandingTheme |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.UpdateBrandingTheme(context.Background(), environmentID, themeID).BrandingTheme(brandingTheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.UpdateBrandingTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingTheme`: BrandingTheme
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.UpdateBrandingTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingThemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **brandingTheme** | [**BrandingTheme**](BrandingTheme.md) |  | 

### Return type

[**BrandingTheme**](BrandingTheme.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrandingThemeDefault

> BrandingThemeDefault UpdateBrandingThemeDefault(ctx, environmentID, themeID).BrandingThemeDefault(brandingThemeDefault).Execute()

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
    brandingThemeDefault := *openapiclient.NewBrandingThemeDefault(false) // BrandingThemeDefault |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingThemesApi.UpdateBrandingThemeDefault(context.Background(), environmentID, themeID).BrandingThemeDefault(brandingThemeDefault).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingThemesApi.UpdateBrandingThemeDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingThemeDefault`: BrandingThemeDefault
    fmt.Fprintf(os.Stdout, "Response from `BrandingThemesApi.UpdateBrandingThemeDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**themeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingThemeDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **brandingThemeDefault** | [**BrandingThemeDefault**](BrandingThemeDefault.md) |  | 

### Return type

[**BrandingThemeDefault**](BrandingThemeDefault.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

