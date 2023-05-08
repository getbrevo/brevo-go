# \MasterAccountApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorporateMasterAccountGet**](MasterAccountApi.md#CorporateMasterAccountGet) | **Get** /corporate/masterAccount | Get the details of requested master account
[**CorporateSubAccountGet**](MasterAccountApi.md#CorporateSubAccountGet) | **Get** /corporate/subAccount | Get the list of all the sub-accounts of the master account.
[**CorporateSubAccountIdDelete**](MasterAccountApi.md#CorporateSubAccountIdDelete) | **Delete** /corporate/subAccount/{id} | Delete a sub-account
[**CorporateSubAccountIdGet**](MasterAccountApi.md#CorporateSubAccountIdGet) | **Get** /corporate/subAccount/{id} | Get sub-account details
[**CorporateSubAccountIdPlanPut**](MasterAccountApi.md#CorporateSubAccountIdPlanPut) | **Put** /corporate/subAccount/{id}/plan | Update sub-account plan
[**CorporateSubAccountKeyPost**](MasterAccountApi.md#CorporateSubAccountKeyPost) | **Post** /corporate/subAccount/key | Create an API key for a sub-account
[**CorporateSubAccountPost**](MasterAccountApi.md#CorporateSubAccountPost) | **Post** /corporate/subAccount | Create a new sub-account under a master account.
[**CorporateSubAccountSsoTokenPost**](MasterAccountApi.md#CorporateSubAccountSsoTokenPost) | **Post** /corporate/subAccount/ssoToken | Generate SSO token to access Brevo


# **CorporateMasterAccountGet**
> MasterDetailsResponse CorporateMasterAccountGet(ctx, )
Get the details of requested master account

This endpoint will provide the details of the master account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MasterDetailsResponse**](MasterDetailsResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountGet**
> SubAccountsResponse CorporateSubAccountGet(ctx, offset, limit)
Get the list of all the sub-accounts of the master account.

This endpoint will provide the list all the sub-accounts of the master account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **offset** | **int32**| Index of the first sub-account in the page | 
  **limit** | **int32**| Number of sub-accounts to be displayed on each page | 

### Return type

[**SubAccountsResponse**](SubAccountsResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountIdDelete**
> CorporateSubAccountIdDelete(ctx, id)
Delete a sub-account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the sub-account organization to be deleted | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountIdGet**
> SubAccountDetailsResponse CorporateSubAccountIdGet(ctx, id)
Get sub-account details

This endpoint will provide the details for the specified sub-account company

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the sub-account organization | 

### Return type

[**SubAccountDetailsResponse**](SubAccountDetailsResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountIdPlanPut**
> CorporateSubAccountIdPlanPut(ctx, id, updatePlanDetails)
Update sub-account plan

This endpoint will update the sub-account plan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the sub-account organization | 
  **updatePlanDetails** | [**SubAccountUpdatePlanRequest**](SubAccountUpdatePlanRequest.md)| Values to update a sub-account plan | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountKeyPost**
> CreateApiKeyResponse CorporateSubAccountKeyPost(ctx, createApiKeyRequest)
Create an API key for a sub-account

This endpoint will generate an API v3 key for a sub account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md)| Values to generate API key for sub-account | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountPost**
> CreateSubAccountResponse CorporateSubAccountPost(ctx, subAccountCreate)
Create a new sub-account under a master account.

This endpoint will create a new sub-account under a master account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subAccountCreate** | [**CreateSubAccount**](CreateSubAccount.md)| values to create new sub-account | 

### Return type

[**CreateSubAccountResponse**](CreateSubAccountResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountSsoTokenPost**
> GetSsoToken CorporateSubAccountSsoTokenPost(ctx, ssoTokenRequest)
Generate SSO token to access Brevo

This endpoint generates an sso token to authenticate and access a sub-account of the master using the account endpoint https://account-app.brevo.com/account/login/sub-account/sso/[token], where [token] will be replaced by the actual token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ssoTokenRequest** | [**SsoTokenRequest**](SsoTokenRequest.md)| Values to generate SSO token for sub-account | 

### Return type

[**GetSsoToken**](GetSsoToken.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

