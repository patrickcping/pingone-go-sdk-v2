# \ResourceAttributesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResourceAttribute**](ResourceAttributesApi.md#CreateResourceAttribute) | **Post** /environments/{environmentID}/resources/{resourceID}/attributes | CREATE Resource Attribute
[**DeleteResourceAttribute**](ResourceAttributesApi.md#DeleteResourceAttribute) | **Delete** /environments/{environmentID}/resources/{resourceID}/attributes/{resourceAttrID} | DELETE Resource Attribute
[**ReadAllResourceAttributes**](ResourceAttributesApi.md#ReadAllResourceAttributes) | **Get** /environments/{environmentID}/resources/{resourceID}/attributes | READ All Resource Attributes
[**ReadOneResourceAttribute**](ResourceAttributesApi.md#ReadOneResourceAttribute) | **Get** /environments/{environmentID}/resources/{resourceID}/attributes/{resourceAttrID} | READ One Resource Attribute
[**UpdateResourceAttribute**](ResourceAttributesApi.md#UpdateResourceAttribute) | **Put** /environments/{environmentID}/resources/{resourceID}/attributes/{resourceAttrID} | UPDATE Resource Attribute



## CreateResourceAttribute

> ResourceAttribute CreateResourceAttribute(ctx, environmentID, resourceID).ResourceAttribute(resourceAttribute).Execute()

CREATE Resource Attribute

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
    resourceAttribute := *openapiclient.NewResourceAttribute("Name_example", "Value_example") // ResourceAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceAttributesApi.CreateResourceAttribute(context.Background(), environmentID, resourceID).ResourceAttribute(resourceAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAttributesApi.CreateResourceAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateResourceAttribute`: ResourceAttribute
    fmt.Fprintf(os.Stdout, "Response from `ResourceAttributesApi.CreateResourceAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceAttribute** | [**ResourceAttribute**](ResourceAttribute.md) |  | 

### Return type

[**ResourceAttribute**](ResourceAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceAttribute

> DeleteResourceAttribute(ctx, environmentID, resourceID, resourceAttrID).Execute()

DELETE Resource Attribute

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
    resourceAttrID := "resourceAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ResourceAttributesApi.DeleteResourceAttribute(context.Background(), environmentID, resourceID, resourceAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAttributesApi.DeleteResourceAttribute``: %v\n", err)
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
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceAttributeRequest struct via the builder pattern


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


## ReadAllResourceAttributes

READ All Resource Attributes

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllResourceAttributes(ctx, environmentID, resourceID).Execute()

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
    pagedIterator := api.ReadAllResourceAttributes(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllResourceAttributes``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllResourceAttributes`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllResourceAttributes`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllResourceAttributes(ctx, environmentID, resourceID).ExecuteInitialPage()

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
    resp, r, err := apiClient.ResourceAttributesApi.ReadAllResourceAttributes(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAttributesApi.ReadAllResourceAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllResourceAttributes`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `ResourceAttributesApi.ReadAllResourceAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllResourceAttributesRequest struct via the builder pattern


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


## ReadOneResourceAttribute

> ResourceAttribute ReadOneResourceAttribute(ctx, environmentID, resourceID, resourceAttrID).Execute()

READ One Resource Attribute

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
    resourceAttrID := "resourceAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceAttributesApi.ReadOneResourceAttribute(context.Background(), environmentID, resourceID, resourceAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAttributesApi.ReadOneResourceAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneResourceAttribute`: ResourceAttribute
    fmt.Fprintf(os.Stdout, "Response from `ResourceAttributesApi.ReadOneResourceAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneResourceAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ResourceAttribute**](ResourceAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceAttribute

> ResourceAttribute UpdateResourceAttribute(ctx, environmentID, resourceID, resourceAttrID).ResourceAttribute(resourceAttribute).Execute()

UPDATE Resource Attribute

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
    resourceAttrID := "resourceAttrID_example" // string | 
    resourceAttribute := *openapiclient.NewResourceAttribute("Name_example", "Value_example") // ResourceAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceAttributesApi.UpdateResourceAttribute(context.Background(), environmentID, resourceID, resourceAttrID).ResourceAttribute(resourceAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceAttributesApi.UpdateResourceAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceAttribute`: ResourceAttribute
    fmt.Fprintf(os.Stdout, "Response from `ResourceAttributesApi.UpdateResourceAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 
**resourceAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceAttribute** | [**ResourceAttribute**](ResourceAttribute.md) |  | 

### Return type

[**ResourceAttribute**](ResourceAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

