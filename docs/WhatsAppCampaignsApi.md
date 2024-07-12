# \WhatsAppCampaignsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWhatsAppCampaign**](WhatsAppCampaignsApi.md#CreateWhatsAppCampaign) | **Post** /whatsappCampaigns | Create and Send a WhatsApp campaign
[**CreateWhatsAppTemplate**](WhatsAppCampaignsApi.md#CreateWhatsAppTemplate) | **Post** /whatsppCampaigns/template | Create a WhatsApp template
[**DeleteWhatsAppCampaign**](WhatsAppCampaignsApi.md#DeleteWhatsAppCampaign) | **Delete** /whatsappCampaigns/{campaignId} | Delete a WhatsApp campaign
[**GetWhatsAppCampaign**](WhatsAppCampaignsApi.md#GetWhatsAppCampaign) | **Get** /whatsappCampaigns/{campaignId} | Get a WhatsApp campaign
[**GetWhatsAppCampaigns**](WhatsAppCampaignsApi.md#GetWhatsAppCampaigns) | **Get** /whatsappCampaigns | Return all your created WhatsApp campaigns
[**GetWhatsAppConfig**](WhatsAppCampaignsApi.md#GetWhatsAppConfig) | **Get** /whatsappCampaigns/config | Get your WhatsApp API account information
[**GetWhatsAppTemplates**](WhatsAppCampaignsApi.md#GetWhatsAppTemplates) | **Get** /whatsappCampaigns/template-list | Return all your created WhatsApp templates
[**SendWhatsAppTemplateApproval**](WhatsAppCampaignsApi.md#SendWhatsAppTemplateApproval) | **Post** /whatsappCampaigns/template/approval/{templateId} | Send your WhatsApp template for approval
[**UpdateWhatsAppCampaign**](WhatsAppCampaignsApi.md#UpdateWhatsAppCampaign) | **Put** /whatsappCampaigns/{campaignId} | Update a WhatsApp campaign


# **CreateWhatsAppCampaign**
> CreateModel CreateWhatsAppCampaign(ctx, whatsAppCampaigns)
Create and Send a WhatsApp campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whatsAppCampaigns** | [**CreateWhatsAppCampaign**](CreateWhatsAppCampaign.md)| Values to create a campaign | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWhatsAppTemplate**
> CreateModel CreateWhatsAppTemplate(ctx, whatsAppTemplates)
Create a WhatsApp template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **whatsAppTemplates** | [**CreateWhatsAppTemplate**](CreateWhatsAppTemplate.md)| Values to create a template | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWhatsAppCampaign**
> DeleteWhatsAppCampaign(ctx, campaignId)
Delete a WhatsApp campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhatsAppCampaign**
> GetWhatsappCampaignOverview GetWhatsAppCampaign(ctx, campaignId)
Get a WhatsApp campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 

### Return type

[**GetWhatsappCampaignOverview**](GetWhatsappCampaignOverview.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhatsAppCampaigns**
> GetWhatsappCampaigns GetWhatsAppCampaigns(ctx, optional)
Return all your created WhatsApp campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWhatsAppCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWhatsAppCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| **Mandatory if endDate is used**. Starting (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the campaigns created. **Prefer to pass your timezone in date-time format for accurate result**  | 
 **endDate** | **optional.String**| **Mandatory if startDate is used**. Ending (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the campaigns created. **Prefer to pass your timezone in date-time format for accurate result**  | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record modification. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetWhatsappCampaigns**](GetWhatsappCampaigns.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhatsAppConfig**
> GetWhatsAppConfig GetWhatsAppConfig(ctx, )
Get your WhatsApp API account information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetWhatsAppConfig**](GetWhatsAppConfig.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWhatsAppTemplates**
> GetWaTemplates GetWhatsAppTemplates(ctx, optional)
Return all your created WhatsApp templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWhatsAppTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWhatsAppTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| **Mandatory if endDate is used**. Starting (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the templates created. **Prefer to pass your timezone in date-time format for accurate result**  | 
 **endDate** | **optional.String**| **Mandatory if startDate is used**. Ending (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the templates created. **Prefer to pass your timezone in date-time format for accurate result**  | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record modification. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **source** | **optional.String**| source of the template | 

### Return type

[**GetWaTemplates**](GetWATemplates.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendWhatsAppTemplateApproval**
> SendWhatsAppTemplateApproval(ctx, templateId)
Send your WhatsApp template for approval

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWhatsAppCampaign**
> UpdateWhatsAppCampaign(ctx, campaignId, optional)
Update a WhatsApp campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
 **optional** | ***UpdateWhatsAppCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateWhatsAppCampaignOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatsAppCampaign** | [**optional.Interface of UpdateWhatsAppCampaign**](UpdateWhatsAppCampaign.md)| values to update WhatsApp Campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

