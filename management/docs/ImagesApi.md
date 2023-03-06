# \ImagesApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImagesApi.md#CreateImage) | **Post** /v1/environments/{environmentID}/images | CREATE Image
[**DeleteImage**](ImagesApi.md#DeleteImage) | **Delete** /v1/environments/{environmentID}/images/{imgID} | DELETE Image
[**ReadImage**](ImagesApi.md#ReadImage) | **Get** /v1/environments/{environmentID}/images/{imgID} | READ Image



## CreateImage

> Image CreateImage(ctx, environmentID).ContentType(contentType).ContentDisposition(contentDisposition).Body(body).Execute()

CREATE Image

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
    contentType := "image/jpeg" // string |  (optional)
    contentDisposition := "attachment; filename=name.jpg" // string |  (optional)
    body := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.CreateImage(context.Background(), environmentID).ContentType(contentType).ContentDisposition(contentDisposition).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.CreateImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **contentDisposition** | **string** |  | 
 **body** | ***os.File** |  | 

### Return type

[**Image**](Image.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, environmentID, imgID).Execute()

DELETE Image

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
    imgID := "imgID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ImagesApi.DeleteImage(context.Background(), environmentID, imgID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.DeleteImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**imgID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## ReadImage

> Image ReadImage(ctx, environmentID, imgID).Execute()

READ Image

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
    imgID := "imgID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImagesApi.ReadImage(context.Background(), environmentID, imgID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImagesApi.ReadImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadImage`: Image
    fmt.Fprintf(os.Stdout, "Response from `ImagesApi.ReadImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**imgID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Image**](Image.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

