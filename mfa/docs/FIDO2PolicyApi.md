# \FIDO2PolicyApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFIDO2Policy**](FIDO2PolicyApi.md#CreateFIDO2Policy) | **Post** /environments/{environmentID}/fido2Policies | CREATE FIDO2 Policy
[**DeleteFIDO2Policy**](FIDO2PolicyApi.md#DeleteFIDO2Policy) | **Delete** /environments/{environmentID}/fido2Policies/{fido2PolicyID} | DELETE FIDO2 Policy
[**ReadFIDO2Policies**](FIDO2PolicyApi.md#ReadFIDO2Policies) | **Get** /environments/{environmentID}/fido2Policies | READ FIDO2 Policies
[**ReadOneFIDO2Policy**](FIDO2PolicyApi.md#ReadOneFIDO2Policy) | **Get** /environments/{environmentID}/fido2Policies/{fido2PolicyID} | READ One FIDO2 Policy
[**UpdateFIDO2Policy**](FIDO2PolicyApi.md#UpdateFIDO2Policy) | **Put** /environments/{environmentID}/fido2Policies/{fido2PolicyID} | UPDATE FIDO2 Policy



## CreateFIDO2Policy

> FIDO2Policy CreateFIDO2Policy(ctx, environmentID).FIDO2Policy(fIDO2Policy).Execute()

CREATE FIDO2 Policy

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
	fIDO2Policy := *openapiclient.NewFIDO2Policy(openapiclient.EnumFIDO2PolicyAttestationRequirements("DIRECT"), openapiclient.EnumFIDO2PolicyAuthenticatorAttachment("PLATFORM"), *openapiclient.NewFIDO2PolicyBackupEligibility(false, false), "DeviceDisplayName_example", openapiclient.EnumFIDO2PolicyDiscoverableCredentials("DISCOURAGED"), *openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirements(false, openapiclient.EnumFIDO2PolicyMDSAuthenticatorOption("NONE")), "Name_example", "RelyingPartyId_example", *openapiclient.NewFIDO2PolicyUserDisplayNameAttributes([]openapiclient.FIDO2PolicyUserDisplayNameAttributesAttributesInner{*openapiclient.NewFIDO2PolicyUserDisplayNameAttributesAttributesInner("Name_example")}), *openapiclient.NewFIDO2PolicyUserVerification(false, openapiclient.EnumFIDO2PolicyUserVerificationOption("REQUIRED"))) // FIDO2Policy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDO2PolicyApi.CreateFIDO2Policy(context.Background(), environmentID).FIDO2Policy(fIDO2Policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.CreateFIDO2Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFIDO2Policy`: FIDO2Policy
	fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.CreateFIDO2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFIDO2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fIDO2Policy** | [**FIDO2Policy**](FIDO2Policy.md) |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFIDO2Policy

> DeleteFIDO2Policy(ctx, environmentID, fido2PolicyID).Execute()

DELETE FIDO2 Policy

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
	fido2PolicyID := "fido2PolicyID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FIDO2PolicyApi.DeleteFIDO2Policy(context.Background(), environmentID, fido2PolicyID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.DeleteFIDO2Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFIDO2PolicyRequest struct via the builder pattern


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


## ReadFIDO2Policies

READ FIDO2 Policies

### Paged Response (Recommended)

> PagedIterator[ReadFIDO2Policies200Response] ReadFIDO2Policies(ctx, environmentID).Execute()

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
	pagedIterator := apiClient.FIDO2PolicyApi.ReadFIDO2Policies(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.ReadFIDO2Policies``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadFIDO2Policies` page iteration: ReadFIDO2Policies200Response
		fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.ReadFIDO2Policies` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadFIDO2Policies200Response ReadFIDO2Policies(ctx, environmentID).ExecuteInitialPage()

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
	resp, r, err := apiClient.FIDO2PolicyApi.ReadFIDO2Policies(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.ReadFIDO2Policies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadFIDO2Policies`: ReadFIDO2Policies200Response
	fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.ReadFIDO2Policies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFIDO2PoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

Page Iterator: PagedIterator[[**ReadFIDO2Policies200Response**](ReadFIDO2Policies200Response.md)]

PagedIterator[ReadFIDO2Policies200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadFIDO2Policies200Response**](ReadFIDO2Policies200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadFIDO2Policies200Response**](ReadFIDO2Policies200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneFIDO2Policy

> FIDO2Policy ReadOneFIDO2Policy(ctx, environmentID, fido2PolicyID).Expand(expand).Execute()

READ One FIDO2 Policy

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
	fido2PolicyID := "fido2PolicyID_example" // string | 
	expand := "expand_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDO2PolicyApi.ReadOneFIDO2Policy(context.Background(), environmentID, fido2PolicyID).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.ReadOneFIDO2Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneFIDO2Policy`: FIDO2Policy
	fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.ReadOneFIDO2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFIDO2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFIDO2Policy

> FIDO2Policy UpdateFIDO2Policy(ctx, environmentID, fido2PolicyID).FIDO2Policy(fIDO2Policy).Execute()

UPDATE FIDO2 Policy

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
	fido2PolicyID := "fido2PolicyID_example" // string | 
	fIDO2Policy := *openapiclient.NewFIDO2Policy(openapiclient.EnumFIDO2PolicyAttestationRequirements("DIRECT"), openapiclient.EnumFIDO2PolicyAuthenticatorAttachment("PLATFORM"), *openapiclient.NewFIDO2PolicyBackupEligibility(false, false), "DeviceDisplayName_example", openapiclient.EnumFIDO2PolicyDiscoverableCredentials("DISCOURAGED"), *openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirements(false, openapiclient.EnumFIDO2PolicyMDSAuthenticatorOption("NONE")), "Name_example", "RelyingPartyId_example", *openapiclient.NewFIDO2PolicyUserDisplayNameAttributes([]openapiclient.FIDO2PolicyUserDisplayNameAttributesAttributesInner{*openapiclient.NewFIDO2PolicyUserDisplayNameAttributesAttributesInner("Name_example")}), *openapiclient.NewFIDO2PolicyUserVerification(false, openapiclient.EnumFIDO2PolicyUserVerificationOption("REQUIRED"))) // FIDO2Policy |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDO2PolicyApi.UpdateFIDO2Policy(context.Background(), environmentID, fido2PolicyID).FIDO2Policy(fIDO2Policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.UpdateFIDO2Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFIDO2Policy`: FIDO2Policy
	fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.UpdateFIDO2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFIDO2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fIDO2Policy** | [**FIDO2Policy**](FIDO2Policy.md) |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

