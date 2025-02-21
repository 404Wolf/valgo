/*
Val Town API

Val Town’s public API  OpenAPI JSON endpoint:  https://api.val.town/openapi.json

API version: 1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package valgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BetaAPIService BetaAPI service
type BetaAPIService service

type ApiBranchesGetRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	projectId string
	branchId string
}

func (r ApiBranchesGetRequest) Execute() (*Branch, *http.Response, error) {
	return r.ApiService.BranchesGetExecute(r)
}

/*
BranchesGet Method for BranchesGet

[BETA] Get a branch by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @param branchId Id of a branch
 @return ApiBranchesGetRequest
*/
func (a *BetaAPIService) BranchesGet(ctx context.Context, projectId string, branchId string) ApiBranchesGetRequest {
	return ApiBranchesGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		branchId: branchId,
	}
}

// Execute executes the request
//  @return Branch
func (a *BetaAPIService) BranchesGetExecute(r ApiBranchesGetRequest) (*Branch, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Branch
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.BranchesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}/branches/{branch_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"branch_id"+"}", url.PathEscape(parameterValueToString(r.branchId, "branchId")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiBranchesListRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	offset *int32
	limit *int32
	projectId string
}

// Number of items to skip in order to deliver paginated results
func (r ApiBranchesListRequest) Offset(offset int32) ApiBranchesListRequest {
	r.offset = &offset
	return r
}

// Maximum items to return in each paginated response
func (r ApiBranchesListRequest) Limit(limit int32) ApiBranchesListRequest {
	r.limit = &limit
	return r
}

func (r ApiBranchesListRequest) Execute() (*BranchesList200Response, *http.Response, error) {
	return r.ApiService.BranchesListExecute(r)
}

/*
BranchesList Method for BranchesList

[BETA] List all branches for a project

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @return ApiBranchesListRequest
*/
func (a *BetaAPIService) BranchesList(ctx context.Context, projectId string) ApiBranchesListRequest {
	return ApiBranchesListRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return BranchesList200Response
func (a *BetaAPIService) BranchesListExecute(r ApiBranchesListRequest) (*BranchesList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BranchesList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.BranchesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}/branches"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}
	if *r.offset < 0 {
		return localVarReturnValue, nil, reportError("offset must be greater than 0")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 100 {
		return localVarReturnValue, nil, reportError("limit must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiFilesContentGetRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	projectId string
	path string
	version *int32
	branchId *string
	ifMatch *string
	ifUnmodifiedSince *string
	ifNoneMatch *string
	ifModifiedSince *string
	cacheControl *string
}

// Specific branch version to query
func (r ApiFilesContentGetRequest) Version(version int32) ApiFilesContentGetRequest {
	r.version = &version
	return r
}

// Id to query
func (r ApiFilesContentGetRequest) BranchId(branchId string) ApiFilesContentGetRequest {
	r.branchId = &branchId
	return r
}

func (r ApiFilesContentGetRequest) IfMatch(ifMatch string) ApiFilesContentGetRequest {
	r.ifMatch = &ifMatch
	return r
}

func (r ApiFilesContentGetRequest) IfUnmodifiedSince(ifUnmodifiedSince string) ApiFilesContentGetRequest {
	r.ifUnmodifiedSince = &ifUnmodifiedSince
	return r
}

func (r ApiFilesContentGetRequest) IfNoneMatch(ifNoneMatch string) ApiFilesContentGetRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

func (r ApiFilesContentGetRequest) IfModifiedSince(ifModifiedSince string) ApiFilesContentGetRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiFilesContentGetRequest) CacheControl(cacheControl string) ApiFilesContentGetRequest {
	r.cacheControl = &cacheControl
	return r
}

func (r ApiFilesContentGetRequest) Execute() (interface{}, *http.Response, error) {
	return r.ApiService.FilesContentGetExecute(r)
}

/*
FilesContentGet Method for FilesContentGet

Download file content

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @param path URL encoded path to a file or directory (e.g. 'dir%2Fsubdir%2Ffile.ts')
 @return ApiFilesContentGetRequest
*/
func (a *BetaAPIService) FilesContentGet(ctx context.Context, projectId string, path string) ApiFilesContentGetRequest {
	return ApiFilesContentGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		path: path,
	}
}

// Execute executes the request
//  @return interface{}
func (a *BetaAPIService) FilesContentGetExecute(r ApiFilesContentGetRequest) (interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.FilesContentGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}/files/{path}/content"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.branchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "branch_id", r.branchId, "")
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
	if r.ifMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Match", r.ifMatch, "")
	}
	if r.ifUnmodifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Unmodified-Since", r.ifUnmodifiedSince, "")
	}
	if r.ifNoneMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Modified-Since", r.ifModifiedSince, "")
	}
	if r.cacheControl != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Cache-Control", r.cacheControl, "")
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
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 304 {
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiProjectFilesGetRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	offset *int32
	limit *int32
	projectId string
	path string
	version *int32
	branchId *string
}

// Number of items to skip in order to deliver paginated results
func (r ApiProjectFilesGetRequest) Offset(offset int32) ApiProjectFilesGetRequest {
	r.offset = &offset
	return r
}

// Maximum items to return in each paginated response
func (r ApiProjectFilesGetRequest) Limit(limit int32) ApiProjectFilesGetRequest {
	r.limit = &limit
	return r
}

// Specific branch version to query
func (r ApiProjectFilesGetRequest) Version(version int32) ApiProjectFilesGetRequest {
	r.version = &version
	return r
}

// Id to query
func (r ApiProjectFilesGetRequest) BranchId(branchId string) ApiProjectFilesGetRequest {
	r.branchId = &branchId
	return r
}

func (r ApiProjectFilesGetRequest) Execute() (*RootProjectFilesGet200Response, *http.Response, error) {
	return r.ApiService.ProjectFilesGetExecute(r)
}

/*
ProjectFilesGet Method for ProjectFilesGet

Get metadata for files and directories in a project at the specified path

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @param path URL encoded path to a file or directory (e.g. 'dir%2Fsubdir%2Ffile.ts')
 @return ApiProjectFilesGetRequest
*/
func (a *BetaAPIService) ProjectFilesGet(ctx context.Context, projectId string, path string) ApiProjectFilesGetRequest {
	return ApiProjectFilesGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		path: path,
	}
}

// Execute executes the request
//  @return RootProjectFilesGet200Response
func (a *BetaAPIService) ProjectFilesGetExecute(r ApiProjectFilesGetRequest) (*RootProjectFilesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RootProjectFilesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.ProjectFilesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", url.PathEscape(parameterValueToString(r.path, "path")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}
	if *r.offset < 0 {
		return localVarReturnValue, nil, reportError("offset must be greater than 0")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 100 {
		return localVarReturnValue, nil, reportError("limit must be less than 100")
	}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.branchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "branch_id", r.branchId, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiProjectsGetRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	projectId string
}

func (r ApiProjectsGetRequest) Execute() (*Project, *http.Response, error) {
	return r.ApiService.ProjectsGetExecute(r)
}

/*
ProjectsGet Method for ProjectsGet

[BETA] Get a project by id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @return ApiProjectsGetRequest
*/
func (a *BetaAPIService) ProjectsGet(ctx context.Context, projectId string) ApiProjectsGetRequest {
	return ApiProjectsGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return Project
func (a *BetaAPIService) ProjectsGetExecute(r ApiProjectsGetRequest) (*Project, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Project
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.ProjectsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiProjectsListRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	offset *int32
	limit *int32
}

// Number of items to skip in order to deliver paginated results
func (r ApiProjectsListRequest) Offset(offset int32) ApiProjectsListRequest {
	r.offset = &offset
	return r
}

// Maximum items to return in each paginated response
func (r ApiProjectsListRequest) Limit(limit int32) ApiProjectsListRequest {
	r.limit = &limit
	return r
}

func (r ApiProjectsListRequest) Execute() (*MeProjects200Response, *http.Response, error) {
	return r.ApiService.ProjectsListExecute(r)
}

/*
ProjectsList Method for ProjectsList

[BETA] Lists all public projects

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProjectsListRequest
*/
func (a *BetaAPIService) ProjectsList(ctx context.Context) ApiProjectsListRequest {
	return ApiProjectsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MeProjects200Response
func (a *BetaAPIService) ProjectsListExecute(r ApiProjectsListRequest) (*MeProjects200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MeProjects200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.ProjectsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}
	if *r.offset < 0 {
		return localVarReturnValue, nil, reportError("offset must be greater than 0")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 100 {
		return localVarReturnValue, nil, reportError("limit must be less than 100")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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

type ApiRootProjectFilesGetRequest struct {
	ctx context.Context
	ApiService *BetaAPIService
	offset *int32
	limit *int32
	projectId string
	version *int32
	branchId *string
	recursive *bool
}

// Number of items to skip in order to deliver paginated results
func (r ApiRootProjectFilesGetRequest) Offset(offset int32) ApiRootProjectFilesGetRequest {
	r.offset = &offset
	return r
}

// Maximum items to return in each paginated response
func (r ApiRootProjectFilesGetRequest) Limit(limit int32) ApiRootProjectFilesGetRequest {
	r.limit = &limit
	return r
}

// Specific branch version to query
func (r ApiRootProjectFilesGetRequest) Version(version int32) ApiRootProjectFilesGetRequest {
	r.version = &version
	return r
}

// Id to query
func (r ApiRootProjectFilesGetRequest) BranchId(branchId string) ApiRootProjectFilesGetRequest {
	r.branchId = &branchId
	return r
}

// Whether to recursively list all files in the project
func (r ApiRootProjectFilesGetRequest) Recursive(recursive bool) ApiRootProjectFilesGetRequest {
	r.recursive = &recursive
	return r
}

func (r ApiRootProjectFilesGetRequest) Execute() (*RootProjectFilesGet200Response, *http.Response, error) {
	return r.ApiService.RootProjectFilesGetExecute(r)
}

/*
RootProjectFilesGet Method for RootProjectFilesGet

Get metadata for files and directories in a project at the root directory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId Id of a project
 @return ApiRootProjectFilesGetRequest
*/
func (a *BetaAPIService) RootProjectFilesGet(ctx context.Context, projectId string) ApiRootProjectFilesGetRequest {
	return ApiRootProjectFilesGetRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
	}
}

// Execute executes the request
//  @return RootProjectFilesGet200Response
func (a *BetaAPIService) RootProjectFilesGetExecute(r ApiRootProjectFilesGetRequest) (*RootProjectFilesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RootProjectFilesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaAPIService.RootProjectFilesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/projects/{project_id}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.offset == nil {
		return localVarReturnValue, nil, reportError("offset is required and must be specified")
	}
	if *r.offset < 0 {
		return localVarReturnValue, nil, reportError("offset must be greater than 0")
	}
	if r.limit == nil {
		return localVarReturnValue, nil, reportError("limit is required and must be specified")
	}
	if *r.limit < 1 {
		return localVarReturnValue, nil, reportError("limit must be greater than 1")
	}
	if *r.limit > 100 {
		return localVarReturnValue, nil, reportError("limit must be less than 100")
	}

	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.branchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "branch_id", r.branchId, "")
	}
	if r.recursive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recursive", r.recursive, "")
	} else {
		var defaultValue bool = false
		r.recursive = &defaultValue
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
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
