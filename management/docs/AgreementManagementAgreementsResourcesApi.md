# \AgreementManagementAgreementsResourcesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete**](AgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete) | **Delete** /v1/environments/{environmentID}/agreements/{agreementID} | DELETE Agreement
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet**](AgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet) | **Get** /v1/environments/{environmentID}/agreements/{agreementID} | READ One Agreement
[**V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut**](AgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut) | **Put** /v1/environments/{environmentID}/agreements/{agreementID} | UPDATE Agreement
[**V1EnvironmentsEnvironmentIDAgreementsGet**](AgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsGet) | **Get** /v1/environments/{environmentID}/agreements | READ All Agreements
[**V1EnvironmentsEnvironmentIDAgreementsPost**](AgreementManagementAgreementsResourcesApi.md#V1EnvironmentsEnvironmentIDAgreementsPost) | **Post** /v1/environments/{environmentID}/agreements | CREATE Agreement



## V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete(ctx, environmentID, agreementID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet(ctx, environmentID, agreementID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet(context.Background(), environmentID, agreementID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut

> V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut(ctx, environmentID, agreementID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut(context.Background(), environmentID, agreementID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsAgreementIDPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsAgreementIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDAgreementsGet

> V1EnvironmentsEnvironmentIDAgreementsGet(ctx, environmentID).Execute()

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
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDAgreementsPost

> V1EnvironmentsEnvironmentIDAgreementsPost(ctx, environmentID).Body(body).Execute()

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AgreementManagementAgreementsResourcesApi.V1EnvironmentsEnvironmentIDAgreementsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDAgreementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

