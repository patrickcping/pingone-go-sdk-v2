# \ApplicationsApplicationMFAPushCredentialsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#CreateMFAPushCredential) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials | CREATE MFA Push Credential
[**DeleteMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#DeleteMFAPushCredential) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | DELETE MFA Push Credential
[**ReadAllMFAPushCredentials**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadAllMFAPushCredentials) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials | READ All MFA Push Credentials
[**ReadOneMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadOneMFAPushCredential) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | READ One MFA Push Credential
[**UpdateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#UpdateMFAPushCredential) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | UPDATE MFA Push Credential



## CreateMFAPushCredential

> CreateMFAPushCredential201Response CreateMFAPushCredential(ctx, environmentID, applicationID).CreateMFAPushCredentialRequest(createMFAPushCredentialRequest).Execute()

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    createMFAPushCredentialRequest := openapiclient.createMFAPushCredential_request{MFAPushCredential: openapiclient.NewMFAPushCredential(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example")} // CreateMFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential(context.Background(), environmentID, applicationID).CreateMFAPushCredentialRequest(createMFAPushCredentialRequest).Execute()
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
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

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

> DeleteMFAPushCredential(ctx, environmentID, applicationID, pushCredentialID).Authorization(authorization).Execute()

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 
    authorization := "Bearer {{accessToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).Authorization(authorization).Execute()
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
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**pushCredentialID** | **string** |  | 

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

> EntityArray ReadAllMFAPushCredentials(ctx, environmentID, applicationID).Execute()

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(context.Background(), environmentID, applicationID).Execute()
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
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

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

> CreateMFAPushCredential201Response ReadOneMFAPushCredential(ctx, environmentID, applicationID, pushCredentialID).Execute()

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).Execute()
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
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**pushCredentialID** | **string** |  | 

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

> CreateMFAPushCredential201Response UpdateMFAPushCredential(ctx, environmentID, applicationID, pushCredentialID).UpdateMFAPushCredentialRequest(updateMFAPushCredentialRequest).Execute()

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
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 
    updateMFAPushCredentialRequest := openapiclient.updateMFAPushCredential_request{MFAPushCredential: openapiclient.NewMFAPushCredential(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example")} // UpdateMFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).UpdateMFAPushCredentialRequest(updateMFAPushCredentialRequest).Execute()
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
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**pushCredentialID** | **string** |  | 

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

