# \IdentityProviderAttributesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIdentityProviderAttribute**](IdentityProviderAttributesApi.md#CreateIdentityProviderAttribute) | **Post** /environments/{environmentID}/identityProviders/{providerID}/attributes | CREATE Identity Provider Attribute
[**DeleteIdentityProviderAttribute**](IdentityProviderAttributesApi.md#DeleteIdentityProviderAttribute) | **Delete** /environments/{environmentID}/identityProviders/{providerID}/attributes/{providerAttributeID} | DELETE Identity Provider Attribute
[**ReadAllIdentityProviderAttributes**](IdentityProviderAttributesApi.md#ReadAllIdentityProviderAttributes) | **Get** /environments/{environmentID}/identityProviders/{providerID}/attributes | READ All Identity Provider Attributes
[**ReadOneIdentityProviderAttribute**](IdentityProviderAttributesApi.md#ReadOneIdentityProviderAttribute) | **Get** /environments/{environmentID}/identityProviders/{providerID}/attributes/{providerAttributeID} | READ One Identity Provider Attribute
[**UpdateIdentityProviderAttribute**](IdentityProviderAttributesApi.md#UpdateIdentityProviderAttribute) | **Put** /environments/{environmentID}/identityProviders/{providerID}/attributes/{providerAttributeID} | UPDATE Identity Provider Attribute



## CreateIdentityProviderAttribute

> IdentityProviderAttribute CreateIdentityProviderAttribute(ctx, environmentID, providerID).IdentityProviderAttribute(identityProviderAttribute).Execute()

CREATE Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    identityProviderAttribute := *openapiclient.NewIdentityProviderAttribute("Name_example", "Value_example", openapiclient.EnumIdentityProviderAttributeMappingUpdate("EMPTY_ONLY")) // IdentityProviderAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAttributesApi.CreateIdentityProviderAttribute(context.Background(), environmentID, providerID).IdentityProviderAttribute(identityProviderAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAttributesApi.CreateIdentityProviderAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProviderAttribute`: IdentityProviderAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAttributesApi.CreateIdentityProviderAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityProviderAttribute** | [**IdentityProviderAttribute**](IdentityProviderAttribute.md) |  | 

### Return type

[**IdentityProviderAttribute**](IdentityProviderAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIdentityProviderAttribute

> DeleteIdentityProviderAttribute(ctx, environmentID, providerID, providerAttributeID).Execute()

DELETE Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    providerAttributeID := "providerAttributeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityProviderAttributesApi.DeleteIdentityProviderAttribute(context.Background(), environmentID, providerID, providerAttributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAttributesApi.DeleteIdentityProviderAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**providerAttributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderAttributeRequest struct via the builder pattern


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


## ReadAllIdentityProviderAttributes

READ All Identity Provider Attributes

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllIdentityProviderAttributes(ctx, environmentID, providerID).Execute()

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
    pagedIterator := api.ReadAllIdentityProviderAttributes(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllIdentityProviderAttributes``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllIdentityProviderAttributes`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllIdentityProviderAttributes`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllIdentityProviderAttributes(ctx, environmentID, providerID).ExecuteInitialPage()

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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllIdentityProviderAttributes`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAttributesApi.ReadAllIdentityProviderAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllIdentityProviderAttributesRequest struct via the builder pattern


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


## ReadOneIdentityProviderAttribute

> IdentityProviderAttribute ReadOneIdentityProviderAttribute(ctx, environmentID, providerID, providerAttributeID).Execute()

READ One Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    providerAttributeID := "providerAttributeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAttributesApi.ReadOneIdentityProviderAttribute(context.Background(), environmentID, providerID, providerAttributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAttributesApi.ReadOneIdentityProviderAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneIdentityProviderAttribute`: IdentityProviderAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAttributesApi.ReadOneIdentityProviderAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**providerAttributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneIdentityProviderAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**IdentityProviderAttribute**](IdentityProviderAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProviderAttribute

> IdentityProviderAttribute UpdateIdentityProviderAttribute(ctx, environmentID, providerID, providerAttributeID).IdentityProviderAttribute(identityProviderAttribute).Execute()

UPDATE Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    providerAttributeID := "providerAttributeID_example" // string | 
    identityProviderAttribute := *openapiclient.NewIdentityProviderAttribute("Name_example", "Value_example", openapiclient.EnumIdentityProviderAttributeMappingUpdate("EMPTY_ONLY")) // IdentityProviderAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderAttributesApi.UpdateIdentityProviderAttribute(context.Background(), environmentID, providerID, providerAttributeID).IdentityProviderAttribute(identityProviderAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderAttributesApi.UpdateIdentityProviderAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProviderAttribute`: IdentityProviderAttribute
    fmt.Fprintf(os.Stdout, "Response from `IdentityProviderAttributesApi.UpdateIdentityProviderAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**providerAttributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **identityProviderAttribute** | [**IdentityProviderAttribute**](IdentityProviderAttribute.md) |  | 

### Return type

[**IdentityProviderAttribute**](IdentityProviderAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

