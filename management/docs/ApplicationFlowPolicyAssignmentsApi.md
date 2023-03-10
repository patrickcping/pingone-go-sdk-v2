# \ApplicationFlowPolicyAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlowPolicyAssignment**](ApplicationFlowPolicyAssignmentsApi.md#CreateFlowPolicyAssignment) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/flowPolicyAssignments | CREATE Flow Assignment
[**DeleteFlowPolicyAssignment**](ApplicationFlowPolicyAssignmentsApi.md#DeleteFlowPolicyAssignment) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/flowPolicyAssignments/{flowPolicyAssignmentID} | DELETE Flow Policy Assignment
[**ReadAllFlowPolicyAssignments**](ApplicationFlowPolicyAssignmentsApi.md#ReadAllFlowPolicyAssignments) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/flowPolicyAssignments | READ All Flow Policy Assignments
[**ReadOneFlowPolicyAssignment**](ApplicationFlowPolicyAssignmentsApi.md#ReadOneFlowPolicyAssignment) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/flowPolicyAssignments/{flowPolicyAssignmentID} | READ One Flow Policy Assignment
[**UpdateFlowPolicyAssignment**](ApplicationFlowPolicyAssignmentsApi.md#UpdateFlowPolicyAssignment) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/flowPolicyAssignments/{flowPolicyAssignmentID} | UPDATE Flow Policy Assignment



## CreateFlowPolicyAssignment

> FlowPolicyAssignment CreateFlowPolicyAssignment(ctx, environmentID, applicationID).FlowPolicyAssignment(flowPolicyAssignment).Execute()

CREATE Flow Assignment

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
    applicationID := "applicationID_example" // string | 
    flowPolicyAssignment := *openapiclient.NewFlowPolicyAssignment(*openapiclient.NewApplicationAccessControlGroupGroupsInner("Id_example")) // FlowPolicyAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFlowPolicyAssignmentsApi.CreateFlowPolicyAssignment(context.Background(), environmentID, applicationID).FlowPolicyAssignment(flowPolicyAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFlowPolicyAssignmentsApi.CreateFlowPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFlowPolicyAssignment`: FlowPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFlowPolicyAssignmentsApi.CreateFlowPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flowPolicyAssignment** | [**FlowPolicyAssignment**](FlowPolicyAssignment.md) |  | 

### Return type

[**FlowPolicyAssignment**](FlowPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlowPolicyAssignment

> DeleteFlowPolicyAssignment(ctx, environmentID, applicationID, flowPolicyAssignmentID).Execute()

DELETE Flow Policy Assignment

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
    applicationID := "applicationID_example" // string | 
    flowPolicyAssignmentID := "flowPolicyAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationFlowPolicyAssignmentsApi.DeleteFlowPolicyAssignment(context.Background(), environmentID, applicationID, flowPolicyAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFlowPolicyAssignmentsApi.DeleteFlowPolicyAssignment``: %v\n", err)
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
**flowPolicyAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowPolicyAssignmentRequest struct via the builder pattern


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


## ReadAllFlowPolicyAssignments

> EntityArray ReadAllFlowPolicyAssignments(ctx, environmentID, applicationID).Execute()

READ All Flow Policy Assignments

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
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllFlowPolicyAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFlowPolicyAssignmentsApi.ReadAllFlowPolicyAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllFlowPolicyAssignmentsRequest struct via the builder pattern


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


## ReadOneFlowPolicyAssignment

> FlowPolicyAssignment ReadOneFlowPolicyAssignment(ctx, environmentID, applicationID, flowPolicyAssignmentID).Execute()

READ One Flow Policy Assignment

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
    applicationID := "applicationID_example" // string | 
    flowPolicyAssignmentID := "flowPolicyAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFlowPolicyAssignmentsApi.ReadOneFlowPolicyAssignment(context.Background(), environmentID, applicationID, flowPolicyAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFlowPolicyAssignmentsApi.ReadOneFlowPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneFlowPolicyAssignment`: FlowPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFlowPolicyAssignmentsApi.ReadOneFlowPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**flowPolicyAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFlowPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlowPolicyAssignment**](FlowPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFlowPolicyAssignment

> FlowPolicyAssignment UpdateFlowPolicyAssignment(ctx, environmentID, applicationID, flowPolicyAssignmentID).FlowPolicyAssignment(flowPolicyAssignment).Execute()

UPDATE Flow Policy Assignment

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
    applicationID := "applicationID_example" // string | 
    flowPolicyAssignmentID := "flowPolicyAssignmentID_example" // string | 
    flowPolicyAssignment := *openapiclient.NewFlowPolicyAssignment(*openapiclient.NewApplicationAccessControlGroupGroupsInner("Id_example")) // FlowPolicyAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationFlowPolicyAssignmentsApi.UpdateFlowPolicyAssignment(context.Background(), environmentID, applicationID, flowPolicyAssignmentID).FlowPolicyAssignment(flowPolicyAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationFlowPolicyAssignmentsApi.UpdateFlowPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFlowPolicyAssignment`: FlowPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationFlowPolicyAssignmentsApi.UpdateFlowPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**flowPolicyAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFlowPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **flowPolicyAssignment** | [**FlowPolicyAssignment**](FlowPolicyAssignment.md) |  | 

### Return type

[**FlowPolicyAssignment**](FlowPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

