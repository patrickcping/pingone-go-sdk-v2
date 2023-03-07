# \BrandingSettingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadBrandingSettings**](BrandingSettingsApi.md#ReadBrandingSettings) | **Get** /v1/environments/{environmentID}/brandingSettings | READ Branding Settings
[**UpdateBrandingSettings**](BrandingSettingsApi.md#UpdateBrandingSettings) | **Put** /v1/environments/{environmentID}/brandingSettings | UPDATE Branding Settings



## ReadBrandingSettings

> BrandingSettings ReadBrandingSettings(ctx, environmentID).Execute()

READ Branding Settings

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
    resp, r, err := apiClient.BrandingSettingsApi.ReadBrandingSettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingSettingsApi.ReadBrandingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadBrandingSettings`: BrandingSettings
    fmt.Fprintf(os.Stdout, "Response from `BrandingSettingsApi.ReadBrandingSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadBrandingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrandingSettings**](BrandingSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrandingSettings

> BrandingSettings UpdateBrandingSettings(ctx, environmentID).BrandingSettings(brandingSettings).Execute()

UPDATE Branding Settings

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
    brandingSettings := *openapiclient.NewBrandingSettings() // BrandingSettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingSettingsApi.UpdateBrandingSettings(context.Background(), environmentID).BrandingSettings(brandingSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingSettingsApi.UpdateBrandingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingSettings`: BrandingSettings
    fmt.Fprintf(os.Stdout, "Response from `BrandingSettingsApi.UpdateBrandingSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **brandingSettings** | [**BrandingSettings**](BrandingSettings.md) |  | 

### Return type

[**BrandingSettings**](BrandingSettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

