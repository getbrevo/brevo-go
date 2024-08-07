# CreateWebhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the webhook | [default to null]
**Description** | **string** | Description of the webhook | [optional] [default to null]
**Events** | **[]string** | - Events triggering the webhook. Possible values for **Transactional** type webhook: #### &#x60;sent&#x60; OR &#x60;request&#x60;, &#x60;delivered&#x60;, &#x60;hardBounce&#x60;, &#x60;softBounce&#x60;, &#x60;blocked&#x60;, &#x60;spam&#x60;, &#x60;invalid&#x60;, &#x60;deferred&#x60;, &#x60;click&#x60;, &#x60;opened&#x60;, &#x60;uniqueOpened&#x60; and &#x60;unsubscribed&#x60; - Possible values for **Marketing** type webhook: #### &#x60;spam&#x60;, &#x60;opened&#x60;, &#x60;click&#x60;, &#x60;hardBounce&#x60;, &#x60;softBounce&#x60;, &#x60;unsubscribed&#x60;, &#x60;listAddition&#x60; &amp; &#x60;delivered&#x60; - Possible values for **Inbound** type webhook: #### &#x60;inboundEmailProcessed&#x60;  | [default to null]
**Type_** | **string** | Type of the webhook | [optional] [default to null]
**Domain** | **string** | Inbound domain of webhook, required in case of event type &#x60;inbound&#x60; | [optional] [default to null]
**Batched** | **bool** | To send batched webhooks | [optional] [default to null]
**Auth** | [***GetWebhookAuth**](GetWebhookAuth.md) |  | [optional] [default to null]
**Headers** | [**[]GetWebhookHeaders**](GetWebhookHeaders.md) | Custom headers to be send with webhooks | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


