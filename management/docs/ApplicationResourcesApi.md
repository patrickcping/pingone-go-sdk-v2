# \ApplicationResourcesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationResource**](ApplicationResourcesApi.md#CreateApplicationResource) | **Post** /environments/{environmentID}/resources/{resourceID}/applicationResources | CREATE Application Resource
[**DeleteApplicationResource**](ApplicationResourcesApi.md#DeleteApplicationResource) | **Delete** /environments/{environmentID}/resources/{resourceID}/applicationResources/{applicationResourceID} | DELETE Application Resource
[**ReadAllApplicationResources**](ApplicationResourcesApi.md#ReadAllApplicationResources) | **Get** /environments/{environmentID}/resources/{resourceID}/applicationResources | READ All Application Resources
[**ReadOneApplicationResource**](ApplicationResourcesApi.md#ReadOneApplicationResource) | **Get** /environments/{environmentID}/resources/{resourceID}/applicationResources/{applicationResourceID} | READ One Application Resource
[**UpdateApplicationResource**](ApplicationResourcesApi.md#UpdateApplicationResource) | **Put** /environments/{environmentID}/resources/{resourceID}/applicationResources/{applicationResourceID} | UPDATE Application Resource



## CreateApplicationResource

> ResourceApplicationResource CreateApplicationResource(ctx, environmentID, resourceID).ResourceApplicationResource(resourceApplicationResource).Execute()

CREATE Application Resource

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
    resourceID := "resourceID_example" // string | 
    resourceApplicationResource := *openapiclient.NewResourceApplicationResource("Name_example") // ResourceApplicationResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.CreateApplicationResource(context.Background(), environmentID, resourceID).ResourceApplicationResource(resourceApplicationResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.CreateApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationResource`: ResourceApplicationResource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.CreateApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceApplicationResource** | [**ResourceApplicationResource**](ResourceApplicationResource.md) |  | 

### Return type

[**ResourceApplicationResource**](ResourceApplicationResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationResource

> DeleteApplicationResource(ctx, environmentID, resourceID, applicationResourceID).Execute()

DELETE Application Resource

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
    resourceID := "resourceID_example" // string | 
    applicationResourceID := "applicationResourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationResourcesApi.DeleteApplicationResource(context.Background(), environmentID, resourceID, applicationResourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.DeleteApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationResourceRequest struct via the builder pattern


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


## ReadAllApplicationResources

READ All Application Resources

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllApplicationResources(ctx, environmentID, resourceID).Execute()

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
    pagedIterator := api.ReadAllApplicationResources(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllApplicationResources``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllApplicationResources`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllApplicationResources`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllApplicationResources(ctx, environmentID, resourceID).ExecuteInitialPage()

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
    resourceID := "resourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.ReadAllApplicationResources(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadAllApplicationResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationResources`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ReadAllApplicationResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationResourcesRequest struct via the builder pattern


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


## ReadOneApplicationResource

> ResourceApplicationResource ReadOneApplicationResource(ctx, environmentID, resourceID, applicationResourceID).Execute()

READ One Application Resource

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
    resourceID := "resourceID_example" // string | 
    applicationResourceID := "applicationResourceID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.ReadOneApplicationResource(context.Background(), environmentID, resourceID, applicationResourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadOneApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationResource`: ResourceApplicationResource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ReadOneApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceApplicationResource**](ResourceApplicationResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationResource

> ResourceApplicationResource UpdateApplicationResource(ctx, environmentID, resourceID, applicationResourceID).ResourceApplicationResource(resourceApplicationResource).Execute()

UPDATE Application Resource

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
    resourceID := "resourceID_example" // string | 
    applicationResourceID := "applicationResourceID_example" // string | 
    resourceApplicationResource := *openapiclient.NewResourceApplicationResource("Name_example") // ResourceApplicationResource |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourcesApi.UpdateApplicationResource(context.Background(), environmentID, resourceID, applicationResourceID).ResourceApplicationResource(resourceApplicationResource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.UpdateApplicationResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationResource`: ResourceApplicationResource
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.UpdateApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceApplicationResource** | [**ResourceApplicationResource**](ResourceApplicationResource.md) |  | 

### Return type

[**ResourceApplicationResource**](ResourceApplicationResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

