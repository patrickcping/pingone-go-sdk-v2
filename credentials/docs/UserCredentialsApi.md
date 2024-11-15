# \UserCredentialsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAUserCredential**](UserCredentialsApi.md#CreateAUserCredential) | **Post** /environments/{environmentID}/users/{userID}/credentials | Create a User Credential
[**ReadAllUserCredentials**](UserCredentialsApi.md#ReadAllUserCredentials) | **Get** /environments/{environmentID}/users/{userID}/credentials | Read All User Credentials
[**ReadOneUserCredential**](UserCredentialsApi.md#ReadOneUserCredential) | **Get** /environments/{environmentID}/users/{userID}/credentials/{credentialID} | Read One User Credential
[**ReadOneUserCredentialWallet**](UserCredentialsApi.md#ReadOneUserCredentialWallet) | **Get** /environments/{environmentID}/users/{userID}/credentials/{credentialID}/provisionedCredentials | Read One User Credential Wallet
[**UpdateUserCredential**](UserCredentialsApi.md#UpdateUserCredential) | **Put** /environments/{environmentID}/users/{userID}/credentials/{credentialID} | Update a User Credential



## CreateAUserCredential

> UserCredential CreateAUserCredential(ctx, environmentID, userID).UserCredential(userCredential).Execute()

Create a User Credential

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
    userID := "userID_example" // string | 
    userCredential := *openapiclient.NewUserCredential() // UserCredential |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCredentialsApi.CreateAUserCredential(context.Background(), environmentID, userID).UserCredential(userCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.CreateAUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.CreateAUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userCredential** | [**UserCredential**](UserCredential.md) |  | 

### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllUserCredentials

Read All User Credentials

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllUserCredentials(ctx, environmentID, userID).Execute()

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllUserCredentials(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllUserCredentials``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllUserCredentials`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllUserCredentials`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllUserCredentials(ctx, environmentID, userID).ExecuteInitialPage()

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCredentialsApi.ReadAllUserCredentials(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.ReadAllUserCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllUserCredentials`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.ReadAllUserCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllUserCredentialsRequest struct via the builder pattern


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


## ReadOneUserCredential

> UserCredential ReadOneUserCredential(ctx, environmentID, userID, credentialID).Execute()

Read One User Credential

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCredentialsApi.ReadOneUserCredential(context.Background(), environmentID, userID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.ReadOneUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.ReadOneUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneUserCredentialWallet

Read One User Credential Wallet

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadOneUserCredentialWallet(ctx, environmentID, userID, credentialID).Execute()

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
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadOneUserCredentialWallet(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadOneUserCredentialWallet``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadOneUserCredentialWallet`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadOneUserCredentialWallet`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadOneUserCredentialWallet(ctx, environmentID, userID, credentialID).ExecuteInitialPage()

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCredentialsApi.ReadOneUserCredentialWallet(context.Background(), environmentID, userID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.ReadOneUserCredentialWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneUserCredentialWallet`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.ReadOneUserCredentialWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneUserCredentialWalletRequest struct via the builder pattern


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


## UpdateUserCredential

> UserCredential UpdateUserCredential(ctx, environmentID, userID, credentialID).UserCredential(userCredential).Execute()

Update a User Credential

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 
    userCredential := *openapiclient.NewUserCredential() // UserCredential |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserCredentialsApi.UpdateUserCredential(context.Background(), environmentID, userID, credentialID).UserCredential(userCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserCredentialsApi.UpdateUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `UserCredentialsApi.UpdateUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userCredential** | [**UserCredential**](UserCredential.md) |  | 

### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

