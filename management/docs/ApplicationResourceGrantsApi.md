# \ApplicationResourceGrantsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationGrant**](ApplicationResourceGrantsApi.md#CreateApplicationGrant) | **Post** /v1/environments/{environmentID}/applications/{applicationID}/grants | CREATE Grant
[**DeleteApplicationGrant**](ApplicationResourceGrantsApi.md#DeleteApplicationGrant) | **Delete** /v1/environments/{environmentID}/applications/{applicationID}/grants/{grantID} | DELETE Grant
[**ReadAllApplicationGrants**](ApplicationResourceGrantsApi.md#ReadAllApplicationGrants) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/grants | READ All Grants for an Application
[**ReadOneApplicationGrant**](ApplicationResourceGrantsApi.md#ReadOneApplicationGrant) | **Get** /v1/environments/{environmentID}/applications/{applicationID}/grants/{grantID} | READ One Grant for an Application
[**UpdateApplicationGrant**](ApplicationResourceGrantsApi.md#UpdateApplicationGrant) | **Put** /v1/environments/{environmentID}/applications/{applicationID}/grants/{grantID} | UPDATE Grant



## CreateApplicationGrant

> ApplicationResourceGrant CreateApplicationGrant(ctx, environmentID, applicationID).ApplicationResourceGrant(applicationResourceGrant).Execute()

CREATE Grant

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
    applicationID := "applicationID_example" // string | 
    applicationResourceGrant := *openapiclient.NewApplicationResourceGrant(*openapiclient.NewApplicationResourceGrantResource("Id_example"), []openapiclient.ApplicationResourceGrantScopesInner{*openapiclient.NewApplicationResourceGrantScopesInner("Id_example")}) // ApplicationResourceGrant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourceGrantsApi.CreateApplicationGrant(context.Background(), environmentID, applicationID).ApplicationResourceGrant(applicationResourceGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourceGrantsApi.CreateApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourceGrantsApi.CreateApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **applicationResourceGrant** | [**ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | 

### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationGrant

> DeleteApplicationGrant(ctx, environmentID, applicationID, grantID).Execute()

DELETE Grant

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
    applicationID := "applicationID_example" // string | 
    grantID := "grantID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationResourceGrantsApi.DeleteApplicationGrant(context.Background(), environmentID, applicationID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourceGrantsApi.DeleteApplicationGrant``: %v\n", err)
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
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationGrantRequest struct via the builder pattern


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


## ReadAllApplicationGrants

> EntityArray ReadAllApplicationGrants(ctx, environmentID, applicationID).Execute()

READ All Grants for an Application

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
    applicationID := "applicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourceGrantsApi.ReadAllApplicationGrants(context.Background(), environmentID, applicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourceGrantsApi.ReadAllApplicationGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllApplicationGrants`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourceGrantsApi.ReadAllApplicationGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllApplicationGrantsRequest struct via the builder pattern


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


## ReadOneApplicationGrant

> ApplicationResourceGrant ReadOneApplicationGrant(ctx, environmentID, applicationID, grantID).Execute()

READ One Grant for an Application

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
    applicationID := "applicationID_example" // string | 
    grantID := "grantID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourceGrantsApi.ReadOneApplicationGrant(context.Background(), environmentID, applicationID, grantID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourceGrantsApi.ReadOneApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourceGrantsApi.ReadOneApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationGrant

> ApplicationResourceGrant UpdateApplicationGrant(ctx, environmentID, applicationID, grantID).ApplicationResourceGrant(applicationResourceGrant).Execute()

UPDATE Grant

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
    applicationID := "applicationID_example" // string | 
    grantID := "grantID_example" // string | 
    applicationResourceGrant := *openapiclient.NewApplicationResourceGrant(*openapiclient.NewApplicationResourceGrantResource("Id_example"), []openapiclient.ApplicationResourceGrantScopesInner{*openapiclient.NewApplicationResourceGrantScopesInner("Id_example")}) // ApplicationResourceGrant |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationResourceGrantsApi.UpdateApplicationGrant(context.Background(), environmentID, applicationID, grantID).ApplicationResourceGrant(applicationResourceGrant).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationResourceGrantsApi.UpdateApplicationGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationGrant`: ApplicationResourceGrant
    fmt.Fprintf(os.Stdout, "Response from `ApplicationResourceGrantsApi.UpdateApplicationGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**applicationID** | **string** |  | 
**grantID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **applicationResourceGrant** | [**ApplicationResourceGrant**](ApplicationResourceGrant.md) |  | 

### Return type

[**ApplicationResourceGrant**](ApplicationResourceGrant.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

