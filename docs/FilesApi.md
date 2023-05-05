# \FilesApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmFilesGet**](FilesApi.md#CrmFilesGet) | **Get** /crm/files | Get all files
[**CrmFilesIdDataGet**](FilesApi.md#CrmFilesIdDataGet) | **Get** /crm/files/{id}/data | Get file details
[**CrmFilesIdDelete**](FilesApi.md#CrmFilesIdDelete) | **Delete** /crm/files/{id} | Delete a file
[**CrmFilesIdGet**](FilesApi.md#CrmFilesIdGet) | **Get** /crm/files/{id} | Download a file
[**CrmFilesPost**](FilesApi.md#CrmFilesPost) | **Post** /crm/files | Upload a file


# **CrmFilesGet**
> FileList CrmFilesGet(ctx, optional)
Get all files

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **optional.String**| Filter by file entity type | 
 **entityIds** | **optional.String**| Filter by file entity IDs | 
 **dateFrom** | **optional.Int32**| dateFrom to date range filter type (timestamp in milliseconds) | 
 **dateTo** | **optional.Int32**| dateTo to date range filter type (timestamp in milliseconds) | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 

### Return type

[**FileList**](FileList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmFilesIdDataGet**
> FileData CrmFilesIdDataGet(ctx, id)
Get file details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| File id to get file data. | 

### Return type

[**FileData**](FileData.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmFilesIdDelete**
> CrmFilesIdDelete(ctx, id)
Delete a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| File id to delete. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmFilesIdGet**
> FileDownloadableLink CrmFilesIdGet(ctx, id)
Download a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| File id to download. | 

### Return type

[**FileDownloadableLink**](FileDownloadableLink.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmFilesPost**
> FileData CrmFilesPost(ctx, file, optional)
Upload a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File**| File data to create a file. | 
 **optional** | ***CrmFilesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmFilesPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dealId** | **optional.String**| Deal id linked to a file | 
 **contactId** | **optional.Int64**| Contact id linked to a file | 
 **companyId** | **optional.String**| Company id linked to a file | 

### Return type

[**FileData**](FileData.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

