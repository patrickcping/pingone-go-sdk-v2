# \GatewayRoleAssignmentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayRoleAssignment**](GatewayRoleAssignmentsApi.md#CreateGatewayRoleAssignment) | **Post** /environments/{environmentID}/gateways/{gatewayID}/roleAssignments | CREATE Gateway Role Assignments
[**DeleteGatewayRoleAssignment**](GatewayRoleAssignmentsApi.md#DeleteGatewayRoleAssignment) | **Delete** /environments/{environmentID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | DELETE Gateway Role Assignment
[**ReadGatewayRoleAssignments**](GatewayRoleAssignmentsApi.md#ReadGatewayRoleAssignments) | **Get** /environments/{environmentID}/gateways/{gatewayID}/roleAssignments | READ Gateway Role Assignments
[**ReadOneGatewayRoleAssignment**](GatewayRoleAssignmentsApi.md#ReadOneGatewayRoleAssignment) | **Get** /environments/{environmentID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | READ One Gateway Role Assignment
[**UpdateGatewayRoleAssignment**](GatewayRoleAssignmentsApi.md#UpdateGatewayRoleAssignment) | **Put** /environments/{environmentID}/gateways/{gatewayID}/roleAssignments/{gatewayRoleAssignmentID} | UPDATE Gateway Role Assignments



## CreateGatewayRoleAssignment

> RoleAssignment CreateGatewayRoleAssignment(ctx, environmentID, gatewayID).RoleAssignment(roleAssignment).Execute()

CREATE Gateway Role Assignments

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
    gatewayID := "gatewayID_example" // string | 
    roleAssignment := *openapiclient.NewRoleAssignment(*openapiclient.NewRoleAssignmentRole("Id_example"), *openapiclient.NewRoleAssignmentScope("Id_example", openapiclient.EnumRoleAssignmentScopeType("ORGANIZATION"))) // RoleAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayRoleAssignmentsApi.CreateGatewayRoleAssignment(context.Background(), environmentID, gatewayID).RoleAssignment(roleAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoleAssignmentsApi.CreateGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GatewayRoleAssignmentsApi.CreateGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRoleAssignmentRequest struct via the builder pattern


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


## DeleteGatewayRoleAssignment

> DeleteGatewayRoleAssignment(ctx, environmentID, gatewayID, gatewayRoleAssignmentID).Execute()

DELETE Gateway Role Assignment

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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GatewayRoleAssignmentsApi.DeleteGatewayRoleAssignment(context.Background(), environmentID, gatewayID, gatewayRoleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoleAssignmentsApi.DeleteGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayRoleAssignmentRequest struct via the builder pattern


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


## ReadGatewayRoleAssignments

> EntityArray ReadGatewayRoleAssignments(ctx, environmentID, gatewayID).Execute()

READ Gateway Role Assignments

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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadGatewayRoleAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `GatewayRoleAssignmentsApi.ReadGatewayRoleAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadGatewayRoleAssignmentsRequest struct via the builder pattern


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


## ReadOneGatewayRoleAssignment

> RoleAssignment ReadOneGatewayRoleAssignment(ctx, environmentID, gatewayID, gatewayRoleAssignmentID).Execute()

READ One Gateway Role Assignment

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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment(context.Background(), environmentID, gatewayID, gatewayRoleAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GatewayRoleAssignmentsApi.ReadOneGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayRoleAssignmentRequest struct via the builder pattern


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


## UpdateGatewayRoleAssignment

> RoleAssignment UpdateGatewayRoleAssignment(ctx, environmentID, gatewayID, gatewayRoleAssignmentID).Body(body).Execute()

UPDATE Gateway Role Assignments

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
    gatewayID := "gatewayID_example" // string | 
    gatewayRoleAssignmentID := "gatewayRoleAssignmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment(context.Background(), environmentID, gatewayID, gatewayRoleAssignmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGatewayRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `GatewayRoleAssignmentsApi.UpdateGatewayRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**gatewayRoleAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **map[string]interface{}** |  | 

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

