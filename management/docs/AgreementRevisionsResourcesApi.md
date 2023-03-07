# \AgreementRevisionsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgreementLanguageRevision**](AgreementRevisionsResourcesApi.md#CreateAgreementLanguageRevision) | **Post** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | CREATE Revision
[**DeleteAgreementLanguageRevision**](AgreementRevisionsResourcesApi.md#DeleteAgreementLanguageRevision) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | DELETE Revision
[**ReadAllAgreementLanguageRevisions**](AgreementRevisionsResourcesApi.md#ReadAllAgreementLanguageRevisions) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions | READ All Revisions
[**ReadOneAgreementLanguageRevision**](AgreementRevisionsResourcesApi.md#ReadOneAgreementLanguageRevision) | **Get** /v1/environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID} | READ One Revision



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
    "time"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 
    agreementLanguageRevision := *openapiclient.NewAgreementLanguageRevision(openapiclient.EnumAgreementRevisionContentType("text/html"), time.Now(), false, "Text_example") // AgreementLanguageRevision |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementRevisionsResourcesApi.CreateAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID).AgreementLanguageRevision(agreementLanguageRevision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementRevisionsResourcesApi.CreateAgreementLanguageRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgreementLanguageRevision`: AgreementLanguageRevision
    fmt.Fprintf(os.Stdout, "Response from `AgreementRevisionsResourcesApi.CreateAgreementLanguageRevision`: %v\n", resp)
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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 
    revisionID := "revisionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AgreementRevisionsResourcesApi.DeleteAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementRevisionsResourcesApi.DeleteAgreementLanguageRevision``: %v\n", err)
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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions(context.Background(), environmentID, agreementID, languageID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAgreementLanguageRevisions`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AgreementRevisionsResourcesApi.ReadAllAgreementLanguageRevisions`: %v\n", resp)
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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    agreementID := "agreementID_example" // string | 
    languageID := "languageID_example" // string | 
    revisionID := "revisionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision(context.Background(), environmentID, agreementID, languageID, revisionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAgreementLanguageRevision`: AgreementLanguageRevision
    fmt.Fprintf(os.Stdout, "Response from `AgreementRevisionsResourcesApi.ReadOneAgreementLanguageRevision`: %v\n", resp)
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

