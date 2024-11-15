# \AuditActivitiesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDActivitiesActivityIDGet**](AuditActivitiesApi.md#EnvironmentsEnvironmentIDActivitiesActivityIDGet) | **Get** /environments/{environmentID}/activities/{activityID} | GET One User Activity
[**EnvironmentsEnvironmentIDActivitiesGet**](AuditActivitiesApi.md#EnvironmentsEnvironmentIDActivitiesGet) | **Get** /environments/{environmentID}/activities | GET User Activities
[**EnvironmentsEnvironmentIDActivitiesPost**](AuditActivitiesApi.md#EnvironmentsEnvironmentIDActivitiesPost) | **Post** /environments/{environmentID}/activities | GET User Activities



## EnvironmentsEnvironmentIDActivitiesActivityIDGet

> EnvironmentsEnvironmentIDActivitiesActivityIDGet(ctx, environmentID, activityID).Execute()

GET One User Activity

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
    activityID := "activityID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesActivityIDGet(context.Background(), environmentID, activityID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesActivityIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**activityID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDActivitiesActivityIDGetRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIDActivitiesGet

> EnvironmentsEnvironmentIDActivitiesGet(ctx, environmentID).Limit(limit).Filter(filter).Execute()

GET User Activities

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
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    filter := "recordedat gt "2018-08-20T00:00:00Z" AND recordedat lt "2018-08-22T23:59:00Z"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesGet(context.Background(), environmentID).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDActivitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** |  | 

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


## EnvironmentsEnvironmentIDActivitiesPost

> EnvironmentsEnvironmentIDActivitiesPost(ctx, environmentID).Body(body).Execute()

GET User Activities

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditActivitiesApi.EnvironmentsEnvironmentIDActivitiesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDActivitiesPostRequest struct via the builder pattern


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

