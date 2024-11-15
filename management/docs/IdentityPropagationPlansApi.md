# \IdentityPropagationPlansApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlan**](IdentityPropagationPlansApi.md#CreatePlan) | **Post** /environments/{environmentID}/propagation/plans | CREATE Plan
[**DeletePlan**](IdentityPropagationPlansApi.md#DeletePlan) | **Delete** /environments/{environmentID}/propagation/plans/{planID} | DELETE Plan
[**ReadAllPlans**](IdentityPropagationPlansApi.md#ReadAllPlans) | **Get** /environments/{environmentID}/propagation/plans | READ All Plans
[**ReadOnePlan**](IdentityPropagationPlansApi.md#ReadOnePlan) | **Get** /environments/{environmentID}/propagation/plans/{planID} | READ One Plan
[**UpdatePlan**](IdentityPropagationPlansApi.md#UpdatePlan) | **Put** /environments/{environmentID}/propagation/plans/{planID} | UPDATE Plan



## CreatePlan

> IdentityPropagationPlan CreatePlan(ctx, environmentID).IdentityPropagationPlan(identityPropagationPlan).Execute()

CREATE Plan

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
    identityPropagationPlan := *openapiclient.NewIdentityPropagationPlan("Name_example") // IdentityPropagationPlan |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationPlansApi.CreatePlan(context.Background(), environmentID).IdentityPropagationPlan(identityPropagationPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationPlansApi.CreatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlan`: IdentityPropagationPlan
    fmt.Fprintf(os.Stdout, "Response from `IdentityPropagationPlansApi.CreatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **identityPropagationPlan** | [**IdentityPropagationPlan**](IdentityPropagationPlan.md) |  | 

### Return type

[**IdentityPropagationPlan**](IdentityPropagationPlan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlan(ctx, environmentID, planID).Accept(accept).Execute()

DELETE Plan

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
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IdentityPropagationPlansApi.DeletePlan(context.Background(), environmentID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationPlansApi.DeletePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

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


## ReadAllPlans

> EntityArrayPagedIterator ReadAllPlans(ctx, environmentID).Accept(accept).Execute()

READ All Plans

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
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationPlansApi.ReadAllPlans(context.Background(), environmentID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationPlansApi.ReadAllPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllPlans`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `IdentityPropagationPlansApi.ReadAllPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 

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


## ReadOnePlan

> IdentityPropagationPlan ReadOnePlan(ctx, environmentID, planID).Accept(accept).Execute()

READ One Plan

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
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationPlansApi.ReadOnePlan(context.Background(), environmentID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationPlansApi.ReadOnePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePlan`: IdentityPropagationPlan
    fmt.Fprintf(os.Stdout, "Response from `IdentityPropagationPlansApi.ReadOnePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**IdentityPropagationPlan**](IdentityPropagationPlan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> IdentityPropagationPlan UpdatePlan(ctx, environmentID, planID).IdentityPropagationPlan(identityPropagationPlan).Execute()

UPDATE Plan

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
    planID := "planID_example" // string | 
    identityPropagationPlan := *openapiclient.NewIdentityPropagationPlan("Name_example") // IdentityPropagationPlan |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationPlansApi.UpdatePlan(context.Background(), environmentID, planID).IdentityPropagationPlan(identityPropagationPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationPlansApi.UpdatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlan`: IdentityPropagationPlan
    fmt.Fprintf(os.Stdout, "Response from `IdentityPropagationPlansApi.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**planID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **identityPropagationPlan** | [**IdentityPropagationPlan**](IdentityPropagationPlan.md) |  | 

### Return type

[**IdentityPropagationPlan**](IdentityPropagationPlan.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

