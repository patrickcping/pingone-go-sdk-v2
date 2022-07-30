# \SignOnPoliciesSignOnPolicyActionsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSignOnPolicyAction**](SignOnPoliciesSignOnPolicyActionsApi.md#CreateSignOnPolicyAction) | **Post** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions | CREATE Sign-On Policy Action
[**DeleteSignOnPolicyAction**](SignOnPoliciesSignOnPolicyActionsApi.md#DeleteSignOnPolicyAction) | **Delete** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | DELETE Sign-On Policy Action
[**ReadAllSignOnPolicyActions**](SignOnPoliciesSignOnPolicyActionsApi.md#ReadAllSignOnPolicyActions) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions | READ All Sign-On Policy Actions
[**ReadOneSignOnPolicyAction**](SignOnPoliciesSignOnPolicyActionsApi.md#ReadOneSignOnPolicyAction) | **Get** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | READ One Sign-On Policy Action
[**UpdateSignOnPolicyAction**](SignOnPoliciesSignOnPolicyActionsApi.md#UpdateSignOnPolicyAction) | **Put** /v1/environments/{environmentID}/signOnPolicies/{policyID}/actions/{actionID} | UPDATE Sign-On Policy Action



## CreateSignOnPolicyAction

> SignOnPolicyAction CreateSignOnPolicyAction(ctx, environmentID, policyID).SignOnPolicyAction(signOnPolicyAction).Execute()

CREATE Sign-On Policy Action

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
    policyID := "policyID_example" // string | 
    signOnPolicyAction := openapiclient.SignOnPolicyAction{SignOnPolicyActionAgreement: openapiclient.NewSignOnPolicyActionAgreement(int32(123), openapiclient.EnumSignOnPolicyType("LOGIN"), *openapiclient.NewSignOnPolicyActionAgreementAllOfAgreement("Id_example"))} // SignOnPolicyAction |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.CreateSignOnPolicyAction(context.Background(), environmentID, policyID).SignOnPolicyAction(signOnPolicyAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.CreateSignOnPolicyAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSignOnPolicyAction`: SignOnPolicyAction
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesSignOnPolicyActionsApi.CreateSignOnPolicyAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignOnPolicyActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signOnPolicyAction** | [**SignOnPolicyAction**](SignOnPolicyAction.md) |  | 

### Return type

[**SignOnPolicyAction**](SignOnPolicyAction.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSignOnPolicyAction

> DeleteSignOnPolicyAction(ctx, environmentID, policyID, actionID).Execute()

DELETE Sign-On Policy Action

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
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.DeleteSignOnPolicyAction(context.Background(), environmentID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.DeleteSignOnPolicyAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSignOnPolicyActionRequest struct via the builder pattern


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


## ReadAllSignOnPolicyActions

> EntityArray ReadAllSignOnPolicyActions(ctx, environmentID, policyID).Execute()

READ All Sign-On Policy Actions

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
    policyID := "policyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.ReadAllSignOnPolicyActions(context.Background(), environmentID, policyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.ReadAllSignOnPolicyActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllSignOnPolicyActions`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesSignOnPolicyActionsApi.ReadAllSignOnPolicyActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllSignOnPolicyActionsRequest struct via the builder pattern


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


## ReadOneSignOnPolicyAction

> SignOnPolicyAction ReadOneSignOnPolicyAction(ctx, environmentID, policyID, actionID).Execute()

READ One Sign-On Policy Action

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
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.ReadOneSignOnPolicyAction(context.Background(), environmentID, policyID, actionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.ReadOneSignOnPolicyAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneSignOnPolicyAction`: SignOnPolicyAction
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesSignOnPolicyActionsApi.ReadOneSignOnPolicyAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneSignOnPolicyActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SignOnPolicyAction**](SignOnPolicyAction.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSignOnPolicyAction

> SignOnPolicyAction UpdateSignOnPolicyAction(ctx, environmentID, policyID, actionID).SignOnPolicyAction(signOnPolicyAction).Execute()

UPDATE Sign-On Policy Action

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
    policyID := "policyID_example" // string | 
    actionID := "actionID_example" // string | 
    signOnPolicyAction := openapiclient.SignOnPolicyAction{SignOnPolicyActionAgreement: openapiclient.NewSignOnPolicyActionAgreement(int32(123), openapiclient.EnumSignOnPolicyType("LOGIN"), *openapiclient.NewSignOnPolicyActionAgreementAllOfAgreement("Id_example"))} // SignOnPolicyAction |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignOnPoliciesSignOnPolicyActionsApi.UpdateSignOnPolicyAction(context.Background(), environmentID, policyID, actionID).SignOnPolicyAction(signOnPolicyAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignOnPoliciesSignOnPolicyActionsApi.UpdateSignOnPolicyAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSignOnPolicyAction`: SignOnPolicyAction
    fmt.Fprintf(os.Stdout, "Response from `SignOnPoliciesSignOnPolicyActionsApi.UpdateSignOnPolicyAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**policyID** | **string** |  | 
**actionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSignOnPolicyActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **signOnPolicyAction** | [**SignOnPolicyAction**](SignOnPolicyAction.md) |  | 

### Return type

[**SignOnPolicyAction**](SignOnPolicyAction.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

