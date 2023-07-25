# \PasswordPoliciesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePasswordPolicy**](PasswordPoliciesApi.md#CreatePasswordPolicy) | **Post** /environments/{environmentID}/passwordPolicies | CREATE Password Policy
[**DeletePasswordPolicy**](PasswordPoliciesApi.md#DeletePasswordPolicy) | **Delete** /environments/{environmentID}/passwordPolicies/{passwordPolicyID} | DELETE Password Policy
[**ReadAllPasswordPolicies**](PasswordPoliciesApi.md#ReadAllPasswordPolicies) | **Get** /environments/{environmentID}/passwordPolicies | READ All Password Policies
[**ReadOnePasswordPolicy**](PasswordPoliciesApi.md#ReadOnePasswordPolicy) | **Get** /environments/{environmentID}/passwordPolicies/{passwordPolicyID} | READ One Password Policy
[**UpdatePasswordPolicy**](PasswordPoliciesApi.md#UpdatePasswordPolicy) | **Put** /environments/{environmentID}/passwordPolicies/{passwordPolicyID} | UPDATE Password Policy



## CreatePasswordPolicy

> PasswordPolicy CreatePasswordPolicy(ctx, environmentID).PasswordPolicy(passwordPolicy).Execute()

CREATE Password Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    passwordPolicy := *openapiclient.NewPasswordPolicy(false, false, "Name_example", false) // PasswordPolicy | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.CreatePasswordPolicy(context.Background(), environmentID).PasswordPolicy(passwordPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.CreatePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesApi.CreatePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePasswordPolicy

> DeletePasswordPolicy(ctx, environmentID, passwordPolicyID).Execute()

DELETE Password Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PasswordPoliciesApi.DeletePasswordPolicy(context.Background(), environmentID, passwordPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.DeletePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePasswordPolicyRequest struct via the builder pattern


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


## ReadAllPasswordPolicies

> EntityArray ReadAllPasswordPolicies(ctx, environmentID).Execute()

READ All Password Policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.ReadAllPasswordPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.ReadAllPasswordPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllPasswordPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesApi.ReadAllPasswordPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllPasswordPoliciesRequest struct via the builder pattern


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


## ReadOnePasswordPolicy

> PasswordPolicy ReadOnePasswordPolicy(ctx, environmentID, passwordPolicyID).Execute()

READ One Password Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.ReadOnePasswordPolicy(context.Background(), environmentID, passwordPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.ReadOnePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesApi.ReadOnePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePasswordPolicy

> PasswordPolicy UpdatePasswordPolicy(ctx, environmentID, passwordPolicyID).PasswordPolicy(passwordPolicy).Execute()

UPDATE Password Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    passwordPolicyID := "passwordPolicyID_example" // string | 
    passwordPolicy := *openapiclient.NewPasswordPolicy(false, false, "Name_example", false) // PasswordPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PasswordPoliciesApi.UpdatePasswordPolicy(context.Background(), environmentID, passwordPolicyID).PasswordPolicy(passwordPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPoliciesApi.UpdatePasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPoliciesApi.UpdatePasswordPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**passwordPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

