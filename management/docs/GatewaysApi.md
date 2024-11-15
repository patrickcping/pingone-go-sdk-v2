# \GatewaysApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGateway**](GatewaysApi.md#CreateGateway) | **Post** /environments/{environmentID}/gateways | CREATE Gateway
[**DeleteGateway**](GatewaysApi.md#DeleteGateway) | **Delete** /environments/{environmentID}/gateways/{gatewayID} | DELETE Gateway
[**ReadAllGateways**](GatewaysApi.md#ReadAllGateways) | **Get** /environments/{environmentID}/gateways | READ All Gateways
[**ReadOneGateway**](GatewaysApi.md#ReadOneGateway) | **Get** /environments/{environmentID}/gateways/{gatewayID} | READ One Gateway
[**UpdateGateway**](GatewaysApi.md#UpdateGateway) | **Put** /environments/{environmentID}/gateways/{gatewayID} | UPDATE Gateway



## CreateGateway

> CreateGateway201Response CreateGateway(ctx, environmentID).CreateGatewayRequest(createGatewayRequest).Execute()

CREATE Gateway

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
    createGatewayRequest := openapiclient.createGateway_request{Gateway: openapiclient.NewGateway("Name_example", openapiclient.EnumGatewayType("LDAP"), false)} // CreateGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.CreateGateway(context.Background(), environmentID).CreateGatewayRequest(createGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.CreateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.CreateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGatewayRequest** | [**CreateGatewayRequest**](CreateGatewayRequest.md) |  | 

### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGateway

> DeleteGateway(ctx, environmentID, gatewayID).Execute()

DELETE Gateway

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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GatewaysApi.DeleteGateway(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.DeleteGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayRequest struct via the builder pattern


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


## ReadAllGateways

READ All Gateways

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllGateways(ctx, environmentID).Execute()

#### Example

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllGateways(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllGateways``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllGateways`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllGateways`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllGateways(ctx, environmentID).ExecuteInitialPage()

#### Example

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
    resp, r, err := apiClient.GatewaysApi.ReadAllGateways(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.ReadAllGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGateways`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.ReadAllGateways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneGateway

> CreateGateway201Response ReadOneGateway(ctx, environmentID, gatewayID).Execute()

READ One Gateway

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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.ReadOneGateway(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.ReadOneGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.ReadOneGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGateway

> CreateGateway201Response UpdateGateway(ctx, environmentID, gatewayID).CreateGatewayRequest(createGatewayRequest).Execute()

UPDATE Gateway

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
    gatewayID := "gatewayID_example" // string | 
    createGatewayRequest := openapiclient.createGateway_request{Gateway: openapiclient.NewGateway("Name_example", openapiclient.EnumGatewayType("LDAP"), false)} // CreateGatewayRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewaysApi.UpdateGateway(context.Background(), environmentID, gatewayID).CreateGatewayRequest(createGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewaysApi.UpdateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGateway`: CreateGateway201Response
    fmt.Fprintf(os.Stdout, "Response from `GatewaysApi.UpdateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createGatewayRequest** | [**CreateGatewayRequest**](CreateGatewayRequest.md) |  | 

### Return type

[**CreateGateway201Response**](CreateGateway201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

