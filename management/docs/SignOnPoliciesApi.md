# \SignOnPoliciesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSignOnPolicy**](SignOnPoliciesApi.md#CreateSignOnPolicy) | **Post** /v1/environments/{environmentID}/signOnPolicies | CREATE Sign On Policy
[**DeleteSignOnPolicy**](SignOnPoliciesApi.md#DeleteSignOnPolicy) | **Delete** /v1/environments/{environmentID}/signOnPolicies/{policyID} | DELETE Sign On Policy
[**ReadAllSignOnPolicies**](SignOnPoliciesApi.md#ReadAllSignOnPolicies) | **Get** /v1/environments/{environmentID}/signOnPolicies | READ All Sign On Policies
[**ReadOneSignOnPolicy**](SignOnPoliciesApi.md#ReadOneSignOnPolicy) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID} | READ One Sign On Policy
[**UpdateSignOnPolicy**](SignOnPoliciesApi.md#UpdateSignOnPolicy) | **Put** /v1/environments/{environmentID}/signOnPolicies/{policyID} | UPDATE Sign On Policy



## CreateSignOnPolicy

> SignOnPolicy CreateSignOnPolicy(ctx, environmentID).SignOnPolicy(signOnPolicy).Execute()

CREATE Sign On Policy

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
    signOnPolicy := *openapiclient.NewSignOnPolicy("Name_example") // SignOnPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesApi.CreateSignOnPolicy(context.Background(), environmentID).SignOnPolicy(signOnPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesApi.CreateSignOnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSignOnPolicy`: SignOnPolicy
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesApi.CreateSignOnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignOnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signOnPolicy** | [**SignOnPolicy**](SignOnPolicy.md) |  | 

### Return type

[**SignOnPolicy**](SignOnPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSignOnPolicy

> DeleteSignOnPolicy(ctx, environmentID, policyID).Execute()

DELETE Sign On Policy

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
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SignOnPoliciesApi.DeleteSignOnPolicy(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesApi.DeleteSignOnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSignOnPolicyRequest struct via the builder pattern


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


## ReadAllSignOnPolicies

> EntityArray ReadAllSignOnPolicies(ctx, environmentID).Execute()

READ All Sign On Policies

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesApi.ReadAllSignOnPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesApi.ReadAllSignOnPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSignOnPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesApi.ReadAllSignOnPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSignOnPoliciesRequest struct via the builder pattern


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


## ReadOneSignOnPolicy

> SignOnPolicy ReadOneSignOnPolicy(ctx, environmentID, policyID).Execute()

READ One Sign On Policy

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
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesApi.ReadOneSignOnPolicy(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesApi.ReadOneSignOnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneSignOnPolicy`: SignOnPolicy
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesApi.ReadOneSignOnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneSignOnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SignOnPolicy**](SignOnPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignOnPolicy

> SignOnPolicy UpdateSignOnPolicy(ctx, environmentID, policyID).SignOnPolicy(signOnPolicy).Execute()

UPDATE Sign On Policy

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
    policyID := "policyID_example" // string | 
    signOnPolicy := *openapiclient.NewSignOnPolicy("Name_example") // SignOnPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesApi.UpdateSignOnPolicy(context.Background(), environmentID, policyID).SignOnPolicy(signOnPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesApi.UpdateSignOnPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSignOnPolicy`: SignOnPolicy
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesApi.UpdateSignOnPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignOnPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signOnPolicy** | [**SignOnPolicy**](SignOnPolicy.md) |  | 

### Return type

[**SignOnPolicy**](SignOnPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

