# \CustomAdminRolesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomAdminRole**](CustomAdminRolesApi.md#CreateCustomAdminRole) | **Post** /environments/{environmentID}/roles | CREATE Custom Role
[**DeleteCustomAdminRole**](CustomAdminRolesApi.md#DeleteCustomAdminRole) | **Delete** /environments/{environmentID}/roles/{customRoleID} | DELETE Custom Role
[**ReadAllCustomAdminRoles**](CustomAdminRolesApi.md#ReadAllCustomAdminRoles) | **Get** /environments/{environmentID}/roles | READ All Custom Admin Roles
[**ReadOneCustomAdminRole**](CustomAdminRolesApi.md#ReadOneCustomAdminRole) | **Get** /environments/{environmentID}/roles/{customRoleID} | READ One Custom Role
[**UpdateCustomAdminRole**](CustomAdminRolesApi.md#UpdateCustomAdminRole) | **Put** /environments/{environmentID}/roles/{customRoleID} | UPDATE Custom Role



## CreateCustomAdminRole

> CustomAdminRole CreateCustomAdminRole(ctx, environmentID).CustomAdminRole(customAdminRole).Execute()

CREATE Custom Role

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
    customAdminRole := *openapiclient.NewCustomAdminRole([]openapiclient.CustomAdminRoleCanAssignInner{*openapiclient.NewCustomAdminRoleCanAssignInner("Id_example")}, "Name_example", []openapiclient.CustomAdminRolePermissionsInner{*openapiclient.NewCustomAdminRolePermissionsInner("Id_example")}) // CustomAdminRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAdminRolesApi.CreateCustomAdminRole(context.Background(), environmentID).CustomAdminRole(customAdminRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAdminRolesApi.CreateCustomAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomAdminRole`: CustomAdminRole
    fmt.Fprintf(os.Stdout, "Response from `CustomAdminRolesApi.CreateCustomAdminRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAdminRole** | [**CustomAdminRole**](CustomAdminRole.md) |  | 

### Return type

[**CustomAdminRole**](CustomAdminRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomAdminRole

> DeleteCustomAdminRole(ctx, environmentID, customRoleID).Execute()

DELETE Custom Role

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
    customRoleID := "customRoleID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomAdminRolesApi.DeleteCustomAdminRole(context.Background(), environmentID, customRoleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAdminRolesApi.DeleteCustomAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomAdminRoleRequest struct via the builder pattern


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


## ReadAllCustomAdminRoles

READ All Custom Admin Roles

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllCustomAdminRoles(ctx, environmentID).Execute()

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
    pagedIterator := api.ReadAllCustomAdminRoles(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllCustomAdminRoles``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllCustomAdminRoles`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllCustomAdminRoles`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllCustomAdminRoles(ctx, environmentID).ExecuteInitialPage()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAdminRolesApi.ReadAllCustomAdminRoles(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAdminRolesApi.ReadAllCustomAdminRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllCustomAdminRoles`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `CustomAdminRolesApi.ReadAllCustomAdminRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllCustomAdminRolesRequest struct via the builder pattern


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


## ReadOneCustomAdminRole

> CustomAdminRole ReadOneCustomAdminRole(ctx, environmentID, customRoleID).Execute()

READ One Custom Role

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
    customRoleID := "customRoleID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAdminRolesApi.ReadOneCustomAdminRole(context.Background(), environmentID, customRoleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAdminRolesApi.ReadOneCustomAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneCustomAdminRole`: CustomAdminRole
    fmt.Fprintf(os.Stdout, "Response from `CustomAdminRolesApi.ReadOneCustomAdminRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneCustomAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomAdminRole**](CustomAdminRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomAdminRole

> CustomAdminRole UpdateCustomAdminRole(ctx, environmentID, customRoleID).CustomAdminRole(customAdminRole).Execute()

UPDATE Custom Role

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
    customRoleID := "customRoleID_example" // string | 
    customAdminRole := *openapiclient.NewCustomAdminRole([]openapiclient.CustomAdminRoleCanAssignInner{*openapiclient.NewCustomAdminRoleCanAssignInner("Id_example")}, "Name_example", []openapiclient.CustomAdminRolePermissionsInner{*openapiclient.NewCustomAdminRolePermissionsInner("Id_example")}) // CustomAdminRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomAdminRolesApi.UpdateCustomAdminRole(context.Background(), environmentID, customRoleID).CustomAdminRole(customAdminRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomAdminRolesApi.UpdateCustomAdminRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomAdminRole`: CustomAdminRole
    fmt.Fprintf(os.Stdout, "Response from `CustomAdminRolesApi.UpdateCustomAdminRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**customRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomAdminRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customAdminRole** | [**CustomAdminRole**](CustomAdminRole.md) |  | 

### Return type

[**CustomAdminRole**](CustomAdminRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

