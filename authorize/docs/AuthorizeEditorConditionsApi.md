# \AuthorizeEditorConditionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCondition**](AuthorizeEditorConditionsApi.md#CreateCondition) | **Post** /environments/{environmentID}/authorizationConditions | Create a Condition
[**DeleteCondition**](AuthorizeEditorConditionsApi.md#DeleteCondition) | **Delete** /environments/{environmentID}/authorizationConditions/{conditionID} | Delete a Condition
[**GetCondition**](AuthorizeEditorConditionsApi.md#GetCondition) | **Get** /environments/{environmentID}/authorizationConditions/{conditionID} | Get a Condition by ID
[**ListConditions**](AuthorizeEditorConditionsApi.md#ListConditions) | **Get** /environments/{environmentID}/authorizationConditions | List Conditions
[**TestCondition**](AuthorizeEditorConditionsApi.md#TestCondition) | **Post** /environments/{environmentID}/authorizationConditions/{conditionID} | Test a Condition by ID
[**UpdateCondition**](AuthorizeEditorConditionsApi.md#UpdateCondition) | **Put** /environments/{environmentID}/authorizationConditions/{conditionID} | Update a Condition



## CreateCondition

> AuthorizeEditorDataDefinitionsConditionDefinitionDTO CreateCondition(ctx, environmentID).AuthorizeEditorDataDefinitionsConditionDefinitionDTO(authorizeEditorDataDefinitionsConditionDefinitionDTO).Execute()

Create a Condition



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the condition.
    authorizeEditorDataDefinitionsConditionDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsConditionDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataConditionDTO("Type_example")) // AuthorizeEditorDataDefinitionsConditionDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConditionsApi.CreateCondition(context.Background(), environmentID).AuthorizeEditorDataDefinitionsConditionDefinitionDTO(authorizeEditorDataDefinitionsConditionDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.CreateCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCondition`: AuthorizeEditorDataDefinitionsConditionDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConditionsApi.CreateCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the condition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataDefinitionsConditionDefinitionDTO** | [**AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCondition

> DeleteCondition(ctx, environmentID, conditionID).Execute()

Delete a Condition



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the condition.
    conditionID := "conditionID_example" // string | The ID of the condition to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorConditionsApi.DeleteCondition(context.Background(), environmentID, conditionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.DeleteCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the condition. | 
**conditionID** | **string** | The ID of the condition to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConditionRequest struct via the builder pattern


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


## GetCondition

> AuthorizeEditorDataDefinitionsConditionDefinitionDTO GetCondition(ctx, environmentID, conditionID).Execute()

Get a Condition by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the condition.
    conditionID := "conditionID_example" // string | The ID of the condition to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConditionsApi.GetCondition(context.Background(), environmentID, conditionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.GetCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCondition`: AuthorizeEditorDataDefinitionsConditionDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConditionsApi.GetCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the condition. | 
**conditionID** | **string** | The ID of the condition to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConditions

> ListConditions200Response ListConditions(ctx, environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()

List Conditions



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list conditions.
    filter := "filter_example" // string | The parent filter that may be used to list the root conditions or children of a condition. The following filter options can be applied to conditions:     * `parent not pr` (root conditions)     * `parent eq id` (conditions with parent id equal to id) (optional)
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConditionsApi.ListConditions(context.Background(), environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.ListConditions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConditions`: ListConditions200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConditionsApi.ListConditions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list conditions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConditionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The parent filter that may be used to list the root conditions or children of a condition. The following filter options can be applied to conditions:     * &#x60;parent not pr&#x60; (root conditions)     * &#x60;parent eq id&#x60; (conditions with parent id equal to id) | 
 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListConditions200Response**](ListConditions200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCondition

> AuthorizeEditorDataEntityTestResponseDTO TestCondition(ctx, environmentID, conditionID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()

Test a Condition by ID



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
    conditionID := "conditionID_example" // string | 
    authorizeEditorDataEntityTestRequestDTO := *openapiclient.NewAuthorizeEditorDataEntityTestRequestDTO([]openapiclient.AuthorizeEditorDataEntityTestingParameterDTO{*openapiclient.NewAuthorizeEditorDataEntityTestingParameterDTO("Key_example", "Value_example")}) // AuthorizeEditorDataEntityTestRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConditionsApi.TestCondition(context.Background(), environmentID, conditionID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.TestCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestCondition`: AuthorizeEditorDataEntityTestResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConditionsApi.TestCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**conditionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestConditionRequest struct via the builder pattern


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


## UpdateCondition

> AuthorizeEditorDataDefinitionsConditionDefinitionDTO UpdateCondition(ctx, environmentID, conditionID).AuthorizeEditorDataDefinitionsConditionDefinitionDTO(authorizeEditorDataDefinitionsConditionDefinitionDTO).Execute()

Update a Condition



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the condition.
    conditionID := "conditionID_example" // string | The ID of the condition to be updated
    authorizeEditorDataDefinitionsConditionDefinitionDTO := *openapiclient.NewAuthorizeEditorDataDefinitionsConditionDefinitionDTO("Name_example", *openapiclient.NewAuthorizeEditorDataConditionDTO("Type_example")) // AuthorizeEditorDataDefinitionsConditionDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorConditionsApi.UpdateCondition(context.Background(), environmentID, conditionID).AuthorizeEditorDataDefinitionsConditionDefinitionDTO(authorizeEditorDataDefinitionsConditionDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorConditionsApi.UpdateCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCondition`: AuthorizeEditorDataDefinitionsConditionDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorConditionsApi.UpdateCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the condition. | 
**conditionID** | **string** | The ID of the condition to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataDefinitionsConditionDefinitionDTO** | [**AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsConditionDefinitionDTO**](AuthorizeEditorDataDefinitionsConditionDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
