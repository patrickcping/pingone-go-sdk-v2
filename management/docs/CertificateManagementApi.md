# \CertificateManagementApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCertificateFromFile**](CertificateManagementApi.md#CreateCertificateFromFile) | **Post** /v1/environments/{environmentID}/certificates | CREATE Certificate with PKCS7 or PEM File
[**CreateKey**](CertificateManagementApi.md#CreateKey) | **Post** /v1/environments/{environmentID}/keys | CREATE Key
[**DeleteCertificate**](CertificateManagementApi.md#DeleteCertificate) | **Delete** /v1/environments/{environmentID}/certificates/{certID} | DELETE Certificate
[**DeleteKey**](CertificateManagementApi.md#DeleteKey) | **Delete** /v1/environments/{environmentID}/keys/{keyID} | DELETE Key
[**ExportCSR**](CertificateManagementApi.md#ExportCSR) | **Get** /v1/environments/{environmentID}/keys/{keyID}/csr | Export a certificate signing request (CSR)
[**GetCertificate**](CertificateManagementApi.md#GetCertificate) | **Get** /v1/environments/{environmentID}/certificates/{certID} | GET Certificate
[**GetCertificateApplications**](CertificateManagementApi.md#GetCertificateApplications) | **Get** /v1/environments/{environmentID}/certificates/{certID}/applications | GET Certificate Applications
[**GetCertificates**](CertificateManagementApi.md#GetCertificates) | **Get** /v1/environments/{environmentID}/certificates | GET Certificates
[**GetKey**](CertificateManagementApi.md#GetKey) | **Get** /v1/environments/{environmentID}/keys/{keyID} | GET Key
[**GetKeyApplications**](CertificateManagementApi.md#GetKeyApplications) | **Get** /v1/environments/{environmentID}/keys/{keyID}/applications | GET Key Applications
[**GetKeys**](CertificateManagementApi.md#GetKeys) | **Get** /v1/environments/{environmentID}/keys | GET Keys
[**ImportCSRResponse**](CertificateManagementApi.md#ImportCSRResponse) | **Post** /v1/environments/{environmentID}/keys/{keyID}/csr | Import Certificate Authority (CA) Response to a CSR
[**UpdateKey**](CertificateManagementApi.md#UpdateKey) | **Put** /v1/environments/{environmentID}/keys/{keyID} | UPDATE Key
[**V1EnvironmentsEnvironmentIDDecryptionsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDDecryptionsPost) | **Post** /v1/environments/{environmentID}/decryptions | DECRYPT Data
[**V1EnvironmentsEnvironmentIDEncryptionsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDEncryptionsPost) | **Post** /v1/environments/{environmentID}/encryptions | ENCRYPT Data
[**V1EnvironmentsEnvironmentIDSigningsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDSigningsPost) | **Post** /v1/environments/{environmentID}/signings | SIGN Data
[**V1EnvironmentsEnvironmentIDVerificationsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDVerificationsPost) | **Post** /v1/environments/{environmentID}/verifications | VERIFY Signed Data



## CreateCertificateFromFile

> Certificate CreateCertificateFromFile(ctx, environmentID).UsageType(usageType).File(file).ContentType(contentType).Execute()

CREATE Certificate with PKCS7 or PEM File

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    usageType := "usageType_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 
    contentType := openapiclient.EnumCreateCertificateAcceptHeader("application/x-pkcs7-certificates") // EnumCreateCertificateAcceptHeader |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.CreateCertificateFromFile(context.Background(), environmentID).UsageType(usageType).File(file).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.CreateCertificateFromFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCertificateFromFile`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.CreateCertificateFromFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCertificateFromFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usageType** | **string** |  | 
 **file** | ***os.File** |  | 
 **contentType** | [**EnumCreateCertificateAcceptHeader**](EnumCreateCertificateAcceptHeader.md) |  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> Certificate CreateKey(ctx, environmentID).ContentType(contentType).Certificate(certificate).Execute()

CREATE Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    contentType := "application/x-pkcs12-certificates" // string |  (optional)
    certificate := *openapiclient.NewCertificate(openapiclient.EnumCertificateKeyAlgorithm("RSA"), int32(123), "Name_example", openapiclient.EnumCertificateKeySignagureAlgorithm("SHA224withRSA"), "SubjectDN_example", openapiclient.EnumCertificateKeyUsageType("ENCRYPTION"), int32(123)) // Certificate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.CreateKey(context.Background(), environmentID).ContentType(contentType).Certificate(certificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.CreateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKey`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.CreateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **certificate** | [**Certificate**](Certificate.md) |  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCertificate

> DeleteCertificate(ctx, environmentID, certID).Execute()

DELETE Certificate

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.DeleteCertificate(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.DeleteCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCertificateRequest struct via the builder pattern


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


## DeleteKey

> DeleteKey(ctx, environmentID, keyID).Execute()

DELETE Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.DeleteKey(context.Background(), environmentID, keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.DeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyRequest struct via the builder pattern


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


## ExportCSR

> string ExportCSR(ctx, environmentID, keyID).Accept(accept).Execute()

Export a certificate signing request (CSR)

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 
    accept := openapiclient.EnumCSRExportHeader("application/pkcs10") // EnumCSRExportHeader |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.ExportCSR(context.Background(), environmentID, keyID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.ExportCSR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportCSR`: string
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.ExportCSR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCSRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | [**EnumCSRExportHeader**](EnumCSRExportHeader.md) |  | 

### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/pkcs10, application/x-pem-file

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificate

> Certificate GetCertificate(ctx, environmentID, certID).Execute()

GET Certificate

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetCertificate(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetCertificate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificate`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetCertificate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCertificateApplications

> EntityArray GetCertificateApplications(ctx, environmentID, certID).Execute()

GET Certificate Applications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetCertificateApplications(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetCertificateApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificateApplications`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetCertificateApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificateApplicationsRequest struct via the builder pattern


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


## GetCertificates

> EntityArray GetCertificates(ctx, environmentID).Execute()

GET Certificates

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetCertificates(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetCertificates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCertificates`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetCertificates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCertificatesRequest struct via the builder pattern


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


## GetKey

> Certificate GetKey(ctx, environmentID, keyID).Accept(accept).Execute()

GET Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 
    accept := openapiclient.EnumGetKeyAcceptHeader("application/json") // EnumGetKeyAcceptHeader |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetKey(context.Background(), environmentID, keyID).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKey`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | [**EnumGetKeyAcceptHeader**](EnumGetKeyAcceptHeader.md) |  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-x509-ca-cert, application/x-pkcs7-certificates

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyApplications

> EntityArray GetKeyApplications(ctx, environmentID, keyID).Execute()

GET Key Applications

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetKeyApplications(context.Background(), environmentID, keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetKeyApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeyApplications`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetKeyApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyApplicationsRequest struct via the builder pattern


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


## GetKeys

> EntityArray GetKeys(ctx, environmentID).Execute()

GET Keys

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.GetKeys(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.GetKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeys`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.GetKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeysRequest struct via the builder pattern


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


## ImportCSRResponse

> Certificate ImportCSRResponse(ctx, environmentID, keyID).File(file).Execute()

Import Certificate Authority (CA) Response to a CSR

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 
    file := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.ImportCSRResponse(context.Background(), environmentID, keyID).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.ImportCSRResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportCSRResponse`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.ImportCSRResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCSRResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** |  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKey

> Certificate UpdateKey(ctx, environmentID, keyID).CertificateKeyUpdate(certificateKeyUpdate).Execute()

UPDATE Key

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    keyID := "keyID_example" // string | 
    certificateKeyUpdate := *openapiclient.NewCertificateKeyUpdate(false, openapiclient.EnumCertificateKeyUsageType("ENCRYPTION")) // CertificateKeyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.UpdateKey(context.Background(), environmentID, keyID).CertificateKeyUpdate(certificateKeyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.UpdateKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKey`: Certificate
    fmt.Fprintf(os.Stdout, "Response from `CertificateManagementApi.UpdateKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **certificateKeyUpdate** | [**CertificateKeyUpdate**](CertificateKeyUpdate.md) |  | 

### Return type

[**Certificate**](Certificate.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDDecryptionsPost

> V1EnvironmentsEnvironmentIDDecryptionsPost(ctx, environmentID).Body(body).Execute()

DECRYPT Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDDecryptionsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDDecryptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDDecryptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDEncryptionsPost

> V1EnvironmentsEnvironmentIDEncryptionsPost(ctx, environmentID).Body(body).Execute()

ENCRYPT Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDEncryptionsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDEncryptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEncryptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDSigningsPost

> V1EnvironmentsEnvironmentIDSigningsPost(ctx, environmentID).Body(body).Execute()

SIGN Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDSigningsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDSigningsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSigningsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDVerificationsPost

> V1EnvironmentsEnvironmentIDVerificationsPost(ctx, environmentID).Body(body).Execute()

VERIFY Signed Data

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDVerificationsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDVerificationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDVerificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

