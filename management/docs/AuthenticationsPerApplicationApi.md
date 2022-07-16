# \AuthenticationsPerApplicationApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDApplicationSignonsGet**](AuthenticationsPerApplicationApi.md#V1EnvironmentsEnvironmentIDApplicationSignonsGet) | **Get** /v1/environments/{environmentID}/applicationSignons | READ Authentications Per Application (Partial)



## V1EnvironmentsEnvironmentIDApplicationSignonsGet

> V1EnvironmentsEnvironmentIDApplicationSignonsGet(ctx, environmentID).Limit(limit).SamplePeriod(samplePeriod).SamplePeriodCount(samplePeriodCount).Filter(filter).Execute()

READ Authentications Per Application (Partial)

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
    limit := int32(1) // int32 |  (optional)
    samplePeriod := int32(1) // int32 |  (optional)
    samplePeriodCount := int32(100) // int32 |  (optional)
    filter := "occurredAt ge "2019-10-03T00:00:00Z""" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationsPerApplicationApi.V1EnvironmentsEnvironmentIDApplicationSignonsGet(context.Background(), environmentID).Limit(limit).SamplePeriod(samplePeriod).SamplePeriodCount(samplePeriodCount).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationsPerApplicationApi.V1EnvironmentsEnvironmentIDApplicationSignonsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **samplePeriod** | **int32** |  | 
 **samplePeriodCount** | **int32** |  | 
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

