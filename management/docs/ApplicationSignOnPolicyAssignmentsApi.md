# \ApplicationSignOnPolicyAssignmentsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSignOnPolicyAssignment**](ApplicationSignOnPolicyAssignmentsApi.md#CreateSignOnPolicyAssignment) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments | CREATE SOP Assignment
[**DeleteSignOnPolicyAssignment**](ApplicationSignOnPolicyAssignmentsApi.md#DeleteSignOnPolicyAssignment) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | DELETE SOP Assignment
[**ReadAllSignOnPolicyAssignments**](ApplicationSignOnPolicyAssignmentsApi.md#ReadAllSignOnPolicyAssignments) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments | READ All SOP Assignments
[**ReadOneSignOnPolicyAssignment**](ApplicationSignOnPolicyAssignmentsApi.md#ReadOneSignOnPolicyAssignment) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | READ One SOP Assignment
[**UpdateSignOnPolicyAssignment**](ApplicationSignOnPolicyAssignmentsApi.md#UpdateSignOnPolicyAssignment) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/signOnPolicyAssignments/{SOPAssignmentID} | UPDATE SOP Assignment



## CreateSignOnPolicyAssignment

> SignOnPolicyAssignment CreateSignOnPolicyAssignment(ctx, environmentID, applicationID).SignOnPolicyAssignment(signOnPolicyAssignment).Execute()

CREATE SOP Assignment

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
    signOnPolicyAssignment := *openapiclient.NewSignOnPolicyAssignment(int32(123), *openapiclient.NewSignOnPolicyActionCommonSignOnPolicy("Id_example")) // SignOnPolicyAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSignOnPolicyAssignmentsApi.CreateSignOnPolicyAssignment(context.Background(), environmentID, applicationID).SignOnPolicyAssignment(signOnPolicyAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSignOnPolicyAssignmentsApi.CreateSignOnPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSignOnPolicyAssignment`: SignOnPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSignOnPolicyAssignmentsApi.CreateSignOnPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignOnPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signOnPolicyAssignment** | [**SignOnPolicyAssignment**](SignOnPolicyAssignment.md) |  | 

### Return type

[**SignOnPolicyAssignment**](SignOnPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSignOnPolicyAssignment

> DeleteSignOnPolicyAssignment(ctx, environmentID, applicationID, sOPAssignmentID).Execute()

DELETE SOP Assignment

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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationSignOnPolicyAssignmentsApi.DeleteSignOnPolicyAssignment(context.Background(), environmentID, applicationID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSignOnPolicyAssignmentsApi.DeleteSignOnPolicyAssignment``: %v\n", err)
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
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSignOnPolicyAssignmentRequest struct via the builder pattern


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


## ReadAllSignOnPolicyAssignments

> EntityArray ReadAllSignOnPolicyAssignments(ctx, environmentID, applicationID).Execute()

READ All SOP Assignments

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
    resp, r, err := apiClient.ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSignOnPolicyAssignments`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSignOnPolicyAssignmentsApi.ReadAllSignOnPolicyAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSignOnPolicyAssignmentsRequest struct via the builder pattern


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


## ReadOneSignOnPolicyAssignment

> SignOnPolicyAssignment ReadOneSignOnPolicyAssignment(ctx, environmentID, applicationID, sOPAssignmentID).Execute()

READ One SOP Assignment

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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSignOnPolicyAssignmentsApi.ReadOneSignOnPolicyAssignment(context.Background(), environmentID, applicationID, sOPAssignmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSignOnPolicyAssignmentsApi.ReadOneSignOnPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneSignOnPolicyAssignment`: SignOnPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSignOnPolicyAssignmentsApi.ReadOneSignOnPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneSignOnPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SignOnPolicyAssignment**](SignOnPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignOnPolicyAssignment

> SignOnPolicyAssignment UpdateSignOnPolicyAssignment(ctx, environmentID, applicationID, sOPAssignmentID).SignOnPolicyAssignment(signOnPolicyAssignment).Execute()

UPDATE SOP Assignment

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
    sOPAssignmentID := "sOPAssignmentID_example" // string | 
    signOnPolicyAssignment := *openapiclient.NewSignOnPolicyAssignment(int32(123), *openapiclient.NewSignOnPolicyActionCommonSignOnPolicy("Id_example")) // SignOnPolicyAssignment |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationSignOnPolicyAssignmentsApi.UpdateSignOnPolicyAssignment(context.Background(), environmentID, applicationID, sOPAssignmentID).SignOnPolicyAssignment(signOnPolicyAssignment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationSignOnPolicyAssignmentsApi.UpdateSignOnPolicyAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSignOnPolicyAssignment`: SignOnPolicyAssignment
    fmt.Fprintf(os.Stdout, "Response from `ApplicationSignOnPolicyAssignmentsApi.UpdateSignOnPolicyAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**sOPAssignmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignOnPolicyAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **signOnPolicyAssignment** | [**SignOnPolicyAssignment**](SignOnPolicyAssignment.md) |  | 

### Return type

[**SignOnPolicyAssignment**](SignOnPolicyAssignment.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

