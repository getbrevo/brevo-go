# \AccountApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /account | Get your account information, plan and credits details
[**GetAccountActivity**](AccountApi.md#GetAccountActivity) | **Get** /organization/activities | Get user activity logs


# **GetAccount**
> GetAccount GetAccount(ctx, )
Get your account information, plan and credits details

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAccount**](GetAccount.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAccountActivity**
> GetAccountActivity GetAccountActivity(ctx, optional)
Get user activity logs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAccountActivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAccountActivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Enter start date in UTC date (YYYY-MM-DD) format to filter the activity in your account. Maximum time period that can be selected is one month. Additionally, you can retrieve activity logs from the past 12 months from the date of your search. | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Enter end date in UTC date (YYYY-MM-DD) format to filter the activity in your account. Maximum time period that can be selected is one month. | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document in the page. | [default to 0]

### Return type

[**GetAccountActivity**](GetAccountActivity.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

