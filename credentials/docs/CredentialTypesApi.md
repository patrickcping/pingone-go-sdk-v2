# \CredentialTypesApi

All URIs are relative to *https://api.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialType**](CredentialTypesApi.md#CreateCredentialType) | **Post** /environments/{environmentID}/credentialTypes | Create Credential Type
[**DeleteACredentialType**](CredentialTypesApi.md#DeleteACredentialType) | **Delete** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Delete a Credential Type
[**ReadAllCredentialTypes**](CredentialTypesApi.md#ReadAllCredentialTypes) | **Get** /environments/{environmentID}/credentialTypes | Read All Credential Types
[**ReadOneCredentialType**](CredentialTypesApi.md#ReadOneCredentialType) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Read One Credential Type
[**UpdateACredentialType**](CredentialTypesApi.md#UpdateACredentialType) | **Put** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Update a Credential Type



## CreateCredentialType

> CredentialType CreateCredentialType(ctx, environmentID).CredentialType(credentialType).Execute()

Create Credential Type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialType := *openapiclient.NewCredentialType("CardDesignTemplate_example", *openapiclient.NewCredentialTypeMetaData(), "Title_example") // CredentialType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypesApi.CreateCredentialType(context.Background(), environmentID).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.CreateCredentialType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialType`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.CreateCredentialType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialType** | [**CredentialType**](CredentialType.md) |  | 

### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteACredentialType

> DeleteACredentialType(ctx, environmentID, credentialTypeID).Execute()

Delete a Credential Type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CredentialTypesApi.DeleteACredentialType(context.Background(), environmentID, credentialTypeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.DeleteACredentialType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteACredentialTypeRequest struct via the builder pattern


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


## ReadAllCredentialTypes

> EntityArray ReadAllCredentialTypes(ctx, environmentID).Execute()

Read All Credential Types

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypesApi.ReadAllCredentialTypes(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.ReadAllCredentialTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllCredentialTypes`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.ReadAllCredentialTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllCredentialTypesRequest struct via the builder pattern


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


## ReadOneCredentialType

> CredentialType ReadOneCredentialType(ctx, environmentID, credentialTypeID).Execute()

Read One Credential Type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypesApi.ReadOneCredentialType(context.Background(), environmentID, credentialTypeID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.ReadOneCredentialType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneCredentialType`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.ReadOneCredentialType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneCredentialTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateACredentialType

> CredentialType UpdateACredentialType(ctx, environmentID, credentialTypeID).CredentialType(credentialType).Execute()

Update a Credential Type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialType := *openapiclient.NewCredentialType("CardDesignTemplate_example", *openapiclient.NewCredentialTypeMetaData(), "Title_example") // CredentialType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialTypesApi.UpdateACredentialType(context.Background(), environmentID, credentialTypeID).CredentialType(credentialType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.UpdateACredentialType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateACredentialType`: CredentialType
    fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.UpdateACredentialType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateACredentialTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialType** | [**CredentialType**](CredentialType.md) |  | 

### Return type

[**CredentialType**](CredentialType.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

