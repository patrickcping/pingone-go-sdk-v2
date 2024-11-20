# \ApplicationResourcesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApplicationResources**](ApplicationResourcesApi.md#ReadApplicationResources) | **Get** /environments/{environmentID}/applicationResources | READ Application Resources
[**ReadOneApplicationResource**](ApplicationResourcesApi.md#ReadOneApplicationResource) | **Get** /environments/{environmentID}/applicationResources/{applicationResourceID} | READ One Application Resource



## ReadApplicationResources

READ Application Resources

### Paged Response (Recommended)

> PagedIterator[EntityArray] ReadApplicationResources(ctx, environmentID).Execute()

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
	pagedIterator := apiClient.ApplicationResourcesApi.ReadApplicationResources(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadApplicationResources``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadApplicationResources` page iteration: EntityArray
		fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcesApi.ReadApplicationResources` page iteration: %v\n", pageCursor.Data)
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
	resp, r, err := apiClient.ApplicationResourcesApi.ReadApplicationResources(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcesApi.ReadApplicationResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApplicationResources`: EntityArray
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

Page Iterator: PagedIterator[[**EntityArray**](EntityArray.md)]

PagedIterator[EntityArray] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**EntityArray**](EntityArray.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**EntityArray**](EntityArray.md)

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

