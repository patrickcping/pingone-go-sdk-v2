# \RecaptchaConfigurationApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecaptchaConfiguration**](RecaptchaConfigurationApi.md#DeleteRecaptchaConfiguration) | **Delete** /environments/{environmentID}/recaptchaV2Config | DELETE Recaptcha Configuration
[**ReadRecaptchaConfiguration**](RecaptchaConfigurationApi.md#ReadRecaptchaConfiguration) | **Get** /environments/{environmentID}/recaptchaV2Config | READ Recaptcha Configuration
[**UpdateRecaptchaConfiguration**](RecaptchaConfigurationApi.md#UpdateRecaptchaConfiguration) | **Put** /environments/{environmentID}/recaptchaV2Config | UPDATE Recaptcha Configuration



## DeleteRecaptchaConfiguration

> DeleteRecaptchaConfiguration(ctx, environmentID).Execute()

DELETE Recaptcha Configuration

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
    r, err := apiClient.RecaptchaConfigurationApi.DeleteRecaptchaConfiguration(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecaptchaConfigurationApi.DeleteRecaptchaConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRecaptchaConfigurationRequest struct via the builder pattern


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


## ReadRecaptchaConfiguration

> RecaptchaConfiguration ReadRecaptchaConfiguration(ctx, environmentID).Execute()

READ Recaptcha Configuration

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
    resp, r, err := apiClient.RecaptchaConfigurationApi.ReadRecaptchaConfiguration(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecaptchaConfigurationApi.ReadRecaptchaConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRecaptchaConfiguration`: RecaptchaConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RecaptchaConfigurationApi.ReadRecaptchaConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRecaptchaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecaptchaConfiguration**](RecaptchaConfiguration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecaptchaConfiguration

> RecaptchaConfiguration UpdateRecaptchaConfiguration(ctx, environmentID).RecaptchaConfiguration(recaptchaConfiguration).Execute()

UPDATE Recaptcha Configuration

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
    recaptchaConfiguration := *openapiclient.NewRecaptchaConfiguration("SiteKey_example", "SecretKey_example") // RecaptchaConfiguration |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecaptchaConfigurationApi.UpdateRecaptchaConfiguration(context.Background(), environmentID).RecaptchaConfiguration(recaptchaConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecaptchaConfigurationApi.UpdateRecaptchaConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRecaptchaConfiguration`: RecaptchaConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RecaptchaConfigurationApi.UpdateRecaptchaConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecaptchaConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recaptchaConfiguration** | [**RecaptchaConfiguration**](RecaptchaConfiguration.md) |  | 

### Return type

[**RecaptchaConfiguration**](RecaptchaConfiguration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

