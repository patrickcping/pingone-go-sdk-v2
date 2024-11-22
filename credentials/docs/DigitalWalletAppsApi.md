# \DigitalWalletAppsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWalletApp**](DigitalWalletAppsApi.md#CreateDigitalWalletApp) | **Post** /environments/{environmentID}/digitalWalletApplications | Create Digital Wallet App
[**DeleteDigitalWalletApp**](DigitalWalletAppsApi.md#DeleteDigitalWalletApp) | **Delete** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Delete Digital Wallet App
[**ReadAllDigitalWalletApps**](DigitalWalletAppsApi.md#ReadAllDigitalWalletApps) | **Get** /environments/{environmentID}/digitalWalletApplications | Read All Digital Wallet Apps
[**ReadOneDigitalWalletApp**](DigitalWalletAppsApi.md#ReadOneDigitalWalletApp) | **Get** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Read One Digital Wallet App
[**UpdateDigitalWalletApp**](DigitalWalletAppsApi.md#UpdateDigitalWalletApp) | **Put** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Update Digital Wallet App



## CreateDigitalWalletApp

> DigitalWalletApplication CreateDigitalWalletApp(ctx, environmentID).DigitalWalletApplication(digitalWalletApplication).Execute()

Create Digital Wallet App

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
	digitalWalletApplication := *openapiclient.NewDigitalWalletApplication(*openapiclient.NewObjectApplication(), "AppOpenUrl_example", "Name_example") // DigitalWalletApplication |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletAppsApi.CreateDigitalWalletApp(context.Background(), environmentID).DigitalWalletApplication(digitalWalletApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.CreateDigitalWalletApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDigitalWalletApp`: DigitalWalletApplication
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletAppsApi.CreateDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digitalWalletApplication** | [**DigitalWalletApplication**](DigitalWalletApplication.md) |  | 

### Return type

[**DigitalWalletApplication**](DigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDigitalWalletApp

> DeleteDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).Execute()

Delete Digital Wallet App

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
	digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DigitalWalletAppsApi.DeleteDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.DeleteDigitalWalletApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletAppRequest struct via the builder pattern


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


## ReadAllDigitalWalletApps

Read All Digital Wallet Apps

### Paged Response (Recommended)

> PagedIterator[ReadAllDigitalWalletApps200Response] ReadAllDigitalWalletApps(ctx, environmentID).Execute()

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
	pagedIterator := apiClient.DigitalWalletAppsApi.ReadAllDigitalWalletApps(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.ReadAllDigitalWalletApps``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadAllDigitalWalletApps` page iteration: ReadAllDigitalWalletApps200Response
		fmt.Fprintf(os.Stdout, "Response from `DigitalWalletAppsApi.ReadAllDigitalWalletApps` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadAllDigitalWalletApps200Response ReadAllDigitalWalletApps(ctx, environmentID).ExecuteInitialPage()

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
	resp, r, err := apiClient.DigitalWalletAppsApi.ReadAllDigitalWalletApps(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.ReadAllDigitalWalletApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAllDigitalWalletApps`: ReadAllDigitalWalletApps200Response
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletAppsApi.ReadAllDigitalWalletApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDigitalWalletAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

Page Iterator: PagedIterator[[**ReadAllDigitalWalletApps200Response**](ReadAllDigitalWalletApps200Response.md)]

PagedIterator[ReadAllDigitalWalletApps200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadAllDigitalWalletApps200Response**](ReadAllDigitalWalletApps200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadAllDigitalWalletApps200Response**](ReadAllDigitalWalletApps200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneDigitalWalletApp

> DigitalWalletApplication ReadOneDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).Execute()

Read One Digital Wallet App

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
	digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletAppsApi.ReadOneDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.ReadOneDigitalWalletApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneDigitalWalletApp`: DigitalWalletApplication
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletAppsApi.ReadOneDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DigitalWalletApplication**](DigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDigitalWalletApp

> DigitalWalletApplication UpdateDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).DigitalWalletApplication(digitalWalletApplication).Execute()

Update Digital Wallet App

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
	digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 
	digitalWalletApplication := *openapiclient.NewDigitalWalletApplication(*openapiclient.NewObjectApplication(), "AppOpenUrl_example", "Name_example") // DigitalWalletApplication |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DigitalWalletAppsApi.UpdateDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).DigitalWalletApplication(digitalWalletApplication).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletAppsApi.UpdateDigitalWalletApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDigitalWalletApp`: DigitalWalletApplication
	fmt.Fprintf(os.Stdout, "Response from `DigitalWalletAppsApi.UpdateDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **digitalWalletApplication** | [**DigitalWalletApplication**](DigitalWalletApplication.md) |  | 

### Return type

[**DigitalWalletApplication**](DigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

