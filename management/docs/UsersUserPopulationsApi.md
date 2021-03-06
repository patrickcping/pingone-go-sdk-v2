# \UsersUserPopulationsApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadUserPopulation**](UsersUserPopulationsApi.md#ReadUserPopulation) | **Get** /v1/environments/{environmentID}/users/{userID}/population | READ User Population
[**UpdateUserPopulation**](UsersUserPopulationsApi.md#UpdateUserPopulation) | **Put** /v1/environments/{environmentID}/users/{userID}/population | UPDATE User Population



## ReadUserPopulation

> UserPopulation ReadUserPopulation(ctx, environmentID, userID).Execute()

READ User Population

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserPopulationsApi.ReadUserPopulation(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserPopulationsApi.ReadUserPopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserPopulation`: UserPopulation
    fmt.Fprintf(os.Stdout, "Response from `UsersUserPopulationsApi.ReadUserPopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserPopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserPopulation**](UserPopulation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPopulation

> UserPopulation UpdateUserPopulation(ctx, environmentID, userID).UserPopulation(userPopulation).Execute()

UPDATE User Population

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
    userID := "userID_example" // string | 
    userPopulation := *openapiclient.NewUserPopulation("Id_example") // UserPopulation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersUserPopulationsApi.UpdateUserPopulation(context.Background(), environmentID, userID).UserPopulation(userPopulation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersUserPopulationsApi.UpdateUserPopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPopulation`: UserPopulation
    fmt.Fprintf(os.Stdout, "Response from `UsersUserPopulationsApi.UpdateUserPopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userPopulation** | [**UserPopulation**](UserPopulation.md) |  | 

### Return type

[**UserPopulation**](UserPopulation.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

