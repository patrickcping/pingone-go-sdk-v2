# \IntegrationCatalogApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadOneIntegrationVersionAsset**](IntegrationCatalogApi.md#DownloadOneIntegrationVersionAsset) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID}/asset | Download Integration Version Asset
[**ReadAllIntegrationMetadata**](IntegrationCatalogApi.md#ReadAllIntegrationMetadata) | **Get** /environments/{environmentID}/integrations | READ All Integration Metadata
[**ReadAllIntegrationVersionAttributes**](IntegrationCatalogApi.md#ReadAllIntegrationVersionAttributes) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID}/attributes | Read All Attributes of an Integration Version (SAML only)
[**ReadIntegrationVersionMetadata**](IntegrationCatalogApi.md#ReadIntegrationVersionMetadata) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions | Read All Integration Versions Metadata
[**ReadOneIntegrationMetadata**](IntegrationCatalogApi.md#ReadOneIntegrationMetadata) | **Get** /environments/{environmentID}/integrations/{integrationID} | READ One Integration Metadata
[**ReadOneIntegrationVersionAttributes**](IntegrationCatalogApi.md#ReadOneIntegrationVersionAttributes) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID}/attributes/{integrationVersionAttributeID} | Read One Attributes of an Integration Version (SAML only)
[**ReadOneIntegrationVersionMetadata**](IntegrationCatalogApi.md#ReadOneIntegrationVersionMetadata) | **Get** /environments/{environmentID}/integrations/{integrationID}/versions/{integrationVersionID} | READ One Integration Version Metadata



## DownloadOneIntegrationVersionAsset

> string DownloadOneIntegrationVersionAsset(ctx, environmentID, integrationID, integrationVersionID).Execute()

Download Integration Version Asset

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.DownloadOneIntegrationVersionAsset(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.DownloadOneIntegrationVersionAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadOneIntegrationVersionAsset`: string
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.DownloadOneIntegrationVersionAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadOneIntegrationVersionAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/zip, application/x-zip-compressed

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllIntegrationMetadata

> EntityArrayPagedIterator ReadAllIntegrationMetadata(ctx, environmentID).Filter(filter).Expand(expand).Execute()

READ All Integration Metadata

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
    filter := "(pingProductNames eq "PINGFEDERATE" and version.releasedOn ge "2020-05-20")" // string |  (optional)
    expand := "versions" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadAllIntegrationMetadata(context.Background(), environmentID).Filter(filter).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadAllIntegrationMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllIntegrationMetadata`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadAllIntegrationMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllIntegrationMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 
 **expand** | **string** |  | 

### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllIntegrationVersionAttributes

> EntityArrayPagedIterator ReadAllIntegrationVersionAttributes(ctx, environmentID, integrationID, integrationVersionID).Execute()

Read All Attributes of an Integration Version (SAML only)

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadAllIntegrationVersionAttributes(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadAllIntegrationVersionAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllIntegrationVersionAttributes`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadAllIntegrationVersionAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllIntegrationVersionAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIntegrationVersionMetadata

Read All Integration Versions Metadata

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadIntegrationVersionMetadata(ctx, environmentID, integrationID).Execute()

#### Example

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadIntegrationVersionMetadata(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadIntegrationVersionMetadata``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadIntegrationVersionMetadata`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadIntegrationVersionMetadata`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadIntegrationVersionMetadata(ctx, environmentID, integrationID).ExecuteInitialPage()

#### Example

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadIntegrationVersionMetadata(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadIntegrationVersionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIntegrationVersionMetadata`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadIntegrationVersionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadIntegrationVersionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneIntegrationMetadata

> Integration ReadOneIntegrationMetadata(ctx, environmentID, integrationID).Execute()

READ One Integration Metadata

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
    integrationID := "integrationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadOneIntegrationMetadata(context.Background(), environmentID, integrationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadOneIntegrationMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneIntegrationMetadata`: Integration
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadOneIntegrationMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneIntegrationMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Integration**](Integration.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneIntegrationVersionAttributes

> IntegrationVersionAttribute ReadOneIntegrationVersionAttributes(ctx, environmentID, integrationID, integrationVersionID, integrationVersionAttributeID).Execute()

Read One Attributes of an Integration Version (SAML only)

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 
    integrationVersionAttributeID := "integrationVersionAttributeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadOneIntegrationVersionAttributes(context.Background(), environmentID, integrationID, integrationVersionID, integrationVersionAttributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadOneIntegrationVersionAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneIntegrationVersionAttributes`: IntegrationVersionAttribute
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadOneIntegrationVersionAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 
**integrationVersionAttributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneIntegrationVersionAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**IntegrationVersionAttribute**](IntegrationVersionAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneIntegrationVersionMetadata

> IntegrationVersion ReadOneIntegrationVersionMetadata(ctx, environmentID, integrationID, integrationVersionID).Execute()

READ One Integration Version Metadata

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
    integrationID := "integrationID_example" // string | 
    integrationVersionID := "integrationVersionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IntegrationCatalogApi.ReadOneIntegrationVersionMetadata(context.Background(), environmentID, integrationID, integrationVersionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IntegrationCatalogApi.ReadOneIntegrationVersionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneIntegrationVersionMetadata`: IntegrationVersion
    fmt.Fprintf(os.Stdout, "Response from `IntegrationCatalogApi.ReadOneIntegrationVersionMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**integrationID** | **string** |  | 
**integrationVersionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneIntegrationVersionMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IntegrationVersion**](IntegrationVersion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

