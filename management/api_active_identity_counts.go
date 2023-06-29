/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ActiveIdentityCountsApiService ActiveIdentityCountsApi service
type ActiveIdentityCountsApiService service

type ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	environmentID string
	filter *string
	limit *int32
	order *string
}

func (r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Filter(filter string) ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Limit(limit int32) ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Order(order string) ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r)
}

/*
EnvironmentsEnvironmentIDActiveIdentityCountsGet READ Active Identity Counts (Deprecated)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest
*/
func (a *ActiveIdentityCountsApiService) EnvironmentsEnvironmentIDActiveIdentityCountsGet(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	return ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) EnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) (*http.Response, error) {
	_, response, error := processResponse(
		func() (interface{}, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r)
			return nil, resp, err
		},
	)
	return response, error
}
			
func (a *ActiveIdentityCountsApiService) internalEnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r ApiEnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.EnvironmentsEnvironmentIDActiveIdentityCountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
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

type ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	environmentID string
	filter *string
	limit *int32
	order *string
	samplePeriod *string
}

func (r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Filter(filter string) ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Limit(limit int32) ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Order(order string) ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) SamplePeriod(samplePeriod string) ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}

func (r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r)
}

/*
EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet READ Active Identity Counts by Date Range

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest
*/
func (a *ActiveIdentityCountsApiService) EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	return ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) (*http.Response, error) {
	_, response, error := processResponse(
		func() (interface{}, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r)
			return nil, resp, err
		},
	)
	return response, error
}
			
func (a *ActiveIdentityCountsApiService) internalEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r ApiEnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.samplePeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "samplePeriod", r.samplePeriod, "")
	}
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

type ApiReadActiveIdentityCountRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	organizationID string
	licenseID string
	aggregatedBy *string
	limit *int32
	order *string
}

func (r ApiReadActiveIdentityCountRequest) AggregatedBy(aggregatedBy string) ApiReadActiveIdentityCountRequest {
	r.aggregatedBy = &aggregatedBy
	return r
}

func (r ApiReadActiveIdentityCountRequest) Limit(limit int32) ApiReadActiveIdentityCountRequest {
	r.limit = &limit
	return r
}

func (r ApiReadActiveIdentityCountRequest) Order(order string) ApiReadActiveIdentityCountRequest {
	r.order = &order
	return r
}

func (r ApiReadActiveIdentityCountRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadActiveIdentityCountExecute(r)
}

/*
ReadActiveIdentityCount READ Active Identity Counts by License

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationID
 @param licenseID
 @return ApiReadActiveIdentityCountRequest
*/
func (a *ActiveIdentityCountsApiService) ReadActiveIdentityCount(ctx context.Context, organizationID string, licenseID string) ApiReadActiveIdentityCountRequest {
	return ApiReadActiveIdentityCountRequest{
		ApiService: a,
		ctx: ctx,
		organizationID: organizationID,
		licenseID: licenseID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) ReadActiveIdentityCountExecute(r ApiReadActiveIdentityCountRequest) (*http.Response, error) {
	_, response, error := processResponse(
		func() (interface{}, *http.Response, error) {
			resp, err := r.ApiService.internalReadActiveIdentityCountExecute(r)
			return nil, resp, err
		},
	)
	return response, error
}
			
func (a *ActiveIdentityCountsApiService) internalReadActiveIdentityCountExecute(r ApiReadActiveIdentityCountRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.ReadActiveIdentityCount")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationID}/licenses/{licenseID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationID"+"}", url.PathEscape(parameterValueToString(r.organizationID, "organizationID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"licenseID"+"}", url.PathEscape(parameterValueToString(r.licenseID, "licenseID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aggregatedBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aggregatedBy", r.aggregatedBy, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
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
