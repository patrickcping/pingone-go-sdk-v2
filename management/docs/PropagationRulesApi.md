# \PropagationRulesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet) | **Get** /v1/environments/{environmentID}/propagation/plans/{planID}/rules | READ One Plan&#39;s Rules
[**V1EnvironmentsEnvironmentIDPropagationRulesGet**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationRulesGet) | **Get** /v1/environments/{environmentID}/propagation/rules | READ All Rules
[**V1EnvironmentsEnvironmentIDPropagationRulesPost**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationRulesPost) | **Post** /v1/environments/{environmentID}/propagation/rules | CREATE Rule
[**V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete) | **Delete** /v1/environments/{environmentID}/propagation/rules/{ruleID} | DELETE Rule
[**V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet) | **Get** /v1/environments/{environmentID}/propagation/rules/{ruleID} | READ One Rule
[**V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut**](PropagationRulesApi.md#V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut) | **Put** /v1/environments/{environmentID}/propagation/rules/{storeID} | UPDATE Rule



## V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet

> V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet(ctx, environmentID, planID).Accept(accept).ContentType(contentType).Execute()

READ One Plan's Rules

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
    contentType := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet(context.Background(), environmentID, planID).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 
 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDPropagationRulesGet

> V1EnvironmentsEnvironmentIDPropagationRulesGet(ctx, environmentID).Accept(accept).Authorization(authorization).Execute()

READ All Rules

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
    authorization := "authorization_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesGet(context.Background(), environmentID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** |  | 
 **authorization** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDPropagationRulesPost

> V1EnvironmentsEnvironmentIDPropagationRulesPost(ctx, environmentID).Body(body).Execute()

CREATE Rule

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesPostRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete

> V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete(ctx, environmentID, ruleID).Accept(accept).Execute()

DELETE Rule

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
    ruleID := "ruleID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete(context.Background(), environmentID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**ruleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet

> V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet(ctx, environmentID, ruleID).Accept(accept).Execute()

READ One Rule

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
    ruleID := "ruleID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet(context.Background(), environmentID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**ruleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut

> V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut(ctx, environmentID, storeID).Body(body).Execute()

UPDATE Rule

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
    storeID := "storeID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut(context.Background(), environmentID, storeID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.V1EnvironmentsEnvironmentIDPropagationRulesStoreIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**storeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesStoreIDPutRequest struct via the builder pattern


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

