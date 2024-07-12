# \MasterAccountApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorporateGroupIdDelete**](MasterAccountApi.md#CorporateGroupIdDelete) | **Delete** /corporate/group/{id} | Delete a group
[**CorporateGroupIdGet**](MasterAccountApi.md#CorporateGroupIdGet) | **Get** /corporate/group/{id} | GET a group details
[**CorporateGroupIdPut**](MasterAccountApi.md#CorporateGroupIdPut) | **Put** /corporate/group/{id} | Update a group of sub-accounts
[**CorporateGroupPost**](MasterAccountApi.md#CorporateGroupPost) | **Post** /corporate/group | Create a new group of sub-accounts
[**CorporateGroupUnlinkGroupIdSubAccountsPut**](MasterAccountApi.md#CorporateGroupUnlinkGroupIdSubAccountsPut) | **Put** /corporate/group/unlink/{groupId}/subAccounts | Delete sub-account from group
[**CorporateMasterAccountGet**](MasterAccountApi.md#CorporateMasterAccountGet) | **Get** /corporate/masterAccount | Get the details of requested master account
[**CorporateSsoTokenPost**](MasterAccountApi.md#CorporateSsoTokenPost) | **Post** /corporate/ssoToken | Generate SSO token to access admin account
[**CorporateSubAccountGet**](MasterAccountApi.md#CorporateSubAccountGet) | **Get** /corporate/subAccount | Get the list of all the sub-accounts of the master account.
[**CorporateSubAccountIdApplicationsTogglePut**](MasterAccountApi.md#CorporateSubAccountIdApplicationsTogglePut) | **Put** /corporate/subAccount/{id}/applications/toggle | Enable/disable sub-account application(s)
[**CorporateSubAccountIdDelete**](MasterAccountApi.md#CorporateSubAccountIdDelete) | **Delete** /corporate/subAccount/{id} | Delete a sub-account
[**CorporateSubAccountIdGet**](MasterAccountApi.md#CorporateSubAccountIdGet) | **Get** /corporate/subAccount/{id} | Get sub-account details
[**CorporateSubAccountIdPlanPut**](MasterAccountApi.md#CorporateSubAccountIdPlanPut) | **Put** /corporate/subAccount/{id}/plan | Update sub-account plan
[**CorporateSubAccountIpAssociatePost**](MasterAccountApi.md#CorporateSubAccountIpAssociatePost) | **Post** /corporate/subAccount/ip/associate | Associate an IP to sub-accounts
[**CorporateSubAccountIpDissociatePut**](MasterAccountApi.md#CorporateSubAccountIpDissociatePut) | **Put** /corporate/subAccount/ip/dissociate | Dissociate an IP from sub-accounts
[**CorporateSubAccountKeyPost**](MasterAccountApi.md#CorporateSubAccountKeyPost) | **Post** /corporate/subAccount/key | Create an API key for a sub-account
[**CorporateSubAccountPost**](MasterAccountApi.md#CorporateSubAccountPost) | **Post** /corporate/subAccount | Create a new sub-account under a master account.
[**CorporateSubAccountSsoTokenPost**](MasterAccountApi.md#CorporateSubAccountSsoTokenPost) | **Post** /corporate/subAccount/ssoToken | Generate SSO token to access sub-account
[**CorporateUserInvitationActionEmailPut**](MasterAccountApi.md#CorporateUserInvitationActionEmailPut) | **Put** /corporate/user/invitation/{action}/{email} | Resend / cancel admin user invitation
[**CorporateUserRevokeEmailDelete**](MasterAccountApi.md#CorporateUserRevokeEmailDelete) | **Delete** /corporate/user/revoke/{email} | Revoke an admin user
[**GetAccountActivity**](MasterAccountApi.md#GetAccountActivity) | **Get** /organization/activities | Get user activity logs
[**GetCorporateInvitedUsersList**](MasterAccountApi.md#GetCorporateInvitedUsersList) | **Get** /corporate/invited/users | Get the list of all admin users
[**GetCorporateUserPermission**](MasterAccountApi.md#GetCorporateUserPermission) | **Get** /corporate/user/{email}/permissions | Check admin user permissions
[**GetSubAccountGroups**](MasterAccountApi.md#GetSubAccountGroups) | **Get** /corporate/groups | Get the list of groups
[**InviteAdminUser**](MasterAccountApi.md#InviteAdminUser) | **Post** /corporate/user/invitation/send | Send invitation to an admin user


# **CorporateGroupIdDelete**
> CorporateGroupIdDelete(ctx, id)
Delete a group

This endpoint allows you to delete a group of sub-organizations. When a group is deleted, the sub-organizations are no longer part of this group. The users associated with the group are no longer associated with the group once deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the group | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateGroupIdGet**
> CorporateGroupDetailsResponse CorporateGroupIdGet(ctx, id)
GET a group details

This endpoint allows you to retrieve a specific group’s information such as the list of sub-organizations and the user associated with the group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the group of sub-organization | 

### Return type

[**CorporateGroupDetailsResponse**](CorporateGroupDetailsResponse.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateGroupIdPut**
> CorporateGroupIdPut(ctx, id, body)
Update a group of sub-accounts

This endpoint allows you to update a group of sub-accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Id of the group | 
  **body** | [**Body3**](Body3.md)| Group details to be updated. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateGroupPost**
> InlineResponse201 CorporateGroupPost(ctx, body)
Create a new group of sub-accounts

This endpoint allows to create a group of sub-accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| Group details to be created. | 

### Return type

[**InlineResponse201**](InlineResponse201.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateGroupUnlinkGroupIdSubAccountsPut**
> CorporateGroupUnlinkGroupIdSubAccountsPut(ctx, groupId, body)
Delete sub-account from group

This endpoint allows you to remove a sub-organization from a group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **string**| Id of the group | 
  **body** | [**Body4**](Body4.md)| List of sub-account ids | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CorporateSsoTokenPost**
> GetSsoToken CorporateSsoTokenPost(ctx, ssoTokenRequestCorporate)
Generate SSO token to access admin account

This endpoint generates an SSO token to authenticate and access the admin account using the endpoint https://account-app.brevo.com/account/login/corporate/sso/[token], where [token] will be replaced by the actual token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ssoTokenRequestCorporate** | [**SsoTokenRequestCorporate**](SsoTokenRequestCorporate.md)| User email of admin account | 

### Return type

[**GetSsoToken**](GetSsoToken.md)

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

# **CorporateSubAccountIdApplicationsTogglePut**
> CorporateSubAccountIdApplicationsTogglePut(ctx, id, toggleApplications)
Enable/disable sub-account application(s)

API endpoint for the Corporate owner to enable/disable applications on the sub-account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the sub-account organization (mandatory) | 
  **toggleApplications** | [**SubAccountAppsToggleRequest**](SubAccountAppsToggleRequest.md)| List of applications to activate or deactivate on a sub-account | 

### Return type

 (empty response body)

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

This endpoint will update the sub-account plan. On the Corporate solution new version v2, you can set an unlimited number of credits in your sub-organization. Please pass the value “-1\" to set the consumable in unlimited mode.

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

# **CorporateSubAccountIpAssociatePost**
> interface{} CorporateSubAccountIpAssociatePost(ctx, body)
Associate an IP to sub-accounts

This endpoint allows to associate an IP to sub-accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body1**](Body1.md)| Ip address association details | 

### Return type

[*map[string]interface{}**]

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateSubAccountIpDissociatePut**
> CorporateSubAccountIpDissociatePut(ctx, body)
Dissociate an IP from sub-accounts

This endpoint allows to dissociate an IP from sub-accounts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body2**](Body2.md)| Ip address dissociation details | 

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
Generate SSO token to access sub-account

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

# **CorporateUserInvitationActionEmailPut**
> InlineResponse200 CorporateUserInvitationActionEmailPut(ctx, action, email)
Resend / cancel admin user invitation

This endpoint will allow the user to: - Resend an admin user invitation - Cancel an admin user invitation 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **action** | **string**| Action to be performed (cancel / resend) | 
  **email** | **string**| Email address of the recipient | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CorporateUserRevokeEmailDelete**
> CorporateUserRevokeEmailDelete(ctx, email)
Revoke an admin user

This endpoint allows to revoke/remove an invited member of your Admin account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email of the invited user | 

### Return type

 (empty response body)

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

# **GetCorporateInvitedUsersList**
> GetCorporateInvitedUsersList GetCorporateInvitedUsersList(ctx, )
Get the list of all admin users

This endpoint allows you to list all Admin users of your Admin account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCorporateInvitedUsersList**](GetCorporateInvitedUsersList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCorporateUserPermission**
> GetCorporateUserPermission GetCorporateUserPermission(ctx, email)
Check admin user permissions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email of the invited user | 

### Return type

[**GetCorporateUserPermission**](GetCorporateUserPermission.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubAccountGroups**
> []InlineResponse2001 GetSubAccountGroups(ctx, )
Get the list of groups

This endpoint allows you to list all groups created on your Admin account.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](InlineResponse2001.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteAdminUser**
> InviteAdminUser InviteAdminUser(ctx, sendInvitation)
Send invitation to an admin user

`This endpoint allows you to invite a member to manage the Admin account  Features and their respective permissions are as below:  - `my_plan`:   - \"all\" - `api`:   - \"none\" - `user_management`:   - \"all\" - `app_management` | Not available in ENTv2:   - \"all\"  **Note**: - If `all_features_access: false` then only privileges are required otherwise if `true` then it's assumed that all permissions will be there for the invited admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendInvitation** | [**InviteAdminUser**](InviteAdminUser.md)| Payload to send an invitation | 

### Return type

[**InviteAdminUser**](InviteAdminUser.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

