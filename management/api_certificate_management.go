/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

package management

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CertificateManagementApiService CertificateManagementApi service
type CertificateManagementApiService service

type ApiCreateCertificateFromFileRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	usageType                  *string
	file                       **[]byte
	contentType                *EnumCreateCertificateAcceptHeader
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateCertificateFromFileRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiCreateCertificateFromFileRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateCertificateFromFileRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiCreateCertificateFromFileRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiCreateCertificateFromFileRequest) UsageType(usageType string) ApiCreateCertificateFromFileRequest {
	r.usageType = &usageType
	return r
}

func (r ApiCreateCertificateFromFileRequest) File(file *[]byte) ApiCreateCertificateFromFileRequest {
	r.file = &file
	return r
}

func (r ApiCreateCertificateFromFileRequest) ContentType(contentType EnumCreateCertificateAcceptHeader) ApiCreateCertificateFromFileRequest {
	r.contentType = &contentType
	return r
}

func (r ApiCreateCertificateFromFileRequest) Execute() (*Certificate, *http.Response, error) {
	return r.ApiService.CreateCertificateFromFileExecute(r)
}

/*
CreateCertificateFromFile CREATE Certificate with PKCS7 or PEM File

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiCreateCertificateFromFileRequest
*/
func (a *CertificateManagementApiService) CreateCertificateFromFile(ctx context.Context, environmentID string) ApiCreateCertificateFromFileRequest {
	return ApiCreateCertificateFromFileRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) CreateCertificateFromFileExecute(r ApiCreateCertificateFromFileRequest) (*Certificate, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *Certificate
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateCertificateFromFileExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalCreateCertificateFromFileExecute(r ApiCreateCertificateFromFileRequest) (*Certificate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.CreateCertificateFromFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/certificates"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.usageType == nil {
		return localVarReturnValue, nil, reportError("usageType is required and must be specified")
	}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "usageType", r.usageType, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileBytes *[]byte
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}

	fileLocalVarFormFileName = "file"

	fileLocalVarFileBytes = *r.file

	formFiles = append(formFiles, formFile{fileBytes: *fileLocalVarFileBytes, fileName: fileLocalVarFormFileName, formFileName: fileLocalVarFormFileName})

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		if localVarHTTPResponse.StatusCode >= 300 {
			newErr := &GenericOpenAPIError{
				body:  localVarBody,
				error: localVarHTTPResponse.Status,
			}
			if localVarHTTPResponse.StatusCode == 400 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 401 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 403 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				baseDelay := time.Second
				// check if environment created recently - DOCS-8830
				retryEnvironmentResponse, retryVarHTTPResponse, err := a.client.EnvironmentsApi.ReadOneEnvironment(r.ctx, r.environmentID).Execute()
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				if retryVarHTTPResponse.StatusCode == 200 && retryEnvironmentResponse != nil && retryEnvironmentResponse.CreatedAt != nil {
					// Check if the retryEnvironmentResponse.CreatedAt is within the last 17 seconds
					createdAt, nestedErr := time.Parse(time.RFC3339, *retryEnvironmentResponse.CreatedAt)
					if nestedErr != nil {
						slog.Error("Invalid RFC3339 string", "environment created at", *retryEnvironmentResponse.CreatedAt)
					} else {
						if time.Since(createdAt) < 17*time.Second {
							slog.Debug("The environment was created within the last 17 seconds, retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
							// Retry the request
							delay, nestedErr := calculateExponentialBackoff(i, baseDelay)
							if nestedErr != nil {
								slog.Error("Invalid backoff delay duration", "error", nestedErr, "baseDelay", baseDelay, "retry", false)
							} else {
								time.Sleep(delay)
								continue
							}
						}
					}
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 405 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 429 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 500 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		break
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateKeyRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	contentType                *string
	certificate                *Certificate
	file                       **[]byte
	password                   *string
	usageType                  *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateKeyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiCreateKeyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateKeyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiCreateKeyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiCreateKeyRequest) ContentType(contentType string) ApiCreateKeyRequest {
	r.contentType = &contentType
	return r
}

func (r ApiCreateKeyRequest) Certificate(certificate Certificate) ApiCreateKeyRequest {
	r.certificate = &certificate
	return r
}

func (r ApiCreateKeyRequest) UsageType(usageType string) ApiCreateKeyRequest {
	r.usageType = &usageType
	return r
}

func (r ApiCreateKeyRequest) Password(password string) ApiCreateKeyRequest {
	r.password = &password
	return r
}

func (r ApiCreateKeyRequest) File(file *[]byte) ApiCreateKeyRequest {
	r.file = &file
	return r
}

func (r ApiCreateKeyRequest) Execute() (*Certificate, *http.Response, error) {
	return r.ApiService.CreateKeyExecute(r)
}

/*
CreateKey CREATE Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiCreateKeyRequest
*/
func (a *CertificateManagementApiService) CreateKey(ctx context.Context, environmentID string) ApiCreateKeyRequest {
	return ApiCreateKeyRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) CreateKeyExecute(r ApiCreateKeyRequest) (*Certificate, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *Certificate
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateKeyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalCreateKeyExecute(r ApiCreateKeyRequest) (*Certificate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.CreateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}

	if localVarHeaderParams["Content-Type"] == "multipart/form-data" {
		if r.usageType == nil {
			return localVarReturnValue, nil, reportError("usageType is required and must be specified")
		}
		if r.file == nil {
			return localVarReturnValue, nil, reportError("file is required and must be specified")
		}
	}

	if localVarHeaderParams["Content-Type"] == "application/json" {
		if r.certificate == nil {
			return localVarReturnValue, nil, reportError("certificate is required and must be specified")
		}
	}

	// body params
	var req *http.Request
	if localVarHeaderParams["Content-Type"] == "multipart/form-data" {

		localVarFormParams.Add("usageType", parameterValueToString(*r.usageType, ""))
		if r.password != nil {
			localVarFormParams.Add("password", parameterValueToString(*r.password, ""))
		}
		var fileLocalVarFormFileName string
		var fileLocalVarFileBytes *[]byte

		fileLocalVarFormFileName = "file"

		fileLocalVarFileBytes = *r.file

		formFiles = append(formFiles, formFile{fileBytes: *fileLocalVarFileBytes, fileName: fileLocalVarFormFileName, formFileName: fileLocalVarFormFileName})
	}

	if localVarHeaderParams["Content-Type"] == "application/json" {
		localVarPostBody = r.certificate
	}

	req, err = a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteCertificateRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	certID                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeleteCertificateRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiDeleteCertificateRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeleteCertificateRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiDeleteCertificateRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiDeleteCertificateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCertificateExecute(r)
}

/*
DeleteCertificate DELETE Certificate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param certID
	@return ApiDeleteCertificateRequest
*/
func (a *CertificateManagementApiService) DeleteCertificate(ctx context.Context, environmentID string, certID string) ApiDeleteCertificateRequest {
	return ApiDeleteCertificateRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		certID:        certID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) DeleteCertificateExecute(r ApiDeleteCertificateRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalDeleteCertificateExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalDeleteCertificateExecute(r ApiDeleteCertificateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.DeleteCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/certificates/{certID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certID"+"}", url.PathEscape(parameterValueToString(r.certID, "certID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteKeyRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeleteKeyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiDeleteKeyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeleteKeyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiDeleteKeyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiDeleteKeyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteKeyExecute(r)
}

/*
DeleteKey DELETE Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiDeleteKeyRequest
*/
func (a *CertificateManagementApiService) DeleteKey(ctx context.Context, environmentID string, keyID string) ApiDeleteKeyRequest {
	return ApiDeleteKeyRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) DeleteKeyExecute(r ApiDeleteKeyRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalDeleteKeyExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalDeleteKeyExecute(r ApiDeleteKeyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.DeleteKey")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiExportCSRRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	accept                     *EnumCSRExportHeader
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiExportCSRRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiExportCSRRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiExportCSRRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiExportCSRRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiExportCSRRequest) Accept(accept EnumCSRExportHeader) ApiExportCSRRequest {
	r.accept = &accept
	return r
}

func (r ApiExportCSRRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ExportCSRExecute(r)
}

/*
ExportCSR Export a certificate signing request (CSR)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiExportCSRRequest
*/
func (a *CertificateManagementApiService) ExportCSR(ctx context.Context, environmentID string, keyID string) ApiExportCSRRequest {
	return ApiExportCSRRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
//
//	@return string
func (a *CertificateManagementApiService) ExportCSRExecute(r ApiExportCSRRequest) (string, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue string
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalExportCSRExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalExportCSRExecute(r ApiExportCSRRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.ExportCSR")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}/csr"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/pkcs10", "application/x-pem-file"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCertificateRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	certID                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificateRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetCertificateRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificateRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetCertificateRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetCertificateRequest) Execute() (*Certificate, *http.Response, error) {
	return r.ApiService.GetCertificateExecute(r)
}

/*
GetCertificate GET Certificate

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param certID
	@return ApiGetCertificateRequest
*/
func (a *CertificateManagementApiService) GetCertificate(ctx context.Context, environmentID string, certID string) ApiGetCertificateRequest {
	return ApiGetCertificateRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		certID:        certID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) GetCertificateExecute(r ApiGetCertificateRequest) (*Certificate, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *Certificate
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetCertificateExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetCertificateExecute(r ApiGetCertificateRequest) (*Certificate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/certificates/{certID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certID"+"}", url.PathEscape(parameterValueToString(r.certID, "certID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCertificateApplicationsRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	certID                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificateApplicationsRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetCertificateApplicationsRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificateApplicationsRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetCertificateApplicationsRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetCertificateApplicationsRequest) Execute() EntityArrayPagedIterator {
	return r.ApiService.GetCertificateApplicationsExecute(r)
}

func (r ApiGetCertificateApplicationsRequest) ExecuteInitialPage() (*EntityArray, *http.Response, error) {
	return r.ApiService.GetCertificateApplicationsExecuteInitialPage(r)
}

/*
GetCertificateApplications GET Certificate Applications

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param certID
	@return ApiGetCertificateApplicationsRequest
*/
func (a *CertificateManagementApiService) GetCertificateApplications(ctx context.Context, environmentID string, certID string) ApiGetCertificateApplicationsRequest {
	return ApiGetCertificateApplicationsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		certID:        certID,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *CertificateManagementApiService) GetCertificateApplicationsExecute(r ApiGetCertificateApplicationsRequest) EntityArrayPagedIterator {
	return a.client.paginationIterator(r.ctx, r.ExecuteInitialPage)
}

func (a *CertificateManagementApiService) GetCertificateApplicationsExecuteInitialPage(r ApiGetCertificateApplicationsRequest) (*EntityArray, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EntityArray
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetCertificateApplicationsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetCertificateApplicationsExecute(r ApiGetCertificateApplicationsRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetCertificateApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/certificates/{certID}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certID"+"}", url.PathEscape(parameterValueToString(r.certID, "certID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetCertificatesRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificatesRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetCertificatesRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetCertificatesRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetCertificatesRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetCertificatesRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.GetCertificatesExecute(r)
}

/*
GetCertificates GET Certificates

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiGetCertificatesRequest
*/
func (a *CertificateManagementApiService) GetCertificates(ctx context.Context, environmentID string) ApiGetCertificatesRequest {
	return ApiGetCertificatesRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *CertificateManagementApiService) GetCertificatesExecute(r ApiGetCertificatesRequest) (*EntityArray, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EntityArray
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetCertificatesExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetCertificatesExecute(r ApiGetCertificatesRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/certificates"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetKeyRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	accept                     *EnumGetKeyAcceptHeader
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetKeyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetKeyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetKeyRequest) Accept(accept EnumGetKeyAcceptHeader) ApiGetKeyRequest {
	r.accept = &accept
	return r
}

func (r ApiGetKeyRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.GetKeyExecute(r)
}

/*
GetKey GET Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiGetKeyRequest
*/
func (a *CertificateManagementApiService) GetKey(ctx context.Context, environmentID string, keyID string) ApiGetKeyRequest {
	return ApiGetKeyRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) GetKeyExecute(r ApiGetKeyRequest) (interface{}, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue interface{}
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetKeyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetKeyExecute(r ApiGetKeyRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/x-x509-ca-cert", "application/x-pkcs7-certificates"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	if localVarHTTPResponse.Header.Get("Content-Type") == "application/x-x509-ca-cert" {
		return string(localVarBody), localVarHTTPResponse, nil
	}

	if localVarHTTPResponse.Header.Get("Content-Type") == "application/x-pkcs7-certificates" {
		return localVarBody, localVarHTTPResponse, nil
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetKeyApplicationsRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeyApplicationsRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetKeyApplicationsRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeyApplicationsRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetKeyApplicationsRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetKeyApplicationsRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.GetKeyApplicationsExecute(r)
}

/*
GetKeyApplications GET Key Applications

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiGetKeyApplicationsRequest
*/
func (a *CertificateManagementApiService) GetKeyApplications(ctx context.Context, environmentID string, keyID string) ApiGetKeyApplicationsRequest {
	return ApiGetKeyApplicationsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *CertificateManagementApiService) GetKeyApplicationsExecute(r ApiGetKeyApplicationsRequest) (*EntityArray, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EntityArray
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetKeyApplicationsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetKeyApplicationsExecute(r ApiGetKeyApplicationsRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetKeyApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}/applications"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetKeysRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeysRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiGetKeysRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiGetKeysRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiGetKeysRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiGetKeysRequest) Execute() (*EntityArray, *http.Response, error) {
	return r.ApiService.GetKeysExecute(r)
}

/*
GetKeys GET Keys

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiGetKeysRequest
*/
func (a *CertificateManagementApiService) GetKeys(ctx context.Context, environmentID string) ApiGetKeysRequest {
	return ApiGetKeysRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *CertificateManagementApiService) GetKeysExecute(r ApiGetKeysRequest) (*EntityArray, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EntityArray
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalGetKeysExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalGetKeysExecute(r ApiGetKeysRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.GetKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiImportCSRResponseRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	file                       **[]byte
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiImportCSRResponseRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiImportCSRResponseRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiImportCSRResponseRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiImportCSRResponseRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiImportCSRResponseRequest) File(file *[]byte) ApiImportCSRResponseRequest {
	r.file = &file
	return r
}

func (r ApiImportCSRResponseRequest) Execute() (*Certificate, *http.Response, error) {
	return r.ApiService.ImportCSRResponseExecute(r)
}

/*
ImportCSRResponse Import Certificate Authority (CA) Response to a CSR

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiImportCSRResponseRequest
*/
func (a *CertificateManagementApiService) ImportCSRResponse(ctx context.Context, environmentID string, keyID string) ApiImportCSRResponseRequest {
	return ApiImportCSRResponseRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) ImportCSRResponseExecute(r ApiImportCSRResponseRequest) (*Certificate, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *Certificate
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalImportCSRResponseExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalImportCSRResponseExecute(r ApiImportCSRResponseRequest) (*Certificate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.ImportCSRResponse")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}/csr"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.file == nil {
		return localVarReturnValue, nil, reportError("file is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	var fileLocalVarFormFileName string
	var fileLocalVarFileBytes *[]byte

	fileLocalVarFormFileName = "file"

	fileLocalVarFileBytes = *r.file

	formFiles = append(formFiles, formFile{fileBytes: *fileLocalVarFileBytes, fileName: fileLocalVarFormFileName, formFileName: fileLocalVarFormFileName})
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateKeyRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	keyID                      string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	certificateKeyUpdate       *CertificateKeyUpdate
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateKeyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiUpdateKeyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateKeyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiUpdateKeyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiUpdateKeyRequest) CertificateKeyUpdate(certificateKeyUpdate CertificateKeyUpdate) ApiUpdateKeyRequest {
	r.certificateKeyUpdate = &certificateKeyUpdate
	return r
}

func (r ApiUpdateKeyRequest) Execute() (*Certificate, *http.Response, error) {
	return r.ApiService.UpdateKeyExecute(r)
}

/*
UpdateKey UPDATE Key

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param keyID
	@return ApiUpdateKeyRequest
*/
func (a *CertificateManagementApiService) UpdateKey(ctx context.Context, environmentID string, keyID string) ApiUpdateKeyRequest {
	return ApiUpdateKeyRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		keyID:         keyID,
	}
}

// Execute executes the request
//
//	@return Certificate
func (a *CertificateManagementApiService) UpdateKeyExecute(r ApiUpdateKeyRequest) (*Certificate, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *Certificate
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateKeyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CertificateManagementApiService) internalUpdateKeyExecute(r ApiUpdateKeyRequest) (*Certificate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Certificate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.UpdateKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/keys/{keyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"keyID"+"}", url.PathEscape(parameterValueToString(r.keyID, "keyID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.certificateKeyUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDDecryptionsPostRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	body                       *map[string]interface{}
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDDecryptionsPostRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDDecryptionsPostRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) Body(body map[string]interface{}) ApiEnvironmentsEnvironmentIDDecryptionsPostRequest {
	r.body = &body
	return r
}

func (r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDDecryptionsPostExecute(r)
}

/*
EnvironmentsEnvironmentIDDecryptionsPost DECRYPT Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiEnvironmentsEnvironmentIDDecryptionsPostRequest
*/
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDDecryptionsPost(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDDecryptionsPostRequest {
	return ApiEnvironmentsEnvironmentIDDecryptionsPostRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDDecryptionsPostExecute(r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDDecryptionsPostExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalEnvironmentsEnvironmentIDDecryptionsPostExecute(r ApiEnvironmentsEnvironmentIDDecryptionsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.EnvironmentsEnvironmentIDDecryptionsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/decryptions"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDEncryptionsPostRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	body                       *map[string]interface{}
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDEncryptionsPostRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDEncryptionsPostRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) Body(body map[string]interface{}) ApiEnvironmentsEnvironmentIDEncryptionsPostRequest {
	r.body = &body
	return r
}

func (r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDEncryptionsPostExecute(r)
}

/*
EnvironmentsEnvironmentIDEncryptionsPost ENCRYPT Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiEnvironmentsEnvironmentIDEncryptionsPostRequest
*/
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDEncryptionsPost(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDEncryptionsPostRequest {
	return ApiEnvironmentsEnvironmentIDEncryptionsPostRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDEncryptionsPostExecute(r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDEncryptionsPostExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalEnvironmentsEnvironmentIDEncryptionsPostExecute(r ApiEnvironmentsEnvironmentIDEncryptionsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.EnvironmentsEnvironmentIDEncryptionsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/encryptions"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDSigningsPostRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	body                       *map[string]interface{}
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDSigningsPostRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDSigningsPostRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDSigningsPostRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDSigningsPostRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDSigningsPostRequest) Body(body map[string]interface{}) ApiEnvironmentsEnvironmentIDSigningsPostRequest {
	r.body = &body
	return r
}

func (r ApiEnvironmentsEnvironmentIDSigningsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDSigningsPostExecute(r)
}

/*
EnvironmentsEnvironmentIDSigningsPost SIGN Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiEnvironmentsEnvironmentIDSigningsPostRequest
*/
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDSigningsPost(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDSigningsPostRequest {
	return ApiEnvironmentsEnvironmentIDSigningsPostRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDSigningsPostExecute(r ApiEnvironmentsEnvironmentIDSigningsPostRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDSigningsPostExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalEnvironmentsEnvironmentIDSigningsPostExecute(r ApiEnvironmentsEnvironmentIDSigningsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.EnvironmentsEnvironmentIDSigningsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/signings"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnvironmentsEnvironmentIDVerificationsPostRequest struct {
	ctx                        context.Context
	ApiService                 *CertificateManagementApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	body                       *map[string]interface{}
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiEnvironmentsEnvironmentIDVerificationsPostRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiEnvironmentsEnvironmentIDVerificationsPostRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) Body(body map[string]interface{}) ApiEnvironmentsEnvironmentIDVerificationsPostRequest {
	r.body = &body
	return r
}

func (r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDVerificationsPostExecute(r)
}

/*
EnvironmentsEnvironmentIDVerificationsPost VERIFY Signed Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiEnvironmentsEnvironmentIDVerificationsPostRequest
*/
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDVerificationsPost(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDVerificationsPostRequest {
	return ApiEnvironmentsEnvironmentIDVerificationsPostRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *CertificateManagementApiService) EnvironmentsEnvironmentIDVerificationsPostExecute(r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDVerificationsPostExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *CertificateManagementApiService) internalEnvironmentsEnvironmentIDVerificationsPostExecute(r ApiEnvironmentsEnvironmentIDVerificationsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CertificateManagementApiService.EnvironmentsEnvironmentIDVerificationsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/verifications"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
