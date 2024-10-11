# \AuthorizeEditorPoliciesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicy**](AuthorizeEditorPoliciesApi.md#CreatePolicy) | **Post** /environments/{environmentID}/authorizationPolicies | Create a Policy
[**DeletePolicy**](AuthorizeEditorPoliciesApi.md#DeletePolicy) | **Delete** /environments/{environmentID}/authorizationPolicies/{policyID} | Delete a Policy
[**GetPolicy**](AuthorizeEditorPoliciesApi.md#GetPolicy) | **Get** /environments/{environmentID}/authorizationPolicies/{policyID} | Get a Policy by ID
[**ListRootPolicies**](AuthorizeEditorPoliciesApi.md#ListRootPolicies) | **Get** /environments/{environmentID}/authorizationPolicies | List Root Policies
[**TestPolicy**](AuthorizeEditorPoliciesApi.md#TestPolicy) | **Post** /environments/{environmentID}/authorizationPolicies/{policyID} | Test a Policy by ID
[**UpdatePolicy**](AuthorizeEditorPoliciesApi.md#UpdatePolicy) | **Put** /environments/{environmentID}/authorizationPolicies/{policyID} | Update a Policy



## CreatePolicy

> AuthorizeEditorDataPoliciesReferenceablePolicyDTO CreatePolicy(ctx, environmentID).AuthorizeEditorDataPoliciesPolicyDTO(authorizeEditorDataPoliciesPolicyDTO).Accept(accept).Execute()

Create a Policy



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the policy.
    authorizeEditorDataPoliciesPolicyDTO := *openapiclient.NewAuthorizeEditorDataPoliciesPolicyDTO("Name_example", *openapiclient.NewAuthorizeEditorDataPoliciesCombiningAlgorithmDTO(openapiclient.EnumAuthorizeEditorDataPoliciesCombiningAlgorithmDTOAlgorithm("PERMIT_OVERRIDES"))) // AuthorizeEditorDataPoliciesPolicyDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorPoliciesApi.CreatePolicy(context.Background(), environmentID).AuthorizeEditorDataPoliciesPolicyDTO(authorizeEditorDataPoliciesPolicyDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: AuthorizeEditorDataPoliciesReferenceablePolicyDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorPoliciesApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the policy. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataPoliciesPolicyDTO** | [**AuthorizeEditorDataPoliciesPolicyDTO**](AuthorizeEditorDataPoliciesPolicyDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, environmentID, policyID).Execute()

Delete a Policy



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the policy.
    policyID := "policyID_example" // string | The ID of the policy to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorPoliciesApi.DeletePolicy(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the policy. | 
**policyID** | **string** | The ID of the policy to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


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


## GetPolicy

> AuthorizeEditorDataPoliciesReferenceablePolicyDTO GetPolicy(ctx, environmentID, policyID).Accept(accept).Execute()

Get a Policy by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the policy.
    policyID := "policyID_example" // string | The ID of the policy to retrieve
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorPoliciesApi.GetPolicy(context.Background(), environmentID, policyID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: AuthorizeEditorDataPoliciesReferenceablePolicyDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorPoliciesApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the policy. | 
**policyID** | **string** | The ID of the policy to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRootPolicies

> EntityArray ListRootPolicies(ctx, environmentID).Limit(limit).Cursor(cursor).Execute()

List Root Policies



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list root Policies
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorPoliciesApi.ListRootPolicies(context.Background(), environmentID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.ListRootPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRootPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorPoliciesApi.ListRootPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list root Policies | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRootPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestPolicy

> AuthorizeEditorDataEntityTestResponseDTO TestPolicy(ctx, environmentID, policyID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()

Test a Policy by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    policyID := "policyID_example" // string | 
    authorizeEditorDataEntityTestRequestDTO := *openapiclient.NewAuthorizeEditorDataEntityTestRequestDTO([]openapiclient.AuthorizeEditorDataEntityTestingParameterDTO{*openapiclient.NewAuthorizeEditorDataEntityTestingParameterDTO("Key_example", "Value_example")}) // AuthorizeEditorDataEntityTestRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorPoliciesApi.TestPolicy(context.Background(), environmentID, policyID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.TestPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestPolicy`: AuthorizeEditorDataEntityTestResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorPoliciesApi.TestPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataEntityTestRequestDTO** | [**AuthorizeEditorDataEntityTestRequestDTO**](AuthorizeEditorDataEntityTestRequestDTO.md) |  | 

### Return type

[**AuthorizeEditorDataEntityTestResponseDTO**](AuthorizeEditorDataEntityTestResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicy

> AuthorizeEditorDataPoliciesReferenceablePolicyDTO UpdatePolicy(ctx, environmentID, policyID).AuthorizeEditorDataPoliciesReferenceablePolicyDTO(authorizeEditorDataPoliciesReferenceablePolicyDTO).Accept(accept).Execute()

Update a Policy



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the policy.
    policyID := "policyID_example" // string | The ID of the policy to be updated
    authorizeEditorDataPoliciesReferenceablePolicyDTO := *openapiclient.NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO("Id_example", "Name_example", *openapiclient.NewAuthorizeEditorDataPoliciesCombiningAlgorithmDTO(openapiclient.EnumAuthorizeEditorDataPoliciesCombiningAlgorithmDTOAlgorithm("PERMIT_OVERRIDES")), "Version_example") // AuthorizeEditorDataPoliciesReferenceablePolicyDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorPoliciesApi.UpdatePolicy(context.Background(), environmentID, policyID).AuthorizeEditorDataPoliciesReferenceablePolicyDTO(authorizeEditorDataPoliciesReferenceablePolicyDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorPoliciesApi.UpdatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicy`: AuthorizeEditorDataPoliciesReferenceablePolicyDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorPoliciesApi.UpdatePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the policy. | 
**policyID** | **string** | The ID of the policy to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataPoliciesReferenceablePolicyDTO** | [**AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataPoliciesReferenceablePolicyDTO**](AuthorizeEditorDataPoliciesReferenceablePolicyDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

