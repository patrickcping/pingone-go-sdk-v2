# \SchemasApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAttribute**](SchemasApi.md#CreateAttribute) | **Post** /v1/environments/{environmentID}/schemas/{schemaID}/attributes | CREATE Attribute
[**DeleteAttribute**](SchemasApi.md#DeleteAttribute) | **Delete** /v1/environments/{environmentID}/schemas/{schemaID}/attributes/{attributeID} | DELETE Attribute
[**ReadAllSchemaAttributes**](SchemasApi.md#ReadAllSchemaAttributes) | **Get** /v1/environments/{environmentID}/schemas/{schemaID}/attributes | READ All (Schema) Attributes
[**ReadAllSchemas**](SchemasApi.md#ReadAllSchemas) | **Get** /v1/environments/{environmentID}/schemas | READ All Schemas
[**ReadOneAttribute**](SchemasApi.md#ReadOneAttribute) | **Get** /v1/environments/{environmentID}/schemas/{schemaID}/attributes/{attributeID} | READ One Attribute
[**ReadOneSchema**](SchemasApi.md#ReadOneSchema) | **Get** /v1/environments/{environmentID}/schemas/{schemaID} | READ One Schema
[**UpdateAttributePatch**](SchemasApi.md#UpdateAttributePatch) | **Patch** /v1/environments/{environmentID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Patch)
[**UpdateAttributePut**](SchemasApi.md#UpdateAttributePut) | **Put** /v1/environments/{environmentID}/schemas/{schemaID}/attributes/{attributeID} | UPDATE Attribute (Put)



## CreateAttribute

> SchemaAttribute CreateAttribute(ctx, environmentID, schemaID).SchemaAttribute(schemaAttribute).Execute()

CREATE Attribute

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
    schemaID := "schemaID_example" // string | 
    schemaAttribute := *openapiclient.NewSchemaAttribute(false, "Name_example", openapiclient.EnumSchemaAttributeType("STRING")) // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.CreateAttribute(context.Background(), environmentID, schemaID).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.CreateAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAttribute`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.CreateAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAttribute

> DeleteAttribute(ctx, environmentID, schemaID, attributeID).Execute()

DELETE Attribute

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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SchemasApi.DeleteAttribute(context.Background(), environmentID, schemaID, attributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.DeleteAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributeRequest struct via the builder pattern


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


## ReadAllSchemaAttributes

> EntityArray ReadAllSchemaAttributes(ctx, environmentID, schemaID).Execute()

READ All (Schema) Attributes

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
    schemaID := "schemaID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ReadAllSchemaAttributes(context.Background(), environmentID, schemaID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReadAllSchemaAttributes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSchemaAttributes`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReadAllSchemaAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSchemaAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllSchemas

> EntityArray ReadAllSchemas(ctx, environmentID).Execute()

READ All Schemas

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ReadAllSchemas(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReadAllSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSchemas`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReadAllSchemas`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneAttribute

> SchemaAttribute ReadOneAttribute(ctx, environmentID, schemaID, attributeID).Execute()

READ One Attribute

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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ReadOneAttribute(context.Background(), environmentID, schemaID, attributeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReadOneAttribute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAttribute`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReadOneAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneSchema

> Schema ReadOneSchema(ctx, environmentID, schemaID).Execute()

READ One Schema

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
    schemaID := "schemaID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.ReadOneSchema(context.Background(), environmentID, schemaID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.ReadOneSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneSchema`: Schema
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.ReadOneSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Schema**](Schema.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributePatch

> SchemaAttribute UpdateAttributePatch(ctx, environmentID, schemaID, attributeID).SchemaAttribute(schemaAttribute).Execute()

UPDATE Attribute (Patch)

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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    schemaAttribute := *openapiclient.NewSchemaAttribute(false, "Name_example", openapiclient.EnumSchemaAttributeType("STRING")) // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.UpdateAttributePatch(context.Background(), environmentID, schemaID, attributeID).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.UpdateAttributePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributePatch`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.UpdateAttributePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttributePut

> SchemaAttribute UpdateAttributePut(ctx, environmentID, schemaID, attributeID).SchemaAttribute(schemaAttribute).Execute()

UPDATE Attribute (Put)

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
    schemaID := "schemaID_example" // string | 
    attributeID := "attributeID_example" // string | 
    schemaAttribute := *openapiclient.NewSchemaAttribute(false, "Name_example", openapiclient.EnumSchemaAttributeType("STRING")) // SchemaAttribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchemasApi.UpdateAttributePut(context.Background(), environmentID, schemaID, attributeID).SchemaAttribute(schemaAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchemasApi.UpdateAttributePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAttributePut`: SchemaAttribute
    fmt.Fprintf(os.Stdout, "Response from `SchemasApi.UpdateAttributePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**schemaID** | **string** |  | 
**attributeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **schemaAttribute** | [**SchemaAttribute**](SchemaAttribute.md) |  | 

### Return type

[**SchemaAttribute**](SchemaAttribute.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

