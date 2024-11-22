# \FIDODeviceApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFidoDevice**](FIDODeviceApi.md#CreateFidoDevice) | **Post** /environments/{environmentID}/fidoDevicesMetadata | CREATE FIDO Device
[**DeleteFidoDevice**](FIDODeviceApi.md#DeleteFidoDevice) | **Delete** /environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | DELETE FIDO Device
[**ReadFidoDevices**](FIDODeviceApi.md#ReadFidoDevices) | **Get** /environments/{environmentID}/fidoDevicesMetadata | READ All FIDO Devices
[**ReadOneFidoDevice**](FIDODeviceApi.md#ReadOneFidoDevice) | **Get** /environments/{environmentID}/fidoDevicesMetadata/{fidoDeviceID} | READ One FIDO Device



## CreateFidoDevice

> map[string]interface{} CreateFidoDevice(ctx, environmentID).Body(body).Execute()

CREATE FIDO Device

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
	environmentID := "environmentID_example" // string | 
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDODeviceApi.CreateFidoDevice(context.Background(), environmentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDODeviceApi.CreateFidoDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFidoDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FIDODeviceApi.CreateFidoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFidoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFidoDevice

> DeleteFidoDevice(ctx, environmentID, fidoDeviceID).Execute()

DELETE FIDO Device

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
	environmentID := "environmentID_example" // string | 
	fidoDeviceID := "fidoDeviceID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FIDODeviceApi.DeleteFidoDevice(context.Background(), environmentID, fidoDeviceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDODeviceApi.DeleteFidoDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoDeviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFidoDeviceRequest struct via the builder pattern


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


## ReadFidoDevices

READ All FIDO Devices

### Paged Response (Recommended)

> PagedIterator[ReadFidoDevices200Response] ReadFidoDevices(ctx, environmentID).Execute()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.FIDODeviceApi.ReadFidoDevices(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FIDODeviceApi.ReadFidoDevices``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadFidoDevices` page iteration: ReadFidoDevices200Response
		fmt.Fprintf(os.Stdout, "Response from `FIDODeviceApi.ReadFidoDevices` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadFidoDevices200Response ReadFidoDevices(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDODeviceApi.ReadFidoDevices(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDODeviceApi.ReadFidoDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadFidoDevices`: ReadFidoDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `FIDODeviceApi.ReadFidoDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFidoDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

Page Iterator: PagedIterator[[**ReadFidoDevices200Response**](ReadFidoDevices200Response.md)]

PagedIterator[ReadFidoDevices200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadFidoDevices200Response**](ReadFidoDevices200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadFidoDevices200Response**](ReadFidoDevices200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneFidoDevice

> map[string]interface{} ReadOneFidoDevice(ctx, environmentID, fidoDeviceID).Execute()

READ One FIDO Device

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
	environmentID := "environmentID_example" // string | 
	fidoDeviceID := "fidoDeviceID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDODeviceApi.ReadOneFidoDevice(context.Background(), environmentID, fidoDeviceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDODeviceApi.ReadOneFidoDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneFidoDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FIDODeviceApi.ReadOneFidoDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fidoDeviceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFidoDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

