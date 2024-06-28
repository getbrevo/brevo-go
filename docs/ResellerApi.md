# \ResellerApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredits**](ResellerApi.md#AddCredits) | **Post** /reseller/children/{childIdentifier}/credits/add | Add Email and/or SMS credits to a specific child account
[**AssociateIpToChild**](ResellerApi.md#AssociateIpToChild) | **Post** /reseller/children/{childIdentifier}/ips/associate | Associate a dedicated IP to the child
[**CreateChildDomain**](ResellerApi.md#CreateChildDomain) | **Post** /reseller/children/{childIdentifier}/domains | Create a domain for a child account
[**CreateResellerChild**](ResellerApi.md#CreateResellerChild) | **Post** /reseller/children | Creates a reseller child
[**DeleteChildDomain**](ResellerApi.md#DeleteChildDomain) | **Delete** /reseller/children/{childIdentifier}/domains/{domainName} | Delete the sender domain of the reseller child based on the childIdentifier and domainName passed
[**DeleteResellerChild**](ResellerApi.md#DeleteResellerChild) | **Delete** /reseller/children/{childIdentifier} | Delete a single reseller child based on the child identifier supplied
[**DissociateIpFromChild**](ResellerApi.md#DissociateIpFromChild) | **Post** /reseller/children/{childIdentifier}/ips/dissociate | Dissociate a dedicated IP to the child
[**GetChildAccountCreationStatus**](ResellerApi.md#GetChildAccountCreationStatus) | **Get** /reseller/children/{childIdentifier}/accountCreationStatus | Get the status of a reseller&#39;s child account creation, whether it is successfully created (exists) or not based on the identifier supplied
[**GetChildDomains**](ResellerApi.md#GetChildDomains) | **Get** /reseller/children/{childIdentifier}/domains | Get all sender domains for a specific child account
[**GetChildInfo**](ResellerApi.md#GetChildInfo) | **Get** /reseller/children/{childIdentifier} | Get a child account&#39;s details
[**GetResellerChilds**](ResellerApi.md#GetResellerChilds) | **Get** /reseller/children | Get the list of all children accounts
[**GetSsoToken**](ResellerApi.md#GetSsoToken) | **Get** /reseller/children/{childIdentifier}/auth | Get session token to access Brevo (SSO)
[**RemoveCredits**](ResellerApi.md#RemoveCredits) | **Post** /reseller/children/{childIdentifier}/credits/remove | Remove Email and/or SMS credits from a specific child account
[**UpdateChildAccountStatus**](ResellerApi.md#UpdateChildAccountStatus) | **Put** /reseller/children/{childIdentifier}/accountStatus | Update info of reseller&#39;s child account status based on the childIdentifier supplied
[**UpdateChildDomain**](ResellerApi.md#UpdateChildDomain) | **Put** /reseller/children/{childIdentifier}/domains/{domainName} | Update the sender domain of reseller&#39;s child based on the childIdentifier and domainName passed
[**UpdateResellerChild**](ResellerApi.md#UpdateResellerChild) | **Put** /reseller/children/{childIdentifier} | Update info of reseller&#39;s child based on the child identifier supplied


# **AddCredits**
> RemainingCreditModel AddCredits(ctx, childIdentifier, addCredits)
Add Email and/or SMS credits to a specific child account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **addCredits** | [**AddCredits**](AddCredits.md)| Values to post to add credit to a specific child account | 

### Return type

[**RemainingCreditModel**](remainingCreditModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssociateIpToChild**
> AssociateIpToChild(ctx, childIdentifier, ip)
Associate a dedicated IP to the child

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **ip** | [**ManageIp**](ManageIp.md)| IP to associate | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateChildDomain**
> CreateChildDomain(ctx, childIdentifier, addChildDomain)
Create a domain for a child account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **addChildDomain** | [**AddChildDomain**](AddChildDomain.md)| Sender domain to add for a specific child account. This will not be displayed to the parent account. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateResellerChild**
> CreateReseller CreateResellerChild(ctx, optional)
Creates a reseller child

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CreateResellerChildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreateResellerChildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resellerChild** | [**optional.Interface of CreateChild**](CreateChild.md)| reseller child to add | 

### Return type

[**CreateReseller**](createReseller.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChildDomain**
> DeleteChildDomain(ctx, childIdentifier, domainName)
Delete the sender domain of the reseller child based on the childIdentifier and domainName passed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **domainName** | **string**| Pass the existing domain that needs to be deleted | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteResellerChild**
> DeleteResellerChild(ctx, childIdentifier)
Delete a single reseller child based on the child identifier supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or child id of reseller&#39;s child | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DissociateIpFromChild**
> DissociateIpFromChild(ctx, childIdentifier, ip)
Dissociate a dedicated IP to the child

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **ip** | [**ManageIp**](ManageIp.md)| IP to dissociate | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildAccountCreationStatus**
> GetChildAccountCreationStatus GetChildAccountCreationStatus(ctx, childIdentifier)
Get the status of a reseller's child account creation, whether it is successfully created (exists) or not based on the identifier supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 

### Return type

[**GetChildAccountCreationStatus**](getChildAccountCreationStatus.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildDomains**
> GetChildDomains GetChildDomains(ctx, childIdentifier)
Get all sender domains for a specific child account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 

### Return type

[**GetChildDomains**](getChildDomains.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChildInfo**
> GetChildInfo GetChildInfo(ctx, childIdentifier)
Get a child account's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 

### Return type

[**GetChildInfo**](getChildInfo.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetResellerChilds**
> GetChildrenList GetResellerChilds(ctx, optional)
Get the list of all children accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetResellerChildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetResellerChildsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents for child accounts information per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]

### Return type

[**GetChildrenList**](getChildrenList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSsoToken**
> GetSsoToken GetSsoToken(ctx, childIdentifier)
Get session token to access Brevo (SSO)

It returns a session [token] which will remain valid for a short period of time. A child account will be able to access a white-labeled section by using the following url pattern => https:/email.mydomain.com/login/sso?token=[token]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 

### Return type

[**GetSsoToken**](getSsoToken.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveCredits**
> RemainingCreditModel RemoveCredits(ctx, childIdentifier, removeCredits)
Remove Email and/or SMS credits from a specific child account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **removeCredits** | [**RemoveCredits**](RemoveCredits.md)| Values to post to remove email or SMS credits from a specific child account | 

### Return type

[**RemainingCreditModel**](remainingCreditModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChildAccountStatus**
> UpdateChildAccountStatus(ctx, childIdentifier, updateChildAccountStatus)
Update info of reseller's child account status based on the childIdentifier supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **updateChildAccountStatus** | [**UpdateChildAccountStatus**](UpdateChildAccountStatus.md)| values to update in child account status | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChildDomain**
> UpdateChildDomain(ctx, childIdentifier, domainName, updateChildDomain)
Update the sender domain of reseller's child based on the childIdentifier and domainName passed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **domainName** | **string**| Pass the existing domain that needs to be updated | 
  **updateChildDomain** | [**UpdateChildDomain**](UpdateChildDomain.md)| value to update for sender domain | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateResellerChild**
> UpdateResellerChild(ctx, childIdentifier, resellerChild)
Update info of reseller's child based on the child identifier supplied

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **childIdentifier** | **string**| Either auth key or id of reseller&#39;s child | 
  **resellerChild** | [**UpdateChild**](UpdateChild.md)| values to update in child profile | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

