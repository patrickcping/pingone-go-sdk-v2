# \GatewayCredentialsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGatewayCredential**](GatewayCredentialsApi.md#CreateGatewayCredential) | **Post** /environments/{environmentID}/gateways/{gatewayID}/credentials | CREATE Gateway Credentials
[**DeleteGatewayCredential**](GatewayCredentialsApi.md#DeleteGatewayCredential) | **Delete** /environments/{environmentID}/gateways/{gatewayID}/credentials/{credentialID} | DELETE Gateway Credentials
[**ReadAllGatewayCredentials**](GatewayCredentialsApi.md#ReadAllGatewayCredentials) | **Get** /environments/{environmentID}/gateways/{gatewayID}/credentials | READ All Gateway Credentials
[**ReadOneGatewayCredential**](GatewayCredentialsApi.md#ReadOneGatewayCredential) | **Get** /environments/{environmentID}/gateways/{gatewayID}/credentials/{credentialID} | READ One Gateway Credential



## CreateGatewayCredential

> GatewayCredential CreateGatewayCredential(ctx, environmentID, gatewayID).Execute()

CREATE Gateway Credentials

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
    resp, r, err := apiClient.GatewayCredentialsApi.CreateGatewayCredential(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayCredentialsApi.CreateGatewayCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGatewayCredential`: GatewayCredential
    fmt.Fprintf(os.Stdout, "Response from `GatewayCredentialsApi.CreateGatewayCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGatewayCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GatewayCredential**](GatewayCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGatewayCredential

> DeleteGatewayCredential(ctx, environmentID, gatewayID, credentialID).Execute()

DELETE Gateway Credentials

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
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GatewayCredentialsApi.DeleteGatewayCredential(context.Background(), environmentID, gatewayID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayCredentialsApi.DeleteGatewayCredential``: %v\n", err)
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
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGatewayCredentialRequest struct via the builder pattern


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


## ReadAllGatewayCredentials

READ All Gateway Credentials

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllGatewayCredentials(ctx, environmentID, gatewayID).Execute()

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
    pagedIterator := api.ReadAllGatewayCredentials(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllGatewayCredentials``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllGatewayCredentials`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllGatewayCredentials`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllGatewayCredentials(ctx, environmentID, gatewayID).ExecuteInitialPage()

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
    gatewayID := "gatewayID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayCredentialsApi.ReadAllGatewayCredentials(context.Background(), environmentID, gatewayID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayCredentialsApi.ReadAllGatewayCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGatewayCredentials`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `GatewayCredentialsApi.ReadAllGatewayCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGatewayCredentialsRequest struct via the builder pattern


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


## ReadOneGatewayCredential

> GatewayCredential ReadOneGatewayCredential(ctx, environmentID, gatewayID, credentialID).Execute()

READ One Gateway Credential

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
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GatewayCredentialsApi.ReadOneGatewayCredential(context.Background(), environmentID, gatewayID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GatewayCredentialsApi.ReadOneGatewayCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGatewayCredential`: GatewayCredential
    fmt.Fprintf(os.Stdout, "Response from `GatewayCredentialsApi.ReadOneGatewayCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**gatewayID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGatewayCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GatewayCredential**](GatewayCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

