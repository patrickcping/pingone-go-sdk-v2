# \AuthorizationVersionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTag**](AuthorizationVersionsApi.md#DeleteTag) | **Delete** /environments/{environmentId}/authorizationVersions/{authorizationVersionId}/tag | Remove a Tag from an Authorize Version
[**GetPolicyInVersion**](AuthorizationVersionsApi.md#GetPolicyInVersion) | **Get** /environments/{environmentId}/authorizationVersions/{authorizationVersionId}/policies/{policyId} | Get a Policy from a specific Authorize Version
[**GetTag**](AuthorizationVersionsApi.md#GetTag) | **Get** /environments/{environmentId}/authorizationVersions/{authorizationVersionId}/tag | Get a Tag for an Authorize Version
[**GetVersion**](AuthorizationVersionsApi.md#GetVersion) | **Get** /environments/{environmentId}/authorizationVersions/{authorizationVersionId} | Get an Authorize Version by ID
[**ListVersions**](AuthorizationVersionsApi.md#ListVersions) | **Get** /environments/{environmentId}/authorizationVersions | List Authorize Versions
[**UpdateTag**](AuthorizationVersionsApi.md#UpdateTag) | **Put** /environments/{environmentId}/authorizationVersions/{authorizationVersionId}/tag | Create or Update a Tag for an Authorize Version



## DeleteTag

> DeleteTag(ctx, environmentId, authorizationVersionId).Execute()

Remove a Tag from an Authorize Version



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to remove the version tag.
    authorizationVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the version to remove the tag from.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizationVersionsApi.DeleteTag(context.Background(), environmentId, authorizationVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to remove the version tag. | 
**authorizationVersionId** | **string** | The ID of the version to remove the tag from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetPolicyInVersion

> AuthorizeEditorDataPoliciesReferenceablePolicyDTO GetPolicyInVersion(ctx, environmentId, authorizationVersionId, policyId).Execute()

Get a Policy from a specific Authorize Version



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the policy.
    authorizationVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Authorize Version from which to get the policy.
    policyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the policy to retrieve.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationVersionsApi.GetPolicyInVersion(context.Background(), environmentId, authorizationVersionId, policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.GetPolicyInVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyInVersion`: AuthorizeEditorDataPoliciesReferenceablePolicyDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationVersionsApi.GetPolicyInVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to get the policy. | 
**authorizationVersionId** | **string** | The ID of the Authorize Version from which to get the policy. | 
**policyId** | **string** | The ID of the policy to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyInVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## GetTag

> AuthorizeEditorDataTagResponseDTO GetTag(ctx, environmentId, authorizationVersionId).Execute()

Get a Tag for an Authorize Version



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the version tag.
    authorizationVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the version to retrieve the tag for.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationVersionsApi.GetTag(context.Background(), environmentId, authorizationVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.GetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTag`: AuthorizeEditorDataTagResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationVersionsApi.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to get the version tag. | 
**authorizationVersionId** | **string** | The ID of the version to retrieve the tag for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataTagResponseDTO**](AuthorizeEditorDataTagResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> AuthorizeEditorDataAuthorizationVersionDTO GetVersion(ctx, environmentId, authorizationVersionId).Execute()

Get an Authorize Version by ID



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the version.
    authorizationVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the version to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationVersionsApi.GetVersion(context.Background(), environmentId, authorizationVersionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.GetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVersion`: AuthorizeEditorDataAuthorizationVersionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationVersionsApi.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to get the version. | 
**authorizationVersionId** | **string** | The ID of the version to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataAuthorizationVersionDTO**](AuthorizeEditorDataAuthorizationVersionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVersions

> ListVersions200Response ListVersions(ctx, environmentId).Filter(filter).Limit(limit).Cursor(cursor).Filter2(filter2).Execute()

List Authorize Versions



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list versions.
    filter := "filter_example" // string | The SCIM filter (RFC 7644 Section 3.4.2.2) that should be used to determine the resources to return. Filterable attributes: tag (pr)
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)
    filter2 := *openapiclient.NewFilter() // Filter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationVersionsApi.ListVersions(context.Background(), environmentId).Filter(filter).Limit(limit).Cursor(cursor).Filter2(filter2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.ListVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVersions`: ListVersions200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationVersionsApi.ListVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to list versions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The SCIM filter (RFC 7644 Section 3.4.2.2) that should be used to determine the resources to return. Filterable attributes: tag (pr) | 
 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 
 **filter2** | [**Filter**](Filter.md) |  | 

### Return type

[**ListVersions200Response**](ListVersions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> AuthorizeEditorDataTagResponseDTO UpdateTag(ctx, environmentId, authorizationVersionId).AuthorizeEditorDataTagRequestDTO(authorizeEditorDataTagRequestDTO).Execute()

Create or Update a Tag for an Authorize Version



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create or update the version tag.
    authorizationVersionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the version to create or update the tag for.
    authorizeEditorDataTagRequestDTO := *openapiclient.NewAuthorizeEditorDataTagRequestDTO() // AuthorizeEditorDataTagRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationVersionsApi.UpdateTag(context.Background(), environmentId, authorizationVersionId).AuthorizeEditorDataTagRequestDTO(authorizeEditorDataTagRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationVersionsApi.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: AuthorizeEditorDataTagResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationVersionsApi.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The ID of the environment from which to create or update the version tag. | 
**authorizationVersionId** | **string** | The ID of the version to create or update the tag for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataTagRequestDTO** | [**AuthorizeEditorDataTagRequestDTO**](AuthorizeEditorDataTagRequestDTO.md) |  | 

### Return type

[**AuthorizeEditorDataTagResponseDTO**](AuthorizeEditorDataTagResponseDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

