# \PopulationsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePopulation**](PopulationsApi.md#CreatePopulation) | **Post** /environments/{environmentID}/populations | CREATE Population
[**DeletePopulation**](PopulationsApi.md#DeletePopulation) | **Delete** /environments/{environmentID}/populations/{populationID} | DELETE Population
[**ReadAllPopulations**](PopulationsApi.md#ReadAllPopulations) | **Get** /environments/{environmentID}/populations | READ All Populations
[**ReadOnePopulation**](PopulationsApi.md#ReadOnePopulation) | **Get** /environments/{environmentID}/populations/{populationID} | READ One Population
[**UpdatePopulation**](PopulationsApi.md#UpdatePopulation) | **Put** /environments/{environmentID}/populations/{populationID} | UPDATE Population



## CreatePopulation

> Population CreatePopulation(ctx, environmentID).Population(population).Execute()

CREATE Population

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
    population := *openapiclient.NewPopulation("Name_example") // Population | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PopulationsApi.CreatePopulation(context.Background(), environmentID).Population(population).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PopulationsApi.CreatePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePopulation`: Population
    fmt.Fprintf(os.Stdout, "Response from `PopulationsApi.CreatePopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **population** | [**Population**](Population.md) |  | 

### Return type

[**Population**](Population.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePopulation

> DeletePopulation(ctx, environmentID, populationID).Execute()

DELETE Population

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
    populationID := "populationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PopulationsApi.DeletePopulation(context.Background(), environmentID, populationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PopulationsApi.DeletePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**populationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePopulationRequest struct via the builder pattern


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


## ReadAllPopulations

> EntityArray ReadAllPopulations(ctx, environmentID).Limit(limit).Filter(filter).Cursor(cursor).Execute()

READ All Populations

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
    filter := "id eq "60971d3b-cc5a-4601-9c44-2be541f91bf1"" // string | Adding a SCIM filter for a population ID or population name to display only those resources associated with the specified population. Only the id and name parameters are supported (optional)
    cursor := "cursor_example" // string | Adding a cursor value to retrieve the next page of results, used with the `limit` parameter. The cursor value is returned in the `_links.next.href` link in the response payload. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PopulationsApi.ReadAllPopulations(context.Background(), environmentID).Limit(limit).Filter(filter).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PopulationsApi.ReadAllPopulations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllPopulations`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PopulationsApi.ReadAllPopulations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllPopulationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Adding a paging value to limit the number of resources displayed per page | 
 **filter** | **string** | Adding a SCIM filter for a population ID or population name to display only those resources associated with the specified population. Only the id and name parameters are supported | 
 **cursor** | **string** | Adding a cursor value to retrieve the next page of results, used with the &#x60;limit&#x60; parameter. The cursor value is returned in the &#x60;_links.next.href&#x60; link in the response payload. | 

### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOnePopulation

> Population ReadOnePopulation(ctx, environmentID, populationID).Execute()

READ One Population

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
    populationID := "populationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PopulationsApi.ReadOnePopulation(context.Background(), environmentID, populationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PopulationsApi.ReadOnePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOnePopulation`: Population
    fmt.Fprintf(os.Stdout, "Response from `PopulationsApi.ReadOnePopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**populationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOnePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Population**](Population.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePopulation

> Population UpdatePopulation(ctx, environmentID, populationID).Population(population).Execute()

UPDATE Population

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
    populationID := "populationID_example" // string | 
    population := *openapiclient.NewPopulation("Name_example") // Population |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PopulationsApi.UpdatePopulation(context.Background(), environmentID, populationID).Population(population).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PopulationsApi.UpdatePopulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePopulation`: Population
    fmt.Fprintf(os.Stdout, "Response from `PopulationsApi.UpdatePopulation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**populationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePopulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **population** | [**Population**](Population.md) |  | 

### Return type

[**Population**](Population.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

