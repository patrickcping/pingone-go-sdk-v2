# \UserPopulationsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadUserPopulation**](UserPopulationsApi.md#ReadUserPopulation) | **Get** /environments/{environmentID}/users/{userID}/population | READ User Population
[**UpdateUserPopulation**](UserPopulationsApi.md#UpdateUserPopulation) | **Put** /environments/{environmentID}/users/{userID}/population | UPDATE User Population



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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserPopulationsApi.ReadUserPopulation(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserPopulationsApi.ReadUserPopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadUserPopulation`: UserPopulation
    fmt.Fprintf(os.Stdout, "Response from `UserPopulationsApi.ReadUserPopulation`: %v\n", resp)
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
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    userPopulation := *openapiclient.NewUserPopulation("Id_example") // UserPopulation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserPopulationsApi.UpdateUserPopulation(context.Background(), environmentID, userID).UserPopulation(userPopulation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserPopulationsApi.UpdateUserPopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserPopulation`: UserPopulation
    fmt.Fprintf(os.Stdout, "Response from `UserPopulationsApi.UpdateUserPopulation`: %v\n", resp)
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

