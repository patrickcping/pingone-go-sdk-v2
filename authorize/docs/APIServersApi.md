# \APIServersApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIServer**](APIServersApi.md#CreateAPIServer) | **Post** /environments/{environmentID}/apiServers | CREATE API Server
[**DeleteAPIServer**](APIServersApi.md#DeleteAPIServer) | **Delete** /environments/{environmentID}/apiServers/{apiServerID} | DELETE API Server
[**ReadAllAPIServers**](APIServersApi.md#ReadAllAPIServers) | **Get** /environments/{environmentID}/apiServers | READ All API Servers
[**ReadOneAPIServer**](APIServersApi.md#ReadOneAPIServer) | **Get** /environments/{environmentID}/apiServers/{apiServerID} | READ One API Server
[**UpdateAPIServer**](APIServersApi.md#UpdateAPIServer) | **Put** /environments/{environmentID}/apiServers/{apiServerID} | UPDATE API Server



## CreateAPIServer

> APIServer CreateAPIServer(ctx, environmentID).APIServer(aPIServer).Execute()

CREATE API Server

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
	environmentID := "environmentID_example" // string | 
	aPIServer := *openapiclient.NewAPIServer(*openapiclient.NewAPIServerAuthorizationServer(openapiclient.EnumAPIServerAuthorizationServerType("PINGONE_SSO")), []string{"BaseUrls_example"}, "Name_example") // APIServer |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIServersApi.CreateAPIServer(context.Background(), environmentID).APIServer(aPIServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.CreateAPIServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPIServer`: APIServer
	fmt.Fprintf(os.Stdout, "Response from `APIServersApi.CreateAPIServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIServer** | [**APIServer**](APIServer.md) |  | 

### Return type

[**APIServer**](APIServer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIServer

> DeleteAPIServer(ctx, environmentID, apiServerID).Execute()

DELETE API Server

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
	environmentID := "environmentID_example" // string | 
	apiServerID := "apiServerID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.APIServersApi.DeleteAPIServer(context.Background(), environmentID, apiServerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.DeleteAPIServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIServerRequest struct via the builder pattern


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


## ReadAllAPIServers

READ All API Servers

### Paged Response (Recommended)

> PagedIterator[ReadAllAPIServers200Response] ReadAllAPIServers(ctx, environmentID).Execute()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.APIServersApi.ReadAllAPIServers(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.ReadAllAPIServers``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadAllAPIServers` page iteration: ReadAllAPIServers200Response
		fmt.Fprintf(os.Stdout, "Response from `APIServersApi.ReadAllAPIServers` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadAllAPIServers200Response ReadAllAPIServers(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/authorize"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIServersApi.ReadAllAPIServers(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.ReadAllAPIServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAllAPIServers`: ReadAllAPIServers200Response
	fmt.Fprintf(os.Stdout, "Response from `APIServersApi.ReadAllAPIServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllAPIServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

Page Iterator: PagedIterator[[**ReadAllAPIServers200Response**](ReadAllAPIServers200Response.md)]

PagedIterator[ReadAllAPIServers200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadAllAPIServers200Response**](ReadAllAPIServers200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadAllAPIServers200Response**](ReadAllAPIServers200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneAPIServer

> APIServer ReadOneAPIServer(ctx, environmentID, apiServerID).Execute()

READ One API Server

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
	environmentID := "environmentID_example" // string | 
	apiServerID := "apiServerID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIServersApi.ReadOneAPIServer(context.Background(), environmentID, apiServerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.ReadOneAPIServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneAPIServer`: APIServer
	fmt.Fprintf(os.Stdout, "Response from `APIServersApi.ReadOneAPIServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAPIServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**APIServer**](APIServer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIServer

> APIServer UpdateAPIServer(ctx, environmentID, apiServerID).APIServer(aPIServer).Execute()

UPDATE API Server

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
	environmentID := "environmentID_example" // string | 
	apiServerID := "apiServerID_example" // string | 
	aPIServer := *openapiclient.NewAPIServer(*openapiclient.NewAPIServerAuthorizationServer(openapiclient.EnumAPIServerAuthorizationServerType("PINGONE_SSO")), []string{"BaseUrls_example"}, "Name_example") // APIServer |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIServersApi.UpdateAPIServer(context.Background(), environmentID, apiServerID).APIServer(aPIServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIServersApi.UpdateAPIServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAPIServer`: APIServer
	fmt.Fprintf(os.Stdout, "Response from `APIServersApi.UpdateAPIServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**apiServerID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPIServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aPIServer** | [**APIServer**](APIServer.md) |  | 

### Return type

[**APIServer**](APIServer.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

