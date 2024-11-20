# \ApplicationResourcePermissionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationPermission**](ApplicationResourcePermissionsApi.md#CreateApplicationPermission) | **Post** /environments/{environmentID}/applicationResources/{applicationResourceID}/permissions | CREATE Application Permission
[**DeleteApplicationPermission**](ApplicationResourcePermissionsApi.md#DeleteApplicationPermission) | **Delete** /environments/{environmentID}/applicationResources/{applicationResourceID}/permissions/{applicationResourcePermissionID} | DELETE Application Permission
[**ReadApplicationPermissions**](ApplicationResourcePermissionsApi.md#ReadApplicationPermissions) | **Get** /environments/{environmentID}/applicationResources/{applicationResourceID}/permissions | READ Application Permissions
[**ReadOneApplicationPermission**](ApplicationResourcePermissionsApi.md#ReadOneApplicationPermission) | **Get** /environments/{environmentID}/applicationResources/{applicationResourceID}/permissions/{applicationResourcePermissionID} | READ One Application Permission
[**UpdateApplicationPermission**](ApplicationResourcePermissionsApi.md#UpdateApplicationPermission) | **Put** /environments/{environmentID}/applicationResources/{applicationResourceID}/permissions/{applicationResourcePermissionID} | UPDATE Application Permission



## CreateApplicationPermission

> ApplicationResourcePermission CreateApplicationPermission(ctx, environmentID, applicationResourceID).ApplicationResourcePermission(applicationResourcePermission).Execute()

CREATE Application Permission

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
	applicationResourcePermission := *openapiclient.NewApplicationResourcePermission("Action_example") // ApplicationResourcePermission |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcePermissionsApi.CreateApplicationPermission(context.Background(), environmentID, applicationResourceID).ApplicationResourcePermission(applicationResourcePermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.CreateApplicationPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationPermission`: ApplicationResourcePermission
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcePermissionsApi.CreateApplicationPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationResourcePermission** | [**ApplicationResourcePermission**](ApplicationResourcePermission.md) |  | 

### Return type

[**ApplicationResourcePermission**](ApplicationResourcePermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationPermission

> DeleteApplicationPermission(ctx, environmentID, applicationResourceID, applicationResourcePermissionID).Execute()

DELETE Application Permission

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
	applicationResourcePermissionID := "applicationResourcePermissionID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplicationResourcePermissionsApi.DeleteApplicationPermission(context.Background(), environmentID, applicationResourceID, applicationResourcePermissionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.DeleteApplicationPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 
**applicationResourcePermissionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationPermissionRequest struct via the builder pattern


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


## ReadApplicationPermissions

READ Application Permissions

### Paged Response (Recommended)

> PagedIterator[EntityArray] ReadApplicationPermissions(ctx, environmentID, applicationResourceID).Execute()

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
	applicationResourceID := "applicationResourceID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.ApplicationResourcePermissionsApi.ReadApplicationPermissions(context.Background(), environmentID, applicationResourceID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.ReadApplicationPermissions``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadApplicationPermissions` page iteration: EntityArray
		fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcePermissionsApi.ReadApplicationPermissions` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> EntityArray ReadApplicationPermissions(ctx, environmentID, applicationResourceID).ExecuteInitialPage()

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
	applicationResourceID := "applicationResourceID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcePermissionsApi.ReadApplicationPermissions(context.Background(), environmentID, applicationResourceID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.ReadApplicationPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApplicationPermissions`: EntityArray
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcePermissionsApi.ReadApplicationPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationPermissionsRequest struct via the builder pattern


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


## ReadOneApplicationPermission

> ApplicationResourcePermission ReadOneApplicationPermission(ctx, environmentID, applicationResourceID, applicationResourcePermissionID).Execute()

READ One Application Permission

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
	applicationResourcePermissionID := "applicationResourcePermissionID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcePermissionsApi.ReadOneApplicationPermission(context.Background(), environmentID, applicationResourceID, applicationResourcePermissionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.ReadOneApplicationPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneApplicationPermission`: ApplicationResourcePermission
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcePermissionsApi.ReadOneApplicationPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 
**applicationResourcePermissionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationResourcePermission**](ApplicationResourcePermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPermission

> ApplicationResourcePermission UpdateApplicationPermission(ctx, environmentID, applicationResourceID, applicationResourcePermissionID).ApplicationResourcePermission(applicationResourcePermission).Execute()

UPDATE Application Permission

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
	applicationResourcePermissionID := "applicationResourcePermissionID_example" // string | 
	applicationResourcePermission := *openapiclient.NewApplicationResourcePermission("Action_example") // ApplicationResourcePermission |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationResourcePermissionsApi.UpdateApplicationPermission(context.Background(), environmentID, applicationResourceID, applicationResourcePermissionID).ApplicationResourcePermission(applicationResourcePermission).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourcePermissionsApi.UpdateApplicationPermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationPermission`: ApplicationResourcePermission
	fmt.Fprintf(os.Stdout, "Response from `ApplicationResourcePermissionsApi.UpdateApplicationPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationResourceID** | **string** |  | 
**applicationResourcePermissionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationResourcePermission** | [**ApplicationResourcePermission**](ApplicationResourcePermission.md) |  | 

### Return type

[**ApplicationResourcePermission**](ApplicationResourcePermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

