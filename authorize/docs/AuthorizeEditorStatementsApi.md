# \AuthorizeEditorStatementsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStatement**](AuthorizeEditorStatementsApi.md#CreateStatement) | **Post** /environments/{environmentID}/authorizationStatements | Create a Statement
[**DeleteStatement**](AuthorizeEditorStatementsApi.md#DeleteStatement) | **Delete** /environments/{environmentID}/authorizationStatements/{statementID} | Delete a Statement
[**GetStatement**](AuthorizeEditorStatementsApi.md#GetStatement) | **Get** /environments/{environmentID}/authorizationStatements/{statementID} | Get a Statement by ID
[**ListStatements**](AuthorizeEditorStatementsApi.md#ListStatements) | **Get** /environments/{environmentID}/authorizationStatements | List Statements
[**UpdateStatement**](AuthorizeEditorStatementsApi.md#UpdateStatement) | **Put** /environments/{environmentID}/authorizationStatements/{statementID} | Update a Statement



## CreateStatement

> AuthorizeEditorDataStatementsReferenceableStatementDTO CreateStatement(ctx, environmentID).AuthorizeEditorDataStatementsStatementDTO(authorizeEditorDataStatementsStatementDTO).Accept(accept).Execute()

Create a Statement



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the statement.
    authorizeEditorDataStatementsStatementDTO := *openapiclient.NewAuthorizeEditorDataStatementsStatementDTO("Name_example", "Code_example", "AppliesTo_example", "AppliesIf_example", "Payload_example", []openapiclient.AuthorizeEditorDataReferenceObjectDTO{*openapiclient.NewAuthorizeEditorDataReferenceObjectDTO("Id_example")}) // AuthorizeEditorDataStatementsStatementDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorStatementsApi.CreateStatement(context.Background(), environmentID).AuthorizeEditorDataStatementsStatementDTO(authorizeEditorDataStatementsStatementDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorStatementsApi.CreateStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStatement`: AuthorizeEditorDataStatementsReferenceableStatementDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorStatementsApi.CreateStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the statement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataStatementsStatementDTO** | [**AuthorizeEditorDataStatementsStatementDTO**](AuthorizeEditorDataStatementsStatementDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStatement

> DeleteStatement(ctx, environmentID, statementID).Execute()

Delete a Statement



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the statement.
    statementID := "statementID_example" // string | The ID of the statement to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorStatementsApi.DeleteStatement(context.Background(), environmentID, statementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorStatementsApi.DeleteStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the statement. | 
**statementID** | **string** | The ID of the statement to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStatementRequest struct via the builder pattern


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


## GetStatement

> AuthorizeEditorDataStatementsReferenceableStatementDTO GetStatement(ctx, environmentID, statementID).Accept(accept).Execute()

Get a Statement by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the statement.
    statementID := "statementID_example" // string | The ID of the statement to retrieve
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorStatementsApi.GetStatement(context.Background(), environmentID, statementID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorStatementsApi.GetStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatement`: AuthorizeEditorDataStatementsReferenceableStatementDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorStatementsApi.GetStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the statement. | 
**statementID** | **string** | The ID of the statement to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStatements

> ListStatements200Response ListStatements(ctx, environmentID).Limit(limit).Cursor(cursor).Execute()

List Statements



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list statements.
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorStatementsApi.ListStatements(context.Background(), environmentID).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorStatementsApi.ListStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStatements`: ListStatements200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorStatementsApi.ListStatements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list statements. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStatementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListStatements200Response**](ListStatements200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatement

> AuthorizeEditorDataStatementsReferenceableStatementDTO UpdateStatement(ctx, environmentID, statementID).AuthorizeEditorDataStatementsReferenceableStatementDTO(authorizeEditorDataStatementsReferenceableStatementDTO).Accept(accept).Execute()

Update a Statement



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the statement.
    statementID := "statementID_example" // string | The ID of the statement to be updated
    authorizeEditorDataStatementsReferenceableStatementDTO := *openapiclient.NewAuthorizeEditorDataStatementsReferenceableStatementDTO("Id_example", "Name_example", "Code_example", "AppliesTo_example", "AppliesIf_example", "Payload_example", []openapiclient.AuthorizeEditorDataReferenceObjectDTO{*openapiclient.NewAuthorizeEditorDataReferenceObjectDTO("Id_example")}, "Version_example") // AuthorizeEditorDataStatementsReferenceableStatementDTO | 
    accept := "accept_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorStatementsApi.UpdateStatement(context.Background(), environmentID, statementID).AuthorizeEditorDataStatementsReferenceableStatementDTO(authorizeEditorDataStatementsReferenceableStatementDTO).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorStatementsApi.UpdateStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStatement`: AuthorizeEditorDataStatementsReferenceableStatementDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorStatementsApi.UpdateStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the statement. | 
**statementID** | **string** | The ID of the statement to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataStatementsReferenceableStatementDTO** | [**AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md) |  | 
 **accept** | **string** |  | 

### Return type

[**AuthorizeEditorDataStatementsReferenceableStatementDTO**](AuthorizeEditorDataStatementsReferenceableStatementDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
