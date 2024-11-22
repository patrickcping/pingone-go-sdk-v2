# \CredentialTypesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialType**](CredentialTypesApi.md#CreateCredentialType) | **Post** /environments/{environmentID}/credentialTypes | Create Credential Type
[**DeleteCredentialType**](CredentialTypesApi.md#DeleteCredentialType) | **Delete** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Delete a Credential Type
[**ReadAllCredentialTypes**](CredentialTypesApi.md#ReadAllCredentialTypes) | **Get** /environments/{environmentID}/credentialTypes | Read All Credential Types
[**ReadOneCredentialType**](CredentialTypesApi.md#ReadOneCredentialType) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Read One Credential Type
[**UpdateCredentialType**](CredentialTypesApi.md#UpdateCredentialType) | **Put** /environments/{environmentID}/credentialTypes/{credentialTypeID} | Update a Credential Type



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


## DeleteCredentialType

> DeleteCredentialType(ctx, environmentID, credentialTypeID).Execute()

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
	r, err := apiClient.CredentialTypesApi.DeleteCredentialType(context.Background(), environmentID, credentialTypeID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.DeleteCredentialType``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCredentialTypeRequest struct via the builder pattern


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

Read All Credential Types

### Paged Response (Recommended)

> PagedIterator[ReadAllCredentialTypes200Response] ReadAllCredentialTypes(ctx, environmentID).Execute()

#### Example

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
	pagedIterator := apiClient.CredentialTypesApi.ReadAllCredentialTypes(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.ReadAllCredentialTypes``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadAllCredentialTypes` page iteration: ReadAllCredentialTypes200Response
		fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.ReadAllCredentialTypes` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadAllCredentialTypes200Response ReadAllCredentialTypes(ctx, environmentID).ExecuteInitialPage()

#### Example

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
	resp, r, err := apiClient.CredentialTypesApi.ReadAllCredentialTypes(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.ReadAllCredentialTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAllCredentialTypes`: ReadAllCredentialTypes200Response
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

Page Iterator: PagedIterator[[**ReadAllCredentialTypes200Response**](ReadAllCredentialTypes200Response.md)]

PagedIterator[ReadAllCredentialTypes200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadAllCredentialTypes200Response**](ReadAllCredentialTypes200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadAllCredentialTypes200Response**](ReadAllCredentialTypes200Response.md)

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


## UpdateCredentialType

> CredentialType UpdateCredentialType(ctx, environmentID, credentialTypeID).CredentialType(credentialType).Execute()

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
	resp, r, err := apiClient.CredentialTypesApi.UpdateCredentialType(context.Background(), environmentID, credentialTypeID).CredentialType(credentialType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CredentialTypesApi.UpdateCredentialType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCredentialType`: CredentialType
	fmt.Fprintf(os.Stdout, "Response from `CredentialTypesApi.UpdateCredentialType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialTypeRequest struct via the builder pattern


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

