/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

package management

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PasswordPoliciesApiService PasswordPoliciesApi service
type PasswordPoliciesApiService service

type ApiCreatePasswordPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *PasswordPoliciesApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	passwordPolicy             *PasswordPolicy
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreatePasswordPolicyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiCreatePasswordPolicyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreatePasswordPolicyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiCreatePasswordPolicyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiCreatePasswordPolicyRequest) PasswordPolicy(passwordPolicy PasswordPolicy) ApiCreatePasswordPolicyRequest {
	r.passwordPolicy = &passwordPolicy
	return r
}

func (r ApiCreatePasswordPolicyRequest) Execute() (*PasswordPolicy, *http.Response, error) {
	return r.ApiService.CreatePasswordPolicyExecute(r)
}

/*
CreatePasswordPolicy CREATE Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiCreatePasswordPolicyRequest
*/
func (a *PasswordPoliciesApiService) CreatePasswordPolicy(ctx context.Context, environmentID string) ApiCreatePasswordPolicyRequest {
	return ApiCreatePasswordPolicyRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return PasswordPolicy
func (a *PasswordPoliciesApiService) CreatePasswordPolicyExecute(r ApiCreatePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *PasswordPolicy
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreatePasswordPolicyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *PasswordPoliciesApiService) internalCreatePasswordPolicyExecute(r ApiCreatePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPoliciesApiService.CreatePasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/passwordPolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.passwordPolicy == nil {
		return localVarReturnValue, nil, reportError("passwordPolicy is required and must be specified")
	}

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
	localVarPostBody = r.passwordPolicy
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

	err = decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

func decode(v interface{}, b []byte, contentType string) (err error) {
	if jsonCheck.MatchString(contentType) {
		bStr := string(b[:])
		bStr = strings.ReplaceAll(bStr, "~!@#$%^&*()-_=+[]{}|;:,.<>/?", "specialchar")

		err := json.Unmarshal([]byte(bStr), v)
		if err != nil { // simple model
			return err
		}

		return nil
	}
	return errors.New("undefined response type")
}

type ApiDeletePasswordPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *PasswordPoliciesApiService
	environmentID              string
	passwordPolicyID           string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeletePasswordPolicyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiDeletePasswordPolicyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiDeletePasswordPolicyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiDeletePasswordPolicyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiDeletePasswordPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePasswordPolicyExecute(r)
}

/*
DeletePasswordPolicy DELETE Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param passwordPolicyID
	@return ApiDeletePasswordPolicyRequest
*/
func (a *PasswordPoliciesApiService) DeletePasswordPolicy(ctx context.Context, environmentID string, passwordPolicyID string) ApiDeletePasswordPolicyRequest {
	return ApiDeletePasswordPolicyRequest{
		ApiService:       a,
		ctx:              ctx,
		environmentID:    environmentID,
		passwordPolicyID: passwordPolicyID,
	}
}

// Execute executes the request
func (a *PasswordPoliciesApiService) DeletePasswordPolicyExecute(r ApiDeletePasswordPolicyRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalDeletePasswordPolicyExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *PasswordPoliciesApiService) internalDeletePasswordPolicyExecute(r ApiDeletePasswordPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPoliciesApiService.DeletePasswordPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/passwordPolicies/{passwordPolicyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"passwordPolicyID"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyID, "passwordPolicyID")), -1)

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

type ApiReadAllPasswordPoliciesRequest struct {
	ctx                        context.Context
	ApiService                 *PasswordPoliciesApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadAllPasswordPoliciesRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiReadAllPasswordPoliciesRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadAllPasswordPoliciesRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiReadAllPasswordPoliciesRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiReadAllPasswordPoliciesRequest) Execute() EntityArrayPagedIterator {
	return r.ApiService.ReadAllPasswordPoliciesExecute(r)
}

func (r ApiReadAllPasswordPoliciesRequest) ExecuteInitialPage() (*EntityArray, *http.Response, error) {
	return r.ApiService.ReadAllPasswordPoliciesExecuteInitialPage(r)
}

/*
ReadAllPasswordPolicies READ All Password Policies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiReadAllPasswordPoliciesRequest
*/
func (a *PasswordPoliciesApiService) ReadAllPasswordPolicies(ctx context.Context, environmentID string) ApiReadAllPasswordPoliciesRequest {
	return ApiReadAllPasswordPoliciesRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *PasswordPoliciesApiService) ReadAllPasswordPoliciesExecute(r ApiReadAllPasswordPoliciesRequest) EntityArrayPagedIterator {
	return a.client.paginationIterator(r.ctx, r.ExecuteInitialPage)
}

func (a *PasswordPoliciesApiService) ReadAllPasswordPoliciesExecuteInitialPage(r ApiReadAllPasswordPoliciesRequest) (*EntityArray, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *EntityArray
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalReadAllPasswordPoliciesExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *PasswordPoliciesApiService) internalReadAllPasswordPoliciesExecute(r ApiReadAllPasswordPoliciesRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPoliciesApiService.ReadAllPasswordPolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/passwordPolicies"
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

	err = decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReadOnePasswordPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *PasswordPoliciesApiService
	environmentID              string
	passwordPolicyID           string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadOnePasswordPolicyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiReadOnePasswordPolicyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadOnePasswordPolicyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiReadOnePasswordPolicyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiReadOnePasswordPolicyRequest) Execute() (*PasswordPolicy, *http.Response, error) {
	return r.ApiService.ReadOnePasswordPolicyExecute(r)
}

/*
ReadOnePasswordPolicy READ One Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param passwordPolicyID
	@return ApiReadOnePasswordPolicyRequest
*/
func (a *PasswordPoliciesApiService) ReadOnePasswordPolicy(ctx context.Context, environmentID string, passwordPolicyID string) ApiReadOnePasswordPolicyRequest {
	return ApiReadOnePasswordPolicyRequest{
		ApiService:       a,
		ctx:              ctx,
		environmentID:    environmentID,
		passwordPolicyID: passwordPolicyID,
	}
}

// Execute executes the request
//
//	@return PasswordPolicy
func (a *PasswordPoliciesApiService) ReadOnePasswordPolicyExecute(r ApiReadOnePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *PasswordPolicy
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalReadOnePasswordPolicyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *PasswordPoliciesApiService) internalReadOnePasswordPolicyExecute(r ApiReadOnePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPoliciesApiService.ReadOnePasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/passwordPolicies/{passwordPolicyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"passwordPolicyID"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyID, "passwordPolicyID")), -1)

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

	err = decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdatePasswordPolicyRequest struct {
	ctx                        context.Context
	ApiService                 *PasswordPoliciesApiService
	environmentID              string
	passwordPolicyID           string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	passwordPolicy             *PasswordPolicy
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdatePasswordPolicyRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiUpdatePasswordPolicyRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdatePasswordPolicyRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiUpdatePasswordPolicyRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiUpdatePasswordPolicyRequest) PasswordPolicy(passwordPolicy PasswordPolicy) ApiUpdatePasswordPolicyRequest {
	r.passwordPolicy = &passwordPolicy
	return r
}

func (r ApiUpdatePasswordPolicyRequest) Execute() (*PasswordPolicy, *http.Response, error) {
	return r.ApiService.UpdatePasswordPolicyExecute(r)
}

/*
UpdatePasswordPolicy UPDATE Password Policy

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param passwordPolicyID
	@return ApiUpdatePasswordPolicyRequest
*/
func (a *PasswordPoliciesApiService) UpdatePasswordPolicy(ctx context.Context, environmentID string, passwordPolicyID string) ApiUpdatePasswordPolicyRequest {
	return ApiUpdatePasswordPolicyRequest{
		ApiService:       a,
		ctx:              ctx,
		environmentID:    environmentID,
		passwordPolicyID: passwordPolicyID,
	}
}

// Execute executes the request
//
//	@return PasswordPolicy
func (a *PasswordPoliciesApiService) UpdatePasswordPolicyExecute(r ApiUpdatePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		err                 error
		response            *http.Response
		localVarReturnValue *PasswordPolicy
	)

	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdatePasswordPolicyExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *PasswordPoliciesApiService) internalUpdatePasswordPolicyExecute(r ApiUpdatePasswordPolicyRequest) (*PasswordPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PasswordPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PasswordPoliciesApiService.UpdatePasswordPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/passwordPolicies/{passwordPolicyID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"passwordPolicyID"+"}", url.PathEscape(parameterValueToString(r.passwordPolicyID, "passwordPolicyID")), -1)

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
	localVarPostBody = r.passwordPolicy
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

	err = decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
