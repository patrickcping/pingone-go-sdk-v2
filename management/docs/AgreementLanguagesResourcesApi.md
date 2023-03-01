# \AgreementLanguagesResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgreementLanguage**](AgreementLanguagesResourcesApi.md#CreateAgreementLanguage) | **Post** /v1/environments/{environmentID}/agreements/{agreementID}/languages | CREATE Language
[**DeleteAgreementLanguage**](AgreementLanguagesResourcesApi.md#DeleteAgreementLanguage) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | DELETE Language
[**ReadAllAgreementLanguages**](AgreementLanguagesResourcesApi.md#ReadAllAgreementLanguages) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages | READ All Languages
[**ReadOneAgreementLanguage**](AgreementLanguagesResourcesApi.md#ReadOneAgreementLanguage) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | READ One Language
[**UpdateAgreementLanguage**](AgreementLanguagesResourcesApi.md#UpdateAgreementLanguage) | **Put** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | UPDATE Language



## CreateAgreementLanguage

> AgreementLanguage CreateAgreementLanguage(ctx, environmentID, agreementID).AgreementLanguage(agreementLanguage).Execute()

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
    agreementID := "agreementID_example" // string | 
    agreementLanguage := *openapiclient.NewAgreementLanguage("DisplayName_example", false, "Locale_example") // AgreementLanguage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementLanguagesResourcesApi.CreateAgreementLanguage(context.Background(), environmentID, agreementID).AgreementLanguage(agreementLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementLanguagesResourcesApi.CreateAgreementLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgreementLanguage`: AgreementLanguage
    fmt.Fprintf(os.Stdout, "Response from `AgreementLanguagesResourcesApi.CreateAgreementLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgreementLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agreementLanguage** | [**AgreementLanguage**](AgreementLanguage.md) |  | 

### Return type

[**AgreementLanguage**](AgreementLanguage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgreementLanguage

> DeleteAgreementLanguage(ctx, environmentID, agreementID, languageID).Execute()

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
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementLanguagesResourcesApi.DeleteAgreementLanguage(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementLanguagesResourcesApi.DeleteAgreementLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgreementLanguageRequest struct via the builder pattern


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


## ReadAllAgreementLanguages

> EntityArray ReadAllAgreementLanguages(ctx, environmentID, agreementID).Execute()

READ All Languages

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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementLanguagesResourcesApi.ReadAllAgreementLanguages(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementLanguagesResourcesApi.ReadAllAgreementLanguages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAgreementLanguages`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AgreementLanguagesResourcesApi.ReadAllAgreementLanguages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllAgreementLanguagesRequest struct via the builder pattern


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


## ReadOneAgreementLanguage

> AgreementLanguage ReadOneAgreementLanguage(ctx, environmentID, agreementID, languageID).Execute()

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
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementLanguagesResourcesApi.ReadOneAgreementLanguage(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementLanguagesResourcesApi.ReadOneAgreementLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAgreementLanguage`: AgreementLanguage
    fmt.Fprintf(os.Stdout, "Response from `AgreementLanguagesResourcesApi.ReadOneAgreementLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAgreementLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AgreementLanguage**](AgreementLanguage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgreementLanguage

> AgreementLanguage UpdateAgreementLanguage(ctx, environmentID, agreementID, languageID).AgreementLanguage(agreementLanguage).Execute()

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
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 
    agreementLanguage := *openapiclient.NewAgreementLanguage("DisplayName_example", false, "Locale_example") // AgreementLanguage |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementLanguagesResourcesApi.UpdateAgreementLanguage(context.Background(), environmentID, agreementID, languageID).AgreementLanguage(agreementLanguage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementLanguagesResourcesApi.UpdateAgreementLanguage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgreementLanguage`: AgreementLanguage
    fmt.Fprintf(os.Stdout, "Response from `AgreementLanguagesResourcesApi.UpdateAgreementLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 
**languageID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgreementLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **agreementLanguage** | [**AgreementLanguage**](AgreementLanguage.md) |  | 

### Return type

[**AgreementLanguage**](AgreementLanguage.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

