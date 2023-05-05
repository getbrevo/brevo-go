# \WebhooksApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /webhooks | Create a webhook
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /webhooks/{webhookId} | Delete a webhook
[**ExportWebhooksHistory**](WebhooksApi.md#ExportWebhooksHistory) | **Post** /webhooks/export | Export all transactional events
[**GetWebhook**](WebhooksApi.md#GetWebhook) | **Get** /webhooks/{webhookId} | Get a webhook details
[**GetWebhooks**](WebhooksApi.md#GetWebhooks) | **Get** /webhooks | Get all webhooks
[**UpdateWebhook**](WebhooksApi.md#UpdateWebhook) | **Put** /webhooks/{webhookId} | Update a webhook


# **CreateWebhook**
> CreateModel CreateWebhook(ctx, createWebhook)
Create a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createWebhook** | [**CreateWebhook**](CreateWebhook.md)| Values to create a webhook | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhook**
> DeleteWebhook(ctx, webhookId)
Delete a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Id of the webhook | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportWebhooksHistory**
> CreatedProcessId ExportWebhooksHistory(ctx, exportWebhookHistory)
Export all transactional events

This endpoint will submit a request to get the history of webhooks in the CSV file. The link to download that CSV file will send to the webhook link provided in notifyURL of the body payload.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **exportWebhookHistory** | [**ExportWebhooksHistory**](ExportWebhooksHistory.md)| Values to submit for webhooks history | 

### Return type

[**CreatedProcessId**](createdProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhook**
> GetWebhook GetWebhook(ctx, webhookId)
Get a webhook details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Id of the webhook | 

### Return type

[**GetWebhook**](getWebhook.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWebhooks**
> GetWebhooks GetWebhooks(ctx, optional)
Get all webhooks

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWebhooksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWebhooksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Filter on webhook type | [default to transactional]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of webhook creation | [default to desc]

### Return type

[**GetWebhooks**](getWebhooks.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> UpdateWebhook(ctx, webhookId, updateWebhook)
Update a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookId** | **int64**| Id of the webhook | 
  **updateWebhook** | [**UpdateWebhook**](UpdateWebhook.md)| Values to update a webhook | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

