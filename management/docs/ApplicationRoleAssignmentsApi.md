# \ApplicationRoleAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationRoleAssignment**](ApplicationRoleAssignmentsApi.md#CreateApplicationRoleAssignment) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/roleAssignments | CREATE Application Role Assignments
[**DeleteApplicationRoleAssignment**](ApplicationRoleAssignmentsApi.md#DeleteApplicationRoleAssignment) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/roleAssignments/{roleAssignmentID} | DELETE Application Role Assignment
[**ReadApplicationRoleAssignments**](ApplicationRoleAssignmentsApi.md#ReadApplicationRoleAssignments) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/roleAssignments | READ Application Role Assignments
[**ReadOneApplicationRoleAssignment**](ApplicationRoleAssignmentsApi.md#ReadOneApplicationRoleAssignment) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/roleAssignments/{roleAssignmentID} | READ One Application Role Assignment



## CreateApplicationRoleAssignment

> RoleAssignment CreateApplicationRoleAssignment(ctx, environmentID, applicationID).RoleAssignment(roleAssignment).Execute()

CREATE Application Role Assignments

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
    applicationID := "applicationID_example" // string | 
    roleAssignment := *openapiclient.NewRoleAssignment(*openapiclient.NewRoleAssignmentRole("Id_example"), *openapiclient.NewRoleAssignmentScope("Id_example", openapiclient.EnumRoleAssignmentScopeType("ORGANIZATION"))) // RoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment(context.Background(), environmentID, applicationID).RoleAssignment(roleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRoleAssignmentsApi.CreateApplicationRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRoleAssignmentRequest struct via the builder pattern


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


## DeleteApplicationRoleAssignment

> DeleteApplicationRoleAssignment(ctx, environmentID, applicationID, roleAssignmentID).Execute()

DELETE Application Role Assignment

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
    applicationID := "applicationID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleAssignmentsApi.DeleteApplicationRoleAssignment(context.Background(), environmentID, applicationID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleAssignmentsApi.DeleteApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRoleAssignmentRequest struct via the builder pattern


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


## ReadApplicationRoleAssignments

> EntityArray ReadApplicationRoleAssignments(ctx, environmentID, applicationID).Execute()

READ Application Role Assignments

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
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadApplicationRoleAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadApplicationRoleAssignmentsRequest struct via the builder pattern


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


## ReadOneApplicationRoleAssignment

> RoleAssignment ReadOneApplicationRoleAssignment(ctx, environmentID, applicationID, roleAssignmentID).Execute()

READ One Application Role Assignment

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
    applicationID := "applicationID_example" // string | 
    roleAssignmentID := "roleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment(context.Background(), environmentID, applicationID, roleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRoleAssignmentsApi.ReadOneApplicationRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**roleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationRoleAssignmentRequest struct via the builder pattern


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

