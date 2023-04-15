# \DigitalWalletsApi

All URIs are relative to *https://api.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWallet**](DigitalWalletsApi.md#CreateDigitalWallet) | **Post** /environments/{environmentID}/users/{userID}/digitalWallets | Create Digital Wallet
[**DeleteDigitalWallet**](DigitalWalletsApi.md#DeleteDigitalWallet) | **Delete** /environments/{environmentID}/users/{userID}/digitalWallets/{digitalWalletID} | Delete Digital Wallet
[**ReadAllDigitalWallets**](DigitalWalletsApi.md#ReadAllDigitalWallets) | **Get** /environments/{environmentID}/users/{userID}/digitalWallets | Read All Digital Wallets
[**ReadOneDigitalWallet**](DigitalWalletsApi.md#ReadOneDigitalWallet) | **Get** /environments/{environmentID}/users/{userID}/digitalWallets/{digitalWalletID} | Read One Digital Wallet
[**ReadOneDigitalWalletCredential**](DigitalWalletsApi.md#ReadOneDigitalWalletCredential) | **Get** /environments/{environmentID}/users/{userID}/digitalWallets/{digitalWalletID}/provisionedCredentials | Read One Digital Wallet Credential
[**UpdateDigitalWallet**](DigitalWalletsApi.md#UpdateDigitalWallet) | **Put** /environments/{environmentID}/users/{userID}/digitalWallets/{digitalWalletID} | Update Digital Wallet



## CreateDigitalWallet

> CredentialDigitalWallet CreateDigitalWallet(ctx, environmentID, userID).CredentialDigitalWallet(credentialDigitalWallet).Execute()

Create Digital Wallet

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
    credentialDigitalWallet := *openapiclient.NewCredentialDigitalWallet() // CredentialDigitalWallet |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsApi.CreateDigitalWallet(context.Background(), environmentID, userID).CredentialDigitalWallet(credentialDigitalWallet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.CreateDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDigitalWallet`: CredentialDigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.CreateDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialDigitalWallet** | [**CredentialDigitalWallet**](CredentialDigitalWallet.md) |  | 

### Return type

[**CredentialDigitalWallet**](CredentialDigitalWallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDigitalWallet

> DeleteDigitalWallet(ctx, environmentID, userID, digitalWalletID).Execute()

Delete Digital Wallet

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
    digitalWalletID := "digitalWalletID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DigitalWalletsApi.DeleteDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.DeleteDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletRequest struct via the builder pattern


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


## ReadAllDigitalWallets

> EntityArray ReadAllDigitalWallets(ctx, environmentID, userID).Execute()

Read All Digital Wallets

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsApi.ReadAllDigitalWallets(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.ReadAllDigitalWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllDigitalWallets`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.ReadAllDigitalWallets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDigitalWalletsRequest struct via the builder pattern


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


## ReadOneDigitalWallet

> CredentialDigitalWallet ReadOneDigitalWallet(ctx, environmentID, userID, digitalWalletID).Execute()

Read One Digital Wallet

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
    digitalWalletID := "digitalWalletID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsApi.ReadOneDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.ReadOneDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDigitalWallet`: CredentialDigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.ReadOneDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**CredentialDigitalWallet**](CredentialDigitalWallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneDigitalWalletCredential

> EntityArray ReadOneDigitalWalletCredential(ctx, environmentID, userID, digitalWalletID).Execute()

Read One Digital Wallet Credential

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
    digitalWalletID := "digitalWalletID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsApi.ReadOneDigitalWalletCredential(context.Background(), environmentID, userID, digitalWalletID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.ReadOneDigitalWalletCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDigitalWalletCredential`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.ReadOneDigitalWalletCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletCredentialRequest struct via the builder pattern


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


## UpdateDigitalWallet

> CredentialDigitalWallet UpdateDigitalWallet(ctx, environmentID, userID, digitalWalletID).CredentialDigitalWallet(credentialDigitalWallet).Execute()

Update Digital Wallet

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
    digitalWalletID := "digitalWalletID_example" // string | 
    credentialDigitalWallet := *openapiclient.NewCredentialDigitalWallet() // CredentialDigitalWallet |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DigitalWalletsApi.UpdateDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).CredentialDigitalWallet(credentialDigitalWallet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DigitalWalletsApi.UpdateDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalWallet`: CredentialDigitalWallet
    fmt.Fprintf(os.Stdout, "Response from `DigitalWalletsApi.UpdateDigitalWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **credentialDigitalWallet** | [**CredentialDigitalWallet**](CredentialDigitalWallet.md) |  | 

### Return type

[**CredentialDigitalWallet**](CredentialDigitalWallet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

