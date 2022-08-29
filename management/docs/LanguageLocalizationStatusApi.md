# \LanguageLocalizationStatusApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet**](LanguageLocalizationStatusApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet) | **Get** /v1/environments/{environmentID}/languages/{languageID}/status | READ Language Localization Status
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete**](LanguageLocalizationStatusApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete) | **Delete** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | DELETE Language Localization Status
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet**](LanguageLocalizationStatusApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet) | **Get** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | READ One Language Localization Status
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut**](LanguageLocalizationStatusApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut) | **Put** /v1/environments/{environmentID}/languages/{languageID}/status/{l10nStatusID} | CREATE Language Localization Status
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost**](LanguageLocalizationStatusApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost) | **Post** /v1/environments/{environmentID}/languages/{languageID}/status | CREATE Language Localization Status



## V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet(ctx, environmentID, languageID).Execute()

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
    resp, r, err := apiClient.LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet(context.Background(), environmentID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete(ctx, environmentID, languageID, l10nStatusID).Execute()

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
    resp, r, err := apiClient.LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete(context.Background(), environmentID, languageID, l10nStatusID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet(ctx, environmentID, languageID, l10nStatusID).Execute()

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
    resp, r, err := apiClient.LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet(context.Background(), environmentID, languageID, l10nStatusID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut(ctx, environmentID, languageID, l10nStatusID).Body(body).Execute()

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
    l10nStatusID := "l10nStatusID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut(context.Background(), environmentID, languageID, l10nStatusID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusL10nStatusIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost(ctx, environmentID, languageID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost(context.Background(), environmentID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguageLocalizationStatusApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDStatusPostRequest struct via the builder pattern


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

