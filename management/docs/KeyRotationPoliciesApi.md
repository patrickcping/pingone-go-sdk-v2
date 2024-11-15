# \KeyRotationPoliciesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeyRotationPolicy**](KeyRotationPoliciesApi.md#CreateKeyRotationPolicy) | **Post** /environments/{environmentID}/keyRotationPolicies | CREATE Key Rotation Policy
[**DeleteKeyRotationPolicy**](KeyRotationPoliciesApi.md#DeleteKeyRotationPolicy) | **Delete** /environments/{environmentID}/keyRotationPolicies/{keyRotationPolicyID} | DELETE Key Rotation Policy
[**GetKeyRotationPolicies**](KeyRotationPoliciesApi.md#GetKeyRotationPolicies) | **Get** /environments/{environmentID}/keyRotationPolicies | GET Key Rotation Policies
[**GetKeyRotationPolicy**](KeyRotationPoliciesApi.md#GetKeyRotationPolicy) | **Get** /environments/{environmentID}/keyRotationPolicies/{keyRotationPolicyID} | GET Key Rotation Policy
[**UpdateKeyRotationPolicy**](KeyRotationPoliciesApi.md#UpdateKeyRotationPolicy) | **Put** /environments/{environmentID}/keyRotationPolicies/{keyRotationPolicyID} | UPDATE Key Rotation Policy



## CreateKeyRotationPolicy

> KeyRotationPolicy CreateKeyRotationPolicy(ctx, environmentID).KeyRotationPolicy(keyRotationPolicy).Execute()

CREATE Key Rotation Policy

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
    keyRotationPolicy := *openapiclient.NewKeyRotationPolicy(openapiclient.EnumKeyRotationPolicyAlgorithm("RSA"), "Dn_example", int32(123), "Name_example", openapiclient.EnumKeyRotationPolicySigAlgorithm("SHA256withRSA"), openapiclient.EnumKeyRotationPolicyUsageType("SIGNING")) // KeyRotationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationPoliciesApi.CreateKeyRotationPolicy(context.Background(), environmentID).KeyRotationPolicy(keyRotationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationPoliciesApi.CreateKeyRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeyRotationPolicy`: KeyRotationPolicy
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationPoliciesApi.CreateKeyRotationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyRotationPolicy** | [**KeyRotationPolicy**](KeyRotationPolicy.md) |  | 

### Return type

[**KeyRotationPolicy**](KeyRotationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeyRotationPolicy

> DeleteKeyRotationPolicy(ctx, environmentID, keyRotationPolicyID).Execute()

DELETE Key Rotation Policy

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
    keyRotationPolicyID := "keyRotationPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KeyRotationPoliciesApi.DeleteKeyRotationPolicy(context.Background(), environmentID, keyRotationPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationPoliciesApi.DeleteKeyRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyRotationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRotationPolicyRequest struct via the builder pattern


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


## GetKeyRotationPolicies

GET Key Rotation Policies

### Paged Response (Recommended)

> EntityArrayPagedIterator GetKeyRotationPolicies(ctx, environmentID).Execute()

#### Example

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.GetKeyRotationPolicies(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.GetKeyRotationPolicies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `GetKeyRotationPolicies`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.GetKeyRotationPolicies`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray GetKeyRotationPolicies(ctx, environmentID).ExecuteInitialPage()

#### Example

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
    resp, r, err := apiClient.KeyRotationPoliciesApi.GetKeyRotationPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationPoliciesApi.GetKeyRotationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyRotationPolicies`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationPoliciesApi.GetKeyRotationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRotationPoliciesRequest struct via the builder pattern


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


## GetKeyRotationPolicy

> KeyRotationPolicy GetKeyRotationPolicy(ctx, environmentID, keyRotationPolicyID).Execute()

GET Key Rotation Policy

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
    keyRotationPolicyID := "keyRotationPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationPoliciesApi.GetKeyRotationPolicy(context.Background(), environmentID, keyRotationPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationPoliciesApi.GetKeyRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyRotationPolicy`: KeyRotationPolicy
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationPoliciesApi.GetKeyRotationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyRotationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyRotationPolicy**](KeyRotationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKeyRotationPolicy

> KeyRotationPolicy UpdateKeyRotationPolicy(ctx, environmentID, keyRotationPolicyID).KeyRotationPolicy(keyRotationPolicy).Execute()

UPDATE Key Rotation Policy

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
    keyRotationPolicyID := "keyRotationPolicyID_example" // string | 
    keyRotationPolicy := *openapiclient.NewKeyRotationPolicy(openapiclient.EnumKeyRotationPolicyAlgorithm("RSA"), "Dn_example", int32(123), "Name_example", openapiclient.EnumKeyRotationPolicySigAlgorithm("SHA256withRSA"), openapiclient.EnumKeyRotationPolicyUsageType("SIGNING")) // KeyRotationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KeyRotationPoliciesApi.UpdateKeyRotationPolicy(context.Background(), environmentID, keyRotationPolicyID).KeyRotationPolicy(keyRotationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KeyRotationPoliciesApi.UpdateKeyRotationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKeyRotationPolicy`: KeyRotationPolicy
    fmt.Fprintf(os.Stdout, "Response from `KeyRotationPoliciesApi.UpdateKeyRotationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyRotationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRotationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **keyRotationPolicy** | [**KeyRotationPolicy**](KeyRotationPolicy.md) |  | 

### Return type

[**KeyRotationPolicy**](KeyRotationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

