# \ApplicationRoleAssignmentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadApplicationRoleAssignments**](ApplicationRoleAssignmentsApi.md#ReadApplicationRoleAssignments) | **Get** /environments/{environmentID}/applicationRoles/{applicationRoleID}/assignments | READ Application Role Assignments



## ReadApplicationRoleAssignments

> EntityArray ReadApplicationRoleAssignments(ctx, environmentID, applicationRoleID).Execute()

READ Application Role Assignments

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleAssignmentsApi.ReadApplicationRoleAssignments(context.Background(), environmentID, applicationRoleID).Execute()
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
**applicationRoleID** | **string** |  | 

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

