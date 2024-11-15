# \ApplicationsApplicationMFAPushCredentialsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#CreateMFAPushCredential) | **Post** /environments/{environmentID}/applications/{applicationID}/pushCredentials | CREATE MFA Push Credential
[**DeleteMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#DeleteMFAPushCredential) | **Delete** /environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | DELETE MFA Push Credential
[**ReadAllMFAPushCredentials**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadAllMFAPushCredentials) | **Get** /environments/{environmentID}/applications/{applicationID}/pushCredentials | READ All MFA Push Credentials
[**ReadOneMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#ReadOneMFAPushCredential) | **Get** /environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | READ One MFA Push Credential
[**UpdateMFAPushCredential**](ApplicationsApplicationMFAPushCredentialsApi.md#UpdateMFAPushCredential) | **Put** /environments/{environmentID}/applications/{applicationID}/pushCredentials/{pushCredentialID} | UPDATE MFA Push Credential



## CreateMFAPushCredential

> MFAPushCredentialResponse CreateMFAPushCredential(ctx, environmentID, applicationID).MFAPushCredentialRequest(mFAPushCredentialRequest).Execute()

CREATE MFA Push Credential

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
    applicationID := "applicationID_example" // string | 
    mFAPushCredentialRequest := openapiclient.MFAPushCredentialRequest{MFAPushCredentialAPNS: openapiclient.NewMFAPushCredentialAPNS(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example", "TeamId_example", "Token_example")} // MFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential(context.Background(), environmentID, applicationID).MFAPushCredentialRequest(mFAPushCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.CreateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMFAPushCredential`: MFAPushCredentialResponse
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


 **mFAPushCredentialRequest** | [**MFAPushCredentialRequest**](MFAPushCredentialRequest.md) |  | 

### Return type

[**MFAPushCredentialResponse**](MFAPushCredentialResponse.md)

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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 
    authorization := "Bearer {{accessToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.DeleteMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).Authorization(authorization).Execute()
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

READ All MFA Push Credentials

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllMFAPushCredentials(ctx, environmentID, applicationID).Execute()

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllMFAPushCredentials(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllMFAPushCredentials``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllMFAPushCredentials`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllMFAPushCredentials`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllMFAPushCredentials(ctx, environmentID, applicationID).ExecuteInitialPage()

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
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.ReadAllMFAPushCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllMFAPushCredentials`: EntityArrayPagedIterator
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

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneMFAPushCredential

> MFAPushCredentialResponse ReadOneMFAPushCredential(ctx, environmentID, applicationID, pushCredentialID).Execute()

READ One MFA Push Credential

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
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.ReadOneMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneMFAPushCredential`: MFAPushCredentialResponse
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

[**MFAPushCredentialResponse**](MFAPushCredentialResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMFAPushCredential

> MFAPushCredentialResponse UpdateMFAPushCredential(ctx, environmentID, applicationID, pushCredentialID).MFAPushCredentialRequest(mFAPushCredentialRequest).Execute()

UPDATE MFA Push Credential

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
    applicationID := "applicationID_example" // string | 
    pushCredentialID := "pushCredentialID_example" // string | 
    mFAPushCredentialRequest := openapiclient.MFAPushCredentialRequest{MFAPushCredentialAPNS: openapiclient.NewMFAPushCredentialAPNS(openapiclient.EnumMFAPushCredentialAttrType("APNS"), "Key_example", "TeamId_example", "Token_example")} // MFAPushCredentialRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential(context.Background(), environmentID, applicationID, pushCredentialID).MFAPushCredentialRequest(mFAPushCredentialRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApplicationMFAPushCredentialsApi.UpdateMFAPushCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMFAPushCredential`: MFAPushCredentialResponse
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



 **mFAPushCredentialRequest** | [**MFAPushCredentialRequest**](MFAPushCredentialRequest.md) |  | 

### Return type

[**MFAPushCredentialResponse**](MFAPushCredentialResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

