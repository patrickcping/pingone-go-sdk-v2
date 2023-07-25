# \PropagationStoreMetadataApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost**](PropagationStoreMetadataApi.md#EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost) | **Post** /environments/{environmentID}/propagation/storeMetadata/Aquera | Identity Propagation Store Metadata (Aquera)
[**EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost**](PropagationStoreMetadataApi.md#EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost) | **Post** /environments/{environmentID}/propagation/storeMetadata/SalesforceContacts | Identity Propagation Store Metadata (SalesforceContacts)
[**EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost**](PropagationStoreMetadataApi.md#EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost) | **Post** /environments/{environmentID}/propagation/storeMetadata/Salesforce | Identity Propagation Store Metadata (Salesforce)
[**EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost**](PropagationStoreMetadataApi.md#EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost) | **Post** /environments/{environmentID}/propagation/storeMetadata/scim | Identity Propagation Store Metadata (SCIM)



## EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost

> EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (Aquera)

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost

> EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (SalesforceContacts)

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost

> EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (Salesforce)

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost

> EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (SCIM)

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationStoreMetadataApi.EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationStoreMetadataScimPostRequest struct via the builder pattern


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

