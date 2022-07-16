# \ActiveIdentityCountsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDActiveIdentityCountsGet**](ActiveIdentityCountsApi.md#V1EnvironmentsEnvironmentIDActiveIdentityCountsGet) | **Get** /v1/environments/{environmentID}/activeIdentityCounts | READ Active Identity Counts (Deprecated)
[**V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet**](ActiveIdentityCountsApi.md#V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet) | **Get** /v1/environments/{environmentID}/metrics/activeIdentityCounts | READ Active Identity Counts by Date Range
[**V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet**](ActiveIdentityCountsApi.md#V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet) | **Get** /v1/organizations/{organizationID}/licenses/{licenseID}/metrics/activeIdentityCounts | READ Active Identity Counts by License



## V1EnvironmentsEnvironmentIDActiveIdentityCountsGet

> V1EnvironmentsEnvironmentIDActiveIdentityCountsGet(ctx, environmentID).Filter(filter).Limit(limit).Order(order).Execute()

READ Active Identity Counts (Deprecated)

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
    filter := "startDate ge "2019-05-01T19:00:00Z" and samplingPeriod eq "10"" // string |  (optional)
    limit := int32(10) // int32 |  (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveIdentityCountsApi.V1EnvironmentsEnvironmentIDActiveIdentityCountsGet(context.Background(), environmentID).Filter(filter).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.V1EnvironmentsEnvironmentIDActiveIdentityCountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** |  | 
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


## V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet

> V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(ctx, environmentID).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()

READ Active Identity Counts by Date Range

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
    filter := "startDate ge "2020-05-01T19:00:00Z"" // string |  (optional)
    limit := int32(10) // int32 |  (optional)
    order := "-startDate" // string |  (optional)
    samplePeriod := "MONTH" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveIdentityCountsApi.V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(context.Background(), environmentID).Filter(filter).Limit(limit).Order(order).SamplePeriod(samplePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest struct via the builder pattern


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


## V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet

> V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet(ctx, organizationID, licenseID).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()

READ Active Identity Counts by License

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
    organizationID := "organizationID_example" // string | 
    licenseID := "licenseID_example" // string | 
    aggregatedBy := "calendarMonth" // string |  (optional)
    limit := int32(20) // int32 |  (optional)
    order := "-startDate" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveIdentityCountsApi.V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet(context.Background(), organizationID, licenseID).AggregatedBy(aggregatedBy).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveIdentityCountsApi.V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest struct via the builder pattern


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

