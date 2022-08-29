# \LanguagesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDLanguagesGet**](LanguagesApi.md#V1EnvironmentsEnvironmentIDLanguagesGet) | **Get** /v1/environments/{environmentID}/languages | READ Languages
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete**](LanguagesApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete) | **Delete** /v1/environments/{environmentID}/languages/{languageID} | DELETE Language
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet**](LanguagesApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet) | **Get** /v1/environments/{environmentID}/languages/{languageID} | READ One Language
[**V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut**](LanguagesApi.md#V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut) | **Put** /v1/environments/{environmentID}/languages/{languageID} | UPDATE Language 
[**V1EnvironmentsEnvironmentIDLanguagesPost**](LanguagesApi.md#V1EnvironmentsEnvironmentIDLanguagesPost) | **Post** /v1/environments/{environmentID}/languages/ | CREATE Language



## V1EnvironmentsEnvironmentIDLanguagesGet

> V1EnvironmentsEnvironmentIDLanguagesGet(ctx, environmentID).Execute()

READ Languages

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
    resp, r, err := apiClient.LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete(ctx, environmentID, languageID).Execute()

DELETE Language

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
    resp, r, err := apiClient.LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete(context.Background(), environmentID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet(ctx, environmentID, languageID).Execute()

READ One Language

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
    resp, r, err := apiClient.LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet(context.Background(), environmentID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut

> V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut(ctx, environmentID, languageID).Body(body).Execute()

UPDATE Language 

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
    resp, r, err := apiClient.LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut(context.Background(), environmentID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesLanguageIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesLanguageIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDLanguagesPost

> V1EnvironmentsEnvironmentIDLanguagesPost(ctx, environmentID).Body(body).Execute()

CREATE Language

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanguagesApi.V1EnvironmentsEnvironmentIDLanguagesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDLanguagesPostRequest struct via the builder pattern


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

