# \AgreementRevisionsResourcesApi

All URIs are relative to *https://agreement-mgmt.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadOneAgreementLanguageRevision**](AgreementRevisionsResourcesApi.md#ReadOneAgreementLanguageRevision) | **Get** /environments/{environmentID}/agreements/{agreementID}/languages/{languageID}/revisions/{revisionID}.json | READ One Agreement Language Revision



## ReadOneAgreementLanguageRevision

> AgreementRevisionText ReadOneAgreementLanguageRevision(ctx, environmentID, agreementID, languageID, revisionID).Execute()

READ One Agreement Language Revision

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/agreementmanagement"
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
    // response from `ReadOneAgreementLanguageRevision`: AgreementRevisionText
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

[**AgreementRevisionText**](AgreementRevisionText.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

