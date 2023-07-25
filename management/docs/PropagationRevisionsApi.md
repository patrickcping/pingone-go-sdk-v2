# \PropagationRevisionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet**](PropagationRevisionsApi.md#EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet) | **Get** /environments/{environmentID}/propagation/revisions/id:latest | READ Latest Revision
[**EnvironmentsEnvironmentIDPropagationRevisionsPost**](PropagationRevisionsApi.md#EnvironmentsEnvironmentIDPropagationRevisionsPost) | **Post** /environments/{environmentID}/propagation/revisions | CREATE Revision
[**EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet**](PropagationRevisionsApi.md#EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet) | **Get** /environments/{environmentID}/propagation/revisions/{previousRevisionID} | READ Previous Revision



## EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet

> EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet(ctx, environmentID).Accept(accept).Execute()

READ Latest Revision

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
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet(context.Background(), environmentID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsIdlatestGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRevisionsIdlatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

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


## EnvironmentsEnvironmentIDPropagationRevisionsPost

> EnvironmentsEnvironmentIDPropagationRevisionsPost(ctx, environmentID).Execute()

CREATE Revision

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
    r, err := apiClient.PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsPost(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRevisionsPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet

> EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet(ctx, environmentID, previousRevisionID).Accept(accept).Execute()

READ Previous Revision

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
    previousRevisionID := "previousRevisionID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet(context.Background(), environmentID, previousRevisionID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRevisionsApi.EnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**previousRevisionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRevisionsPreviousRevisionIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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

