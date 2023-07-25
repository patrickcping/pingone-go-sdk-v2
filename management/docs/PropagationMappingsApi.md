# \PropagationMappingsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete**](PropagationMappingsApi.md#EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete) | **Delete** /environments/{environmentID}/propagation/mapping/{mappingID} | DELETE Mapping
[**EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet**](PropagationMappingsApi.md#EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet) | **Get** /environments/{environmentID}/propagation/mappings/{mappingID} | READ One Mapping
[**EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut**](PropagationMappingsApi.md#EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut) | **Put** /environments/{environmentID}/propagation/mappings/{mappingID} | UPDATE Mapping
[**EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet**](PropagationMappingsApi.md#EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet) | **Get** /environments/{environmentID}/propagation/rules/{ruleID}/mappings | READ One Rule  Mapping
[**EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost**](PropagationMappingsApi.md#EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost) | **Post** /environments/{environmentID}/propagation/rules/{ruleID}/mappings | CREATE Rule Mapping



## EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete

> EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete(ctx, environmentID, mappingID).Accept(accept).Execute()

DELETE Mapping

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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete(context.Background(), environmentID, mappingID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet

> EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet(ctx, environmentID, mappingID).Accept(accept).Execute()

READ One Mapping

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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet(context.Background(), environmentID, mappingID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut

> EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut(ctx, environmentID, mappingID).Body(body).Execute()

UPDATE Mapping

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
    mappingID := "mappingID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut(context.Background(), environmentID, mappingID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**mappingID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet

> EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet(ctx, environmentID, ruleID).Accept(accept).ContentType(contentType).Execute()

READ One Rule  Mapping

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
    contentType := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet(context.Background(), environmentID, ruleID).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost

> EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost(ctx, environmentID, ruleID).Body(body).Execute()

CREATE Rule Mapping

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost(context.Background(), environmentID, ruleID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PropagationMappingsApi.EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest struct via the builder pattern


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

