# \IdentityPropagationProvisioningPropagationStoreMetadataApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost**](IdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost) | **Post** /v1/environments/{environmentID}/propagation/storeMetadata/Aquera | Identity Propagation Store Metadata (Aquera)
[**V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost**](IdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost) | **Post** /v1/environments/{environmentID}/propagation/storeMetadata/SalesforceContacts | Identity Propagation Store Metadata (SalesforceContacts)
[**V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost**](IdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost) | **Post** /v1/environments/{environmentID}/propagation/storeMetadata/Salesforce | Identity Propagation Store Metadata (Salesforce)
[**V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost**](IdentityPropagationProvisioningPropagationStoreMetadataApi.md#V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost) | **Post** /v1/environments/{environmentID}/propagation/storeMetadata/scim | Identity Propagation Store Metadata (SCIM)



## V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost

> V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (Aquera)

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
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoreMetadataAqueraPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost

> V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (SalesforceContacts)

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
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforceContactsPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost

> V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (Salesforce)

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
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoreMetadataSalesforcePostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost

> V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost(ctx, environmentID).Body(body).Execute()

Identity Propagation Store Metadata (SCIM)

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
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationStoreMetadataApi.V1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationStoreMetadataScimPostRequest struct via the builder pattern


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

