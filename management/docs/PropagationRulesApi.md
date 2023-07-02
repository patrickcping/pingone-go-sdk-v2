# \PropagationRulesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet) | **Get** /environments/{environmentID}/propagation/plans/{planID}/rules | READ One Plan&#39;s Rules
[**EnvironmentsEnvironmentIDPropagationRulesGet**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationRulesGet) | **Get** /environments/{environmentID}/propagation/rules | READ All Rules
[**EnvironmentsEnvironmentIDPropagationRulesPost**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationRulesPost) | **Post** /environments/{environmentID}/propagation/rules | CREATE Rule
[**EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete) | **Delete** /environments/{environmentID}/propagation/rules/{ruleID} | DELETE Rule
[**EnvironmentsEnvironmentIDPropagationRulesRuleIDGet**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationRulesRuleIDGet) | **Get** /environments/{environmentID}/propagation/rules/{ruleID} | READ One Rule
[**EnvironmentsEnvironmentIDPropagationRulesStoreIDPut**](PropagationRulesApi.md#EnvironmentsEnvironmentIDPropagationRulesStoreIDPut) | **Put** /environments/{environmentID}/propagation/rules/{storeID} | UPDATE Rule



## EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet

> EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet(ctx, environmentID, planID).Accept(accept).ContentType(contentType).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet(context.Background(), environmentID, planID).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationPlansPlanIDRulesGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesGet

> EnvironmentsEnvironmentIDPropagationRulesGet(ctx, environmentID).Accept(accept).Authorization(authorization).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesGet(context.Background(), environmentID).Accept(accept).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesPost

> EnvironmentsEnvironmentIDPropagationRulesPost(ctx, environmentID).Body(body).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesPostRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete

> EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete(ctx, environmentID, ruleID).Accept(accept).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete(context.Background(), environmentID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesRuleIDDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesRuleIDGet

> EnvironmentsEnvironmentIDPropagationRulesRuleIDGet(ctx, environmentID, ruleID).Accept(accept).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDGet(context.Background(), environmentID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesRuleIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesStoreIDPut

> EnvironmentsEnvironmentIDPropagationRulesStoreIDPut(ctx, environmentID, storeID).Body(body).Execute()

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
    r, err := apiClient.PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesStoreIDPut(context.Background(), environmentID, storeID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationRulesApi.EnvironmentsEnvironmentIDPropagationRulesStoreIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesStoreIDPutRequest struct via the builder pattern


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

