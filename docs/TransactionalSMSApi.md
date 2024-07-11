# \TransactionalSMSApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSmsEvents**](TransactionalSMSApi.md#GetSmsEvents) | **Get** /transactionalSMS/statistics/events | Get all your SMS activity (unaggregated events)
[**GetTransacAggregatedSmsReport**](TransactionalSMSApi.md#GetTransacAggregatedSmsReport) | **Get** /transactionalSMS/statistics/aggregatedReport | Get your SMS activity aggregated over a period of time
[**GetTransacSmsReport**](TransactionalSMSApi.md#GetTransacSmsReport) | **Get** /transactionalSMS/statistics/reports | Get your SMS activity aggregated per day
[**SendTransacSms**](TransactionalSMSApi.md#SendTransacSms) | **Post** /transactionalSMS/sms | Send SMS message to a mobile number


# **GetSmsEvents**
> GetSmsEventReport GetSmsEvents(ctx, optional)
Get all your SMS activity (unaggregated events)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSmsEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSmsEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | 
 **offset** | **optional.Int64**| Index of the first document of the page | [default to 0]
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | 
 **phoneNumber** | **optional.String**| Filter the report for a specific phone number | 
 **event** | **optional.String**| Filter the report for specific events | 
 **tags** | **optional.String**| Filter the report for specific tags passed as a serialized urlencoded array | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetSmsEventReport**](GetSmsEventReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacAggregatedSmsReport**
> GetTransacAggregatedSmsReport GetTransacAggregatedSmsReport(ctx, optional)
Get your SMS activity aggregated over a period of time

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransacAggregatedSmsReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTransacAggregatedSmsReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with startDate and endDate | 
 **tag** | **optional.String**| Filter on a tag | 

### Return type

[**GetTransacAggregatedSmsReport**](GetTransacAggregatedSmsReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacSmsReport**
> GetTransacSmsReport GetTransacSmsReport(ctx, optional)
Get your SMS activity aggregated per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransacSmsReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTransacSmsReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | 
 **tag** | **optional.String**| Filter on a tag | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetTransacSmsReport**](GetTransacSmsReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTransacSms**
> SendSms SendTransacSms(ctx, sendTransacSms)
Send SMS message to a mobile number

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendTransacSms** | [**SendTransacSms**](SendTransacSms.md)| Values to send a transactional SMS | 

### Return type

[**SendSms**](sendSms.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

