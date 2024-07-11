# \PaymentsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentRequest**](PaymentsApi.md#CreatePaymentRequest) | **Post** /payments/requests | Create a payment request
[**DeletePaymentRequest**](PaymentsApi.md#DeletePaymentRequest) | **Delete** /payments/requests/{id} | Delete a payment request.
[**GetPaymentRequest**](PaymentsApi.md#GetPaymentRequest) | **Get** /payments/requests/{id} | Get payment request details


# **CreatePaymentRequest**
> CreatePaymentResponse CreatePaymentRequest(ctx, createPaymentRquest)
Create a payment request

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createPaymentRquest** | [**CreatePaymentRequest**](CreatePaymentRequest.md)| Create a payment request  | 

### Return type

[**CreatePaymentResponse**](CreatePaymentResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentRequest**
> DeletePaymentRequest(ctx, id)
Delete a payment request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the payment request. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPaymentRequest**
> GetPaymentRequest GetPaymentRequest(ctx, id)
Get payment request details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the payment Request | 

### Return type

[**GetPaymentRequest**](GetPaymentRequest.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

