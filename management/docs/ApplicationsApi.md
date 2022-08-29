# \ApplicationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationsApi.md#CreateApplication) | **Post** /v1/environments/{environmentID}/applications | CREATE Application
[**DeleteApplication**](ApplicationsApi.md#DeleteApplication) | **Delete** /v1/environments/{environmentID}/applications/{applicationID} | DELETE Application
[**ReadAllApplications**](ApplicationsApi.md#ReadAllApplications) | **Get** /v1/environments/{environmentID}/applications | READ All Applications
[**ReadOneApplication**](ApplicationsApi.md#ReadOneApplication) | **Get** /v1/environments/{environmentID}/applications/{applicationID} | READ One Application
[**UpdateApplication**](ApplicationsApi.md#UpdateApplication) | **Put** /v1/environments/{environmentID}/applications/{applicationID} | UPDATE Application



## CreateApplication

> CreateApplication201Response CreateApplication(ctx, environmentID).CreateApplicationRequest(createApplicationRequest).Execute()

CREATE Application

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
    createApplicationRequest := openapiclient.createApplication_request{ApplicationOIDC: openapiclient.NewApplicationOIDC(false, "Name_example", openapiclient.EnumApplicationProtocol("OPENID_CONNECT"), openapiclient.EnumApplicationType("WEB_APP"), []openapiclient.EnumApplicationOIDCGrantType{openapiclient.EnumApplicationOIDCGrantType("AUTHORIZATION_CODE")}, openapiclient.EnumApplicationOIDCTokenAuthMethod("NONE"))} // CreateApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.CreateApplication(context.Background(), environmentID).CreateApplicationRequest(createApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.CreateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.CreateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md) |  | 

### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, environmentID, applicationID).Execute()

DELETE Application

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
    resp, r, err := apiClient.ApplicationsApi.DeleteApplication(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.DeleteApplication``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## ReadAllApplications

> EntityArray ReadAllApplications(ctx, environmentID).Execute()

READ All Applications

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
    resp, r, err := apiClient.ApplicationsApi.ReadAllApplications(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.ReadAllApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplications`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.ReadAllApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationsRequest struct via the builder pattern


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


## ReadOneApplication

> CreateApplication201Response ReadOneApplication(ctx, environmentID, applicationID).Execute()

READ One Application

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
    resp, r, err := apiClient.ApplicationsApi.ReadOneApplication(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.ReadOneApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.ReadOneApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> CreateApplication201Response UpdateApplication(ctx, environmentID, applicationID).UpdateApplicationRequest(updateApplicationRequest).Execute()

UPDATE Application

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
    updateApplicationRequest := openapiclient.updateApplication_request{ApplicationOIDC: openapiclient.NewApplicationOIDC(false, "Name_example", openapiclient.EnumApplicationProtocol("OPENID_CONNECT"), openapiclient.EnumApplicationType("WEB_APP"), []openapiclient.EnumApplicationOIDCGrantType{openapiclient.EnumApplicationOIDCGrantType("AUTHORIZATION_CODE")}, openapiclient.EnumApplicationOIDCTokenAuthMethod("NONE"))} // UpdateApplicationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsApi.UpdateApplication(context.Background(), environmentID, applicationID).UpdateApplicationRequest(updateApplicationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsApi.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: CreateApplication201Response
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsApi.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateApplicationRequest** | [**UpdateApplicationRequest**](UpdateApplicationRequest.md) |  | 

### Return type

[**CreateApplication201Response**](CreateApplication201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

