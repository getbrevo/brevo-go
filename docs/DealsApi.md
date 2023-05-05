# \DealsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CrmAttributesDealsGet**](DealsApi.md#CrmAttributesDealsGet) | **Get** /crm/attributes/deals | Get deal attributes
[**CrmDealsGet**](DealsApi.md#CrmDealsGet) | **Get** /crm/deals | Get all deals
[**CrmDealsIdDelete**](DealsApi.md#CrmDealsIdDelete) | **Delete** /crm/deals/{id} | Delete a deal
[**CrmDealsIdGet**](DealsApi.md#CrmDealsIdGet) | **Get** /crm/deals/{id} | Get a deal
[**CrmDealsIdPatch**](DealsApi.md#CrmDealsIdPatch) | **Patch** /crm/deals/{id} | Update a deal
[**CrmDealsLinkUnlinkIdPatch**](DealsApi.md#CrmDealsLinkUnlinkIdPatch) | **Patch** /crm/deals/link-unlink/{id} | Link and Unlink a deal with contacts and companies
[**CrmDealsPost**](DealsApi.md#CrmDealsPost) | **Post** /crm/deals | Create a deal
[**CrmPipelineDetailsGet**](DealsApi.md#CrmPipelineDetailsGet) | **Get** /crm/pipeline/details | Get pipeline stages


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
 **filtersAttributes** | **optional.String**| Filter by attrbutes. If you have filter for owner on your side please send it as &#x60;attributes.owner&#x60;.\&quot; | 
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
  **body** | [**Body4**](Body4.md)| Updated deal details. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
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
  **body** | [**Body5**](Body5.md)| Linked / Unlinked contacts and companies ids. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmDealsPost**
> InlineResponse201 CrmDealsPost(ctx, body)
Create a deal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body3**](Body3.md)| Deal create data. | 

### Return type

[**InlineResponse201**](inlineResponse201.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CrmPipelineDetailsGet**
> Pipeline CrmPipelineDetailsGet(ctx, )
Get pipeline stages

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

