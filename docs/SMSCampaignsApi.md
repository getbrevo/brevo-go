# \SMSCampaignsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSmsCampaign**](SMSCampaignsApi.md#CreateSmsCampaign) | **Post** /smsCampaigns | Creates an SMS campaign
[**DeleteSmsCampaign**](SMSCampaignsApi.md#DeleteSmsCampaign) | **Delete** /smsCampaigns/{campaignId} | Delete an SMS campaign
[**GetSmsCampaign**](SMSCampaignsApi.md#GetSmsCampaign) | **Get** /smsCampaigns/{campaignId} | Get an SMS campaign
[**GetSmsCampaigns**](SMSCampaignsApi.md#GetSmsCampaigns) | **Get** /smsCampaigns | Returns the information for all your created SMS campaigns
[**RequestSmsRecipientExport**](SMSCampaignsApi.md#RequestSmsRecipientExport) | **Post** /smsCampaigns/{campaignId}/exportRecipients | Export an SMS campaign&#39;s recipients
[**SendSmsCampaignNow**](SMSCampaignsApi.md#SendSmsCampaignNow) | **Post** /smsCampaigns/{campaignId}/sendNow | Send your SMS campaign immediately
[**SendSmsReport**](SMSCampaignsApi.md#SendSmsReport) | **Post** /smsCampaigns/{campaignId}/sendReport | Send an SMS campaign&#39;s report
[**SendTestSms**](SMSCampaignsApi.md#SendTestSms) | **Post** /smsCampaigns/{campaignId}/sendTest | Send a test SMS campaign
[**UpdateSmsCampaign**](SMSCampaignsApi.md#UpdateSmsCampaign) | **Put** /smsCampaigns/{campaignId} | Update an SMS campaign
[**UpdateSmsCampaignStatus**](SMSCampaignsApi.md#UpdateSmsCampaignStatus) | **Put** /smsCampaigns/{campaignId}/status | Update a campaign&#39;s status


# **CreateSmsCampaign**
> CreateModel CreateSmsCampaign(ctx, createSmsCampaign)
Creates an SMS campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createSmsCampaign** | [**CreateSmsCampaign**](CreateSmsCampaign.md)| Values to create an SMS Campaign | 

### Return type

[**CreateModel**](CreateModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmsCampaign**
> DeleteSmsCampaign(ctx, campaignId)
Delete an SMS campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the SMS campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmsCampaign**
> GetSmsCampaign GetSmsCampaign(ctx, campaignId)
Get an SMS campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the SMS campaign | 

### Return type

[**GetSmsCampaign**](getSmsCampaign.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmsCampaigns**
> GetSmsCampaigns GetSmsCampaigns(ctx, optional)
Returns the information for all your created SMS campaigns

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSmsCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSmsCampaignsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Status of campaign. | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent sms campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#39;status&#39; not passed and if passed is set to &#39;sent&#39; ) | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending (urlencoded) UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ) to filter the sent sms campaigns. Prefer to pass your timezone in date-time format for accurate result ( only available if either &#39;status&#39; not passed and if passed is set to &#39;sent&#39; ) | 
 **limit** | **optional.Int64**| Number limitation for the result returned | [default to 500]
 **offset** | **optional.Int64**| Beginning point in the list to retrieve from. | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetSmsCampaigns**](GetSmsCampaigns.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSmsRecipientExport**
> CreatedProcessId RequestSmsRecipientExport(ctx, campaignId, optional)
Export an SMS campaign's recipients

It returns the background process ID which on completion calls the notify URL that you have set in the input.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the campaign | 
 **optional** | ***RequestSmsRecipientExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestSmsRecipientExportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recipientExport** | [**optional.Interface of RequestSmsRecipientExport**](RequestSmsRecipientExport.md)| Values to send for a recipient export request | 

### Return type

[**CreatedProcessId**](CreatedProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendSmsCampaignNow**
> SendSmsCampaignNow(ctx, campaignId)
Send your SMS campaign immediately

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

# **SendSmsReport**
> SendSmsReport(ctx, campaignId, sendReport)
Send an SMS campaign's report

Send report of Sent and Archived campaign, to the specified email addresses, with respective data and a pdf attachment in detail.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the campaign | 
  **sendReport** | [**SendReport**](SendReport.md)| Values for send a report | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestSms**
> SendTestSms(ctx, campaignId, phoneNumber)
Send a test SMS campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| Id of the SMS campaign | 
  **phoneNumber** | [**SendTestSms**](SendTestSms.md)| Mobile number of the recipient with the country code. This number must belong to one of your contacts in Brevo account and must not be blacklisted | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmsCampaign**
> UpdateSmsCampaign(ctx, campaignId, updateSmsCampaign)
Update an SMS campaign

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the SMS campaign | 
  **updateSmsCampaign** | [**UpdateSmsCampaign**](UpdateSmsCampaign.md)| Values to update an SMS Campaign | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmsCampaignStatus**
> UpdateSmsCampaignStatus(ctx, campaignId, status)
Update a campaign's status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **campaignId** | **int64**| id of the campaign | 
  **status** | [**UpdateCampaignStatus**](UpdateCampaignStatus.md)| Status of the campaign. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

