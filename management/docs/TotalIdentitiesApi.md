# \TotalIdentitiesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDTotalIdentitiesGet**](TotalIdentitiesApi.md#V1EnvironmentsEnvironmentIDTotalIdentitiesGet) | **Get** /v1/environments/{environmentID}/totalIdentities | READ Total Identity Counts



## V1EnvironmentsEnvironmentIDTotalIdentitiesGet

> V1EnvironmentsEnvironmentIDTotalIdentitiesGet(ctx, environmentID).Filter(filter).Execute()

READ Total Identity Counts

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
    filter := "startDate eq "2019-05-01T19:00:00Z" and endDate eq "2019-05-31T19:00:00Z"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TotalIdentitiesApi.V1EnvironmentsEnvironmentIDTotalIdentitiesGet(context.Background(), environmentID).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TotalIdentitiesApi.V1EnvironmentsEnvironmentIDTotalIdentitiesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDTotalIdentitiesGetRequest struct via the builder pattern


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

