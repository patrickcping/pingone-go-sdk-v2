# \ApplicationResourcesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApplicationResources**](ApplicationResourcesApi.md#ReadApplicationResources) | **Get** /environments/{environmentID}/applicationResources | READ Application Resources
[**ReadOneApplicationResource**](ApplicationResourcesApi.md#ReadOneApplicationResource) | **Get** /environments/{environmentID}/applicationResources/{applicationResourceID} | READ One Application Resource



## ReadApplicationResources

READ Application Resources

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadApplicationResources(ctx, environmentID).Execute()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
    environmentID := "environmentID_example" // string | 
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadApplicationResources(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadApplicationResources``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadApplicationResources`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadApplicationResources`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadApplicationResources(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcesApi.ReadApplicationResources(context.Background(), environmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadApplicationResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApplicationResources`: EntityArrayPagedIterator
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ReadApplicationResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationResourcesRequest struct via the builder pattern


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

> ApplicationResource ReadOneApplicationResource(ctx, environmentID, applicationResourceID).Execute()

READ One Application Resource

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
	environmentID := "environmentID_example" // string | 
	applicationResourceID := "applicationResourceID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcesApi.ReadOneApplicationResource(context.Background(), environmentID, applicationResourceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadOneApplicationResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneApplicationResource`: ApplicationResource
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ReadOneApplicationResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationResource**](ApplicationResource.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

