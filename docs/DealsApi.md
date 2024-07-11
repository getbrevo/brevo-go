# \DealsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmAttributesDealsGet**](DealsApi.md#CrmAttributesDealsGet) | **Get** /crm/attributes/deals | Get deal attributes
[**CrmDealsGet**](DealsApi.md#CrmDealsGet) | **Get** /crm/deals | Get all deals
[**CrmDealsIdDelete**](DealsApi.md#CrmDealsIdDelete) | **Delete** /crm/deals/{id} | Delete a deal
[**CrmDealsIdGet**](DealsApi.md#CrmDealsIdGet) | **Get** /crm/deals/{id} | Get a deal
[**CrmDealsIdPatch**](DealsApi.md#CrmDealsIdPatch) | **Patch** /crm/deals/{id} | Update a deal
[**CrmDealsImportPost**](DealsApi.md#CrmDealsImportPost) | **Post** /crm/deals/import | Import deals(creation and updation)
[**CrmDealsLinkUnlinkIdPatch**](DealsApi.md#CrmDealsLinkUnlinkIdPatch) | **Patch** /crm/deals/link-unlink/{id} | Link and Unlink a deal with contacts and companies
[**CrmDealsPost**](DealsApi.md#CrmDealsPost) | **Post** /crm/deals | Create a deal
[**CrmPipelineDetailsAllGet**](DealsApi.md#CrmPipelineDetailsAllGet) | **Get** /crm/pipeline/details/all | Get all pipelines
[**CrmPipelineDetailsGet**](DealsApi.md#CrmPipelineDetailsGet) | **Get** /crm/pipeline/details | Get pipeline stages
[**CrmPipelineDetailsPipelineIDGet**](DealsApi.md#CrmPipelineDetailsPipelineIDGet) | **Get** /crm/pipeline/details/{pipelineID} | Get a pipeline


# **CrmAttributesDealsGet**
> DealAttributes CrmAttributesDealsGet(ctx, )
Get deal attributes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DealAttributes**](DealAttributes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsGet**
> DealsList CrmDealsGet(ctx, optional)
Get all deals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CrmDealsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CrmDealsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersAttributesDealName** | **optional.String**| Filter by attributes. If you have a filter for the owner on your end, please send it as filters[attributes.deal_owner] and utilize the account email for the filtering. | 
 **filtersLinkedCompaniesIds** | **optional.String**| Filter by linked companies ids | 
 **filtersLinkedContactsIds** | **optional.String**| Filter by linked companies ids | 
 **offset** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 
 **sortBy** | **optional.String**| The field used to sort field names. | 

### Return type

[**DealsList**](DealsList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsIdDelete**
> CrmDealsIdDelete(ctx, id)
Delete a deal

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

# **CrmDealsIdGet**
> Deal CrmDealsIdGet(ctx, id)
Get a deal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Deal**](Deal.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsIdPatch**
> CrmDealsIdPatch(ctx, id, body)
Update a deal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body9**](Body9.md)| Updated deal details. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsImportPost**
> InlineResponse2003 CrmDealsImportPost(ctx, file, mapping)
Import deals(creation and updation)

Import deals from a CSV file with mapping options.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **file** | ***os.File**| The CSV file to upload.The file should have the first row as the mapping attribute. Some default attribute names are (a) deal_id [brevo mongoID to update deals] (b) associated_contact (c) associated_company (f) any other attribute with internal name  | 
  **mapping** | **string**| The mapping options in JSON format.   json    {       \&quot;link_entities\&quot;: true, // Determines whether to link related entities during the import process       \&quot;unlink_entities\&quot;: false, //Determines whether to unlink related entities during the import process.       \&quot;update_existing_records\&quot;: true, // Determines whether to update based on deal ID or treat every row as create       \&quot;unset_empty_attributes\&quot;: false // Determines whether unset a specific attribute during update if values input is blank     }  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsLinkUnlinkIdPatch**
> CrmDealsLinkUnlinkIdPatch(ctx, id, body)
Link and Unlink a deal with contacts and companies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body10**](Body10.md)| Linked / Unlinked contacts and companies ids. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsPost**
> InlineResponse2011 CrmDealsPost(ctx, body)
Create a deal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body8**](Body8.md)| Deal create data. | 

### Return type

[**InlineResponse2011**](InlineResponse2011.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmPipelineDetailsAllGet**
> Pipelines CrmPipelineDetailsAllGet(ctx, )
Get all pipelines

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Pipelines**](Pipelines.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmPipelineDetailsGet**
> Pipeline CrmPipelineDetailsGet(ctx, )
Get pipeline stages

This endpoint is deprecated. Prefer /crm/pipeline/details/{pipelineID} instead.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmPipelineDetailsPipelineIDGet**
> Pipelines CrmPipelineDetailsPipelineIDGet(ctx, pipelineID)
Get a pipeline

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineID** | **string**|  | 

### Return type

[**Pipelines**](Pipelines.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

