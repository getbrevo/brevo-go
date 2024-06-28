# \EmailCampaignsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailCampaign**](EmailCampaignsApi.md#CreateEmailCampaign) | **Post** /emailCampaigns | Create an email campaign
[**DeleteEmailCampaign**](EmailCampaignsApi.md#DeleteEmailCampaign) | **Delete** /emailCampaigns/{campaignId} | Delete an email campaign
[**EmailExportRecipients**](EmailCampaignsApi.md#EmailExportRecipients) | **Post** /emailCampaigns/{campaignId}/exportRecipients | Export the recipients of an email campaign
[**GetAbTestCampaignResult**](EmailCampaignsApi.md#GetAbTestCampaignResult) | **Get** /emailCampaigns/{campaignId}/abTestCampaignResult | Get an A/B test email campaign results
[**GetEmailCampaign**](EmailCampaignsApi.md#GetEmailCampaign) | **Get** /emailCampaigns/{campaignId} | Get an email campaign report
[**GetEmailCampaigns**](EmailCampaignsApi.md#GetEmailCampaigns) | **Get** /emailCampaigns | Return all your created email campaigns
[**GetSharedTemplateUrl**](EmailCampaignsApi.md#GetSharedTemplateUrl) | **Get** /emailCampaigns/{campaignId}/sharedUrl | Get a shared template url
[**SendEmailCampaignNow**](EmailCampaignsApi.md#SendEmailCampaignNow) | **Post** /emailCampaigns/{campaignId}/sendNow | Send an email campaign immediately, based on campaignId
[**SendReport**](EmailCampaignsApi.md#SendReport) | **Post** /emailCampaigns/{campaignId}/sendReport | Send the report of a campaign
[**SendTestEmail**](EmailCampaignsApi.md#SendTestEmail) | **Post** /emailCampaigns/{campaignId}/sendTest | Send an email campaign to your test list
[**UpdateCampaignStatus**](EmailCampaignsApi.md#UpdateCampaignStatus) | **Put** /emailCampaigns/{campaignId}/status | Update an email campaign status
[**UpdateEmailCampaign**](EmailCampaignsApi.md#UpdateEmailCampaign) | **Put** /emailCampaigns/{campaignId} | Update an email campaign
[**UploadImageToGallery**](EmailCampaignsApi.md#UploadImageToGallery) | **Post** /emailCampaigns/images | Upload an image to your account&#39;s image gallery


# **CreateEmailCampaign**
> CreateModel CreateEmailCampaign(ctx, emailCampaigns)
Create an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **emailCampaigns** | [**CreateEmailCampaign**](CreateEmailCampaign.md)| Values to create a campaign | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEmailCampaign**
> DeleteEmailCampaign(ctx, campaignId)
Delete an email campaign

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

# **EmailExportRecipients**
> CreatedProcessId EmailExportRecipients(ctx, campaignId, optional)
Export the recipients of an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
 **optional** | ***EmailExportRecipientsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailExportRecipientsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recipientExport** | [**optional.Interface of EmailExportRecipients**](EmailExportRecipients.md)| Values to send for a recipient export request | 

### Return type

[**CreatedProcessId**](createdProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAbTestCampaignResult**
> AbTestCampaignResult GetAbTestCampaignResult(ctx, campaignId)
Get an A/B test email campaign results

Obtain winning version of an A/B test email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the A/B test campaign | 

### Return type

[**AbTestCampaignResult**](abTestCampaignResult.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaign**
> GetEmailCampaign GetEmailCampaign(ctx, campaignId, optional)
Get an email campaign report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
 **optional** | ***GetEmailCampaignOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEmailCampaignOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statistics** | **optional.String**| Filter on the type of statistics required. Example **globalStats** value will only fetch globalStats info of the campaign in returned response. | 

### Return type

[**GetEmailCampaign**](getEmailCampaign.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailCampaigns**
> GetEmailCampaigns GetEmailCampaigns(ctx, optional)
Return all your created email campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEmailCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Filter on the type of the campaigns | 
 **status** | **optional.String**| Filter on the status of the campaign | 
 **statistics** | **optional.String**| Filter on the type of statistics required. Example **globalStats** value will only fetch globalStats info of the campaign in returned response.This option only returns data for events occurred in the last 6 months.For older campaigns, it’s advisable to use the **Get Campaign Report** endpoint. | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent email campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#39;status&#39; not passed and if passed is set to &#39;sent&#39; ) | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent email campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#39;status&#39; not passed and if passed is set to &#39;sent&#39; ) | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **excludeHtmlContent** | **optional.Bool**| Use this flag to exclude htmlContent from the response body. If set to **true**, htmlContent field will be returned as empty string in the response body | 

### Return type

[**GetEmailCampaigns**](getEmailCampaigns.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharedTemplateUrl**
> GetSharedTemplateUrl GetSharedTemplateUrl(ctx, campaignId)
Get a shared template url

Get a unique URL to share & import an email template from one Brevo account to another.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign or template | 

### Return type

[**GetSharedTemplateUrl**](getSharedTemplateUrl.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendEmailCampaignNow**
> SendEmailCampaignNow(ctx, campaignId)
Send an email campaign immediately, based on campaignId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendReport**
> SendReport(ctx, campaignId, sendReport)
Send the report of a campaign

A PDF will be sent to the specified email addresses

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
  **sendReport** | [**SendReport**](SendReport.md)| Values for send a report | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestEmail**
> SendTestEmail(ctx, campaignId, emailTo)
Send an email campaign to your test list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
  **emailTo** | [**SendTestEmail**](SendTestEmail.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaignStatus**
> UpdateCampaignStatus(ctx, campaignId, status)
Update an email campaign status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
  **status** | [**UpdateCampaignStatus**](UpdateCampaignStatus.md)| Status of the campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmailCampaign**
> UpdateEmailCampaign(ctx, campaignId, emailCampaign)
Update an email campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the campaign | 
  **emailCampaign** | [**UpdateEmailCampaign**](UpdateEmailCampaign.md)| Values to update a campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadImageToGallery**
> UploadImageModel UploadImageToGallery(ctx, uploadImage)
Upload an image to your account's image gallery

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uploadImage** | [**UploadImageToGallery**](UploadImageToGallery.md)| Parameters to upload an image | 

### Return type

[**UploadImageModel**](uploadImageModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

