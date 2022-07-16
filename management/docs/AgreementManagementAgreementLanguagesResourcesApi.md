# \AgreementManagementAgreementLanguagesResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet**](AgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages | READ All Languages
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete**](AgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | DELETE Language
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet**](AgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | READ One Language
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut**](AgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut) | **Put** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID} | UPDATE Language
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost**](AgreementManagementAgreementLanguagesResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost) | **Post** /v1/environments/{environmentID}/agreements/{agreementID}/languages | CREATE Language



## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet(ctx, environmentID, agreementID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete(ctx, environmentID, agreementID, languageID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet(ctx, environmentID, agreementID, languageID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut(ctx, environmentID, agreementID, languageID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut(context.Background(), environmentID, agreementID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost(ctx, environmentID, agreementID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost(context.Background(), environmentID, agreementID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementLanguagesResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPost``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesPostRequest struct via the builder pattern


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

