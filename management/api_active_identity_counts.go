/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ActiveIdentityCountsApiService ActiveIdentityCountsApi service
type ActiveIdentityCountsApiService service

type ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	environmentID string
	filter *string
	limit *int32
	order *string
}

func (r ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Order(order string) ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r)
}

/*
V1EnvironmentsEnvironmentIDActiveIdentityCountsGet READ Active Identity Counts (Deprecated)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest
*/
func (a *ActiveIdentityCountsApiService) V1EnvironmentsEnvironmentIDActiveIdentityCountsGet(ctx context.Context, environmentID string) ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest {
	return ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) V1EnvironmentsEnvironmentIDActiveIdentityCountsGetExecute(r ApiV1EnvironmentsEnvironmentIDActiveIdentityCountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.V1EnvironmentsEnvironmentIDActiveIdentityCountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
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
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	environmentID string
	filter *string
	limit *int32
	order *string
	samplePeriod *string
}

func (r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Order(order string) ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) SamplePeriod(samplePeriod string) ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r)
}

/*
V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet READ Active Identity Counts by Date Range

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest
*/
func (a *ActiveIdentityCountsApiService) V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet(ctx context.Context, environmentID string) ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest {
	return ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetExecute(r ApiV1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.V1EnvironmentsEnvironmentIDMetricsActiveIdentityCountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
	}
	if r.samplePeriod != nil {
		localVarQueryParams.Add("samplePeriod", parameterToString(*r.samplePeriod, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
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
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest struct {
	ctx context.Context
	ApiService *ActiveIdentityCountsApiService
	organizationID string
	licenseID string
	aggregatedBy *string
	limit *int32
	order *string
}

func (r ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) AggregatedBy(aggregatedBy string) ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.aggregatedBy = &aggregatedBy
	return r
}

func (r ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Limit(limit int32) ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Order(order string) ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	r.order = &order
	return r
}

func (r ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetExecute(r)
}

/*
V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet READ Active Identity Counts by License

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationID
 @param licenseID
 @return ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest
*/
func (a *ActiveIdentityCountsApiService) V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet(ctx context.Context, organizationID string, licenseID string) ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest {
	return ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest{
		ApiService: a,
		ctx: ctx,
		organizationID: organizationID,
		licenseID: licenseID,
	}
}

// Execute executes the request
func (a *ActiveIdentityCountsApiService) V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetExecute(r ApiV1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActiveIdentityCountsApiService.V1OrganizationsOrganizationIDLicensesLicenseIDMetricsActiveIdentityCountsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationID}/licenses/{licenseID}/metrics/activeIdentityCounts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationID"+"}", url.PathEscape(parameterToString(r.organizationID, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"licenseID"+"}", url.PathEscape(parameterToString(r.licenseID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aggregatedBy != nil {
		localVarQueryParams.Add("aggregatedBy", parameterToString(*r.aggregatedBy, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.order != nil {
		localVarQueryParams.Add("order", parameterToString(*r.order, ""))
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
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
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
