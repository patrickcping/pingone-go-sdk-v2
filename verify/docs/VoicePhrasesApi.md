# \VoicePhrasesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVoicePhrase**](VoicePhrasesApi.md#CreateVoicePhrase) | **Post** /environments/{environmentID}/voicePhrases | CREATE Voice Phrase
[**DeleteVoicePhrase**](VoicePhrasesApi.md#DeleteVoicePhrase) | **Delete** /environments/{environmentID}/voicePhrases/{voicePhraseID} | Delete Voice Phrase
[**ReadAllVoicePhraseContents**](VoicePhrasesApi.md#ReadAllVoicePhraseContents) | **Get** /environments/{environmentID}/voicePhrases/{voicePhraseID}/contents | READ All Voice Phrase Contents
[**ReadAllVoicePhrases**](VoicePhrasesApi.md#ReadAllVoicePhrases) | **Get** /environments/{environmentID}/voicePhrases | READ All Voice Phrases
[**ReadOneVoicePhrase**](VoicePhrasesApi.md#ReadOneVoicePhrase) | **Get** /environments/{environmentID}/voicePhrases/{voicePhraseID} | READ One Voice Phrase
[**UpdateVoicePhrase**](VoicePhrasesApi.md#UpdateVoicePhrase) | **Put** /environments/{environmentID}/voicePhrases/{voicePhraseID} | UPDATE Voice Phrase



## CreateVoicePhrase

> VoicePhrase CreateVoicePhrase(ctx, environmentID).VoicePhrase(voicePhrase).Execute()

CREATE Voice Phrase

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
    voicePhrase := *openapiclient.NewVoicePhrase("DisplayName_example") // VoicePhrase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicePhrasesApi.CreateVoicePhrase(context.Background(), environmentID).VoicePhrase(voicePhrase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.CreateVoicePhrase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVoicePhrase`: VoicePhrase
    fmt.Fprintf(os.Stdout, "Response from `VoicePhrasesApi.CreateVoicePhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVoicePhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **voicePhrase** | [**VoicePhrase**](VoicePhrase.md) |  | 

### Return type

[**VoicePhrase**](VoicePhrase.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVoicePhrase

> DeleteVoicePhrase(ctx, environmentID, voicePhraseID).Execute()

Delete Voice Phrase

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VoicePhrasesApi.DeleteVoicePhrase(context.Background(), environmentID, voicePhraseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.DeleteVoicePhrase``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVoicePhraseRequest struct via the builder pattern


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

> EntityArray ReadAllVoicePhraseContents(ctx, environmentID, voicePhraseID).Execute()

READ All Voice Phrase Contents

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicePhrasesApi.ReadAllVoicePhraseContents(context.Background(), environmentID, voicePhraseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.ReadAllVoicePhraseContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllVoicePhraseContents`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `VoicePhrasesApi.ReadAllVoicePhraseContents`: %v\n", resp)
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

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllVoicePhrases

> EntityArray ReadAllVoicePhrases(ctx, environmentID).Execute()

READ All Voice Phrases

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicePhrasesApi.ReadAllVoicePhrases(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.ReadAllVoicePhrases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllVoicePhrases`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `VoicePhrasesApi.ReadAllVoicePhrases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllVoicePhrasesRequest struct via the builder pattern


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


## ReadOneVoicePhrase

> VoicePhrase ReadOneVoicePhrase(ctx, environmentID, voicePhraseID).Execute()

READ One Voice Phrase

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicePhrasesApi.ReadOneVoicePhrase(context.Background(), environmentID, voicePhraseID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.ReadOneVoicePhrase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneVoicePhrase`: VoicePhrase
    fmt.Fprintf(os.Stdout, "Response from `VoicePhrasesApi.ReadOneVoicePhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneVoicePhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VoicePhrase**](VoicePhrase.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVoicePhrase

> VoicePhrase UpdateVoicePhrase(ctx, environmentID, voicePhraseID).VoicePhrase(voicePhrase).Execute()

UPDATE Voice Phrase

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
    voicePhrase := *openapiclient.NewVoicePhrase("DisplayName_example") // VoicePhrase |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VoicePhrasesApi.UpdateVoicePhrase(context.Background(), environmentID, voicePhraseID).VoicePhrase(voicePhrase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VoicePhrasesApi.UpdateVoicePhrase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVoicePhrase`: VoicePhrase
    fmt.Fprintf(os.Stdout, "Response from `VoicePhrasesApi.UpdateVoicePhrase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**voicePhraseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVoicePhraseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **voicePhrase** | [**VoicePhrase**](VoicePhrase.md) |  | 

### Return type

[**VoicePhrase**](VoicePhrase.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

