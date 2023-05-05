# \NotesApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmNotesGet**](NotesApi.md#CrmNotesGet) | **Get** /crm/notes | Get all notes
[**CrmNotesIdDelete**](NotesApi.md#CrmNotesIdDelete) | **Delete** /crm/notes/{id} | Delete a note
[**CrmNotesIdGet**](NotesApi.md#CrmNotesIdGet) | **Get** /crm/notes/{id} | Get a note
[**CrmNotesIdPatch**](NotesApi.md#CrmNotesIdPatch) | **Patch** /crm/notes/{id} | Update a note
[**CrmNotesPost**](NotesApi.md#CrmNotesPost) | **Post** /crm/notes | Create a note


# **CrmNotesGet**
> NoteList CrmNotesGet(ctx, optional)
Get all notes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmNotesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmNotesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entity** | **optional.String**| Filter by note entity type | 
 **entityIds** | **optional.String**| Filter by note entity IDs | 
 **dateFrom** | **optional.Int32**| dateFrom to date range filter type (timestamp in milliseconds) | 
 **dateTo** | **optional.Int32**| dateTo to date range filter type (timestamp in milliseconds) | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 

### Return type

[**NoteList**](NoteList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdDelete**
> CrmNotesIdDelete(ctx, id)
Delete a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to delete | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdGet**
> Note CrmNotesIdGet(ctx, id)
Get a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to get | 

### Return type

[**Note**](Note.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesIdPatch**
> CrmNotesIdPatch(ctx, id, body)
Update a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Note ID to update | 
  **body** | [**NoteData**](NoteData.md)| Note data to update a note | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmNotesPost**
> NoteId CrmNotesPost(ctx, body)
Create a note

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NoteData**](NoteData.md)| Note data to create a note. | 

### Return type

[**NoteId**](NoteId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

