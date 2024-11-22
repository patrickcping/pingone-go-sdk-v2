# \VoicePhraseContentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoicePhraseContent**](VoicePhraseContentsApi.md#CreateVoicePhraseContent) | **Post** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents | CREATE Voice Phrase Content
[**DeleteVoicePhraseContent**](VoicePhraseContentsApi.md#DeleteVoicePhraseContent) | **Delete** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents/{voiceContentsID} | Delete Voice Phrase Content
[**ReadAllVoicePhraseContents**](VoicePhraseContentsApi.md#ReadAllVoicePhraseContents) | **Get** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents | READ All Voice Phrase Contents
[**ReadOneVoicePhraseContent**](VoicePhraseContentsApi.md#ReadOneVoicePhraseContent) | **Get** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents/{voiceContentsID} | READ One Voice Phrase Content
[**UpdateVoicePhraseContent**](VoicePhraseContentsApi.md#UpdateVoicePhraseContent) | **Put** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents/{voiceContentsID} | UPDATE Voice Phrase Content



## CreateVoicePhraseContent

> VoicePhraseContents CreateVoicePhraseContent(ctx, environmentID, voicePhraseID).VoicePhraseContents(voicePhraseContents).Execute()

CREATE Voice Phrase Content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 
	voicePhraseContents := *openapiclient.NewVoicePhraseContents(*openapiclient.NewVoicePhraseContentsVoicePhrase("Id_example"), "Locale_example", "Content_example") // VoicePhraseContents |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicePhraseContentsApi.CreateVoicePhraseContent(context.Background(), environmentID, voicePhraseID).VoicePhraseContents(voicePhraseContents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.CreateVoicePhraseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVoicePhraseContent`: VoicePhraseContents
	fmt.Fprintf(os.Stdout, "Response from `VoicePhraseContentsApi.CreateVoicePhraseContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoicePhraseContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voicePhraseContents** | [**VoicePhraseContents**](VoicePhraseContents.md) |  | 

### Return type

[**VoicePhraseContents**](VoicePhraseContents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVoicePhraseContent

> DeleteVoicePhraseContent(ctx, environmentID, voicePhraseID, voiceContentsID).Execute()

Delete Voice Phrase Content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 
	voiceContentsID := "voiceContentsID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VoicePhraseContentsApi.DeleteVoicePhraseContent(context.Background(), environmentID, voicePhraseID, voiceContentsID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.DeleteVoicePhraseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 
**voiceContentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoicePhraseContentRequest struct via the builder pattern


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


## ReadAllVoicePhraseContents

READ All Voice Phrase Contents

### Paged Response (Recommended)

> PagedIterator[ReadAllVoicePhraseContents200Response] ReadAllVoicePhraseContents(ctx, environmentID, voicePhraseID).Execute()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.VoicePhraseContentsApi.ReadAllVoicePhraseContents(context.Background(), environmentID, voicePhraseID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.ReadAllVoicePhraseContents``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadAllVoicePhraseContents` page iteration: ReadAllVoicePhraseContents200Response
		fmt.Fprintf(os.Stdout, "Response from `VoicePhraseContentsApi.ReadAllVoicePhraseContents` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadAllVoicePhraseContents200Response ReadAllVoicePhraseContents(ctx, environmentID, voicePhraseID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicePhraseContentsApi.ReadAllVoicePhraseContents(context.Background(), environmentID, voicePhraseID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.ReadAllVoicePhraseContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadAllVoicePhraseContents`: ReadAllVoicePhraseContents200Response
	fmt.Fprintf(os.Stdout, "Response from `VoicePhraseContentsApi.ReadAllVoicePhraseContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllVoicePhraseContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

Page Iterator: PagedIterator[[**ReadAllVoicePhraseContents200Response**](ReadAllVoicePhraseContents200Response.md)]

PagedIterator[ReadAllVoicePhraseContents200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadAllVoicePhraseContents200Response**](ReadAllVoicePhraseContents200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadAllVoicePhraseContents200Response**](ReadAllVoicePhraseContents200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneVoicePhraseContent

> VoicePhraseContents ReadOneVoicePhraseContent(ctx, environmentID, voicePhraseID, voiceContentsID).Execute()

READ One Voice Phrase Content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 
	voiceContentsID := "voiceContentsID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicePhraseContentsApi.ReadOneVoicePhraseContent(context.Background(), environmentID, voicePhraseID, voiceContentsID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.ReadOneVoicePhraseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneVoicePhraseContent`: VoicePhraseContents
	fmt.Fprintf(os.Stdout, "Response from `VoicePhraseContentsApi.ReadOneVoicePhraseContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 
**voiceContentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneVoicePhraseContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**VoicePhraseContents**](VoicePhraseContents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoicePhraseContent

> VoicePhraseContents UpdateVoicePhraseContent(ctx, environmentID, voicePhraseID, voiceContentsID).VoicePhraseContents(voicePhraseContents).Execute()

UPDATE Voice Phrase Content

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/verify"
)

func main() {
	environmentID := "environmentID_example" // string | 
	voicePhraseID := "voicePhraseID_example" // string | 
	voiceContentsID := "voiceContentsID_example" // string | 
	voicePhraseContents := *openapiclient.NewVoicePhraseContents(*openapiclient.NewVoicePhraseContentsVoicePhrase("Id_example"), "Locale_example", "Content_example") // VoicePhraseContents |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VoicePhraseContentsApi.UpdateVoicePhraseContent(context.Background(), environmentID, voicePhraseID, voiceContentsID).VoicePhraseContents(voicePhraseContents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VoicePhraseContentsApi.UpdateVoicePhraseContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVoicePhraseContent`: VoicePhraseContents
	fmt.Fprintf(os.Stdout, "Response from `VoicePhraseContentsApi.UpdateVoicePhraseContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 
**voiceContentsID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicePhraseContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **voicePhraseContents** | [**VoicePhraseContents**](VoicePhraseContents.md) |  | 

### Return type

[**VoicePhraseContents**](VoicePhraseContents.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

