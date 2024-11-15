# \RolesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllRoles**](RolesApi.md#ReadAllRoles) | **Get** /roles | READ All Roles
[**ReadOneRole**](RolesApi.md#ReadOneRole) | **Get** /roles/{roleID} | READ One Role



## ReadAllRoles

READ All Roles

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllRoles(ctx).Execute()

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
    pagedIterator := api.ReadAllRoles(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllRoles``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllRoles`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllRoles`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllRoles(ctx).ExecuteInitialPage()

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ReadAllRoles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ReadAllRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllRoles`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ReadAllRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllRolesRequest struct via the builder pattern


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


## ReadOneRole

> Role ReadOneRole(ctx, roleID).Execute()

READ One Role

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
    roleID := "roleID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.ReadOneRole(context.Background(), roleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ReadOneRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRole`: Role
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ReadOneRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Role**](Role.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

