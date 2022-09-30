# \LanguageLocalizationStatusApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLanguageLocalizationStatus**](LanguageLocalizationStatusApi.md#CreateLanguageLocalizationStatus) | **Post** /v1/environments/{environmentID}/languages/{languageID}/status | CREATE Language Localization Status
[**DeleteLanguageLocalizationStatus**](LanguageLocalizationStatusApi.md#DeleteLanguageLocalizationStatus) | **Delete** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | DELETE Language Localization Status
[**ReadLanguageLocalizationStatus**](LanguageLocalizationStatusApi.md#ReadLanguageLocalizationStatus) | **Get** /v1/environments/{environmentID}/languages/{languageID}/status | READ Language Localization Status
[**ReadOneLanguageLocalizationStatus**](LanguageLocalizationStatusApi.md#ReadOneLanguageLocalizationStatus) | **Get** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | READ One Language Localization Status
[**UpdateLanguageLocalizationStatus**](LanguageLocalizationStatusApi.md#UpdateLanguageLocalizationStatus) | **Put** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | UPDATE Language Localization Status



## CreateLanguageLocalizationStatus

> LanguageLocalizationStatus CreateLanguageLocalizationStatus(ctx, environmentID, languageID).LanguageLocalizationStatus(languageLocalizationStatus).Execute()

CREATE Language Localization Status

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
    languageID := "languageID_example" // string | 
    languageLocalizationStatus := *openapiclient.NewLanguageLocalizationStatus("Service_example") // LanguageLocalizationStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.CreateLanguageLocalizationStatus(context.Background(), environmentID, languageID).LanguageLocalizationStatus(languageLocalizationStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.CreateLanguageLocalizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLanguageLocalizationStatus`: LanguageLocalizationStatus
    fmt.Fprintf(os.Stdout, "Response from `LanguageLocalizationStatusApi.CreateLanguageLocalizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLanguageLocalizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **languageLocalizationStatus** | [**LanguageLocalizationStatus**](LanguageLocalizationStatus.md) |  | 

### Return type

[**LanguageLocalizationStatus**](LanguageLocalizationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLanguageLocalizationStatus

> DeleteLanguageLocalizationStatus(ctx, environmentID, languageID, l10nStatusID).Execute()

DELETE Language Localization Status

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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.DeleteLanguageLocalizationStatus(context.Background(), environmentID, languageID, l10nStatusID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.DeleteLanguageLocalizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLanguageLocalizationStatusRequest struct via the builder pattern


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


## ReadLanguageLocalizationStatus

> LanguageLocalizationStatus ReadLanguageLocalizationStatus(ctx, environmentID, languageID).Execute()

READ Language Localization Status

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
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.ReadLanguageLocalizationStatus(context.Background(), environmentID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.ReadLanguageLocalizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadLanguageLocalizationStatus`: LanguageLocalizationStatus
    fmt.Fprintf(os.Stdout, "Response from `LanguageLocalizationStatusApi.ReadLanguageLocalizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadLanguageLocalizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LanguageLocalizationStatus**](LanguageLocalizationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneLanguageLocalizationStatus

> LanguageLocalizationStatus ReadOneLanguageLocalizationStatus(ctx, environmentID, languageID, l10nStatusID).Execute()

READ One Language Localization Status

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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.ReadOneLanguageLocalizationStatus(context.Background(), environmentID, languageID, l10nStatusID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.ReadOneLanguageLocalizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneLanguageLocalizationStatus`: LanguageLocalizationStatus
    fmt.Fprintf(os.Stdout, "Response from `LanguageLocalizationStatusApi.ReadOneLanguageLocalizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneLanguageLocalizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**LanguageLocalizationStatus**](LanguageLocalizationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLanguageLocalizationStatus

> LanguageLocalizationStatus UpdateLanguageLocalizationStatus(ctx, environmentID, languageID, l10nStatusID).LanguageLocalizationStatus(languageLocalizationStatus).Execute()

UPDATE Language Localization Status

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
    languageID := "languageID_example" // string | 
    l10nStatusID := "l10nStatusID_example" // string | 
    languageLocalizationStatus := *openapiclient.NewLanguageLocalizationStatus("Service_example") // LanguageLocalizationStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.UpdateLanguageLocalizationStatus(context.Background(), environmentID, languageID, l10nStatusID).LanguageLocalizationStatus(languageLocalizationStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.UpdateLanguageLocalizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLanguageLocalizationStatus`: LanguageLocalizationStatus
    fmt.Fprintf(os.Stdout, "Response from `LanguageLocalizationStatusApi.UpdateLanguageLocalizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**languageID** | **string** |  | 
**l10nStatusID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLanguageLocalizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **languageLocalizationStatus** | [**LanguageLocalizationStatus**](LanguageLocalizationStatus.md) |  | 

### Return type

[**LanguageLocalizationStatus**](LanguageLocalizationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

