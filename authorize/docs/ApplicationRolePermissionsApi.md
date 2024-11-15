# \ApplicationRolePermissionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationRolePermission**](ApplicationRolePermissionsApi.md#CreateApplicationRolePermission) | **Post** /environments/{environmentID}/applicationRoles/{applicationRoleID}/permissions | CREATE Application Role Permission
[**DeleteApplicationRolePermission**](ApplicationRolePermissionsApi.md#DeleteApplicationRolePermission) | **Delete** /environments/{environmentID}/applicationRoles/{applicationRoleID}/permissions/{applicationRolePermissionID} | DELETE Application Role Permission
[**ReadApplicationRolePermissions**](ApplicationRolePermissionsApi.md#ReadApplicationRolePermissions) | **Get** /environments/{environmentID}/applicationRoles/{applicationRoleID}/permissions | READ Application Role Permissions



## CreateApplicationRolePermission

> ApplicationRolePermission CreateApplicationRolePermission(ctx, environmentID, applicationRoleID).ApplicationRolePermission(applicationRolePermission).Execute()

CREATE Application Role Permission

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
    applicationRoleID := "applicationRoleID_example" // string | 
    applicationRolePermission := *openapiclient.NewApplicationRolePermission("Id_example") // ApplicationRolePermission |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRolePermissionsApi.CreateApplicationRolePermission(context.Background(), environmentID, applicationRoleID).ApplicationRolePermission(applicationRolePermission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRolePermissionsApi.CreateApplicationRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationRolePermission`: ApplicationRolePermission
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRolePermissionsApi.CreateApplicationRolePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRolePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationRolePermission** | [**ApplicationRolePermission**](ApplicationRolePermission.md) |  | 

### Return type

[**ApplicationRolePermission**](ApplicationRolePermission.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationRolePermission

> DeleteApplicationRolePermission(ctx, environmentID, applicationRoleID, applicationRolePermissionID).Execute()

DELETE Application Role Permission

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
    applicationRoleID := "applicationRoleID_example" // string | 
    applicationRolePermissionID := "applicationRolePermissionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationRolePermissionsApi.DeleteApplicationRolePermission(context.Background(), environmentID, applicationRoleID, applicationRolePermissionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRolePermissionsApi.DeleteApplicationRolePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationRoleID** | **string** |  | 
**applicationRolePermissionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRolePermissionRequest struct via the builder pattern


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


## ReadApplicationRolePermissions

READ Application Role Permissions

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadApplicationRolePermissions(ctx, environmentID, applicationRoleID).Execute()

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
    pagedIterator := api.ReadApplicationRolePermissions(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadApplicationRolePermissions``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadApplicationRolePermissions`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadApplicationRolePermissions`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadApplicationRolePermissions(ctx, environmentID, applicationRoleID).ExecuteInitialPage()

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
    applicationRoleID := "applicationRoleID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRolePermissionsApi.ReadApplicationRolePermissions(context.Background(), environmentID, applicationRoleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRolePermissionsApi.ReadApplicationRolePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApplicationRolePermissions`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRolePermissionsApi.ReadApplicationRolePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationRolePermissionsRequest struct via the builder pattern


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

