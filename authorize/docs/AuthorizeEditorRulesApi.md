# \AuthorizeEditorRulesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](AuthorizeEditorRulesApi.md#CreateRule) | **Post** /environments/{environmentID}/authorizationRules | Create a Rule
[**DeleteRule**](AuthorizeEditorRulesApi.md#DeleteRule) | **Delete** /environments/{environmentID}/authorizationRules/{ruleID} | Delete a Rule
[**GetRule**](AuthorizeEditorRulesApi.md#GetRule) | **Get** /environments/{environmentID}/authorizationRules/{ruleID} | Get a Rule by ID
[**ListRules**](AuthorizeEditorRulesApi.md#ListRules) | **Get** /environments/{environmentID}/authorizationRules | List Rules
[**TestRule**](AuthorizeEditorRulesApi.md#TestRule) | **Post** /environments/{environmentID}/authorizationRules/{ruleID} | Test a Rule by ID
[**UpdateRule**](AuthorizeEditorRulesApi.md#UpdateRule) | **Put** /environments/{environmentID}/authorizationRules/{ruleID} | Update a Rule



## CreateRule

> AuthorizeEditorDataRulesReferenceableRuleDTO CreateRule(ctx, environmentID).AuthorizeEditorDataRulesRuleDTO(authorizeEditorDataRulesRuleDTO).Accept(accept).Execute()

Create a Rule



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the rule.
    authorizeEditorDataRulesRuleDTO := *openapiclient.NewAuthorizeEditorDataRulesRuleDTO("Name_example", openapiclient.AuthorizeEditorDataRulesRuleDTO_effectSettings{AuthorizeEditorDataRulesEffectSettingsConditionalDenyElsePermitDTO: openapiclient.NewAuthorizeEditorDataRulesEffectSettingsConditionalDenyElsePermitDTO(*openapiclient.NewAuthorizeEditorDataConditionDTO("Type_example"), "Type_example")}) // AuthorizeEditorDataRulesRuleDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorRulesApi.CreateRule(context.Background(), environmentID).AuthorizeEditorDataRulesRuleDTO(authorizeEditorDataRulesRuleDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: AuthorizeEditorDataRulesReferenceableRuleDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorRulesApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataRulesRuleDTO** | [**AuthorizeEditorDataRulesRuleDTO**](AuthorizeEditorDataRulesRuleDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, environmentID, ruleID).Execute()

Delete a Rule



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the rule.
    ruleID := "ruleID_example" // string | The ID of the rule to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorRulesApi.DeleteRule(context.Background(), environmentID, ruleID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the rule. | 
**ruleID** | **string** | The ID of the rule to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetRule

> AuthorizeEditorDataRulesReferenceableRuleDTO GetRule(ctx, environmentID, ruleID).Accept(accept).Execute()

Get a Rule by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the rule.
    ruleID := "ruleID_example" // string | The ID of the rule to retrieve
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorRulesApi.GetRule(context.Background(), environmentID, ruleID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: AuthorizeEditorDataRulesReferenceableRuleDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorRulesApi.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the rule. | 
**ruleID** | **string** | The ID of the rule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRules

> ListRules200Response ListRules(ctx, environmentID).Limit(limit).Cursor(cursor).Execute()

List Rules



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list Rules
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorRulesApi.ListRules(context.Background(), environmentID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.ListRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRules`: ListRules200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorRulesApi.ListRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list Rules | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListRules200Response**](ListRules200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestRule

> AuthorizeEditorDataEntityTestResponseDTO TestRule(ctx, environmentID, ruleID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()

Test a Rule by ID



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
    ruleID := "ruleID_example" // string | 
    authorizeEditorDataEntityTestRequestDTO := *openapiclient.NewAuthorizeEditorDataEntityTestRequestDTO([]openapiclient.AuthorizeEditorDataEntityTestingParameterDTO{*openapiclient.NewAuthorizeEditorDataEntityTestingParameterDTO("Key_example", "Value_example")}) // AuthorizeEditorDataEntityTestRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorRulesApi.TestRule(context.Background(), environmentID, ruleID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.TestRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestRule`: AuthorizeEditorDataEntityTestResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorRulesApi.TestRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**ruleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRuleRequest struct via the builder pattern


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


## UpdateRule

> AuthorizeEditorDataRulesReferenceableRuleDTO UpdateRule(ctx, environmentID, ruleID).AuthorizeEditorDataRulesReferenceableRuleDTO(authorizeEditorDataRulesReferenceableRuleDTO).Accept(accept).Execute()

Update a Rule



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the rule.
    ruleID := "ruleID_example" // string | The ID of the rule to be updated
    authorizeEditorDataRulesReferenceableRuleDTO := *openapiclient.NewAuthorizeEditorDataRulesReferenceableRuleDTO(*openapiclient.NewAuthorizeEditorDataRulesEffectSettingsDTO("Type_example"), "Version_example") // AuthorizeEditorDataRulesReferenceableRuleDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorRulesApi.UpdateRule(context.Background(), environmentID, ruleID).AuthorizeEditorDataRulesReferenceableRuleDTO(authorizeEditorDataRulesReferenceableRuleDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorRulesApi.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: AuthorizeEditorDataRulesReferenceableRuleDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorRulesApi.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the rule. | 
**ruleID** | **string** | The ID of the rule to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataRulesReferenceableRuleDTO** | [**AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataRulesReferenceableRuleDTO**](AuthorizeEditorDataRulesReferenceableRuleDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

