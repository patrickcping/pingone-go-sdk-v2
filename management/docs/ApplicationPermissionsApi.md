# \ApplicationPermissionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllApplicationPermissions**](ApplicationPermissionsApi.md#ReadAllApplicationPermissions) | **Get** /environments/{environmentID}/resources/{resourceID}/applicationPermissions | READ All Application Permissions



## ReadAllApplicationPermissions

READ All Application Permissions

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllApplicationPermissions(ctx, environmentID, resourceID).Execute()

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
    pagedIterator := api.ReadAllApplicationPermissions(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllApplicationPermissions``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllApplicationPermissions`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllApplicationPermissions`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllApplicationPermissions(ctx, environmentID, resourceID).ExecuteInitialPage()

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
    resp, r, err := apiClient.ApplicationPermissionsApi.ReadAllApplicationPermissions(context.Background(), environmentID, resourceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationPermissionsApi.ReadAllApplicationPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationPermissions`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `ApplicationPermissionsApi.ReadAllApplicationPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**resourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationPermissionsRequest struct via the builder pattern


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

