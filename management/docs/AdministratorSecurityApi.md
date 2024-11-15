# \AdministratorSecurityApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadAdministratorSecurity**](AdministratorSecurityApi.md#ReadAdministratorSecurity) | **Get** /environments/{environmentID}/adminConfig | READ Administrator Security
[**UpdateAdministratorSecurity**](AdministratorSecurityApi.md#UpdateAdministratorSecurity) | **Put** /environments/{environmentID}/adminConfig | UPDATE Administrator Security



## ReadAdministratorSecurity

> AdministratorSecurity ReadAdministratorSecurity(ctx, environmentID).Execute()

READ Administrator Security

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
    resp, r, err := apiClient.AdministratorSecurityApi.ReadAdministratorSecurity(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorSecurityApi.ReadAdministratorSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAdministratorSecurity`: AdministratorSecurity
    fmt.Fprintf(os.Stdout, "Response from `AdministratorSecurityApi.ReadAdministratorSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAdministratorSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AdministratorSecurity**](AdministratorSecurity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdministratorSecurity

> AdministratorSecurity UpdateAdministratorSecurity(ctx, environmentID).AdministratorSecurity(administratorSecurity).Execute()

UPDATE Administrator Security

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
    administratorSecurity := *openapiclient.NewAdministratorSecurity(openapiclient.EnumAdministratorSecurityAuthenticationMethod("PINGONE"), false) // AdministratorSecurity |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdministratorSecurityApi.UpdateAdministratorSecurity(context.Background(), environmentID).AdministratorSecurity(administratorSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdministratorSecurityApi.UpdateAdministratorSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdministratorSecurity`: AdministratorSecurity
    fmt.Fprintf(os.Stdout, "Response from `AdministratorSecurityApi.UpdateAdministratorSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdministratorSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **administratorSecurity** | [**AdministratorSecurity**](AdministratorSecurity.md) |  | 

### Return type

[**AdministratorSecurity**](AdministratorSecurity.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

