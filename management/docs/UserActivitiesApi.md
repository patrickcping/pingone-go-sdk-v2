# \UserActivitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDUserActivitiesGet**](UserActivitiesApi.md#V1EnvironmentsEnvironmentIDUserActivitiesGet) | **Get** /v1/environments/{environmentID}/userActivities | READ User Activities



## V1EnvironmentsEnvironmentIDUserActivitiesGet

> V1EnvironmentsEnvironmentIDUserActivitiesGet(ctx, environmentID).Filter(filter).Execute()

READ User Activities

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
    filter := "startDate eq "2018-02-17T09:10:12-04:00" and endDate eq "2018-02-23T09:10:12-04:00"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserActivitiesApi.V1EnvironmentsEnvironmentIDUserActivitiesGet(context.Background(), environmentID).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserActivitiesApi.V1EnvironmentsEnvironmentIDUserActivitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDUserActivitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

