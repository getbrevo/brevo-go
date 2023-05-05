# \TransactionalEmailsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockNewDomain**](TransactionalEmailsApi.md#BlockNewDomain) | **Post** /smtp/blockedDomains | Add a new domain to the list of blocked domains
[**CreateSmtpTemplate**](TransactionalEmailsApi.md#CreateSmtpTemplate) | **Post** /smtp/templates | Create an email template
[**DeleteBlockedDomain**](TransactionalEmailsApi.md#DeleteBlockedDomain) | **Delete** /smtp/blockedDomains/{domain} | Unblock an existing domain from the list of blocked domains
[**DeleteHardbounces**](TransactionalEmailsApi.md#DeleteHardbounces) | **Post** /smtp/deleteHardbounces | Delete hardbounces
[**DeleteScheduledEmailById**](TransactionalEmailsApi.md#DeleteScheduledEmailById) | **Delete** /smtp/email/{identifier} | Delete scheduled emails by batchId or messageId
[**DeleteSmtpTemplate**](TransactionalEmailsApi.md#DeleteSmtpTemplate) | **Delete** /smtp/templates/{templateId} | Delete an inactive email template
[**GetAggregatedSmtpReport**](TransactionalEmailsApi.md#GetAggregatedSmtpReport) | **Get** /smtp/statistics/aggregatedReport | Get your transactional email activity aggregated over a period of time
[**GetBlockedDomains**](TransactionalEmailsApi.md#GetBlockedDomains) | **Get** /smtp/blockedDomains | Get the list of blocked domains
[**GetEmailEventReport**](TransactionalEmailsApi.md#GetEmailEventReport) | **Get** /smtp/statistics/events | Get all your transactional email activity (unaggregated events)
[**GetScheduledEmailByBatchId**](TransactionalEmailsApi.md#GetScheduledEmailByBatchId) | **Get** /smtp/emailStatus/{batchId} | Fetch scheduled emails by batchId
[**GetScheduledEmailByMessageId**](TransactionalEmailsApi.md#GetScheduledEmailByMessageId) | **Get** /smtp/emailStatus/{messageId} | Fetch scheduled email by messageId
[**GetSmtpReport**](TransactionalEmailsApi.md#GetSmtpReport) | **Get** /smtp/statistics/reports | Get your transactional email activity aggregated per day
[**GetSmtpTemplate**](TransactionalEmailsApi.md#GetSmtpTemplate) | **Get** /smtp/templates/{templateId} | Returns the template information
[**GetSmtpTemplates**](TransactionalEmailsApi.md#GetSmtpTemplates) | **Get** /smtp/templates | Get the list of email templates
[**GetTransacBlockedContacts**](TransactionalEmailsApi.md#GetTransacBlockedContacts) | **Get** /smtp/blockedContacts | Get the list of blocked or unsubscribed transactional contacts
[**GetTransacEmailContent**](TransactionalEmailsApi.md#GetTransacEmailContent) | **Get** /smtp/emails/{uuid} | Get the personalized content of a sent transactional email
[**GetTransacEmailsList**](TransactionalEmailsApi.md#GetTransacEmailsList) | **Get** /smtp/emails | Get the list of transactional emails on the basis of allowed filters
[**SendTestTemplate**](TransactionalEmailsApi.md#SendTestTemplate) | **Post** /smtp/templates/{templateId}/sendTest | Send a template to your test list
[**SendTransacEmail**](TransactionalEmailsApi.md#SendTransacEmail) | **Post** /smtp/email | Send a transactional email
[**SmtpBlockedContactsEmailDelete**](TransactionalEmailsApi.md#SmtpBlockedContactsEmailDelete) | **Delete** /smtp/blockedContacts/{email} | Unblock or resubscribe a transactional contact
[**SmtpLogMessageIdDelete**](TransactionalEmailsApi.md#SmtpLogMessageIdDelete) | **Delete** /smtp/log/{messageId} | Delete an SMTP transactional log
[**UpdateSmtpTemplate**](TransactionalEmailsApi.md#UpdateSmtpTemplate) | **Put** /smtp/templates/{templateId} | Update an email template


# **BlockNewDomain**
> BlockNewDomain(ctx, blockDomain)
Add a new domain to the list of blocked domains

Blocks a new domain in order to avoid messages being sent to the same

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockDomain** | [**BlockDomain**](BlockDomain.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSmtpTemplate**
> CreateModel CreateSmtpTemplate(ctx, smtpTemplate)
Create an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smtpTemplate** | [**CreateSmtpTemplate**](CreateSmtpTemplate.md)| values to update in transactional email template | 

### Return type

[**CreateModel**](createModel.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBlockedDomain**
> DeleteBlockedDomain(ctx, domain)
Unblock an existing domain from the list of blocked domains

Unblocks an existing domain from the list of blocked domains

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| The name of the domain to be deleted | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHardbounces**
> DeleteHardbounces(ctx, optional)
Delete hardbounces

Delete hardbounces. To use carefully (e.g. in case of temporary ISP failures)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteHardbouncesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteHardbouncesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteHardbounces** | [**optional.Interface of DeleteHardbounces**](DeleteHardbounces.md)| values to delete hardbounces | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledEmailById**
> DeleteScheduledEmailById(ctx, identifier)
Delete scheduled emails by batchId or messageId

Delete scheduled batch of emails by batchId or single scheduled email by messageId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identifier** | **string**| The &#x60;batchId&#x60; of scheduled emails batch (Should be a valid UUIDv4) or the &#x60;messageId&#x60; of scheduled email. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSmtpTemplate**
> DeleteSmtpTemplate(ctx, templateId)
Delete an inactive email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the template | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregatedSmtpReport**
> GetAggregatedReport GetAggregatedSmtpReport(ctx, optional)
Get your transactional email activity aggregated over a period of time

This endpoint will show the aggregated stats for past 90 days by default if `startDate` and `endDate` OR `days` is not passed. The date range can not exceed 90 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAggregatedSmtpReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAggregatedSmtpReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | 
 **tag** | **optional.String**| Tag of the emails | 

### Return type

[**GetAggregatedReport**](getAggregatedReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockedDomains**
> GetBlockedDomains GetBlockedDomains(ctx, )
Get the list of blocked domains

Get the list of blocked domains

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBlockedDomains**](getBlockedDomains.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmailEventReport**
> GetEmailEventReport GetEmailEventReport(ctx, optional)
Get all your transactional email activity (unaggregated events)

This endpoint will show the aggregated stats for past 30 days by default if `startDate` and `endDate` OR `days` is not passed. The date range can not exceed 90 days

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetEmailEventReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEmailEventReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number limitation for the result returned | [default to 2500]
 **offset** | **optional.Int64**| Beginning point in the list to retrieve from. | [default to 0]
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD). Must be lower than equal to endDate | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD). Must be greater than equal to startDate | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | 
 **email** | **optional.String**| Filter the report for a specific email addresses | 
 **event** | **optional.String**| Filter the report for a specific event type | 
 **tags** | **optional.String**| Filter the report for tags (serialized and urlencoded array) | 
 **messageId** | **optional.String**| Filter on a specific message id | 
 **templateId** | **optional.Int64**| Filter on a specific template id | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetEmailEventReport**](getEmailEventReport.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledEmailByBatchId**
> GetScheduledEmailByBatchId GetScheduledEmailByBatchId(ctx, batchId, optional)
Fetch scheduled emails by batchId

Fetch scheduled batch of emails by batchId (Can retrieve data upto 30 days old)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **batchId** | **string**| The batchId of scheduled emails batch (Should be a valid UUIDv4) | 
 **optional** | ***GetScheduledEmailByBatchIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetScheduledEmailByBatchIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| Mandatory if &#x60;endDate&#x60; is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Can be maximum 30 days older tha current date. | 
 **endDate** | **optional.String**| Mandatory if &#x60;startDate&#x60; is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **status** | **optional.String**| Filter the records by &#x60;status&#x60; of the scheduled email batch or message. | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 100]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]

### Return type

[**GetScheduledEmailByBatchId**](getScheduledEmailByBatchId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledEmailByMessageId**
> GetScheduledEmailByMessageId GetScheduledEmailByMessageId(ctx, messageId, optional)
Fetch scheduled email by messageId

Fetch scheduled email by messageId (Can retrieve data upto 30 days old)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| The messageId of scheduled email | 
 **optional** | ***GetScheduledEmailByMessageIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetScheduledEmailByMessageIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Can be maximum 30 days older tha current date. | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 

### Return type

[**GetScheduledEmailByMessageId**](getScheduledEmailByMessageId.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpReport**
> GetReports GetSmtpReport(ctx, optional)
Get your transactional email activity aggregated per day

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSmtpReportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSmtpReportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 10]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date of the report (YYYY-MM-DD) | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date of the report (YYYY-MM-DD) | 
 **days** | **optional.Int64**| Number of days in the past including today (positive integer). Not compatible with &#39;startDate&#39; and &#39;endDate&#39; | 
 **tag** | **optional.String**| Tag of the emails | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetReports**](getReports.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplate**
> GetSmtpTemplateOverview GetSmtpTemplate(ctx, templateId)
Returns the template information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the template | 

### Return type

[**GetSmtpTemplateOverview**](getSmtpTemplateOverview.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSmtpTemplates**
> GetSmtpTemplates GetSmtpTemplates(ctx, optional)
Get the list of email templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSmtpTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSmtpTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateStatus** | **optional.Bool**| Filter on the status of the template. Active &#x3D; true, inactive &#x3D; false | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetSmtpTemplates**](getSmtpTemplates.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacBlockedContacts**
> GetTransacBlockedContacts GetTransacBlockedContacts(ctx, optional)
Get the list of blocked or unsubscribed transactional contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransacBlockedContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTransacBlockedContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the blocked or unsubscribed contacts | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the blocked or unsubscribed contacts | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 50]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **senders** | [**optional.Interface of []string**](string.md)| Comma separated list of emails of the senders from which contacts are blocked or unsubscribed | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]

### Return type

[**GetTransacBlockedContacts**](getTransacBlockedContacts.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacEmailContent**
> GetTransacEmailContent GetTransacEmailContent(ctx, uuid)
Get the personalized content of a sent transactional email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| Unique id of the transactional email that has been sent to a particular contact | 

### Return type

[**GetTransacEmailContent**](getTransacEmailContent.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTransacEmailsList**
> GetTransacEmailsList GetTransacEmailsList(ctx, optional)
Get the list of transactional emails on the basis of allowed filters

This endpoint will show the list of emails for past 30 days by default. To retrieve emails before that time, please pass startDate and endDate in query filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTransacEmailsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTransacEmailsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Mandatory if templateId and messageId are not passed in query filters. Email address to which transactional email has been sent. | 
 **templateId** | **optional.Int64**| Mandatory if email and messageId are not passed in query filters. Id of the template that was used to compose transactional email. | 
 **messageId** | **optional.String**| Mandatory if templateId and email are not passed in query filters. Message ID of the transactional email sent. | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation. Default order is **descending** if &#x60;sort&#x60; is not passed | [default to desc]
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 500]
 **offset** | **optional.Int64**| Index of the first document in the page | [default to 0]

### Return type

[**GetTransacEmailsList**](getTransacEmailsList.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTestTemplate**
> SendTestTemplate(ctx, templateId, sendTestEmail)
Send a template to your test list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| Id of the template | 
  **sendTestEmail** | [**SendTestEmail**](SendTestEmail.md)|  | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTransacEmail**
> CreateSmtpEmail SendTransacEmail(ctx, sendSmtpEmail)
Send a transactional email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendSmtpEmail** | [**SendSmtpEmail**](SendSmtpEmail.md)| Values to send a transactional email | 

### Return type

[**CreateSmtpEmail**](createSmtpEmail.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmtpBlockedContactsEmailDelete**
> SmtpBlockedContactsEmailDelete(ctx, email)
Unblock or resubscribe a transactional contact

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| contact email (urlencoded) to unblock. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SmtpLogMessageIdDelete**
> SmtpLogMessageIdDelete(ctx, messageId)
Delete an SMTP transactional log

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **messageId** | **string**| MessageId of the transactional log to delete | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmtpTemplate**
> UpdateSmtpTemplate(ctx, templateId, smtpTemplate)
Update an email template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **int64**| id of the template | 
  **smtpTemplate** | [**UpdateSmtpTemplate**](UpdateSmtpTemplate.md)| values to update in transactional email template | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

