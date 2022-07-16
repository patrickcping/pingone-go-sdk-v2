# \IdentityProviderManagementIdentityProviderAttributesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet**](IdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet) | **Get** /v1/environments/{environmentID}/identityProviders/{providerID}/attributes | READ All Identity Provider Attributes
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete**](IdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete) | **Delete** /v1/environments/{environmentID}/identityProviders/{providerID}/attributes/{idpAttrID} | DELETE Identity Provider Attribute
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet**](IdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet) | **Get** /v1/environments/{environmentID}/identityProviders/{providerID}/attributes/{idpAttrID} | READ One Identity Provider Attribute
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut**](IdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut) | **Put** /v1/environments/{environmentID}/identityProviders/{providerID}/attributes/{idpAttrID} | UPDATE Identity Provider Attribute
[**V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost**](IdentityProviderManagementIdentityProviderAttributesApi.md#V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost) | **Post** /v1/environments/{environmentID}/identityProviders/{providerID}/attributes | CREATE Identity Provider Attribute (SAML)



## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet(ctx, environmentID, providerID).Execute()

READ All Identity Provider Attributes

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
    providerID := "providerID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet(context.Background(), environmentID, providerID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete(ctx, environmentID, providerID, idpAttrID).Execute()

DELETE Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete(context.Background(), environmentID, providerID, idpAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet(ctx, environmentID, providerID, idpAttrID).Execute()

READ One Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet(context.Background(), environmentID, providerID, idpAttrID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut(ctx, environmentID, providerID, idpAttrID).Body(body).Execute()

UPDATE Identity Provider Attribute

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
    providerID := "providerID_example" // string | 
    idpAttrID := "idpAttrID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut(context.Background(), environmentID, providerID, idpAttrID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 
**idpAttrID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesIdpAttrIDPutRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost

> V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost(ctx, environmentID, providerID).Body(body).Execute()

CREATE Identity Provider Attribute (SAML)

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
    providerID := "providerID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost(context.Background(), environmentID, providerID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityProviderManagementIdentityProviderAttributesApi.V1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**providerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDIdentityProvidersProviderIDAttributesPostRequest struct via the builder pattern


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

