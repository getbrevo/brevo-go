# ConversationsMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Message ID. It can be used for further manipulations with the message. | [optional] [default to null]
**Type_** | **string** | &#x60;\&quot;agent\&quot;&#x60; for agents’ messages, &#x60;\&quot;visitor\&quot;&#x60; for visitors’ messages. | [optional] [default to null]
**Text** | **string** | Message text or name of the attached file | [optional] [default to null]
**VisitorId** | **string** | visitor’s ID | [optional] [default to null]
**AgentId** | **string** | ID of the agent on whose behalf the message was sent (only in messages sent by an agent). | [optional] [default to null]
**AgentName** | **string** | Agent’s name as displayed to the visitor. Only in the messages sent by an agent. | [optional] [default to null]
**CreatedAt** | **int64** | Timestamp in milliseconds. | [optional] [default to null]
**IsPushed** | **bool** | &#x60;true&#x60; for pushed messages | [optional] [default to null]
**ReceivedFrom** | **string** | In two-way integrations, messages sent via REST API can be marked with receivedFrom property and then filtered out when received in a webhook to avoid infinite loop. | [optional] [default to null]
**File** | [***ConversationsMessageFile**](ConversationsMessage_file.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


