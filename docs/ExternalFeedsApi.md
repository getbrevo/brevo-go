# \ExternalFeedsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalFeed**](ExternalFeedsApi.md#CreateExternalFeed) | **Post** /feeds | Create an external feed
[**DeleteExternalFeed**](ExternalFeedsApi.md#DeleteExternalFeed) | **Delete** /feeds/{uuid} | Delete an external feed
[**GetAllExternalFeeds**](ExternalFeedsApi.md#GetAllExternalFeeds) | **Get** /feeds | Fetch all external feeds
[**GetExternalFeedByUUID**](ExternalFeedsApi.md#GetExternalFeedByUUID) | **Get** /feeds/{uuid} | Get an external feed by UUID
[**UpdateExternalFeed**](ExternalFeedsApi.md#UpdateExternalFeed) | **Put** /feeds/{uuid} | Update an external feed


# **CreateExternalFeed**
> InlineResponse2015 CreateExternalFeed(ctx, createExternalFeed)
Create an external feed

This endpoint will create an external feed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createExternalFeed** | [**CreateExternalFeed**](CreateExternalFeed.md)| Values to create a feed | 

### Return type

[**InlineResponse2015**](inline_response_201_5.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteExternalFeed**
> DeleteExternalFeed(ctx, uuid)
Delete an external feed

This endpoint will delete an external feed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| UUID of the feed to delete | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllExternalFeeds**
> GetAllExternalFeeds GetAllExternalFeeds(ctx, optional)
Fetch all external feeds

This endpoint can fetch all created external feeds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllExternalFeedsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllExternalFeedsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Can be used to filter records by search keyword on feed name | 
 **startDate** | **optional.String**| Mandatory if &#x60;endDate&#x60; is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Can be maximum 30 days older than current date. | 
 **endDate** | **optional.String**| Mandatory if &#x60;startDate&#x60; is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed. | [default to desc]
 **authType** | **optional.String**| Filter the records by &#x60;authType&#x60; of the feed. | 
 **limit** | **optional.Int64**| Number of documents returned per page. | [default to 50]
 **offset** | **optional.Int64**| Index of the first document on the page. | [default to 0]

### Return type

[**GetAllExternalFeeds**](getAllExternalFeeds.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExternalFeedByUUID**
> GetExternalFeedByUuid GetExternalFeedByUUID(ctx, uuid)
Get an external feed by UUID

This endpoint will update an external feed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| UUID of the feed to fetch | 

### Return type

[**GetExternalFeedByUuid**](getExternalFeedByUUID.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateExternalFeed**
> UpdateExternalFeed(ctx, uuid, updateExternalFeed)
Update an external feed

This endpoint will update an external feed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| UUID of the feed to update | 
  **updateExternalFeed** | [**UpdateExternalFeed**](UpdateExternalFeed.md)| Values to update a feed | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

