# \RiskPoliciesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskPolicySet**](RiskPoliciesApi.md#CreateRiskPolicySet) | **Post** /environments/{environmentID}/riskPolicySets | CREATE Risk Policy Set
[**DeleteRiskPolicySet**](RiskPoliciesApi.md#DeleteRiskPolicySet) | **Delete** /environments/{environmentID}/riskPolicySets/{riskPolicySetID} | DELETE Risk Policy Set 
[**ReadOneRiskPolicySet**](RiskPoliciesApi.md#ReadOneRiskPolicySet) | **Get** /environments/{environmentID}/riskPolicySets/{riskPolicySetID} | READ One Risk Policy Set
[**ReadRiskPolicySets**](RiskPoliciesApi.md#ReadRiskPolicySets) | **Get** /environments/{environmentID}/riskPolicySets | READ Risk Policy Sets
[**UpdateRiskPolicySet**](RiskPoliciesApi.md#UpdateRiskPolicySet) | **Put** /environments/{environmentID}/riskPolicySets/{riskPolicySetID} | UPDATE Risk Policy Set



## CreateRiskPolicySet

> RiskPolicySet CreateRiskPolicySet(ctx, environmentID).RiskPolicySet(riskPolicySet).Execute()

CREATE Risk Policy Set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 
	riskPolicySet := *openapiclient.NewRiskPolicySet("Name_example") // RiskPolicySet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskPoliciesApi.CreateRiskPolicySet(context.Background(), environmentID).RiskPolicySet(riskPolicySet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.CreateRiskPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRiskPolicySet`: RiskPolicySet
	fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.CreateRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **riskPolicySet** | [**RiskPolicySet**](RiskPolicySet.md) |  | 

### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskPolicySet

> DeleteRiskPolicySet(ctx, environmentID, riskPolicySetID).Execute()

DELETE Risk Policy Set 

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 
	riskPolicySetID := "riskPolicySetID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RiskPoliciesApi.DeleteRiskPolicySet(context.Background(), environmentID, riskPolicySetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.DeleteRiskPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskPolicySetRequest struct via the builder pattern


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


## ReadOneRiskPolicySet

> RiskPolicySet ReadOneRiskPolicySet(ctx, environmentID, riskPolicySetID).Execute()

READ One Risk Policy Set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 
	riskPolicySetID := "riskPolicySetID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskPoliciesApi.ReadOneRiskPolicySet(context.Background(), environmentID, riskPolicySetID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.ReadOneRiskPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadOneRiskPolicySet`: RiskPolicySet
	fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.ReadOneRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRiskPolicySets

READ Risk Policy Sets

### Paged Response (Recommended)

> PagedIterator[ReadRiskPolicySets200Response] ReadRiskPolicySets(ctx, environmentID).Execute()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.RiskPoliciesApi.ReadRiskPolicySets(context.Background(), environmentID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.ReadRiskPolicySets``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `ReadRiskPolicySets` page iteration: ReadRiskPolicySets200Response
		fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.ReadRiskPolicySets` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> ReadRiskPolicySets200Response ReadRiskPolicySets(ctx, environmentID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskPoliciesApi.ReadRiskPolicySets(context.Background(), environmentID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.ReadRiskPolicySets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadRiskPolicySets`: ReadRiskPolicySets200Response
	fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.ReadRiskPolicySets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRiskPolicySetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

Page Iterator: PagedIterator[[**ReadRiskPolicySets200Response**](ReadRiskPolicySets200Response.md)]

PagedIterator[ReadRiskPolicySets200Response] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**ReadRiskPolicySets200Response**](ReadRiskPolicySets200Response.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**ReadRiskPolicySets200Response**](ReadRiskPolicySets200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRiskPolicySet

> RiskPolicySet UpdateRiskPolicySet(ctx, environmentID, riskPolicySetID).RiskPolicySet(riskPolicySet).Execute()

UPDATE Risk Policy Set

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/patrickcping/pingone-go-sdk-v2/risk"
)

func main() {
	environmentID := "environmentID_example" // string | 
	riskPolicySetID := "riskPolicySetID_example" // string | 
	riskPolicySet := *openapiclient.NewRiskPolicySet("Name_example") // RiskPolicySet |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskPoliciesApi.UpdateRiskPolicySet(context.Background(), environmentID, riskPolicySetID).RiskPolicySet(riskPolicySet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskPoliciesApi.UpdateRiskPolicySet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRiskPolicySet`: RiskPolicySet
	fmt.Fprintf(os.Stdout, "Response from `RiskPoliciesApi.UpdateRiskPolicySet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**riskPolicySetID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRiskPolicySetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **riskPolicySet** | [**RiskPolicySet**](RiskPolicySet.md) |  | 

### Return type

[**RiskPolicySet**](RiskPolicySet.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

