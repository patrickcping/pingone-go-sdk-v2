# \RiskEvaluationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskEvaluation**](RiskEvaluationsApi.md#CreateRiskEvaluation) | **Post** /v1/environments/{environmentID}/riskEvaluations | CREATE Risk Evaluation
[**ReadOneRiskEvaluation**](RiskEvaluationsApi.md#ReadOneRiskEvaluation) | **Get** /v1/environments/{environmentID}/riskEvaluations/{riskEvaluationID} | READ One Risk Evaluation
[**UpdateRiskEvaluation**](RiskEvaluationsApi.md#UpdateRiskEvaluation) | **Put** /v1/environments/{environmentID}/riskEvaluations/{riskEvaluationID}/event | UPDATE Risk Evaluation



## CreateRiskEvaluation

> RiskEvaluation CreateRiskEvaluation(ctx, environmentID).RiskEvaluation(riskEvaluation).Execute()

CREATE Risk Evaluation

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
    riskEvaluation := *openapiclient.NewRiskEvaluation(*openapiclient.NewRiskEvaluationEvent("Ip_example", *openapiclient.NewRiskEvaluationEventUser("Id_example", "Type_example"))) // RiskEvaluation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskEvaluationsApi.CreateRiskEvaluation(context.Background(), environmentID).RiskEvaluation(riskEvaluation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskEvaluationsApi.CreateRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `RiskEvaluationsApi.CreateRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskEvaluation** | [**RiskEvaluation**](RiskEvaluation.md) |  | 

### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneRiskEvaluation

> RiskEvaluation ReadOneRiskEvaluation(ctx, environmentID, riskEvaluationID).Execute()

READ One Risk Evaluation

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
    riskEvaluationID := "riskEvaluationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskEvaluationsApi.ReadOneRiskEvaluation(context.Background(), environmentID, riskEvaluationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskEvaluationsApi.ReadOneRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `RiskEvaluationsApi.ReadOneRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskEvaluationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskEvaluation

> RiskEvaluation UpdateRiskEvaluation(ctx, environmentID, riskEvaluationID).RiskEvaluationEvent(riskEvaluationEvent).Execute()

UPDATE Risk Evaluation

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
    riskEvaluationID := "riskEvaluationID_example" // string | 
    riskEvaluationEvent := *openapiclient.NewRiskEvaluationEvent("Ip_example", *openapiclient.NewRiskEvaluationEventUser("Id_example", "Type_example")) // RiskEvaluationEvent |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RiskEvaluationsApi.UpdateRiskEvaluation(context.Background(), environmentID, riskEvaluationID).RiskEvaluationEvent(riskEvaluationEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RiskEvaluationsApi.UpdateRiskEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRiskEvaluation`: RiskEvaluation
    fmt.Fprintf(os.Stdout, "Response from `RiskEvaluationsApi.UpdateRiskEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskEvaluationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskEvaluationEvent** | [**RiskEvaluationEvent**](RiskEvaluationEvent.md) |  | 

### Return type

[**RiskEvaluation**](RiskEvaluation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

