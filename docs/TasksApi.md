# \TasksApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmTasksGet**](TasksApi.md#CrmTasksGet) | **Get** /crm/tasks | Get all tasks
[**CrmTasksIdDelete**](TasksApi.md#CrmTasksIdDelete) | **Delete** /crm/tasks/{id} | Delete a task
[**CrmTasksIdGet**](TasksApi.md#CrmTasksIdGet) | **Get** /crm/tasks/{id} | Get a task
[**CrmTasksIdPatch**](TasksApi.md#CrmTasksIdPatch) | **Patch** /crm/tasks/{id} | Update a task
[**CrmTasksPost**](TasksApi.md#CrmTasksPost) | **Post** /crm/tasks | Create a task
[**CrmTasktypesGet**](TasksApi.md#CrmTasktypesGet) | **Get** /crm/tasktypes | Get all task types


# **CrmTasksGet**
> TaskList CrmTasksGet(ctx, optional)
Get all tasks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmTasksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmTasksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterType** | **optional.String**| Filter by task type (ID) | 
 **filterStatus** | **optional.String**| Filter by task status | 
 **filterDate** | **optional.String**| Filter by date | 
 **filterAssignTo** | **optional.String**| Filter by assignTo id | 
 **filterContacts** | **optional.String**| Filter by contact ids | 
 **filterDeals** | **optional.String**| Filter by deals ids | 
 **filterCompanies** | **optional.String**| Filter by companies ids | 
 **dateFrom** | **optional.Int32**| dateFrom to date range filter type (timestamp in milliseconds) | 
 **dateTo** | **optional.Int32**| dateTo to date range filter type (timestamp in milliseconds) | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 
 **sortBy** | **optional.String**| The field used to sort field names. | 

### Return type

[**TaskList**](TaskList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdDelete**
> CrmTasksIdDelete(ctx, id)
Delete a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdGet**
> Task CrmTasksIdGet(ctx, id)
Get a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Task**](Task.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksIdPatch**
> CrmTasksIdPatch(ctx, id, body)
Update a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body7**](Body7.md)| Updated task details. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasksPost**
> InlineResponse2011 CrmTasksPost(ctx, body)
Create a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body6**](Body6.md)| Task name. | 

### Return type

[**InlineResponse2011**](inlineResponse2011.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmTasktypesGet**
> TaskTypes CrmTasktypesGet(ctx, )
Get all task types

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TaskTypes**](TaskTypes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

