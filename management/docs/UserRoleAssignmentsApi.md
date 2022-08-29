# \UserRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserRoleAssignment**](UserRoleAssignmentsApi.md#CreateUserRoleAssignment) | **Post** /v1/environments/{environmentID}/users/{userID}/roleAssignments | CREATE User Role Assignment
[**DeleteUserRoleAssignment**](UserRoleAssignmentsApi.md#DeleteUserRoleAssignment) | **Delete** /v1/environments/{environmentID}/users/{userID}/roleAssignments/{roleAssignmentID} | DELETE User&#39;s Role Assignment
[**ReadOneUserRoleAssignment**](UserRoleAssignmentsApi.md#ReadOneUserRoleAssignment) | **Get** /v1/environments/{environmentID}/users/{userID}/roleAssignments/{roleAssignmentID} | READ One Role Assignment
[**ReadUserRoleAssignments**](UserRoleAssignmentsApi.md#ReadUserRoleAssignments) | **Get** /v1/environments/{environmentID}/users/{userID}/roleAssignments | READ Role Assignments



## CreateUserRoleAssignment

> RoleAssignment CreateUserRoleAssignment(ctx, environmentID, userID).RoleAssignment(roleAssignment).Execute()

CREATE User Role Assignment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    roleAssignment := *openapiclient.NewRoleAssignment(*openapiclient.NewRoleAssignmentRole("Id_example"), *openapiclient.NewRoleAssignmentScope("Id_example", openapiclient.EnumRoleAssignmentScopeType("ORGANIZATION"))) // RoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRoleAssignmentsApi.CreateUserRoleAssignment(context.Background(), environmentID, userID).RoleAssignment(roleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsApi.CreateUserRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UserRoleAssignmentsApi.CreateUserRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **roleAssignment** | [**RoleAssignment**](RoleAssignment.md) |  | 

### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserRoleAssignment

> DeleteUserRoleAssignment(ctx, environmentID, userID, roleAssignmentID).Execute()

DELETE User's Role Assignment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRoleAssignmentsApi.DeleteUserRoleAssignment(context.Background(), environmentID, userID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsApi.DeleteUserRoleAssignment``: %v\n", err)
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
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRoleAssignmentRequest struct via the builder pattern


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


## ReadOneUserRoleAssignment

> RoleAssignment ReadOneUserRoleAssignment(ctx, environmentID, userID, roleAssignmentID).Execute()

READ One Role Assignment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRoleAssignmentsApi.ReadOneUserRoleAssignment(context.Background(), environmentID, userID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsApi.ReadOneUserRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneUserRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `UserRoleAssignmentsApi.ReadOneUserRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneUserRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadUserRoleAssignments

> EntityArray ReadUserRoleAssignments(ctx, environmentID, userID).Execute()

READ Role Assignments

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserRoleAssignmentsApi.ReadUserRoleAssignments(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserRoleAssignmentsApi.ReadUserRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserRoleAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `UserRoleAssignmentsApi.ReadUserRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

