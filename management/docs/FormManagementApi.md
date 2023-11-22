# \FormManagementApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForm**](FormManagementApi.md#CreateForm) | **Post** /environments/{environmentID}/forms | CREATE Form
[**DeleteForm**](FormManagementApi.md#DeleteForm) | **Delete** /environments/{environmentID}/forms/{formID} | DELETE Form
[**ReadAllForms**](FormManagementApi.md#ReadAllForms) | **Get** /environments/{environmentID}/forms | READ All Forms
[**ReadForm**](FormManagementApi.md#ReadForm) | **Get** /environments/{environmentID}/forms/{formID} | READ One Form
[**UpdateForm**](FormManagementApi.md#UpdateForm) | **Put** /environments/{environmentID}/forms/{formID} | UPDATE Form



## CreateForm

> Form CreateForm(ctx, environmentID).Form(form).Execute()

CREATE Form

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
    form := *openapiclient.NewForm("Name_example", openapiclient.EnumFormCategory("CUSTOM"), *openapiclient.NewFormComponents([]openapiclient.FormField{openapiclient.FormField{FormFieldCheckbox: openapiclient.NewFormFieldCheckbox(openapiclient.EnumFormFieldType("TEXT"), *openapiclient.NewFormFieldCommonPosition(int32(123), int32(123)), "Key_example", false, openapiclient.EnumFormElementLayout("HORIZONTAL"), []string{"Options_example"})}}), false, false) // Form |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FormManagementApi.CreateForm(context.Background(), environmentID).Form(form).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormManagementApi.CreateForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateForm`: Form
    fmt.Fprintf(os.Stdout, "Response from `FormManagementApi.CreateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **form** | [**Form**](Form.md) |  | 

### Return type

[**Form**](Form.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForm

> DeleteForm(ctx, environmentID, formID).Execute()

DELETE Form

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
    formID := "formID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FormManagementApi.DeleteForm(context.Background(), environmentID, formID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormManagementApi.DeleteForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**formID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFormRequest struct via the builder pattern


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


## ReadAllForms

> EntityArray ReadAllForms(ctx, environmentID).Execute()

READ All Forms

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
    resp, r, err := apiClient.FormManagementApi.ReadAllForms(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormManagementApi.ReadAllForms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllForms`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `FormManagementApi.ReadAllForms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllFormsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ReadForm

> Form ReadForm(ctx, environmentID, formID).Include(include).Execute()

READ One Form

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
    formID := "formID_example" // string | 
    include := openapiclient.EnumFormsIncludeParameter("components") // EnumFormsIncludeParameter |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FormManagementApi.ReadForm(context.Background(), environmentID, formID).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormManagementApi.ReadForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadForm`: Form
    fmt.Fprintf(os.Stdout, "Response from `FormManagementApi.ReadForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**formID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | [**EnumFormsIncludeParameter**](EnumFormsIncludeParameter.md) |  | 

### Return type

[**Form**](Form.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateForm

> Form UpdateForm(ctx, environmentID, formID).Form(form).Execute()

UPDATE Form

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
    formID := "formID_example" // string | 
    form := *openapiclient.NewForm("Name_example", openapiclient.EnumFormCategory("CUSTOM"), *openapiclient.NewFormComponents([]openapiclient.FormField{openapiclient.FormField{FormFieldCheckbox: openapiclient.NewFormFieldCheckbox(openapiclient.EnumFormFieldType("TEXT"), *openapiclient.NewFormFieldCommonPosition(int32(123), int32(123)), "Key_example", false, openapiclient.EnumFormElementLayout("HORIZONTAL"), []string{"Options_example"})}}), false, false) // Form |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FormManagementApi.UpdateForm(context.Background(), environmentID, formID).Form(form).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FormManagementApi.UpdateForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateForm`: Form
    fmt.Fprintf(os.Stdout, "Response from `FormManagementApi.UpdateForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**formID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **form** | [**Form**](Form.md) |  | 

### Return type

[**Form**](Form.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

