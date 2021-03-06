# \ApplicationsApplicationMFAPushCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#CreateMFAPushCredential) | **Post** /v1/environments/{envID}/applications/{appID}/pushCredentials | CREATE MFA Push Credential
[**DeleteMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#DeleteMFAPushCredential) | **Delete** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | DELETE MFA Push Credential
[**ReadAllMFAPushCredentials**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadAllMFAPushCredentials) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials | READ All MFA Push Credentials
[**ReadOneMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadOneMFAPushCredential) | **Get** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | READ One MFA Push Credential
[**UpdateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#UpdateMFAPushCredential) | **Put** /v1/environments/{envID}/applications/{appID}/pushCredentials/{pushCredID} | UPDATE MFA Push Credential



## CreateMFAPushCredential

> CreateMFAPushCredential201Response CreateMFAPushCredential(ctx, envID, appID).CreateMFAPushCredentialRequest(createMFAPushCredentialRequest).Execute()

CREATE MFA Push Credential



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
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    createMFAPushCredentialRequest := openapiclient.createMFAPushCredential_request{MFAPushCredential: openapiclient.NewMFAPushCredential(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example")} // CreateMFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential(context.Background(), envID, appID).CreateMFAPushCredentialRequest(createMFAPushCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMFAPushCredential`: CreateMFAPushCredential201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createMFAPushCredentialRequest** | [**CreateMFAPushCredentialRequest**](CreateMFAPushCredentialRequest.md) |  | 

### Return type

[**CreateMFAPushCredential201Response**](CreateMFAPushCredential201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMFAPushCredential

> DeleteMFAPushCredential(ctx, envID, appID, pushCredID).Authorization(authorization).Execute()

DELETE MFA Push Credential



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
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 
    authorization := "Bearer {{accessToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential(context.Background(), envID, appID, pushCredID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

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


## ReadAllMFAPushCredentials

> EntityArray ReadAllMFAPushCredentials(ctx, envID, appID).Execute()

READ All MFA Push Credentials



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
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(context.Background(), envID, appID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllMFAPushCredentials`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllMFAPushCredentialsRequest struct via the builder pattern


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


## ReadOneMFAPushCredential

> CreateMFAPushCredential201Response ReadOneMFAPushCredential(ctx, envID, appID, pushCredID).Execute()

READ One MFA Push Credential



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
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential(context.Background(), envID, appID, pushCredID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneMFAPushCredential`: CreateMFAPushCredential201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CreateMFAPushCredential201Response**](CreateMFAPushCredential201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMFAPushCredential

> CreateMFAPushCredential201Response UpdateMFAPushCredential(ctx, envID, appID, pushCredID).UpdateMFAPushCredentialRequest(updateMFAPushCredentialRequest).Execute()

UPDATE MFA Push Credential



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
    envID := "envID_example" // string | 
    appID := "appID_example" // string | 
    pushCredID := "pushCredID_example" // string | 
    updateMFAPushCredentialRequest := openapiclient.updateMFAPushCredential_request{MFAPushCredential: openapiclient.NewMFAPushCredential(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example")} // UpdateMFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential(context.Background(), envID, appID, pushCredID).UpdateMFAPushCredentialRequest(updateMFAPushCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFAPushCredential`: CreateMFAPushCredential201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**appID** | **string** |  | 
**pushCredID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMFAPushCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateMFAPushCredentialRequest** | [**UpdateMFAPushCredentialRequest**](UpdateMFAPushCredentialRequest.md) |  | 

### Return type

[**CreateMFAPushCredential201Response**](CreateMFAPushCredential201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

