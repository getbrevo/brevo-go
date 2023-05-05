# \CompaniesApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompaniesAttributesGet**](CompaniesApi.md#CompaniesAttributesGet) | **Get** /companies/attributes | Get company attributes
[**CompaniesGet**](CompaniesApi.md#CompaniesGet) | **Get** /companies | Get all companies
[**CompaniesIdDelete**](CompaniesApi.md#CompaniesIdDelete) | **Delete** /companies/{id} | Delete a company
[**CompaniesIdGet**](CompaniesApi.md#CompaniesIdGet) | **Get** /companies/{id} | Get a company
[**CompaniesIdPatch**](CompaniesApi.md#CompaniesIdPatch) | **Patch** /companies/{id} | Update a company
[**CompaniesLinkUnlinkIdPatch**](CompaniesApi.md#CompaniesLinkUnlinkIdPatch) | **Patch** /companies/link-unlink/{id} | Link and Unlink company with contacts and deals
[**CompaniesPost**](CompaniesApi.md#CompaniesPost) | **Post** /companies | Create a company


# **CompaniesAttributesGet**
> CompanyAttributes CompaniesAttributesGet(ctx, )
Get company attributes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CompanyAttributes**](CompanyAttributes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompaniesGet**
> CompaniesList CompaniesGet(ctx, optional)
Get all companies

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CompaniesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CompaniesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filter by attrbutes. If you have filter for owner on your side please send it as {\&quot;attributes.owner\&quot;:\&quot;5b1a17d914b73d35a76ca0c7\&quot;} | 
 **linkedContactsIds** | **optional.Int64**| Filter by linked contacts ids | 
 **linkedDealsIds** | **optional.String**| Filter by linked deals ids | 
 **page** | **optional.Int64**| Index of the first document of the page | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **sort** | **optional.String**| Sort the results in the ascending/descending order. Default order is **descending** by creation if &#x60;sort&#x60; is not passed | 
 **sortBy** | **optional.String**| The field used to sort field names. | 

### Return type

[**CompaniesList**](CompaniesList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompaniesIdDelete**
> CompaniesIdDelete(ctx, id)
Delete a company

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

# **CompaniesIdGet**
> Company CompaniesIdGet(ctx, id)
Get a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**Company**](Company.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompaniesIdPatch**
> Company CompaniesIdPatch(ctx, id, body)
Update a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body1**](Body1.md)| Updated company details. | 

### Return type

[**Company**](Company.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompaniesLinkUnlinkIdPatch**
> CompaniesLinkUnlinkIdPatch(ctx, id, body)
Link and Unlink company with contacts and deals

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
  **body** | [**Body2**](Body2.md)| Linked / Unlinked contacts and deals ids. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompaniesPost**
> InlineResponse200 CompaniesPost(ctx, body)
Create a company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| Company create data. | 

### Return type

[**InlineResponse200**](inlineResponse200.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

