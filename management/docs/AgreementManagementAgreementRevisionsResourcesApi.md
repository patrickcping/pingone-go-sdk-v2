# \AgreementManagementAgreementRevisionsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgreementLanguageRevision**](AgreementManagementAgreementRevisionsResourcesApi.md#CreateAgreementLanguageRevision) | **Post** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | CREATE Revision
[**DeleteAgreementLanguageRevision**](AgreementManagementAgreementRevisionsResourcesApi.md#DeleteAgreementLanguageRevision) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | DELETE Revision
[**ReadAllAgreementLanguageRevisions**](AgreementManagementAgreementRevisionsResourcesApi.md#ReadAllAgreementLanguageRevisions) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | READ All Revisions
[**ReadOneAgreementLanguageRevision**](AgreementManagementAgreementRevisionsResourcesApi.md#ReadOneAgreementLanguageRevision) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | READ One Revision



## CreateAgreementLanguageRevision

> AgreementLanguageRevision CreateAgreementLanguageRevision(ctx, environmentID, agreementID, languageID).AgreementLanguageRevision(agreementLanguageRevision).Execute()

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
    agreementLanguageRevision := *openapiclient.NewAgreementLanguageRevision() // AgreementLanguageRevision |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.CreateAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID).AgreementLanguageRevision(agreementLanguageRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.CreateAgreementLanguageRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgreementLanguageRevision`: AgreementLanguageRevision
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementRevisionsResourcesApi.CreateAgreementLanguageRevision`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiCreateAgreementLanguageRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **agreementLanguageRevision** | [**AgreementLanguageRevision**](AgreementLanguageRevision.md) |  | 

### Return type

[**AgreementLanguageRevision**](AgreementLanguageRevision.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgreementLanguageRevision

> DeleteAgreementLanguageRevision(ctx, environmentID, agreementID, languageID, revisionID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.DeleteAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.DeleteAgreementLanguageRevision``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteAgreementLanguageRevisionRequest struct via the builder pattern


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


## ReadAllAgreementLanguageRevisions

> EntityArray ReadAllAgreementLanguageRevisions(ctx, environmentID, agreementID, languageID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAgreementLanguageRevisions`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadAllAgreementLanguageRevisionsRequest struct via the builder pattern


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


## ReadOneAgreementLanguageRevision

> AgreementLanguageRevision ReadOneAgreementLanguageRevision(ctx, environmentID, agreementID, languageID, revisionID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAgreementLanguageRevision`: AgreementLanguageRevision
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiReadOneAgreementLanguageRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**AgreementLanguageRevision**](AgreementLanguageRevision.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

