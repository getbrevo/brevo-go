# \DomainsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateDomain**](DomainsApi.md#AuthenticateDomain) | **Put** /senders/domains/{domainName}/authenticate | Authenticate a domain
[**CreateDomain**](DomainsApi.md#CreateDomain) | **Post** /senders/domains | Create a new domain
[**DeleteDomain**](DomainsApi.md#DeleteDomain) | **Delete** /senders/domains/{domainName} | Delete a domain
[**GetDomainonfiguration**](DomainsApi.md#GetDomainonfiguration) | **Get** /senders/domains/{domainName} | Validate domain configuration
[**GetDomains**](DomainsApi.md#GetDomains) | **Get** /senders/domains | Get the list of all your domains


# **AuthenticateDomain**
> AuthenticateDomainModel AuthenticateDomain(ctx, domainName)
Authenticate a domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| Domain name | 

### Return type

[**AuthenticateDomainModel**](AuthenticateDomainModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDomain**
> CreateDomainModel CreateDomain(ctx, optional)
Create a new domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateDomainOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateDomainOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainName** | [**optional.Interface of CreateDomain**](CreateDomain.md)| domain&#39;s name | 

### Return type

[**CreateDomainModel**](CreateDomainModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDomain**
> DeleteDomain(ctx, domainName)
Delete a domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| Domain name | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomainonfiguration**
> GetDomainonfigurationModel GetDomainonfiguration(ctx, domainName)
Validate domain configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domainName** | **string**| Domain name | 

### Return type

[**GetDomainonfigurationModel**](GetDomainonfigurationModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDomains**
> GetDomainsList GetDomains(ctx, )
Get the list of all your domains

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetDomainsList**](GetDomainsList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

