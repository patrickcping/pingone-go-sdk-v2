# \MFASettingsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadMFASettings**](MFASettingsApi.md#ReadMFASettings) | **Get** /v1/environments/{environmentID}/mfaSettings | READ MFA Settings
[**ResetMFASettings**](MFASettingsApi.md#ResetMFASettings) | **Delete** /v1/environments/{environmentID}/mfaSettings | RESET MFA Settings
[**UpdateMFASettings**](MFASettingsApi.md#UpdateMFASettings) | **Put** /v1/environments/{environmentID}/mfaSettings | UPDATE MFA Settings



## ReadMFASettings

> MFASettings ReadMFASettings(ctx, environmentID).Execute()

READ MFA Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFASettingsApi.ReadMFASettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFASettingsApi.ReadMFASettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadMFASettings`: MFASettings
    fmt.Fprintf(os.Stdout, "Response from `MFASettingsApi.ReadMFASettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadMFASettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MFASettings**](MFASettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetMFASettings

> MFASettings ResetMFASettings(ctx, environmentID).Execute()

RESET MFA Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFASettingsApi.ResetMFASettings(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFASettingsApi.ResetMFASettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetMFASettings`: MFASettings
    fmt.Fprintf(os.Stdout, "Response from `MFASettingsApi.ResetMFASettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetMFASettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MFASettings**](MFASettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMFASettings

> MFASettings UpdateMFASettings(ctx, environmentID).MFASettings(mFASettings).Execute()

UPDATE MFA Settings

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    mFASettings := *openapiclient.NewMFASettings(*openapiclient.NewMFASettingsPairing(int32(123), openapiclient.EnumMFASettingsPairingKeyFormat("NUMERIC"))) // MFASettings |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MFASettingsApi.UpdateMFASettings(context.Background(), environmentID).MFASettings(mFASettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MFASettingsApi.UpdateMFASettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFASettings`: MFASettings
    fmt.Fprintf(os.Stdout, "Response from `MFASettingsApi.UpdateMFASettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMFASettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFASettings** | [**MFASettings**](MFASettings.md) |  | 

### Return type

[**MFASettings**](MFASettings.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

