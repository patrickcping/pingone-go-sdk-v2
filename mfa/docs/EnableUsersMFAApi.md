# \EnableUsersMFAApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReadUserMFAEnabled**](EnableUsersMFAApi.md#ReadUserMFAEnabled) | **Get** /environments/{environmentID}/users/{userID}/mfaEnabled | READ User MFA Enabled
[**UpdateUserMFAEnabled**](EnableUsersMFAApi.md#UpdateUserMFAEnabled) | **Put** /environments/{environmentID}/users/{userID}/mfaEnabled | UPDATE User MFA Enabled



## ReadUserMFAEnabled

> UserMFAEnabled ReadUserMFAEnabled(ctx, environmentID, userID).Execute()

READ User MFA Enabled

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
	userID := "userID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnableUsersMFAApi.ReadUserMFAEnabled(context.Background(), environmentID, userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnableUsersMFAApi.ReadUserMFAEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadUserMFAEnabled`: UserMFAEnabled
	fmt.Fprintf(os.Stdout, "Response from `EnableUsersMFAApi.ReadUserMFAEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadUserMFAEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserMFAEnabled**](UserMFAEnabled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserMFAEnabled

> UserMFAEnabled UpdateUserMFAEnabled(ctx, environmentID, userID).UserMFAEnabled(userMFAEnabled).Execute()

UPDATE User MFA Enabled

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
	userID := "userID_example" // string | 
	userMFAEnabled := *openapiclient.NewUserMFAEnabled(false) // UserMFAEnabled |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnableUsersMFAApi.UpdateUserMFAEnabled(context.Background(), environmentID, userID).UserMFAEnabled(userMFAEnabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnableUsersMFAApi.UpdateUserMFAEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserMFAEnabled`: UserMFAEnabled
	fmt.Fprintf(os.Stdout, "Response from `EnableUsersMFAApi.UpdateUserMFAEnabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserMFAEnabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userMFAEnabled** | [**UserMFAEnabled**](UserMFAEnabled.md) |  | 

### Return type

[**UserMFAEnabled**](UserMFAEnabled.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

