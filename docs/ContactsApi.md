# \ContactsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContactToList**](ContactsApi.md#AddContactToList) | **Post** /contacts/lists/{listId}/contacts/add | Add existing contacts to a list
[**CreateAttribute**](ContactsApi.md#CreateAttribute) | **Post** /contacts/attributes/{attributeCategory}/{attributeName} | Create contact attribute
[**CreateContact**](ContactsApi.md#CreateContact) | **Post** /contacts | Create a contact
[**CreateDoiContact**](ContactsApi.md#CreateDoiContact) | **Post** /contacts/doubleOptinConfirmation | Create Contact via DOI (Double-Opt-In) Flow
[**CreateFolder**](ContactsApi.md#CreateFolder) | **Post** /contacts/folders | Create a folder
[**CreateList**](ContactsApi.md#CreateList) | **Post** /contacts/lists | Create a list
[**DeleteAttribute**](ContactsApi.md#DeleteAttribute) | **Delete** /contacts/attributes/{attributeCategory}/{attributeName} | Delete an attribute
[**DeleteContact**](ContactsApi.md#DeleteContact) | **Delete** /contacts/{identifier} | Delete a contact
[**DeleteFolder**](ContactsApi.md#DeleteFolder) | **Delete** /contacts/folders/{folderId} | Delete a folder (and all its lists)
[**DeleteList**](ContactsApi.md#DeleteList) | **Delete** /contacts/lists/{listId} | Delete a list
[**GetAttributes**](ContactsApi.md#GetAttributes) | **Get** /contacts/attributes | List all attributes
[**GetContactInfo**](ContactsApi.md#GetContactInfo) | **Get** /contacts/{identifier} | Get a contact&#39;s details
[**GetContactStats**](ContactsApi.md#GetContactStats) | **Get** /contacts/{identifier}/campaignStats | Get email campaigns&#39; statistics for a contact
[**GetContacts**](ContactsApi.md#GetContacts) | **Get** /contacts | Get all the contacts
[**GetContactsFromList**](ContactsApi.md#GetContactsFromList) | **Get** /contacts/lists/{listId}/contacts | Get contacts in a list
[**GetFolder**](ContactsApi.md#GetFolder) | **Get** /contacts/folders/{folderId} | Returns a folder&#39;s details
[**GetFolderLists**](ContactsApi.md#GetFolderLists) | **Get** /contacts/folders/{folderId}/lists | Get lists in a folder
[**GetFolders**](ContactsApi.md#GetFolders) | **Get** /contacts/folders | Get all folders
[**GetList**](ContactsApi.md#GetList) | **Get** /contacts/lists/{listId} | Get a list&#39;s details
[**GetLists**](ContactsApi.md#GetLists) | **Get** /contacts/lists | Get all the lists
[**GetSegments**](ContactsApi.md#GetSegments) | **Get** /contacts/segments | Get all the Segments
[**ImportContacts**](ContactsApi.md#ImportContacts) | **Post** /contacts/import | Import contacts
[**RemoveContactFromList**](ContactsApi.md#RemoveContactFromList) | **Post** /contacts/lists/{listId}/contacts/remove | Delete a contact from a list
[**RequestContactExport**](ContactsApi.md#RequestContactExport) | **Post** /contacts/export | Export contacts
[**UpdateAttribute**](ContactsApi.md#UpdateAttribute) | **Put** /contacts/attributes/{attributeCategory}/{attributeName} | Update contact attribute
[**UpdateBatchContacts**](ContactsApi.md#UpdateBatchContacts) | **Post** /contacts/batch | Update multiple contacts
[**UpdateContact**](ContactsApi.md#UpdateContact) | **Put** /contacts/{identifier} | Update a contact
[**UpdateFolder**](ContactsApi.md#UpdateFolder) | **Put** /contacts/folders/{folderId} | Update a folder
[**UpdateList**](ContactsApi.md#UpdateList) | **Put** /contacts/lists/{listId} | Update a list


# **AddContactToList**
> PostContactInfo AddContactToList(ctx, listId, contactEmails)
Add existing contacts to a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 
  **contactEmails** | [**AddContactToList**](AddContactToList.md)| Emails addresses OR IDs of the contacts | 

### Return type

[**PostContactInfo**](postContactInfo.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAttribute**
> CreateAttribute(ctx, attributeCategory, attributeName, createAttribute)
Create contact attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the attribute | 
  **createAttribute** | [**CreateAttribute**](CreateAttribute.md)| Values to create an attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContact**
> CreateUpdateContactModel CreateContact(ctx, createContact)
Create a contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContact** | [**CreateContact**](CreateContact.md)| Values to create a contact | 

### Return type

[**CreateUpdateContactModel**](createUpdateContactModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDoiContact**
> CreateDoiContact(ctx, createDoiContact)
Create Contact via DOI (Double-Opt-In) Flow

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createDoiContact** | [**CreateDoiContact**](CreateDoiContact.md)| Values to create the Double opt-in (DOI) contact | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFolder**
> CreateModel CreateFolder(ctx, createFolder)
Create a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createFolder** | [**CreateUpdateFolder**](CreateUpdateFolder.md)| Name of the folder | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateList**
> CreateModel CreateList(ctx, createList)
Create a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createList** | [**CreateList**](CreateList.md)| Values to create a list | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAttribute**
> DeleteAttribute(ctx, attributeCategory, attributeName)
Delete an attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the existing attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContact**
> DeleteContact(ctx, identifier)
Delete a contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Email (urlencoded) OR ID of the contact | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFolder**
> DeleteFolder(ctx, folderId)
Delete a folder (and all its lists)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **int64**| Id of the folder | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteList**
> DeleteList(ctx, listId)
Delete a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAttributes**
> GetAttributes GetAttributes(ctx, )
List all attributes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetAttributes**](getAttributes.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactInfo**
> GetExtendedContactDetails GetContactInfo(ctx, identifier, optional)
Get a contact's details

Along with the contact details, this endpoint will show the statistics of contact for the recent 90 days by default. To fetch the earlier statistics, please use Get contact campaign stats (https://developers.brevo.com/reference/contacts-7#getcontactstats) endpoint with the appropriate date ranges.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Email (urlencoded) OR ID of the contact OR its SMS attribute value | 
 **optional** | ***GetContactInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| **Mandatory if endDate is used.** Starting date (YYYY-MM-DD) of the statistic events specific to campaigns. Must be lower than equal to endDate  | 
 **endDate** | **optional.String**| **Mandatory if startDate is used.** Ending date (YYYY-MM-DD) of the statistic events specific to campaigns. Must be greater than equal to startDate.  | 

### Return type

[**GetExtendedContactDetails**](getExtendedContactDetails.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactStats**
> GetContactCampaignStats GetContactStats(ctx, identifier, optional)
Get email campaigns' statistics for a contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Email (urlencoded) OR ID of the contact | 
 **optional** | ***GetContactStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the statistic events specific to campaigns. Must be lower than equal to endDate | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the statistic events specific to campaigns. Must be greater than equal to startDate. Maximum difference between startDate and endDate should not be greater than 90 days | 

### Return type

[**GetContactCampaignStats**](getContactCampaignStats.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> GetContacts GetContacts(ctx, optional)
Get all the contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document of the page | [default to 0]
 **modifiedSince** | **optional.String**| Filter (urlencoded) the contacts modified after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result. | 
 **createdSince** | **optional.String**| Filter (urlencoded) the contacts created after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result. | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetContacts**](getContacts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsFromList**
> GetContacts GetContactsFromList(ctx, listId, optional)
Get contacts in a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 
 **optional** | ***GetContactsFromListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsFromListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modifiedSince** | **optional.String**| Filter (urlencoded) the contacts modified after a given UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result. | 
 **limit** | **optional.Int64**| Number of documents per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document of the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetContacts**](getContacts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolder**
> GetFolder GetFolder(ctx, folderId)
Returns a folder's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **int64**| id of the folder | 

### Return type

[**GetFolder**](getFolder.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolderLists**
> GetFolderLists GetFolderLists(ctx, folderId, optional)
Get lists in a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **int64**| Id of the folder | 
 **optional** | ***GetFolderListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFolderListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int64**| Number of documents per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document of the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetFolderLists**](getFolderLists.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFolders**
> GetFolders GetFolders(ctx, limit, offset, optional)
Get all folders

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int64**| Number of documents per page | [default to 10]
  **offset** | **int64**| Index of the first document of the page | [default to 0]
 **optional** | ***GetFoldersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFoldersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetFolders**](getFolders.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> GetExtendedList GetList(ctx, listId)
Get a list's details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 

### Return type

[**GetExtendedList**](getExtendedList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLists**
> GetLists GetLists(ctx, optional)
Get all the lists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document of the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetLists**](getLists.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSegments**
> GetSegments GetSegments(ctx, limit, offset, optional)
Get all the Segments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **limit** | **int64**| Number of documents per page | [default to 10]
  **offset** | **int64**| Index of the first document of the page | [default to 0]
 **optional** | ***GetSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSegmentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetSegments**](getSegments.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportContacts**
> CreatedProcessId ImportContacts(ctx, requestContactImport)
Import contacts

It returns the background process ID which on completion calls the notify URL that you have set in the input.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestContactImport** | [**RequestContactImport**](RequestContactImport.md)| Values to import contacts in Brevo. To know more about the expected format, please have a look at &#x60;&#x60;https://help.brevo.com/hc/en-us/articles/209499265-Build-contacts-lists-for-your-email-marketing-campaigns&#x60;&#x60; | 

### Return type

[**CreatedProcessId**](createdProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveContactFromList**
> PostContactInfo RemoveContactFromList(ctx, listId, contactEmails)
Delete a contact from a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 
  **contactEmails** | [**RemoveContactFromList**](RemoveContactFromList.md)| Emails addresses OR IDs of the contacts | 

### Return type

[**PostContactInfo**](postContactInfo.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestContactExport**
> CreatedProcessId RequestContactExport(ctx, requestContactExport)
Export contacts

It returns the background process ID which on completion calls the notify URL that you have set in the input. File will be available in csv.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestContactExport** | [**RequestContactExport**](RequestContactExport.md)| Values to request a contact export | 

### Return type

[**CreatedProcessId**](createdProcessId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAttribute**
> UpdateAttribute(ctx, attributeCategory, attributeName, updateAttribute)
Update contact attribute

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **attributeCategory** | **string**| Category of the attribute | 
  **attributeName** | **string**| Name of the existing attribute | 
  **updateAttribute** | [**UpdateAttribute**](UpdateAttribute.md)| Values to update an attribute | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBatchContacts**
> UpdateBatchContacts(ctx, updateBatchContacts)
Update multiple contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateBatchContacts** | [**UpdateBatchContacts**](UpdateBatchContacts.md)| Values to update multiple contacts | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> UpdateContact(ctx, identifier, updateContact)
Update a contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| Email (urlencoded) OR ID of the contact | 
  **updateContact** | [**UpdateContact**](UpdateContact.md)| Values to update a contact | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolder**
> UpdateFolder(ctx, folderId, updateFolder)
Update a folder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderId** | **int64**| Id of the folder | 
  **updateFolder** | [**CreateUpdateFolder**](CreateUpdateFolder.md)| Name of the folder | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateList**
> UpdateList(ctx, listId, updateList)
Update a list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **listId** | **int64**| Id of the list | 
  **updateList** | [**UpdateList**](UpdateList.md)| Values to update a list | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

