# \PropagationPlansApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPropagationPlansGet**](PropagationPlansApi.md#V1EnvironmentsEnvironmentIDPropagationPlansGet) | **Get** /v1/environments/{environmentID}/propagation/plans | READ All Plans
[**V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete**](PropagationPlansApi.md#V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete) | **Delete** /v1/environments/{environmentID}/propagation/plans/{planID} | DELETE Plan
[**V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet**](PropagationPlansApi.md#V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet) | **Get** /v1/environments/{environmentID}/propagation/plans/{planID} | READ One Plan
[**V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut**](PropagationPlansApi.md#V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut) | **Put** /v1/environments/{environmentID}/propagation/plans/{planID} | UPDATE Plan
[**V1EnvironmentsEnvironmentIDPropagationPlansPost**](PropagationPlansApi.md#V1EnvironmentsEnvironmentIDPropagationPlansPost) | **Post** /v1/environments/{environmentID}/propagation/plans | CREATE Plan



## V1EnvironmentsEnvironmentIDPropagationPlansGet

> V1EnvironmentsEnvironmentIDPropagationPlansGet(ctx, environmentID).Accept(accept).Execute()

READ All Plans

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansGet(context.Background(), environmentID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete

> V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete(ctx, environmentID, planID).Accept(accept).Execute()

DELETE Plan

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete(context.Background(), environmentID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansPlanIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet

> V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet(ctx, environmentID, planID).Accept(accept).Execute()

READ One Plan

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    planID := "planID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet(context.Background(), environmentID, planID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansPlanIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut

> V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut(ctx, environmentID, planID).Body(body).Execute()

UPDATE Plan

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    planID := "planID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut(context.Background(), environmentID, planID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansPlanIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDPropagationPlansPost

> V1EnvironmentsEnvironmentIDPropagationPlansPost(ctx, environmentID).Body(body).Execute()

CREATE Plan

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationPlansApi.V1EnvironmentsEnvironmentIDPropagationPlansPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

