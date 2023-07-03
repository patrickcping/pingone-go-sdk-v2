# \DeviceAuthenticationPolicyApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceAuthenticationPolicies**](DeviceAuthenticationPolicyApi.md#CreateDeviceAuthenticationPolicies) | **Post** /environments/{environmentID}/deviceAuthenticationPolicies | CREATE Device Authentication Policy
[**DeleteDeviceAuthenticationPolicy**](DeviceAuthenticationPolicyApi.md#DeleteDeviceAuthenticationPolicy) | **Delete** /environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | DELETE Device Authentication Policy
[**ReadDeviceAuthenticationPolicies**](DeviceAuthenticationPolicyApi.md#ReadDeviceAuthenticationPolicies) | **Get** /environments/{environmentID}/deviceAuthenticationPolicies | READ Device Authentication Policies
[**ReadOneDeviceAuthenticationPolicy**](DeviceAuthenticationPolicyApi.md#ReadOneDeviceAuthenticationPolicy) | **Get** /environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | READ One Device Authentication Policy
[**UpdateDeviceAuthenticationPolicy**](DeviceAuthenticationPolicyApi.md#UpdateDeviceAuthenticationPolicy) | **Put** /environments/{environmentID}/deviceAuthenticationPolicies/{deviceAuthenticationPolicyID} | UPDATE Device Authentication Policy



## CreateDeviceAuthenticationPolicies

> DeviceAuthenticationPolicyPost CreateDeviceAuthenticationPolicies(ctx, environmentID).DeviceAuthenticationPolicyPost(deviceAuthenticationPolicyPost).Execute()

CREATE Device Authentication Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    deviceAuthenticationPolicyPost := openapiclient.DeviceAuthenticationPolicyPost{DeviceAuthenticationPolicy: openapiclient.NewDeviceAuthenticationPolicy("Name_example", *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyMobile(false, *openapiclient.NewDeviceAuthenticationPolicyMobileOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))), *openapiclient.NewDeviceAuthenticationPolicyMobileOtpWindow(*openapiclient.NewDeviceAuthenticationPolicyMobileOtpWindowStepSize(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyTotp(false, *openapiclient.NewDeviceAuthenticationPolicyTotpOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), false, false)} // DeviceAuthenticationPolicyPost |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies(context.Background(), environmentID).DeviceAuthenticationPolicyPost(deviceAuthenticationPolicyPost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceAuthenticationPolicies`: DeviceAuthenticationPolicyPost
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.CreateDeviceAuthenticationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceAuthenticationPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceAuthenticationPolicyPost** | [**DeviceAuthenticationPolicyPost**](DeviceAuthenticationPolicyPost.md) |  | 

### Return type

[**DeviceAuthenticationPolicyPost**](DeviceAuthenticationPolicyPost.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json, application/vnd.pingidentity.deviceAuthenticationPolicy.fido2.migrate+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceAuthenticationPolicy

> DeleteDeviceAuthenticationPolicy(ctx, environmentID, deviceAuthenticationPolicyID).Execute()

DELETE Device Authentication Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    deviceAuthenticationPolicyID := "deviceAuthenticationPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeviceAuthenticationPolicyApi.DeleteDeviceAuthenticationPolicy(context.Background(), environmentID, deviceAuthenticationPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.DeleteDeviceAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**deviceAuthenticationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceAuthenticationPolicyRequest struct via the builder pattern


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


## ReadDeviceAuthenticationPolicies

> EntityArray ReadDeviceAuthenticationPolicies(ctx, environmentID).Execute()

READ Device Authentication Policies

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDeviceAuthenticationPolicies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.ReadDeviceAuthenticationPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDeviceAuthenticationPoliciesRequest struct via the builder pattern


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


## ReadOneDeviceAuthenticationPolicy

> DeviceAuthenticationPolicy ReadOneDeviceAuthenticationPolicy(ctx, environmentID, deviceAuthenticationPolicyID).Execute()

READ One Device Authentication Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    deviceAuthenticationPolicyID := "deviceAuthenticationPolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy(context.Background(), environmentID, deviceAuthenticationPolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDeviceAuthenticationPolicy`: DeviceAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.ReadOneDeviceAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**deviceAuthenticationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDeviceAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceAuthenticationPolicy

> DeviceAuthenticationPolicy UpdateDeviceAuthenticationPolicy(ctx, environmentID, deviceAuthenticationPolicyID).DeviceAuthenticationPolicy(deviceAuthenticationPolicy).Execute()

UPDATE Device Authentication Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/mfa"
)

func main() {
    environmentID := "environmentID_example" // string | 
    deviceAuthenticationPolicyID := "deviceAuthenticationPolicyID_example" // string | 
    deviceAuthenticationPolicy := *openapiclient.NewDeviceAuthenticationPolicy("Name_example", *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyOfflineDevice(false, *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpLifeTime(int32(123), openapiclient.EnumTimeUnit("MINUTES")), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyMobile(false, *openapiclient.NewDeviceAuthenticationPolicyMobileOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))), *openapiclient.NewDeviceAuthenticationPolicyMobileOtpWindow(*openapiclient.NewDeviceAuthenticationPolicyMobileOtpWindowStepSize(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), *openapiclient.NewDeviceAuthenticationPolicyTotp(false, *openapiclient.NewDeviceAuthenticationPolicyTotpOtp(*openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(int32(123), *openapiclient.NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown(int32(123), openapiclient.EnumTimeUnit("MINUTES"))))), false, false) // DeviceAuthenticationPolicy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceAuthenticationPolicyApi.UpdateDeviceAuthenticationPolicy(context.Background(), environmentID, deviceAuthenticationPolicyID).DeviceAuthenticationPolicy(deviceAuthenticationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceAuthenticationPolicyApi.UpdateDeviceAuthenticationPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceAuthenticationPolicy`: DeviceAuthenticationPolicy
    fmt.Fprintf(os.Stdout, "Response from `DeviceAuthenticationPolicyApi.UpdateDeviceAuthenticationPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**deviceAuthenticationPolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceAuthenticationPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deviceAuthenticationPolicy** | [**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md) |  | 

### Return type

[**DeviceAuthenticationPolicy**](DeviceAuthenticationPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

