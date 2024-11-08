# \ActiveIdentityCountsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIDActiveIdentityCountsGet**](ActiveIdentityCountsApi.md#EnvironmentsEnvironmentIDActiveIdentityCountsGet) | **Get** /environments/{environmentID}/activeIdentityCounts | READ Active Identity Counts (Deprecated)
[**EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet**](ActiveIdentityCountsApi.md#EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet) | **Get** /environments/{environmentID}/metrics/activeIdentityCounts | READ Active Identity Counts by Date Range
[**ReadActiveIdentityCount**](ActiveIdentityCountsApi.md#ReadActiveIdentityCount) | **Get** /organizations/{organizationID}/licenses/{licenseID}/metrics/activeIdentityCounts | READ Active Identity Counts by License



## EnvironmentsEnvironmentIDActiveIdentityCountsGet

> EnvironmentsEnvironmentIDActiveIdentityCountsGet(ctx, environmentID).Filter(filter).Limit(limit).Order(order).Execute()

READ Active Identity Counts (Deprecated)

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
    filter := "startDate ge "2019-05-01T19:00:00Z" and samplingPeriod eq "10"" // string |  (optional)
    limit := int32(56) // int32 | Adding a paging value to limit the number of resources displayed per page (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveIdentityCountsApi.EnvironmentsEnvironmentIDActiveIdentityCountsGet(context.Background(), environmentID).Filter(filter).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.EnvironmentsEnvironmentIDActiveIdentityCountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 
 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **order** | **string** |  | 

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


## EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet

> EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(ctx, environmentID).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()

READ Active Identity Counts by Date Range

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
    filter := "startDate ge "2020-05-01T19:00:00Z"" // string |  (optional)
    limit := int32(10) // int32 |  (optional)
    order := "-startDate" // string |  (optional)
    samplePeriod := "MONTH" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveIdentityCountsApi.EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(context.Background(), environmentID).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 
 **limit** | **int32** |  | 
 **order** | **string** |  | 
 **samplePeriod** | **string** |  | 

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


## ReadActiveIdentityCount

> ReadActiveIdentityCount(ctx, organizationID, licenseID).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()

READ Active Identity Counts by License

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
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 
    aggregatedBy := "calendarMonth" // string |  (optional)
    limit := int32(20) // int32 |  (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveIdentityCountsApi.ReadActiveIdentityCount(context.Background(), organizationID, licenseID).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.ReadActiveIdentityCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationID** | **string** |  | 
**licenseID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadActiveIdentityCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **aggregatedBy** | **string** |  | 
 **limit** | **int32** |  | 
 **order** | **string** |  | 

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

