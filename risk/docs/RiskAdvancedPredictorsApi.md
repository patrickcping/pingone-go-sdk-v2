# \RiskAdvancedPredictorsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskPredictor**](RiskAdvancedPredictorsApi.md#CreateRiskPredictor) | **Post** /environments/{environmentID}/riskPredictors | CREATE Risk Predictor
[**DeleteRiskAdvancedPredictor**](RiskAdvancedPredictorsApi.md#DeleteRiskAdvancedPredictor) | **Delete** /environments/{environmentID}/riskPredictors/{riskPredictorID} | DELETE Risk Advanced Predictor
[**ReadAllRiskPredictors**](RiskAdvancedPredictorsApi.md#ReadAllRiskPredictors) | **Get** /environments/{environmentID}/riskPredictors | READ All Risk Predictors
[**ReadOneRiskPredictor**](RiskAdvancedPredictorsApi.md#ReadOneRiskPredictor) | **Get** /environments/{environmentID}/riskPredictors/{riskPredictorID} | READ One Risk Predictor
[**UpdateRiskPredictor**](RiskAdvancedPredictorsApi.md#UpdateRiskPredictor) | **Put** /environments/{environmentID}/riskPredictors/{riskPredictorID} | UPDATE Risk Predictor



## CreateRiskPredictor

> RiskPredictor CreateRiskPredictor(ctx, environmentID).RiskPredictor(riskPredictor).Execute()

CREATE Risk Predictor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 
    riskPredictor := openapiclient.RiskPredictor{RiskPredictorAdversaryInTheMiddle: openapiclient.NewRiskPredictorAdversaryInTheMiddle("Name_example", "CompactName_example", openapiclient.EnumPredictorType("ADVERSARY_IN_THE_MIDDLE"))} // RiskPredictor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskAdvancedPredictorsApi.CreateRiskPredictor(context.Background(), environmentID).RiskPredictor(riskPredictor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAdvancedPredictorsApi.CreateRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `RiskAdvancedPredictorsApi.CreateRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskPredictor** | [**RiskPredictor**](RiskPredictor.md) |  | 

### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskAdvancedPredictor

> DeleteRiskAdvancedPredictor(ctx, environmentID, riskPredictorID).Execute()

DELETE Risk Advanced Predictor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 
    riskPredictorID := "riskPredictorID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RiskAdvancedPredictorsApi.DeleteRiskAdvancedPredictor(context.Background(), environmentID, riskPredictorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAdvancedPredictorsApi.DeleteRiskAdvancedPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskAdvancedPredictorRequest struct via the builder pattern


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


## ReadAllRiskPredictors

READ All Risk Predictors

### Paged Response (Recommended)

> EntityArrayPagedIterator ReadAllRiskPredictors(ctx, environmentID).Execute()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 
	// ... other parameters

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
	api := apiClient. // .... API function
    pagedIterator := api.ReadAllRiskPredictors(context.Background(), environmentID, /* ... other parameters */).Execute()
	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `api.ReadAllRiskPredictors``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
			break
		}

		// response from `ReadAllRiskPredictors`: EntityArrayPagedIterator
		fmt.Fprintf(os.Stdout, "Response from `api.ReadAllRiskPredictors`: %v\n", pageCursor.EntityArray)
	}
}
```

### Initial Page Response

> EntityArray ReadAllRiskPredictors(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskAdvancedPredictorsApi.ReadAllRiskPredictors(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAdvancedPredictorsApi.ReadAllRiskPredictors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllRiskPredictors`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `RiskAdvancedPredictorsApi.ReadAllRiskPredictors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllRiskPredictorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneRiskPredictor

> RiskPredictor ReadOneRiskPredictor(ctx, environmentID, riskPredictorID).Execute()

READ One Risk Predictor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 
    riskPredictorID := "riskPredictorID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskAdvancedPredictorsApi.ReadOneRiskPredictor(context.Background(), environmentID, riskPredictorID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAdvancedPredictorsApi.ReadOneRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `RiskAdvancedPredictorsApi.ReadOneRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskPredictor

> RiskPredictor UpdateRiskPredictor(ctx, environmentID, riskPredictorID).RiskPredictor(riskPredictor).Execute()

UPDATE Risk Predictor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
    environmentID := "environmentID_example" // string | 
    riskPredictorID := "riskPredictorID_example" // string | 
    riskPredictor := openapiclient.RiskPredictor{RiskPredictorAdversaryInTheMiddle: openapiclient.NewRiskPredictorAdversaryInTheMiddle("Name_example", "CompactName_example", openapiclient.EnumPredictorType("ADVERSARY_IN_THE_MIDDLE"))} // RiskPredictor |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskAdvancedPredictorsApi.UpdateRiskPredictor(context.Background(), environmentID, riskPredictorID).RiskPredictor(riskPredictor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskAdvancedPredictorsApi.UpdateRiskPredictor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskPredictor`: RiskPredictor
    fmt.Fprintf(os.Stdout, "Response from `RiskAdvancedPredictorsApi.UpdateRiskPredictor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPredictorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskPredictorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskPredictor** | [**RiskPredictor**](RiskPredictor.md) |  | 

### Return type

[**RiskPredictor**](RiskPredictor.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

