# {{classname}}

All URIs are relative to *https://api.twitter.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBatchComplianceJob**](ComplianceApi.md#CreateBatchComplianceJob) | **Post** /2/compliance/jobs | Create compliance job
[**GetBatchComplianceJob**](ComplianceApi.md#GetBatchComplianceJob) | **Get** /2/compliance/jobs/{id} | Get compliance job
[**ListBatchComplianceJobs**](ComplianceApi.md#ListBatchComplianceJobs) | **Get** /2/compliance/jobs | List compliance jobs

# **CreateBatchComplianceJob**
> SingleComplianceJobResponse CreateBatchComplianceJob(ctx, body)
Create compliance job

Creates a compliance for the given job type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ComplianceJobsBody**](ComplianceJobsBody.md)|  | 

### Return type

[**SingleComplianceJobResponse**](SingleComplianceJobResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBatchComplianceJob**
> SingleComplianceJobResponse GetBatchComplianceJob(ctx, id)
Get compliance job

Returns a single compliance job by ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| ID of the compliance job to retrieve. | 

### Return type

[**SingleComplianceJobResponse**](SingleComplianceJobResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBatchComplianceJobs**
> MultiComplianceJobResponse ListBatchComplianceJobs(ctx, type_, optional)
List compliance jobs

Returns recent compliance jobs for a given job type and optional job status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | [**ComplianceJobType**](.md)| Type of compliance job to list. | 
 **optional** | ***ComplianceApiListBatchComplianceJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComplianceApiListBatchComplianceJobsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of ComplianceJobStatus**](.md)| Status of compliance job to list. | 

### Return type

[**MultiComplianceJobResponse**](MultiComplianceJobResponse.md)

### Authorization

[BearerToken](../README.md#BearerToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

