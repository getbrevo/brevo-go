# Body12

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | **string** | agent ID. It can be found on agent’s page or received &lt;a href&#x3D;\&quot;https://developers.brevo.com/docs/conversations-webhooks\&quot;&gt;from a webhook&lt;/a&gt;. Alternatively, you can use &#x60;agentEmail&#x60; + &#x60;agentName&#x60; + &#x60;receivedFrom&#x60; instead (all 3 fields required). | [optional] [default to null]
**ReceivedFrom** | **string** | mark your messages to distinguish messages created by you from the others. | [optional] [default to null]
**AgentEmail** | **string** | agent email. When sending online pings from a standalone system, it’s hard to maintain a 1-to-1 relationship between the users of both systems. In this case, an agent can be specified by their email address. If there’s no agent with the specified email address in your Brevo organization, a dummy agent will be created automatically. | [optional] [default to null]
**AgentName** | **string** | agent name. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


