# \ApplicationRoleApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteApplicationRole**](ApplicationRoleApi.md#DeleteApplicationRole) | **Delete** /environments/{environmentID}/applicationRoles/{applicationRoleID} | DELETE Application Role
[**UpdateApplicationRole**](ApplicationRoleApi.md#UpdateApplicationRole) | **Put** /environments/{environmentID}/applicationRoles/{applicationRoleID} | UPDATE Application Role



## DeleteApplicationRole

> DeleteApplicationRole(ctx, environmentID, applicationRoleID).Execute()

DELETE Application Role

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
    r, err := apiClient.ApplicationRoleApi.DeleteApplicationRole(context.Background(), environmentID, applicationRoleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleApi.DeleteApplicationRole``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRoleRequest struct via the builder pattern


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


## UpdateApplicationRole

> ApplicationRole UpdateApplicationRole(ctx, environmentID, applicationRoleID).ApplicationRole(applicationRole).Execute()

UPDATE Application Role

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
    applicationRole := *openapiclient.NewApplicationRole("Name_example") // ApplicationRole |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationRoleApi.UpdateApplicationRole(context.Background(), environmentID, applicationRoleID).ApplicationRole(applicationRole).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationRoleApi.UpdateApplicationRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationRole`: ApplicationRole
    fmt.Fprintf(os.Stdout, "Response from `ApplicationRoleApi.UpdateApplicationRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationRoleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationRole** | [**ApplicationRole**](ApplicationRole.md) |  | 

### Return type

[**ApplicationRole**](ApplicationRole.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

