# \AgreementManagementAgreementRevisionsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet**](AgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | READ All Revisions
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost**](AgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost) | **Post** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | CREATE Revision
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete**](AgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | DELETE Revision
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet**](AgreementManagementAgreementRevisionsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | READ One Revision



## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet(ctx, environmentID, agreementID, languageID).Execute()

READ All Revisions

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
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost(ctx, environmentID, agreementID, languageID).Body(body).Execute()

CREATE Revision

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
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost(context.Background(), environmentID, agreementID, languageID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete(ctx, environmentID, agreementID, languageID, revisionID).Execute()

DELETE Revision

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
    revisionID := "revisionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDelete``: %v\n", err)
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
**revisionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet(ctx, environmentID, agreementID, languageID, revisionID).Execute()

READ One Revision

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
    revisionID := "revisionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGet``: %v\n", err)
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
**revisionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDLanguagesLanguageIDRevisionsRevisionIDGetRequest struct via the builder pattern


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

