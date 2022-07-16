# \IdentityPropagationProvisioningPropagationMappingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete**](IdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete) | **Delete** /v1/environments/{environmentID}/propagation/mapping/{mappingID} | DELETE Mapping
[**V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet**](IdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet) | **Get** /v1/environments/{environmentID}/propagation/mappings/{mappingID} | READ One Mapping
[**V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut**](IdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut) | **Put** /v1/environments/{environmentID}/propagation/mappings/{mappingID} | UPDATE Mapping
[**V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet**](IdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet) | **Get** /v1/environments/{environmentID}/propagation/rules/{ruleID}/mappings | READ One Rule  Mapping
[**V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost**](IdentityPropagationProvisioningPropagationMappingsApi.md#V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost) | **Post** /v1/environments/{environmentID}/propagation/rules/{ruleID}/mappings | CREATE Rule Mapping



## V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete

> V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete(ctx, environmentID, mappingID).Accept(accept).Execute()

DELETE Mapping

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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete(context.Background(), environmentID, mappingID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingMappingIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationMappingMappingIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet

> V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet(ctx, environmentID, mappingID).Accept(accept).Execute()

READ One Mapping

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
    mappingID := "mappingID_example" // string | 
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet(context.Background(), environmentID, mappingID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut

> V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut(ctx, environmentID, mappingID).Body(body).Execute()

UPDATE Mapping

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
    mappingID := "mappingID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut(context.Background(), environmentID, mappingID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationMappingsMappingIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet

> V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet(ctx, environmentID, ruleID).Accept(accept).ContentType(contentType).Execute()

READ One Rule  Mapping

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
    ruleID := "ruleID_example" // string | 
    accept := "application/json" // string |  (optional)
    contentType := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet(context.Background(), environmentID, ruleID).Accept(accept).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost

> V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost(ctx, environmentID, ruleID).Body(body).Execute()

CREATE Rule Mapping

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
    ruleID := "ruleID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost(context.Background(), environmentID, ruleID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPropagationProvisioningPropagationMappingsApi.V1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDPropagationRulesRuleIDMappingsPostRequest struct via the builder pattern


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

