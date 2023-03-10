# \FlowPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAllFlowPolicies**](FlowPoliciesApi.md#ReadAllFlowPolicies) | **Get** /v1/environments/{environmentID}/flowPolicies | READ All Flow Policies
[**ReadOneFlowPolicy**](FlowPoliciesApi.md#ReadOneFlowPolicy) | **Get** /v1/environments/{environmentID}/flowPolicies/{flowPolicyID} | READ ONE Flow Policy



## ReadAllFlowPolicies

> EntityArray ReadAllFlowPolicies(ctx, environmentID).Filter(filter).Execute()

READ All Flow Policies

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
    filter := "trigger.type eq "AUTHENTICATION"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowPoliciesApi.ReadAllFlowPolicies(context.Background(), environmentID).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowPoliciesApi.ReadAllFlowPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllFlowPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `FlowPoliciesApi.ReadAllFlowPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllFlowPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 

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


## ReadOneFlowPolicy

> FlowPolicy ReadOneFlowPolicy(ctx, environmentID, flowPolicyID).Execute()

READ ONE Flow Policy

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
    flowPolicyID := "flowPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlowPoliciesApi.ReadOneFlowPolicy(context.Background(), environmentID, flowPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlowPoliciesApi.ReadOneFlowPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneFlowPolicy`: FlowPolicy
    fmt.Fprintf(os.Stdout, "Response from `FlowPoliciesApi.ReadOneFlowPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**flowPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFlowPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FlowPolicy**](FlowPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

