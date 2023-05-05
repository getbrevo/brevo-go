# \ConversationsApi

All URIs are relative to *https://api.brevo.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversationsAgentOnlinePingPost**](ConversationsApi.md#ConversationsAgentOnlinePingPost) | **Post** /conversations/agentOnlinePing | Sets agent’s status to online for 2-3 minutes
[**ConversationsMessagesIdDelete**](ConversationsApi.md#ConversationsMessagesIdDelete) | **Delete** /conversations/messages/{id} | Delete a message sent by an agent
[**ConversationsMessagesIdGet**](ConversationsApi.md#ConversationsMessagesIdGet) | **Get** /conversations/messages/{id} | Get a message
[**ConversationsMessagesIdPut**](ConversationsApi.md#ConversationsMessagesIdPut) | **Put** /conversations/messages/{id} | Update a message sent by an agent
[**ConversationsMessagesPost**](ConversationsApi.md#ConversationsMessagesPost) | **Post** /conversations/messages | Send a message as an agent
[**ConversationsPushedMessagesIdDelete**](ConversationsApi.md#ConversationsPushedMessagesIdDelete) | **Delete** /conversations/pushedMessages/{id} | Delete an automated message
[**ConversationsPushedMessagesIdGet**](ConversationsApi.md#ConversationsPushedMessagesIdGet) | **Get** /conversations/pushedMessages/{id} | Get an automated message
[**ConversationsPushedMessagesIdPut**](ConversationsApi.md#ConversationsPushedMessagesIdPut) | **Put** /conversations/pushedMessages/{id} | Update an automated message
[**ConversationsPushedMessagesPost**](ConversationsApi.md#ConversationsPushedMessagesPost) | **Post** /conversations/pushedMessages | Send an automated message to a visitor


# **ConversationsAgentOnlinePingPost**
> ConversationsAgentOnlinePingPost(ctx, body)
Sets agent’s status to online for 2-3 minutes

We recommend pinging this endpoint every minute for as long as the agent has to be considered online.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body12**](Body12.md)| Agent fields. | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMessagesIdDelete**
> ConversationsMessagesIdDelete(ctx, id)
Delete a message sent by an agent

Only agents’ messages can be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMessagesIdGet**
> ConversationsMessage ConversationsMessagesIdGet(ctx, id)
Get a message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMessagesIdPut**
> ConversationsMessage ConversationsMessagesIdPut(ctx, id, optional)
Update a message sent by an agent

Only agents’ messages can be edited.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message | 
 **optional** | ***ConversationsMessagesIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsMessagesIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of Body9**](Body9.md)|  | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMessagesPost**
> ConversationsMessage ConversationsMessagesPost(ctx, body)
Send a message as an agent

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body8**](Body8.md)| Message fields. | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsPushedMessagesIdDelete**
> ConversationsPushedMessagesIdDelete(ctx, id)
Delete an automated message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message | 

### Return type

 (empty response body)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsPushedMessagesIdGet**
> ConversationsMessage ConversationsPushedMessagesIdGet(ctx, id)
Get an automated message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message sent previously | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsPushedMessagesIdPut**
> ConversationsMessage ConversationsPushedMessagesIdPut(ctx, id, body)
Update an automated message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID of the message | 
  **body** | [**Body11**](Body11.md)|  | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsPushedMessagesPost**
> ConversationsMessage ConversationsPushedMessagesPost(ctx, body)
Send an automated message to a visitor

Example of automated messages: order status, announce new features in your web app, etc.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body10**](Body10.md)|  | 

### Return type

[**ConversationsMessage**](ConversationsMessage.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

