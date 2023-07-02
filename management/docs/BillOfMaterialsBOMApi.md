# \BillOfMaterialsBOMApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadOneBillOfMaterials**](BillOfMaterialsBOMApi.md#ReadOneBillOfMaterials) | **Get** /environments/{environmentID}/billOfMaterials | READ One Bill of Materials
[**UpdateBillOfMaterials**](BillOfMaterialsBOMApi.md#UpdateBillOfMaterials) | **Put** /environments/{environmentID}/billOfMaterials | UPDATE Bill of Materials



## ReadOneBillOfMaterials

> BillOfMaterials ReadOneBillOfMaterials(ctx, environmentID).Execute()

READ One Bill of Materials

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillOfMaterialsBOMApi.ReadOneBillOfMaterials(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfMaterialsBOMApi.ReadOneBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneBillOfMaterials`: BillOfMaterials
    fmt.Fprintf(os.Stdout, "Response from `BillOfMaterialsBOMApi.ReadOneBillOfMaterials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillOfMaterials**](BillOfMaterials.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBillOfMaterials

> BillOfMaterials UpdateBillOfMaterials(ctx, environmentID).BillOfMaterials(billOfMaterials).Execute()

UPDATE Bill of Materials

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
    billOfMaterials := *openapiclient.NewBillOfMaterials([]openapiclient.BillOfMaterialsProductsInner{*openapiclient.NewBillOfMaterialsProductsInner(openapiclient.EnumProductType("PING_ONE_MFA"))}) // BillOfMaterials |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillOfMaterialsBOMApi.UpdateBillOfMaterials(context.Background(), environmentID).BillOfMaterials(billOfMaterials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillOfMaterialsBOMApi.UpdateBillOfMaterials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBillOfMaterials`: BillOfMaterials
    fmt.Fprintf(os.Stdout, "Response from `BillOfMaterialsBOMApi.UpdateBillOfMaterials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBillOfMaterialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billOfMaterials** | [**BillOfMaterials**](BillOfMaterials.md) |  | 

### Return type

[**BillOfMaterials**](BillOfMaterials.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

