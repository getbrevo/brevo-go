# \TransactionalWhatsAppApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWhatsappEventReport**](TransactionalWhatsAppApi.md#GetWhatsappEventReport) | **Get** /whatsapp/statistics/events | Get all your WhatsApp activity (unaggregated events)
[**SendWhatsappMessage**](TransactionalWhatsAppApi.md#SendWhatsappMessage) | **Post** /whatsapp/sendMessage | Send a WhatsApp message


# **GetWhatsappEventReport**
> GetWhatsappEventReport GetWhatsappEventReport(ctx, optional)
Get all your WhatsApp activity (unaggregated events)

This endpoint will show the unaggregated statistics for WhatsApp activity (30 days by default if `startDate` and `endDate` or `days` is not passed. The date range can not exceed 90 days)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetWhatsappEventReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetWhatsappEventReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number limitation for the result returned | [default to 2500]
 **offset** | **optional.Int64**| Beginning point in the list to retrieve from | [default to 0]
 **startDate** | **optional.String**| **Mandatory if endDate is used.** Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate  | 
 **endDate** | **optional.String**| **Mandatory if startDate is used.** Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate  | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). _Not compatible with &#39;startDate&#39; and &#39;endDate&#39;_  | 
 **contactNumber** | **optional.String**| Filter results for specific contact (WhatsApp Number with country code. Example, 85264318721) | 
 **event** | **optional.String**| Filter the report for a specific event type | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetWhatsappEventReport**](GetWhatsappEventReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendWhatsappMessage**
> InlineResponse2014 SendWhatsappMessage(ctx, sendWhatsappMessage)
Send a WhatsApp message

This endpoint is used to send a WhatsApp message. <br/>(**The first message you send using the API must contain a Template ID. You must create a template on WhatsApp on the Brevo platform to fetch the Template ID.**)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendWhatsappMessage** | [**SendWhatsappMessage**](SendWhatsappMessage.md)| Values to send WhatsApp message | 

### Return type

[**InlineResponse2014**](inline_response_201_4.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

