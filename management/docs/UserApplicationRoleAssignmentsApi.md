# \UserApplicationRoleAssignmentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserApplicationRoleAssignment**](UserApplicationRoleAssignmentsApi.md#CreateUserApplicationRoleAssignment) | **Post** /environments/{environmentID}/users/{userID}/applicationRoles | CREATE User Application Role Assignment
[**DeleteUserApplicationRoleAssignment**](UserApplicationRoleAssignmentsApi.md#DeleteUserApplicationRoleAssignment) | **Delete** /environments/{environmentID}/users/{userID}/applicationRoles/{applicationRoleID} | DELETE User Application Role Assignment
[**ReadUserApplicationRoleAssignments**](UserApplicationRoleAssignmentsApi.md#ReadUserApplicationRoleAssignments) | **Get** /environments/{environmentID}/users/{userID}/applicationRoles | READ User Application Role Assignments



## CreateUserApplicationRoleAssignment

> UserApplicationRoleAssignment CreateUserApplicationRoleAssignment(ctx, environmentID, userID).UserApplicationRoleAssignment(userApplicationRoleAssignment).Execute()

CREATE User Application Role Assignment

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
    userID := "userID_example" // string | 
    userApplicationRoleAssignment := *openapiclient.NewUserApplicationRoleAssignment("Id_example") // UserApplicationRoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApplicationRoleAssignmentsApi.CreateUserApplicationRoleAssignment(context.Background(), environmentID, userID).UserApplicationRoleAssignment(userApplicationRoleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationRoleAssignmentsApi.CreateUserApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserApplicationRoleAssignment`: UserApplicationRoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UserApplicationRoleAssignmentsApi.CreateUserApplicationRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserApplicationRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userApplicationRoleAssignment** | [**UserApplicationRoleAssignment**](UserApplicationRoleAssignment.md) |  | 

### Return type

[**UserApplicationRoleAssignment**](UserApplicationRoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserApplicationRoleAssignment

> DeleteUserApplicationRoleAssignment(ctx, environmentID, userID, applicationRoleID).Execute()

DELETE User Application Role Assignment

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
    userID := "userID_example" // string | 
    applicationRoleID := "applicationRoleID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserApplicationRoleAssignmentsApi.DeleteUserApplicationRoleAssignment(context.Background(), environmentID, userID, applicationRoleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationRoleAssignmentsApi.DeleteUserApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**applicationRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserApplicationRoleAssignmentRequest struct via the builder pattern


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


## ReadUserApplicationRoleAssignments

READ User Application Role Assignments

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadUserApplicationRoleAssignments(ctx, environmentID, userID).Execute()

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
    pagedIterator := api.ReadUserApplicationRoleAssignments(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadUserApplicationRoleAssignments``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadUserApplicationRoleAssignments`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadUserApplicationRoleAssignments`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadUserApplicationRoleAssignments(ctx, environmentID, userID).ExecuteInitialPage()

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApplicationRoleAssignmentsApi.ReadUserApplicationRoleAssignments(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationRoleAssignmentsApi.ReadUserApplicationRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserApplicationRoleAssignments`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `UserApplicationRoleAssignmentsApi.ReadUserApplicationRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserApplicationRoleAssignmentsRequest struct via the builder pattern


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

