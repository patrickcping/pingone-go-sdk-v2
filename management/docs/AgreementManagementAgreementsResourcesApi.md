# \AgreementManagementAgreementsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAgreement**](AgreementManagementAgreementsResourcesApi.md#CreateAgreement) | **Post** /v1/environments/{environmentID}/agreements | CREATE Agreement
[**DeleteAgreement**](AgreementManagementAgreementsResourcesApi.md#DeleteAgreement) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID} | DELETE Agreement
[**ReadAllAgreements**](AgreementManagementAgreementsResourcesApi.md#ReadAllAgreements) | **Get** /v1/environments/{environmentID}/agreements | READ All Agreements
[**ReadOneAgreement**](AgreementManagementAgreementsResourcesApi.md#ReadOneAgreement) | **Get** /v1/environments/{environmentID}/agreements/{agreementID} | READ One Agreement
[**UpdateAgreement**](AgreementManagementAgreementsResourcesApi.md#UpdateAgreement) | **Put** /v1/environments/{environmentID}/agreements/{agreementID} | UPDATE Agreement



## CreateAgreement

> Agreement CreateAgreement(ctx, environmentID).Agreement(agreement).Execute()

CREATE Agreement

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
    agreement := *openapiclient.NewAgreement(false, "Name_example") // Agreement |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.CreateAgreement(context.Background(), environmentID).Agreement(agreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.CreateAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAgreement`: Agreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementsResourcesApi.CreateAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **agreement** | [**Agreement**](Agreement.md) |  | 

### Return type

[**Agreement**](Agreement.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAgreement

> DeleteAgreement(ctx, environmentID, agreementID).Execute()

DELETE Agreement

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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.DeleteAgreement(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.DeleteAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAgreementRequest struct via the builder pattern


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


## ReadAllAgreements

> EntityArray ReadAllAgreements(ctx, environmentID).Execute()

READ All Agreements

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
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.ReadAllAgreements(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.ReadAllAgreements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllAgreements`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementsResourcesApi.ReadAllAgreements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllAgreementsRequest struct via the builder pattern


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


## ReadOneAgreement

> Agreement ReadOneAgreement(ctx, environmentID, agreementID).Execute()

READ One Agreement

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
    agreementID := "agreementID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.ReadOneAgreement(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.ReadOneAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneAgreement`: Agreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementsResourcesApi.ReadOneAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Agreement**](Agreement.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAgreement

> Agreement UpdateAgreement(ctx, environmentID, agreementID).Agreement(agreement).Execute()

UPDATE Agreement

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
    agreementID := "agreementID_example" // string | 
    agreement := *openapiclient.NewAgreement(false, "Name_example") // Agreement |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.UpdateAgreement(context.Background(), environmentID, agreementID).Agreement(agreement).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.UpdateAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAgreement`: Agreement
    fmt.Fprintf(os.Stdout, "Response from `AgreementManagementAgreementsResourcesApi.UpdateAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**agreementID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **agreement** | [**Agreement**](Agreement.md) |  | 

### Return type

[**Agreement**](Agreement.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

