# \FIDO2PolicyApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFido2Policy**](FIDO2PolicyApi.md#CreateFido2Policy) | **Post** /v1/environments/{environmentID}/fido2Policies | CREATE FIDO2 Policy
[**DeleteFido2Policy**](FIDO2PolicyApi.md#DeleteFido2Policy) | **Delete** /v1/environments/{environmentID}/fido2Policies/{fido2PolicyID} | DELETE FIDO2 Policy
[**ReadFido2Policies**](FIDO2PolicyApi.md#ReadFido2Policies) | **Get** /v1/environments/{environmentID}/fido2Policies | READ FIDO2 Policies
[**ReadOneFido2Policy**](FIDO2PolicyApi.md#ReadOneFido2Policy) | **Get** /v1/environments/{environmentID}/fido2Policies/{fido2PolicyID} | READ One FIDO2 Policy
[**UpdateFIDO2Policy**](FIDO2PolicyApi.md#UpdateFIDO2Policy) | **Put** /v1/environments/{environmentID}/fido2Policies/{fido2PolicyID} | UPDATE FIDO2 Policy



## CreateFido2Policy

> FIDO2Policy CreateFido2Policy(ctx, environmentID).FIDO2Policy(fIDO2Policy).Execute()

CREATE FIDO2 Policy

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
    fIDO2Policy := *openapiclient.NewFIDO2Policy(openapiclient.EnumFIDO2PolicyAttestationRequirements("DIRECT"), openapiclient.EnumFIDO2PolicyAuthenticatorAttachment("PLATFORM"), *openapiclient.NewFIDO2PolicyBackupEligibility(false, false), "DeviceDisplayName_example", openapiclient.EnumFIDO2PolicyDiscoverableCredentials("DISCOURAGED"), *openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirements([]openapiclient.FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner{*openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner("Id_example")}, false, openapiclient.EnumFIDO2PolicyMDSAuthenticatorOption("NONE")), "Name_example", "RelyingPartyId_example", *openapiclient.NewFIDO2PolicyUserDisplayNameAttributes([]openapiclient.FIDO2PolicyUserDisplayNameAttributesAttributesInner{*openapiclient.NewFIDO2PolicyUserDisplayNameAttributesAttributesInner("Name_example")}), *openapiclient.NewFIDO2PolicyUserVerification(false, openapiclient.EnumFIDO2PolicyUserVerificationOption("REQUIRED"))) // FIDO2Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDO2PolicyApi.CreateFido2Policy(context.Background(), environmentID).FIDO2Policy(fIDO2Policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.CreateFido2Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFido2Policy`: FIDO2Policy
    fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.CreateFido2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFido2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fIDO2Policy** | [**FIDO2Policy**](FIDO2Policy.md) |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFido2Policy

> DeleteFido2Policy(ctx, environmentID, fido2PolicyID).Execute()

DELETE FIDO2 Policy

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
    fido2PolicyID := "fido2PolicyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FIDO2PolicyApi.DeleteFido2Policy(context.Background(), environmentID, fido2PolicyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.DeleteFido2Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFido2PolicyRequest struct via the builder pattern


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


## ReadFido2Policies

> EntityArray ReadFido2Policies(ctx, environmentID).Execute()

READ FIDO2 Policies

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
    resp, r, err := apiClient.FIDO2PolicyApi.ReadFido2Policies(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.ReadFido2Policies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadFido2Policies`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.ReadFido2Policies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadFido2PoliciesRequest struct via the builder pattern


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


## ReadOneFido2Policy

> FIDO2Policy ReadOneFido2Policy(ctx, environmentID, fido2PolicyID).Expand(expand).Execute()

READ One FIDO2 Policy

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
    fido2PolicyID := "fido2PolicyID_example" // string | 
    expand := "expand_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDO2PolicyApi.ReadOneFido2Policy(context.Background(), environmentID, fido2PolicyID).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.ReadOneFido2Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneFido2Policy`: FIDO2Policy
    fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.ReadOneFido2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneFido2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFIDO2Policy

> FIDO2Policy UpdateFIDO2Policy(ctx, environmentID, fido2PolicyID).FIDO2Policy(fIDO2Policy).Execute()

UPDATE FIDO2 Policy

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
    fido2PolicyID := "fido2PolicyID_example" // string | 
    fIDO2Policy := *openapiclient.NewFIDO2Policy(openapiclient.EnumFIDO2PolicyAttestationRequirements("DIRECT"), openapiclient.EnumFIDO2PolicyAuthenticatorAttachment("PLATFORM"), *openapiclient.NewFIDO2PolicyBackupEligibility(false, false), "DeviceDisplayName_example", openapiclient.EnumFIDO2PolicyDiscoverableCredentials("DISCOURAGED"), *openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirements([]openapiclient.FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner{*openapiclient.NewFIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner("Id_example")}, false, openapiclient.EnumFIDO2PolicyMDSAuthenticatorOption("NONE")), "Name_example", "RelyingPartyId_example", *openapiclient.NewFIDO2PolicyUserDisplayNameAttributes([]openapiclient.FIDO2PolicyUserDisplayNameAttributesAttributesInner{*openapiclient.NewFIDO2PolicyUserDisplayNameAttributesAttributesInner("Name_example")}), *openapiclient.NewFIDO2PolicyUserVerification(false, openapiclient.EnumFIDO2PolicyUserVerificationOption("REQUIRED"))) // FIDO2Policy |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FIDO2PolicyApi.UpdateFIDO2Policy(context.Background(), environmentID, fido2PolicyID).FIDO2Policy(fIDO2Policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FIDO2PolicyApi.UpdateFIDO2Policy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFIDO2Policy`: FIDO2Policy
    fmt.Fprintf(os.Stdout, "Response from `FIDO2PolicyApi.UpdateFIDO2Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**fido2PolicyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFIDO2PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fIDO2Policy** | [**FIDO2Policy**](FIDO2Policy.md) |  | 

### Return type

[**FIDO2Policy**](FIDO2Policy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

