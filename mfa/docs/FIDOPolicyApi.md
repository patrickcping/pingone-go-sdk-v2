# \FIDOPolicyApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFidoPolicy**](FIDOPolicyApi.md#CreateFidoPolicy) | **Post** /v1/environments/{environmentID}/fidoPolicies | CREATE FIDO Policy
[**DeleteFidoPolicy**](FIDOPolicyApi.md#DeleteFidoPolicy) | **Delete** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | DELETE FIDO Policy
[**ReadFidoPolicies**](FIDOPolicyApi.md#ReadFidoPolicies) | **Get** /v1/environments/{environmentID}/fidoPolicies | READ FIDO Policies
[**ReadOneFidoPolicy**](FIDOPolicyApi.md#ReadOneFidoPolicy) | **Get** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | READ One FIDO Policy
[**UpdateFIDOPolicy**](FIDOPolicyApi.md#UpdateFIDOPolicy) | **Put** /v1/environments/{environmentID}/fidoPolicies/{fidoPolicyID} | UPDATE FIDO Policy



## CreateFidoPolicy

> FIDOPolicy CreateFidoPolicy(ctx, environmentID).FIDOPolicy(fIDOPolicy).Execute()

CREATE FIDO Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    fIDOPolicy := *openapiclient.NewFIDOPolicy("Name_example", openapiclient.EnumFIDOAttestationRequirements("NONE"), openapiclient.EnumFIDOResidentKeyRequirement("DISCOURAGED")) // FIDOPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDOPolicyApi.CreateFidoPolicy(context.Background(), environmentID).FIDOPolicy(fIDOPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDOPolicyApi.CreateFidoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFidoPolicy`: FIDOPolicy
    fmt.Fprintf(os.Stdout, "Response from `FIDOPolicyApi.CreateFidoPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFidoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fIDOPolicy** | [**FIDOPolicy**](FIDOPolicy.md) |  | 

### Return type

[**FIDOPolicy**](FIDOPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFidoPolicy

> DeleteFidoPolicy(ctx, environmentID, fidoPolicyID).Execute()

DELETE FIDO Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    fidoPolicyID := "fidoPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FIDOPolicyApi.DeleteFidoPolicy(context.Background(), environmentID, fidoPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDOPolicyApi.DeleteFidoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFidoPolicyRequest struct via the builder pattern


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


## ReadFidoPolicies

> EntityArray ReadFidoPolicies(ctx, environmentID).Execute()

READ FIDO Policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDOPolicyApi.ReadFidoPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDOPolicyApi.ReadFidoPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFidoPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `FIDOPolicyApi.ReadFidoPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFidoPoliciesRequest struct via the builder pattern


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


## ReadOneFidoPolicy

> FIDOPolicy ReadOneFidoPolicy(ctx, environmentID, fidoPolicyID).Execute()

READ One FIDO Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    fidoPolicyID := "fidoPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDOPolicyApi.ReadOneFidoPolicy(context.Background(), environmentID, fidoPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDOPolicyApi.ReadOneFidoPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneFidoPolicy`: FIDOPolicy
    fmt.Fprintf(os.Stdout, "Response from `FIDOPolicyApi.ReadOneFidoPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFidoPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FIDOPolicy**](FIDOPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFIDOPolicy

> FIDOPolicy UpdateFIDOPolicy(ctx, environmentID, fidoPolicyID).FIDOPolicy(fIDOPolicy).Execute()

UPDATE FIDO Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    fidoPolicyID := "fidoPolicyID_example" // string | 
    fIDOPolicy := *openapiclient.NewFIDOPolicy("Name_example", openapiclient.EnumFIDOAttestationRequirements("NONE"), openapiclient.EnumFIDOResidentKeyRequirement("DISCOURAGED")) // FIDOPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDOPolicyApi.UpdateFIDOPolicy(context.Background(), environmentID, fidoPolicyID).FIDOPolicy(fIDOPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDOPolicyApi.UpdateFIDOPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFIDOPolicy`: FIDOPolicy
    fmt.Fprintf(os.Stdout, "Response from `FIDOPolicyApi.UpdateFIDOPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFIDOPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fIDOPolicy** | [**FIDOPolicy**](FIDOPolicy.md) |  | 

### Return type

[**FIDOPolicy**](FIDOPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

