# \VerifyPoliciesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVerifyPolicy**](VerifyPoliciesApi.md#CreateVerifyPolicy) | **Post** /environments/{environmentID}/verifyPolicies | CREATE Verify Policy
[**DeleteVerifyPolicy**](VerifyPoliciesApi.md#DeleteVerifyPolicy) | **Delete** /environments/{environmentID}/verifyPolicies/{verifyPolicyID} | Delete Verify Policy
[**ReadAllVerifyPolicies**](VerifyPoliciesApi.md#ReadAllVerifyPolicies) | **Get** /environments/{environmentID}/verifyPolicies | READ All Verify Policies
[**ReadOneVerifyPolicy**](VerifyPoliciesApi.md#ReadOneVerifyPolicy) | **Get** /environments/{environmentID}/verifyPolicies/{verifyPolicyID} | READ One Verify Policy
[**UpdateVerifyPolicy**](VerifyPoliciesApi.md#UpdateVerifyPolicy) | **Put** /environments/{environmentID}/verifyPolicies/{verifyPolicyID} | UPDATE Verify Policy



## CreateVerifyPolicy

> VerifyPolicy CreateVerifyPolicy(ctx, environmentID).VerifyPolicy(verifyPolicy).Execute()

CREATE Verify Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 
    verifyPolicy := *openapiclient.NewVerifyPolicy("Name_example") // VerifyPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyPoliciesApi.CreateVerifyPolicy(context.Background(), environmentID).VerifyPolicy(verifyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyPoliciesApi.CreateVerifyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVerifyPolicy`: VerifyPolicy
    fmt.Fprintf(os.Stdout, "Response from `VerifyPoliciesApi.CreateVerifyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVerifyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyPolicy** | [**VerifyPolicy**](VerifyPolicy.md) |  | 

### Return type

[**VerifyPolicy**](VerifyPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVerifyPolicy

> DeleteVerifyPolicy(ctx, environmentID, verifyPolicyID).Execute()

Delete Verify Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 
    verifyPolicyID := "verifyPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VerifyPoliciesApi.DeleteVerifyPolicy(context.Background(), environmentID, verifyPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyPoliciesApi.DeleteVerifyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**verifyPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVerifyPolicyRequest struct via the builder pattern


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


## ReadAllVerifyPolicies

READ All Verify Policies

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllVerifyPolicies(ctx, environmentID).Execute()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllVerifyPolicies(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllVerifyPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllVerifyPolicies`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllVerifyPolicies`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllVerifyPolicies(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyPoliciesApi.ReadAllVerifyPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyPoliciesApi.ReadAllVerifyPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllVerifyPolicies`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `VerifyPoliciesApi.ReadAllVerifyPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllVerifyPoliciesRequest struct via the builder pattern


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


## ReadOneVerifyPolicy

> VerifyPolicy ReadOneVerifyPolicy(ctx, environmentID, verifyPolicyID).Execute()

READ One Verify Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 
    verifyPolicyID := "verifyPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyPoliciesApi.ReadOneVerifyPolicy(context.Background(), environmentID, verifyPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyPoliciesApi.ReadOneVerifyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneVerifyPolicy`: VerifyPolicy
    fmt.Fprintf(os.Stdout, "Response from `VerifyPoliciesApi.ReadOneVerifyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**verifyPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneVerifyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VerifyPolicy**](VerifyPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVerifyPolicy

> VerifyPolicy UpdateVerifyPolicy(ctx, environmentID, verifyPolicyID).VerifyPolicy(verifyPolicy).Execute()

UPDATE Verify Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
    environmentID := "environmentID_example" // string | 
    verifyPolicyID := "verifyPolicyID_example" // string | 
    verifyPolicy := *openapiclient.NewVerifyPolicy("Name_example") // VerifyPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VerifyPoliciesApi.UpdateVerifyPolicy(context.Background(), environmentID, verifyPolicyID).VerifyPolicy(verifyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VerifyPoliciesApi.UpdateVerifyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVerifyPolicy`: VerifyPolicy
    fmt.Fprintf(os.Stdout, "Response from `VerifyPoliciesApi.UpdateVerifyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**verifyPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVerifyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **verifyPolicy** | [**VerifyPolicy**](VerifyPolicy.md) |  | 

### Return type

[**VerifyPolicy**](VerifyPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

