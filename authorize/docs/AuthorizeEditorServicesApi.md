# \AuthorizeEditorServicesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateService**](AuthorizeEditorServicesApi.md#CreateService) | **Post** /environments/{environmentID}/authorizationServices | Create a Service
[**DeleteService**](AuthorizeEditorServicesApi.md#DeleteService) | **Delete** /environments/{environmentID}/authorizationServices/{serviceID} | Delete a Service
[**GetService**](AuthorizeEditorServicesApi.md#GetService) | **Get** /environments/{environmentID}/authorizationServices/{serviceID} | Get a Service by ID
[**ListServices**](AuthorizeEditorServicesApi.md#ListServices) | **Get** /environments/{environmentID}/authorizationServices | List Services
[**TestService**](AuthorizeEditorServicesApi.md#TestService) | **Post** /environments/{environmentID}/authorizationServices/{serviceID} | Test a Service by ID
[**UpdateService**](AuthorizeEditorServicesApi.md#UpdateService) | **Put** /environments/{environmentID}/authorizationServices/{serviceID} | Update a Service



## CreateService

> CreateService201Response CreateService(ctx, environmentID).AuthorizeEditorDataDefinitionsServiceDefinitionDTO(authorizeEditorDataDefinitionsServiceDefinitionDTO).Execute()

Create a Service



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to create the service.
    authorizeEditorDataDefinitionsServiceDefinitionDTO := openapiclient.AuthorizeEditorDataDefinitionsServiceDefinitionDTO{AuthorizeEditorDataServicesConnectorServiceDefinitionDTO: openapiclient.NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO("Name_example", openapiclient.EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType("NONE"), *openapiclient.NewAuthorizeEditorDataValueTypeDTO(openapiclient.EnumAuthorizeEditorDataValueTypeDTO("BOOLEAN")), *openapiclient.NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO(openapiclient.EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel("AUTHORIZE"), openapiclient.EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode("P1_RISK"), "Capability_example", []openapiclient.AuthorizeEditorDataInputMappingDTO{openapiclient.AuthorizeEditorDataInputMappingDTO{AuthorizeEditorDataInputMappingsAttributeInputMappingDTO: openapiclient.NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTO("Property_example", openapiclient.EnumAuthorizeEditorDataInputMappingDTOType("ATTRIBUTE"), *openapiclient.NewAuthorizeEditorDataReferenceObjectDTO("Id_example"))}}))} // AuthorizeEditorDataDefinitionsServiceDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorServicesApi.CreateService(context.Background(), environmentID).AuthorizeEditorDataDefinitionsServiceDefinitionDTO(authorizeEditorDataDefinitionsServiceDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.CreateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateService`: CreateService201Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorServicesApi.CreateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to create the service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizeEditorDataDefinitionsServiceDefinitionDTO** | [**AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md) |  | 

### Return type

[**CreateService201Response**](CreateService201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteService

> DeleteService(ctx, environmentID, serviceID).Execute()

Delete a Service



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to delete the service.
    serviceID := "serviceID_example" // string | The ID of the service to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuthorizeEditorServicesApi.DeleteService(context.Background(), environmentID, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.DeleteService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to delete the service. | 
**serviceID** | **string** | The ID of the service to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceRequest struct via the builder pattern


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


## GetService

> AuthorizeEditorDataDefinitionsServiceDefinitionDTO GetService(ctx, environmentID, serviceID).Execute()

Get a Service by ID



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to get the service.
    serviceID := "serviceID_example" // string | The ID of the service to retrieve

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorServicesApi.GetService(context.Background(), environmentID, serviceID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.GetService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetService`: AuthorizeEditorDataDefinitionsServiceDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorServicesApi.GetService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to get the service. | 
**serviceID** | **string** | The ID of the service to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServices

> ListServices200Response ListServices(ctx, environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()

List Services



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to list Services
    filter := "filter_example" // string | The parent filter that may be used to list the root services or children of a service. The following filter options can be applied to services:     * `parent not pr` (root services)     * `parent eq id` (services with parent id equal to id) (optional)
    limit := int32(56) // int32 | The maximum number of resources to return in the page (optional) (default to 25)
    cursor := "cursor_example" // string | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the 'next' link in the response body (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorServicesApi.ListServices(context.Background(), environmentID).Filter(filter).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.ListServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServices`: ListServices200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorServicesApi.ListServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to list Services | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | The parent filter that may be used to list the root services or children of a service. The following filter options can be applied to services:     * &#x60;parent not pr&#x60; (root services)     * &#x60;parent eq id&#x60; (services with parent id equal to id) | 
 **limit** | **int32** | The maximum number of resources to return in the page | [default to 25]
 **cursor** | **string** | An optional cursor that may be provided to start paging from a certain location. This cursor will be included in the &#39;next&#39; link in the response body | 

### Return type

[**ListServices200Response**](ListServices200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestService

> AuthorizeEditorDataEntityTestResponseDTO TestService(ctx, environmentID, serviceID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()

Test a Service by ID



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
    serviceID := "serviceID_example" // string | 
    authorizeEditorDataEntityTestRequestDTO := *openapiclient.NewAuthorizeEditorDataEntityTestRequestDTO([]openapiclient.AuthorizeEditorDataEntityTestingParameterDTO{*openapiclient.NewAuthorizeEditorDataEntityTestingParameterDTO("Key_example", "Value_example")}) // AuthorizeEditorDataEntityTestRequestDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorServicesApi.TestService(context.Background(), environmentID, serviceID).AuthorizeEditorDataEntityTestRequestDTO(authorizeEditorDataEntityTestRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.TestService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestService`: AuthorizeEditorDataEntityTestResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorServicesApi.TestService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**serviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestServiceRequest struct via the builder pattern


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


## UpdateService

> AuthorizeEditorDataDefinitionsServiceDefinitionDTO UpdateService(ctx, environmentID, serviceID).AuthorizeEditorDataDefinitionsServiceDefinitionDTO(authorizeEditorDataDefinitionsServiceDefinitionDTO).Execute()

Update a Service



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
    environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment from which to update the service.
    serviceID := "serviceID_example" // string | The ID of the service to be updated
    authorizeEditorDataDefinitionsServiceDefinitionDTO := openapiclient.AuthorizeEditorDataDefinitionsServiceDefinitionDTO{AuthorizeEditorDataServicesConnectorServiceDefinitionDTO: openapiclient.NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO("Name_example", openapiclient.EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType("NONE"), *openapiclient.NewAuthorizeEditorDataValueTypeDTO(openapiclient.EnumAuthorizeEditorDataValueTypeDTO("BOOLEAN")), *openapiclient.NewAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO(openapiclient.EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel("AUTHORIZE"), openapiclient.EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOCode("P1_RISK"), "Capability_example", []openapiclient.AuthorizeEditorDataInputMappingDTO{openapiclient.AuthorizeEditorDataInputMappingDTO{AuthorizeEditorDataInputMappingsAttributeInputMappingDTO: openapiclient.NewAuthorizeEditorDataInputMappingsAttributeInputMappingDTO("Property_example", openapiclient.EnumAuthorizeEditorDataInputMappingDTOType("ATTRIBUTE"), *openapiclient.NewAuthorizeEditorDataReferenceObjectDTO("Id_example"))}}))} // AuthorizeEditorDataDefinitionsServiceDefinitionDTO | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizeEditorServicesApi.UpdateService(context.Background(), environmentID, serviceID).AuthorizeEditorDataDefinitionsServiceDefinitionDTO(authorizeEditorDataDefinitionsServiceDefinitionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizeEditorServicesApi.UpdateService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateService`: AuthorizeEditorDataDefinitionsServiceDefinitionDTO
    fmt.Fprintf(os.Stdout, "Response from `AuthorizeEditorServicesApi.UpdateService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** | The ID of the environment from which to update the service. | 
**serviceID** | **string** | The ID of the service to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizeEditorDataDefinitionsServiceDefinitionDTO** | [**AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md) |  | 

### Return type

[**AuthorizeEditorDataDefinitionsServiceDefinitionDTO**](AuthorizeEditorDataDefinitionsServiceDefinitionDTO.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

